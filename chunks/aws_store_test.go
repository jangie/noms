package chunks

import (
	"bytes"
	"errors"
	"io/ioutil"
	"testing"

	"github.com/attic-labs/noms/ref"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/stretchr/testify/suite"
)

func TestAWSStoreTestSuite(t *testing.T) {
	suite.Run(t, &AWSStoreTestSuite{})
}

type AWSStoreTestSuite struct {
	suite.Suite
	store AWSStore
}

func (suite *AWSStoreTestSuite) SetupTest() {
	suite.store = AWSStore{
		"bucket",
		"table",
		&mockS3{},
		nil,
	}
}

type mockS3 map[string][]byte

func (m *mockS3) GetObject(input *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	result, ok := (*m)[*input.Key]
	if !ok {
		return nil, errors.New("not here")
	}

	return &s3.GetObjectOutput{
		Body: ioutil.NopCloser(bytes.NewReader(result)),
	}, nil
}

func (m *mockS3) PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	bytes, _ := ioutil.ReadAll(input.Body)
	(*m)[*input.Key] = bytes
	return nil, nil
}

func (suite *AWSStoreTestSuite) TestAWSStorePut() {
	input := "abc"

	w := suite.store.Put()
	_, err := w.Write([]byte(input))
	suite.NoError(err)

	r1, err := w.Ref()
	suite.NoError(err)

	// See http://www.di-mgt.com.au/sha_testvectors.html
	suite.Equal("sha1-a9993e364706816aba3e25717850c26c9cd0d89d", r1.String())

	// And reading it via the API should work...
	assertInputInStore(input, r1, suite.store, suite.Assert())

	// Reading a non-existing ref fails
	hash := ref.NewHash()
	hash.Write([]byte("Non-existent"))
	_, err = suite.store.Get(ref.FromHash(hash))
	suite.Error(err)
}

func (suite *AWSStoreTestSuite) TestAWSStorePutRefAfterClose() {
	input := "abc"

	w := suite.store.Put()
	_, err := w.Write([]byte(input))
	suite.NoError(err)

	suite.NoError(w.Close())
	r1, err := w.Ref()
	suite.NoError(err)

	// See http://www.di-mgt.com.au/sha_testvectors.html
	suite.Equal("sha1-a9993e364706816aba3e25717850c26c9cd0d89d", r1.String())

	// And reading it via the API should work...
	assertInputInStore(input, r1, suite.store, suite.Assert())
}

func (suite *AWSStoreTestSuite) TestAWSStorePutMultiRef() {
	input := "abc"

	w := suite.store.Put()
	_, err := w.Write([]byte(input))
	suite.NoError(err)

	_, _ = w.Ref()
	r1, err := w.Ref()
	suite.NoError(err)

	// See http://www.di-mgt.com.au/sha_testvectors.html
	suite.Equal("sha1-a9993e364706816aba3e25717850c26c9cd0d89d", r1.String())

	// And reading it via the API should work...
	assertInputInStore(input, r1, suite.store, suite.Assert())
}

type mockAWSError string

func (m mockAWSError) Error() string   { return string(m) }
func (m mockAWSError) Code() string    { return string(m) }
func (m mockAWSError) Message() string { return string(m) }
func (m mockAWSError) OrigErr() error  { return nil }

type mockDDB string

func (m *mockDDB) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	if *m == "" {
		return &dynamodb.GetItemOutput{
			Item: map[string]*dynamodb.AttributeValue{},
		}, nil
	}

	return &dynamodb.GetItemOutput{
		Item: map[string]*dynamodb.AttributeValue{
			"name":    {S: aws.String("root")},
			"hashRef": {S: aws.String(string(*m))},
		},
	}, nil
}

func (m *mockDDB) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	initial := *(input.ConditionExpression) == "attribute_not_exists(hashRef)"

	if (initial && string(*m) != "") || (!initial && string(*m) != *(input.ExpressionAttributeValues[":prev"].S)) {
		return nil, mockAWSError("ConditionalCheckFailedException")
	}

	*m = mockDDB(*(input.Item["hashRef"].S))
	return &dynamodb.PutItemOutput{}, nil
}

func (suite *AWSStoreTestSuite) TestAWSStoreRoot() {
	m := mockDDB("")

	suite.store = AWSStore{
		"bucket",
		"table",
		nil,
		&m,
	}

	oldRoot := suite.store.Root()
	suite.Equal(oldRoot, ref.Ref{})

	bogusRoot, err := ref.Parse("sha1-81c870618113ba29b6f2b396ea3a69c6f1d626c5") // sha1("Bogus, Dude")
	suite.NoError(err)
	newRoot, err := ref.Parse("sha1-907d14fb3af2b0d4f18c2d46abe8aedce17367bd") // sha1("Hello, World")
	suite.NoError(err)

	// Try to update root with bogus oldRoot
	result := suite.store.UpdateRoot(newRoot, bogusRoot)
	suite.False(result)
	suite.Equal(ref.Ref{}, suite.store.Root())

	// No do a valid update
	result = suite.store.UpdateRoot(newRoot, oldRoot)
	suite.True(result)
	suite.Equal(suite.store.Root(), newRoot)

	// Now that there is a valid root, try to start a new lineage
	result = suite.store.UpdateRoot(bogusRoot, ref.Ref{})
	suite.False(result)
	suite.Equal(suite.store.Root(), newRoot)
}