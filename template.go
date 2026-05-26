package gendoc

import (
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/pseudomuto/protokit"
)

// Template is a type for encapsulating all the parsed files, messages, fields, enums, services, extensions, etc. into
// an object that will be supplied to a go template.
type Template struct {
	// The files that were parsed
	Files []*File `json:"files"`
	// Details about the scalar values and their respective types in supported languages.
	Scalars []*ScalarValue `json:"scalarValueTypes"`
}

// NewTemplate creates a Template object from a set of descriptors.
func NewTemplate(descs []*protokit.FileDescriptor, pluginOptions *PluginOptions) *Template {
	_ = "STUB: not implemented"
	return nil
}

// Recursively add nested types from messages

func makeScalars() []*ScalarValue { _ = "STUB: not implemented"; return nil }

func mergeOptions(opts ...map[string]interface{}) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func camelCase(s string) string { _ = "STUB: not implemented"; return "" }

// CommonOptions are options common to all descriptor types.
type commonOptions interface {
	GetDeprecated() bool
}

func extractOptions(opts commonOptions) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

// File wraps all the relevant parsed info about a proto file. File objects guarantee that their top-level enums,
// extensions, messages, and services are sorted alphabetically based on their "long name". Other values (enum values,
// fields, service methods) will be in the order that they're defined within their respective proto files.
//
// In the case of proto3 files, HasExtensions will always be false, and Extensions will be empty.
type File struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Package     string `json:"package"`

	HasEnums      bool `json:"hasEnums"`
	HasExtensions bool `json:"hasExtensions"`
	HasMessages   bool `json:"hasMessages"`
	HasServices   bool `json:"hasServices"`

	Enums      orderedEnums      `json:"enums"`
	Extensions orderedExtensions `json:"extensions"`
	Messages   orderedMessages   `json:"messages"`
	Services   orderedServices   `json:"services"`

	Options map[string]interface{} `json:"options,omitempty"`
}

// Option returns the named option.
func (f File) Option(name string) interface{} { _ = "STUB: not implemented"; return nil }

// FileExtension contains details about top-level extensions within a proto(2) file.
type FileExtension struct {
	Name               string `json:"name"`
	LongName           string `json:"longName"`
	FullName           string `json:"fullName"`
	Description        string `json:"description"`
	Label              string `json:"label"`
	Type               string `json:"type"`
	LongType           string `json:"longType"`
	FullType           string `json:"fullType"`
	Number             int    `json:"number"`
	DefaultValue       string `json:"defaultValue"`
	ContainingType     string `json:"containingType"`
	ContainingLongType string `json:"containingLongType"`
	ContainingFullType string `json:"containingFullType"`
}

// Message contains details about a protobuf message.
//
// In the case of proto3 files, HasExtensions will always be false, and Extensions will be empty.
type Message struct {
	Name        string `json:"name"`
	LongName    string `json:"longName"`
	FullName    string `json:"fullName"`
	Description string `json:"description"`

	HasExtensions bool `json:"hasExtensions"`
	HasFields     bool `json:"hasFields"`
	HasOneofs     bool `json:"hasOneofs"`

	Extensions []*MessageExtension `json:"extensions"`
	Fields     []*MessageField     `json:"fields"`

	Options map[string]interface{} `json:"options,omitempty"`
}

// Option returns the named option.
func (m Message) Option(name string) interface{} { _ = "STUB: not implemented"; return nil }

// FieldOptions returns all options that are set on the fields in this message.
func (m Message) FieldOptions() []string { _ = "STUB: not implemented"; return nil }

// FieldsWithOption returns all fields that have the given option set.
// If no single value has the option set, this returns nil.
func (m Message) FieldsWithOption(optionName string) []*MessageField {
	_ = "STUB: not implemented"
	return nil
}

// MessageField contains details about an individual field within a message.
//
// In the case of proto3 files, DefaultValue will always be empty. Similarly, label will be empty unless the field is
// repeated (in which case it'll be "repeated").
type MessageField struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Label        string `json:"label"`
	Type         string `json:"type"`
	LongType     string `json:"longType"`
	FullType     string `json:"fullType"`
	IsMap        bool   `json:"ismap"`
	IsOneof      bool   `json:"isoneof"`
	OneofDecl    string `json:"oneofdecl"`
	DefaultValue string `json:"defaultValue"`

	Options map[string]interface{} `json:"options,omitempty"`
}

// Option returns the named option.
func (f MessageField) Option(name string) interface{} { _ = "STUB: not implemented"; return nil }

// MessageExtension contains details about message-scoped extensions in proto(2) files.
type MessageExtension struct {
	FileExtension

	ScopeType     string `json:"scopeType"`
	ScopeLongType string `json:"scopeLongType"`
	ScopeFullType string `json:"scopeFullType"`
}

// Enum contains details about enumerations. These can be either top level enums, or nested (defined within a message).
type Enum struct {
	Name        string       `json:"name"`
	LongName    string       `json:"longName"`
	FullName    string       `json:"fullName"`
	Description string       `json:"description"`
	Values      []*EnumValue `json:"values"`

	Options map[string]interface{} `json:"options,omitempty"`
}

// Option returns the named option.
func (e Enum) Option(name string) interface{} { _ = "STUB: not implemented"; return nil }

// ValueOptions returns all options that are set on the values in this enum.
func (e Enum) ValueOptions() []string { _ = "STUB: not implemented"; return nil }

