/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// ApplicationsApplicationAttributeMappingApiService ApplicationsApplicationAttributeMappingApi service
type ApplicationsApplicationAttributeMappingApiService service

type ApiCreateApplicationAttributeMappingRequest struct {
	ctx context.Context
	ApiService *ApplicationsApplicationAttributeMappingApiService
	environmentID string
	applicationID string
	applicationAttributeMapping *ApplicationAttributeMapping
}

func (r ApiCreateApplicationAttributeMappingRequest) ApplicationAttributeMapping(applicationAttributeMapping ApplicationAttributeMapping) ApiCreateApplicationAttributeMappingRequest {
	r.applicationAttributeMapping = &applicationAttributeMapping
	return r
}

func (r ApiCreateApplicationAttributeMappingRequest) Execute() (*ApplicationAttributeMapping, *http.Response, error) {
	return r.ApiService.CreateApplicationAttributeMappingExecute(r)
}

/*
CreateApplicationAttributeMapping CREATE Application Attribute Mapping

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentID
 @param applicationID
 @return ApiCreateApplicationAttributeMappingRequest
*/
func (a *ApplicationsApplicationAttributeMappingApiService) CreateApplicationAttributeMapping(ctx context.Context, environmentID string, applicationID string) ApiCreateApplicationAttributeMappingRequest {
	return ApiCreateApplicationAttributeMappingRequest{
		ApiService: a,
		ctx: ctx,
		environmentID: environmentID,
		applicationID: applicationID,
	}
}

