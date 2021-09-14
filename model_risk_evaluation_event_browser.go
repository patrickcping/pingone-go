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

// RiskEvaluationEventBrowser An object that specifies the browser fingerprint attributes. Browser data can be retrieved using browser fingerprint JS. For more information, see Overview of the PingOne Risk Integration Kit.
type RiskEvaluationEventBrowser struct {
	UserAgent *string `json:"userAgent,omitempty"`
	Language *string `json:"language,omitempty"`
	ColorDepth *float32 `json:"colorDepth,omitempty"`
	DeviceMemory *float32 `json:"deviceMemory,omitempty"`
	HardwareConcurrency *float32 `json:"hardwareConcurrency,omitempty"`
	ScreenResolution *[]float32 `json:"screenResolution,omitempty"`
	AvailableScreenResolution *[]float32 `json:"availableScreenResolution,omitempty"`
	TimezoneOffset *float32 `json:"timezoneOffset,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
	SessionStorage *bool `json:"sessionStorage,omitempty"`
	LocalStorage *bool `json:"localStorage,omitempty"`
	IndexedDb *bool `json:"indexedDb,omitempty"`
	AddBehaviour *map[string]interface{} `json:"addBehaviour,omitempty"`
	OpenDatabase *bool `json:"openDatabase,omitempty"`
	CpuClass *string `json:"cpuClass,omitempty"`
	Platform *string `json:"platform,omitempty"`
	Plugins *[]map[string]interface{} `json:"plugins,omitempty"`
	WebglVendorAndRenderer *string `json:"webglVendorAndRenderer,omitempty"`
	Webgl *[]string `json:"webgl,omitempty"`
	AdBlock *bool `json:"adBlock,omitempty"`
	HasLiedLanguages *bool `json:"hasLiedLanguages,omitempty"`
	HasLiedResolution *bool `json:"hasLiedResolution,omitempty"`
	HasLiedOs *bool `json:"hasLiedOs,omitempty"`
	HasLiedBrowser *bool `json:"hasLiedBrowser,omitempty"`
	TouchSupport *[]string `json:"touchSupport,omitempty"`
	Fonts *[]string `json:"fonts,omitempty"`
	Audio *string `json:"audio,omitempty"`
}

// NewRiskEvaluationEventBrowser instantiates a new RiskEvaluationEventBrowser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskEvaluationEventBrowser() *RiskEvaluationEventBrowser {
	this := RiskEvaluationEventBrowser{}
	return &this
}

// NewRiskEvaluationEventBrowserWithDefaults instantiates a new RiskEvaluationEventBrowser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskEvaluationEventBrowserWithDefaults() *RiskEvaluationEventBrowser {
	this := RiskEvaluationEventBrowser{}
	return &this
}

// GetUserAgent returns the UserAgent field value if set, zero value otherwise.
func (o *RiskEvaluationEventBrowser) GetUserAgent() string {
	if o == nil || o.UserAgent == nil {
		var ret string
		return ret
	}
	return *o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEventBrowser) GetUserAgentOk() (*string, bool) {
	if o == nil || o.UserAgent == nil {
		return nil, false
	}
	return o.UserAgent, true
}

// HasUserAgent returns a boolean if a field has been set.
func (o *RiskEvaluationEventBrowser) HasUserAgent() bool {
	if o != nil && o.UserAgent != nil {
		return true
	}

	return false
}

// SetUserAgent gets a reference to the given string and assigns it to the UserAgent field.
func (o *RiskEvaluationEventBrowser) SetUserAgent(v string) {
	o.UserAgent = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *RiskEvaluationEventBrowser) GetLanguage() string {
	if o == nil || o.Language == nil {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEventBrowser) GetLanguageOk() (*string, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *RiskEvaluationEventBrowser) HasLanguage() bool {
	if o != nil && o.Language != nil {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *RiskEvaluationEventBrowser) SetLanguage(v string) {
	o.Language = &v
}

// GetColorDepth returns the ColorDepth field value if set, zero value otherwise.
func (o *RiskEvaluationEventBrowser) GetColorDepth() float32 {
	if o == nil || o.ColorDepth == nil {
		var ret float32
		return ret
	}
	return *o.ColorDepth
}

// GetColorDepthOk returns a tuple with the ColorDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEventBrowser) GetColorDepthOk() (*float32, bool) {
	if o == nil || o.ColorDepth == nil {
		return nil, false
	}
	return o.ColorDepth, true
}

// HasColorDepth returns a boolean if a field has been set.
func (o *RiskEvaluationEventBrowser) HasColorDepth() bool {
	if o != nil && o.ColorDepth != nil {
		return true
	}

	return false
}

// SetColorDepth gets a reference to the given float32 and assigns it to the ColorDepth field.
func (o *RiskEvaluationEventBrowser) SetColorDepth(v float32) {
	o.ColorDepth = &v
}

// GetDeviceMemory returns the DeviceMemory field value if set, zero value otherwise.
func (o *RiskEvaluationEventBrowser) GetDeviceMemory() float32 {
	if o == nil || o.DeviceMemory == nil {
		var ret float32
		return ret
	}
	return *o.DeviceMemory
}

// GetDeviceMemoryOk returns a tuple with the DeviceMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEventBrowser) GetDeviceMemoryOk() (*float32, bool) {
	if o == nil || o.DeviceMemory == nil {
		return nil, false
	}
	return o.DeviceMemory, true
}

// HasDeviceMemory returns a boolean if a field has been set.
func (o *RiskEvaluationEventBrowser) HasDeviceMemory() bool {
	if o != nil && o.DeviceMemory != nil {
		return true
	}

	return false
}

// SetDeviceMemory gets a reference to the given float32 and assigns it to the DeviceMemory field.
func (o *RiskEvaluationEventBrowser) SetDeviceMemory(v float32) {
	o.DeviceMemory = &v
}

// GetHardwareConcurrency returns the HardwareConcurrency field value if set, zero value otherwise.
func (o *RiskEvaluationEventBrowser) GetHardwareConcurrency() float32 {
	if o == nil || o.HardwareConcurrency == nil {
		var ret float32
		return ret
	}
	return *o.HardwareConcurrency
}

// GetHardwareConcurrencyOk returns a tuple with the HardwareConcurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEventBrowser) GetHardwareConcurrencyOk() (*float32, bool) {
	if o == nil || o.HardwareConcurrency == nil {
		return nil, false
	}
	return o.HardwareConcurrency, true
}

// HasHardwareConcurrency returns a boolean if a field has been set.
func (o *RiskEvaluationEventBrowser) HasHardwareConcurrency() bool {
	if o != nil && o.HardwareConcurrency != nil {
		return true
	}

	return false
}

// SetHardwareConcurrency gets a reference to the given float32 and assigns it to the HardwareConcurrency field.
func (o *RiskEvaluationEventBrowser) SetHardwareConcurrency(v float32) {
	o.HardwareConcurrency = &v
}

// GetScreenResolution returns the ScreenResolution field value if set, zero value otherwise.
func (o *RiskEvaluationEventBrowser) GetScreenResolution() []float32 {
	if o == nil || o.ScreenResolution == nil {
		var ret []float32
		return ret
	}
	return *o.ScreenResolution
}

// GetScreenResolutionOk returns a tuple with the ScreenResolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEventBrowser) GetScreenResolutionOk() (*[]float32, bool) {
	if o == nil || o.ScreenResolution == nil {
		return nil, false
	}
	return o.ScreenResolution, true
}

// HasScreenResolution returns a boolean if a field has been set.
func (o *RiskEvaluationEventBrowser) HasScreenResolution() bool {
	if o != nil && o.ScreenResolution != nil {
		return true
	}

	return false
}

// SetScreenResolution gets a reference to the given []float32 and assigns it to the ScreenResolution field.
func (o *RiskEvaluationEventBrowser) SetScreenResolution(v []float32) {
	o.ScreenResolution = &v
}

// GetAvailableScreenResolution returns the AvailableScreenResolution field value if set, zero value otherwise.
func (o *RiskEvaluationEventBrowser) GetAvailableScreenResolution() []float32 {
	if o == nil || o.AvailableScreenResolution == nil {
		var ret []float32
		return ret
	}
	return *o.AvailableScreenResolution
}

// GetAvailableScreenResolutionOk returns a tuple with the AvailableScreenResolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEventBrowser) GetAvailableScreenResolutionOk() (*[]float32, bool) {
	if o == nil || o.AvailableScreenResolution == nil {
		return nil, false
	}
	return o.AvailableScreenResolution, true
}

// HasAvailableScreenResolution returns a boolean if a field has been set.
func (o *RiskEvaluationEventBrowser) HasAvailableScreenResolution() bool {
	if o != nil && o.AvailableScreenResolution != nil {
		return true
	}

	return false
}

// SetAvailableScreenResolution gets a reference to the given []float32 and assigns it to the AvailableScreenResolution field.
func (o *RiskEvaluationEventBrowser) SetAvailableScreenResolution(v []float32) {
	o.AvailableScreenResolution = &v
}

// GetTimezoneOffset returns the TimezoneOffset field value if set, zero value otherwise.
func (o *RiskEvaluationEventBrowser) GetTimezoneOffset() float32 {
	if o == nil || o.TimezoneOffset == nil {
		var ret float32
		return ret
	}
	return *o.TimezoneOffset
}

// GetTimezoneOffsetOk returns a tuple with the TimezoneOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEventBrowser) GetTimezoneOffsetOk() (*float32, bool) {
	if o == nil || o.TimezoneOffset == nil {
		return nil, false
	}
	return o.TimezoneOffset, true
}

// HasTimezoneOffset returns a boolean if a field has been set.
func (o *RiskEvaluationEventBrowser) HasTimezoneOffset() bool {
	if o != nil && o.TimezoneOffset != nil {
		return true
	}

	return false
}

// SetTimezoneOffset gets a reference to the given float32 and assigns it to the TimezoneOffset field.
func (o *RiskEvaluationEventBrowser) SetTimezoneOffset(v float32) {
	o.TimezoneOffset = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *RiskEvaluationEventBrowser) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEventBrowser) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *RiskEvaluationEventBrowser) HasTimezone() bool {
	if o != nil && o.Timezone != nil {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *RiskEvaluationEventBrowser) SetTimezone(v string) {
	o.Timezone = &v
}

// GetSessionStorage returns the SessionStorage field value if set, zero value otherwise.
func (o *RiskEvaluationEventBrowser) GetSessionStorage() bool {
	if o == nil || o.SessionStorage == nil {
		var ret bool
		return ret
	}
	return *o.SessionStorage
}

// GetSessionStorageOk returns a tuple with the SessionStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEventBrowser) GetSessionStorageOk() (*bool, bool) {
	if o == nil || o.SessionStorage == nil {
		return nil, false
	}
	return o.SessionStorage, true
}

// HasSessionStorage returns a boolean if a field has been set.
func (o *RiskEvaluationEventBrowser) HasSessionStorage() bool {
	if o != nil && o.SessionStorage != nil {
		return true
	}

	return false
}

// SetSessionStorage gets a reference to the given bool and assigns it to the SessionStorage field.
func (o *RiskEvaluationEventBrowser) SetSessionStorage(v bool) {
	o.SessionStorage = &v
}

// GetLocalStorage returns the LocalStorage field value if set, zero value otherwise.
func (o *RiskEvaluationEventBrowser) GetLocalStorage() bool {
	if o == nil || o.LocalStorage == nil {
		var ret bool
		return ret
	}
	return *o.LocalStorage
}

// GetLocalStorageOk returns a tuple with the LocalStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEventBrowser) GetLocalStorageOk() (*bool, bool) {
	if o == nil || o.LocalStorage == nil {
		return nil, false
	}
	return o.LocalStorage, true
}

// HasLocalStorage returns a boolean if a field has been set.
func (o *RiskEvaluationEventBrowser) HasLocalStorage() bool {
	if o != nil && o.LocalStorage != nil {
		return true
	}

	return false
}

// SetLocalStorage gets a reference to the given bool and assigns it to the LocalStorage field.
func (o *RiskEvaluationEventBrowser) SetLocalStorage(v bool) {
	o.LocalStorage = &v
}

// GetIndexedDb returns the IndexedDb field value if set, zero value otherwise.
func (o *RiskEvaluationEventBrowser) GetIndexedDb() bool {
	if o == nil || o.IndexedDb == nil {
		var ret bool
		return ret
	}
	return *o.IndexedDb
}

// GetIndexedDbOk returns a tuple with the IndexedDb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEventBrowser) GetIndexedDbOk() (*bool, bool) {
	if o == nil || o.IndexedDb == nil {
		return nil, false
	}
	return o.IndexedDb, true
}

// HasIndexedDb returns a boolean if a field has been set.
func (o *RiskEvaluationEventBrowser) HasIndexedDb() bool {
	if o != nil && o.IndexedDb != nil {
		return true
	}

	return false
}

// SetIndexedDb gets a reference to the given bool and assigns it to the IndexedDb field.
func (o *RiskEvaluationEventBrowser) SetIndexedDb(v bool) {
	o.IndexedDb = &v
}

// GetAddBehaviour returns the AddBehaviour field value if set, zero value otherwise.
func (o *RiskEvaluationEventBrowser) GetAddBehaviour() map[string]interface{} {
	if o == nil || o.AddBehaviour == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.AddBehaviour
}

// GetAddBehaviourOk returns a tuple with the AddBehaviour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEventBrowser) GetAddBehaviourOk() (*map[string]interface{}, bool) {
	if o == nil || o.AddBehaviour == nil {
		return nil, false
	}
	return o.AddBehaviour, true
}

// HasAddBehaviour returns a boolean if a field has been set.
func (o *RiskEvaluationEventBrowser) HasAddBehaviour() bool {
	if o != nil && o.AddBehaviour != nil {
		return true
	}

	return false
}

// SetAddBehaviour gets a reference to the given map[string]interface{} and assigns it to the AddBehaviour field.
func (o *RiskEvaluationEventBrowser) SetAddBehaviour(v map[string]interface{}) {
	o.AddBehaviour = &v
}

// GetOpenDatabase returns the OpenDatabase field value if set, zero value otherwise.
func (o *RiskEvaluationEventBrowser) GetOpenDatabase() bool {
	if o == nil || o.OpenDatabase == nil {
		var ret bool
		return ret
	}
	return *o.OpenDatabase
}

// GetOpenDatabaseOk returns a tuple with the OpenDatabase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEventBrowser) GetOpenDatabaseOk() (*bool, bool) {
	if o == nil || o.OpenDatabase == nil {
		return nil, false
	}
	return o.OpenDatabase, true
}

// HasOpenDatabase returns a boolean if a field has been set.
func (o *RiskEvaluationEventBrowser) HasOpenDatabase() bool {
	if o != nil && o.OpenDatabase != nil {
		return true
	}

	return false
}

// SetOpenDatabase gets a reference to the given bool and assigns it to the OpenDatabase field.
func (o *RiskEvaluationEventBrowser) SetOpenDatabase(v bool) {
	o.OpenDatabase = &v
}

// GetCpuClass returns the CpuClass field value if set, zero value otherwise.
func (o *RiskEvaluationEventBrowser) GetCpuClass() string {
	if o == nil || o.CpuClass == nil {
		var ret string
		return ret
	}
	return *o.CpuClass
}

// GetCpuClassOk returns a tuple with the CpuClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEventBrowser) GetCpuClassOk() (*string, bool) {
	if o == nil || o.CpuClass == nil {
		return nil, false
	}
	return o.CpuClass, true
}

// HasCpuClass returns a boolean if a field has been set.
func (o *RiskEvaluationEventBrowser) HasCpuClass() bool {
	if o != nil && o.CpuClass != nil {
		return true
	}

	return false
}

// SetCpuClass gets a reference to the given string and assigns it to the CpuClass field.
func (o *RiskEvaluationEventBrowser) SetCpuClass(v string) {
	o.CpuClass = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *RiskEvaluationEventBrowser) GetPlatform() string {
	if o == nil || o.Platform == nil {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEventBrowser) GetPlatformOk() (*string, bool) {
	if o == nil || o.Platform == nil {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *RiskEvaluationEventBrowser) HasPlatform() bool {
	if o != nil && o.Platform != nil {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *RiskEvaluationEventBrowser) SetPlatform(v string) {
	o.Platform = &v
}

// GetPlugins returns the Plugins field value if set, zero value otherwise.
func (o *RiskEvaluationEventBrowser) GetPlugins() []map[string]interface{} {
	if o == nil || o.Plugins == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Plugins
}

// GetPluginsOk returns a tuple with the Plugins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEventBrowser) GetPluginsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Plugins == nil {
		return nil, false
	}
	return o.Plugins, true
}

// HasPlugins returns a boolean if a field has been set.
func (o *RiskEvaluationEventBrowser) HasPlugins() bool {
	if o != nil && o.Plugins != nil {
		return true
	}

	return false
}

// SetPlugins gets a reference to the given []map[string]interface{} and assigns it to the Plugins field.
func (o *RiskEvaluationEventBrowser) SetPlugins(v []map[string]interface{}) {
	o.Plugins = &v
}

// GetWebglVendorAndRenderer returns the WebglVendorAndRenderer field value if set, zero value otherwise.
func (o *RiskEvaluationEventBrowser) GetWebglVendorAndRenderer() string {
	if o == nil || o.WebglVendorAndRenderer == nil {
		var ret string
		return ret
	}
	return *o.WebglVendorAndRenderer
}

// GetWebglVendorAndRendererOk returns a tuple with the WebglVendorAndRenderer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEventBrowser) GetWebglVendorAndRendererOk() (*string, bool) {
	if o == nil || o.WebglVendorAndRenderer == nil {
		return nil, false
	}
	return o.WebglVendorAndRenderer, true
}

// HasWebglVendorAndRenderer returns a boolean if a field has been set.
func (o *RiskEvaluationEventBrowser) HasWebglVendorAndRenderer() bool {
	if o != nil && o.WebglVendorAndRenderer != nil {
		return true
	}

	return false
}

// SetWebglVendorAndRenderer gets a reference to the given string and assigns it to the WebglVendorAndRenderer field.
func (o *RiskEvaluationEventBrowser) SetWebglVendorAndRenderer(v string) {
	o.WebglVendorAndRenderer = &v
}

// GetWebgl returns the Webgl field value if set, zero value otherwise.
func (o *RiskEvaluationEventBrowser) GetWebgl() []string {
	if o == nil || o.Webgl == nil {
		var ret []string
		return ret
	}
	return *o.Webgl
}

// GetWebglOk returns a tuple with the Webgl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEventBrowser) GetWebglOk() (*[]string, bool) {
	if o == nil || o.Webgl == nil {
		return nil, false
	}
	return o.Webgl, true
}

// HasWebgl returns a boolean if a field has been set.
func (o *RiskEvaluationEventBrowser) HasWebgl() bool {
	if o != nil && o.Webgl != nil {
		return true
	}

	return false
}

// SetWebgl gets a reference to the given []string and assigns it to the Webgl field.
func (o *RiskEvaluationEventBrowser) SetWebgl(v []string) {
	o.Webgl = &v
}

// GetAdBlock returns the AdBlock field value if set, zero value otherwise.
func (o *RiskEvaluationEventBrowser) GetAdBlock() bool {
	if o == nil || o.AdBlock == nil {
		var ret bool
		return ret
	}
	return *o.AdBlock
}

// GetAdBlockOk returns a tuple with the AdBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEventBrowser) GetAdBlockOk() (*bool, bool) {
	if o == nil || o.AdBlock == nil {
		return nil, false
	}
	return o.AdBlock, true
}

// HasAdBlock returns a boolean if a field has been set.
func (o *RiskEvaluationEventBrowser) HasAdBlock() bool {
	if o != nil && o.AdBlock != nil {
		return true
	}

	return false
}

// SetAdBlock gets a reference to the given bool and assigns it to the AdBlock field.
func (o *RiskEvaluationEventBrowser) SetAdBlock(v bool) {
	o.AdBlock = &v
}

// GetHasLiedLanguages returns the HasLiedLanguages field value if set, zero value otherwise.
func (o *RiskEvaluationEventBrowser) GetHasLiedLanguages() bool {
	if o == nil || o.HasLiedLanguages == nil {
		var ret bool
		return ret
	}
	return *o.HasLiedLanguages
}

// GetHasLiedLanguagesOk returns a tuple with the HasLiedLanguages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEventBrowser) GetHasLiedLanguagesOk() (*bool, bool) {
	if o == nil || o.HasLiedLanguages == nil {
		return nil, false
	}
	return o.HasLiedLanguages, true
}

// HasHasLiedLanguages returns a boolean if a field has been set.
func (o *RiskEvaluationEventBrowser) HasHasLiedLanguages() bool {
	if o != nil && o.HasLiedLanguages != nil {
		return true
	}

	return false
}

// SetHasLiedLanguages gets a reference to the given bool and assigns it to the HasLiedLanguages field.
func (o *RiskEvaluationEventBrowser) SetHasLiedLanguages(v bool) {
	o.HasLiedLanguages = &v
}

// GetHasLiedResolution returns the HasLiedResolution field value if set, zero value otherwise.
func (o *RiskEvaluationEventBrowser) GetHasLiedResolution() bool {
	if o == nil || o.HasLiedResolution == nil {
		var ret bool
		return ret
	}
	return *o.HasLiedResolution
}

// GetHasLiedResolutionOk returns a tuple with the HasLiedResolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEventBrowser) GetHasLiedResolutionOk() (*bool, bool) {
	if o == nil || o.HasLiedResolution == nil {
		return nil, false
	}
	return o.HasLiedResolution, true
}

// HasHasLiedResolution returns a boolean if a field has been set.
func (o *RiskEvaluationEventBrowser) HasHasLiedResolution() bool {
	if o != nil && o.HasLiedResolution != nil {
		return true
	}

	return false
}

// SetHasLiedResolution gets a reference to the given bool and assigns it to the HasLiedResolution field.
func (o *RiskEvaluationEventBrowser) SetHasLiedResolution(v bool) {
	o.HasLiedResolution = &v
}

// GetHasLiedOs returns the HasLiedOs field value if set, zero value otherwise.
func (o *RiskEvaluationEventBrowser) GetHasLiedOs() bool {
	if o == nil || o.HasLiedOs == nil {
		var ret bool
		return ret
	}
	return *o.HasLiedOs
}

// GetHasLiedOsOk returns a tuple with the HasLiedOs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEventBrowser) GetHasLiedOsOk() (*bool, bool) {
	if o == nil || o.HasLiedOs == nil {
		return nil, false
	}
	return o.HasLiedOs, true
}

// HasHasLiedOs returns a boolean if a field has been set.
func (o *RiskEvaluationEventBrowser) HasHasLiedOs() bool {
	if o != nil && o.HasLiedOs != nil {
		return true
	}

	return false
}

// SetHasLiedOs gets a reference to the given bool and assigns it to the HasLiedOs field.
func (o *RiskEvaluationEventBrowser) SetHasLiedOs(v bool) {
	o.HasLiedOs = &v
}

// GetHasLiedBrowser returns the HasLiedBrowser field value if set, zero value otherwise.
func (o *RiskEvaluationEventBrowser) GetHasLiedBrowser() bool {
	if o == nil || o.HasLiedBrowser == nil {
		var ret bool
		return ret
	}
	return *o.HasLiedBrowser
}

// GetHasLiedBrowserOk returns a tuple with the HasLiedBrowser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEventBrowser) GetHasLiedBrowserOk() (*bool, bool) {
	if o == nil || o.HasLiedBrowser == nil {
		return nil, false
	}
	return o.HasLiedBrowser, true
}

// HasHasLiedBrowser returns a boolean if a field has been set.
func (o *RiskEvaluationEventBrowser) HasHasLiedBrowser() bool {
	if o != nil && o.HasLiedBrowser != nil {
		return true
	}

	return false
}

// SetHasLiedBrowser gets a reference to the given bool and assigns it to the HasLiedBrowser field.
func (o *RiskEvaluationEventBrowser) SetHasLiedBrowser(v bool) {
	o.HasLiedBrowser = &v
}

// GetTouchSupport returns the TouchSupport field value if set, zero value otherwise.
func (o *RiskEvaluationEventBrowser) GetTouchSupport() []string {
	if o == nil || o.TouchSupport == nil {
		var ret []string
		return ret
	}
	return *o.TouchSupport
}

// GetTouchSupportOk returns a tuple with the TouchSupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEventBrowser) GetTouchSupportOk() (*[]string, bool) {
	if o == nil || o.TouchSupport == nil {
		return nil, false
	}
	return o.TouchSupport, true
}

// HasTouchSupport returns a boolean if a field has been set.
func (o *RiskEvaluationEventBrowser) HasTouchSupport() bool {
	if o != nil && o.TouchSupport != nil {
		return true
	}

	return false
}

// SetTouchSupport gets a reference to the given []string and assigns it to the TouchSupport field.
func (o *RiskEvaluationEventBrowser) SetTouchSupport(v []string) {
	o.TouchSupport = &v
}

// GetFonts returns the Fonts field value if set, zero value otherwise.
func (o *RiskEvaluationEventBrowser) GetFonts() []string {
	if o == nil || o.Fonts == nil {
		var ret []string
		return ret
	}
	return *o.Fonts
}

// GetFontsOk returns a tuple with the Fonts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEventBrowser) GetFontsOk() (*[]string, bool) {
	if o == nil || o.Fonts == nil {
		return nil, false
	}
	return o.Fonts, true
}

// HasFonts returns a boolean if a field has been set.
func (o *RiskEvaluationEventBrowser) HasFonts() bool {
	if o != nil && o.Fonts != nil {
		return true
	}

	return false
}

// SetFonts gets a reference to the given []string and assigns it to the Fonts field.
func (o *RiskEvaluationEventBrowser) SetFonts(v []string) {
	o.Fonts = &v
}

// GetAudio returns the Audio field value if set, zero value otherwise.
func (o *RiskEvaluationEventBrowser) GetAudio() string {
	if o == nil || o.Audio == nil {
		var ret string
		return ret
	}
	return *o.Audio
}

// GetAudioOk returns a tuple with the Audio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEventBrowser) GetAudioOk() (*string, bool) {
	if o == nil || o.Audio == nil {
		return nil, false
	}
	return o.Audio, true
}

// HasAudio returns a boolean if a field has been set.
func (o *RiskEvaluationEventBrowser) HasAudio() bool {
	if o != nil && o.Audio != nil {
		return true
	}

	return false
}

// SetAudio gets a reference to the given string and assigns it to the Audio field.
func (o *RiskEvaluationEventBrowser) SetAudio(v string) {
	o.Audio = &v
}

func (o RiskEvaluationEventBrowser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UserAgent != nil {
		toSerialize["userAgent"] = o.UserAgent
	}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	if o.ColorDepth != nil {
		toSerialize["colorDepth"] = o.ColorDepth
	}
	if o.DeviceMemory != nil {
		toSerialize["deviceMemory"] = o.DeviceMemory
	}
	if o.HardwareConcurrency != nil {
		toSerialize["hardwareConcurrency"] = o.HardwareConcurrency
	}
	if o.ScreenResolution != nil {
		toSerialize["screenResolution"] = o.ScreenResolution
	}
	if o.AvailableScreenResolution != nil {
		toSerialize["availableScreenResolution"] = o.AvailableScreenResolution
	}
	if o.TimezoneOffset != nil {
		toSerialize["timezoneOffset"] = o.TimezoneOffset
	}
	if o.Timezone != nil {
		toSerialize["timezone"] = o.Timezone
	}
	if o.SessionStorage != nil {
		toSerialize["sessionStorage"] = o.SessionStorage
	}
	if o.LocalStorage != nil {
		toSerialize["localStorage"] = o.LocalStorage
	}
	if o.IndexedDb != nil {
		toSerialize["indexedDb"] = o.IndexedDb
	}
	if o.AddBehaviour != nil {
		toSerialize["addBehaviour"] = o.AddBehaviour
	}
	if o.OpenDatabase != nil {
		toSerialize["openDatabase"] = o.OpenDatabase
	}
	if o.CpuClass != nil {
		toSerialize["cpuClass"] = o.CpuClass
	}
	if o.Platform != nil {
		toSerialize["platform"] = o.Platform
	}
	if o.Plugins != nil {
		toSerialize["plugins"] = o.Plugins
	}
	if o.WebglVendorAndRenderer != nil {
		toSerialize["webglVendorAndRenderer"] = o.WebglVendorAndRenderer
	}
	if o.Webgl != nil {
		toSerialize["webgl"] = o.Webgl
	}
	if o.AdBlock != nil {
		toSerialize["adBlock"] = o.AdBlock
	}
	if o.HasLiedLanguages != nil {
		toSerialize["hasLiedLanguages"] = o.HasLiedLanguages
	}
	if o.HasLiedResolution != nil {
		toSerialize["hasLiedResolution"] = o.HasLiedResolution
	}
	if o.HasLiedOs != nil {
		toSerialize["hasLiedOs"] = o.HasLiedOs
	}
	if o.HasLiedBrowser != nil {
		toSerialize["hasLiedBrowser"] = o.HasLiedBrowser
	}
	if o.TouchSupport != nil {
		toSerialize["touchSupport"] = o.TouchSupport
	}
	if o.Fonts != nil {
		toSerialize["fonts"] = o.Fonts
	}
	if o.Audio != nil {
		toSerialize["audio"] = o.Audio
	}
	return json.Marshal(toSerialize)
}

type NullableRiskEvaluationEventBrowser struct {
	value *RiskEvaluationEventBrowser
	isSet bool
}

func (v NullableRiskEvaluationEventBrowser) Get() *RiskEvaluationEventBrowser {
	return v.value
}

func (v *NullableRiskEvaluationEventBrowser) Set(val *RiskEvaluationEventBrowser) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskEvaluationEventBrowser) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskEvaluationEventBrowser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskEvaluationEventBrowser(val *RiskEvaluationEventBrowser) *NullableRiskEvaluationEventBrowser {
	return &NullableRiskEvaluationEventBrowser{value: val, isSet: true}
}

func (v NullableRiskEvaluationEventBrowser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskEvaluationEventBrowser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


