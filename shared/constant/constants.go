// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/stainless-sdks/llamacloud-prod-go/internal/encoding/json"
)

type Constant[T any] interface {
	Default() T
}

// ValueOf gives the default value of a constant from its type. It's helpful when
// constructing constants as variants in a one-of. Note that empty structs are
// marshalled by default. Usage: constant.ValueOf[constant.Foo]()
func ValueOf[T Constant[T]]() T {
	var t T
	return t.Default()
}

type ClassifyV2 string    // Always "classify_v2"
type ExtractV2 string     // Always "extract_v2"
type ParseV2 string       // Always "parse_v2"
type SplitV1 string       // Always "split_v1"
type SpreadsheetV1 string // Always "spreadsheet_v1"
type Unknown string       // Always "unknown"

func (c ClassifyV2) Default() ClassifyV2       { return "classify_v2" }
func (c ExtractV2) Default() ExtractV2         { return "extract_v2" }
func (c ParseV2) Default() ParseV2             { return "parse_v2" }
func (c SplitV1) Default() SplitV1             { return "split_v1" }
func (c SpreadsheetV1) Default() SpreadsheetV1 { return "spreadsheet_v1" }
func (c Unknown) Default() Unknown             { return "unknown" }

func (c ClassifyV2) MarshalJSON() ([]byte, error)    { return marshalString(c) }
func (c ExtractV2) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c ParseV2) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c SplitV1) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c SpreadsheetV1) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c Unknown) MarshalJSON() ([]byte, error)       { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return shimjson.Marshal(string(v))
}
