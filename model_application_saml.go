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

// ApplicationSAML struct for ApplicationSAML
type ApplicationSAML struct {
	AccessControl *ApplicationAccessControl `json:"accessControl,omitempty"`
	// A boolean that specifies whether the permissions service should assign default roles to the application. This property is set only on the POST request. The property is ignored when included in a PUT request.
	AssignActorRoles *bool `json:"assignActorRoles,omitempty"`
	// The time the resource was created.
	CreatedAt *string `json:"createdAt,omitempty"`
	// A string that specifies the description of the application.
	Description *string `json:"description,omitempty"`
	// A string that specifies the current enabled state of the application. Options are ENABLED or DISABLED.
	Enabled *string `json:"enabled,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	Icon *ApplicationIcon `json:"icon,omitempty"`
	// A string that specifies the application ID.
	Id *string `json:"id,omitempty"`
	// A string that specifies the custom login page URL for the application. If you set the loginPageUrl property for applications in an environment that sets a custom domain, the URL should include the top-level domain and at least one additional domain level. Warning To avoid issues with third-party cookies in some browsers, a custom domain must be used, giving your PingOne environment the same parent domain as your authentication application. For more information about custom domains, see Custom domains.
	LoginPageUrl *string `json:"loginPageUrl,omitempty"`
	// A string that specifies the name of the application. This is a required property.
	Name *string `json:"name,omitempty"`
	// A string that specifies the protocol for the Application. Options are OPENID_CONNECT and SAML.
	Protocol *string `json:"protocol,omitempty"`
	// An array that specifies the list of labels associated with the application. Options are PING_FED_CONNECTION_INTEGRATION.
	Tags *[]string `json:"tags,omitempty"`
	// A string that specifies the type associated with the application. This is a required property. Options are WEB_APP, NATIVE_APP, SINGLE_PAGE_APP, and WORKER.
	Type *string `json:"type,omitempty"`
	// The time the resource was last updated.
	UpdatedAt *string `json:"updatedAt,omitempty"`
	Mobile *ApplicationMobile `json:"mobile,omitempty"`
	// A boolean that specifies whether the request query parameter JWT is allowed to be unsigned. If false or null (default), an unsigned request object is not allowed.
	SupportUnsignedRequestObject *bool `json:"supportUnsignedRequestObject,omitempty"`
	// A string that specifies the Assertion Consumer Service URLs. The first URL in the list is used as default (there must be at least one URL). This is a required property.
	AcsUrls *[]string `json:"acsUrls,omitempty"`
	// An integer that specifies the assertion validity duration in seconds. This is a required property.
	AssertionDuration *int32 `json:"assertionDuration,omitempty"`
	// A boolean that specifies whether the SAML assertion itself should be signed. The default value is true.
	AssertionSigned *bool `json:"assertionSigned,omitempty"`
	IdpSigningtype *ApplicationSAMLAllOfIdpSigningtype `json:"idpSigningtype,omitempty"`
	// A string that specifies the format of the Subject NameID attibute in the SAML assertion
	NameIdFormat *string `json:"nameIdFormat,omitempty"`
	// A boolean that specifies whether the SAML assertion response itself should be signed. The default value is False.
	ResponseSigned *bool `json:"responseSigned,omitempty"`
	// A string that specifies the binding protocol to be used for the logout response. Options are HTTP_REDIRECT or HTTP_POST. The default is HTTP_POST; existing configurations with no data default to HTTP_POST. This is an optional property.
	SloBinding *string `json:"sloBinding,omitempty"`
	// A string that specifies the logout endpoint URL. This is an optional property. However, if a sloEndpoint logout endpoint URL is not defined, logout actions result in an error.
	SloEndpoint *string `json:"sloEndpoint,omitempty"`
	// A string that specifies the endpoint URL to submit the logout response. If a value is not provided, the sloEndpoint property value is used to submit SLO response.
	SloResponseEndpoint *string `json:"sloResponseEndpoint,omitempty"`
	// A string that specifies the service provider entity ID used to lookup the application. This is a required property and is unique within the environment.
	SpEntityId *string `json:"spEntityId,omitempty"`
	SpVerification *ApplicationSAMLAllOfSpVerification `json:"spVerification,omitempty"`
}

// NewApplicationSAML instantiates a new ApplicationSAML object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationSAML() *ApplicationSAML {
	this := ApplicationSAML{}
	return &this
}

// NewApplicationSAMLWithDefaults instantiates a new ApplicationSAML object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationSAMLWithDefaults() *ApplicationSAML {
	this := ApplicationSAML{}
	return &this
}

// GetAccessControl returns the AccessControl field value if set, zero value otherwise.
func (o *ApplicationSAML) GetAccessControl() ApplicationAccessControl {
	if o == nil || o.AccessControl == nil {
		var ret ApplicationAccessControl
		return ret
	}
	return *o.AccessControl
}

// GetAccessControlOk returns a tuple with the AccessControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetAccessControlOk() (*ApplicationAccessControl, bool) {
	if o == nil || o.AccessControl == nil {
		return nil, false
	}
	return o.AccessControl, true
}

// HasAccessControl returns a boolean if a field has been set.
func (o *ApplicationSAML) HasAccessControl() bool {
	if o != nil && o.AccessControl != nil {
		return true
	}

	return false
}

// SetAccessControl gets a reference to the given ApplicationAccessControl and assigns it to the AccessControl field.
func (o *ApplicationSAML) SetAccessControl(v ApplicationAccessControl) {
	o.AccessControl = &v
}

// GetAssignActorRoles returns the AssignActorRoles field value if set, zero value otherwise.
func (o *ApplicationSAML) GetAssignActorRoles() bool {
	if o == nil || o.AssignActorRoles == nil {
		var ret bool
		return ret
	}
	return *o.AssignActorRoles
}

// GetAssignActorRolesOk returns a tuple with the AssignActorRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetAssignActorRolesOk() (*bool, bool) {
	if o == nil || o.AssignActorRoles == nil {
		return nil, false
	}
	return o.AssignActorRoles, true
}

// HasAssignActorRoles returns a boolean if a field has been set.
func (o *ApplicationSAML) HasAssignActorRoles() bool {
	if o != nil && o.AssignActorRoles != nil {
		return true
	}

	return false
}

// SetAssignActorRoles gets a reference to the given bool and assigns it to the AssignActorRoles field.
func (o *ApplicationSAML) SetAssignActorRoles(v bool) {
	o.AssignActorRoles = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ApplicationSAML) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ApplicationSAML) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ApplicationSAML) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApplicationSAML) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApplicationSAML) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApplicationSAML) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ApplicationSAML) GetEnabled() string {
	if o == nil || o.Enabled == nil {
		var ret string
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetEnabledOk() (*string, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ApplicationSAML) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given string and assigns it to the Enabled field.
func (o *ApplicationSAML) SetEnabled(v string) {
	o.Enabled = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *ApplicationSAML) GetEnvironment() ObjectEnvironment {
	if o == nil || o.Environment == nil {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *ApplicationSAML) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *ApplicationSAML) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *ApplicationSAML) GetIcon() ApplicationIcon {
	if o == nil || o.Icon == nil {
		var ret ApplicationIcon
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetIconOk() (*ApplicationIcon, bool) {
	if o == nil || o.Icon == nil {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *ApplicationSAML) HasIcon() bool {
	if o != nil && o.Icon != nil {
		return true
	}

	return false
}

// SetIcon gets a reference to the given ApplicationIcon and assigns it to the Icon field.
func (o *ApplicationSAML) SetIcon(v ApplicationIcon) {
	o.Icon = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApplicationSAML) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApplicationSAML) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApplicationSAML) SetId(v string) {
	o.Id = &v
}

// GetLoginPageUrl returns the LoginPageUrl field value if set, zero value otherwise.
func (o *ApplicationSAML) GetLoginPageUrl() string {
	if o == nil || o.LoginPageUrl == nil {
		var ret string
		return ret
	}
	return *o.LoginPageUrl
}

// GetLoginPageUrlOk returns a tuple with the LoginPageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetLoginPageUrlOk() (*string, bool) {
	if o == nil || o.LoginPageUrl == nil {
		return nil, false
	}
	return o.LoginPageUrl, true
}

// HasLoginPageUrl returns a boolean if a field has been set.
func (o *ApplicationSAML) HasLoginPageUrl() bool {
	if o != nil && o.LoginPageUrl != nil {
		return true
	}

	return false
}

// SetLoginPageUrl gets a reference to the given string and assigns it to the LoginPageUrl field.
func (o *ApplicationSAML) SetLoginPageUrl(v string) {
	o.LoginPageUrl = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApplicationSAML) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApplicationSAML) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApplicationSAML) SetName(v string) {
	o.Name = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *ApplicationSAML) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *ApplicationSAML) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *ApplicationSAML) SetProtocol(v string) {
	o.Protocol = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ApplicationSAML) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ApplicationSAML) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ApplicationSAML) SetTags(v []string) {
	o.Tags = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApplicationSAML) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApplicationSAML) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ApplicationSAML) SetType(v string) {
	o.Type = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ApplicationSAML) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ApplicationSAML) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ApplicationSAML) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetMobile returns the Mobile field value if set, zero value otherwise.
func (o *ApplicationSAML) GetMobile() ApplicationMobile {
	if o == nil || o.Mobile == nil {
		var ret ApplicationMobile
		return ret
	}
	return *o.Mobile
}

// GetMobileOk returns a tuple with the Mobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetMobileOk() (*ApplicationMobile, bool) {
	if o == nil || o.Mobile == nil {
		return nil, false
	}
	return o.Mobile, true
}

// HasMobile returns a boolean if a field has been set.
func (o *ApplicationSAML) HasMobile() bool {
	if o != nil && o.Mobile != nil {
		return true
	}

	return false
}

// SetMobile gets a reference to the given ApplicationMobile and assigns it to the Mobile field.
func (o *ApplicationSAML) SetMobile(v ApplicationMobile) {
	o.Mobile = &v
}

// GetSupportUnsignedRequestObject returns the SupportUnsignedRequestObject field value if set, zero value otherwise.
func (o *ApplicationSAML) GetSupportUnsignedRequestObject() bool {
	if o == nil || o.SupportUnsignedRequestObject == nil {
		var ret bool
		return ret
	}
	return *o.SupportUnsignedRequestObject
}

// GetSupportUnsignedRequestObjectOk returns a tuple with the SupportUnsignedRequestObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetSupportUnsignedRequestObjectOk() (*bool, bool) {
	if o == nil || o.SupportUnsignedRequestObject == nil {
		return nil, false
	}
	return o.SupportUnsignedRequestObject, true
}

// HasSupportUnsignedRequestObject returns a boolean if a field has been set.
func (o *ApplicationSAML) HasSupportUnsignedRequestObject() bool {
	if o != nil && o.SupportUnsignedRequestObject != nil {
		return true
	}

	return false
}

// SetSupportUnsignedRequestObject gets a reference to the given bool and assigns it to the SupportUnsignedRequestObject field.
func (o *ApplicationSAML) SetSupportUnsignedRequestObject(v bool) {
	o.SupportUnsignedRequestObject = &v
}

// GetAcsUrls returns the AcsUrls field value if set, zero value otherwise.
func (o *ApplicationSAML) GetAcsUrls() []string {
	if o == nil || o.AcsUrls == nil {
		var ret []string
		return ret
	}
	return *o.AcsUrls
}

// GetAcsUrlsOk returns a tuple with the AcsUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetAcsUrlsOk() (*[]string, bool) {
	if o == nil || o.AcsUrls == nil {
		return nil, false
	}
	return o.AcsUrls, true
}

// HasAcsUrls returns a boolean if a field has been set.
func (o *ApplicationSAML) HasAcsUrls() bool {
	if o != nil && o.AcsUrls != nil {
		return true
	}

	return false
}

// SetAcsUrls gets a reference to the given []string and assigns it to the AcsUrls field.
func (o *ApplicationSAML) SetAcsUrls(v []string) {
	o.AcsUrls = &v
}

// GetAssertionDuration returns the AssertionDuration field value if set, zero value otherwise.
func (o *ApplicationSAML) GetAssertionDuration() int32 {
	if o == nil || o.AssertionDuration == nil {
		var ret int32
		return ret
	}
	return *o.AssertionDuration
}

// GetAssertionDurationOk returns a tuple with the AssertionDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetAssertionDurationOk() (*int32, bool) {
	if o == nil || o.AssertionDuration == nil {
		return nil, false
	}
	return o.AssertionDuration, true
}

// HasAssertionDuration returns a boolean if a field has been set.
func (o *ApplicationSAML) HasAssertionDuration() bool {
	if o != nil && o.AssertionDuration != nil {
		return true
	}

	return false
}

// SetAssertionDuration gets a reference to the given int32 and assigns it to the AssertionDuration field.
func (o *ApplicationSAML) SetAssertionDuration(v int32) {
	o.AssertionDuration = &v
}

// GetAssertionSigned returns the AssertionSigned field value if set, zero value otherwise.
func (o *ApplicationSAML) GetAssertionSigned() bool {
	if o == nil || o.AssertionSigned == nil {
		var ret bool
		return ret
	}
	return *o.AssertionSigned
}

// GetAssertionSignedOk returns a tuple with the AssertionSigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetAssertionSignedOk() (*bool, bool) {
	if o == nil || o.AssertionSigned == nil {
		return nil, false
	}
	return o.AssertionSigned, true
}

// HasAssertionSigned returns a boolean if a field has been set.
func (o *ApplicationSAML) HasAssertionSigned() bool {
	if o != nil && o.AssertionSigned != nil {
		return true
	}

	return false
}

// SetAssertionSigned gets a reference to the given bool and assigns it to the AssertionSigned field.
func (o *ApplicationSAML) SetAssertionSigned(v bool) {
	o.AssertionSigned = &v
}

// GetIdpSigningtype returns the IdpSigningtype field value if set, zero value otherwise.
func (o *ApplicationSAML) GetIdpSigningtype() ApplicationSAMLAllOfIdpSigningtype {
	if o == nil || o.IdpSigningtype == nil {
		var ret ApplicationSAMLAllOfIdpSigningtype
		return ret
	}
	return *o.IdpSigningtype
}

// GetIdpSigningtypeOk returns a tuple with the IdpSigningtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetIdpSigningtypeOk() (*ApplicationSAMLAllOfIdpSigningtype, bool) {
	if o == nil || o.IdpSigningtype == nil {
		return nil, false
	}
	return o.IdpSigningtype, true
}

// HasIdpSigningtype returns a boolean if a field has been set.
func (o *ApplicationSAML) HasIdpSigningtype() bool {
	if o != nil && o.IdpSigningtype != nil {
		return true
	}

	return false
}

// SetIdpSigningtype gets a reference to the given ApplicationSAMLAllOfIdpSigningtype and assigns it to the IdpSigningtype field.
func (o *ApplicationSAML) SetIdpSigningtype(v ApplicationSAMLAllOfIdpSigningtype) {
	o.IdpSigningtype = &v
}

// GetNameIdFormat returns the NameIdFormat field value if set, zero value otherwise.
func (o *ApplicationSAML) GetNameIdFormat() string {
	if o == nil || o.NameIdFormat == nil {
		var ret string
		return ret
	}
	return *o.NameIdFormat
}

// GetNameIdFormatOk returns a tuple with the NameIdFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetNameIdFormatOk() (*string, bool) {
	if o == nil || o.NameIdFormat == nil {
		return nil, false
	}
	return o.NameIdFormat, true
}

// HasNameIdFormat returns a boolean if a field has been set.
func (o *ApplicationSAML) HasNameIdFormat() bool {
	if o != nil && o.NameIdFormat != nil {
		return true
	}

	return false
}

// SetNameIdFormat gets a reference to the given string and assigns it to the NameIdFormat field.
func (o *ApplicationSAML) SetNameIdFormat(v string) {
	o.NameIdFormat = &v
}

// GetResponseSigned returns the ResponseSigned field value if set, zero value otherwise.
func (o *ApplicationSAML) GetResponseSigned() bool {
	if o == nil || o.ResponseSigned == nil {
		var ret bool
		return ret
	}
	return *o.ResponseSigned
}

// GetResponseSignedOk returns a tuple with the ResponseSigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetResponseSignedOk() (*bool, bool) {
	if o == nil || o.ResponseSigned == nil {
		return nil, false
	}
	return o.ResponseSigned, true
}

// HasResponseSigned returns a boolean if a field has been set.
func (o *ApplicationSAML) HasResponseSigned() bool {
	if o != nil && o.ResponseSigned != nil {
		return true
	}

	return false
}

// SetResponseSigned gets a reference to the given bool and assigns it to the ResponseSigned field.
func (o *ApplicationSAML) SetResponseSigned(v bool) {
	o.ResponseSigned = &v
}

// GetSloBinding returns the SloBinding field value if set, zero value otherwise.
func (o *ApplicationSAML) GetSloBinding() string {
	if o == nil || o.SloBinding == nil {
		var ret string
		return ret
	}
	return *o.SloBinding
}

// GetSloBindingOk returns a tuple with the SloBinding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetSloBindingOk() (*string, bool) {
	if o == nil || o.SloBinding == nil {
		return nil, false
	}
	return o.SloBinding, true
}

// HasSloBinding returns a boolean if a field has been set.
func (o *ApplicationSAML) HasSloBinding() bool {
	if o != nil && o.SloBinding != nil {
		return true
	}

	return false
}

// SetSloBinding gets a reference to the given string and assigns it to the SloBinding field.
func (o *ApplicationSAML) SetSloBinding(v string) {
	o.SloBinding = &v
}

// GetSloEndpoint returns the SloEndpoint field value if set, zero value otherwise.
func (o *ApplicationSAML) GetSloEndpoint() string {
	if o == nil || o.SloEndpoint == nil {
		var ret string
		return ret
	}
	return *o.SloEndpoint
}

// GetSloEndpointOk returns a tuple with the SloEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetSloEndpointOk() (*string, bool) {
	if o == nil || o.SloEndpoint == nil {
		return nil, false
	}
	return o.SloEndpoint, true
}

// HasSloEndpoint returns a boolean if a field has been set.
func (o *ApplicationSAML) HasSloEndpoint() bool {
	if o != nil && o.SloEndpoint != nil {
		return true
	}

	return false
}

// SetSloEndpoint gets a reference to the given string and assigns it to the SloEndpoint field.
func (o *ApplicationSAML) SetSloEndpoint(v string) {
	o.SloEndpoint = &v
}

// GetSloResponseEndpoint returns the SloResponseEndpoint field value if set, zero value otherwise.
func (o *ApplicationSAML) GetSloResponseEndpoint() string {
	if o == nil || o.SloResponseEndpoint == nil {
		var ret string
		return ret
	}
	return *o.SloResponseEndpoint
}

// GetSloResponseEndpointOk returns a tuple with the SloResponseEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetSloResponseEndpointOk() (*string, bool) {
	if o == nil || o.SloResponseEndpoint == nil {
		return nil, false
	}
	return o.SloResponseEndpoint, true
}

// HasSloResponseEndpoint returns a boolean if a field has been set.
func (o *ApplicationSAML) HasSloResponseEndpoint() bool {
	if o != nil && o.SloResponseEndpoint != nil {
		return true
	}

	return false
}

// SetSloResponseEndpoint gets a reference to the given string and assigns it to the SloResponseEndpoint field.
func (o *ApplicationSAML) SetSloResponseEndpoint(v string) {
	o.SloResponseEndpoint = &v
}

// GetSpEntityId returns the SpEntityId field value if set, zero value otherwise.
func (o *ApplicationSAML) GetSpEntityId() string {
	if o == nil || o.SpEntityId == nil {
		var ret string
		return ret
	}
	return *o.SpEntityId
}

// GetSpEntityIdOk returns a tuple with the SpEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetSpEntityIdOk() (*string, bool) {
	if o == nil || o.SpEntityId == nil {
		return nil, false
	}
	return o.SpEntityId, true
}

// HasSpEntityId returns a boolean if a field has been set.
func (o *ApplicationSAML) HasSpEntityId() bool {
	if o != nil && o.SpEntityId != nil {
		return true
	}

	return false
}

// SetSpEntityId gets a reference to the given string and assigns it to the SpEntityId field.
func (o *ApplicationSAML) SetSpEntityId(v string) {
	o.SpEntityId = &v
}

// GetSpVerification returns the SpVerification field value if set, zero value otherwise.
func (o *ApplicationSAML) GetSpVerification() ApplicationSAMLAllOfSpVerification {
	if o == nil || o.SpVerification == nil {
		var ret ApplicationSAMLAllOfSpVerification
		return ret
	}
	return *o.SpVerification
}

// GetSpVerificationOk returns a tuple with the SpVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetSpVerificationOk() (*ApplicationSAMLAllOfSpVerification, bool) {
	if o == nil || o.SpVerification == nil {
		return nil, false
	}
	return o.SpVerification, true
}

// HasSpVerification returns a boolean if a field has been set.
func (o *ApplicationSAML) HasSpVerification() bool {
	if o != nil && o.SpVerification != nil {
		return true
	}

	return false
}

// SetSpVerification gets a reference to the given ApplicationSAMLAllOfSpVerification and assigns it to the SpVerification field.
func (o *ApplicationSAML) SetSpVerification(v ApplicationSAMLAllOfSpVerification) {
	o.SpVerification = &v
}

func (o ApplicationSAML) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessControl != nil {
		toSerialize["accessControl"] = o.AccessControl
	}
	if o.AssignActorRoles != nil {
		toSerialize["assignActorRoles"] = o.AssignActorRoles
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.Icon != nil {
		toSerialize["icon"] = o.Icon
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LoginPageUrl != nil {
		toSerialize["loginPageUrl"] = o.LoginPageUrl
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Protocol != nil {
		toSerialize["protocol"] = o.Protocol
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.Mobile != nil {
		toSerialize["mobile"] = o.Mobile
	}
	if o.SupportUnsignedRequestObject != nil {
		toSerialize["supportUnsignedRequestObject"] = o.SupportUnsignedRequestObject
	}
	if o.AcsUrls != nil {
		toSerialize["acsUrls"] = o.AcsUrls
	}
	if o.AssertionDuration != nil {
		toSerialize["assertionDuration"] = o.AssertionDuration
	}
	if o.AssertionSigned != nil {
		toSerialize["assertionSigned"] = o.AssertionSigned
	}
	if o.IdpSigningtype != nil {
		toSerialize["idpSigningtype"] = o.IdpSigningtype
	}
	if o.NameIdFormat != nil {
		toSerialize["nameIdFormat"] = o.NameIdFormat
	}
	if o.ResponseSigned != nil {
		toSerialize["responseSigned"] = o.ResponseSigned
	}
	if o.SloBinding != nil {
		toSerialize["sloBinding"] = o.SloBinding
	}
	if o.SloEndpoint != nil {
		toSerialize["sloEndpoint"] = o.SloEndpoint
	}
	if o.SloResponseEndpoint != nil {
		toSerialize["sloResponseEndpoint"] = o.SloResponseEndpoint
	}
	if o.SpEntityId != nil {
		toSerialize["spEntityId"] = o.SpEntityId
	}
	if o.SpVerification != nil {
		toSerialize["spVerification"] = o.SpVerification
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationSAML struct {
	value *ApplicationSAML
	isSet bool
}

func (v NullableApplicationSAML) Get() *ApplicationSAML {
	return v.value
}

func (v *NullableApplicationSAML) Set(val *ApplicationSAML) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationSAML) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationSAML) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationSAML(val *ApplicationSAML) *NullableApplicationSAML {
	return &NullableApplicationSAML{value: val, isSet: true}
}

func (v NullableApplicationSAML) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationSAML) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


