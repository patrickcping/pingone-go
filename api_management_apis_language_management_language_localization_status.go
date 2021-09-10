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

// ManagementAPIsLanguageManagementLanguageLocalizationStatusApiService ManagementAPIsLanguageManagementLanguageLocalizationStatusApi service
type ManagementAPIsLanguageManagementLanguageLocalizationStatusApiService service

type ApiV1EnvironmentsEnvIDLanguagesLanguageIDStatusGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsLanguageManagementLanguageLocalizationStatusApiService
	envID string
	languageID string
}


func (r ApiV1EnvironmentsEnvIDLanguagesLanguageIDStatusGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDLanguagesLanguageIDStatusGetExecute(r)
}

/*
V1EnvironmentsEnvIDLanguagesLanguageIDStatusGet READ Language Localization Status

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param languageID
 @return ApiV1EnvironmentsEnvIDLanguagesLanguageIDStatusGetRequest
*/
func (a *ManagementAPIsLanguageManagementLanguageLocalizationStatusApiService) V1EnvironmentsEnvIDLanguagesLanguageIDStatusGet(ctx _context.Context, envID string, languageID string) ApiV1EnvironmentsEnvIDLanguagesLanguageIDStatusGetRequest {
	return ApiV1EnvironmentsEnvIDLanguagesLanguageIDStatusGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		languageID: languageID,
	}
}

// Execute executes the request
func (a *ManagementAPIsLanguageManagementLanguageLocalizationStatusApiService) V1EnvironmentsEnvIDLanguagesLanguageIDStatusGetExecute(r ApiV1EnvironmentsEnvIDLanguagesLanguageIDStatusGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsLanguageManagementLanguageLocalizationStatusApiService.V1EnvironmentsEnvIDLanguagesLanguageIDStatusGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/languages/{languageID}/status"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"languageID"+"}", _neturl.PathEscape(parameterToString(r.languageID, "")), -1)

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

type ApiV1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDDeleteRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsLanguageManagementLanguageLocalizationStatusApiService
	envID string
	languageID string
	l10nStatusID string
}


func (r ApiV1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDDeleteRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDDeleteExecute(r)
}

/*
V1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDDelete DELETE Language Localization Status

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param languageID
 @param l10nStatusID
 @return ApiV1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDDeleteRequest
*/
func (a *ManagementAPIsLanguageManagementLanguageLocalizationStatusApiService) V1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDDelete(ctx _context.Context, envID string, languageID string, l10nStatusID string) ApiV1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDDeleteRequest {
	return ApiV1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDDeleteRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		languageID: languageID,
		l10nStatusID: l10nStatusID,
	}
}

// Execute executes the request
func (a *ManagementAPIsLanguageManagementLanguageLocalizationStatusApiService) V1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDDeleteExecute(r ApiV1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDDeleteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsLanguageManagementLanguageLocalizationStatusApiService.V1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDDelete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/languages/{languageID}/status/{l10nStatusID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"languageID"+"}", _neturl.PathEscape(parameterToString(r.languageID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"l10nStatusID"+"}", _neturl.PathEscape(parameterToString(r.l10nStatusID, "")), -1)

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

type ApiV1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsLanguageManagementLanguageLocalizationStatusApiService
	envID string
	languageID string
	l10nStatusID string
}


func (r ApiV1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDGetExecute(r)
}

/*
V1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDGet READ One Language Localization Status

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param languageID
 @param l10nStatusID
 @return ApiV1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDGetRequest
*/
func (a *ManagementAPIsLanguageManagementLanguageLocalizationStatusApiService) V1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDGet(ctx _context.Context, envID string, languageID string, l10nStatusID string) ApiV1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDGetRequest {
	return ApiV1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		languageID: languageID,
		l10nStatusID: l10nStatusID,
	}
}

// Execute executes the request
func (a *ManagementAPIsLanguageManagementLanguageLocalizationStatusApiService) V1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDGetExecute(r ApiV1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsLanguageManagementLanguageLocalizationStatusApiService.V1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/languages/{languageID}/status/{l10nStatusID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"languageID"+"}", _neturl.PathEscape(parameterToString(r.languageID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"l10nStatusID"+"}", _neturl.PathEscape(parameterToString(r.l10nStatusID, "")), -1)

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

type ApiV1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDPutRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsLanguageManagementLanguageLocalizationStatusApiService
	envID string
	languageID string
	l10nStatusID string
	contentType *string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDPutRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDPutRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDPutRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDPutRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDPutRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDPutExecute(r)
}

/*
V1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDPut CREATE Language Localization Status

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param languageID
 @param l10nStatusID
 @return ApiV1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDPutRequest
*/
func (a *ManagementAPIsLanguageManagementLanguageLocalizationStatusApiService) V1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDPut(ctx _context.Context, envID string, languageID string, l10nStatusID string) ApiV1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDPutRequest {
	return ApiV1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDPutRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		languageID: languageID,
		l10nStatusID: l10nStatusID,
	}
}

// Execute executes the request
func (a *ManagementAPIsLanguageManagementLanguageLocalizationStatusApiService) V1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDPutExecute(r ApiV1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDPutRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsLanguageManagementLanguageLocalizationStatusApiService.V1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDPut")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/languages/{languageID}/status/{l10nStatusID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"languageID"+"}", _neturl.PathEscape(parameterToString(r.languageID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"l10nStatusID"+"}", _neturl.PathEscape(parameterToString(r.l10nStatusID, "")), -1)

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

type ApiV1EnvironmentsEnvIDLanguagesLanguageIDStatusPostRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsLanguageManagementLanguageLocalizationStatusApiService
	envID string
	languageID string
	contentType *string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDLanguagesLanguageIDStatusPostRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDLanguagesLanguageIDStatusPostRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDLanguagesLanguageIDStatusPostRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDLanguagesLanguageIDStatusPostRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDLanguagesLanguageIDStatusPostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDLanguagesLanguageIDStatusPostExecute(r)
}

/*
V1EnvironmentsEnvIDLanguagesLanguageIDStatusPost CREATE Language Localization Status

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param languageID
 @return ApiV1EnvironmentsEnvIDLanguagesLanguageIDStatusPostRequest
*/
func (a *ManagementAPIsLanguageManagementLanguageLocalizationStatusApiService) V1EnvironmentsEnvIDLanguagesLanguageIDStatusPost(ctx _context.Context, envID string, languageID string) ApiV1EnvironmentsEnvIDLanguagesLanguageIDStatusPostRequest {
	return ApiV1EnvironmentsEnvIDLanguagesLanguageIDStatusPostRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		languageID: languageID,
	}
}

// Execute executes the request
func (a *ManagementAPIsLanguageManagementLanguageLocalizationStatusApiService) V1EnvironmentsEnvIDLanguagesLanguageIDStatusPostExecute(r ApiV1EnvironmentsEnvIDLanguagesLanguageIDStatusPostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsLanguageManagementLanguageLocalizationStatusApiService.V1EnvironmentsEnvIDLanguagesLanguageIDStatusPost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/languages/{languageID}/status"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"languageID"+"}", _neturl.PathEscape(parameterToString(r.languageID, "")), -1)

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
