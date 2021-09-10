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

// User struct for User
type User struct {
	Account *UserAccount `json:"account,omitempty"`
	Address *UserAddress `json:"address,omitempty"`
	// The time the resource was created.
	CreatedAt *string `json:"createdAt,omitempty"`
	// A string that specifies the user’s email address, which must be provided and valid. For more information about email address formatting, see section 3.4 of RFC 2822, Internet Message Format.
	Email *string `json:"email,omitempty"`
	// A read-only boolean attribute that specifies whether the user is enabled. This attribute is set to ‘true’ by default when the user is created.
	Enabled *bool `json:"enabled,omitempty"`
	Environment *UserEnvironment `json:"environment,omitempty"`
	// A string that specifies an identifier for the user resource as defined by the provisioning client. This is optional. This may be explicitly set to null when updating a user to unset it. The externalId attribute simplifies the correlation of the user in PingOne with the user’s account in another system of record. The platform does not use this attribute directly in any way, but it is used by Ping Identity’s Data Sync product. It can have a length of no more than 1024 characters (min/max=1/1024).
	ExternalId *string `json:"externalId,omitempty"`
	// A string that specifies the user resource’s unique identifier.
	Id *string `json:"id,omitempty"`
	IdentityProvider *UserIdentityProvider `json:"identityProvider,omitempty"`
	LastSignOn *UserLastSignOn `json:"lastSignOn,omitempty"`
	Lifecycle *UserLifecycle `json:"lifecycle,omitempty"`
	// A string that specifies the user’s default location, which is optional. This may be explicitly set to null when updating a user to unset it. This is used for purposes of localizing such items as currency, date time format, or numerical representations. If provided, it must be a valid language tag as defined in RFC 5646. The following are example tags fr, en-US, es-419, az-Arab, man-Nkoo-GN. The string can contain any letters, numbers, combining characters, math and currency symbols, dingbats and drawing characters, and invisible whitespace (regex ^[\\p{L}\\p{M}\\p{Zs}\\p{S}\\p{N}\\p{P}]*$). It can have a length of no more than 256 characters (min/max=1/256).
	Locale *string `json:"locale,omitempty"`
	// A read-only array of IDs for the groups the user is a member of. This property is returned for GET /environments/{envID}/users/{userID} when include=memberOfGroupIDs is appended to the request. This property is not returned with a list of users.
	MemberOfGroupIDs *string `json:"memberOfGroupIDs,omitempty"`
	// A read-only array of names for the groups the user is a member of. This property is returned for GET /environments/{envID}/users/{userID} when include=memberOfGroupNames is appended to the request. This property is not returned with a list of users.
	MemberOfGroupNames *string `json:"memberOfGroupNames,omitempty"`
	// A boolean attribute that specifies whether multi-factor authentication is enabled. This attribute is set to false by default when the user is created. You can set mfaEnabled to true with POST CREATE User, POST CREATE User (Import), or PUT UPDATE User MFA Enabled. You cannot update mfaEnabled with PUT UPDATE User or PATCH UPDATE User.
	MfaEnabled *bool `json:"mfaEnabled,omitempty"`
	// A string that specifies the user’s native phone number, which is optional. This might also match the primaryPhone attribute. This may be explicitly set to null when updating a user to unset it. Valid phone numbers must have at least one number and a maximum character length of 32.
	MobilePhone *string `json:"mobilePhone,omitempty"`
	Name *UserName `json:"name,omitempty"`
	// A string that specifies the user’s nickname, which is optional. This can be explicitly set to null when updating a user to unset it. The string can contain any letters, numbers, combining characters, math and currency symbols, dingbats and drawing characters, and invisible whitespace (regex ^[\\p{L}\\p{M}\\p{Zs}\\p{S}\\p{N}\\p{P}]*$). It can have a length of no more than 256 characters (min/max=1/256).
	Nickname *string `json:"nickname,omitempty"`
	Password *UserPassword `json:"password,omitempty"`
	Photo *UserPhoto `json:"photo,omitempty"`
	Population *UserPopulation `json:"population,omitempty"`
	// A string that specifies the user’s preferred written or spoken languages, which are optional. This may be explicitly set to null when updating a user to unset it. If provided, the format of the value must be a valid language range and the same as the HTTP Accept-Language header field (not including Accept-Language:) and is specified in Section 5.3.5 of RFC 7231. For example en-US, en-gb;q=0.8, en;q=0.7.
	PreferredLanguage *string `json:"preferredLanguage,omitempty"`
	// A string that specifies the user’s primary phone number, which is optional. This might also match the mobilePhone attribute. This may be explicitly set to null when updating a user to unset it. Valid phone numbers must have at least one number and a maximum character length of 32.
	PrimaryPhone *string `json:"primaryPhone,omitempty"`
	// A string that specifies the user’s time zone, which is optional. This can be explicitly set to null when updating a user to unset it. If provided, it must conform with the IANA Time Zone database format [RFC6557], also known as the “Olson” time zone database format [Olson-TZ] for example, “America/Los_Angeles” (regex ^\\w+/\\w+$).
	Timezone *string `json:"timezone,omitempty"`
	// A string that specifies the user’s title, which is optional, such as \"Vice President\". This can be explicitly set to null when updating a user to unset it. The string can contain any letters, numbers, combining characters, math and currency symbols, dingbats and drawing characters, and invisible whitespace (regex ^[\\p{L}\\p{M}\\p{Zs}\\p{S}\\p{N}\\p{P}]*$). It can have a length of no more than 256 characters (min/max=1/256).
	Title *string `json:"title,omitempty"`
	// A string that specifies the user’s type, which is optional. This can be explicitly set to null when updating a user to unset it. This attribute is organization-specific and has no special meaning within the PingOne platform. It could have values of \"Contractor\", \"Employee\", \"Intern\", \"Temp\", \"External\", and “Unknown”. The string can contain any letters, numbers, combining characters, math and currency symbols, dingbats and drawing characters, and invisible whitespace (regex ^[\\p{L}\\p{M}\\p{Zs}\\p{S}\\p{N}\\p{P}]*$). It can have a length of no more than 256 characters (min/max=1/256).
	Type *string `json:"type,omitempty"`
	// The time the resource was last updated.
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// A string that specifies the user name, which must be provided and must be unique within an environment. The username must either be a well-formed email address or a string. The string can contain any letters, numbers, combining characters, math and currency symbols, dingbats and drawing characters, and invisible whitespace (regex ^[\\p{L}\\p{M}\\p{Zs}\\p{S}\\p{N}\\p{P}]*$). It can have a length of no more than 128 characters (min/max=1/128).
	Username *string `json:"username,omitempty"`
	// Indicates whether ID verification can be done for the user. This value can be NOT_INITIATED (the initial value), ENABLED, or DISABLED. If the user verification status is DISABLED, a new verification status cannot be created for that user until the status is changed to ENABLED.
	VerifyStatus *string `json:"verifyStatus,omitempty"`
}

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser() *User {
	this := User{}
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *User) GetAccount() UserAccount {
	if o == nil || o.Account == nil {
		var ret UserAccount
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetAccountOk() (*UserAccount, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *User) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given UserAccount and assigns it to the Account field.
func (o *User) SetAccount(v UserAccount) {
	o.Account = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *User) GetAddress() UserAddress {
	if o == nil || o.Address == nil {
		var ret UserAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetAddressOk() (*UserAddress, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *User) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given UserAddress and assigns it to the Address field.
func (o *User) SetAddress(v UserAddress) {
	o.Address = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *User) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *User) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *User) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *User) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *User) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *User) SetEmail(v string) {
	o.Email = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *User) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *User) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *User) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *User) GetEnvironment() UserEnvironment {
	if o == nil || o.Environment == nil {
		var ret UserEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetEnvironmentOk() (*UserEnvironment, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *User) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given UserEnvironment and assigns it to the Environment field.
func (o *User) SetEnvironment(v UserEnvironment) {
	o.Environment = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *User) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *User) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *User) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *User) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *User) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *User) SetId(v string) {
	o.Id = &v
}

