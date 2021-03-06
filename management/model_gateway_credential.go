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

// GatewayCredential struct for GatewayCredential
type GatewayCredential struct {
	// A string that specifies the auto-generated ID for this credential. This is the JWT's jti claim. This is a required property.
	Id *string `json:"id,omitempty"`
	// A date that specifies the date the credential was created in Coordinated Universal Time (UTC). This is a required property.
	CreatedAt *string `json:"createdAt,omitempty"`
	// A date that specifies the date the credential was created in Coordinated Universal Time (UTC). This is a required property.
	GatewayType *string `json:"gatewayType,omitempty"`
	// A date that specifies the date the credential was last used in UTC. This is a required property.
	LastUsedAt *string `json:"lastUsedAt,omitempty"`
	ConsoleUrl *string `json:"consoleUrl,omitempty"`
	ApiUrl *string `json:"apiUrl,omitempty"`
	AuthUrl *string `json:"authUrl,omitempty"`
	// A string that specifies the signed JWT for the gateway credential. This property is present only when the gateway credential is created.
	Credential *string `json:"credential,omitempty"`
}

// NewGatewayCredential instantiates a new GatewayCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayCredential() *GatewayCredential {
	this := GatewayCredential{}
	return &this
}

// NewGatewayCredentialWithDefaults instantiates a new GatewayCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayCredentialWithDefaults() *GatewayCredential {
	this := GatewayCredential{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GatewayCredential) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCredential) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GatewayCredential) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GatewayCredential) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GatewayCredential) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCredential) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GatewayCredential) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GatewayCredential) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetGatewayType returns the GatewayType field value if set, zero value otherwise.
func (o *GatewayCredential) GetGatewayType() string {
	if o == nil || o.GatewayType == nil {
		var ret string
		return ret
	}
	return *o.GatewayType
}

// GetGatewayTypeOk returns a tuple with the GatewayType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCredential) GetGatewayTypeOk() (*string, bool) {
	if o == nil || o.GatewayType == nil {
		return nil, false
	}
	return o.GatewayType, true
}

// HasGatewayType returns a boolean if a field has been set.
func (o *GatewayCredential) HasGatewayType() bool {
	if o != nil && o.GatewayType != nil {
		return true
	}

	return false
}

// SetGatewayType gets a reference to the given string and assigns it to the GatewayType field.
func (o *GatewayCredential) SetGatewayType(v string) {
	o.GatewayType = &v
}

// GetLastUsedAt returns the LastUsedAt field value if set, zero value otherwise.
func (o *GatewayCredential) GetLastUsedAt() string {
	if o == nil || o.LastUsedAt == nil {
		var ret string
		return ret
	}
	return *o.LastUsedAt
}

// GetLastUsedAtOk returns a tuple with the LastUsedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCredential) GetLastUsedAtOk() (*string, bool) {
	if o == nil || o.LastUsedAt == nil {
		return nil, false
	}
	return o.LastUsedAt, true
}

// HasLastUsedAt returns a boolean if a field has been set.
func (o *GatewayCredential) HasLastUsedAt() bool {
	if o != nil && o.LastUsedAt != nil {
		return true
	}

	return false
}

// SetLastUsedAt gets a reference to the given string and assigns it to the LastUsedAt field.
func (o *GatewayCredential) SetLastUsedAt(v string) {
	o.LastUsedAt = &v
}

// GetConsoleUrl returns the ConsoleUrl field value if set, zero value otherwise.
func (o *GatewayCredential) GetConsoleUrl() string {
	if o == nil || o.ConsoleUrl == nil {
		var ret string
		return ret
	}
	return *o.ConsoleUrl
}

// GetConsoleUrlOk returns a tuple with the ConsoleUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCredential) GetConsoleUrlOk() (*string, bool) {
	if o == nil || o.ConsoleUrl == nil {
		return nil, false
	}
	return o.ConsoleUrl, true
}

