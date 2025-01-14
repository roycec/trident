// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ClusterPeerSetupResponse cluster peer setup response
//
// swagger:model cluster_peer_setup_response
type ClusterPeerSetupResponse struct {

	// links
	Links *ClusterPeerSetupResponseInlineLinks `json:"_links,omitempty"`

	// authentication
	Authentication *ClusterPeerSetupResponseInlineAuthentication `json:"authentication,omitempty"`

	// A local intercluster IP address that a remote cluster can use, together with the passphrase, to create a cluster peer relationship with the local cluster.
	IPAddress *IPAddress `json:"ip_address,omitempty"`

	// Optional name for the cluster peer relationship. By default, it is the name of the remote cluster, or a temporary name might be autogenerated for anonymous cluster peer offers.
	// Example: cluster2
	Name *string `json:"name,omitempty"`
}

// Validate validates this cluster peer setup response
func (m *ClusterPeerSetupResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthentication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterPeerSetupResponse) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterPeerSetupResponse) validateAuthentication(formats strfmt.Registry) error {
	if swag.IsZero(m.Authentication) { // not required
		return nil
	}

	if m.Authentication != nil {
		if err := m.Authentication.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authentication")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterPeerSetupResponse) validateIPAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.IPAddress) { // not required
		return nil
	}

	if m.IPAddress != nil {
		if err := m.IPAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ip_address")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster peer setup response based on the context it is used
func (m *ClusterPeerSetupResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAuthentication(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIPAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterPeerSetupResponse) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterPeerSetupResponse) contextValidateAuthentication(ctx context.Context, formats strfmt.Registry) error {

	if m.Authentication != nil {
		if err := m.Authentication.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authentication")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterPeerSetupResponse) contextValidateIPAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.IPAddress != nil {
		if err := m.IPAddress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ip_address")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterPeerSetupResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterPeerSetupResponse) UnmarshalBinary(b []byte) error {
	var res ClusterPeerSetupResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClusterPeerSetupResponseInlineAuthentication cluster peer setup response inline authentication
//
// swagger:model cluster_peer_setup_response_inline_authentication
type ClusterPeerSetupResponseInlineAuthentication struct {

	// The date and time the passphrase will expire.  The default expiry time is one hour.
	// Example: 2017-01-25T11:20:13Z
	// Format: date-time
	ExpiryTime *strfmt.DateTime `json:"expiry_time,omitempty"`

	// A password to authenticate the cluster peer relationship.
	Passphrase *string `json:"passphrase,omitempty"`
}

// Validate validates this cluster peer setup response inline authentication
func (m *ClusterPeerSetupResponseInlineAuthentication) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpiryTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterPeerSetupResponseInlineAuthentication) validateExpiryTime(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpiryTime) { // not required
		return nil
	}

	if err := validate.FormatOf("authentication"+"."+"expiry_time", "body", "date-time", m.ExpiryTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cluster peer setup response inline authentication based on context it is used
func (m *ClusterPeerSetupResponseInlineAuthentication) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterPeerSetupResponseInlineAuthentication) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterPeerSetupResponseInlineAuthentication) UnmarshalBinary(b []byte) error {
	var res ClusterPeerSetupResponseInlineAuthentication
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClusterPeerSetupResponseInlineLinks cluster peer setup response inline links
//
// swagger:model cluster_peer_setup_response_inline__links
type ClusterPeerSetupResponseInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this cluster peer setup response inline links
func (m *ClusterPeerSetupResponseInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterPeerSetupResponseInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster peer setup response inline links based on the context it is used
func (m *ClusterPeerSetupResponseInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterPeerSetupResponseInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterPeerSetupResponseInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterPeerSetupResponseInlineLinks) UnmarshalBinary(b []byte) error {
	var res ClusterPeerSetupResponseInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
