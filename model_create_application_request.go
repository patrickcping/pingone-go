/*
PingOne Platform API - Management

A bare-bones collection for the PingOne API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pingone

import (
	"encoding/json"
	"fmt"
)

// CreateApplicationRequest - struct for CreateApplicationRequest
type CreateApplicationRequest struct {
	ApplicationOIDC *ApplicationOIDC
	ApplicationSAML *ApplicationSAML
}

// ApplicationOIDCAsCreateApplicationRequest is a convenience function that returns ApplicationOIDC wrapped in CreateApplicationRequest
func ApplicationOIDCAsCreateApplicationRequest(v *ApplicationOIDC) CreateApplicationRequest {
	return CreateApplicationRequest{
		ApplicationOIDC: v,
	}
}

// ApplicationSAMLAsCreateApplicationRequest is a convenience function that returns ApplicationSAML wrapped in CreateApplicationRequest
func ApplicationSAMLAsCreateApplicationRequest(v *ApplicationSAML) CreateApplicationRequest {
	return CreateApplicationRequest{
		ApplicationSAML: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateApplicationRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ApplicationOIDC
	err = newStrictDecoder(data).Decode(&dst.ApplicationOIDC)
	if err == nil {
		jsonApplicationOIDC, _ := json.Marshal(dst.ApplicationOIDC)
		if string(jsonApplicationOIDC) == "{}" { // empty struct
			dst.ApplicationOIDC = nil
		} else {
			match++
		}
	} else {
		dst.ApplicationOIDC = nil
	}

	// try to unmarshal data into ApplicationSAML
	err = newStrictDecoder(data).Decode(&dst.ApplicationSAML)
	if err == nil {
		jsonApplicationSAML, _ := json.Marshal(dst.ApplicationSAML)
		if string(jsonApplicationSAML) == "{}" { // empty struct
			dst.ApplicationSAML = nil
		} else {
			match++
		}
	} else {
		dst.ApplicationSAML = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ApplicationOIDC = nil
		dst.ApplicationSAML = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(CreateApplicationRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(CreateApplicationRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateApplicationRequest) MarshalJSON() ([]byte, error) {
	if src.ApplicationOIDC != nil {
		return json.Marshal(&src.ApplicationOIDC)
	}

	if src.ApplicationSAML != nil {
		return json.Marshal(&src.ApplicationSAML)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateApplicationRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ApplicationOIDC != nil {
		return obj.ApplicationOIDC
	}

	if obj.ApplicationSAML != nil {
		return obj.ApplicationSAML
	}

	// all schemas are nil
	return nil
}

type NullableCreateApplicationRequest struct {
	value *CreateApplicationRequest
	isSet bool
}

func (v NullableCreateApplicationRequest) Get() *CreateApplicationRequest {
	return v.value
}

func (v *NullableCreateApplicationRequest) Set(val *CreateApplicationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateApplicationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateApplicationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateApplicationRequest(val *CreateApplicationRequest) *NullableCreateApplicationRequest {
	return &NullableCreateApplicationRequest{value: val, isSet: true}
}

func (v NullableCreateApplicationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateApplicationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

