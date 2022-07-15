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


// IdentityPropagationProvisioningPropagationMappingsApiService IdentityPropagationProvisioningPropagationMappingsApi service
type IdentityPropagationProvisioningPropagationMappingsApiService service

type ApiV1EnvironmentsEnvironmentIDPropagationMappingMappingIDDeleteRequest struct {
	ctx context.Context
	ApiService *IdentityPropagationProvisioningPropagationMappingsApiService
	environmentID string
	mappingID string
	accept *string
}

func (r ApiV1EnvironmentsEnvironmentIDPropagationMappingMappingIDDeleteRequest) Accept(accept string) ApiV1EnvironmentsEnvironmentIDPropagationMappingMappingIDDeleteRequest {
	r.accept = &accept
	return r
}

func (r ApiV1EnvironmentsEnvironmentIDPropagationMappingMappingIDDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1EnvironmentsEnvironmentIDPropagationMappingMappingIDDeleteExecute(r)
}

/*
V1EnvironmentsEnvironmentIDPropagationMappingMappingIDDelete DELETE Mapping

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentID
 @param mappingID
 @return ApiV1EnvironmentsEnvironmentIDPropagationMappingMappingIDDeleteRequest
*/
func (a *IdentityPropagationProvisioningPropagationMappingsApiService) V1EnvironmentsEnvironmentIDPropagationMappingMappingIDDelete(ctx context.Context, environmentID string, mappingID string) ApiV1EnvironmentsEnvironmentIDPropagationMappingMappingIDDeleteRequest {
	return ApiV1EnvironmentsEnvironmentIDPropagationMappingMappingIDDeleteRequest{
		ApiService: a,
		ctx: ctx,
		environmentID: environmentID,
		mappingID: mappingID,
	}
}

