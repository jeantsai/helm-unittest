package validators_test

import (
	"testing"

	. "github.com/lrills/helm-unittest/pkg/unittest/validators"
	"github.com/stretchr/testify/assert"
)

func TestHasDocumentsValidatorOk(t *testing.T) {
	data := map[string]interface{}{}

	validator := HasDocumentsValidator{2}
	pass, diff := validator.Validate(&ValidateContext{
		Docs: []map[string]interface{}{data, data},
	})

	assert.True(t, pass)
	assert.Equal(t, []string{}, diff)
}

func TestHasDocumentsValidatorWhenNegativeAndOk(t *testing.T) {
	data := map[string]interface{}{}

	validator := HasDocumentsValidator{2}
	pass, diff := validator.Validate(&ValidateContext{
		Docs:     []map[string]interface{}{data},
		Negative: true,
	})

	assert.True(t, pass)
	assert.Equal(t, []string{}, diff)
}

func TestHasDocumentsValidatorWhenFail(t *testing.T) {
	data := map[string]interface{}{}

	validator := HasDocumentsValidator{1}
	pass, diff := validator.Validate(&ValidateContext{
		Docs: []map[string]interface{}{data, data},
	})

	assert.False(t, pass)
	assert.Equal(t, []string{
		"Expected documents count to be:",
		"	1",
		"Actual:",
		"	2",
	}, diff)
}

func TestHasDocumentsValidatorWhenNegativeAndFail(t *testing.T) {
	data := map[string]interface{}{}

	validator := HasDocumentsValidator{2}
	pass, diff := validator.Validate(&ValidateContext{
		Docs:     []map[string]interface{}{data, data},
		Negative: true,
	})

	assert.False(t, pass)
	assert.Equal(t, []string{
		"Expected NOT documents count to be:",
		"	2",
	}, diff)
}
