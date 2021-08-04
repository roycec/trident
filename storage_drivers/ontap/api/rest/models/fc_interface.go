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

// FcInterface A Fibre Channel (FC) interface is the logical endpoint for FC network connections to an SVM. An FC interface provides FC access to storage within the interface SVM using either Fibre Channel Protocol or NVMe over Fibre Channel (NVMe/FC).<br/>
// An FC interface is created on an FC port which is located on a cluster node. The FC port must be specified to identify the location of the interface for a POST or PATCH operation that relocates an interface. You can identify the port by supplying either the node and port names or the port UUID.
//
// swagger:model fc_interface
type FcInterface struct {

	// links
	Links *FcInterfaceLinks `json:"_links,omitempty"`

	// A user configurable comment. Optional in POST; valid in PATCH. To clear a prior comment, set the property to an empty string in PATCH.
	//
	Comment string `json:"comment,omitempty"`

	// The data protocol for which the FC interface is configured. Required in POST.
	//
	// Enum: [fcp fc_nvme]
	DataProtocol string `json:"data_protocol,omitempty"`

	// The administrative state of the FC interface. The FC interface can be disabled to block all FC communication with the SVM through this interface. Optional in POST and PATCH; defaults to _true_ (enabled) in POST.
	//
	Enabled *bool `json:"enabled,omitempty"`

	// location
	Location *FcInterfaceLocation `json:"location,omitempty"`

	// The name of the FC interface. Required in POST; optional in PATCH.
	//
	// Example: lif1
	Name string `json:"name,omitempty"`

	// The port address of the FC interface. Each FC port in an FC switched fabric has its own unique FC port address for routing purposes. The FC port address is assigned by a switch in the fabric when that port logs in to the fabric. This property refers to the address given by a switch to the FC interface when the SVM performs a port login (PLOGI).<br/>
	// This is useful for obtaining statistics and diagnostic information from FC switches.<br/>
	// This is a hexadecimal encoded numeric value.
	//
	// Example: 5060F
	// Read Only: true
	PortAddress string `json:"port_address,omitempty"`

	// The current operational state of the FC interface. The state is set to _down_ if the interface is not enabled.<br/>
	// If the node hosting the port is down or unavailable, no state value is returned.
	//
	// Read Only: true
	// Enum: [up down]
	State string `json:"state,omitempty"`

	// svm
	Svm *FcInterfaceSvmType `json:"svm,omitempty"`

	// The unique identifier of the FC interface. Required in the URL.
	//
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	// Read Only: true
	UUID string `json:"uuid,omitempty"`

	// The world wide node name (WWNN) of the FC interface SVM. The WWNN is generated by ONTAP when Fibre Channel Protocol or the NVMe service is created for the FC interface SVM.
	//
	// Example: 20:00:00:50:56:b4:13:01
	// Read Only: true
	Wwnn string `json:"wwnn,omitempty"`

	// The world wide port name (WWPN) of the FC interface. The WWPN is generated by ONTAP when the FC interface is created.
	//
	// Example: 20:00:00:50:56:b4:13:a8
	// Read Only: true
	Wwpn string `json:"wwpn,omitempty"`
}

// Validate validates this fc interface
func (m *FcInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcInterface) validateLinks(formats strfmt.Registry) error {
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

var fcInterfaceTypeDataProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["fcp","fc_nvme"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fcInterfaceTypeDataProtocolPropEnum = append(fcInterfaceTypeDataProtocolPropEnum, v)
	}
}

const (

	// BEGIN RIPPY DEBUGGING
	// fc_interface
	// FcInterface
	// data_protocol
	// DataProtocol
	// fcp
	// END RIPPY DEBUGGING
	// FcInterfaceDataProtocolFcp captures enum value "fcp"
	FcInterfaceDataProtocolFcp string = "fcp"

	// BEGIN RIPPY DEBUGGING
	// fc_interface
	// FcInterface
	// data_protocol
	// DataProtocol
	// fc_nvme
	// END RIPPY DEBUGGING
	// FcInterfaceDataProtocolFcNvme captures enum value "fc_nvme"
	FcInterfaceDataProtocolFcNvme string = "fc_nvme"
)

// prop value enum
func (m *FcInterface) validateDataProtocolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fcInterfaceTypeDataProtocolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FcInterface) validateDataProtocol(formats strfmt.Registry) error {
	if swag.IsZero(m.DataProtocol) { // not required
		return nil
	}

	// value enum
	if err := m.validateDataProtocolEnum("data_protocol", "body", m.DataProtocol); err != nil {
		return err
	}

	return nil
}

func (m *FcInterface) validateLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

var fcInterfaceTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["up","down"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fcInterfaceTypeStatePropEnum = append(fcInterfaceTypeStatePropEnum, v)
	}
}

const (

	// BEGIN RIPPY DEBUGGING
	// fc_interface
	// FcInterface
	// state
	// State
	// up
	// END RIPPY DEBUGGING
	// FcInterfaceStateUp captures enum value "up"
	FcInterfaceStateUp string = "up"

	// BEGIN RIPPY DEBUGGING
	// fc_interface
	// FcInterface
	// state
	// State
	// down
	// END RIPPY DEBUGGING
	// FcInterfaceStateDown captures enum value "down"
	FcInterfaceStateDown string = "down"
)

