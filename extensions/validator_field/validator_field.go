package extensions

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	validator "github.com/mwitkow/go-proto-validators"
	"github.com/pseudomuto/protoc-gen-doc/extensions"
)

func init() {
	// NOTE: mwitkow/go-proto-validators uses gogo/profobuf/proto and therefore
	// only registers the extension under gogo. We need to register it under
	// golang/protobuf/proto with the same properties, except using the
	// golang/protobuf FieldOptions descriptor.
	proto.RegisterExtension(&proto.ExtensionDesc{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: validator.E_Field.ExtensionType,
		Field:         validator.E_Field.Field,
		Name:          validator.E_Field.Name,
		Tag:           validator.E_Field.Tag,
		Filename:      validator.E_Field.Filename,
	})
}

// ValidatorRule represents a single validator rule from the (validator.field) method option extension.
type ValidatorRule struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

// ValidatorExtension contains the rules set by the (validator.field) method option extension.
type ValidatorExtension struct {
	*validator.FieldValidator
	rules []ValidatorRule // memoized so that we don't have to use reflection more than we need.
}

// MarshalJSON implements the json.Marshaler interface.
func (v ValidatorExtension) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Rules returns all active rules
}

func (v ValidatorExtension) Rules() []ValidatorRule { _ = "STUB: not implemented"; return nil }

func init() {
	extensions.SetTransformer("validator.field", func(payload interface{}) interface{} {
		validator, ok := payload.(*validator.FieldValidator)
		if !ok {
			return nil
		}
		return ValidatorExtension{FieldValidator: validator}
	})
}
