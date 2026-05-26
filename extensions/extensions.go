// Package extensions implements a system for working with extended options.
package extensions

// Transformer functions for transforming payloads of an extension option into
// something that can be rendered by a template.
type Transformer func(payload interface{}) interface{}

var transformers = make(map[string]Transformer)

// SetTransformer sets the transformer function for the given extension name
func SetTransformer(extensionName string, f Transformer) { _ = "STUB: not implemented"; return }

// Transform the extensions using the registered transformers.
func Transform(extensions map[string]interface{}) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

// No transformer registered, skip.

// Transformer returned nothing, skip.