// ValuesWithOption returns all values that have the given option set.
// If no single value has the option set, this returns nil.
func (e Enum) ValuesWithOption(optionName string) []*EnumValue {
	_ = "STUB: not implemented"
	return nil
}

// EnumValue contains details about an individual value within an enumeration.
type EnumValue struct {
	Name        string `json:"name"`
	Number      string `json:"number"`
	Description string `json:"description"`

	Options map[string]interface{} `json:"options,omitempty"`
}

// Option returns the named option.
func (v EnumValue) Option(name string) interface{} { _ = "STUB: not implemented"; return nil }

// Service contains details about a service definition within a proto file.
type Service struct {
	Name        string           `json:"name"`
	LongName    string           `json:"longName"`
	FullName    string           `json:"fullName"`
	Description string           `json:"description"`
	Methods     []*ServiceMethod `json:"methods"`

	Options map[string]interface{} `json:"options,omitempty"`
}

// Option returns the named option.
func (s Service) Option(name string) interface{} { _ = "STUB: not implemented"; return nil }

// MethodOptions returns all options that are set on the methods in this service.
func (s Service) MethodOptions() []string { _ = "STUB: not implemented"; return nil }

// MethodsWithOption returns all methods that have the given option set.
// If no single method has the option set, this returns nil.
func (s Service) MethodsWithOption(optionName string) []*ServiceMethod {
	_ = "STUB: not implemented"
	return nil
}

// ServiceMethod contains details about an individual method within a service.
type ServiceMethod struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	RequestType       string `json:"requestType"`
	RequestLongType   string `json:"requestLongType"`
	RequestFullType   string `json:"requestFullType"`
	RequestStreaming  bool   `json:"requestStreaming"`
	ResponseType      string `json:"responseType"`
	ResponseLongType  string `json:"responseLongType"`
	ResponseFullType  string `json:"responseFullType"`
	ResponseStreaming bool   `json:"responseStreaming"`

	Options map[string]interface{} `json:"options,omitempty"`
}

// Option returns the named option.
func (m ServiceMethod) Option(name string) interface{} { _ = "STUB: not implemented"; return nil }

// ScalarValue contains information about scalar value types in protobuf. The common use case for this type is to know
// which language specific type maps to the protobuf type.
//
// For example, the protobuf type `int64` maps to `long` in C#, and `Bignum` in Ruby. For the full list, take a look at
// https://developers.google.com/protocol-buffers/docs/proto3#scalar
type ScalarValue struct {
	ProtoType  string `json:"protoType"`
	Notes      string `json:"notes"`
	CppType    string `json:"cppType"`
	CSharp     string `json:"csType"`
	GoType     string `json:"goType"`
	JavaType   string `json:"javaType"`
	PhpType    string `json:"phpType"`
	PythonType string `json:"pythonType"`
	RubyType   string `json:"rubyType"`
}

func parseEnum(pe *protokit.EnumDescriptor) *Enum { _ = "STUB: not implemented"; return nil }

func parseFileExtension(pe *protokit.ExtensionDescriptor) *FileExtension {
	_ = "STUB: not implemented"
	return nil
}

func parseMessage(pm *protokit.Descriptor, pluginOptions *PluginOptions) *Message {
	_ = "STUB: not implemented"
	return nil
}

func parseMessageExtension(pe *protokit.ExtensionDescriptor) *MessageExtension {
	_ = "STUB: not implemented"
	return nil
}

func parseMessageField(pf *protokit.FieldDescriptor, oneofDecls []*descriptor.OneofDescriptorProto, pluginOptions *PluginOptions) *MessageField {
	_ = "STUB: not implemented"
	return nil
}

// Check if this is a map.
// See https://github.com/golang/protobuf/blob/master/protoc-gen-go/descriptor/descriptor.pb.go#L1556
// for more information

func parseService(ps *protokit.ServiceDescriptor) *Service { _ = "STUB: not implemented"; return nil }

func parseServiceMethod(pm *protokit.MethodDescriptor) *ServiceMethod {
	_ = "STUB: not implemented"
	return nil
}

func baseName(name string) string { _ = "STUB: not implemented"; return "" }

func labelName(lbl descriptor.FieldDescriptorProto_Label, proto3 bool, proto3Opt bool) string {
	_ = "STUB: not implemented"
	return ""
}

type typeContainer interface {
	GetType() descriptor.FieldDescriptorProto_Type
	GetTypeName() string
	GetPackage() string
}

func parseType(tc typeContainer) (string, string, string) {
	_ = "STUB: not implemented"
	return "", "", ""
}

func description(comment string) string { _ = "STUB: not implemented"; return "" }

type orderedEnums []*Enum

func (oe orderedEnums) Len() int           { _ = "STUB: not implemented"; return 0 }
func (oe orderedEnums) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (oe orderedEnums) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

type orderedExtensions []*FileExtension

func (oe orderedExtensions) Len() int           { _ = "STUB: not implemented"; return 0 }
func (oe orderedExtensions) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (oe orderedExtensions) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

type orderedMessages []*Message

func (om orderedMessages) Len() int           { _ = "STUB: not implemented"; return 0 }
func (om orderedMessages) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (om orderedMessages) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

type orderedServices []*Service

func (os orderedServices) Len() int           { _ = "STUB: not implemented"; return 0 }
func (os orderedServices) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (os orderedServices) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
