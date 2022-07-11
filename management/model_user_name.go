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

// UserName struct for UserName
type UserName struct {
	// A string that specifies the family name of the user, or Last in most Western languages (for example, ‘Jensen’ given the full name ‘Ms. Barbara J Jensen, III’). This may be explicitly set to null when updating a name to unset it. Valid characters consist of any Unicode letter, mark (for example, accent, umlaut), space, dot, apostrophe, or hyphen (regex `^[\\p{L}\\p{M}\\p{N}' .-]*$`). It can have a length of no more than 256 characters (min/max=1/256).
	Family *string `json:"family,omitempty"`
	// A string that specifies the fully formatted name of the user (for example ‘Ms. Barbara J Jensen, III’). This can be explicitly set to null when updating a name to unset it. Valid characters consist of any Unicode letter, mark (for example, accent, umlaut), space, dot, apostrophe, or hyphen (regex `^[\\p{L}\\p{M}\\p{N}' .-]*$`). It can have a length of no more than 256 characters (min/max=1/256).
	Formatted *string `json:"formatted,omitempty"`
	// A string that specifies the given name of the user, or First in most Western languages (for example, ‘Barbara’ given the full name ‘Ms. Barbara J Jensen, III’). This may be explicitly set to null when updating a name to unset it. The string can contain any letters, numbers, combining characters, math and currency symbols, dingbats and drawing characters, and invisible whitespace (regex `^[\\p{L}\\p{M}\\p{Zs}\\p{S}\\p{N}\\p{P}]*$`). It can have a length of no more than 256 characters (min/max=1/256).
	Given *string `json:"given,omitempty"`
	// A string that specifies the honorific prefix(es) of the user, or title in most Western languages (for example, ‘Ms.’ given the full name ‘Ms. Barbara Jane Jensen, III’). This can be explicitly set to null when updating a name to unset it.
	HonorificPrefix *string `json:"honorificPrefix,omitempty"`
	// A string that specifies the honorific suffix(es) of the user, or suffix in most Western languages (for example, ‘III’ given the full name ‘Ms. Barbara Jane Jensen, III’). This can be explicitly set to null when updating a name to unset it.
	HonorificSuffix *string `json:"honorificSuffix,omitempty"`
	// A string that specifies the the middle name(s) of the user (for exmple, ‘Jane’ given the full name ‘Ms. Barbara Jane Jensen, III’). This can be explicitly set to null when updating a name to unset it. The string can contain any letters, numbers, combining characters, math and currency symbols, dingbats and drawing characters, and invisible whitespace (regex `^[\\p{L}\\p{M}\\p{Zs}\\p{S}\\p{N}\\p{P}]*$`). It can have a length of no more than 256 characters (min/max=1/256).
	Middle *string `json:"middle,omitempty"`
}

// NewUserName instantiates a new UserName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserName() *UserName {
	this := UserName{}
	return &this
}

// NewUserNameWithDefaults instantiates a new UserName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserNameWithDefaults() *UserName {
	this := UserName{}
	return &this
}

// GetFamily returns the Family field value if set, zero value otherwise.
func (o *UserName) GetFamily() string {
	if o == nil || o.Family == nil {
		var ret string
		return ret
	}
	return *o.Family
}

// GetFamilyOk returns a tuple with the Family field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserName) GetFamilyOk() (*string, bool) {
	if o == nil || o.Family == nil {
		return nil, false
	}
	return o.Family, true
}

// HasFamily returns a boolean if a field has been set.
func (o *UserName) HasFamily() bool {
	if o != nil && o.Family != nil {
		return true
	}

	return false
}

// SetFamily gets a reference to the given string and assigns it to the Family field.
func (o *UserName) SetFamily(v string) {
	o.Family = &v
}

// GetFormatted returns the Formatted field value if set, zero value otherwise.
func (o *UserName) GetFormatted() string {
	if o == nil || o.Formatted == nil {
		var ret string
		return ret
	}
	return *o.Formatted
}

// GetFormattedOk returns a tuple with the Formatted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserName) GetFormattedOk() (*string, bool) {
	if o == nil || o.Formatted == nil {
		return nil, false
	}
	return o.Formatted, true
}

// HasFormatted returns a boolean if a field has been set.
func (o *UserName) HasFormatted() bool {
	if o != nil && o.Formatted != nil {
		return true
	}

	return false
}

// SetFormatted gets a reference to the given string and assigns it to the Formatted field.
func (o *UserName) SetFormatted(v string) {
	o.Formatted = &v
}

