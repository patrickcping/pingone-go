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

// ManagementAPIsUsersUserAgreementConsentsApiService ManagementAPIsUsersUserAgreementConsentsApi service
type ManagementAPIsUsersUserAgreementConsentsApiService service

type ApiV1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsUsersUserAgreementConsentsApiService
	envID string
	userID string
	agreementID string
	authorization *string
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDGetRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDGetRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDGetExecute(r)
}

/*
 * V1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDGet READ One User Agreement Consent
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @param userID
 * @param agreementID
 * @return ApiV1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDGetRequest
 */
func (a *ManagementAPIsUsersUserAgreementConsentsApiService) V1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDGet(ctx _context.Context, envID string, userID string, agreementID string) ApiV1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDGetRequest {
	return ApiV1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		userID: userID,
		agreementID: agreementID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsUsersUserAgreementConsentsApiService) V1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDGetExecute(r ApiV1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsUsersUserAgreementConsentsApiService.V1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/users/{userID}/agreementConsents/{agreementID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userID"+"}", _neturl.PathEscape(parameterToString(r.userID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"agreementID"+"}", _neturl.PathEscape(parameterToString(r.agreementID, "")), -1)

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

type ApiV1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDPostRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsUsersUserAgreementConsentsApiService
	envID string
	userID string
	agreementID string
	authorization *string
	contentType *string
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDPostRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDPostRequest {
	r.authorization = &authorization
	return r
}
func (r ApiV1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDPostRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDPostRequest {
	r.contentType = &contentType
	return r
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDPostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDPostExecute(r)
}

/*
 * V1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDPost Revoke Agreement
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @param userID
 * @param agreementID
 * @return ApiV1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDPostRequest
 */
func (a *ManagementAPIsUsersUserAgreementConsentsApiService) V1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDPost(ctx _context.Context, envID string, userID string, agreementID string) ApiV1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDPostRequest {
	return ApiV1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDPostRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		userID: userID,
		agreementID: agreementID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsUsersUserAgreementConsentsApiService) V1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDPostExecute(r ApiV1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDPostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsUsersUserAgreementConsentsApiService.V1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDPost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/users/{userID}/agreementConsents/{agreementID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userID"+"}", _neturl.PathEscape(parameterToString(r.userID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"agreementID"+"}", _neturl.PathEscape(parameterToString(r.agreementID, "")), -1)

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
	if r.contentType != nil {
		localVarHeaderParams["Content-Type"] = parameterToString(*r.contentType, "")
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

type ApiV1EnvironmentsEnvIDUsersUserIDAgreementConsentsGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsUsersUserAgreementConsentsApiService
	envID string
	userID string
	authorization *string
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDAgreementConsentsGetRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDUsersUserIDAgreementConsentsGetRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDAgreementConsentsGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDUsersUserIDAgreementConsentsGetExecute(r)
}

/*
 * V1EnvironmentsEnvIDUsersUserIDAgreementConsentsGet READ All User Agreement Consents
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @param userID
 * @return ApiV1EnvironmentsEnvIDUsersUserIDAgreementConsentsGetRequest
 */
func (a *ManagementAPIsUsersUserAgreementConsentsApiService) V1EnvironmentsEnvIDUsersUserIDAgreementConsentsGet(ctx _context.Context, envID string, userID string) ApiV1EnvironmentsEnvIDUsersUserIDAgreementConsentsGetRequest {
	return ApiV1EnvironmentsEnvIDUsersUserIDAgreementConsentsGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		userID: userID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsUsersUserAgreementConsentsApiService) V1EnvironmentsEnvIDUsersUserIDAgreementConsentsGetExecute(r ApiV1EnvironmentsEnvIDUsersUserIDAgreementConsentsGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsUsersUserAgreementConsentsApiService.V1EnvironmentsEnvIDUsersUserIDAgreementConsentsGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/users/{userID}/agreementConsents"
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
