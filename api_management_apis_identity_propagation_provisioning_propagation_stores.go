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

// ManagementAPIsIdentityPropagationProvisioningPropagationStoresApiService ManagementAPIsIdentityPropagationProvisioningPropagationStoresApi service
type ManagementAPIsIdentityPropagationProvisioningPropagationStoresApiService service

type ApiV1EnvironmentsEnvIDPropagationStoresConnectionStatusPostRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsIdentityPropagationProvisioningPropagationStoresApiService
	envID string
	contentType *string
	authorization *string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDPropagationStoresConnectionStatusPostRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDPropagationStoresConnectionStatusPostRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDPropagationStoresConnectionStatusPostRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDPropagationStoresConnectionStatusPostRequest {
	r.authorization = &authorization
	return r
}
func (r ApiV1EnvironmentsEnvIDPropagationStoresConnectionStatusPostRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDPropagationStoresConnectionStatusPostRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDPropagationStoresConnectionStatusPostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDPropagationStoresConnectionStatusPostExecute(r)
}

/*
 * V1EnvironmentsEnvIDPropagationStoresConnectionStatusPost TEST Connection Configuration
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @return ApiV1EnvironmentsEnvIDPropagationStoresConnectionStatusPostRequest
 */
func (a *ManagementAPIsIdentityPropagationProvisioningPropagationStoresApiService) V1EnvironmentsEnvIDPropagationStoresConnectionStatusPost(ctx _context.Context, envID string) ApiV1EnvironmentsEnvIDPropagationStoresConnectionStatusPostRequest {
	return ApiV1EnvironmentsEnvIDPropagationStoresConnectionStatusPostRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsIdentityPropagationProvisioningPropagationStoresApiService) V1EnvironmentsEnvIDPropagationStoresConnectionStatusPostExecute(r ApiV1EnvironmentsEnvIDPropagationStoresConnectionStatusPostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsIdentityPropagationProvisioningPropagationStoresApiService.V1EnvironmentsEnvIDPropagationStoresConnectionStatusPost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/propagation/stores/connection/status"
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
		localVarHeaderParams["Content-Type"] = parameterToString(*r.contentType, "")
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

type ApiV1EnvironmentsEnvIDPropagationStoresGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsIdentityPropagationProvisioningPropagationStoresApiService
	envID string
	accept *string
	authorization *string
}

func (r ApiV1EnvironmentsEnvIDPropagationStoresGetRequest) Accept(accept string) ApiV1EnvironmentsEnvIDPropagationStoresGetRequest {
	r.accept = &accept
	return r
}
func (r ApiV1EnvironmentsEnvIDPropagationStoresGetRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDPropagationStoresGetRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1EnvironmentsEnvIDPropagationStoresGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDPropagationStoresGetExecute(r)
}

/*
 * V1EnvironmentsEnvIDPropagationStoresGet READ All Stores
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @return ApiV1EnvironmentsEnvIDPropagationStoresGetRequest
 */