// Execute executes the request
func (a *IdentityPropagationProvisioningPropagationMappingsApiService) V1EnvironmentsEnvironmentIDPropagationMappingMappingIDDeleteExecute(r ApiV1EnvironmentsEnvironmentIDPropagationMappingMappingIDDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityPropagationProvisioningPropagationMappingsApiService.V1EnvironmentsEnvironmentIDPropagationMappingMappingIDDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{environmentID}/propagation/mapping/{mappingID}"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterToString(r.environmentID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"mappingID"+"}", url.PathEscape(parameterToString(r.mappingID, "")), -1)

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
	if r.accept != nil {
		localVarHeaderParams["Accept"] = parameterToString(*r.accept, "")
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
		if localVarHTTPResponse.StatusCode == 403 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
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

type ApiV1EnvironmentsEnvironmentIDPropagationMappingsMappingIDGetRequest struct {
	ctx context.Context
	ApiService *IdentityPropagationProvisioningPropagationMappingsApiService
	environmentID string
	mappingID string
	accept *string
}

func (r ApiV1EnvironmentsEnvironmentIDPropagationMappingsMappingIDGetRequest) Accept(accept string) ApiV1EnvironmentsEnvironmentIDPropagationMappingsMappingIDGetRequest {
	r.accept = &accept
	return r
}

func (r ApiV1EnvironmentsEnvironmentIDPropagationMappingsMappingIDGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1EnvironmentsEnvironmentIDPropagationMappingsMappingIDGetExecute(r)
}

/*
V1EnvironmentsEnvironmentIDPropagationMappingsMappingIDGet READ One Mapping

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentID
 @param mappingID
 @return ApiV1EnvironmentsEnvironmentIDPropagationMappingsMappingIDGetRequest
*/
func (a *IdentityPropagationProvisioningPropagationMappingsApiService) V1EnvironmentsEnvironmentIDPropagationMappingsMappingIDGet(ctx context.Context, environmentID string, mappingID string) ApiV1EnvironmentsEnvironmentIDPropagationMappingsMappingIDGetRequest {
	return ApiV1EnvironmentsEnvironmentIDPropagationMappingsMappingIDGetRequest{
		ApiService: a,
		ctx: ctx,
		environmentID: environmentID,
		mappingID: mappingID,
	}
}

// Execute executes the request
func (a *IdentityPropagationProvisioningPropagationMappingsApiService) V1EnvironmentsEnvironmentIDPropagationMappingsMappingIDGetExecute(r ApiV1EnvironmentsEnvironmentIDPropagationMappingsMappingIDGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityPropagationProvisioningPropagationMappingsApiService.V1EnvironmentsEnvironmentIDPropagationMappingsMappingIDGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{environmentID}/propagation/mappings/{mappingID}"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterToString(r.environmentID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"mappingID"+"}", url.PathEscape(parameterToString(r.mappingID, "")), -1)

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
	if r.accept != nil {
		localVarHeaderParams["Accept"] = parameterToString(*r.accept, "")
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
		if localVarHTTPResponse.StatusCode == 403 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
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

type ApiV1EnvironmentsEnvironmentIDPropagationMappingsMappingIDPutRequest struct {
	ctx context.Context
	ApiService *IdentityPropagationProvisioningPropagationMappingsApiService
	environmentID string
	mappingID string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvironmentIDPropagationMappingsMappingIDPutRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvironmentIDPropagationMappingsMappingIDPutRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvironmentIDPropagationMappingsMappingIDPutRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1EnvironmentsEnvironmentIDPropagationMappingsMappingIDPutExecute(r)
}

/*
V1EnvironmentsEnvironmentIDPropagationMappingsMappingIDPut UPDATE Mapping

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentID
 @param mappingID
 @return ApiV1EnvironmentsEnvironmentIDPropagationMappingsMappingIDPutRequest
*/
func (a *IdentityPropagationProvisioningPropagationMappingsApiService) V1EnvironmentsEnvironmentIDPropagationMappingsMappingIDPut(ctx context.Context, environmentID string, mappingID string) ApiV1EnvironmentsEnvironmentIDPropagationMappingsMappingIDPutRequest {
	return ApiV1EnvironmentsEnvironmentIDPropagationMappingsMappingIDPutRequest{
		ApiService: a,
		ctx: ctx,
		environmentID: environmentID,
		mappingID: mappingID,
	}
}

// Execute executes the request
func (a *IdentityPropagationProvisioningPropagationMappingsApiService) V1EnvironmentsEnvironmentIDPropagationMappingsMappingIDPutExecute(r ApiV1EnvironmentsEnvironmentIDPropagationMappingsMappingIDPutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityPropagationProvisioningPropagationMappingsApiService.V1EnvironmentsEnvironmentIDPropagationMappingsMappingIDPut")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{environmentID}/propagation/mappings/{mappingID}"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterToString(r.environmentID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"mappingID"+"}", url.PathEscape(parameterToString(r.mappingID, "")), -1)

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
	localVarPostBody = r.body
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
		if localVarHTTPResponse.StatusCode == 403 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
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

type ApiV1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsGetRequest struct {
	ctx context.Context
	ApiService *IdentityPropagationProvisioningPropagationMappingsApiService
	environmentID string
	ruleID string
	accept *string
	contentType *string
}

func (r ApiV1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsGetRequest) Accept(accept string) ApiV1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsGetRequest {
	r.accept = &accept
	return r
}

func (r ApiV1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsGetRequest) ContentType(contentType string) ApiV1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsGetRequest {
	r.contentType = &contentType
	return r
}

func (r ApiV1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsGetExecute(r)
}

/*
V1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsGet READ One Rule  Mapping

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentID
 @param ruleID
 @return ApiV1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsGetRequest
*/
func (a *IdentityPropagationProvisioningPropagationMappingsApiService) V1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsGet(ctx context.Context, environmentID string, ruleID string) ApiV1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsGetRequest {
	return ApiV1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsGetRequest{
		ApiService: a,
		ctx: ctx,
		environmentID: environmentID,
		ruleID: ruleID,
	}
}

// Execute executes the request
func (a *IdentityPropagationProvisioningPropagationMappingsApiService) V1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsGetExecute(r ApiV1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityPropagationProvisioningPropagationMappingsApiService.V1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{environmentID}/propagation/rules/{ruleID}/mappings"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterToString(r.environmentID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"ruleID"+"}", url.PathEscape(parameterToString(r.ruleID, "")), -1)

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
	if r.accept != nil {
		localVarHeaderParams["Accept"] = parameterToString(*r.accept, "")
	}
	if r.contentType != nil {
		localVarHeaderParams["Content-Type"] = parameterToString(*r.contentType, "")
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
		if localVarHTTPResponse.StatusCode == 403 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
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

type ApiV1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsPostRequest struct {
	ctx context.Context
	ApiService *IdentityPropagationProvisioningPropagationMappingsApiService
	environmentID string
	ruleID string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsPostRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsPostRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsPostExecute(r)
}

/*
V1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsPost CREATE Rule Mapping

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentID
 @param ruleID
 @return ApiV1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsPostRequest
*/
func (a *IdentityPropagationProvisioningPropagationMappingsApiService) V1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsPost(ctx context.Context, environmentID string, ruleID string) ApiV1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsPostRequest {
	return ApiV1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsPostRequest{
		ApiService: a,
		ctx: ctx,
		environmentID: environmentID,
		ruleID: ruleID,
	}
}

// Execute executes the request
func (a *IdentityPropagationProvisioningPropagationMappingsApiService) V1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsPostExecute(r ApiV1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityPropagationProvisioningPropagationMappingsApiService.V1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{environmentID}/propagation/rules/{ruleID}/mappings"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterToString(r.environmentID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"ruleID"+"}", url.PathEscape(parameterToString(r.ruleID, "")), -1)

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
	localVarPostBody = r.body
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
		if localVarHTTPResponse.StatusCode == 403 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
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
