package gendoc

import (
	"html/template"
	"regexp"
)

var (
	paraPattern         = regexp.MustCompile(`(\n|\r|\r\n)\s*`)
	spacePattern        = regexp.MustCompile("( )+")
	multiNewlinePattern = regexp.MustCompile(`(\r\n|\r|\n){2,}`)
	specialCharsPattern = regexp.MustCompile(`[^a-zA-Z0-9_-]`)
)

// PFilter splits the content by new lines and wraps each one in a <p> tag.
func PFilter(content string) template.HTML { _ = "STUB: not implemented"; return *new(template.HTML) }

// ParaFilter splits the content by new lines and wraps each one in a <para> tag.
func ParaFilter(content string) string { _ = "STUB: not implemented"; return "" }

// NoBrFilter removes single CR and LF from content.
func NoBrFilter(content string) string { _ = "STUB: not implemented"; return "" }

// AnchorFilter replaces all special characters with URL friendly dashes
func AnchorFilter(str string) string { _ = "STUB: not implemented"; return "" }