func (a *ManagementAPIsIdentityPropagationProvisioningPropagationStoresApiService) V1EnvironmentsEnvIDPropagationStoresGet(ctx _context.Context, envID string) ApiV1EnvironmentsEnvIDPropagationStoresGetRequest {
	return ApiV1EnvironmentsEnvIDPropagationStoresGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsIdentityPropagationProvisioningPropagationStoresApiService) V1EnvironmentsEnvIDPropagationStoresGetExecute(r ApiV1EnvironmentsEnvIDPropagationStoresGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsIdentityPropagationProvisioningPropagationStoresApiService.V1EnvironmentsEnvIDPropagationStoresGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/propagation/stores"
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
	if r.accept != nil {
		localVarHeaderParams["Accept"] = parameterToString(*r.accept, "")
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

type ApiV1EnvironmentsEnvIDPropagationStoresPostRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsIdentityPropagationProvisioningPropagationStoresApiService
	envID string
	contentType *string
	authorization *string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDPropagationStoresPostRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDPropagationStoresPostRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDPropagationStoresPostRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDPropagationStoresPostRequest {
	r.authorization = &authorization
	return r
}
func (r ApiV1EnvironmentsEnvIDPropagationStoresPostRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDPropagationStoresPostRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDPropagationStoresPostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDPropagationStoresPostExecute(r)
}

/*
 * V1EnvironmentsEnvIDPropagationStoresPost CREATE Store (Aquera)
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @return ApiV1EnvironmentsEnvIDPropagationStoresPostRequest
 */
func (a *ManagementAPIsIdentityPropagationProvisioningPropagationStoresApiService) V1EnvironmentsEnvIDPropagationStoresPost(ctx _context.Context, envID string) ApiV1EnvironmentsEnvIDPropagationStoresPostRequest {
	return ApiV1EnvironmentsEnvIDPropagationStoresPostRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsIdentityPropagationProvisioningPropagationStoresApiService) V1EnvironmentsEnvIDPropagationStoresPostExecute(r ApiV1EnvironmentsEnvIDPropagationStoresPostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsIdentityPropagationProvisioningPropagationStoresApiService.V1EnvironmentsEnvIDPropagationStoresPost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/propagation/stores"
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
		localVarHeaderParams["Content-Type"] = parameterToString(*r.contentType, "")
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

type ApiV1EnvironmentsEnvIDPropagationStoresStoreIDDeleteRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsIdentityPropagationProvisioningPropagationStoresApiService
	envID string
	storeID string
	accept *string
	authorization *string
}

func (r ApiV1EnvironmentsEnvIDPropagationStoresStoreIDDeleteRequest) Accept(accept string) ApiV1EnvironmentsEnvIDPropagationStoresStoreIDDeleteRequest {
	r.accept = &accept
	return r
}
func (r ApiV1EnvironmentsEnvIDPropagationStoresStoreIDDeleteRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDPropagationStoresStoreIDDeleteRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1EnvironmentsEnvIDPropagationStoresStoreIDDeleteRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDPropagationStoresStoreIDDeleteExecute(r)
}

/*
 * V1EnvironmentsEnvIDPropagationStoresStoreIDDelete DELETE Store
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @param storeID
 * @return ApiV1EnvironmentsEnvIDPropagationStoresStoreIDDeleteRequest
 */
func (a *ManagementAPIsIdentityPropagationProvisioningPropagationStoresApiService) V1EnvironmentsEnvIDPropagationStoresStoreIDDelete(ctx _context.Context, envID string, storeID string) ApiV1EnvironmentsEnvIDPropagationStoresStoreIDDeleteRequest {
	return ApiV1EnvironmentsEnvIDPropagationStoresStoreIDDeleteRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		storeID: storeID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsIdentityPropagationProvisioningPropagationStoresApiService) V1EnvironmentsEnvIDPropagationStoresStoreIDDeleteExecute(r ApiV1EnvironmentsEnvIDPropagationStoresStoreIDDeleteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsIdentityPropagationProvisioningPropagationStoresApiService.V1EnvironmentsEnvIDPropagationStoresStoreIDDelete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/propagation/stores/{storeID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"storeID"+"}", _neturl.PathEscape(parameterToString(r.storeID, "")), -1)

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
	if r.accept != nil {
		localVarHeaderParams["Accept"] = parameterToString(*r.accept, "")
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

type ApiV1EnvironmentsEnvIDPropagationStoresStoreIDGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsIdentityPropagationProvisioningPropagationStoresApiService
	envID string
	storeID string
	accept *string
	authorization *string
}

func (r ApiV1EnvironmentsEnvIDPropagationStoresStoreIDGetRequest) Accept(accept string) ApiV1EnvironmentsEnvIDPropagationStoresStoreIDGetRequest {
	r.accept = &accept
	return r
}
func (r ApiV1EnvironmentsEnvIDPropagationStoresStoreIDGetRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDPropagationStoresStoreIDGetRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1EnvironmentsEnvIDPropagationStoresStoreIDGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDPropagationStoresStoreIDGetExecute(r)
}

/*
 * V1EnvironmentsEnvIDPropagationStoresStoreIDGet READ One Store
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @param storeID
 * @return ApiV1EnvironmentsEnvIDPropagationStoresStoreIDGetRequest
 */
func (a *ManagementAPIsIdentityPropagationProvisioningPropagationStoresApiService) V1EnvironmentsEnvIDPropagationStoresStoreIDGet(ctx _context.Context, envID string, storeID string) ApiV1EnvironmentsEnvIDPropagationStoresStoreIDGetRequest {
	return ApiV1EnvironmentsEnvIDPropagationStoresStoreIDGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		storeID: storeID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsIdentityPropagationProvisioningPropagationStoresApiService) V1EnvironmentsEnvIDPropagationStoresStoreIDGetExecute(r ApiV1EnvironmentsEnvIDPropagationStoresStoreIDGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsIdentityPropagationProvisioningPropagationStoresApiService.V1EnvironmentsEnvIDPropagationStoresStoreIDGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/propagation/stores/{storeID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"storeID"+"}", _neturl.PathEscape(parameterToString(r.storeID, "")), -1)

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
	if r.accept != nil {
		localVarHeaderParams["Accept"] = parameterToString(*r.accept, "")
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

type ApiV1EnvironmentsEnvIDPropagationStoresStoreIDPutRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsIdentityPropagationProvisioningPropagationStoresApiService
	envID string
	storeID string
	contentType *string
	authorization *string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDPropagationStoresStoreIDPutRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDPropagationStoresStoreIDPutRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDPropagationStoresStoreIDPutRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDPropagationStoresStoreIDPutRequest {
	r.authorization = &authorization
	return r
}
func (r ApiV1EnvironmentsEnvIDPropagationStoresStoreIDPutRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDPropagationStoresStoreIDPutRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDPropagationStoresStoreIDPutRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDPropagationStoresStoreIDPutExecute(r)
}

/*
 * V1EnvironmentsEnvIDPropagationStoresStoreIDPut UPDATE Store
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @param storeID
 * @return ApiV1EnvironmentsEnvIDPropagationStoresStoreIDPutRequest
 */
func (a *ManagementAPIsIdentityPropagationProvisioningPropagationStoresApiService) V1EnvironmentsEnvIDPropagationStoresStoreIDPut(ctx _context.Context, envID string, storeID string) ApiV1EnvironmentsEnvIDPropagationStoresStoreIDPutRequest {
	return ApiV1EnvironmentsEnvIDPropagationStoresStoreIDPutRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		storeID: storeID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsIdentityPropagationProvisioningPropagationStoresApiService) V1EnvironmentsEnvIDPropagationStoresStoreIDPutExecute(r ApiV1EnvironmentsEnvIDPropagationStoresStoreIDPutRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsIdentityPropagationProvisioningPropagationStoresApiService.V1EnvironmentsEnvIDPropagationStoresStoreIDPut")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/propagation/stores/{storeID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"storeID"+"}", _neturl.PathEscape(parameterToString(r.storeID, "")), -1)

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
		localVarHeaderParams["Content-Type"] = parameterToString(*r.contentType, "")
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