// GetIdentityProvider returns the IdentityProvider field value if set, zero value otherwise.
func (o *User) GetIdentityProvider() UserIdentityProvider {
	if o == nil || o.IdentityProvider == nil {
		var ret UserIdentityProvider
		return ret
	}
	return *o.IdentityProvider
}

// GetIdentityProviderOk returns a tuple with the IdentityProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetIdentityProviderOk() (*UserIdentityProvider, bool) {
	if o == nil || o.IdentityProvider == nil {
		return nil, false
	}
	return o.IdentityProvider, true
}

// HasIdentityProvider returns a boolean if a field has been set.
func (o *User) HasIdentityProvider() bool {
	if o != nil && o.IdentityProvider != nil {
		return true
	}

	return false
}

// SetIdentityProvider gets a reference to the given UserIdentityProvider and assigns it to the IdentityProvider field.
func (o *User) SetIdentityProvider(v UserIdentityProvider) {
	o.IdentityProvider = &v
}

// GetLastSignOn returns the LastSignOn field value if set, zero value otherwise.
func (o *User) GetLastSignOn() UserLastSignOn {
	if o == nil || o.LastSignOn == nil {
		var ret UserLastSignOn
		return ret
	}
	return *o.LastSignOn
}

// GetLastSignOnOk returns a tuple with the LastSignOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLastSignOnOk() (*UserLastSignOn, bool) {
	if o == nil || o.LastSignOn == nil {
		return nil, false
	}
	return o.LastSignOn, true
}

