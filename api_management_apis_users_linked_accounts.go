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

// ManagementAPIsUsersLinkedAccountsApiService ManagementAPIsUsersLinkedAccountsApi service
type ManagementAPIsUsersLinkedAccountsApiService service

type ApiV1EnvironmentsEnvIDUsersUserIDLinkedAccountsGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsUsersLinkedAccountsApiService
	envID string
	userID string
	authorization *string
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDLinkedAccountsGetRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDUsersUserIDLinkedAccountsGetRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDLinkedAccountsGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDUsersUserIDLinkedAccountsGetExecute(r)
}

/*
 * V1EnvironmentsEnvIDUsersUserIDLinkedAccountsGet READ Linked Accounts
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @param userID
 * @return ApiV1EnvironmentsEnvIDUsersUserIDLinkedAccountsGetRequest
 */
func (a *ManagementAPIsUsersLinkedAccountsApiService) V1EnvironmentsEnvIDUsersUserIDLinkedAccountsGet(ctx _context.Context, envID string, userID string) ApiV1EnvironmentsEnvIDUsersUserIDLinkedAccountsGetRequest {
	return ApiV1EnvironmentsEnvIDUsersUserIDLinkedAccountsGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		userID: userID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsUsersLinkedAccountsApiService) V1EnvironmentsEnvIDUsersUserIDLinkedAccountsGetExecute(r ApiV1EnvironmentsEnvIDUsersUserIDLinkedAccountsGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsUsersLinkedAccountsApiService.V1EnvironmentsEnvIDUsersUserIDLinkedAccountsGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/users/{userID}/linkedAccounts"
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

type ApiV1EnvironmentsEnvIDUsersUserIDLinkedAccountsLinkedAccountIDDeleteRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsUsersLinkedAccountsApiService
	envID string
	userID string
	linkedAccountID string
	authorization *string
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDLinkedAccountsLinkedAccountIDDeleteRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDUsersUserIDLinkedAccountsLinkedAccountIDDeleteRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDLinkedAccountsLinkedAccountIDDeleteRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDUsersUserIDLinkedAccountsLinkedAccountIDDeleteExecute(r)
}

/*
 * V1EnvironmentsEnvIDUsersUserIDLinkedAccountsLinkedAccountIDDelete DELETE Linked Account
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @param userID
 * @param linkedAccountID
 * @return ApiV1EnvironmentsEnvIDUsersUserIDLinkedAccountsLinkedAccountIDDeleteRequest
 */
func (a *ManagementAPIsUsersLinkedAccountsApiService) V1EnvironmentsEnvIDUsersUserIDLinkedAccountsLinkedAccountIDDelete(ctx _context.Context, envID string, userID string, linkedAccountID string) ApiV1EnvironmentsEnvIDUsersUserIDLinkedAccountsLinkedAccountIDDeleteRequest {
	return ApiV1EnvironmentsEnvIDUsersUserIDLinkedAccountsLinkedAccountIDDeleteRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		userID: userID,
		linkedAccountID: linkedAccountID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsUsersLinkedAccountsApiService) V1EnvironmentsEnvIDUsersUserIDLinkedAccountsLinkedAccountIDDeleteExecute(r ApiV1EnvironmentsEnvIDUsersUserIDLinkedAccountsLinkedAccountIDDeleteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsUsersLinkedAccountsApiService.V1EnvironmentsEnvIDUsersUserIDLinkedAccountsLinkedAccountIDDelete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/users/{userID}/linkedAccounts/{linkedAccountID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userID"+"}", _neturl.PathEscape(parameterToString(r.userID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"linkedAccountID"+"}", _neturl.PathEscape(parameterToString(r.linkedAccountID, "")), -1)

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

type ApiV1EnvironmentsEnvIDUsersUserIDLinkedAccountsLinkedAccountIDGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsUsersLinkedAccountsApiService
	envID string
	userID string
	linkedAccountID string
	authorization *string
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDLinkedAccountsLinkedAccountIDGetRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDUsersUserIDLinkedAccountsLinkedAccountIDGetRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDLinkedAccountsLinkedAccountIDGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDUsersUserIDLinkedAccountsLinkedAccountIDGetExecute(r)
}

/*
 * V1EnvironmentsEnvIDUsersUserIDLinkedAccountsLinkedAccountIDGet READ One Linked Account
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @param userID
 * @param linkedAccountID
 * @return ApiV1EnvironmentsEnvIDUsersUserIDLinkedAccountsLinkedAccountIDGetRequest
 */
func (a *ManagementAPIsUsersLinkedAccountsApiService) V1EnvironmentsEnvIDUsersUserIDLinkedAccountsLinkedAccountIDGet(ctx _context.Context, envID string, userID string, linkedAccountID string) ApiV1EnvironmentsEnvIDUsersUserIDLinkedAccountsLinkedAccountIDGetRequest {
	return ApiV1EnvironmentsEnvIDUsersUserIDLinkedAccountsLinkedAccountIDGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		userID: userID,
		linkedAccountID: linkedAccountID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsUsersLinkedAccountsApiService) V1EnvironmentsEnvIDUsersUserIDLinkedAccountsLinkedAccountIDGetExecute(r ApiV1EnvironmentsEnvIDUsersUserIDLinkedAccountsLinkedAccountIDGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsUsersLinkedAccountsApiService.V1EnvironmentsEnvIDUsersUserIDLinkedAccountsLinkedAccountIDGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/users/{userID}/linkedAccounts/{linkedAccountID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userID"+"}", _neturl.PathEscape(parameterToString(r.userID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"linkedAccountID"+"}", _neturl.PathEscape(parameterToString(r.linkedAccountID, "")), -1)

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
