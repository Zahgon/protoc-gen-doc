package extensions

import (
	"github.com/pseudomuto/protoc-gen-doc/extensions"
	"google.golang.org/genproto/googleapis/api/annotations"
)

// HTTPRule represents a single HTTP rule from the (google.api.http) method option extension.
type HTTPRule struct {
	Method  string `json:"method"`
	Pattern string `json:"pattern"`
	Body    string `json:"body,omitempty"`
}

// HTTPExtension contains the rules set by the (google.api.http) method option extension.
type HTTPExtension struct {
	Rules []HTTPRule `json:"rules"`
}

func getRule(r *annotations.HttpRule) (rule HTTPRule) {
	_ = "STUB: not implemented"
	return *new(HTTPRule)
}

func init() {
	extensions.SetTransformer("google.api.http", func(payload interface{}) interface{} {
		var rules []HTTPRule
		rule, ok := payload.(*annotations.HttpRule)
		if !ok {
			return nil
		}

		rules = append(rules, getRule(rule))

		// NOTE: The option can only have one level of nested AdditionalBindings.
		for _, rule := range rule.AdditionalBindings {
			rules = append(rules, getRule(rule))
		}

		return HTTPExtension{Rules: rules}
	})
}
