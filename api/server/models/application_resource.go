// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ApplicationResource application resource
//
// swagger:model ApplicationResource
type ApplicationResource struct {

	// applications
	Applications uint32 `json:"applications,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// packages
	Packages uint32 `json:"packages,omitempty"`

	// reporting s b o m analyzers
	ReportingSBOMAnalyzers []string `json:"reportingSBOMAnalyzers"`

	// resource hash
	ResourceHash string `json:"resourceHash,omitempty"`

	// resource name
	ResourceName string `json:"resourceName,omitempty"`

	// resource type
	ResourceType ResourceType `json:"resourceType,omitempty"`

	// vulnerabilities
	Vulnerabilities []*VulnerabilityCount `json:"vulnerabilities"`
}

// Validate validates this application resource
func (m *ApplicationResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResourceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVulnerabilities(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationResource) validateResourceType(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceType) { // not required
		return nil
	}

	if err := m.ResourceType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("resourceType")
		}
		return err
	}

	return nil
}

func (m *ApplicationResource) validateVulnerabilities(formats strfmt.Registry) error {
	if swag.IsZero(m.Vulnerabilities) { // not required
		return nil
	}

	for i := 0; i < len(m.Vulnerabilities); i++ {
		if swag.IsZero(m.Vulnerabilities[i]) { // not required
			continue
		}

		if m.Vulnerabilities[i] != nil {
			if err := m.Vulnerabilities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vulnerabilities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this application resource based on the context it is used
func (m *ApplicationResource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResourceType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVulnerabilities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationResource) contextValidateResourceType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ResourceType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("resourceType")
		}
		return err
	}

	return nil
}

func (m *ApplicationResource) contextValidateVulnerabilities(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Vulnerabilities); i++ {

		if m.Vulnerabilities[i] != nil {
			if err := m.Vulnerabilities[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vulnerabilities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationResource) UnmarshalBinary(b []byte) error {
	var res ApplicationResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
