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

// ManagementAPIsNotificationsTrustedEmailDomainsApiService ManagementAPIsNotificationsTrustedEmailDomainsApi service
type ManagementAPIsNotificationsTrustedEmailDomainsApiService service

type ApiV1EnvironmentsEnvIDEmailDomainsEmailDomainIdDeleteRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsNotificationsTrustedEmailDomainsApiService
	envID string
	emailDomainId string
}


func (r ApiV1EnvironmentsEnvIDEmailDomainsEmailDomainIdDeleteRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDEmailDomainsEmailDomainIdDeleteExecute(r)
}

/*
V1EnvironmentsEnvIDEmailDomainsEmailDomainIdDelete DELETE Trusted Email Domain

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param emailDomainId
 @return ApiV1EnvironmentsEnvIDEmailDomainsEmailDomainIdDeleteRequest
*/
func (a *ManagementAPIsNotificationsTrustedEmailDomainsApiService) V1EnvironmentsEnvIDEmailDomainsEmailDomainIdDelete(ctx _context.Context, envID string, emailDomainId string) ApiV1EnvironmentsEnvIDEmailDomainsEmailDomainIdDeleteRequest {
	return ApiV1EnvironmentsEnvIDEmailDomainsEmailDomainIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		emailDomainId: emailDomainId,
	}
}

// Execute executes the request
func (a *ManagementAPIsNotificationsTrustedEmailDomainsApiService) V1EnvironmentsEnvIDEmailDomainsEmailDomainIdDeleteExecute(r ApiV1EnvironmentsEnvIDEmailDomainsEmailDomainIdDeleteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsNotificationsTrustedEmailDomainsApiService.V1EnvironmentsEnvIDEmailDomainsEmailDomainIdDelete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/emailDomains/{emailDomainId}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"emailDomainId"+"}", _neturl.PathEscape(parameterToString(r.emailDomainId, "")), -1)

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

type ApiV1EnvironmentsEnvIDEmailDomainsEmailDomainIdDkimGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsNotificationsTrustedEmailDomainsApiService
	envID string
	emailDomainId string
}


func (r ApiV1EnvironmentsEnvIDEmailDomainsEmailDomainIdDkimGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDEmailDomainsEmailDomainIdDkimGetExecute(r)
}

/*
V1EnvironmentsEnvIDEmailDomainsEmailDomainIdDkimGet READ Trusted Email Domain DKIM Status

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param emailDomainId
 @return ApiV1EnvironmentsEnvIDEmailDomainsEmailDomainIdDkimGetRequest
*/
func (a *ManagementAPIsNotificationsTrustedEmailDomainsApiService) V1EnvironmentsEnvIDEmailDomainsEmailDomainIdDkimGet(ctx _context.Context, envID string, emailDomainId string) ApiV1EnvironmentsEnvIDEmailDomainsEmailDomainIdDkimGetRequest {
	return ApiV1EnvironmentsEnvIDEmailDomainsEmailDomainIdDkimGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		emailDomainId: emailDomainId,
	}
}

// Execute executes the request
func (a *ManagementAPIsNotificationsTrustedEmailDomainsApiService) V1EnvironmentsEnvIDEmailDomainsEmailDomainIdDkimGetExecute(r ApiV1EnvironmentsEnvIDEmailDomainsEmailDomainIdDkimGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsNotificationsTrustedEmailDomainsApiService.V1EnvironmentsEnvIDEmailDomainsEmailDomainIdDkimGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/emailDomains/{emailDomainId}/dkim"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"emailDomainId"+"}", _neturl.PathEscape(parameterToString(r.emailDomainId, "")), -1)

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

type ApiV1EnvironmentsEnvIDEmailDomainsEmailDomainIdGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsNotificationsTrustedEmailDomainsApiService
	envID string
	emailDomainId string
}


func (r ApiV1EnvironmentsEnvIDEmailDomainsEmailDomainIdGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDEmailDomainsEmailDomainIdGetExecute(r)
}

/*
V1EnvironmentsEnvIDEmailDomainsEmailDomainIdGet READ One Trusted Email Domain

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param emailDomainId
 @return ApiV1EnvironmentsEnvIDEmailDomainsEmailDomainIdGetRequest
*/
func (a *ManagementAPIsNotificationsTrustedEmailDomainsApiService) V1EnvironmentsEnvIDEmailDomainsEmailDomainIdGet(ctx _context.Context, envID string, emailDomainId string) ApiV1EnvironmentsEnvIDEmailDomainsEmailDomainIdGetRequest {
	return ApiV1EnvironmentsEnvIDEmailDomainsEmailDomainIdGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		emailDomainId: emailDomainId,
	}
}

// Execute executes the request
func (a *ManagementAPIsNotificationsTrustedEmailDomainsApiService) V1EnvironmentsEnvIDEmailDomainsEmailDomainIdGetExecute(r ApiV1EnvironmentsEnvIDEmailDomainsEmailDomainIdGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsNotificationsTrustedEmailDomainsApiService.V1EnvironmentsEnvIDEmailDomainsEmailDomainIdGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/emailDomains/{emailDomainId}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"emailDomainId"+"}", _neturl.PathEscape(parameterToString(r.emailDomainId, "")), -1)

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

type ApiV1EnvironmentsEnvIDEmailDomainsEmailDomainIdOwnershipGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsNotificationsTrustedEmailDomainsApiService
	envID string
	emailDomainId string
}


func (r ApiV1EnvironmentsEnvIDEmailDomainsEmailDomainIdOwnershipGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDEmailDomainsEmailDomainIdOwnershipGetExecute(r)
}

