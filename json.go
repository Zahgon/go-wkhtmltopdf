package wkhtmltopdf

import (
	"io"
)

type jsonPDFGenerator struct {
	GlobalOptions  globalOptions
	OutlineOptions outlineOptions
	Cover          cover
	TOC            toc
	Pages          []jsonPage
}

type jsonPage struct {
	PageOptions    PageOptions
	InputFile      string
	Base64PageData string
}

// ToJSON creates JSON of the complete representation of the PDFGenerator.
// It also saves all pages. For a PageReader page, the content is stored as a Base64 string in the JSON.
func (pdfg *PDFGenerator) ToJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// NewPDFGeneratorFromJSON creates a new PDFGenerator and restores all the settings and pages
// from a JSON byte slice which should be created using PDFGenerator.ToJSON().
func NewPDFGeneratorFromJSON(jsonReader io.Reader) (*PDFGenerator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type jsonBoolOption struct {
	Option string
	Value  bool
}

func (bo *boolOption) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (bo *boolOption) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

type jsonStringOption struct {
	Option string
	Value  string
}

func (so *stringOption) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (so *stringOption) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

type jsonUintOption struct {
	Option string
	IsSet  bool
	Value  uint
}

func (io *uintOption) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (io *uintOption) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

type jsonFloatOption struct {
	Option string
	IsSet  bool
	Value  float64
}

func (fo *floatOption) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (fo *floatOption) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

type jsonMapOption struct {
	Option string
	Value  map[string]string
}

func (mo *mapOption) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (mo *mapOption) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

type jsonSliceOption struct {
	Option string
	Value  []string
}

func (so *sliceOption) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (so *sliceOption) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }
