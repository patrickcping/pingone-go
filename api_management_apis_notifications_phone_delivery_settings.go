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

// ManagementAPIsNotificationsPhoneDeliverySettingsApiService ManagementAPIsNotificationsPhoneDeliverySettingsApi service
type ManagementAPIsNotificationsPhoneDeliverySettingsApiService service

type ApiV1EnvironmentsEnvIDNotificationsSettingsEmailDeliverySettingsDeleteRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsNotificationsPhoneDeliverySettingsApiService
	envID string
}


func (r ApiV1EnvironmentsEnvIDNotificationsSettingsEmailDeliverySettingsDeleteRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDNotificationsSettingsEmailDeliverySettingsDeleteExecute(r)
}

/*
V1EnvironmentsEnvIDNotificationsSettingsEmailDeliverySettingsDelete DELETE Phone Delivery Settings

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @return ApiV1EnvironmentsEnvIDNotificationsSettingsEmailDeliverySettingsDeleteRequest
*/
func (a *ManagementAPIsNotificationsPhoneDeliverySettingsApiService) V1EnvironmentsEnvIDNotificationsSettingsEmailDeliverySettingsDelete(ctx _context.Context, envID string) ApiV1EnvironmentsEnvIDNotificationsSettingsEmailDeliverySettingsDeleteRequest {
	return ApiV1EnvironmentsEnvIDNotificationsSettingsEmailDeliverySettingsDeleteRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
	}
}

// Execute executes the request
func (a *ManagementAPIsNotificationsPhoneDeliverySettingsApiService) V1EnvironmentsEnvIDNotificationsSettingsEmailDeliverySettingsDeleteExecute(r ApiV1EnvironmentsEnvIDNotificationsSettingsEmailDeliverySettingsDeleteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsNotificationsPhoneDeliverySettingsApiService.V1EnvironmentsEnvIDNotificationsSettingsEmailDeliverySettingsDelete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/notificationsSettings/emailDeliverySettings"
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

type ApiV1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsNotificationsPhoneDeliverySettingsApiService
	envID string
}


func (r ApiV1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsGetExecute(r)
}

/*
V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsGet READ All Phone Delivery Settings

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @return ApiV1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsGetRequest
*/
func (a *ManagementAPIsNotificationsPhoneDeliverySettingsApiService) V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsGet(ctx _context.Context, envID string) ApiV1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsGetRequest {
	return ApiV1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
	}
}

// Execute executes the request
func (a *ManagementAPIsNotificationsPhoneDeliverySettingsApiService) V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsGetExecute(r ApiV1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsNotificationsPhoneDeliverySettingsApiService.V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/notificationsSettings/phoneDeliverySettings"
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

type ApiV1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsNotificationsPhoneDeliverySettingsApiService
	envID string
	phoneDeliverySettingsId string
}


func (r ApiV1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdGetExecute(r)
}

/*
V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdGet READ One Phone Delivery Settings

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param phoneDeliverySettingsId
 @return ApiV1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdGetRequest
*/
func (a *ManagementAPIsNotificationsPhoneDeliverySettingsApiService) V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdGet(ctx _context.Context, envID string, phoneDeliverySettingsId string) ApiV1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdGetRequest {
	return ApiV1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		phoneDeliverySettingsId: phoneDeliverySettingsId,
	}
}

// Execute executes the request
func (a *ManagementAPIsNotificationsPhoneDeliverySettingsApiService) V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdGetExecute(r ApiV1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsNotificationsPhoneDeliverySettingsApiService.V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/notificationsSettings/phoneDeliverySettings/{phoneDeliverySettingsId}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"phoneDeliverySettingsId"+"}", _neturl.PathEscape(parameterToString(r.phoneDeliverySettingsId, "")), -1)

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

type ApiV1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdPutRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsNotificationsPhoneDeliverySettingsApiService
	envID string
	phoneDeliverySettingsId string
	contentType *string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdPutRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdPutRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdPutRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdPutRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdPutRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdPutExecute(r)
}

/*
V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdPut UPDATE Phone Delivery Settings

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param phoneDeliverySettingsId
 @return ApiV1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdPutRequest
*/
func (a *ManagementAPIsNotificationsPhoneDeliverySettingsApiService) V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdPut(ctx _context.Context, envID string, phoneDeliverySettingsId string) ApiV1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdPutRequest {
	return ApiV1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdPutRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		phoneDeliverySettingsId: phoneDeliverySettingsId,
	}
}

// Execute executes the request
func (a *ManagementAPIsNotificationsPhoneDeliverySettingsApiService) V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdPutExecute(r ApiV1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdPutRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsNotificationsPhoneDeliverySettingsApiService.V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdPut")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/notificationsSettings/phoneDeliverySettings/{phoneDeliverySettingsId}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"phoneDeliverySettingsId"+"}", _neturl.PathEscape(parameterToString(r.phoneDeliverySettingsId, "")), -1)

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

type ApiV1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPostRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsNotificationsPhoneDeliverySettingsApiService
	envID string
	contentType *string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPostRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPostRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPostRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPostRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPostExecute(r)
}

/*
V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPost CREATE Phone Delivery Settings (Syniverse)

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @return ApiV1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPostRequest
*/
func (a *ManagementAPIsNotificationsPhoneDeliverySettingsApiService) V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPost(ctx _context.Context, envID string) ApiV1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPostRequest {
	return ApiV1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPostRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
	}
}

// Execute executes the request
func (a *ManagementAPIsNotificationsPhoneDeliverySettingsApiService) V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPostExecute(r ApiV1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsNotificationsPhoneDeliverySettingsApiService.V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/notificationsSettings/phoneDeliverySettings"
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
