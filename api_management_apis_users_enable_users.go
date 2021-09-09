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

// ManagementAPIsUsersEnableUsersApiService ManagementAPIsUsersEnableUsersApi service
type ManagementAPIsUsersEnableUsersApiService service

type ApiV1EnvironmentsEnvIDUsersUserIDEnabledGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsUsersEnableUsersApiService
	envID string
	userID string
}


func (r ApiV1EnvironmentsEnvIDUsersUserIDEnabledGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDUsersUserIDEnabledGetExecute(r)
}

/*
V1EnvironmentsEnvIDUsersUserIDEnabledGet READ User Enabled

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param userID
 @return ApiV1EnvironmentsEnvIDUsersUserIDEnabledGetRequest
*/
func (a *ManagementAPIsUsersEnableUsersApiService) V1EnvironmentsEnvIDUsersUserIDEnabledGet(ctx _context.Context, envID string, userID string) ApiV1EnvironmentsEnvIDUsersUserIDEnabledGetRequest {
	return ApiV1EnvironmentsEnvIDUsersUserIDEnabledGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		userID: userID,
	}
}

// Execute executes the request
func (a *ManagementAPIsUsersEnableUsersApiService) V1EnvironmentsEnvIDUsersUserIDEnabledGetExecute(r ApiV1EnvironmentsEnvIDUsersUserIDEnabledGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsUsersEnableUsersApiService.V1EnvironmentsEnvIDUsersUserIDEnabledGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/users/{userID}/enabled"
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

type ApiV1EnvironmentsEnvIDUsersUserIDEnabledPutRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsUsersEnableUsersApiService
	envID string
	userID string
	contentType *string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDEnabledPutRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDUsersUserIDEnabledPutRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDUsersUserIDEnabledPutRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDUsersUserIDEnabledPutRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDEnabledPutRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDUsersUserIDEnabledPutExecute(r)
}

/*
V1EnvironmentsEnvIDUsersUserIDEnabledPut UPDATE User Enabled

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param userID
 @return ApiV1EnvironmentsEnvIDUsersUserIDEnabledPutRequest
*/
func (a *ManagementAPIsUsersEnableUsersApiService) V1EnvironmentsEnvIDUsersUserIDEnabledPut(ctx _context.Context, envID string, userID string) ApiV1EnvironmentsEnvIDUsersUserIDEnabledPutRequest {
	return ApiV1EnvironmentsEnvIDUsersUserIDEnabledPutRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		userID: userID,
	}
}

// Execute executes the request
func (a *ManagementAPIsUsersEnableUsersApiService) V1EnvironmentsEnvIDUsersUserIDEnabledPutExecute(r ApiV1EnvironmentsEnvIDUsersUserIDEnabledPutRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsUsersEnableUsersApiService.V1EnvironmentsEnvIDUsersUserIDEnabledPut")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/users/{userID}/enabled"
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
