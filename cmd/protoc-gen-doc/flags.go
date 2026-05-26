package main

import (
	"flag"
	"io"
)

const helpMessage = `
This is a protoc plugin that is used to generate documentation from your protobuf files. Invocation is controlled by
using the doc_opt and doc_out options for protoc.

EXAMPLE: Generate HTML docs
protoc --doc_out=. --doc_opt=html,index.html protos/*.proto

EXAMPLE: Exclude file patterns
protoc --doc_out=. --doc_opt=html,index.html:google/*,somedir/* protos/*.proto

EXAMPLE: Use a custom template
protoc --doc_out=. --doc_opt=custom.tmpl,docs.txt protos/*.proto

EXAMPLE: Generate docs relative to source protos
protoc --doc_out=. --doc_opt=html,index.html,source_relative protos/*.proto

See https://github.com/pseudomuto/protoc-gen-doc for more details.
`

// Version returns the currently running version of protoc-gen-doc
func Version() string { _ = "STUB: not implemented"; return "" }

// Flags contains details about the CLI invocation of protoc-gen-doc
type Flags struct {
	appName     string
	flagSet     *flag.FlagSet
	err         error
	showHelp    bool
	showVersion bool
	writer      io.Writer
}

// Code returns the status code to exit with after handling the supplied flags
func (f *Flags) Code() int { _ = "STUB: not implemented"; return 0 }

// HasMatch returns whether or not the supplied args are matches. For example, passing `--help` will match, or some
// unknown parameter, but passing nothing will not.
func (f *Flags) HasMatch() bool { _ = "STUB: not implemented"; return false }

// ShowHelp determines whether or not to show the help message
func (f *Flags) ShowHelp() bool { _ = "STUB: not implemented"; return false }

// ShowVersion determines whether or not to show the version message
func (f *Flags) ShowVersion() bool { _ = "STUB: not implemented"; return false }

// PrintHelp prints the usage string including all flags to the `io.Writer` that was supplied to the `Flags` object.
func (f *Flags) PrintHelp() { _ = "STUB: not implemented"; return }

// PrintVersion prints the version string to the `io.Writer` that was supplied to the `Flags` object.
func (f *Flags) PrintVersion() { _ = "STUB: not implemented"; return }

// ParseFlags parses the supplied options are returns a `Flags` object to the caller.
//
// Parameters:
//   - `w` - the `io.Writer` to use for printing messages (help, version, etc.)
//   - `args` - the set of args the program was invoked with (typically `os.Args`)
func ParseFlags(w io.Writer, args []string) *Flags { _ = "STUB: not implemented"; return nil }

// prevent showing help on parse error
