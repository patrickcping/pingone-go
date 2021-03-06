/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"fmt"
)

// CreateGatewayRequest - struct for CreateGatewayRequest
type CreateGatewayRequest struct {
	Gateway *Gateway
	GatewayLDAP *GatewayLDAP
}

// GatewayAsCreateGatewayRequest is a convenience function that returns Gateway wrapped in CreateGatewayRequest
func GatewayAsCreateGatewayRequest(v *Gateway) CreateGatewayRequest {
	return CreateGatewayRequest{
		Gateway: v,
	}
}

// GatewayLDAPAsCreateGatewayRequest is a convenience function that returns GatewayLDAP wrapped in CreateGatewayRequest
func GatewayLDAPAsCreateGatewayRequest(v *GatewayLDAP) CreateGatewayRequest {
	return CreateGatewayRequest{
		GatewayLDAP: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateGatewayRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Gateway
	err = newStrictDecoder(data).Decode(&dst.Gateway)
	if err == nil {
		jsonGateway, _ := json.Marshal(dst.Gateway)
		if string(jsonGateway) == "{}" { // empty struct
			dst.Gateway = nil
		} else {
			match++
		}
	} else {
		dst.Gateway = nil
	}

	// try to unmarshal data into GatewayLDAP
	err = newStrictDecoder(data).Decode(&dst.GatewayLDAP)
	if err == nil {
		jsonGatewayLDAP, _ := json.Marshal(dst.GatewayLDAP)
		if string(jsonGatewayLDAP) == "{}" { // empty struct
			dst.GatewayLDAP = nil
		} else {
			match++
		}
	} else {
		dst.GatewayLDAP = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Gateway = nil
		dst.GatewayLDAP = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(CreateGatewayRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(CreateGatewayRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateGatewayRequest) MarshalJSON() ([]byte, error) {
	if src.Gateway != nil {
		return json.Marshal(&src.Gateway)
	}

	if src.GatewayLDAP != nil {
		return json.Marshal(&src.GatewayLDAP)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateGatewayRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Gateway != nil {
		return obj.Gateway
	}

	if obj.GatewayLDAP != nil {
		return obj.GatewayLDAP
	}

	// all schemas are nil
	return nil
}

type NullableCreateGatewayRequest struct {
	value *CreateGatewayRequest
	isSet bool
}

func (v NullableCreateGatewayRequest) Get() *CreateGatewayRequest {
	return v.value
}

func (v *NullableCreateGatewayRequest) Set(val *CreateGatewayRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGatewayRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGatewayRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGatewayRequest(val *CreateGatewayRequest) *NullableCreateGatewayRequest {
	return &NullableCreateGatewayRequest{value: val, isSet: true}
}

func (v NullableCreateGatewayRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGatewayRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


