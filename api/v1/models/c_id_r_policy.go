// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CIDRPolicy CIDR endpoint policy
// swagger:model CIDRPolicy

type CIDRPolicy struct {

	// List of CIDR egress rules
	Egress []string `json:"egress"`

	// List of CIDR ingress rules
	Ingress []string `json:"ingress"`
}

/* polymorph CIDRPolicy egress false */

/* polymorph CIDRPolicy ingress false */

// Validate validates this c ID r policy
func (m *CIDRPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEgress(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIngress(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CIDRPolicy) validateEgress(formats strfmt.Registry) error {

	if swag.IsZero(m.Egress) { // not required
		return nil
	}

	return nil
}

func (m *CIDRPolicy) validateIngress(formats strfmt.Registry) error {

	if swag.IsZero(m.Ingress) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CIDRPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CIDRPolicy) UnmarshalBinary(b []byte) error {
	var res CIDRPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
