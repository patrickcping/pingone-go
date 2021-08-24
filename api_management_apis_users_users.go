/*
 * PingOne Platform API
 *
 * A bare-bones collection for the PingOne API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

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

// ManagementAPIsUsersUsersApiService ManagementAPIsUsersUsersApi service
type ManagementAPIsUsersUsersApiService service

type ApiV1EnvironmentsEnvIDUsersPostRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsUsersUsersApiService
	envID string
	contentType *string
	authorization *string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDUsersPostRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDUsersPostRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDUsersPostRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDUsersPostRequest {
	r.authorization = &authorization
	return r
}
func (r ApiV1EnvironmentsEnvIDUsersPostRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDUsersPostRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDUsersPostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDUsersPostExecute(r)
}

/*
 * V1EnvironmentsEnvIDUsersPost CREATE User (Import)
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @return ApiV1EnvironmentsEnvIDUsersPostRequest
 */
func (a *ManagementAPIsUsersUsersApiService) V1EnvironmentsEnvIDUsersPost(ctx _context.Context, envID string) ApiV1EnvironmentsEnvIDUsersPostRequest {
	return ApiV1EnvironmentsEnvIDUsersPostRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsUsersUsersApiService) V1EnvironmentsEnvIDUsersPostExecute(r ApiV1EnvironmentsEnvIDUsersPostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsUsersUsersApiService.V1EnvironmentsEnvIDUsersPost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/users"
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
	if r.contentType != nil {
		localVarHeaderParams["content-type"] = parameterToString(*r.contentType, "")
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

type ApiV1EnvironmentsEnvIDUsersUserIDDeleteRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsUsersUsersApiService
	envID string
	userID string
	authorization *string
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDDeleteRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDUsersUserIDDeleteRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDDeleteRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDUsersUserIDDeleteExecute(r)
}

/*
 * V1EnvironmentsEnvIDUsersUserIDDelete DELETE User
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @param userID
 * @return ApiV1EnvironmentsEnvIDUsersUserIDDeleteRequest
 */
func (a *ManagementAPIsUsersUsersApiService) V1EnvironmentsEnvIDUsersUserIDDelete(ctx _context.Context, envID string, userID string) ApiV1EnvironmentsEnvIDUsersUserIDDeleteRequest {
	return ApiV1EnvironmentsEnvIDUsersUserIDDeleteRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		userID: userID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsUsersUsersApiService) V1EnvironmentsEnvIDUsersUserIDDeleteExecute(r ApiV1EnvironmentsEnvIDUsersUserIDDeleteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsUsersUsersApiService.V1EnvironmentsEnvIDUsersUserIDDelete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/users/{userID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userID"+"}", _neturl.PathEscape(parameterToString(r.userID, "")), -1)

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

type ApiV1EnvironmentsEnvIDUsersUserIDIdentityProviderGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsUsersUsersApiService
	envID string
	userID string
	authorization *string
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDIdentityProviderGetRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDUsersUserIDIdentityProviderGetRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDIdentityProviderGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDUsersUserIDIdentityProviderGetExecute(r)
}

/*
 * V1EnvironmentsEnvIDUsersUserIDIdentityProviderGet READ User Identity Provider
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @param userID
 * @return ApiV1EnvironmentsEnvIDUsersUserIDIdentityProviderGetRequest
 */
func (a *ManagementAPIsUsersUsersApiService) V1EnvironmentsEnvIDUsersUserIDIdentityProviderGet(ctx _context.Context, envID string, userID string) ApiV1EnvironmentsEnvIDUsersUserIDIdentityProviderGetRequest {
	return ApiV1EnvironmentsEnvIDUsersUserIDIdentityProviderGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		userID: userID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsUsersUsersApiService) V1EnvironmentsEnvIDUsersUserIDIdentityProviderGetExecute(r ApiV1EnvironmentsEnvIDUsersUserIDIdentityProviderGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsUsersUsersApiService.V1EnvironmentsEnvIDUsersUserIDIdentityProviderGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/users/{userID}/identityProvider"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userID"+"}", _neturl.PathEscape(parameterToString(r.userID, "")), -1)

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

type ApiV1EnvironmentsEnvIDUsersUserIDIdentityProviderPutRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsUsersUsersApiService
	envID string
	userID string
	authorization *string
	contentType *string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDIdentityProviderPutRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDUsersUserIDIdentityProviderPutRequest {
	r.authorization = &authorization
	return r
}
func (r ApiV1EnvironmentsEnvIDUsersUserIDIdentityProviderPutRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDUsersUserIDIdentityProviderPutRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDUsersUserIDIdentityProviderPutRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDUsersUserIDIdentityProviderPutRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDIdentityProviderPutRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDUsersUserIDIdentityProviderPutExecute(r)
}

/*
 * V1EnvironmentsEnvIDUsersUserIDIdentityProviderPut UPDATE User Identity Provider
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @param userID
 * @return ApiV1EnvironmentsEnvIDUsersUserIDIdentityProviderPutRequest
 */
func (a *ManagementAPIsUsersUsersApiService) V1EnvironmentsEnvIDUsersUserIDIdentityProviderPut(ctx _context.Context, envID string, userID string) ApiV1EnvironmentsEnvIDUsersUserIDIdentityProviderPutRequest {
	return ApiV1EnvironmentsEnvIDUsersUserIDIdentityProviderPutRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		userID: userID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsUsersUsersApiService) V1EnvironmentsEnvIDUsersUserIDIdentityProviderPutExecute(r ApiV1EnvironmentsEnvIDUsersUserIDIdentityProviderPutRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsUsersUsersApiService.V1EnvironmentsEnvIDUsersUserIDIdentityProviderPut")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/users/{userID}/identityProvider"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userID"+"}", _neturl.PathEscape(parameterToString(r.userID, "")), -1)

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
	if r.contentType != nil {
		localVarHeaderParams["Content-Type"] = parameterToString(*r.contentType, "")
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

type ApiV1EnvironmentsEnvIDUsersUserIDPatchRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsUsersUsersApiService
	envID string
	userID string
	contentType *string
	authorization *string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDPatchRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDUsersUserIDPatchRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDUsersUserIDPatchRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDUsersUserIDPatchRequest {
	r.authorization = &authorization
	return r
}
func (r ApiV1EnvironmentsEnvIDUsersUserIDPatchRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDUsersUserIDPatchRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDPatchRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDUsersUserIDPatchExecute(r)
}

/*
 * V1EnvironmentsEnvIDUsersUserIDPatch UPDATE User (Patch)
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @param userID
 * @return ApiV1EnvironmentsEnvIDUsersUserIDPatchRequest
 */
func (a *ManagementAPIsUsersUsersApiService) V1EnvironmentsEnvIDUsersUserIDPatch(ctx _context.Context, envID string, userID string) ApiV1EnvironmentsEnvIDUsersUserIDPatchRequest {
	return ApiV1EnvironmentsEnvIDUsersUserIDPatchRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		userID: userID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsUsersUsersApiService) V1EnvironmentsEnvIDUsersUserIDPatchExecute(r ApiV1EnvironmentsEnvIDUsersUserIDPatchRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsUsersUsersApiService.V1EnvironmentsEnvIDUsersUserIDPatch")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/users/{userID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userID"+"}", _neturl.PathEscape(parameterToString(r.userID, "")), -1)

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
	if r.contentType != nil {
		localVarHeaderParams["content-type"] = parameterToString(*r.contentType, "")
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

type ApiV1EnvironmentsEnvIDUsersUserIDPutRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsUsersUsersApiService
	envID string
	userID string
	contentType *string
	authorization *string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDPutRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDUsersUserIDPutRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDUsersUserIDPutRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDUsersUserIDPutRequest {
	r.authorization = &authorization
	return r
}
func (r ApiV1EnvironmentsEnvIDUsersUserIDPutRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDUsersUserIDPutRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDPutRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDUsersUserIDPutExecute(r)
}

/*
 * V1EnvironmentsEnvIDUsersUserIDPut UPDATE User (Put)
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @param userID
 * @return ApiV1EnvironmentsEnvIDUsersUserIDPutRequest
 */
func (a *ManagementAPIsUsersUsersApiService) V1EnvironmentsEnvIDUsersUserIDPut(ctx _context.Context, envID string, userID string) ApiV1EnvironmentsEnvIDUsersUserIDPutRequest {
	return ApiV1EnvironmentsEnvIDUsersUserIDPutRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		userID: userID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsUsersUsersApiService) V1EnvironmentsEnvIDUsersUserIDPutExecute(r ApiV1EnvironmentsEnvIDUsersUserIDPutRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsUsersUsersApiService.V1EnvironmentsEnvIDUsersUserIDPut")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/users/{userID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userID"+"}", _neturl.PathEscape(parameterToString(r.userID, "")), -1)

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
	if r.contentType != nil {
		localVarHeaderParams["content-type"] = parameterToString(*r.contentType, "")
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

type ApiV1EnvironmentsEnvIDUsersUserIDVerifyStatusGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsUsersUsersApiService
	envID string
	userID string
	contentType *string
	authorization *string
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDVerifyStatusGetRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDUsersUserIDVerifyStatusGetRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDUsersUserIDVerifyStatusGetRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDUsersUserIDVerifyStatusGetRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDVerifyStatusGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDUsersUserIDVerifyStatusGetExecute(r)
}

/*
 * V1EnvironmentsEnvIDUsersUserIDVerifyStatusGet READ user verification status
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @param userID
 * @return ApiV1EnvironmentsEnvIDUsersUserIDVerifyStatusGetRequest
 */
func (a *ManagementAPIsUsersUsersApiService) V1EnvironmentsEnvIDUsersUserIDVerifyStatusGet(ctx _context.Context, envID string, userID string) ApiV1EnvironmentsEnvIDUsersUserIDVerifyStatusGetRequest {
	return ApiV1EnvironmentsEnvIDUsersUserIDVerifyStatusGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		userID: userID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsUsersUsersApiService) V1EnvironmentsEnvIDUsersUserIDVerifyStatusGetExecute(r ApiV1EnvironmentsEnvIDUsersUserIDVerifyStatusGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsUsersUsersApiService.V1EnvironmentsEnvIDUsersUserIDVerifyStatusGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/users/{userID}/verifyStatus"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userID"+"}", _neturl.PathEscape(parameterToString(r.userID, "")), -1)

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
	if r.contentType != nil {
		localVarHeaderParams["content-type"] = parameterToString(*r.contentType, "")
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

type ApiV1EnvironmentsEnvIDUsersUserIDVerifyStatusPutRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsUsersUsersApiService
	envID string
	userID string
	contentType *string
	authorization *string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDVerifyStatusPutRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDUsersUserIDVerifyStatusPutRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDUsersUserIDVerifyStatusPutRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDUsersUserIDVerifyStatusPutRequest {
	r.authorization = &authorization
	return r
}
func (r ApiV1EnvironmentsEnvIDUsersUserIDVerifyStatusPutRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDUsersUserIDVerifyStatusPutRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDVerifyStatusPutRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDUsersUserIDVerifyStatusPutExecute(r)
}

/*
 * V1EnvironmentsEnvIDUsersUserIDVerifyStatusPut UPDATE user verification status
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @param userID
 * @return ApiV1EnvironmentsEnvIDUsersUserIDVerifyStatusPutRequest
 */
func (a *ManagementAPIsUsersUsersApiService) V1EnvironmentsEnvIDUsersUserIDVerifyStatusPut(ctx _context.Context, envID string, userID string) ApiV1EnvironmentsEnvIDUsersUserIDVerifyStatusPutRequest {
	return ApiV1EnvironmentsEnvIDUsersUserIDVerifyStatusPutRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		userID: userID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsUsersUsersApiService) V1EnvironmentsEnvIDUsersUserIDVerifyStatusPutExecute(r ApiV1EnvironmentsEnvIDUsersUserIDVerifyStatusPutRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsUsersUsersApiService.V1EnvironmentsEnvIDUsersUserIDVerifyStatusPut")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/users/{userID}/verifyStatus"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userID"+"}", _neturl.PathEscape(parameterToString(r.userID, "")), -1)

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
	if r.contentType != nil {
		localVarHeaderParams["content-type"] = parameterToString(*r.contentType, "")
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