// HasLastSignOn returns a boolean if a field has been set.
func (o *User) HasLastSignOn() bool {
	if o != nil && o.LastSignOn != nil {
		return true
	}

	return false
}

// SetLastSignOn gets a reference to the given UserLastSignOn and assigns it to the LastSignOn field.
func (o *User) SetLastSignOn(v UserLastSignOn) {
	o.LastSignOn = &v
}

// GetLifecycle returns the Lifecycle field value if set, zero value otherwise.
func (o *User) GetLifecycle() UserLifecycle {
	if o == nil || o.Lifecycle == nil {
		var ret UserLifecycle
		return ret
	}
	return *o.Lifecycle
}

// GetLifecycleOk returns a tuple with the Lifecycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLifecycleOk() (*UserLifecycle, bool) {
	if o == nil || o.Lifecycle == nil {
		return nil, false
	}
	return o.Lifecycle, true
}

// HasLifecycle returns a boolean if a field has been set.
func (o *User) HasLifecycle() bool {
	if o != nil && o.Lifecycle != nil {
		return true
	}

	return false
}

// SetLifecycle gets a reference to the given UserLifecycle and assigns it to the Lifecycle field.
func (o *User) SetLifecycle(v UserLifecycle) {
	o.Lifecycle = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *User) GetLocale() string {
	if o == nil || o.Locale == nil {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLocaleOk() (*string, bool) {
	if o == nil || o.Locale == nil {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *User) HasLocale() bool {
	if o != nil && o.Locale != nil {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *User) SetLocale(v string) {
	o.Locale = &v
}

// GetMemberOfGroupIDs returns the MemberOfGroupIDs field value if set, zero value otherwise.
func (o *User) GetMemberOfGroupIDs() string {
	if o == nil || o.MemberOfGroupIDs == nil {
		var ret string
		return ret
	}
	return *o.MemberOfGroupIDs
}

// GetMemberOfGroupIDsOk returns a tuple with the MemberOfGroupIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetMemberOfGroupIDsOk() (*string, bool) {
	if o == nil || o.MemberOfGroupIDs == nil {
		return nil, false
	}
	return o.MemberOfGroupIDs, true
}

// HasMemberOfGroupIDs returns a boolean if a field has been set.
func (o *User) HasMemberOfGroupIDs() bool {
	if o != nil && o.MemberOfGroupIDs != nil {
		return true
	}

	return false
}

// SetMemberOfGroupIDs gets a reference to the given string and assigns it to the MemberOfGroupIDs field.
func (o *User) SetMemberOfGroupIDs(v string) {
	o.MemberOfGroupIDs = &v
}

// GetMemberOfGroupNames returns the MemberOfGroupNames field value if set, zero value otherwise.
func (o *User) GetMemberOfGroupNames() string {
	if o == nil || o.MemberOfGroupNames == nil {
		var ret string
		return ret
	}
	return *o.MemberOfGroupNames
}

// GetMemberOfGroupNamesOk returns a tuple with the MemberOfGroupNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetMemberOfGroupNamesOk() (*string, bool) {
	if o == nil || o.MemberOfGroupNames == nil {
		return nil, false
	}
	return o.MemberOfGroupNames, true
}

// HasMemberOfGroupNames returns a boolean if a field has been set.
func (o *User) HasMemberOfGroupNames() bool {
	if o != nil && o.MemberOfGroupNames != nil {
		return true
	}

	return false
}

// SetMemberOfGroupNames gets a reference to the given string and assigns it to the MemberOfGroupNames field.
func (o *User) SetMemberOfGroupNames(v string) {
	o.MemberOfGroupNames = &v
}

// GetMfaEnabled returns the MfaEnabled field value if set, zero value otherwise.
func (o *User) GetMfaEnabled() bool {
	if o == nil || o.MfaEnabled == nil {
		var ret bool
		return ret
	}
	return *o.MfaEnabled
}

// GetMfaEnabledOk returns a tuple with the MfaEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetMfaEnabledOk() (*bool, bool) {
	if o == nil || o.MfaEnabled == nil {
		return nil, false
	}
	return o.MfaEnabled, true
}

// HasMfaEnabled returns a boolean if a field has been set.
func (o *User) HasMfaEnabled() bool {
	if o != nil && o.MfaEnabled != nil {
		return true
	}

	return false
}

// SetMfaEnabled gets a reference to the given bool and assigns it to the MfaEnabled field.
func (o *User) SetMfaEnabled(v bool) {
	o.MfaEnabled = &v
}

// GetMobilePhone returns the MobilePhone field value if set, zero value otherwise.
func (o *User) GetMobilePhone() string {
	if o == nil || o.MobilePhone == nil {
		var ret string
		return ret
	}
	return *o.MobilePhone
}

// GetMobilePhoneOk returns a tuple with the MobilePhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetMobilePhoneOk() (*string, bool) {
	if o == nil || o.MobilePhone == nil {
		return nil, false
	}
	return o.MobilePhone, true
}

