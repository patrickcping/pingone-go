/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// ImageTargets struct for ImageTargets
type ImageTargets struct {
	// A string that specifies the URL or fully qualified path to the image source file.
	Href *string `json:"href,omitempty"`
	// A string that specifies the UUID of the target image.
	Id *string `json:"id,omitempty"`
	// A string that specifies the type of format used for the image. Options are jpg, png, and gif.
	Type *string `json:"type,omitempty"`
	// The width of the image (in pixels).
	Width *int32 `json:"width,omitempty"`
	// The height of the image (in pixels).
	Height *int32 `json:"height,omitempty"`
}

// NewImageTargets instantiates a new ImageTargets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageTargets() *ImageTargets {
	this := ImageTargets{}
	return &this
}

// NewImageTargetsWithDefaults instantiates a new ImageTargets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageTargetsWithDefaults() *ImageTargets {
	this := ImageTargets{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *ImageTargets) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageTargets) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *ImageTargets) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *ImageTargets) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ImageTargets) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageTargets) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ImageTargets) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ImageTargets) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ImageTargets) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageTargets) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ImageTargets) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ImageTargets) SetType(v string) {
	o.Type = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *ImageTargets) GetWidth() int32 {
	if o == nil || o.Width == nil {
		var ret int32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageTargets) GetWidthOk() (*int32, bool) {
	if o == nil || o.Width == nil {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *ImageTargets) HasWidth() bool {
	if o != nil && o.Width != nil {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
func (o *ImageTargets) SetWidth(v int32) {
	o.Width = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *ImageTargets) GetHeight() int32 {
	if o == nil || o.Height == nil {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageTargets) GetHeightOk() (*int32, bool) {
	if o == nil || o.Height == nil {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *ImageTargets) HasHeight() bool {
	if o != nil && o.Height != nil {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *ImageTargets) SetHeight(v int32) {
	o.Height = &v
}

func (o ImageTargets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Width != nil {
		toSerialize["width"] = o.Width
	}
	if o.Height != nil {
		toSerialize["height"] = o.Height
	}
	return json.Marshal(toSerialize)
}

type NullableImageTargets struct {
	value *ImageTargets
	isSet bool
}

func (v NullableImageTargets) Get() *ImageTargets {
	return v.value
}

func (v *NullableImageTargets) Set(val *ImageTargets) {
	v.value = val
	v.isSet = true
}

func (v NullableImageTargets) IsSet() bool {
	return v.isSet
}

func (v *NullableImageTargets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageTargets(val *ImageTargets) *NullableImageTargets {
	return &NullableImageTargets{value: val, isSet: true}
}

func (v NullableImageTargets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageTargets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

