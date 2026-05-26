package extensions

import (
	"reflect"

	"github.com/envoyproxy/protoc-gen-validate/validate"
	"github.com/pseudomuto/protoc-gen-doc/extensions"
)

// ValidateRule represents a single validator rule from the (validate.rules) method option extension.
type ValidateRule struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

// ValidateExtension contains the rules set by the (validate.rules) method option extension.
type ValidateExtension struct {
	*validate.FieldRules
	rules []ValidateRule // memoized so that we don't have to use reflection more than we need.
}

// MarshalJSON implements the json.Marshaler interface.
func (v ValidateExtension) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Rules returns the set of rules for this extension.
}

func (v ValidateExtension) Rules() []ValidateRule { _ = "STUB: not implemented"; return nil }

func flattenRules(prefix string, vv reflect.Value) (rules []ValidateRule) {
	_ = "STUB: not implemented"
	return nil
}

func init() {
	extensions.SetTransformer("validate.rules", func(payload interface{}) interface{} {
		rules, ok := payload.(*validate.FieldRules)
		if !ok {
			return nil
		}
		return ValidateExtension{FieldRules: rules}
	})
}
