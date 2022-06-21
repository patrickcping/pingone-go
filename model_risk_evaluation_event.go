/*
PingOne Platform API - Management

A bare-bones collection for the PingOne API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pingone

import (
	"encoding/json"
)

// RiskEvaluationEvent struct for RiskEvaluationEvent
type RiskEvaluationEvent struct {
	Browser *RiskEvaluationEventBrowser `json:"browser,omitempty"`
	// A string that specifies the state of the transaction. Options are FAILED, IN_PROGRESS, and SUCCESS. If a value is not provided, the value defaults to IN_PROGRESS. The value of this property can be changed only if its current state is IN_PROGRESS.
	CompletionStatus *string `json:"completionStatus,omitempty"`
	EvaluatedFactors *RiskEvaluationEventEvaluatedFactors `json:"evaluatedFactors,omitempty"`
	// A string that specifies the origin IP address of the authentication flow. This is a required property.
	Ip string `json:"ip"`
	Flow *RiskEvaluationEventFlow `json:"flow,omitempty"`
	// A string that specifies the calling service.
	Origin *string `json:"origin,omitempty"`
	Session *RiskEvaluationEventSession `json:"session,omitempty"`
	TargetResource *RiskEvaluationEventTargetResource `json:"targetResource,omitempty"`
	User RiskEvaluationEventUser `json:"user"`
	// A string that specifies the device sharing type. Options are UNSPECIFIED, SHARED, and PRIVATE.
	SharingType *string `json:"sharingType,omitempty"`
}

// NewRiskEvaluationEvent instantiates a new RiskEvaluationEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskEvaluationEvent(ip string, user RiskEvaluationEventUser) *RiskEvaluationEvent {
	this := RiskEvaluationEvent{}
	this.Ip = ip
	this.User = user
	return &this
}

// NewRiskEvaluationEventWithDefaults instantiates a new RiskEvaluationEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskEvaluationEventWithDefaults() *RiskEvaluationEvent {
	this := RiskEvaluationEvent{}
	return &this
}

// GetBrowser returns the Browser field value if set, zero value otherwise.
func (o *RiskEvaluationEvent) GetBrowser() RiskEvaluationEventBrowser {
	if o == nil || o.Browser == nil {
		var ret RiskEvaluationEventBrowser
		return ret
	}
	return *o.Browser
}

// GetBrowserOk returns a tuple with the Browser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEvent) GetBrowserOk() (*RiskEvaluationEventBrowser, bool) {
	if o == nil || o.Browser == nil {
		return nil, false
	}
	return o.Browser, true
}

// HasBrowser returns a boolean if a field has been set.
func (o *RiskEvaluationEvent) HasBrowser() bool {
	if o != nil && o.Browser != nil {
		return true
	}

	return false
}

// SetBrowser gets a reference to the given RiskEvaluationEventBrowser and assigns it to the Browser field.
func (o *RiskEvaluationEvent) SetBrowser(v RiskEvaluationEventBrowser) {
	o.Browser = &v
}

// GetCompletionStatus returns the CompletionStatus field value if set, zero value otherwise.
func (o *RiskEvaluationEvent) GetCompletionStatus() string {
	if o == nil || o.CompletionStatus == nil {
		var ret string
		return ret
	}
	return *o.CompletionStatus
}

// GetCompletionStatusOk returns a tuple with the CompletionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEvent) GetCompletionStatusOk() (*string, bool) {
	if o == nil || o.CompletionStatus == nil {
		return nil, false
	}
	return o.CompletionStatus, true
}

// HasCompletionStatus returns a boolean if a field has been set.
func (o *RiskEvaluationEvent) HasCompletionStatus() bool {
	if o != nil && o.CompletionStatus != nil {
		return true
	}

	return false
}

// SetCompletionStatus gets a reference to the given string and assigns it to the CompletionStatus field.
func (o *RiskEvaluationEvent) SetCompletionStatus(v string) {
	o.CompletionStatus = &v
}

// GetEvaluatedFactors returns the EvaluatedFactors field value if set, zero value otherwise.
func (o *RiskEvaluationEvent) GetEvaluatedFactors() RiskEvaluationEventEvaluatedFactors {
	if o == nil || o.EvaluatedFactors == nil {
		var ret RiskEvaluationEventEvaluatedFactors
		return ret
	}
	return *o.EvaluatedFactors
}

// GetEvaluatedFactorsOk returns a tuple with the EvaluatedFactors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEvent) GetEvaluatedFactorsOk() (*RiskEvaluationEventEvaluatedFactors, bool) {
	if o == nil || o.EvaluatedFactors == nil {
		return nil, false
	}
	return o.EvaluatedFactors, true
}

// HasEvaluatedFactors returns a boolean if a field has been set.
func (o *RiskEvaluationEvent) HasEvaluatedFactors() bool {
	if o != nil && o.EvaluatedFactors != nil {
		return true
	}

	return false
}

// SetEvaluatedFactors gets a reference to the given RiskEvaluationEventEvaluatedFactors and assigns it to the EvaluatedFactors field.
func (o *RiskEvaluationEvent) SetEvaluatedFactors(v RiskEvaluationEventEvaluatedFactors) {
	o.EvaluatedFactors = &v
}

// GetIp returns the Ip field value
func (o *RiskEvaluationEvent) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEvent) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *RiskEvaluationEvent) SetIp(v string) {
	o.Ip = v
}

// GetFlow returns the Flow field value if set, zero value otherwise.
func (o *RiskEvaluationEvent) GetFlow() RiskEvaluationEventFlow {
	if o == nil || o.Flow == nil {
		var ret RiskEvaluationEventFlow
		return ret
	}
	return *o.Flow
}

// GetFlowOk returns a tuple with the Flow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEvent) GetFlowOk() (*RiskEvaluationEventFlow, bool) {
	if o == nil || o.Flow == nil {
		return nil, false
	}
	return o.Flow, true
}

// HasFlow returns a boolean if a field has been set.
func (o *RiskEvaluationEvent) HasFlow() bool {
	if o != nil && o.Flow != nil {
		return true
	}

	return false
}

// SetFlow gets a reference to the given RiskEvaluationEventFlow and assigns it to the Flow field.
func (o *RiskEvaluationEvent) SetFlow(v RiskEvaluationEventFlow) {
	o.Flow = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *RiskEvaluationEvent) GetOrigin() string {
	if o == nil || o.Origin == nil {
		var ret string
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEvent) GetOriginOk() (*string, bool) {
	if o == nil || o.Origin == nil {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *RiskEvaluationEvent) HasOrigin() bool {
	if o != nil && o.Origin != nil {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given string and assigns it to the Origin field.
func (o *RiskEvaluationEvent) SetOrigin(v string) {
	o.Origin = &v
}

// GetSession returns the Session field value if set, zero value otherwise.
func (o *RiskEvaluationEvent) GetSession() RiskEvaluationEventSession {
	if o == nil || o.Session == nil {
		var ret RiskEvaluationEventSession
		return ret
	}
	return *o.Session
}

// GetSessionOk returns a tuple with the Session field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEvent) GetSessionOk() (*RiskEvaluationEventSession, bool) {
	if o == nil || o.Session == nil {
		return nil, false
	}
	return o.Session, true
}

// HasSession returns a boolean if a field has been set.
func (o *RiskEvaluationEvent) HasSession() bool {
	if o != nil && o.Session != nil {
		return true
	}

	return false
}

// SetSession gets a reference to the given RiskEvaluationEventSession and assigns it to the Session field.
func (o *RiskEvaluationEvent) SetSession(v RiskEvaluationEventSession) {
	o.Session = &v
}

// GetTargetResource returns the TargetResource field value if set, zero value otherwise.
func (o *RiskEvaluationEvent) GetTargetResource() RiskEvaluationEventTargetResource {
	if o == nil || o.TargetResource == nil {
		var ret RiskEvaluationEventTargetResource
		return ret
	}
	return *o.TargetResource
}

// GetTargetResourceOk returns a tuple with the TargetResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEvent) GetTargetResourceOk() (*RiskEvaluationEventTargetResource, bool) {
	if o == nil || o.TargetResource == nil {
		return nil, false
	}
	return o.TargetResource, true
}

// HasTargetResource returns a boolean if a field has been set.
func (o *RiskEvaluationEvent) HasTargetResource() bool {
	if o != nil && o.TargetResource != nil {
		return true
	}

	return false
}

// SetTargetResource gets a reference to the given RiskEvaluationEventTargetResource and assigns it to the TargetResource field.
func (o *RiskEvaluationEvent) SetTargetResource(v RiskEvaluationEventTargetResource) {
	o.TargetResource = &v
}

// GetUser returns the User field value
func (o *RiskEvaluationEvent) GetUser() RiskEvaluationEventUser {
	if o == nil {
		var ret RiskEvaluationEventUser
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEvent) GetUserOk() (*RiskEvaluationEventUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *RiskEvaluationEvent) SetUser(v RiskEvaluationEventUser) {
	o.User = v
}

// GetSharingType returns the SharingType field value if set, zero value otherwise.
func (o *RiskEvaluationEvent) GetSharingType() string {
	if o == nil || o.SharingType == nil {
		var ret string
		return ret
	}
	return *o.SharingType
}

// GetSharingTypeOk returns a tuple with the SharingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEvent) GetSharingTypeOk() (*string, bool) {
	if o == nil || o.SharingType == nil {
		return nil, false
	}
	return o.SharingType, true
}

// HasSharingType returns a boolean if a field has been set.
func (o *RiskEvaluationEvent) HasSharingType() bool {
	if o != nil && o.SharingType != nil {
		return true
	}

	return false
}

// SetSharingType gets a reference to the given string and assigns it to the SharingType field.
func (o *RiskEvaluationEvent) SetSharingType(v string) {
	o.SharingType = &v
}

func (o RiskEvaluationEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Browser != nil {
		toSerialize["browser"] = o.Browser
	}
	if o.CompletionStatus != nil {
		toSerialize["completionStatus"] = o.CompletionStatus
	}
	if o.EvaluatedFactors != nil {
		toSerialize["evaluatedFactors"] = o.EvaluatedFactors
	}
	if true {
		toSerialize["ip"] = o.Ip
	}
	if o.Flow != nil {
		toSerialize["flow"] = o.Flow
	}
	if o.Origin != nil {
		toSerialize["origin"] = o.Origin
	}
	if o.Session != nil {
		toSerialize["session"] = o.Session
	}
	if o.TargetResource != nil {
		toSerialize["targetResource"] = o.TargetResource
	}
	if true {
		toSerialize["user"] = o.User
	}
	if o.SharingType != nil {
		toSerialize["sharingType"] = o.SharingType
	}
	return json.Marshal(toSerialize)
}

type NullableRiskEvaluationEvent struct {
	value *RiskEvaluationEvent
	isSet bool
}

func (v NullableRiskEvaluationEvent) Get() *RiskEvaluationEvent {
	return v.value
}

func (v *NullableRiskEvaluationEvent) Set(val *RiskEvaluationEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskEvaluationEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskEvaluationEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskEvaluationEvent(val *RiskEvaluationEvent) *NullableRiskEvaluationEvent {
	return &NullableRiskEvaluationEvent{value: val, isSet: true}
}

func (v NullableRiskEvaluationEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskEvaluationEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


