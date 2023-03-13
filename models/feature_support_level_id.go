// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// FeatureSupportLevelID feature support level id
//
// swagger:model feature-support-level-id
type FeatureSupportLevelID string

func NewFeatureSupportLevelID(value FeatureSupportLevelID) *FeatureSupportLevelID {
	return &value
}

// Pointer returns a pointer to a freshly-allocated FeatureSupportLevelID.
func (m FeatureSupportLevelID) Pointer() *FeatureSupportLevelID {
	return &m
}

const (

	// FeatureSupportLevelIDADDITIONALNTPSOURCE captures enum value "ADDITIONAL_NTP_SOURCE"
	FeatureSupportLevelIDADDITIONALNTPSOURCE FeatureSupportLevelID = "ADDITIONAL_NTP_SOURCE"

	// FeatureSupportLevelIDREQUESTEDHOSTNAME captures enum value "REQUESTED_HOSTNAME"
	FeatureSupportLevelIDREQUESTEDHOSTNAME FeatureSupportLevelID = "REQUESTED_HOSTNAME"

	// FeatureSupportLevelIDPROXY captures enum value "PROXY"
	FeatureSupportLevelIDPROXY FeatureSupportLevelID = "PROXY"

	// FeatureSupportLevelIDSNO captures enum value "SNO"
	FeatureSupportLevelIDSNO FeatureSupportLevelID = "SNO"

	// FeatureSupportLevelIDDAY2HOSTS captures enum value "DAY2_HOSTS"
	FeatureSupportLevelIDDAY2HOSTS FeatureSupportLevelID = "DAY2_HOSTS"

	// FeatureSupportLevelIDVIPAUTOALLOC captures enum value "VIP_AUTO_ALLOC"
	FeatureSupportLevelIDVIPAUTOALLOC FeatureSupportLevelID = "VIP_AUTO_ALLOC"

	// FeatureSupportLevelIDDISKSELECTION captures enum value "DISK_SELECTION"
	FeatureSupportLevelIDDISKSELECTION FeatureSupportLevelID = "DISK_SELECTION"

	// FeatureSupportLevelIDOVNNETWORKTYPE captures enum value "OVN_NETWORK_TYPE"
	FeatureSupportLevelIDOVNNETWORKTYPE FeatureSupportLevelID = "OVN_NETWORK_TYPE"

	// FeatureSupportLevelIDSDNNETWORKTYPE captures enum value "SDN_NETWORK_TYPE"
	FeatureSupportLevelIDSDNNETWORKTYPE FeatureSupportLevelID = "SDN_NETWORK_TYPE"

	// FeatureSupportLevelIDSCHEDULABLEMASTERS captures enum value "SCHEDULABLE_MASTERS"
	FeatureSupportLevelIDSCHEDULABLEMASTERS FeatureSupportLevelID = "SCHEDULABLE_MASTERS"

	// FeatureSupportLevelIDAUTOASSIGNROLE captures enum value "AUTO_ASSIGN_ROLE"
	FeatureSupportLevelIDAUTOASSIGNROLE FeatureSupportLevelID = "AUTO_ASSIGN_ROLE"

	// FeatureSupportLevelIDCUSTOMMANIFEST captures enum value "CUSTOM_MANIFEST"
	FeatureSupportLevelIDCUSTOMMANIFEST FeatureSupportLevelID = "CUSTOM_MANIFEST"

	// FeatureSupportLevelIDDISKENCRYPTION captures enum value "DISK_ENCRYPTION"
	FeatureSupportLevelIDDISKENCRYPTION FeatureSupportLevelID = "DISK_ENCRYPTION"

	// FeatureSupportLevelIDCLUSTERMANAGEDNETWORKINGWITHVMS captures enum value "CLUSTER_MANAGED_NETWORKING_WITH_VMS"
	FeatureSupportLevelIDCLUSTERMANAGEDNETWORKINGWITHVMS FeatureSupportLevelID = "CLUSTER_MANAGED_NETWORKING_WITH_VMS"

	// FeatureSupportLevelIDSINGLENODEEXPANSION captures enum value "SINGLE_NODE_EXPANSION"
	FeatureSupportLevelIDSINGLENODEEXPANSION FeatureSupportLevelID = "SINGLE_NODE_EXPANSION"

	// FeatureSupportLevelIDLVM captures enum value "LVM"
	FeatureSupportLevelIDLVM FeatureSupportLevelID = "LVM"

	// FeatureSupportLevelIDODF captures enum value "ODF"
	FeatureSupportLevelIDODF FeatureSupportLevelID = "ODF"

	// FeatureSupportLevelIDCNV captures enum value "CNV"
	FeatureSupportLevelIDCNV FeatureSupportLevelID = "CNV"

	// FeatureSupportLevelIDDUALSTACKNETWORKING captures enum value "DUAL_STACK_NETWORKING"
	FeatureSupportLevelIDDUALSTACKNETWORKING FeatureSupportLevelID = "DUAL_STACK_NETWORKING"

	// FeatureSupportLevelIDNUTANIXINTEGRATION captures enum value "NUTANIX_INTEGRATION"
	FeatureSupportLevelIDNUTANIXINTEGRATION FeatureSupportLevelID = "NUTANIX_INTEGRATION"

	// FeatureSupportLevelIDVSPHEREINTEGRATION captures enum value "VSPHERE_INTEGRATION"
	FeatureSupportLevelIDVSPHEREINTEGRATION FeatureSupportLevelID = "VSPHERE_INTEGRATION"

	// FeatureSupportLevelIDDUALSTACKVIPS captures enum value "DUAL_STACK_VIPS"
	FeatureSupportLevelIDDUALSTACKVIPS FeatureSupportLevelID = "DUAL_STACK_VIPS"

	// FeatureSupportLevelIDUSERMANAGEDNETWORKINGWITHMULTINODE captures enum value "USER_MANAGED_NETWORKING_WITH_MULTI_NODE"
	FeatureSupportLevelIDUSERMANAGEDNETWORKINGWITHMULTINODE FeatureSupportLevelID = "USER_MANAGED_NETWORKING_WITH_MULTI_NODE"

	// FeatureSupportLevelIDCLUSTERMANAGEDNETWORKING captures enum value "CLUSTER_MANAGED_NETWORKING"
	FeatureSupportLevelIDCLUSTERMANAGEDNETWORKING FeatureSupportLevelID = "CLUSTER_MANAGED_NETWORKING"
)

