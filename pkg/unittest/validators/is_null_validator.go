package validators

import (
	"github.com/lrills/helm-unittest/internal/common"
	"github.com/lrills/helm-unittest/pkg/unittest/valueutils"
)

// IsNullValidator validate value of Path id kind
type IsNullValidator struct {
	Path string
}

func (v IsNullValidator) failInfo(actual interface{}, index int, not bool) []string {
	return splitInfof(
		setFailFormat(not, true, false, false, " to be null, got"),
		index,
		v.Path,
		common.TrustedMarshalYAML(actual),
	)
}

// Validate implement Validatable
func (v IsNullValidator) Validate(context *ValidateContext) (bool, []string) {
	manifests, err := context.getManifests()
	if err != nil {
		return false, splitInfof(errorFormat, -1, err.Error())
	}

	validateSuccess := false
	validateErrors := make([]string, 0)

	for idx, manifest := range manifests {
		actual, err := valueutils.GetValueOfSetPath(manifest, v.Path)

		//#### Option 1: Path must exist excepting the last part. but "xxx:" in YAML is not valid in json ####
		// if err != nil {
		// 	validateSuccess = false
		// 	errorMessage := splitInfof(errorFormat, idx, err.Error())
		// 	validateErrors = append(validateErrors, errorMessage...)
		// 	continue
		// }

		// if actual == nil == context.Negative {
		// 	validateSuccess = false
		// 	errorMessage := v.failInfo(actual, idx, context.Negative)
		// 	validateErrors = append(validateErrors, errorMessage...)
		// 	continue
		// }
		

		//#### Option 2: Assume any invalid path as NULL ####
		if context.Negative {
			if err != nil || actual == nil {
				validateSuccess = false
				errorMessage := v.failInfo(actual, idx, context.Negative)
				validateErrors = append(validateErrors, errorMessage...)
				continue
			}
		} else {
			if err == nil && actual != nil {
				validateSuccess = false
				errorMessage := v.failInfo(actual, idx, context.Negative)
				validateErrors = append(validateErrors, errorMessage...)
				continue
			}
		}

		validateSuccess = determineSuccess(idx, validateSuccess, true)
	}

	return validateSuccess, validateErrors
}
