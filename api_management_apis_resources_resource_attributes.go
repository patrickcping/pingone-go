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

// ManagementAPIsResourcesResourceAttributesApiService ManagementAPIsResourcesResourceAttributesApi service
type ManagementAPIsResourcesResourceAttributesApiService service

type ApiV1EnvironmentsEnvIDResourcesResourceIDAttributesGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsResourcesResourceAttributesApiService
	envID string
	resourceID string
}


func (r ApiV1EnvironmentsEnvIDResourcesResourceIDAttributesGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDResourcesResourceIDAttributesGetExecute(r)
}

/*
V1EnvironmentsEnvIDResourcesResourceIDAttributesGet READ All Resource Attributes

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param resourceID
 @return ApiV1EnvironmentsEnvIDResourcesResourceIDAttributesGetRequest
*/
func (a *ManagementAPIsResourcesResourceAttributesApiService) V1EnvironmentsEnvIDResourcesResourceIDAttributesGet(ctx _context.Context, envID string, resourceID string) ApiV1EnvironmentsEnvIDResourcesResourceIDAttributesGetRequest {
	return ApiV1EnvironmentsEnvIDResourcesResourceIDAttributesGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		resourceID: resourceID,
	}
}

// Execute executes the request
func (a *ManagementAPIsResourcesResourceAttributesApiService) V1EnvironmentsEnvIDResourcesResourceIDAttributesGetExecute(r ApiV1EnvironmentsEnvIDResourcesResourceIDAttributesGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsResourcesResourceAttributesApiService.V1EnvironmentsEnvIDResourcesResourceIDAttributesGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/resources/{resourceID}/attributes"
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
			var v P1Error
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

type ApiV1EnvironmentsEnvIDResourcesResourceIDAttributesPostRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsResourcesResourceAttributesApiService
	envID string
	resourceID string
	contentType *string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDResourcesResourceIDAttributesPostRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDResourcesResourceIDAttributesPostRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDResourcesResourceIDAttributesPostRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDResourcesResourceIDAttributesPostRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDResourcesResourceIDAttributesPostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDResourcesResourceIDAttributesPostExecute(r)
}

/*
V1EnvironmentsEnvIDResourcesResourceIDAttributesPost CREATE Resource Attribute

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param resourceID
 @return ApiV1EnvironmentsEnvIDResourcesResourceIDAttributesPostRequest
*/
func (a *ManagementAPIsResourcesResourceAttributesApiService) V1EnvironmentsEnvIDResourcesResourceIDAttributesPost(ctx _context.Context, envID string, resourceID string) ApiV1EnvironmentsEnvIDResourcesResourceIDAttributesPostRequest {
	return ApiV1EnvironmentsEnvIDResourcesResourceIDAttributesPostRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		resourceID: resourceID,
	}
}

// Execute executes the request
func (a *ManagementAPIsResourcesResourceAttributesApiService) V1EnvironmentsEnvIDResourcesResourceIDAttributesPostExecute(r ApiV1EnvironmentsEnvIDResourcesResourceIDAttributesPostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsResourcesResourceAttributesApiService.V1EnvironmentsEnvIDResourcesResourceIDAttributesPost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/resources/{resourceID}/attributes"
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
			var v P1Error
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

type ApiV1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDDeleteRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsResourcesResourceAttributesApiService
	envID string
	resourceID string
	resourceAttrID string
}


func (r ApiV1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDDeleteRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDDeleteExecute(r)
}

/*
V1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDDelete DELETE Resource Attribute

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param resourceID
 @param resourceAttrID
 @return ApiV1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDDeleteRequest
*/
func (a *ManagementAPIsResourcesResourceAttributesApiService) V1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDDelete(ctx _context.Context, envID string, resourceID string, resourceAttrID string) ApiV1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDDeleteRequest {
	return ApiV1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDDeleteRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		resourceID: resourceID,
		resourceAttrID: resourceAttrID,
	}
}

// Execute executes the request
func (a *ManagementAPIsResourcesResourceAttributesApiService) V1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDDeleteExecute(r ApiV1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDDeleteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsResourcesResourceAttributesApiService.V1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDDelete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/resources/{resourceID}/attributes/{resourceAttrID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resourceID"+"}", _neturl.PathEscape(parameterToString(r.resourceID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resourceAttrID"+"}", _neturl.PathEscape(parameterToString(r.resourceAttrID, "")), -1)

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
			var v P1Error
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

type ApiV1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsResourcesResourceAttributesApiService
	envID string
	resourceID string
	resourceAttrID string
}


func (r ApiV1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDGetExecute(r)
}

/*
V1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDGet READ One Resource Attribute

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param resourceID
 @param resourceAttrID
 @return ApiV1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDGetRequest
*/
func (a *ManagementAPIsResourcesResourceAttributesApiService) V1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDGet(ctx _context.Context, envID string, resourceID string, resourceAttrID string) ApiV1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDGetRequest {
	return ApiV1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		resourceID: resourceID,
		resourceAttrID: resourceAttrID,
	}
}

// Execute executes the request
func (a *ManagementAPIsResourcesResourceAttributesApiService) V1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDGetExecute(r ApiV1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsResourcesResourceAttributesApiService.V1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/resources/{resourceID}/attributes/{resourceAttrID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resourceID"+"}", _neturl.PathEscape(parameterToString(r.resourceID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resourceAttrID"+"}", _neturl.PathEscape(parameterToString(r.resourceAttrID, "")), -1)

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
			var v P1Error
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

type ApiV1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDPutRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsResourcesResourceAttributesApiService
	envID string
	resourceID string
	resourceAttrID string
	contentType *string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDPutRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDPutRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDPutRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDPutRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDPutRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDPutExecute(r)
}

/*
V1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDPut UPDATE Resource Attribute

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param resourceID
 @param resourceAttrID
 @return ApiV1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDPutRequest
*/
func (a *ManagementAPIsResourcesResourceAttributesApiService) V1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDPut(ctx _context.Context, envID string, resourceID string, resourceAttrID string) ApiV1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDPutRequest {
	return ApiV1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDPutRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		resourceID: resourceID,
		resourceAttrID: resourceAttrID,
	}
}

// Execute executes the request
func (a *ManagementAPIsResourcesResourceAttributesApiService) V1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDPutExecute(r ApiV1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDPutRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsResourcesResourceAttributesApiService.V1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDPut")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/resources/{resourceID}/attributes/{resourceAttrID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resourceID"+"}", _neturl.PathEscape(parameterToString(r.resourceID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resourceAttrID"+"}", _neturl.PathEscape(parameterToString(r.resourceAttrID, "")), -1)

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
			var v P1Error
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
