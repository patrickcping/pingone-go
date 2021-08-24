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

// ManagementAPIsDeviceAuthenticationPolicyApiService ManagementAPIsDeviceAuthenticationPolicyApi service
type ManagementAPIsDeviceAuthenticationPolicyApiService service

type ApiV1EnvironmentsEnvIDDeviceAuthenticationPolicyDeviceAuthPolicyIDPutRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsDeviceAuthenticationPolicyApiService
	envID string
	deviceAuthPolicyID string
	contentType *string
	authorization *string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDDeviceAuthenticationPolicyDeviceAuthPolicyIDPutRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDDeviceAuthenticationPolicyDeviceAuthPolicyIDPutRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDDeviceAuthenticationPolicyDeviceAuthPolicyIDPutRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDDeviceAuthenticationPolicyDeviceAuthPolicyIDPutRequest {
	r.authorization = &authorization
	return r
}
func (r ApiV1EnvironmentsEnvIDDeviceAuthenticationPolicyDeviceAuthPolicyIDPutRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDDeviceAuthenticationPolicyDeviceAuthPolicyIDPutRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDDeviceAuthenticationPolicyDeviceAuthPolicyIDPutRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDDeviceAuthenticationPolicyDeviceAuthPolicyIDPutExecute(r)
}

/*
 * V1EnvironmentsEnvIDDeviceAuthenticationPolicyDeviceAuthPolicyIDPut UPDATE Device Authentication Policy
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @param deviceAuthPolicyID
 * @return ApiV1EnvironmentsEnvIDDeviceAuthenticationPolicyDeviceAuthPolicyIDPutRequest
 */
func (a *ManagementAPIsDeviceAuthenticationPolicyApiService) V1EnvironmentsEnvIDDeviceAuthenticationPolicyDeviceAuthPolicyIDPut(ctx _context.Context, envID string, deviceAuthPolicyID string) ApiV1EnvironmentsEnvIDDeviceAuthenticationPolicyDeviceAuthPolicyIDPutRequest {
	return ApiV1EnvironmentsEnvIDDeviceAuthenticationPolicyDeviceAuthPolicyIDPutRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		deviceAuthPolicyID: deviceAuthPolicyID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsDeviceAuthenticationPolicyApiService) V1EnvironmentsEnvIDDeviceAuthenticationPolicyDeviceAuthPolicyIDPutExecute(r ApiV1EnvironmentsEnvIDDeviceAuthenticationPolicyDeviceAuthPolicyIDPutRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsDeviceAuthenticationPolicyApiService.V1EnvironmentsEnvIDDeviceAuthenticationPolicyDeviceAuthPolicyIDPut")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/deviceAuthenticationPolicy/{deviceAuthPolicyID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"deviceAuthPolicyID"+"}", _neturl.PathEscape(parameterToString(r.deviceAuthPolicyID, "")), -1)

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

type ApiV1EnvironmentsEnvIDDeviceAuthenticationPolicyGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsDeviceAuthenticationPolicyApiService
	envID string
	authorization *string
}

func (r ApiV1EnvironmentsEnvIDDeviceAuthenticationPolicyGetRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDDeviceAuthenticationPolicyGetRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1EnvironmentsEnvIDDeviceAuthenticationPolicyGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDDeviceAuthenticationPolicyGetExecute(r)
}

/*
 * V1EnvironmentsEnvIDDeviceAuthenticationPolicyGet READ Device Authentication Policy
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @return ApiV1EnvironmentsEnvIDDeviceAuthenticationPolicyGetRequest
 */
func (a *ManagementAPIsDeviceAuthenticationPolicyApiService) V1EnvironmentsEnvIDDeviceAuthenticationPolicyGet(ctx _context.Context, envID string) ApiV1EnvironmentsEnvIDDeviceAuthenticationPolicyGetRequest {
	return ApiV1EnvironmentsEnvIDDeviceAuthenticationPolicyGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsDeviceAuthenticationPolicyApiService) V1EnvironmentsEnvIDDeviceAuthenticationPolicyGetExecute(r ApiV1EnvironmentsEnvIDDeviceAuthenticationPolicyGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsDeviceAuthenticationPolicyApiService.V1EnvironmentsEnvIDDeviceAuthenticationPolicyGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/deviceAuthenticationPolicy"
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
