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

// ManagementAPIsIdentityPropagationProvisioningPropagationPlansApiService ManagementAPIsIdentityPropagationProvisioningPropagationPlansApi service
type ManagementAPIsIdentityPropagationProvisioningPropagationPlansApiService service

type ApiV1EnvironmentsEnvIDPropagationPlansGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsIdentityPropagationProvisioningPropagationPlansApiService
	envID string
	accept *string
}

func (r ApiV1EnvironmentsEnvIDPropagationPlansGetRequest) Accept(accept string) ApiV1EnvironmentsEnvIDPropagationPlansGetRequest {
	r.accept = &accept
	return r
}

func (r ApiV1EnvironmentsEnvIDPropagationPlansGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDPropagationPlansGetExecute(r)
}

/*
V1EnvironmentsEnvIDPropagationPlansGet READ All Plans

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @return ApiV1EnvironmentsEnvIDPropagationPlansGetRequest
*/
func (a *ManagementAPIsIdentityPropagationProvisioningPropagationPlansApiService) V1EnvironmentsEnvIDPropagationPlansGet(ctx _context.Context, envID string) ApiV1EnvironmentsEnvIDPropagationPlansGetRequest {
	return ApiV1EnvironmentsEnvIDPropagationPlansGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
	}
}

// Execute executes the request
func (a *ManagementAPIsIdentityPropagationProvisioningPropagationPlansApiService) V1EnvironmentsEnvIDPropagationPlansGetExecute(r ApiV1EnvironmentsEnvIDPropagationPlansGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsIdentityPropagationProvisioningPropagationPlansApiService.V1EnvironmentsEnvIDPropagationPlansGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/propagation/plans"
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

type ApiV1EnvironmentsEnvIDPropagationPlansPlanIDDeleteRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsIdentityPropagationProvisioningPropagationPlansApiService
	envID string
	planID string
	accept *string
}

func (r ApiV1EnvironmentsEnvIDPropagationPlansPlanIDDeleteRequest) Accept(accept string) ApiV1EnvironmentsEnvIDPropagationPlansPlanIDDeleteRequest {
	r.accept = &accept
	return r
}

func (r ApiV1EnvironmentsEnvIDPropagationPlansPlanIDDeleteRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDPropagationPlansPlanIDDeleteExecute(r)
}

/*
V1EnvironmentsEnvIDPropagationPlansPlanIDDelete DELETE Plan

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param planID
 @return ApiV1EnvironmentsEnvIDPropagationPlansPlanIDDeleteRequest
*/
func (a *ManagementAPIsIdentityPropagationProvisioningPropagationPlansApiService) V1EnvironmentsEnvIDPropagationPlansPlanIDDelete(ctx _context.Context, envID string, planID string) ApiV1EnvironmentsEnvIDPropagationPlansPlanIDDeleteRequest {
	return ApiV1EnvironmentsEnvIDPropagationPlansPlanIDDeleteRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		planID: planID,
	}
}

// Execute executes the request
func (a *ManagementAPIsIdentityPropagationProvisioningPropagationPlansApiService) V1EnvironmentsEnvIDPropagationPlansPlanIDDeleteExecute(r ApiV1EnvironmentsEnvIDPropagationPlansPlanIDDeleteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsIdentityPropagationProvisioningPropagationPlansApiService.V1EnvironmentsEnvIDPropagationPlansPlanIDDelete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/propagation/plans/{planID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"planID"+"}", _neturl.PathEscape(parameterToString(r.planID, "")), -1)

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

type ApiV1EnvironmentsEnvIDPropagationPlansPlanIDGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsIdentityPropagationProvisioningPropagationPlansApiService
	envID string
	planID string
	accept *string
}

func (r ApiV1EnvironmentsEnvIDPropagationPlansPlanIDGetRequest) Accept(accept string) ApiV1EnvironmentsEnvIDPropagationPlansPlanIDGetRequest {
	r.accept = &accept
	return r
}

func (r ApiV1EnvironmentsEnvIDPropagationPlansPlanIDGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDPropagationPlansPlanIDGetExecute(r)
}