// GetGiven returns the Given field value if set, zero value otherwise.
func (o *UserName) GetGiven() string {
	if o == nil || o.Given == nil {
		var ret string
		return ret
	}
	return *o.Given
}

// GetGivenOk returns a tuple with the Given field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserName) GetGivenOk() (*string, bool) {
	if o == nil || o.Given == nil {
		return nil, false
	}
	return o.Given, true
}

// HasGiven returns a boolean if a field has been set.
func (o *UserName) HasGiven() bool {
	if o != nil && o.Given != nil {
		return true
	}

	return false
}

// SetGiven gets a reference to the given string and assigns it to the Given field.
func (o *UserName) SetGiven(v string) {
	o.Given = &v
}

// GetHonorificPrefix returns the HonorificPrefix field value if set, zero value otherwise.
func (o *UserName) GetHonorificPrefix() string {
	if o == nil || o.HonorificPrefix == nil {
		var ret string
		return ret
	}
	return *o.HonorificPrefix
}

// GetHonorificPrefixOk returns a tuple with the HonorificPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserName) GetHonorificPrefixOk() (*string, bool) {
	if o == nil || o.HonorificPrefix == nil {
		return nil, false
	}
	return o.HonorificPrefix, true
}

// HasHonorificPrefix returns a boolean if a field has been set.
func (o *UserName) HasHonorificPrefix() bool {
	if o != nil && o.HonorificPrefix != nil {
		return true
	}

	return false
}

// SetHonorificPrefix gets a reference to the given string and assigns it to the HonorificPrefix field.
func (o *UserName) SetHonorificPrefix(v string) {
	o.HonorificPrefix = &v
}

// GetHonorificSuffix returns the HonorificSuffix field value if set, zero value otherwise.
func (o *UserName) GetHonorificSuffix() string {
	if o == nil || o.HonorificSuffix == nil {
		var ret string
		return ret
	}
	return *o.HonorificSuffix
}

// GetHonorificSuffixOk returns a tuple with the HonorificSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserName) GetHonorificSuffixOk() (*string, bool) {
	if o == nil || o.HonorificSuffix == nil {
		return nil, false
	}
	return o.HonorificSuffix, true
}

// HasHonorificSuffix returns a boolean if a field has been set.
func (o *UserName) HasHonorificSuffix() bool {
	if o != nil && o.HonorificSuffix != nil {
		return true
	}

	return false
}

// SetHonorificSuffix gets a reference to the given string and assigns it to the HonorificSuffix field.
func (o *UserName) SetHonorificSuffix(v string) {
	o.HonorificSuffix = &v
}

// GetMiddle returns the Middle field value if set, zero value otherwise.
func (o *UserName) GetMiddle() string {
	if o == nil || o.Middle == nil {
		var ret string
		return ret
	}
	return *o.Middle
}

// GetMiddleOk returns a tuple with the Middle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserName) GetMiddleOk() (*string, bool) {
	if o == nil || o.Middle == nil {
		return nil, false
	}
	return o.Middle, true
}

// HasMiddle returns a boolean if a field has been set.
func (o *UserName) HasMiddle() bool {
	if o != nil && o.Middle != nil {
		return true
	}

	return false
}

// SetMiddle gets a reference to the given string and assigns it to the Middle field.
func (o *UserName) SetMiddle(v string) {
	o.Middle = &v
}

func (o UserName) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Family != nil {
		toSerialize["family"] = o.Family
	}
	if o.Formatted != nil {
		toSerialize["formatted"] = o.Formatted
	}
	if o.Given != nil {
		toSerialize["given"] = o.Given
	}
	if o.HonorificPrefix != nil {
		toSerialize["honorificPrefix"] = o.HonorificPrefix
	}
	if o.HonorificSuffix != nil {
		toSerialize["honorificSuffix"] = o.HonorificSuffix
	}
	if o.Middle != nil {
		toSerialize["middle"] = o.Middle
	}
	return json.Marshal(toSerialize)
}

type NullableUserName struct {
	value *UserName
	isSet bool
}

func (v NullableUserName) Get() *UserName {
	return v.value
}

func (v *NullableUserName) Set(val *UserName) {
	v.value = val
	v.isSet = true
}

func (v NullableUserName) IsSet() bool {
	return v.isSet
}

func (v *NullableUserName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserName(val *UserName) *NullableUserName {
	return &NullableUserName{value: val, isSet: true}
}

func (v NullableUserName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