// HasMobilePhone returns a boolean if a field has been set.
func (o *User) HasMobilePhone() bool {
	if o != nil && o.MobilePhone != nil {
		return true
	}

	return false
}

// SetMobilePhone gets a reference to the given string and assigns it to the MobilePhone field.
func (o *User) SetMobilePhone(v string) {
	o.MobilePhone = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *User) GetName() UserName {
	if o == nil || o.Name == nil {
		var ret UserName
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetNameOk() (*UserName, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *User) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given UserName and assigns it to the Name field.
func (o *User) SetName(v UserName) {
	o.Name = &v
}

// GetNickname returns the Nickname field value if set, zero value otherwise.
func (o *User) GetNickname() string {
	if o == nil || o.Nickname == nil {
		var ret string
		return ret
	}
	return *o.Nickname
}

// GetNicknameOk returns a tuple with the Nickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetNicknameOk() (*string, bool) {
	if o == nil || o.Nickname == nil {
		return nil, false
	}
	return o.Nickname, true
}

// HasNickname returns a boolean if a field has been set.
func (o *User) HasNickname() bool {
	if o != nil && o.Nickname != nil {
		return true
	}

	return false
}

// SetNickname gets a reference to the given string and assigns it to the Nickname field.
func (o *User) SetNickname(v string) {
	o.Nickname = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *User) GetPassword() UserPassword {
	if o == nil || o.Password == nil {
		var ret UserPassword
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetPasswordOk() (*UserPassword, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *User) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given UserPassword and assigns it to the Password field.
func (o *User) SetPassword(v UserPassword) {
	o.Password = &v
}

// GetPhoto returns the Photo field value if set, zero value otherwise.
func (o *User) GetPhoto() UserPhoto {
	if o == nil || o.Photo == nil {
		var ret UserPhoto
		return ret
	}
	return *o.Photo
}

// GetPhotoOk returns a tuple with the Photo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetPhotoOk() (*UserPhoto, bool) {
	if o == nil || o.Photo == nil {
		return nil, false
	}
	return o.Photo, true
}

// HasPhoto returns a boolean if a field has been set.
func (o *User) HasPhoto() bool {
	if o != nil && o.Photo != nil {
		return true
	}

	return false
}

// SetPhoto gets a reference to the given UserPhoto and assigns it to the Photo field.
func (o *User) SetPhoto(v UserPhoto) {
	o.Photo = &v
}

// GetPopulation returns the Population field value if set, zero value otherwise.
func (o *User) GetPopulation() UserPopulation {
	if o == nil || o.Population == nil {
		var ret UserPopulation
		return ret
	}
	return *o.Population
}

// GetPopulationOk returns a tuple with the Population field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetPopulationOk() (*UserPopulation, bool) {
	if o == nil || o.Population == nil {
		return nil, false
	}
	return o.Population, true
}

// HasPopulation returns a boolean if a field has been set.
func (o *User) HasPopulation() bool {
	if o != nil && o.Population != nil {
		return true
	}

	return false
}

// SetPopulation gets a reference to the given UserPopulation and assigns it to the Population field.
func (o *User) SetPopulation(v UserPopulation) {
	o.Population = &v
}

// GetPreferredLanguage returns the PreferredLanguage field value if set, zero value otherwise.
func (o *User) GetPreferredLanguage() string {
	if o == nil || o.PreferredLanguage == nil {
		var ret string
		return ret
	}
	return *o.PreferredLanguage
}

// GetPreferredLanguageOk returns a tuple with the PreferredLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetPreferredLanguageOk() (*string, bool) {
	if o == nil || o.PreferredLanguage == nil {
		return nil, false
	}
	return o.PreferredLanguage, true
}

