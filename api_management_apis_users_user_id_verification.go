/*
 * PingOne Platform API
 *
 * A bare-bones collection for the PingOne API
 *
 * API version: 1.0.0
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

// ManagementAPIsUsersUserIDVerificationApiService ManagementAPIsUsersUserIDVerificationApi service
type ManagementAPIsUsersUserIDVerificationApiService service

type ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsUsersUserIDVerificationApiService
	envID string
	userID string
	authorization *string
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsGetRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsGetRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsGetExecute(r)
}

/*
 * V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsGet READ All ID Verification Transaction Records for a User
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @param userID
 * @return ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsGetRequest
 */
func (a *ManagementAPIsUsersUserIDVerificationApiService) V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsGet(ctx _context.Context, envID string, userID string) ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsGetRequest {
	return ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		userID: userID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsUsersUserIDVerificationApiService) V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsGetExecute(r ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsUsersUserIDVerificationApiService.V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/users/{userID}/verifyTransactions"
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

type ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsPostRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsUsersUserIDVerificationApiService
	envID string
	userID string
	contentType *string
	authorization *string
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsPostRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsPostRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsPostRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsPostRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsPostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsPostExecute(r)
}

/*
 * V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsPost CREATE ID Verification Transaction Record for a User
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @param userID
 * @return ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsPostRequest
 */
func (a *ManagementAPIsUsersUserIDVerificationApiService) V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsPost(ctx _context.Context, envID string, userID string) ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsPostRequest {
	return ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsPostRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		userID: userID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsUsersUserIDVerificationApiService) V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsPostExecute(r ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsPostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsUsersUserIDVerificationApiService.V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsPost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/users/{userID}/verifyTransactions"
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
	if r.contentType != nil {
		localVarHeaderParams["Content-Type"] = parameterToString(*r.contentType, "")
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

type ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDDeleteRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsUsersUserIDVerificationApiService
	envID string
	userID string
	transactionID string
	authorization *string
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDDeleteRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDDeleteRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDDeleteRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDDeleteExecute(r)
}

/*
 * V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDDelete DELETE ID Verification Transaction Record for a User
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @param userID
 * @param transactionID
 * @return ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDDeleteRequest
 */
func (a *ManagementAPIsUsersUserIDVerificationApiService) V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDDelete(ctx _context.Context, envID string, userID string, transactionID string) ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDDeleteRequest {
	return ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDDeleteRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		userID: userID,
		transactionID: transactionID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsUsersUserIDVerificationApiService) V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDDeleteExecute(r ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDDeleteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsUsersUserIDVerificationApiService.V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDDelete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/users/{userID}/verifyTransactions/{transactionID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userID"+"}", _neturl.PathEscape(parameterToString(r.userID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"transactionID"+"}", _neturl.PathEscape(parameterToString(r.transactionID, "")), -1)

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

type ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsUsersUserIDVerificationApiService
	envID string
	userID string
	transactionID string
	authorization *string
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDGetRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDGetRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDGetExecute(r)
}

/*
 * V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDGet READ ID Verification Transaction Record for a User
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @param userID
 * @param transactionID
 * @return ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDGetRequest
 */
func (a *ManagementAPIsUsersUserIDVerificationApiService) V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDGet(ctx _context.Context, envID string, userID string, transactionID string) ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDGetRequest {
	return ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		userID: userID,
		transactionID: transactionID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsUsersUserIDVerificationApiService) V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDGetExecute(r ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsUsersUserIDVerificationApiService.V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/users/{userID}/verifyTransactions/{transactionID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userID"+"}", _neturl.PathEscape(parameterToString(r.userID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"transactionID"+"}", _neturl.PathEscape(parameterToString(r.transactionID, "")), -1)

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

type ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDPutRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsUsersUserIDVerificationApiService
	envID string
	userID string
	transactionID string
	authorization *string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDPutRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDPutRequest {
	r.authorization = &authorization
	return r
}
func (r ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDPutRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDPutRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDPutRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDPutExecute(r)
}

/*
 * V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDPut UPDATE ID Verification Transaction Record for a User
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @param userID
 * @param transactionID
 * @return ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDPutRequest
 */
func (a *ManagementAPIsUsersUserIDVerificationApiService) V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDPut(ctx _context.Context, envID string, userID string, transactionID string) ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDPutRequest {
	return ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDPutRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		userID: userID,
		transactionID: transactionID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsUsersUserIDVerificationApiService) V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDPutExecute(r ApiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDPutRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsUsersUserIDVerificationApiService.V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDPut")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/users/{userID}/verifyTransactions/{transactionID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userID"+"}", _neturl.PathEscape(parameterToString(r.userID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"transactionID"+"}", _neturl.PathEscape(parameterToString(r.transactionID, "")), -1)

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
