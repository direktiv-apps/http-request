// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostParamsBodyContent Defines the payload of the request. The `kind` value can have three different values:
//  - string: Plain string payload, e.g. JSON
//  - base64: Will be "converted" to binary and attached
//  - file: File payload
// Example: {"kind":"string","value":"This is the payload"}
//
// swagger:model postParamsBodyContent
type PostParamsBodyContent struct {

	// Kind of data
	// Enum: [string file base64]
	Kind *string `json:"kind,omitempty"`

	// Value depends on `kind` value.
	Value string `json:"value,omitempty"`
}

// Validate validates this post params body content
func (m *PostParamsBodyContent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var postParamsBodyContentTypeKindPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["string","file","base64"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postParamsBodyContentTypeKindPropEnum = append(postParamsBodyContentTypeKindPropEnum, v)
	}
}

const (

	// PostParamsBodyContentKindString captures enum value "string"
	PostParamsBodyContentKindString string = "string"

	// PostParamsBodyContentKindFile captures enum value "file"
	PostParamsBodyContentKindFile string = "file"

	// PostParamsBodyContentKindBase64 captures enum value "base64"
	PostParamsBodyContentKindBase64 string = "base64"
)

// prop value enum
func (m *PostParamsBodyContent) validateKindEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postParamsBodyContentTypeKindPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PostParamsBodyContent) validateKind(formats strfmt.Registry) error {
	if swag.IsZero(m.Kind) { // not required
		return nil
	}

	// value enum
	if err := m.validateKindEnum("kind", "body", *m.Kind); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post params body content based on context it is used
func (m *PostParamsBodyContent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostParamsBodyContent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostParamsBodyContent) UnmarshalBinary(b []byte) error {
	var res PostParamsBodyContent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
