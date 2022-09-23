package validators_test

import (
	"testing"

	. "github.com/lrills/helm-unittest/pkg/unittest/validators"
	"github.com/stretchr/testify/assert"
)

func TestIsNullValidatorWhenOk(t *testing.T) {
	doc := "a:"
	manifest := makeManifest(doc)

	v := IsNullValidator{"a"}
	pass, diff := v.Validate(&ValidateContext{
		Docs: []map[string]interface{}{manifest},
	})
	assert.True(t, pass)
	assert.Equal(t, []string{}, diff)
}

func TestIsNullValidatorWhenNegativeAndOk(t *testing.T) {
	doc := "a: 0"
	manifest := makeManifest(doc)

	v := IsNullValidator{"a"}
	pass, diff := v.Validate(&ValidateContext{
		Docs:     []map[string]interface{}{manifest},
		Negative: true,
	})

	assert.True(t, pass)
	assert.Equal(t, []string{}, diff)
}

func TestIsNullValidatorWhenFail(t *testing.T) {
	doc := "a: A"
	manifest := makeManifest(doc)

	v := IsNullValidator{"a"}
	pass, diff := v.Validate(&ValidateContext{
		Docs: []map[string]interface{}{manifest},
	})

	assert.False(t, pass)
	assert.Equal(t, []string{
		"DocumentIndex:	0",
		"Path:	a",
		"Expected to be null, got:",
		"	A",
	}, diff)
}

func TestIsNullValidatorWhenNegativeAndFail(t *testing.T) {
	doc := "a:"
	manifest := makeManifest(doc)

	v := IsNullValidator{"a"}
	pass, diff := v.Validate(&ValidateContext{
		Docs:     []map[string]interface{}{manifest},
		Negative: true,
	})

	assert.False(t, pass)
	assert.Equal(t, []string{
		"DocumentIndex:	0",
		"Path:	a",
		"Expected NOT to be null, got:",
		"	null",
	}, diff)
}

func TestIsNullValidatorWhenInvalidIndex(t *testing.T) {
	doc := "a:"
	manifest := makeManifest(doc)

	validator := IsNullValidator{"a"}
	pass, diff := validator.Validate(&ValidateContext{
		Docs:  []map[string]interface{}{manifest},
		Index: 2,
	})

	assert.False(t, pass)
	assert.Equal(t, []string{
		"Error:",
		"	documentIndex 2 out of range",
	}, diff)
}

func TestIsNullValidatorWhenInvalidPath(t *testing.T) {
	doc := "x:"
	manifest := makeManifest(doc)

	validator := IsNullValidator{"x.b"}
	pass, _ := validator.Validate(&ValidateContext{
		Docs: []map[string]interface{}{manifest},
	})

	// After changed to jsonpath, all cases of invalid path including missing key will pass isNull validator
	assert.True(t, pass) 
	// assert.False(t, pass)
	// assert.Equal(t, []string{
	// 	"DocumentIndex:	0",
	// 	"Error:",
	// 	"	can't get [\"b\"] from a non map type:",
	// 	"	null",
	// }, diff)
}
