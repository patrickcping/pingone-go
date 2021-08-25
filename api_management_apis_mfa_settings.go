/*
 * PingOne Platform API
 *
 * A bare-bones collection for the PingOne API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pingone

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// ManagementAPIsMFASettingsApiService ManagementAPIsMFASettingsApi service
type ManagementAPIsMFASettingsApiService service

type ApiV1EnvironmentsEnvIDMfaSettingsDeleteRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsMFASettingsApiService
	envID string
	authorization *string
}

func (r ApiV1EnvironmentsEnvIDMfaSettingsDeleteRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDMfaSettingsDeleteRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1EnvironmentsEnvIDMfaSettingsDeleteRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDMfaSettingsDeleteExecute(r)
}

/*
 * V1EnvironmentsEnvIDMfaSettingsDelete RESET MFA Settings
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @return ApiV1EnvironmentsEnvIDMfaSettingsDeleteRequest
 */
func (a *ManagementAPIsMFASettingsApiService) V1EnvironmentsEnvIDMfaSettingsDelete(ctx _context.Context, envID string) ApiV1EnvironmentsEnvIDMfaSettingsDeleteRequest {
	return ApiV1EnvironmentsEnvIDMfaSettingsDeleteRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsMFASettingsApiService) V1EnvironmentsEnvIDMfaSettingsDeleteExecute(r ApiV1EnvironmentsEnvIDMfaSettingsDeleteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsMFASettingsApiService.V1EnvironmentsEnvIDMfaSettingsDelete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/mfaSettings"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.authorization != nil {
		localVarHeaderParams["Authorization"] = parameterToString(*r.authorization, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiV1EnvironmentsEnvIDMfaSettingsGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsMFASettingsApiService
	envID string
	authorization *string
}

func (r ApiV1EnvironmentsEnvIDMfaSettingsGetRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDMfaSettingsGetRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1EnvironmentsEnvIDMfaSettingsGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDMfaSettingsGetExecute(r)
}

/*
 * V1EnvironmentsEnvIDMfaSettingsGet READ MFA Settings
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @return ApiV1EnvironmentsEnvIDMfaSettingsGetRequest
 */
func (a *ManagementAPIsMFASettingsApiService) V1EnvironmentsEnvIDMfaSettingsGet(ctx _context.Context, envID string) ApiV1EnvironmentsEnvIDMfaSettingsGetRequest {
	return ApiV1EnvironmentsEnvIDMfaSettingsGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsMFASettingsApiService) V1EnvironmentsEnvIDMfaSettingsGetExecute(r ApiV1EnvironmentsEnvIDMfaSettingsGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsMFASettingsApiService.V1EnvironmentsEnvIDMfaSettingsGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/mfaSettings"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.authorization != nil {
		localVarHeaderParams["Authorization"] = parameterToString(*r.authorization, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiV1EnvironmentsEnvIDMfaSettingsPutRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsMFASettingsApiService
	envID string
	authorization *string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDMfaSettingsPutRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDMfaSettingsPutRequest {
	r.authorization = &authorization
	return r
}
func (r ApiV1EnvironmentsEnvIDMfaSettingsPutRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDMfaSettingsPutRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDMfaSettingsPutRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDMfaSettingsPutExecute(r)
}

/*
 * V1EnvironmentsEnvIDMfaSettingsPut UPDATE MFA Settings
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @return ApiV1EnvironmentsEnvIDMfaSettingsPutRequest
 */
func (a *ManagementAPIsMFASettingsApiService) V1EnvironmentsEnvIDMfaSettingsPut(ctx _context.Context, envID string) ApiV1EnvironmentsEnvIDMfaSettingsPutRequest {
	return ApiV1EnvironmentsEnvIDMfaSettingsPutRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsMFASettingsApiService) V1EnvironmentsEnvIDMfaSettingsPutExecute(r ApiV1EnvironmentsEnvIDMfaSettingsPutRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsMFASettingsApiService.V1EnvironmentsEnvIDMfaSettingsPut")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/mfaSettings"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.authorization != nil {
		localVarHeaderParams["Authorization"] = parameterToString(*r.authorization, "")
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