/*
V1EnvironmentsEnvIDEmailDomainsEmailDomainIdOwnershipGet READ Trusted Email Domain Ownership Status

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param emailDomainId
 @return ApiV1EnvironmentsEnvIDEmailDomainsEmailDomainIdOwnershipGetRequest
*/
func (a *ManagementAPIsNotificationsTrustedEmailDomainsApiService) V1EnvironmentsEnvIDEmailDomainsEmailDomainIdOwnershipGet(ctx _context.Context, envID string, emailDomainId string) ApiV1EnvironmentsEnvIDEmailDomainsEmailDomainIdOwnershipGetRequest {
	return ApiV1EnvironmentsEnvIDEmailDomainsEmailDomainIdOwnershipGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		emailDomainId: emailDomainId,
	}
}

// Execute executes the request
func (a *ManagementAPIsNotificationsTrustedEmailDomainsApiService) V1EnvironmentsEnvIDEmailDomainsEmailDomainIdOwnershipGetExecute(r ApiV1EnvironmentsEnvIDEmailDomainsEmailDomainIdOwnershipGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsNotificationsTrustedEmailDomainsApiService.V1EnvironmentsEnvIDEmailDomainsEmailDomainIdOwnershipGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/emailDomains/{emailDomainId}/ownership"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"emailDomainId"+"}", _neturl.PathEscape(parameterToString(r.emailDomainId, "")), -1)

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

type ApiV1EnvironmentsEnvIDEmailDomainsEmailDomainIdSpfGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsNotificationsTrustedEmailDomainsApiService
	envID string
	emailDomainId string
}


func (r ApiV1EnvironmentsEnvIDEmailDomainsEmailDomainIdSpfGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDEmailDomainsEmailDomainIdSpfGetExecute(r)
}

/*
V1EnvironmentsEnvIDEmailDomainsEmailDomainIdSpfGet READ Trusted Email Domain SPF Status

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param emailDomainId
 @return ApiV1EnvironmentsEnvIDEmailDomainsEmailDomainIdSpfGetRequest
*/
func (a *ManagementAPIsNotificationsTrustedEmailDomainsApiService) V1EnvironmentsEnvIDEmailDomainsEmailDomainIdSpfGet(ctx _context.Context, envID string, emailDomainId string) ApiV1EnvironmentsEnvIDEmailDomainsEmailDomainIdSpfGetRequest {
	return ApiV1EnvironmentsEnvIDEmailDomainsEmailDomainIdSpfGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		emailDomainId: emailDomainId,
	}
}

// Execute executes the request
func (a *ManagementAPIsNotificationsTrustedEmailDomainsApiService) V1EnvironmentsEnvIDEmailDomainsEmailDomainIdSpfGetExecute(r ApiV1EnvironmentsEnvIDEmailDomainsEmailDomainIdSpfGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsNotificationsTrustedEmailDomainsApiService.V1EnvironmentsEnvIDEmailDomainsEmailDomainIdSpfGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/emailDomains/{emailDomainId}/spf"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"emailDomainId"+"}", _neturl.PathEscape(parameterToString(r.emailDomainId, "")), -1)

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

type ApiV1EnvironmentsEnvIDEmailDomainsGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsNotificationsTrustedEmailDomainsApiService
	envID string
}


func (r ApiV1EnvironmentsEnvIDEmailDomainsGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDEmailDomainsGetExecute(r)
}

/*
V1EnvironmentsEnvIDEmailDomainsGet READ All Trusted Email Domains

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @return ApiV1EnvironmentsEnvIDEmailDomainsGetRequest
*/
func (a *ManagementAPIsNotificationsTrustedEmailDomainsApiService) V1EnvironmentsEnvIDEmailDomainsGet(ctx _context.Context, envID string) ApiV1EnvironmentsEnvIDEmailDomainsGetRequest {
	return ApiV1EnvironmentsEnvIDEmailDomainsGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
	}
}

// Execute executes the request
func (a *ManagementAPIsNotificationsTrustedEmailDomainsApiService) V1EnvironmentsEnvIDEmailDomainsGetExecute(r ApiV1EnvironmentsEnvIDEmailDomainsGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsNotificationsTrustedEmailDomainsApiService.V1EnvironmentsEnvIDEmailDomainsGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/emailDomains"
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

type ApiV1EnvironmentsEnvIDEmailDomainsPostRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsNotificationsTrustedEmailDomainsApiService
	envID string
	contentType *string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDEmailDomainsPostRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDEmailDomainsPostRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDEmailDomainsPostRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDEmailDomainsPostRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDEmailDomainsPostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDEmailDomainsPostExecute(r)
}

/*
V1EnvironmentsEnvIDEmailDomainsPost CREATE Trusted Email Domain

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @return ApiV1EnvironmentsEnvIDEmailDomainsPostRequest
*/
func (a *ManagementAPIsNotificationsTrustedEmailDomainsApiService) V1EnvironmentsEnvIDEmailDomainsPost(ctx _context.Context, envID string) ApiV1EnvironmentsEnvIDEmailDomainsPostRequest {
	return ApiV1EnvironmentsEnvIDEmailDomainsPostRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
	}
}

// Execute executes the request
func (a *ManagementAPIsNotificationsTrustedEmailDomainsApiService) V1EnvironmentsEnvIDEmailDomainsPostExecute(r ApiV1EnvironmentsEnvIDEmailDomainsPostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsNotificationsTrustedEmailDomainsApiService.V1EnvironmentsEnvIDEmailDomainsPost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/emailDomains"
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