// HasConsoleUrl returns a boolean if a field has been set.
func (o *GatewayCredential) HasConsoleUrl() bool {
	if o != nil && o.ConsoleUrl != nil {
		return true
	}

	return false
}

// SetConsoleUrl gets a reference to the given string and assigns it to the ConsoleUrl field.
func (o *GatewayCredential) SetConsoleUrl(v string) {
	o.ConsoleUrl = &v
}

// GetApiUrl returns the ApiUrl field value if set, zero value otherwise.
func (o *GatewayCredential) GetApiUrl() string {
	if o == nil || o.ApiUrl == nil {
		var ret string
		return ret
	}
	return *o.ApiUrl
}

// GetApiUrlOk returns a tuple with the ApiUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCredential) GetApiUrlOk() (*string, bool) {
	if o == nil || o.ApiUrl == nil {
		return nil, false
	}
	return o.ApiUrl, true
}

// HasApiUrl returns a boolean if a field has been set.
func (o *GatewayCredential) HasApiUrl() bool {
	if o != nil && o.ApiUrl != nil {
		return true
	}

	return false
}

// SetApiUrl gets a reference to the given string and assigns it to the ApiUrl field.
func (o *GatewayCredential) SetApiUrl(v string) {
	o.ApiUrl = &v
}

// GetAuthUrl returns the AuthUrl field value if set, zero value otherwise.
func (o *GatewayCredential) GetAuthUrl() string {
	if o == nil || o.AuthUrl == nil {
		var ret string
		return ret
	}
	return *o.AuthUrl
}

// GetAuthUrlOk returns a tuple with the AuthUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCredential) GetAuthUrlOk() (*string, bool) {
	if o == nil || o.AuthUrl == nil {
		return nil, false
	}
	return o.AuthUrl, true
}

// HasAuthUrl returns a boolean if a field has been set.
func (o *GatewayCredential) HasAuthUrl() bool {
	if o != nil && o.AuthUrl != nil {
		return true
	}

	return false
}

// SetAuthUrl gets a reference to the given string and assigns it to the AuthUrl field.
func (o *GatewayCredential) SetAuthUrl(v string) {
	o.AuthUrl = &v
}

// GetCredential returns the Credential field value if set, zero value otherwise.
func (o *GatewayCredential) GetCredential() string {
	if o == nil || o.Credential == nil {
		var ret string
		return ret
	}
	return *o.Credential
}

// GetCredentialOk returns a tuple with the Credential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCredential) GetCredentialOk() (*string, bool) {
	if o == nil || o.Credential == nil {
		return nil, false
	}
	return o.Credential, true
}

// HasCredential returns a boolean if a field has been set.
func (o *GatewayCredential) HasCredential() bool {
	if o != nil && o.Credential != nil {
		return true
	}

	return false
}

// SetCredential gets a reference to the given string and assigns it to the Credential field.
func (o *GatewayCredential) SetCredential(v string) {
	o.Credential = &v
}

func (o GatewayCredential) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.GatewayType != nil {
		toSerialize["gatewayType"] = o.GatewayType
	}
	if o.LastUsedAt != nil {
		toSerialize["lastUsedAt"] = o.LastUsedAt
	}
	if o.ConsoleUrl != nil {
		toSerialize["consoleUrl"] = o.ConsoleUrl
	}
	if o.ApiUrl != nil {
		toSerialize["apiUrl"] = o.ApiUrl
	}
	if o.AuthUrl != nil {
		toSerialize["authUrl"] = o.AuthUrl
	}
	if o.Credential != nil {
		toSerialize["credential"] = o.Credential
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayCredential struct {
	value *GatewayCredential
	isSet bool
}

func (v NullableGatewayCredential) Get() *GatewayCredential {
	return v.value
}

func (v *NullableGatewayCredential) Set(val *GatewayCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayCredential(val *GatewayCredential) *NullableGatewayCredential {
	return &NullableGatewayCredential{value: val, isSet: true}
}

func (v NullableGatewayCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


