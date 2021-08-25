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

// ManagementAPIsUsersMFAPairingKeysApiService ManagementAPIsUsersMFAPairingKeysApi service
type ManagementAPIsUsersMFAPairingKeysApiService service

type ApiV1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDDeleteRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsUsersMFAPairingKeysApiService
	envID string
	userID string
	pairingKeyID string
	authorization *string
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDDeleteRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDDeleteRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDDeleteRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDDeleteExecute(r)
}

/*
 * V1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDDelete DELETE MFA Pairing Key
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @param userID
 * @param pairingKeyID
 * @return ApiV1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDDeleteRequest
 */
func (a *ManagementAPIsUsersMFAPairingKeysApiService) V1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDDelete(ctx _context.Context, envID string, userID string, pairingKeyID string) ApiV1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDDeleteRequest {
	return ApiV1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDDeleteRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		userID: userID,
		pairingKeyID: pairingKeyID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsUsersMFAPairingKeysApiService) V1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDDeleteExecute(r ApiV1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDDeleteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsUsersMFAPairingKeysApiService.V1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDDelete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/users/{userID}/pairingKeys/{pairingKeyID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userID"+"}", _neturl.PathEscape(parameterToString(r.userID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pairingKeyID"+"}", _neturl.PathEscape(parameterToString(r.pairingKeyID, "")), -1)

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

type ApiV1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsUsersMFAPairingKeysApiService
	envID string
	userID string
	pairingKeyID string
	authorization *string
	contentType *string
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDGetRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDGetRequest {
	r.authorization = &authorization
	return r
}
func (r ApiV1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDGetRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDGetRequest {
	r.contentType = &contentType
	return r
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDGetExecute(r)
}

/*
 * V1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDGet READ One MFA Pairing Key
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @param userID
 * @param pairingKeyID
 * @return ApiV1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDGetRequest
 */
func (a *ManagementAPIsUsersMFAPairingKeysApiService) V1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDGet(ctx _context.Context, envID string, userID string, pairingKeyID string) ApiV1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDGetRequest {
	return ApiV1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		userID: userID,
		pairingKeyID: pairingKeyID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsUsersMFAPairingKeysApiService) V1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDGetExecute(r ApiV1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsUsersMFAPairingKeysApiService.V1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/users/{userID}/pairingKeys/{pairingKeyID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userID"+"}", _neturl.PathEscape(parameterToString(r.userID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pairingKeyID"+"}", _neturl.PathEscape(parameterToString(r.pairingKeyID, "")), -1)

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

type ApiV1EnvironmentsEnvIDUsersUserIDPairingKeysPostRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsUsersMFAPairingKeysApiService
	envID string
	userID string
	contentType *string
	authorization *string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDPairingKeysPostRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDUsersUserIDPairingKeysPostRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDUsersUserIDPairingKeysPostRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDUsersUserIDPairingKeysPostRequest {
	r.authorization = &authorization
	return r
}
func (r ApiV1EnvironmentsEnvIDUsersUserIDPairingKeysPostRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDUsersUserIDPairingKeysPostRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDPairingKeysPostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDUsersUserIDPairingKeysPostExecute(r)
}

/*
 * V1EnvironmentsEnvIDUsersUserIDPairingKeysPost CREATE MFA Pairing Key
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @param userID
 * @return ApiV1EnvironmentsEnvIDUsersUserIDPairingKeysPostRequest
 */
func (a *ManagementAPIsUsersMFAPairingKeysApiService) V1EnvironmentsEnvIDUsersUserIDPairingKeysPost(ctx _context.Context, envID string, userID string) ApiV1EnvironmentsEnvIDUsersUserIDPairingKeysPostRequest {
	return ApiV1EnvironmentsEnvIDUsersUserIDPairingKeysPostRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		userID: userID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsUsersMFAPairingKeysApiService) V1EnvironmentsEnvIDUsersUserIDPairingKeysPostExecute(r ApiV1EnvironmentsEnvIDUsersUserIDPairingKeysPostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsUsersMFAPairingKeysApiService.V1EnvironmentsEnvIDUsersUserIDPairingKeysPost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/users/{userID}/pairingKeys"
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