// for schema
var featureSupportLevelIdEnum []interface{}

func init() {
	var res []FeatureSupportLevelID
	if err := json.Unmarshal([]byte(`["ADDITIONAL_NTP_SOURCE","REQUESTED_HOSTNAME","PROXY","SNO","DAY2_HOSTS","VIP_AUTO_ALLOC","DISK_SELECTION","OVN_NETWORK_TYPE","SDN_NETWORK_TYPE","SCHEDULABLE_MASTERS","AUTO_ASSIGN_ROLE","CUSTOM_MANIFEST","DISK_ENCRYPTION","CLUSTER_MANAGED_NETWORKING_WITH_VMS","SINGLE_NODE_EXPANSION","LVM","ODF","CNV","DUAL_STACK_NETWORKING","NUTANIX_INTEGRATION","VSPHERE_INTEGRATION","DUAL_STACK_VIPS","USER_MANAGED_NETWORKING_WITH_MULTI_NODE","CLUSTER_MANAGED_NETWORKING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		featureSupportLevelIdEnum = append(featureSupportLevelIdEnum, v)
	}
}

func (m FeatureSupportLevelID) validateFeatureSupportLevelIDEnum(path, location string, value FeatureSupportLevelID) error {
	if err := validate.EnumCase(path, location, value, featureSupportLevelIdEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this feature support level id
func (m FeatureSupportLevelID) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateFeatureSupportLevelIDEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this feature support level id based on context it is used
func (m FeatureSupportLevelID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
