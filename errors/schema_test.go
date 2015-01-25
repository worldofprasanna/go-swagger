package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSchemaErrors(t *testing.T) {
	err := InvalidType("confirmed", "query", "boolean", nil)
	assert.Error(t, err)
	assert.Equal(t, 422, err.Code())
	assert.Equal(t, "confirmed in query must be of type boolean", err.Error())

	err = InvalidType("confirmed", "query", "boolean", "hello")
	assert.Error(t, err)
	assert.Equal(t, 422, err.Code())
	assert.Equal(t, "confirmed in query must be of type boolean: \"hello\"", err.Error())

	err = InvalidType("confirmed", "query", "boolean", errors.New("hello"))
	assert.Error(t, err)
	assert.Equal(t, 422, err.Code())
	assert.Equal(t, "confirmed in query must be of type boolean, because: hello", err.Error())

	err = DuplicateItems("uniques", "query")
	assert.Error(t, err)
	assert.Equal(t, 422, err.Code())
	assert.Equal(t, "uniques in query shouldn't contain duplicates", err.Error())

	err = TooManyItems("something", "query", 5)
	assert.Error(t, err)
	assert.Equal(t, 422, err.Code())
	assert.Equal(t, "something in query should have at most 5 items", err.Error())

	err = TooFewItems("something", "query", 5)
	assert.Error(t, err)
	assert.Equal(t, 422, err.Code())
	assert.Equal(t, "something in query should have at least 5 items", err.Error())

	err = ExceedsMaximum("something", "query", 5, false)
	assert.Error(t, err)
	assert.Equal(t, 422, err.Code())
	assert.Equal(t, "something in query should be less than or equal to 5", err.Error())

	err = ExceedsMaximum("something", "query", 5, true)
	assert.Error(t, err)
	assert.Equal(t, 422, err.Code())
	assert.Equal(t, "something in query should be less than 5", err.Error())

	err = ExceedsMinimum("something", "query", 5, false)
	assert.Error(t, err)
	assert.Equal(t, 422, err.Code())
	assert.Equal(t, "something in query should be greater than or equal to 5", err.Error())

	err = ExceedsMinimum("something", "query", 5, true)
	assert.Error(t, err)
	assert.Equal(t, 422, err.Code())
	assert.Equal(t, "something in query should be greater than 5", err.Error())

	err = NotMultipleOf("something", "query", 5)
	assert.Error(t, err)
	assert.Equal(t, 422, err.Code())
	assert.Equal(t, "something in query should be a multiple of 5", err.Error())

	err = EnumFail("something", "query", "yada", []interface{}{"hello", "world"})
	assert.Error(t, err)
	assert.Equal(t, 422, err.Code())
	assert.Equal(t, "something in query should be one of [hello world]", err.Error())

	err = Required("something", "query")
	assert.Error(t, err)
	assert.Equal(t, 422, err.Code())
	assert.Equal(t, "something in query is required", err.Error())

	err = TooLong("something", "query", 5)
	assert.Error(t, err)
	assert.Equal(t, 422, err.Code())
	assert.Equal(t, "something in query should be at most 5 chars long", err.Error())

	err = TooShort("something", "query", 5)
	assert.Error(t, err)
	assert.Equal(t, 422, err.Code())
	assert.Equal(t, "something in query should be at least 5 chars long", err.Error())

	err = FailedPattern("something", "query", "\\d+")
	assert.Error(t, err)
	assert.Equal(t, 422, err.Code())
	assert.Equal(t, "something in query should match '\\d+'", err.Error())

	err = InvalidTypeName("something")
	assert.Error(t, err)
	assert.Equal(t, 422, err.Code())
	assert.Equal(t, "something is an invalid type name", err.Error())

	err = InvalidCollectionFormat("something", "query", "yada")
	assert.Error(t, err)
	assert.Equal(t, 422, err.Code())
	assert.Equal(t, "the collection format \"yada\" is not supported for the query param \"something\"", err.Error())

	err = CompositeValidationError()
	assert.Error(t, err)
	assert.Equal(t, 422, err.Code())
	assert.Equal(t, "validation failure list", err.Error())
}
