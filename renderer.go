package gendoc

// RenderType is an "enum" for which type of renderer to use.
type RenderType int8

// Available render types.
const (
	_ RenderType = iota
	RenderTypeDocBook
	RenderTypeHTML
	RenderTypeJSON
	RenderTypeMarkdown
)

// NewRenderType creates a RenderType from the supplied string. If the type is not known, (0, error) is returned. It is
// assumed (by the plugin) that invalid render type simply means that the path to a custom template was supplied.
func NewRenderType(renderType string) (RenderType, error) {
	_ = "STUB: not implemented"
	return *new(RenderType), nil
}

func (rt RenderType) renderer() (Processor, error) {
	_ = "STUB: not implemented"
	return *new(Processor), nil
}

func (rt RenderType) template() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

var funcMap = map[string]interface{}{
	"p":      PFilter,
	"para":   ParaFilter,
	"nobr":   NoBrFilter,
	"anchor": AnchorFilter,
}

// Processor is an interface that is satisfied by all built-in processors (text, html, and json).
type Processor interface {
	Apply(template *Template) ([]byte, error)
}

// RenderTemplate renders the template based on the render type. It supports overriding the default input templates by
// supplying a non-empty string as the last parameter.
//
// Example: generating an HTML template (assuming you've got a Template object)
//
//	data, err := RenderTemplate(RenderTypeHTML, &template, "")
//
// Example: generating a custom template (assuming you've got a Template object)
//
//	data, err := RenderTemplate(RenderTypeHTML, &template, "{{range .Files}}{{.Name}}{{end}}")
func RenderTemplate(kind RenderType, template *Template, inputTemplate string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type textRenderer struct {
	inputTemplate string
}

func (mr *textRenderer) Apply(template *Template) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type htmlRenderer struct {
	inputTemplate string
}

func (mr *htmlRenderer) Apply(template *Template) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type jsonRenderer struct{}

func (r *jsonRenderer) Apply(template *Template) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
