package gendoc

import (
	"regexp"

	plugin_go "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/pseudomuto/protokit"
)

// PluginOptions encapsulates options for the plugin. The type of renderer, template file, and the name of the output
// file are included.
type PluginOptions struct {
	Type            RenderType
	TemplateFile    string
	OutputFile      string
	ExcludePatterns []*regexp.Regexp
	SourceRelative  bool
	CamelCaseFields bool
}

// SupportedFeatures describes a flag setting for supported features.
var SupportedFeatures = uint64(plugin_go.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

// Plugin describes a protoc code generate plugin. It's an implementation of Plugin from github.com/pseudomuto/protokit
type Plugin struct{}

// Generate compiles the documentation and generates the CodeGeneratorResponse to send back to protoc. It does this
// by rendering a template based on the options parsed from the CodeGeneratorRequest.
func (p *Plugin) Generate(r *plugin_go.CodeGeneratorRequest) (*plugin_go.CodeGeneratorResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func groupProtosByDirectory(fds []*protokit.FileDescriptor, sourceRelative bool) map[string][]*protokit.FileDescriptor {
	_ = "STUB: not implemented"
	return nil
}

func excludeUnwantedProtos(fds []*protokit.FileDescriptor, excludePatterns []*regexp.Regexp) []*protokit.FileDescriptor {
	_ = "STUB: not implemented"
	return nil
}

// ParseOptions parses plugin options from a CodeGeneratorRequest. It does this by splitting the `Parameter` field from
// the request object and parsing out the type of renderer to use and the name of the file to be generated.
//
// The parameter (`--doc_opt`) must be of the format <TYPE|TEMPLATE_FILE>,<OUTPUT_FILE>[,default|source_relative]:<EXCLUDE_PATTERN>,<EXCLUDE_PATTERN>*:[,camel_case_fields=[true|false]].
// The file will be written to the directory specified with the `--doc_out` argument to protoc.
func ParseOptions(req *plugin_go.CodeGeneratorRequest) (*PluginOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Parse out exclude patterns if any
