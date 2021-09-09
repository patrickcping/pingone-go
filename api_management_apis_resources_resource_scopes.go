/*
PingOne Platform API - Management

A bare-bones collection for the PingOne API

API version: 1.0.0
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

// ManagementAPIsResourcesResourceScopesApiService ManagementAPIsResourcesResourceScopesApi service
type ManagementAPIsResourcesResourceScopesApiService service

type ApiV1EnvironmentsEnvIDResourcesResourceIDScopesGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsResourcesResourceScopesApiService
	envID string
	resourceID string
}


func (r ApiV1EnvironmentsEnvIDResourcesResourceIDScopesGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDResourcesResourceIDScopesGetExecute(r)
}

/*
V1EnvironmentsEnvIDResourcesResourceIDScopesGet READ All Scopes (Resource)

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param resourceID
 @return ApiV1EnvironmentsEnvIDResourcesResourceIDScopesGetRequest
*/
func (a *ManagementAPIsResourcesResourceScopesApiService) V1EnvironmentsEnvIDResourcesResourceIDScopesGet(ctx _context.Context, envID string, resourceID string) ApiV1EnvironmentsEnvIDResourcesResourceIDScopesGetRequest {
	return ApiV1EnvironmentsEnvIDResourcesResourceIDScopesGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		resourceID: resourceID,
	}
}

// Execute executes the request
func (a *ManagementAPIsResourcesResourceScopesApiService) V1EnvironmentsEnvIDResourcesResourceIDScopesGetExecute(r ApiV1EnvironmentsEnvIDResourcesResourceIDScopesGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsResourcesResourceScopesApiService.V1EnvironmentsEnvIDResourcesResourceIDScopesGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/resources/{resourceID}/scopes"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resourceID"+"}", _neturl.PathEscape(parameterToString(r.resourceID, "")), -1)

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
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiV1EnvironmentsEnvIDResourcesResourceIDScopesPostRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsResourcesResourceScopesApiService
	envID string
	resourceID string
	contentType *string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDResourcesResourceIDScopesPostRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDResourcesResourceIDScopesPostRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDResourcesResourceIDScopesPostRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDResourcesResourceIDScopesPostRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDResourcesResourceIDScopesPostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDResourcesResourceIDScopesPostExecute(r)
}

/*
V1EnvironmentsEnvIDResourcesResourceIDScopesPost CREATE PingOne access control scope

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param resourceID
 @return ApiV1EnvironmentsEnvIDResourcesResourceIDScopesPostRequest
*/
func (a *ManagementAPIsResourcesResourceScopesApiService) V1EnvironmentsEnvIDResourcesResourceIDScopesPost(ctx _context.Context, envID string, resourceID string) ApiV1EnvironmentsEnvIDResourcesResourceIDScopesPostRequest {
	return ApiV1EnvironmentsEnvIDResourcesResourceIDScopesPostRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		resourceID: resourceID,
	}
}

// Execute executes the request
func (a *ManagementAPIsResourcesResourceScopesApiService) V1EnvironmentsEnvIDResourcesResourceIDScopesPostExecute(r ApiV1EnvironmentsEnvIDResourcesResourceIDScopesPostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsResourcesResourceScopesApiService.V1EnvironmentsEnvIDResourcesResourceIDScopesPost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/resources/{resourceID}/scopes"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resourceID"+"}", _neturl.PathEscape(parameterToString(r.resourceID, "")), -1)

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
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiV1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDDeleteRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsResourcesResourceScopesApiService
	envID string
	resourceID string
	scopeID string
}


func (r ApiV1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDDeleteRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDDeleteExecute(r)
}

/*
V1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDDelete DELETE Scope

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param resourceID
 @param scopeID
 @return ApiV1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDDeleteRequest
*/
func (a *ManagementAPIsResourcesResourceScopesApiService) V1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDDelete(ctx _context.Context, envID string, resourceID string, scopeID string) ApiV1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDDeleteRequest {
	return ApiV1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDDeleteRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		resourceID: resourceID,
		scopeID: scopeID,
	}
}

// Execute executes the request
func (a *ManagementAPIsResourcesResourceScopesApiService) V1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDDeleteExecute(r ApiV1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDDeleteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsResourcesResourceScopesApiService.V1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDDelete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/resources/{resourceID}/scopes/{scopeID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resourceID"+"}", _neturl.PathEscape(parameterToString(r.resourceID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"scopeID"+"}", _neturl.PathEscape(parameterToString(r.scopeID, "")), -1)

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
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiV1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsResourcesResourceScopesApiService
	envID string
	resourceID string
	scopeID string
}


func (r ApiV1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDGetExecute(r)
}

/*
V1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDGet READ One Scope

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param resourceID
 @param scopeID
 @return ApiV1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDGetRequest
*/
func (a *ManagementAPIsResourcesResourceScopesApiService) V1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDGet(ctx _context.Context, envID string, resourceID string, scopeID string) ApiV1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDGetRequest {
	return ApiV1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		resourceID: resourceID,
		scopeID: scopeID,
	}
}

// Execute executes the request
func (a *ManagementAPIsResourcesResourceScopesApiService) V1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDGetExecute(r ApiV1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsResourcesResourceScopesApiService.V1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/resources/{resourceID}/scopes/{scopeID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resourceID"+"}", _neturl.PathEscape(parameterToString(r.resourceID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"scopeID"+"}", _neturl.PathEscape(parameterToString(r.scopeID, "")), -1)

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
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiV1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDPutRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsResourcesResourceScopesApiService
	envID string
	resourceID string
	scopeID string
	contentType *string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDPutRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDPutRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDPutRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDPutRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDPutRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDPutExecute(r)
}

/*
V1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDPut UPDATE PingOne access control scope

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param resourceID
 @param scopeID
 @return ApiV1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDPutRequest
*/
func (a *ManagementAPIsResourcesResourceScopesApiService) V1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDPut(ctx _context.Context, envID string, resourceID string, scopeID string) ApiV1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDPutRequest {
	return ApiV1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDPutRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		resourceID: resourceID,
		scopeID: scopeID,
	}
}

// Execute executes the request
func (a *ManagementAPIsResourcesResourceScopesApiService) V1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDPutExecute(r ApiV1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDPutRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsResourcesResourceScopesApiService.V1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDPut")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/resources/{resourceID}/scopes/{scopeID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resourceID"+"}", _neturl.PathEscape(parameterToString(r.resourceID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"scopeID"+"}", _neturl.PathEscape(parameterToString(r.scopeID, "")), -1)

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
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