// Execute executes the request
//  @return ApplicationAttributeMapping
func (a *ApplicationsApplicationAttributeMappingApiService) CreateApplicationAttributeMappingExecute(r ApiCreateApplicationAttributeMappingRequest) (*ApplicationAttributeMapping, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApplicationAttributeMapping
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationsApplicationAttributeMappingApiService.CreateApplicationAttributeMapping")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{environmentID}/applications/{applicationID}/attributes"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterToString(r.environmentID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"applicationID"+"}", url.PathEscape(parameterToString(r.applicationID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	// body params
	localVarPostBody = r.applicationAttributeMapping
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteApplicationAttributeMappingRequest struct {
	ctx context.Context
	ApiService *ApplicationsApplicationAttributeMappingApiService
	environmentID string
	applicationID string
	attrMappingID string
}

func (r ApiDeleteApplicationAttributeMappingRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteApplicationAttributeMappingExecute(r)
}

/*
DeleteApplicationAttributeMapping DELETE Application Attribute Mapping

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentID
 @param applicationID
 @param attrMappingID
 @return ApiDeleteApplicationAttributeMappingRequest
*/
func (a *ApplicationsApplicationAttributeMappingApiService) DeleteApplicationAttributeMapping(ctx context.Context, environmentID string, applicationID string, attrMappingID string) ApiDeleteApplicationAttributeMappingRequest {
	return ApiDeleteApplicationAttributeMappingRequest{
		ApiService: a,
		ctx: ctx,
		environmentID: environmentID,
		applicationID: applicationID,
		attrMappingID: attrMappingID,
	}
}

// Execute executes the request
func (a *ApplicationsApplicationAttributeMappingApiService) DeleteApplicationAttributeMappingExecute(r ApiDeleteApplicationAttributeMappingRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationsApplicationAttributeMappingApiService.DeleteApplicationAttributeMapping")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{environmentID}/applications/{applicationID}/attributes/{attrMappingID}"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterToString(r.environmentID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"applicationID"+"}", url.PathEscape(parameterToString(r.applicationID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"attrMappingID"+"}", url.PathEscape(parameterToString(r.attrMappingID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
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

type ApiReadAllApplicationAttributeMappingsRequest struct {
	ctx context.Context
	ApiService *ApplicationsApplicationAttributeMappingApiService
	environmentID string
	applicationID string
}

func (r ApiReadAllApplicationAttributeMappingsRequest) Execute() (*EntityArray, *http.Response, error) {
	return r.ApiService.ReadAllApplicationAttributeMappingsExecute(r)
}

/*
ReadAllApplicationAttributeMappings READ All Application Attribute Mappings

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentID
 @param applicationID
 @return ApiReadAllApplicationAttributeMappingsRequest
*/
func (a *ApplicationsApplicationAttributeMappingApiService) ReadAllApplicationAttributeMappings(ctx context.Context, environmentID string, applicationID string) ApiReadAllApplicationAttributeMappingsRequest {
	return ApiReadAllApplicationAttributeMappingsRequest{
		ApiService: a,
		ctx: ctx,
		environmentID: environmentID,
		applicationID: applicationID,
	}
}

// Execute executes the request
//  @return EntityArray
func (a *ApplicationsApplicationAttributeMappingApiService) ReadAllApplicationAttributeMappingsExecute(r ApiReadAllApplicationAttributeMappingsRequest) (*EntityArray, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EntityArray
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationsApplicationAttributeMappingApiService.ReadAllApplicationAttributeMappings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{environmentID}/applications/{applicationID}/attributes"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterToString(r.environmentID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"applicationID"+"}", url.PathEscape(parameterToString(r.applicationID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiReadOneApplicationAttributeMappingRequest struct {
	ctx context.Context
	ApiService *ApplicationsApplicationAttributeMappingApiService
	environmentID string
	applicationID string
	attrMappingID string
}

func (r ApiReadOneApplicationAttributeMappingRequest) Execute() (*ApplicationAttributeMapping, *http.Response, error) {
	return r.ApiService.ReadOneApplicationAttributeMappingExecute(r)
}

/*
ReadOneApplicationAttributeMapping READ One Application Attribute Mapping

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentID
 @param applicationID
 @param attrMappingID
 @return ApiReadOneApplicationAttributeMappingRequest
*/
func (a *ApplicationsApplicationAttributeMappingApiService) ReadOneApplicationAttributeMapping(ctx context.Context, environmentID string, applicationID string, attrMappingID string) ApiReadOneApplicationAttributeMappingRequest {
	return ApiReadOneApplicationAttributeMappingRequest{
		ApiService: a,
		ctx: ctx,
		environmentID: environmentID,
		applicationID: applicationID,
		attrMappingID: attrMappingID,
	}
}

// Execute executes the request
//  @return ApplicationAttributeMapping
func (a *ApplicationsApplicationAttributeMappingApiService) ReadOneApplicationAttributeMappingExecute(r ApiReadOneApplicationAttributeMappingRequest) (*ApplicationAttributeMapping, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApplicationAttributeMapping
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationsApplicationAttributeMappingApiService.ReadOneApplicationAttributeMapping")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{environmentID}/applications/{applicationID}/attributes/{attrMappingID}"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterToString(r.environmentID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"applicationID"+"}", url.PathEscape(parameterToString(r.applicationID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"attrMappingID"+"}", url.PathEscape(parameterToString(r.attrMappingID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateApplicationAttributeMappingRequest struct {
	ctx context.Context
	ApiService *ApplicationsApplicationAttributeMappingApiService
	environmentID string
	applicationID string
	attrMappingID string
	applicationAttributeMapping *ApplicationAttributeMapping
}

func (r ApiUpdateApplicationAttributeMappingRequest) ApplicationAttributeMapping(applicationAttributeMapping ApplicationAttributeMapping) ApiUpdateApplicationAttributeMappingRequest {
	r.applicationAttributeMapping = &applicationAttributeMapping
	return r
}

func (r ApiUpdateApplicationAttributeMappingRequest) Execute() (*ApplicationAttributeMapping, *http.Response, error) {
	return r.ApiService.UpdateApplicationAttributeMappingExecute(r)
}

/*
UpdateApplicationAttributeMapping UPDATE Application Attribute Mapping

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentID
 @param applicationID
 @param attrMappingID
 @return ApiUpdateApplicationAttributeMappingRequest
*/
func (a *ApplicationsApplicationAttributeMappingApiService) UpdateApplicationAttributeMapping(ctx context.Context, environmentID string, applicationID string, attrMappingID string) ApiUpdateApplicationAttributeMappingRequest {
	return ApiUpdateApplicationAttributeMappingRequest{
		ApiService: a,
		ctx: ctx,
		environmentID: environmentID,
		applicationID: applicationID,
		attrMappingID: attrMappingID,
	}
}

// Execute executes the request
//  @return ApplicationAttributeMapping
func (a *ApplicationsApplicationAttributeMappingApiService) UpdateApplicationAttributeMappingExecute(r ApiUpdateApplicationAttributeMappingRequest) (*ApplicationAttributeMapping, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApplicationAttributeMapping
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationsApplicationAttributeMappingApiService.UpdateApplicationAttributeMapping")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{environmentID}/applications/{applicationID}/attributes/{attrMappingID}"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterToString(r.environmentID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"applicationID"+"}", url.PathEscape(parameterToString(r.applicationID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"attrMappingID"+"}", url.PathEscape(parameterToString(r.attrMappingID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	// body params
	localVarPostBody = r.applicationAttributeMapping
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}