/*
PingOne Platform API - Management

A bare-bones collection for the PingOne API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pingone

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// ManagementAPIsSchemasApiService ManagementAPIsSchemasApi service
type ManagementAPIsSchemasApiService service

type ApiCreateAttributeRequest struct {
	ctx context.Context
	ApiService *ManagementAPIsSchemasApiService
	envID string
	schemaID string
	schemaAttribute *SchemaAttribute
}

func (r ApiCreateAttributeRequest) SchemaAttribute(schemaAttribute SchemaAttribute) ApiCreateAttributeRequest {
	r.schemaAttribute = &schemaAttribute
	return r
}

func (r ApiCreateAttributeRequest) Execute() (*SchemaAttribute, *http.Response, error) {
	return r.ApiService.CreateAttributeExecute(r)
}

/*
CreateAttribute CREATE Attribute

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param schemaID
 @return ApiCreateAttributeRequest
*/
func (a *ManagementAPIsSchemasApiService) CreateAttribute(ctx context.Context, envID string, schemaID string) ApiCreateAttributeRequest {
	return ApiCreateAttributeRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		schemaID: schemaID,
	}
}

// Execute executes the request
//  @return SchemaAttribute
func (a *ManagementAPIsSchemasApiService) CreateAttributeExecute(r ApiCreateAttributeRequest) (*SchemaAttribute, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SchemaAttribute
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsSchemasApiService.CreateAttribute")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/schemas/{schemaID}/attributes"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", url.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"schemaID"+"}", url.PathEscape(parameterToString(r.schemaID, "")), -1)

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
	localVarPostBody = r.schemaAttribute
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

type ApiDeleteAttributeRequest struct {
	ctx context.Context
	ApiService *ManagementAPIsSchemasApiService
	envID string
	schemaID string
	attributeID string
}

func (r ApiDeleteAttributeRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteAttributeExecute(r)
}

/*
DeleteAttribute DELETE Attribute

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param schemaID
 @param attributeID
 @return ApiDeleteAttributeRequest
*/
func (a *ManagementAPIsSchemasApiService) DeleteAttribute(ctx context.Context, envID string, schemaID string, attributeID string) ApiDeleteAttributeRequest {
	return ApiDeleteAttributeRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		schemaID: schemaID,
		attributeID: attributeID,
	}
}

// Execute executes the request
func (a *ManagementAPIsSchemasApiService) DeleteAttributeExecute(r ApiDeleteAttributeRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsSchemasApiService.DeleteAttribute")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/schemas/{schemaID}/attributes/{attributeID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", url.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"schemaID"+"}", url.PathEscape(parameterToString(r.schemaID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"attributeID"+"}", url.PathEscape(parameterToString(r.attributeID, "")), -1)

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

type ApiReadAllSchemaAttributesRequest struct {
	ctx context.Context
	ApiService *ManagementAPIsSchemasApiService
	envID string
	schemaID string
}

func (r ApiReadAllSchemaAttributesRequest) Execute() (*EntityArray, *http.Response, error) {
	return r.ApiService.ReadAllSchemaAttributesExecute(r)
}

/*
ReadAllSchemaAttributes READ All (Schema) Attributes

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param schemaID
 @return ApiReadAllSchemaAttributesRequest
*/
func (a *ManagementAPIsSchemasApiService) ReadAllSchemaAttributes(ctx context.Context, envID string, schemaID string) ApiReadAllSchemaAttributesRequest {
	return ApiReadAllSchemaAttributesRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		schemaID: schemaID,
	}
}

// Execute executes the request
//  @return EntityArray
func (a *ManagementAPIsSchemasApiService) ReadAllSchemaAttributesExecute(r ApiReadAllSchemaAttributesRequest) (*EntityArray, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EntityArray
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsSchemasApiService.ReadAllSchemaAttributes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/schemas/{schemaID}/attributes"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", url.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"schemaID"+"}", url.PathEscape(parameterToString(r.schemaID, "")), -1)

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

type ApiReadAllSchemasRequest struct {
	ctx context.Context
	ApiService *ManagementAPIsSchemasApiService
	envID string
}

func (r ApiReadAllSchemasRequest) Execute() (*EntityArray, *http.Response, error) {
	return r.ApiService.ReadAllSchemasExecute(r)
}

/*
ReadAllSchemas READ All Schemas

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @return ApiReadAllSchemasRequest
*/
func (a *ManagementAPIsSchemasApiService) ReadAllSchemas(ctx context.Context, envID string) ApiReadAllSchemasRequest {
	return ApiReadAllSchemasRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
	}
}

// Execute executes the request
//  @return EntityArray
func (a *ManagementAPIsSchemasApiService) ReadAllSchemasExecute(r ApiReadAllSchemasRequest) (*EntityArray, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EntityArray
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsSchemasApiService.ReadAllSchemas")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/schemas"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", url.PathEscape(parameterToString(r.envID, "")), -1)

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

type ApiReadOneAttributeRequest struct {
	ctx context.Context
	ApiService *ManagementAPIsSchemasApiService
	envID string
	schemaID string
	attributeID string
}

func (r ApiReadOneAttributeRequest) Execute() (*SchemaAttribute, *http.Response, error) {
	return r.ApiService.ReadOneAttributeExecute(r)
}

/*
ReadOneAttribute READ One Attribute

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param schemaID
 @param attributeID
 @return ApiReadOneAttributeRequest
*/
func (a *ManagementAPIsSchemasApiService) ReadOneAttribute(ctx context.Context, envID string, schemaID string, attributeID string) ApiReadOneAttributeRequest {
	return ApiReadOneAttributeRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		schemaID: schemaID,
		attributeID: attributeID,
	}
}