// prop value enum
func (m *FcInterface) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fcInterfaceTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FcInterface) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *FcInterface) validateSvm(formats strfmt.Registry) error {
	if swag.IsZero(m.Svm) { // not required
		return nil
	}

	if m.Svm != nil {
		if err := m.Svm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this fc interface based on the context it is used
func (m *FcInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePortAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWwnn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWwpn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcInterface) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *FcInterface) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.Location != nil {
		if err := m.Location.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *FcInterface) contextValidatePortAddress(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "port_address", "body", string(m.PortAddress)); err != nil {
		return err
	}

	return nil
}

func (m *FcInterface) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "state", "body", string(m.State)); err != nil {
		return err
	}

	return nil
}

func (m *FcInterface) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

	if m.Svm != nil {
		if err := m.Svm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

func (m *FcInterface) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uuid", "body", string(m.UUID)); err != nil {
		return err
	}

	return nil
}

func (m *FcInterface) contextValidateWwnn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "wwnn", "body", string(m.Wwnn)); err != nil {
		return err
	}

	return nil
}

func (m *FcInterface) contextValidateWwpn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "wwpn", "body", string(m.Wwpn)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FcInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FcInterface) UnmarshalBinary(b []byte) error {
	var res FcInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FcInterfaceLinks fc interface links
//
// swagger:model FcInterfaceLinks
type FcInterfaceLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this fc interface links
func (m *FcInterfaceLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcInterfaceLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this fc interface links based on the context it is used
func (m *FcInterfaceLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcInterfaceLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *FcInterfaceLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FcInterfaceLinks) UnmarshalBinary(b []byte) error {
	var res FcInterfaceLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FcInterfaceLocation The location of the FC interface is defined by the location of its port. An FC port is identified by its UUID, or a combination of its node name and port name. Either the UUID or the node name and port name are required for POST. To move an interface, supply either the UUID or the node name and port name in a PATCH.
//
//
// swagger:model FcInterfaceLocation
type FcInterfaceLocation struct {

	// node
	Node *FcInterfaceLocationNode `json:"node,omitempty"`

	// port
	Port *FcPortReference `json:"port,omitempty"`
}

// Validate validates this fc interface location
func (m *FcInterfaceLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcInterfaceLocation) validateNode(formats strfmt.Registry) error {
	if swag.IsZero(m.Node) { // not required
		return nil
	}

	if m.Node != nil {
		if err := m.Node.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location" + "." + "node")
			}
			return err
		}
	}

	return nil
}

func (m *FcInterfaceLocation) validatePort(formats strfmt.Registry) error {
	if swag.IsZero(m.Port) { // not required
		return nil
	}

	if m.Port != nil {
		if err := m.Port.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location" + "." + "port")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this fc interface location based on the context it is used
func (m *FcInterfaceLocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePort(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcInterfaceLocation) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

	if m.Node != nil {
		if err := m.Node.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location" + "." + "node")
			}
			return err
		}
	}

	return nil
}

func (m *FcInterfaceLocation) contextValidatePort(ctx context.Context, formats strfmt.Registry) error {

	if m.Port != nil {
		if err := m.Port.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location" + "." + "port")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FcInterfaceLocation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FcInterfaceLocation) UnmarshalBinary(b []byte) error {
	var res FcInterfaceLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FcInterfaceLocationNode fc interface location node
//
// swagger:model FcInterfaceLocationNode
type FcInterfaceLocationNode struct {

	// links
	Links *FcInterfaceLocationNodeLinks `json:"_links,omitempty"`

	// name
	// Example: node1
	Name string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this fc interface location node
func (m *FcInterfaceLocationNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcInterfaceLocationNode) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location" + "." + "node" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this fc interface location node based on the context it is used
func (m *FcInterfaceLocationNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcInterfaceLocationNode) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location" + "." + "node" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FcInterfaceLocationNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FcInterfaceLocationNode) UnmarshalBinary(b []byte) error {
	var res FcInterfaceLocationNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FcInterfaceLocationNodeLinks fc interface location node links
//
// swagger:model FcInterfaceLocationNodeLinks
type FcInterfaceLocationNodeLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this fc interface location node links
func (m *FcInterfaceLocationNodeLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcInterfaceLocationNodeLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location" + "." + "node" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this fc interface location node links based on the context it is used
func (m *FcInterfaceLocationNodeLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcInterfaceLocationNodeLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location" + "." + "node" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FcInterfaceLocationNodeLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FcInterfaceLocationNodeLinks) UnmarshalBinary(b []byte) error {
	var res FcInterfaceLocationNodeLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FcInterfaceSvmType fc interface svm type
//
// swagger:model FcInterfaceSvmType
type FcInterfaceSvmType struct {

	// links
	Links *FcInterfaceSvmLinksType `json:"_links,omitempty"`

	// The name of the SVM.
	//
	// Example: svm1
	Name string `json:"name,omitempty"`

	// The unique identifier of the SVM.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this fc interface svm type
func (m *FcInterfaceSvmType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcInterfaceSvmType) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this fc interface svm type based on the context it is used
func (m *FcInterfaceSvmType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcInterfaceSvmType) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FcInterfaceSvmType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FcInterfaceSvmType) UnmarshalBinary(b []byte) error {
	var res FcInterfaceSvmType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FcInterfaceSvmLinksType fc interface svm links type
//
// swagger:model FcInterfaceSvmLinksType
type FcInterfaceSvmLinksType struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this fc interface svm links type
func (m *FcInterfaceSvmLinksType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcInterfaceSvmLinksType) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this fc interface svm links type based on the context it is used
func (m *FcInterfaceSvmLinksType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcInterfaceSvmLinksType) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FcInterfaceSvmLinksType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FcInterfaceSvmLinksType) UnmarshalBinary(b []byte) error {
	var res FcInterfaceSvmLinksType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HELLO RIPPY