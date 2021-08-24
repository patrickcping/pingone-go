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

// ManagementAPIsRiskManagementRiskPoliciesApiService ManagementAPIsRiskManagementRiskPoliciesApi service
type ManagementAPIsRiskManagementRiskPoliciesApiService service

type ApiV1EnvironmentsEnvIDRiskPolicySetsGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsRiskManagementRiskPoliciesApiService
	envID string
	authorization *string
}

func (r ApiV1EnvironmentsEnvIDRiskPolicySetsGetRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDRiskPolicySetsGetRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1EnvironmentsEnvIDRiskPolicySetsGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDRiskPolicySetsGetExecute(r)
}

/*
 * V1EnvironmentsEnvIDRiskPolicySetsGet READ Risk Policy Sets
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @return ApiV1EnvironmentsEnvIDRiskPolicySetsGetRequest
 */
func (a *ManagementAPIsRiskManagementRiskPoliciesApiService) V1EnvironmentsEnvIDRiskPolicySetsGet(ctx _context.Context, envID string) ApiV1EnvironmentsEnvIDRiskPolicySetsGetRequest {
	return ApiV1EnvironmentsEnvIDRiskPolicySetsGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsRiskManagementRiskPoliciesApiService) V1EnvironmentsEnvIDRiskPolicySetsGetExecute(r ApiV1EnvironmentsEnvIDRiskPolicySetsGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsRiskManagementRiskPoliciesApiService.V1EnvironmentsEnvIDRiskPolicySetsGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/riskPolicySets"
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

type ApiV1EnvironmentsEnvIDRiskPolicySetsPostRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsRiskManagementRiskPoliciesApiService
	envID string
	authorization *string
	contentType *string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDRiskPolicySetsPostRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDRiskPolicySetsPostRequest {
	r.authorization = &authorization
	return r
}
func (r ApiV1EnvironmentsEnvIDRiskPolicySetsPostRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDRiskPolicySetsPostRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDRiskPolicySetsPostRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDRiskPolicySetsPostRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDRiskPolicySetsPostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDRiskPolicySetsPostExecute(r)
}

/*
 * V1EnvironmentsEnvIDRiskPolicySetsPost CREATE Risk Policy Set
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @return ApiV1EnvironmentsEnvIDRiskPolicySetsPostRequest
 */
func (a *ManagementAPIsRiskManagementRiskPoliciesApiService) V1EnvironmentsEnvIDRiskPolicySetsPost(ctx _context.Context, envID string) ApiV1EnvironmentsEnvIDRiskPolicySetsPostRequest {
	return ApiV1EnvironmentsEnvIDRiskPolicySetsPostRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsRiskManagementRiskPoliciesApiService) V1EnvironmentsEnvIDRiskPolicySetsPostExecute(r ApiV1EnvironmentsEnvIDRiskPolicySetsPostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsRiskManagementRiskPoliciesApiService.V1EnvironmentsEnvIDRiskPolicySetsPost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/riskPolicySets"
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

type ApiV1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDDeleteRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsRiskManagementRiskPoliciesApiService
	envID string
	riskPolicySetID string
	authorization *string
}

func (r ApiV1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDDeleteRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDDeleteRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDDeleteRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDDeleteExecute(r)
}

/*
 * V1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDDelete DELETE Risk Policy Set 
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @param riskPolicySetID
 * @return ApiV1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDDeleteRequest
 */
func (a *ManagementAPIsRiskManagementRiskPoliciesApiService) V1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDDelete(ctx _context.Context, envID string, riskPolicySetID string) ApiV1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDDeleteRequest {
	return ApiV1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDDeleteRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		riskPolicySetID: riskPolicySetID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsRiskManagementRiskPoliciesApiService) V1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDDeleteExecute(r ApiV1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDDeleteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsRiskManagementRiskPoliciesApiService.V1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDDelete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/riskPolicySets/{riskPolicySetID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"riskPolicySetID"+"}", _neturl.PathEscape(parameterToString(r.riskPolicySetID, "")), -1)

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

type ApiV1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsRiskManagementRiskPoliciesApiService
	envID string
	riskPolicySetID string
	authorization *string
}

func (r ApiV1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDGetRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDGetRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDGetExecute(r)
}

/*
 * V1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDGet READ One Risk Policy Set
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @param riskPolicySetID
 * @return ApiV1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDGetRequest
 */
func (a *ManagementAPIsRiskManagementRiskPoliciesApiService) V1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDGet(ctx _context.Context, envID string, riskPolicySetID string) ApiV1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDGetRequest {
	return ApiV1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		riskPolicySetID: riskPolicySetID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsRiskManagementRiskPoliciesApiService) V1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDGetExecute(r ApiV1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsRiskManagementRiskPoliciesApiService.V1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/riskPolicySets/{riskPolicySetID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"riskPolicySetID"+"}", _neturl.PathEscape(parameterToString(r.riskPolicySetID, "")), -1)

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

type ApiV1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDPutRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsRiskManagementRiskPoliciesApiService
	envID string
	riskPolicySetID string
	authorization *string
	contentType *string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDPutRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDPutRequest {
	r.authorization = &authorization
	return r
}
func (r ApiV1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDPutRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDPutRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDPutRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDPutRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDPutRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDPutExecute(r)
}

/*
 * V1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDPut UPDATE Risk Policy Set
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @param riskPolicySetID
 * @return ApiV1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDPutRequest
 */
func (a *ManagementAPIsRiskManagementRiskPoliciesApiService) V1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDPut(ctx _context.Context, envID string, riskPolicySetID string) ApiV1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDPutRequest {
	return ApiV1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDPutRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		riskPolicySetID: riskPolicySetID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsRiskManagementRiskPoliciesApiService) V1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDPutExecute(r ApiV1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDPutRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsRiskManagementRiskPoliciesApiService.V1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDPut")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/riskPolicySets/{riskPolicySetID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"riskPolicySetID"+"}", _neturl.PathEscape(parameterToString(r.riskPolicySetID, "")), -1)

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
