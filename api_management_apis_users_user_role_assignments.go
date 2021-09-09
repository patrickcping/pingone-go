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

// ManagementAPIsUsersUserRoleAssignmentsApiService ManagementAPIsUsersUserRoleAssignmentsApi service
type ManagementAPIsUsersUserRoleAssignmentsApiService service

type ApiV1EnvironmentsEnvIDUsersUserIDRoleAssignmentsGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsUsersUserRoleAssignmentsApiService
	envID string
	userID string
}


func (r ApiV1EnvironmentsEnvIDUsersUserIDRoleAssignmentsGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsGetExecute(r)
}

/*
V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsGet READ Role Assignments

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param userID
 @return ApiV1EnvironmentsEnvIDUsersUserIDRoleAssignmentsGetRequest
*/
func (a *ManagementAPIsUsersUserRoleAssignmentsApiService) V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsGet(ctx _context.Context, envID string, userID string) ApiV1EnvironmentsEnvIDUsersUserIDRoleAssignmentsGetRequest {
	return ApiV1EnvironmentsEnvIDUsersUserIDRoleAssignmentsGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		userID: userID,
	}
}

// Execute executes the request
func (a *ManagementAPIsUsersUserRoleAssignmentsApiService) V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsGetExecute(r ApiV1EnvironmentsEnvIDUsersUserIDRoleAssignmentsGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsUsersUserRoleAssignmentsApiService.V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/users/{userID}/roleAssignments"
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

type ApiV1EnvironmentsEnvIDUsersUserIDRoleAssignmentsPostRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsUsersUserRoleAssignmentsApiService
	envID string
	userID string
	contentType *string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDRoleAssignmentsPostRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDUsersUserIDRoleAssignmentsPostRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDUsersUserIDRoleAssignmentsPostRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDUsersUserIDRoleAssignmentsPostRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDRoleAssignmentsPostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsPostExecute(r)
}

/*
V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsPost CREATE User Role Assignment

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param userID
 @return ApiV1EnvironmentsEnvIDUsersUserIDRoleAssignmentsPostRequest
*/
func (a *ManagementAPIsUsersUserRoleAssignmentsApiService) V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsPost(ctx _context.Context, envID string, userID string) ApiV1EnvironmentsEnvIDUsersUserIDRoleAssignmentsPostRequest {
	return ApiV1EnvironmentsEnvIDUsersUserIDRoleAssignmentsPostRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		userID: userID,
	}
}

// Execute executes the request
func (a *ManagementAPIsUsersUserRoleAssignmentsApiService) V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsPostExecute(r ApiV1EnvironmentsEnvIDUsersUserIDRoleAssignmentsPostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsUsersUserRoleAssignmentsApiService.V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsPost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/users/{userID}/roleAssignments"
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

type ApiV1EnvironmentsEnvIDUsersUserIDRoleAssignmentsRoleAssignmentIDDeleteRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsUsersUserRoleAssignmentsApiService
	envID string
	userID string
	roleAssignmentID string
}


func (r ApiV1EnvironmentsEnvIDUsersUserIDRoleAssignmentsRoleAssignmentIDDeleteRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsRoleAssignmentIDDeleteExecute(r)
}

/*
V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsRoleAssignmentIDDelete DELETE User's Role Assignment

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param userID
 @param roleAssignmentID
 @return ApiV1EnvironmentsEnvIDUsersUserIDRoleAssignmentsRoleAssignmentIDDeleteRequest
*/
func (a *ManagementAPIsUsersUserRoleAssignmentsApiService) V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsRoleAssignmentIDDelete(ctx _context.Context, envID string, userID string, roleAssignmentID string) ApiV1EnvironmentsEnvIDUsersUserIDRoleAssignmentsRoleAssignmentIDDeleteRequest {
	return ApiV1EnvironmentsEnvIDUsersUserIDRoleAssignmentsRoleAssignmentIDDeleteRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		userID: userID,
		roleAssignmentID: roleAssignmentID,
	}
}

// Execute executes the request
func (a *ManagementAPIsUsersUserRoleAssignmentsApiService) V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsRoleAssignmentIDDeleteExecute(r ApiV1EnvironmentsEnvIDUsersUserIDRoleAssignmentsRoleAssignmentIDDeleteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsUsersUserRoleAssignmentsApiService.V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsRoleAssignmentIDDelete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/users/{userID}/roleAssignments/{roleAssignmentID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userID"+"}", _neturl.PathEscape(parameterToString(r.userID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleAssignmentID"+"}", _neturl.PathEscape(parameterToString(r.roleAssignmentID, "")), -1)

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

type ApiV1EnvironmentsEnvIDUsersUserIDRoleAssignmentsRoleAssignmentIDGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsUsersUserRoleAssignmentsApiService
	envID string
	userID string
	roleAssignmentID string
}


func (r ApiV1EnvironmentsEnvIDUsersUserIDRoleAssignmentsRoleAssignmentIDGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsRoleAssignmentIDGetExecute(r)
}

/*
V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsRoleAssignmentIDGet READ One Role Assignment

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param userID
 @param roleAssignmentID
 @return ApiV1EnvironmentsEnvIDUsersUserIDRoleAssignmentsRoleAssignmentIDGetRequest
*/
func (a *ManagementAPIsUsersUserRoleAssignmentsApiService) V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsRoleAssignmentIDGet(ctx _context.Context, envID string, userID string, roleAssignmentID string) ApiV1EnvironmentsEnvIDUsersUserIDRoleAssignmentsRoleAssignmentIDGetRequest {
	return ApiV1EnvironmentsEnvIDUsersUserIDRoleAssignmentsRoleAssignmentIDGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		userID: userID,
		roleAssignmentID: roleAssignmentID,
	}
}

// Execute executes the request
func (a *ManagementAPIsUsersUserRoleAssignmentsApiService) V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsRoleAssignmentIDGetExecute(r ApiV1EnvironmentsEnvIDUsersUserIDRoleAssignmentsRoleAssignmentIDGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsUsersUserRoleAssignmentsApiService.V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsRoleAssignmentIDGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/users/{userID}/roleAssignments/{roleAssignmentID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userID"+"}", _neturl.PathEscape(parameterToString(r.userID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleAssignmentID"+"}", _neturl.PathEscape(parameterToString(r.roleAssignmentID, "")), -1)

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