/*
V1EnvironmentsEnvIDPropagationPlansPlanIDGet READ One Plan

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param planID
 @return ApiV1EnvironmentsEnvIDPropagationPlansPlanIDGetRequest
*/
func (a *ManagementAPIsIdentityPropagationProvisioningPropagationPlansApiService) V1EnvironmentsEnvIDPropagationPlansPlanIDGet(ctx _context.Context, envID string, planID string) ApiV1EnvironmentsEnvIDPropagationPlansPlanIDGetRequest {
	return ApiV1EnvironmentsEnvIDPropagationPlansPlanIDGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		planID: planID,
	}
}

// Execute executes the request
func (a *ManagementAPIsIdentityPropagationProvisioningPropagationPlansApiService) V1EnvironmentsEnvIDPropagationPlansPlanIDGetExecute(r ApiV1EnvironmentsEnvIDPropagationPlansPlanIDGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsIdentityPropagationProvisioningPropagationPlansApiService.V1EnvironmentsEnvIDPropagationPlansPlanIDGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/propagation/plans/{planID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"planID"+"}", _neturl.PathEscape(parameterToString(r.planID, "")), -1)

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

type ApiV1EnvironmentsEnvIDPropagationPlansPlanIDPutRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsIdentityPropagationProvisioningPropagationPlansApiService
	envID string
	planID string
	contentType *string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDPropagationPlansPlanIDPutRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDPropagationPlansPlanIDPutRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDPropagationPlansPlanIDPutRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDPropagationPlansPlanIDPutRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDPropagationPlansPlanIDPutRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDPropagationPlansPlanIDPutExecute(r)
}

/*
V1EnvironmentsEnvIDPropagationPlansPlanIDPut UPDATE Plan

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param planID
 @return ApiV1EnvironmentsEnvIDPropagationPlansPlanIDPutRequest
*/
func (a *ManagementAPIsIdentityPropagationProvisioningPropagationPlansApiService) V1EnvironmentsEnvIDPropagationPlansPlanIDPut(ctx _context.Context, envID string, planID string) ApiV1EnvironmentsEnvIDPropagationPlansPlanIDPutRequest {
	return ApiV1EnvironmentsEnvIDPropagationPlansPlanIDPutRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		planID: planID,
	}
}

// Execute executes the request
func (a *ManagementAPIsIdentityPropagationProvisioningPropagationPlansApiService) V1EnvironmentsEnvIDPropagationPlansPlanIDPutExecute(r ApiV1EnvironmentsEnvIDPropagationPlansPlanIDPutRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsIdentityPropagationProvisioningPropagationPlansApiService.V1EnvironmentsEnvIDPropagationPlansPlanIDPut")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/propagation/plans/{planID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"planID"+"}", _neturl.PathEscape(parameterToString(r.planID, "")), -1)

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

type ApiV1EnvironmentsEnvIDPropagationPlansPostRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsIdentityPropagationProvisioningPropagationPlansApiService
	envID string
	contentType *string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDPropagationPlansPostRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDPropagationPlansPostRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDPropagationPlansPostRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDPropagationPlansPostRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDPropagationPlansPostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDPropagationPlansPostExecute(r)
}

/*
V1EnvironmentsEnvIDPropagationPlansPost CREATE Plan

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @return ApiV1EnvironmentsEnvIDPropagationPlansPostRequest
*/
func (a *ManagementAPIsIdentityPropagationProvisioningPropagationPlansApiService) V1EnvironmentsEnvIDPropagationPlansPost(ctx _context.Context, envID string) ApiV1EnvironmentsEnvIDPropagationPlansPostRequest {
	return ApiV1EnvironmentsEnvIDPropagationPlansPostRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
	}
}

// Execute executes the request
func (a *ManagementAPIsIdentityPropagationProvisioningPropagationPlansApiService) V1EnvironmentsEnvIDPropagationPlansPostExecute(r ApiV1EnvironmentsEnvIDPropagationPlansPostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsIdentityPropagationProvisioningPropagationPlansApiService.V1EnvironmentsEnvIDPropagationPlansPost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/propagation/plans"
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
