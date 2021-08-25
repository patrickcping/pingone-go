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

// AuthenticationAPIsExternalAuthenticationApiService AuthenticationAPIsExternalAuthenticationApi service
type AuthenticationAPIsExternalAuthenticationApiService service

type ApiV1EnvIDRpAuthenticateGetRequest struct {
	ctx _context.Context
	ApiService *AuthenticationAPIsExternalAuthenticationApiService
	envID string
	providerId *string
	flowId *string
}

func (r ApiV1EnvIDRpAuthenticateGetRequest) ProviderId(providerId string) ApiV1EnvIDRpAuthenticateGetRequest {
	r.providerId = &providerId
	return r
}
func (r ApiV1EnvIDRpAuthenticateGetRequest) FlowId(flowId string) ApiV1EnvIDRpAuthenticateGetRequest {
	r.flowId = &flowId
	return r
}

func (r ApiV1EnvIDRpAuthenticateGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvIDRpAuthenticateGetExecute(r)
}

/*
 * V1EnvIDRpAuthenticateGet READ External Authentication Initialization
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @return ApiV1EnvIDRpAuthenticateGetRequest
 */
func (a *AuthenticationAPIsExternalAuthenticationApiService) V1EnvIDRpAuthenticateGet(ctx _context.Context, envID string) ApiV1EnvIDRpAuthenticateGetRequest {
	return ApiV1EnvIDRpAuthenticateGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
	}
}

/*
 * Execute executes the request
 */
func (a *AuthenticationAPIsExternalAuthenticationApiService) V1EnvIDRpAuthenticateGetExecute(r ApiV1EnvIDRpAuthenticateGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthenticationAPIsExternalAuthenticationApiService.V1EnvIDRpAuthenticateGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/{envID}/rp/authenticate"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.providerId != nil {
		localVarQueryParams.Add("providerId", parameterToString(*r.providerId, ""))
	}
	if r.flowId != nil {
		localVarQueryParams.Add("flowId", parameterToString(*r.flowId, ""))
	}
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

type ApiV1EnvIDRpCallbackProviderTypeGetRequest struct {
	ctx _context.Context
	ApiService *AuthenticationAPIsExternalAuthenticationApiService
	envID string
	providerType string
	code *string
	state *string
}

func (r ApiV1EnvIDRpCallbackProviderTypeGetRequest) Code(code string) ApiV1EnvIDRpCallbackProviderTypeGetRequest {
	r.code = &code
	return r
}
func (r ApiV1EnvIDRpCallbackProviderTypeGetRequest) State(state string) ApiV1EnvIDRpCallbackProviderTypeGetRequest {
	r.state = &state
	return r
}

func (r ApiV1EnvIDRpCallbackProviderTypeGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvIDRpCallbackProviderTypeGetExecute(r)
}

/*
 * V1EnvIDRpCallbackProviderTypeGet READ External Authentication Callback
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @param providerType
 * @return ApiV1EnvIDRpCallbackProviderTypeGetRequest
 */
func (a *AuthenticationAPIsExternalAuthenticationApiService) V1EnvIDRpCallbackProviderTypeGet(ctx _context.Context, envID string, providerType string) ApiV1EnvIDRpCallbackProviderTypeGetRequest {
	return ApiV1EnvIDRpCallbackProviderTypeGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		providerType: providerType,
	}
}

/*
 * Execute executes the request
 */
func (a *AuthenticationAPIsExternalAuthenticationApiService) V1EnvIDRpCallbackProviderTypeGetExecute(r ApiV1EnvIDRpCallbackProviderTypeGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthenticationAPIsExternalAuthenticationApiService.V1EnvIDRpCallbackProviderTypeGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/{envID}/rp/callback/{providerType}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"providerType"+"}", _neturl.PathEscape(parameterToString(r.providerType, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.code != nil {
		localVarQueryParams.Add("code", parameterToString(*r.code, ""))
	}
	if r.state != nil {
		localVarQueryParams.Add("state", parameterToString(*r.state, ""))
	}
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

type ApiV1EnvironmentsEnvIDExternalAuthenticationsExtAuthIDGetRequest struct {
	ctx _context.Context
	ApiService *AuthenticationAPIsExternalAuthenticationApiService
	envID string
	extAuthID string
	cookie *string
}

func (r ApiV1EnvironmentsEnvIDExternalAuthenticationsExtAuthIDGetRequest) Cookie(cookie string) ApiV1EnvironmentsEnvIDExternalAuthenticationsExtAuthIDGetRequest {
	r.cookie = &cookie
	return r
}

func (r ApiV1EnvironmentsEnvIDExternalAuthenticationsExtAuthIDGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDExternalAuthenticationsExtAuthIDGetExecute(r)
}

/*
 * V1EnvironmentsEnvIDExternalAuthenticationsExtAuthIDGet READ External Authentication Status
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @param extAuthID
 * @return ApiV1EnvironmentsEnvIDExternalAuthenticationsExtAuthIDGetRequest
 */
func (a *AuthenticationAPIsExternalAuthenticationApiService) V1EnvironmentsEnvIDExternalAuthenticationsExtAuthIDGet(ctx _context.Context, envID string, extAuthID string) ApiV1EnvironmentsEnvIDExternalAuthenticationsExtAuthIDGetRequest {
	return ApiV1EnvironmentsEnvIDExternalAuthenticationsExtAuthIDGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		extAuthID: extAuthID,
	}
}

/*
 * Execute executes the request
 */
func (a *AuthenticationAPIsExternalAuthenticationApiService) V1EnvironmentsEnvIDExternalAuthenticationsExtAuthIDGetExecute(r ApiV1EnvironmentsEnvIDExternalAuthenticationsExtAuthIDGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthenticationAPIsExternalAuthenticationApiService.V1EnvironmentsEnvIDExternalAuthenticationsExtAuthIDGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/externalAuthentications/{extAuthID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"extAuthID"+"}", _neturl.PathEscape(parameterToString(r.extAuthID, "")), -1)

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
	if r.cookie != nil {
		localVarHeaderParams["Cookie"] = parameterToString(*r.cookie, "")
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
