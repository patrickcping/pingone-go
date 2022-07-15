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

// SignOnPolicyAction struct for SignOnPolicyAction
type SignOnPolicyAction struct {
	Conditions *SignOnPolicyActionCommonConditions `json:"conditions,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	// A string that specifies the sign-on policy assignment resource’s unique identifier.
	Id *string `json:"id,omitempty"`
	// An integer that specifies the order in which the policy referenced by this assignment is evaluated during an authentication flow relative to other policies. An assignment with a lower priority will be evaluated first. This is a required property.
	Priority int32 `json:"priority"`
	SignOnPolicy SignOnPolicyActionCommonSignOnPolicy `json:"signOnPolicy"`
	// A string that specifies the type of action. Options are `LOGIN`, `MULTI_FACTOR_AUTHENTICATION`, `IDENTIFIER_FIRST`, `IDENTITY_PROVIDER` `AGREEMENT` and `PROGRESSIVE_PROFILING`.
	Type string `json:"type"`
	// A boolean that specifies whether users must confirm data returned from an identity provider prior to registration. Users can modify the data and omit non-required attributes. Modified attributes are added to the user's profile during account creation. This is an optional property. If omitted, the default value is set to false.
	ConfirmIdentityProviderAttributes *bool `json:"confirmIdentityProviderAttributes,omitempty"`
	// A boolean that if set to true and if the user's account is locked (the account.canAuthenticate attribute is set to false), then social sign on with an external identity provider is prevented.
	EnforceLockoutForIdentityProviders *bool `json:"enforceLockoutForIdentityProviders,omitempty"`
	Recovery *SignOnPolicyActionLoginRecovery `json:"recovery,omitempty"`
	Registration *SignOnPolicyActionIDPRegistration `json:"registration,omitempty"`
	// An array of strings that specifies the IDs of the identity providers that can be used for the social login sign-on flow.
	SocialProviders []SignOnPolicyActionLoginSocialProvidersInner `json:"socialProviders,omitempty"`
	Authenticator *SignOnPolicyActionMFAAuthenticator `json:"authenticator,omitempty"`
	BoundBiometrics *SignOnPolicyActionMFABoundBiometrics `json:"boundBiometrics,omitempty"`
	Email *SignOnPolicyActionMFAEmail `json:"email,omitempty"`
	SecurityKey *SignOnPolicyActionMFASecurityKey `json:"securityKey,omitempty"`
	Sms *SignOnPolicyActionMFASms `json:"sms,omitempty"`
	Voice *SignOnPolicyActionMFAVoice `json:"voice,omitempty"`
	// The applications collection specifies all the native native applications that are allowed in the sign-on policy action.  If the applications collection is empty, a push notification is not allowed for the action.
	Applications []SignOnPolicyActionMFAApplicationsInner `json:"applications,omitempty"`
	// A string that specifies the device mode for the MFA flow. Options are `BYPASS` to allow MFA without a specified device, or `BLOCK` to block the MFA flow if no device is specified. To use this configuration option, the authorize request must include a signed `login_hint_token` property. For more information, see Authorize (Browserless and MFA Only Flows)
	NoDeviceMode *string `json:"noDeviceMode,omitempty"`
	// The list of IDP discovery rules that are evaluated in order when no user is associated with the user identifier. The maximum number of rules is 100. The condition on which this identity provider is used to authenticate the user is expressed using the PingOne policy condition language
	DiscoveryRules []SignOnPolicyActionIDFirstDiscoveryRulesInner `json:"discoveryRules,omitempty"`
	// A string that designates the sign-on policies included in the authorization flow request. Options can include the PingOne predefined sign-on policies, Single_Factor and Multi_Factor, or any custom defined sign-on policy names. Sign-on policy names should be listed in order of preference, and they must be assigned to the application. This property can be configured on the identity provider action and is passed to the identity provider if the identity provider is of type `SAML` or `OPENID_CONNECT`.
	AcrValues *string `json:"acrValues,omitempty"`
	IdentityProvider SignOnPolicyActionIDPIdentityProvider `json:"identityProvider"`
	// A boolean that specifies whether to pass in a login hint to the identity provider on the authentication request. Based on user context, the login hint is set if (1) the user is set on the flow, and (2) the user already has an account link for the identity provider. If both of these conditions are true, then the user is sent to the identity provider with a login hint equal to their externalId for the identity provider (saved on the account link). If these conditions are not true, then the API checks see if there is an OIDC login hint on the flow. If so, that login hint is used. If none of these conditions are true, the login hint parameter is not included on the authorization request to the identity provider.
	PassUserContext *bool `json:"passUserContext,omitempty"`
	Agreement SignOnPolicyActionAgreementAgreement `json:"agreement"`
	Attributes SignOnPolicyActionProgressiveProfilingAttributes `json:"attributes"`
	// A boolean that specifies whether the progressive profiling action will not be executed if another progressive profiling action has already been executed during the flow. This property is required.
	PreventMultiplePromptsPerFlow bool `json:"preventMultiplePromptsPerFlow"`
	// An integer that specifies how often to prompt the user to provide profile data for the configured attributes for which they do not have values. This property is required.
	PromptIntervalSeconds int32 `json:"promptIntervalSeconds"`
	// A string that specifies text to display to the user when prompting for attribute values. This property is required.
	PromptText string `json:"promptText"`
}

// NewSignOnPolicyAction instantiates a new SignOnPolicyAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyAction(priority int32, signOnPolicy SignOnPolicyActionCommonSignOnPolicy, type_ string, identityProvider SignOnPolicyActionIDPIdentityProvider, agreement SignOnPolicyActionAgreementAgreement, attributes SignOnPolicyActionProgressiveProfilingAttributes, preventMultiplePromptsPerFlow bool, promptIntervalSeconds int32, promptText string) *SignOnPolicyAction {
	this := SignOnPolicyAction{}
	this.Priority = priority
	this.SignOnPolicy = signOnPolicy
	this.Type = type_
	var confirmIdentityProviderAttributes bool = false
	this.ConfirmIdentityProviderAttributes = &confirmIdentityProviderAttributes
	this.IdentityProvider = identityProvider
	this.Agreement = agreement
	this.Attributes = attributes
	this.PreventMultiplePromptsPerFlow = preventMultiplePromptsPerFlow
	this.PromptIntervalSeconds = promptIntervalSeconds
	this.PromptText = promptText
	return &this
}

// NewSignOnPolicyActionWithDefaults instantiates a new SignOnPolicyAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionWithDefaults() *SignOnPolicyAction {
	this := SignOnPolicyAction{}
	var confirmIdentityProviderAttributes bool = false
	this.ConfirmIdentityProviderAttributes = &confirmIdentityProviderAttributes
	return &this
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *SignOnPolicyAction) GetConditions() SignOnPolicyActionCommonConditions {
	if o == nil || o.Conditions == nil {
		var ret SignOnPolicyActionCommonConditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyAction) GetConditionsOk() (*SignOnPolicyActionCommonConditions, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *SignOnPolicyAction) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given SignOnPolicyActionCommonConditions and assigns it to the Conditions field.
func (o *SignOnPolicyAction) SetConditions(v SignOnPolicyActionCommonConditions) {
	o.Conditions = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *SignOnPolicyAction) GetEnvironment() ObjectEnvironment {
	if o == nil || o.Environment == nil {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyAction) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *SignOnPolicyAction) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *SignOnPolicyAction) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SignOnPolicyAction) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyAction) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SignOnPolicyAction) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SignOnPolicyAction) SetId(v string) {
	o.Id = &v
}

// GetPriority returns the Priority field value
func (o *SignOnPolicyAction) GetPriority() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyAction) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *SignOnPolicyAction) SetPriority(v int32) {
	o.Priority = v
}

// GetSignOnPolicy returns the SignOnPolicy field value
func (o *SignOnPolicyAction) GetSignOnPolicy() SignOnPolicyActionCommonSignOnPolicy {
	if o == nil {
		var ret SignOnPolicyActionCommonSignOnPolicy
		return ret
	}

	return o.SignOnPolicy
}

// GetSignOnPolicyOk returns a tuple with the SignOnPolicy field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyAction) GetSignOnPolicyOk() (*SignOnPolicyActionCommonSignOnPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignOnPolicy, true
}

// SetSignOnPolicy sets field value
func (o *SignOnPolicyAction) SetSignOnPolicy(v SignOnPolicyActionCommonSignOnPolicy) {
	o.SignOnPolicy = v
}

// GetType returns the Type field value
func (o *SignOnPolicyAction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyAction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SignOnPolicyAction) SetType(v string) {
	o.Type = v
}

// GetConfirmIdentityProviderAttributes returns the ConfirmIdentityProviderAttributes field value if set, zero value otherwise.
func (o *SignOnPolicyAction) GetConfirmIdentityProviderAttributes() bool {
	if o == nil || o.ConfirmIdentityProviderAttributes == nil {
		var ret bool
		return ret
	}
	return *o.ConfirmIdentityProviderAttributes
}

// GetConfirmIdentityProviderAttributesOk returns a tuple with the ConfirmIdentityProviderAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyAction) GetConfirmIdentityProviderAttributesOk() (*bool, bool) {
	if o == nil || o.ConfirmIdentityProviderAttributes == nil {
		return nil, false
	}
	return o.ConfirmIdentityProviderAttributes, true
}

// HasConfirmIdentityProviderAttributes returns a boolean if a field has been set.
func (o *SignOnPolicyAction) HasConfirmIdentityProviderAttributes() bool {
	if o != nil && o.ConfirmIdentityProviderAttributes != nil {
		return true
	}

	return false
}

// SetConfirmIdentityProviderAttributes gets a reference to the given bool and assigns it to the ConfirmIdentityProviderAttributes field.
func (o *SignOnPolicyAction) SetConfirmIdentityProviderAttributes(v bool) {
	o.ConfirmIdentityProviderAttributes = &v
}

// GetEnforceLockoutForIdentityProviders returns the EnforceLockoutForIdentityProviders field value if set, zero value otherwise.
func (o *SignOnPolicyAction) GetEnforceLockoutForIdentityProviders() bool {
	if o == nil || o.EnforceLockoutForIdentityProviders == nil {
		var ret bool
		return ret
	}
	return *o.EnforceLockoutForIdentityProviders
}

// GetEnforceLockoutForIdentityProvidersOk returns a tuple with the EnforceLockoutForIdentityProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyAction) GetEnforceLockoutForIdentityProvidersOk() (*bool, bool) {
	if o == nil || o.EnforceLockoutForIdentityProviders == nil {
		return nil, false
	}
	return o.EnforceLockoutForIdentityProviders, true
}

// HasEnforceLockoutForIdentityProviders returns a boolean if a field has been set.
func (o *SignOnPolicyAction) HasEnforceLockoutForIdentityProviders() bool {
	if o != nil && o.EnforceLockoutForIdentityProviders != nil {
		return true
	}

	return false
}

// SetEnforceLockoutForIdentityProviders gets a reference to the given bool and assigns it to the EnforceLockoutForIdentityProviders field.
func (o *SignOnPolicyAction) SetEnforceLockoutForIdentityProviders(v bool) {
	o.EnforceLockoutForIdentityProviders = &v
}

// GetRecovery returns the Recovery field value if set, zero value otherwise.
func (o *SignOnPolicyAction) GetRecovery() SignOnPolicyActionLoginRecovery {
	if o == nil || o.Recovery == nil {
		var ret SignOnPolicyActionLoginRecovery
		return ret
	}
	return *o.Recovery
}

// GetRecoveryOk returns a tuple with the Recovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyAction) GetRecoveryOk() (*SignOnPolicyActionLoginRecovery, bool) {
	if o == nil || o.Recovery == nil {
		return nil, false
	}
	return o.Recovery, true
}

// HasRecovery returns a boolean if a field has been set.
func (o *SignOnPolicyAction) HasRecovery() bool {
	if o != nil && o.Recovery != nil {
		return true
	}

	return false
}

// SetRecovery gets a reference to the given SignOnPolicyActionLoginRecovery and assigns it to the Recovery field.
func (o *SignOnPolicyAction) SetRecovery(v SignOnPolicyActionLoginRecovery) {
	o.Recovery = &v
}

// GetRegistration returns the Registration field value if set, zero value otherwise.
func (o *SignOnPolicyAction) GetRegistration() SignOnPolicyActionIDPRegistration {
	if o == nil || o.Registration == nil {
		var ret SignOnPolicyActionIDPRegistration
		return ret
	}
	return *o.Registration
}

// GetRegistrationOk returns a tuple with the Registration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyAction) GetRegistrationOk() (*SignOnPolicyActionIDPRegistration, bool) {
	if o == nil || o.Registration == nil {
		return nil, false
	}
	return o.Registration, true
}

// HasRegistration returns a boolean if a field has been set.
func (o *SignOnPolicyAction) HasRegistration() bool {
	if o != nil && o.Registration != nil {
		return true
	}

	return false
}

// SetRegistration gets a reference to the given SignOnPolicyActionIDPRegistration and assigns it to the Registration field.
func (o *SignOnPolicyAction) SetRegistration(v SignOnPolicyActionIDPRegistration) {
	o.Registration = &v
}

// GetSocialProviders returns the SocialProviders field value if set, zero value otherwise.
func (o *SignOnPolicyAction) GetSocialProviders() []SignOnPolicyActionLoginSocialProvidersInner {
	if o == nil || o.SocialProviders == nil {
		var ret []SignOnPolicyActionLoginSocialProvidersInner
		return ret
	}
	return o.SocialProviders
}

// GetSocialProvidersOk returns a tuple with the SocialProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyAction) GetSocialProvidersOk() ([]SignOnPolicyActionLoginSocialProvidersInner, bool) {
	if o == nil || o.SocialProviders == nil {
		return nil, false
	}
	return o.SocialProviders, true
}

// HasSocialProviders returns a boolean if a field has been set.
func (o *SignOnPolicyAction) HasSocialProviders() bool {
	if o != nil && o.SocialProviders != nil {
		return true
	}

	return false
}

// SetSocialProviders gets a reference to the given []SignOnPolicyActionLoginSocialProvidersInner and assigns it to the SocialProviders field.
func (o *SignOnPolicyAction) SetSocialProviders(v []SignOnPolicyActionLoginSocialProvidersInner) {
	o.SocialProviders = v
}

// GetAuthenticator returns the Authenticator field value if set, zero value otherwise.
func (o *SignOnPolicyAction) GetAuthenticator() SignOnPolicyActionMFAAuthenticator {
	if o == nil || o.Authenticator == nil {
		var ret SignOnPolicyActionMFAAuthenticator
		return ret
	}
	return *o.Authenticator
}

// GetAuthenticatorOk returns a tuple with the Authenticator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyAction) GetAuthenticatorOk() (*SignOnPolicyActionMFAAuthenticator, bool) {
	if o == nil || o.Authenticator == nil {
		return nil, false
	}
	return o.Authenticator, true
}

// HasAuthenticator returns a boolean if a field has been set.
func (o *SignOnPolicyAction) HasAuthenticator() bool {
	if o != nil && o.Authenticator != nil {
		return true
	}

	return false
}

// SetAuthenticator gets a reference to the given SignOnPolicyActionMFAAuthenticator and assigns it to the Authenticator field.
func (o *SignOnPolicyAction) SetAuthenticator(v SignOnPolicyActionMFAAuthenticator) {
	o.Authenticator = &v
}

// GetBoundBiometrics returns the BoundBiometrics field value if set, zero value otherwise.
func (o *SignOnPolicyAction) GetBoundBiometrics() SignOnPolicyActionMFABoundBiometrics {
	if o == nil || o.BoundBiometrics == nil {
		var ret SignOnPolicyActionMFABoundBiometrics
		return ret
	}
	return *o.BoundBiometrics
}

// GetBoundBiometricsOk returns a tuple with the BoundBiometrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyAction) GetBoundBiometricsOk() (*SignOnPolicyActionMFABoundBiometrics, bool) {
	if o == nil || o.BoundBiometrics == nil {
		return nil, false
	}
	return o.BoundBiometrics, true
}

// HasBoundBiometrics returns a boolean if a field has been set.
func (o *SignOnPolicyAction) HasBoundBiometrics() bool {
	if o != nil && o.BoundBiometrics != nil {
		return true
	}

	return false
}

// SetBoundBiometrics gets a reference to the given SignOnPolicyActionMFABoundBiometrics and assigns it to the BoundBiometrics field.
func (o *SignOnPolicyAction) SetBoundBiometrics(v SignOnPolicyActionMFABoundBiometrics) {
	o.BoundBiometrics = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *SignOnPolicyAction) GetEmail() SignOnPolicyActionMFAEmail {
	if o == nil || o.Email == nil {
		var ret SignOnPolicyActionMFAEmail
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyAction) GetEmailOk() (*SignOnPolicyActionMFAEmail, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *SignOnPolicyAction) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given SignOnPolicyActionMFAEmail and assigns it to the Email field.
func (o *SignOnPolicyAction) SetEmail(v SignOnPolicyActionMFAEmail) {
	o.Email = &v
}

// GetSecurityKey returns the SecurityKey field value if set, zero value otherwise.
func (o *SignOnPolicyAction) GetSecurityKey() SignOnPolicyActionMFASecurityKey {
	if o == nil || o.SecurityKey == nil {
		var ret SignOnPolicyActionMFASecurityKey
		return ret
	}
	return *o.SecurityKey
}

// GetSecurityKeyOk returns a tuple with the SecurityKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyAction) GetSecurityKeyOk() (*SignOnPolicyActionMFASecurityKey, bool) {
	if o == nil || o.SecurityKey == nil {
		return nil, false
	}
	return o.SecurityKey, true
}

// HasSecurityKey returns a boolean if a field has been set.
func (o *SignOnPolicyAction) HasSecurityKey() bool {
	if o != nil && o.SecurityKey != nil {
		return true
	}

	return false
}

// SetSecurityKey gets a reference to the given SignOnPolicyActionMFASecurityKey and assigns it to the SecurityKey field.
func (o *SignOnPolicyAction) SetSecurityKey(v SignOnPolicyActionMFASecurityKey) {
	o.SecurityKey = &v
}

// GetSms returns the Sms field value if set, zero value otherwise.
func (o *SignOnPolicyAction) GetSms() SignOnPolicyActionMFASms {
	if o == nil || o.Sms == nil {
		var ret SignOnPolicyActionMFASms
		return ret
	}
	return *o.Sms
}

// GetSmsOk returns a tuple with the Sms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyAction) GetSmsOk() (*SignOnPolicyActionMFASms, bool) {
	if o == nil || o.Sms == nil {
		return nil, false
	}
	return o.Sms, true
}

// HasSms returns a boolean if a field has been set.
func (o *SignOnPolicyAction) HasSms() bool {
	if o != nil && o.Sms != nil {
		return true
	}

	return false
}

// SetSms gets a reference to the given SignOnPolicyActionMFASms and assigns it to the Sms field.
func (o *SignOnPolicyAction) SetSms(v SignOnPolicyActionMFASms) {
	o.Sms = &v
}

// GetVoice returns the Voice field value if set, zero value otherwise.
func (o *SignOnPolicyAction) GetVoice() SignOnPolicyActionMFAVoice {
	if o == nil || o.Voice == nil {
		var ret SignOnPolicyActionMFAVoice
		return ret
	}
	return *o.Voice
}

// GetVoiceOk returns a tuple with the Voice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyAction) GetVoiceOk() (*SignOnPolicyActionMFAVoice, bool) {
	if o == nil || o.Voice == nil {
		return nil, false
	}
	return o.Voice, true
}

// HasVoice returns a boolean if a field has been set.
func (o *SignOnPolicyAction) HasVoice() bool {
	if o != nil && o.Voice != nil {
		return true
	}

	return false
}

// SetVoice gets a reference to the given SignOnPolicyActionMFAVoice and assigns it to the Voice field.
func (o *SignOnPolicyAction) SetVoice(v SignOnPolicyActionMFAVoice) {
	o.Voice = &v
}

// GetApplications returns the Applications field value if set, zero value otherwise.
func (o *SignOnPolicyAction) GetApplications() []SignOnPolicyActionMFAApplicationsInner {
	if o == nil || o.Applications == nil {
		var ret []SignOnPolicyActionMFAApplicationsInner
		return ret
	}
	return o.Applications
}

// GetApplicationsOk returns a tuple with the Applications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyAction) GetApplicationsOk() ([]SignOnPolicyActionMFAApplicationsInner, bool) {
	if o == nil || o.Applications == nil {
		return nil, false
	}
	return o.Applications, true
}

// HasApplications returns a boolean if a field has been set.
func (o *SignOnPolicyAction) HasApplications() bool {
	if o != nil && o.Applications != nil {
		return true
	}

	return false
}

// SetApplications gets a reference to the given []SignOnPolicyActionMFAApplicationsInner and assigns it to the Applications field.
func (o *SignOnPolicyAction) SetApplications(v []SignOnPolicyActionMFAApplicationsInner) {
	o.Applications = v
}

// GetNoDeviceMode returns the NoDeviceMode field value if set, zero value otherwise.
func (o *SignOnPolicyAction) GetNoDeviceMode() string {
	if o == nil || o.NoDeviceMode == nil {
		var ret string
		return ret
	}
	return *o.NoDeviceMode
}

// GetNoDeviceModeOk returns a tuple with the NoDeviceMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyAction) GetNoDeviceModeOk() (*string, bool) {
	if o == nil || o.NoDeviceMode == nil {
		return nil, false
	}
	return o.NoDeviceMode, true
}

// HasNoDeviceMode returns a boolean if a field has been set.
func (o *SignOnPolicyAction) HasNoDeviceMode() bool {
	if o != nil && o.NoDeviceMode != nil {
		return true
	}

	return false
}

// SetNoDeviceMode gets a reference to the given string and assigns it to the NoDeviceMode field.
func (o *SignOnPolicyAction) SetNoDeviceMode(v string) {
	o.NoDeviceMode = &v
}

// GetDiscoveryRules returns the DiscoveryRules field value if set, zero value otherwise.
func (o *SignOnPolicyAction) GetDiscoveryRules() []SignOnPolicyActionIDFirstDiscoveryRulesInner {
	if o == nil || o.DiscoveryRules == nil {
		var ret []SignOnPolicyActionIDFirstDiscoveryRulesInner
		return ret
	}
	return o.DiscoveryRules
}

// GetDiscoveryRulesOk returns a tuple with the DiscoveryRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyAction) GetDiscoveryRulesOk() ([]SignOnPolicyActionIDFirstDiscoveryRulesInner, bool) {
	if o == nil || o.DiscoveryRules == nil {
		return nil, false
	}
	return o.DiscoveryRules, true
}

// HasDiscoveryRules returns a boolean if a field has been set.
func (o *SignOnPolicyAction) HasDiscoveryRules() bool {
	if o != nil && o.DiscoveryRules != nil {
		return true
	}

	return false
}

// SetDiscoveryRules gets a reference to the given []SignOnPolicyActionIDFirstDiscoveryRulesInner and assigns it to the DiscoveryRules field.
func (o *SignOnPolicyAction) SetDiscoveryRules(v []SignOnPolicyActionIDFirstDiscoveryRulesInner) {
	o.DiscoveryRules = v
}

// GetAcrValues returns the AcrValues field value if set, zero value otherwise.
func (o *SignOnPolicyAction) GetAcrValues() string {
	if o == nil || o.AcrValues == nil {
		var ret string
		return ret
	}
	return *o.AcrValues
}

// GetAcrValuesOk returns a tuple with the AcrValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyAction) GetAcrValuesOk() (*string, bool) {
	if o == nil || o.AcrValues == nil {
		return nil, false
	}
	return o.AcrValues, true
}

// HasAcrValues returns a boolean if a field has been set.
func (o *SignOnPolicyAction) HasAcrValues() bool {
	if o != nil && o.AcrValues != nil {
		return true
	}

	return false
}

// SetAcrValues gets a reference to the given string and assigns it to the AcrValues field.
func (o *SignOnPolicyAction) SetAcrValues(v string) {
	o.AcrValues = &v
}

// GetIdentityProvider returns the IdentityProvider field value
func (o *SignOnPolicyAction) GetIdentityProvider() SignOnPolicyActionIDPIdentityProvider {
	if o == nil {
		var ret SignOnPolicyActionIDPIdentityProvider
		return ret
	}

	return o.IdentityProvider
}

// GetIdentityProviderOk returns a tuple with the IdentityProvider field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyAction) GetIdentityProviderOk() (*SignOnPolicyActionIDPIdentityProvider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentityProvider, true
}

// SetIdentityProvider sets field value
func (o *SignOnPolicyAction) SetIdentityProvider(v SignOnPolicyActionIDPIdentityProvider) {
	o.IdentityProvider = v
}

// GetPassUserContext returns the PassUserContext field value if set, zero value otherwise.
func (o *SignOnPolicyAction) GetPassUserContext() bool {
	if o == nil || o.PassUserContext == nil {
		var ret bool
		return ret
	}
	return *o.PassUserContext
}

// GetPassUserContextOk returns a tuple with the PassUserContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyAction) GetPassUserContextOk() (*bool, bool) {
	if o == nil || o.PassUserContext == nil {
		return nil, false
	}
	return o.PassUserContext, true
}

// HasPassUserContext returns a boolean if a field has been set.
func (o *SignOnPolicyAction) HasPassUserContext() bool {
	if o != nil && o.PassUserContext != nil {
		return true
	}

	return false
}

// SetPassUserContext gets a reference to the given bool and assigns it to the PassUserContext field.
func (o *SignOnPolicyAction) SetPassUserContext(v bool) {
	o.PassUserContext = &v
}

// GetAgreement returns the Agreement field value
func (o *SignOnPolicyAction) GetAgreement() SignOnPolicyActionAgreementAgreement {
	if o == nil {
		var ret SignOnPolicyActionAgreementAgreement
		return ret
	}

	return o.Agreement
}

// GetAgreementOk returns a tuple with the Agreement field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyAction) GetAgreementOk() (*SignOnPolicyActionAgreementAgreement, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Agreement, true
}

// SetAgreement sets field value
func (o *SignOnPolicyAction) SetAgreement(v SignOnPolicyActionAgreementAgreement) {
	o.Agreement = v
}

// GetAttributes returns the Attributes field value
func (o *SignOnPolicyAction) GetAttributes() SignOnPolicyActionProgressiveProfilingAttributes {
	if o == nil {
		var ret SignOnPolicyActionProgressiveProfilingAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyAction) GetAttributesOk() (*SignOnPolicyActionProgressiveProfilingAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *SignOnPolicyAction) SetAttributes(v SignOnPolicyActionProgressiveProfilingAttributes) {
	o.Attributes = v
}

// GetPreventMultiplePromptsPerFlow returns the PreventMultiplePromptsPerFlow field value
func (o *SignOnPolicyAction) GetPreventMultiplePromptsPerFlow() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PreventMultiplePromptsPerFlow
}

// GetPreventMultiplePromptsPerFlowOk returns a tuple with the PreventMultiplePromptsPerFlow field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyAction) GetPreventMultiplePromptsPerFlowOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreventMultiplePromptsPerFlow, true
}

// SetPreventMultiplePromptsPerFlow sets field value
func (o *SignOnPolicyAction) SetPreventMultiplePromptsPerFlow(v bool) {
	o.PreventMultiplePromptsPerFlow = v
}

// GetPromptIntervalSeconds returns the PromptIntervalSeconds field value
func (o *SignOnPolicyAction) GetPromptIntervalSeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PromptIntervalSeconds
}

// GetPromptIntervalSecondsOk returns a tuple with the PromptIntervalSeconds field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyAction) GetPromptIntervalSecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PromptIntervalSeconds, true
}

// SetPromptIntervalSeconds sets field value
func (o *SignOnPolicyAction) SetPromptIntervalSeconds(v int32) {
	o.PromptIntervalSeconds = v
}

// GetPromptText returns the PromptText field value
func (o *SignOnPolicyAction) GetPromptText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PromptText
}

// GetPromptTextOk returns a tuple with the PromptText field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyAction) GetPromptTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PromptText, true
}

// SetPromptText sets field value
func (o *SignOnPolicyAction) SetPromptText(v string) {
	o.PromptText = v
}

func (o SignOnPolicyAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["priority"] = o.Priority
	}
	if true {
		toSerialize["signOnPolicy"] = o.SignOnPolicy
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.ConfirmIdentityProviderAttributes != nil {
		toSerialize["confirmIdentityProviderAttributes"] = o.ConfirmIdentityProviderAttributes
	}
	if o.EnforceLockoutForIdentityProviders != nil {
		toSerialize["enforceLockoutForIdentityProviders"] = o.EnforceLockoutForIdentityProviders
	}
	if o.Recovery != nil {
		toSerialize["recovery"] = o.Recovery
	}
	if o.Registration != nil {
		toSerialize["registration"] = o.Registration
	}
	if o.SocialProviders != nil {
		toSerialize["socialProviders"] = o.SocialProviders
	}
	if o.Authenticator != nil {
		toSerialize["authenticator"] = o.Authenticator
	}
	if o.BoundBiometrics != nil {
		toSerialize["boundBiometrics"] = o.BoundBiometrics
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.SecurityKey != nil {
		toSerialize["securityKey"] = o.SecurityKey
	}
	if o.Sms != nil {
		toSerialize["sms"] = o.Sms
	}
	if o.Voice != nil {
		toSerialize["voice"] = o.Voice
	}
	if o.Applications != nil {
		toSerialize["applications"] = o.Applications
	}
	if o.NoDeviceMode != nil {
		toSerialize["noDeviceMode"] = o.NoDeviceMode
	}
	if o.DiscoveryRules != nil {
		toSerialize["discoveryRules"] = o.DiscoveryRules
	}
	if o.AcrValues != nil {
		toSerialize["acrValues"] = o.AcrValues
	}
	if true {
		toSerialize["identityProvider"] = o.IdentityProvider
	}
	if o.PassUserContext != nil {
		toSerialize["passUserContext"] = o.PassUserContext
	}
	if true {
		toSerialize["agreement"] = o.Agreement
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if true {
		toSerialize["preventMultiplePromptsPerFlow"] = o.PreventMultiplePromptsPerFlow
	}
	if true {
		toSerialize["promptIntervalSeconds"] = o.PromptIntervalSeconds
	}
	if true {
		toSerialize["promptText"] = o.PromptText
	}
	return json.Marshal(toSerialize)
}

type NullableSignOnPolicyAction struct {
	value *SignOnPolicyAction
	isSet bool
}

func (v NullableSignOnPolicyAction) Get() *SignOnPolicyAction {
	return v.value
}

func (v *NullableSignOnPolicyAction) Set(val *SignOnPolicyAction) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyAction) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyAction(val *SignOnPolicyAction) *NullableSignOnPolicyAction {
	return &NullableSignOnPolicyAction{value: val, isSet: true}
}

func (v NullableSignOnPolicyAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

