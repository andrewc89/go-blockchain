package blockchain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryParameters_String_Multiple(t *testing.T) {
	qpMap := make(map[string]string)
	qpMap["foo"] = "bar"
	qpMap["hello"] = "world"
	qps := QueryParameters(qpMap)
	expected := "foo=bar&hello=world"
	actual := qps.String()
	assert.Equal(t, expected, actual)
}

func TestQueryParameters_String_Single(t *testing.T) {
	qpMap := make(map[string]string)
	qpMap["foo"] = "bar"
	qps := QueryParameters(qpMap)
	expected := "foo=bar"
	actual := qps.String()
	assert.Equal(t, expected, actual)
}

func TestQueryParameters_String_None(t *testing.T) {
	qps := QueryParameters(make(map[string]string))
	actual := qps.String()
	assert.Equal(t, "", actual)
}

func TestEndpoint_String_WithQP(t *testing.T) {
	qpMap := make(map[string]string)
	qpMap["foo"] = "bar"
	qps := QueryParameters(qpMap)
	expected := "http://example.com?foo=bar"
	endpoint := Endpoint{"http://example.com", qps}
	actual := endpoint.String()
	assert.Equal(t, expected, actual)
}

func TestEndpoint_String_WithoutQP(t *testing.T) {
	expected := "http://example.com"
	endpoint := Endpoint{"http://example.com", make(map[string]string)}
	actual := endpoint.String()
	assert.Equal(t, expected, actual)
}