// Execute executes the request
//  @return SchemaAttribute
func (a *ManagementAPIsSchemasApiService) ReadOneAttributeExecute(r ApiReadOneAttributeRequest) (*SchemaAttribute, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SchemaAttribute
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsSchemasApiService.ReadOneAttribute")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/schemas/{schemaID}/attributes/{attributeID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", url.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"schemaID"+"}", url.PathEscape(parameterToString(r.schemaID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"attributeID"+"}", url.PathEscape(parameterToString(r.attributeID, "")), -1)

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

type ApiReadOneSchemaRequest struct {
	ctx context.Context
	ApiService *ManagementAPIsSchemasApiService
	envID string
	schemaID string
}

func (r ApiReadOneSchemaRequest) Execute() (*Schema, *http.Response, error) {
	return r.ApiService.ReadOneSchemaExecute(r)
}

/*
ReadOneSchema READ One Schema

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param schemaID
 @return ApiReadOneSchemaRequest
*/
func (a *ManagementAPIsSchemasApiService) ReadOneSchema(ctx context.Context, envID string, schemaID string) ApiReadOneSchemaRequest {
	return ApiReadOneSchemaRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		schemaID: schemaID,
	}
}

// Execute executes the request
//  @return Schema
func (a *ManagementAPIsSchemasApiService) ReadOneSchemaExecute(r ApiReadOneSchemaRequest) (*Schema, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Schema
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsSchemasApiService.ReadOneSchema")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/schemas/{schemaID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", url.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"schemaID"+"}", url.PathEscape(parameterToString(r.schemaID, "")), -1)

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

type ApiUpdateAttributePatchRequest struct {
	ctx context.Context
	ApiService *ManagementAPIsSchemasApiService
	envID string
	schemaID string
	attributeID string
	schemaAttribute *SchemaAttribute
}

func (r ApiUpdateAttributePatchRequest) SchemaAttribute(schemaAttribute SchemaAttribute) ApiUpdateAttributePatchRequest {
	r.schemaAttribute = &schemaAttribute
	return r
}

func (r ApiUpdateAttributePatchRequest) Execute() (*SchemaAttribute, *http.Response, error) {
	return r.ApiService.UpdateAttributePatchExecute(r)
}

/*
UpdateAttributePatch UPDATE Attribute (Patch)

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param schemaID
 @param attributeID
 @return ApiUpdateAttributePatchRequest
*/
func (a *ManagementAPIsSchemasApiService) UpdateAttributePatch(ctx context.Context, envID string, schemaID string, attributeID string) ApiUpdateAttributePatchRequest {
	return ApiUpdateAttributePatchRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		schemaID: schemaID,
		attributeID: attributeID,
	}
}

// Execute executes the request
//  @return SchemaAttribute
func (a *ManagementAPIsSchemasApiService) UpdateAttributePatchExecute(r ApiUpdateAttributePatchRequest) (*SchemaAttribute, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SchemaAttribute
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsSchemasApiService.UpdateAttributePatch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/schemas/{schemaID}/attributes/{attributeID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", url.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"schemaID"+"}", url.PathEscape(parameterToString(r.schemaID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"attributeID"+"}", url.PathEscape(parameterToString(r.attributeID, "")), -1)

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
	localVarPostBody = r.schemaAttribute
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

type ApiUpdateAttributePutRequest struct {
	ctx context.Context
	ApiService *ManagementAPIsSchemasApiService
	envID string
	schemaID string
	attributeID string
	schemaAttribute *SchemaAttribute
}

func (r ApiUpdateAttributePutRequest) SchemaAttribute(schemaAttribute SchemaAttribute) ApiUpdateAttributePutRequest {
	r.schemaAttribute = &schemaAttribute
	return r
}

func (r ApiUpdateAttributePutRequest) Execute() (*SchemaAttribute, *http.Response, error) {
	return r.ApiService.UpdateAttributePutExecute(r)
}

/*
UpdateAttributePut UPDATE Attribute (Put)

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param schemaID
 @param attributeID
 @return ApiUpdateAttributePutRequest
*/
func (a *ManagementAPIsSchemasApiService) UpdateAttributePut(ctx context.Context, envID string, schemaID string, attributeID string) ApiUpdateAttributePutRequest {
	return ApiUpdateAttributePutRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		schemaID: schemaID,
		attributeID: attributeID,
	}
}

// Execute executes the request
//  @return SchemaAttribute
func (a *ManagementAPIsSchemasApiService) UpdateAttributePutExecute(r ApiUpdateAttributePutRequest) (*SchemaAttribute, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SchemaAttribute
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsSchemasApiService.UpdateAttributePut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/schemas/{schemaID}/attributes/{attributeID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", url.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"schemaID"+"}", url.PathEscape(parameterToString(r.schemaID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"attributeID"+"}", url.PathEscape(parameterToString(r.attributeID, "")), -1)

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
	localVarPostBody = r.schemaAttribute
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
