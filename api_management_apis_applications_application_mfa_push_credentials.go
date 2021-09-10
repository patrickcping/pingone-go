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

// ManagementAPIsApplicationsApplicationMFAPushCredentialsApiService ManagementAPIsApplicationsApplicationMFAPushCredentialsApi service
type ManagementAPIsApplicationsApplicationMFAPushCredentialsApiService service

type ApiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsApplicationsApplicationMFAPushCredentialsApiService
	envID string
	appID string
}


func (r ApiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsGetExecute(r)
}

/*
V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsGet READ All MFA Push Credentials

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param appID
 @return ApiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsGetRequest
*/
func (a *ManagementAPIsApplicationsApplicationMFAPushCredentialsApiService) V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsGet(ctx _context.Context, envID string, appID string) ApiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsGetRequest {
	return ApiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		appID: appID,
	}
}

// Execute executes the request
func (a *ManagementAPIsApplicationsApplicationMFAPushCredentialsApiService) V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsGetExecute(r ApiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsApplicationsApplicationMFAPushCredentialsApiService.V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/applications/{appID}/pushCredentials"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appID"+"}", _neturl.PathEscape(parameterToString(r.appID, "")), -1)

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

type ApiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPostRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsApplicationsApplicationMFAPushCredentialsApiService
	envID string
	appID string
	contentType *string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPostRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPostRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPostRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPostRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPostExecute(r)
}

/*
V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPost CREATE MFA Push Credential (FCM)

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param appID
 @return ApiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPostRequest
*/
func (a *ManagementAPIsApplicationsApplicationMFAPushCredentialsApiService) V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPost(ctx _context.Context, envID string, appID string) ApiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPostRequest {
	return ApiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPostRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		appID: appID,
	}
}

// Execute executes the request
func (a *ManagementAPIsApplicationsApplicationMFAPushCredentialsApiService) V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPostExecute(r ApiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsApplicationsApplicationMFAPushCredentialsApiService.V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/applications/{appID}/pushCredentials"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appID"+"}", _neturl.PathEscape(parameterToString(r.appID, "")), -1)

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

type ApiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDDeleteRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsApplicationsApplicationMFAPushCredentialsApiService
	envID string
	appID string
	pushCredID string
	authorization *string
}

func (r ApiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDDeleteRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDDeleteRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDDeleteRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDDeleteExecute(r)
}

/*
V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDDelete DELETE MFA Push Credential

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param appID
 @param pushCredID
 @return ApiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDDeleteRequest
*/
func (a *ManagementAPIsApplicationsApplicationMFAPushCredentialsApiService) V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDDelete(ctx _context.Context, envID string, appID string, pushCredID string) ApiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDDeleteRequest {
	return ApiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDDeleteRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		appID: appID,
		pushCredID: pushCredID,
	}
}

// Execute executes the request
func (a *ManagementAPIsApplicationsApplicationMFAPushCredentialsApiService) V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDDeleteExecute(r ApiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDDeleteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsApplicationsApplicationMFAPushCredentialsApiService.V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDDelete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/applications/{appID}/pushCredentials/{pushCredID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appID"+"}", _neturl.PathEscape(parameterToString(r.appID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pushCredID"+"}", _neturl.PathEscape(parameterToString(r.pushCredID, "")), -1)

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

type ApiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsApplicationsApplicationMFAPushCredentialsApiService
	envID string
	appID string
	pushCredID string
}


func (r ApiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDGetExecute(r)
}

/*
V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDGet READ One MFA Push Credential

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param appID
 @param pushCredID
 @return ApiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDGetRequest
*/
func (a *ManagementAPIsApplicationsApplicationMFAPushCredentialsApiService) V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDGet(ctx _context.Context, envID string, appID string, pushCredID string) ApiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDGetRequest {
	return ApiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		appID: appID,
		pushCredID: pushCredID,
	}
}

// Execute executes the request
func (a *ManagementAPIsApplicationsApplicationMFAPushCredentialsApiService) V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDGetExecute(r ApiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsApplicationsApplicationMFAPushCredentialsApiService.V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/applications/{appID}/pushCredentials/{pushCredID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appID"+"}", _neturl.PathEscape(parameterToString(r.appID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pushCredID"+"}", _neturl.PathEscape(parameterToString(r.pushCredID, "")), -1)

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

type ApiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDPutRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsApplicationsApplicationMFAPushCredentialsApiService
	envID string
	appID string
	pushCredID string
	contentType *string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDPutRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDPutRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDPutRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDPutRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDPutRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDPutExecute(r)
}

/*
V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDPut UPDATE MFA Push Credential

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param appID
 @param pushCredID
 @return ApiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDPutRequest
*/
func (a *ManagementAPIsApplicationsApplicationMFAPushCredentialsApiService) V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDPut(ctx _context.Context, envID string, appID string, pushCredID string) ApiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDPutRequest {
	return ApiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDPutRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		appID: appID,
		pushCredID: pushCredID,
	}
}

// Execute executes the request
func (a *ManagementAPIsApplicationsApplicationMFAPushCredentialsApiService) V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDPutExecute(r ApiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDPutRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsApplicationsApplicationMFAPushCredentialsApiService.V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDPut")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/applications/{appID}/pushCredentials/{pushCredID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appID"+"}", _neturl.PathEscape(parameterToString(r.appID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pushCredID"+"}", _neturl.PathEscape(parameterToString(r.pushCredID, "")), -1)

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
