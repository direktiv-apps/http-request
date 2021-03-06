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

// PostParamsBody post params body
//
// swagger:model postParamsBody
type PostParamsBody struct {

	// content
	Content *PostParamsBodyContent `json:"content,omitempty"`

	// Prints the full URL and headers to logs.
	// Example: true
	Debug bool `json:"debug,omitempty"`

	// If set to `true` responses with status above 299 will be treated as errors.
	// Example: true
	Error200 bool `json:"error200,omitempty"`

	// List of key/values send as headers with the request.
	// Example: {"myheader":"value"}
	Headers map[string]string `json:"headers,omitempty"`

	// Skips the verification the server certificate chain and host name.
	// Example: true
	Insecure bool `json:"insecure,omitempty"`

	// HTTP method. Defaults to GET.
	// Example: POST
	// Enum: [GET HEAD POST PUT DELETE OPTIONS CONNECT TRACE PATCH]
	Method string `json:"method,omitempty"`

	// List of key/values appended to URL as query parameters.
	// Example: {"query1":"queryvalue"}
	Params map[string]string `json:"params,omitempty"`

	// If username and password are set, it will be used for basic authenitcation for the request.
	// Example: mypassword
	Password string `json:"password,omitempty"`

	// URL for the request.
	// Example: http://www.direktiv.io
	// Required: true
	URL *string `json:"url"`

	// If username and password are set, it will be used for basic authenitcation for the request.
	// Example: myuser
	Username string `json:"username,omitempty"`
}

// Validate validates this post params body
func (m *PostParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostParamsBody) validateContent(formats strfmt.Registry) error {
	if swag.IsZero(m.Content) { // not required
		return nil
	}

	if m.Content != nil {
		if err := m.Content.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("content")
			}
			return err
		}
	}

	return nil
}

var postParamsBodyTypeMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GET","HEAD","POST","PUT","DELETE","OPTIONS","CONNECT","TRACE","PATCH"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postParamsBodyTypeMethodPropEnum = append(postParamsBodyTypeMethodPropEnum, v)
	}
}

const (

	// PostParamsBodyMethodGET captures enum value "GET"
	PostParamsBodyMethodGET string = "GET"

	// PostParamsBodyMethodHEAD captures enum value "HEAD"
	PostParamsBodyMethodHEAD string = "HEAD"

	// PostParamsBodyMethodPOST captures enum value "POST"
	PostParamsBodyMethodPOST string = "POST"

	// PostParamsBodyMethodPUT captures enum value "PUT"
	PostParamsBodyMethodPUT string = "PUT"

	// PostParamsBodyMethodDELETE captures enum value "DELETE"
	PostParamsBodyMethodDELETE string = "DELETE"

	// PostParamsBodyMethodOPTIONS captures enum value "OPTIONS"
	PostParamsBodyMethodOPTIONS string = "OPTIONS"

	// PostParamsBodyMethodCONNECT captures enum value "CONNECT"
	PostParamsBodyMethodCONNECT string = "CONNECT"

	// PostParamsBodyMethodTRACE captures enum value "TRACE"
	PostParamsBodyMethodTRACE string = "TRACE"

	// PostParamsBodyMethodPATCH captures enum value "PATCH"
	PostParamsBodyMethodPATCH string = "PATCH"
)

// prop value enum
func (m *PostParamsBody) validateMethodEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postParamsBodyTypeMethodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PostParamsBody) validateMethod(formats strfmt.Registry) error {
	if swag.IsZero(m.Method) { // not required
		return nil
	}

	// value enum
	if err := m.validateMethodEnum("method", "body", m.Method); err != nil {
		return err
	}

	return nil
}

func (m *PostParamsBody) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this post params body based on the context it is used
func (m *PostParamsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostParamsBody) contextValidateContent(ctx context.Context, formats strfmt.Registry) error {

	if m.Content != nil {
		if err := m.Content.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("content")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostParamsBody) UnmarshalBinary(b []byte) error {
	var res PostParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