// HasPreferredLanguage returns a boolean if a field has been set.
func (o *User) HasPreferredLanguage() bool {
	if o != nil && o.PreferredLanguage != nil {
		return true
	}

	return false
}

// SetPreferredLanguage gets a reference to the given string and assigns it to the PreferredLanguage field.
func (o *User) SetPreferredLanguage(v string) {
	o.PreferredLanguage = &v
}

// GetPrimaryPhone returns the PrimaryPhone field value if set, zero value otherwise.
func (o *User) GetPrimaryPhone() string {
	if o == nil || o.PrimaryPhone == nil {
		var ret string
		return ret
	}
	return *o.PrimaryPhone
}

// GetPrimaryPhoneOk returns a tuple with the PrimaryPhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetPrimaryPhoneOk() (*string, bool) {
	if o == nil || o.PrimaryPhone == nil {
		return nil, false
	}
	return o.PrimaryPhone, true
}

// HasPrimaryPhone returns a boolean if a field has been set.
func (o *User) HasPrimaryPhone() bool {
	if o != nil && o.PrimaryPhone != nil {
		return true
	}

	return false
}

// SetPrimaryPhone gets a reference to the given string and assigns it to the PrimaryPhone field.
func (o *User) SetPrimaryPhone(v string) {
	o.PrimaryPhone = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *User) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *User) HasTimezone() bool {
	if o != nil && o.Timezone != nil {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *User) SetTimezone(v string) {
	o.Timezone = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *User) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *User) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *User) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *User) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *User) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *User) SetType(v string) {
	o.Type = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *User) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *User) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *User) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *User) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *User) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *User) SetUsername(v string) {
	o.Username = &v
}

// GetVerifyStatus returns the VerifyStatus field value if set, zero value otherwise.
func (o *User) GetVerifyStatus() string {
	if o == nil || o.VerifyStatus == nil {
		var ret string
		return ret
	}
	return *o.VerifyStatus
}

// GetVerifyStatusOk returns a tuple with the VerifyStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetVerifyStatusOk() (*string, bool) {
	if o == nil || o.VerifyStatus == nil {
		return nil, false
	}
	return o.VerifyStatus, true
}

// HasVerifyStatus returns a boolean if a field has been set.
func (o *User) HasVerifyStatus() bool {
	if o != nil && o.VerifyStatus != nil {
		return true
	}

	return false
}

// SetVerifyStatus gets a reference to the given string and assigns it to the VerifyStatus field.
func (o *User) SetVerifyStatus(v string) {
	o.VerifyStatus = &v
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Account != nil {
		toSerialize["account"] = o.Account
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.ExternalId != nil {
		toSerialize["externalId"] = o.ExternalId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IdentityProvider != nil {
		toSerialize["identityProvider"] = o.IdentityProvider
	}
	if o.LastSignOn != nil {
		toSerialize["lastSignOn"] = o.LastSignOn
	}
	if o.Lifecycle != nil {
		toSerialize["lifecycle"] = o.Lifecycle
	}
	if o.Locale != nil {
		toSerialize["locale"] = o.Locale
	}
	if o.MemberOfGroupIDs != nil {
		toSerialize["memberOfGroupIDs"] = o.MemberOfGroupIDs
	}
	if o.MemberOfGroupNames != nil {
		toSerialize["memberOfGroupNames"] = o.MemberOfGroupNames
	}
	if o.MfaEnabled != nil {
		toSerialize["mfaEnabled"] = o.MfaEnabled
	}
	if o.MobilePhone != nil {
		toSerialize["mobilePhone"] = o.MobilePhone
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Nickname != nil {
		toSerialize["nickname"] = o.Nickname
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.Photo != nil {
		toSerialize["photo"] = o.Photo
	}
	if o.Population != nil {
		toSerialize["population"] = o.Population
	}
	if o.PreferredLanguage != nil {
		toSerialize["preferredLanguage"] = o.PreferredLanguage
	}
	if o.PrimaryPhone != nil {
		toSerialize["primaryPhone"] = o.PrimaryPhone
	}
	if o.Timezone != nil {
		toSerialize["timezone"] = o.Timezone
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.VerifyStatus != nil {
		toSerialize["verifyStatus"] = o.VerifyStatus
	}
	return json.Marshal(toSerialize)
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

