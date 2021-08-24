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

// ManagementAPIsBrandingBrandingThemesApiService ManagementAPIsBrandingBrandingThemesApi service
type ManagementAPIsBrandingBrandingThemesApiService service

type ApiV1EnvironmentsEnvIDThemesGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsBrandingBrandingThemesApiService
	envID string
	authorization *string
}

func (r ApiV1EnvironmentsEnvIDThemesGetRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDThemesGetRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1EnvironmentsEnvIDThemesGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDThemesGetExecute(r)
}

/*
 * V1EnvironmentsEnvIDThemesGet READ Branding Themes
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @return ApiV1EnvironmentsEnvIDThemesGetRequest
 */
func (a *ManagementAPIsBrandingBrandingThemesApiService) V1EnvironmentsEnvIDThemesGet(ctx _context.Context, envID string) ApiV1EnvironmentsEnvIDThemesGetRequest {
	return ApiV1EnvironmentsEnvIDThemesGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsBrandingBrandingThemesApiService) V1EnvironmentsEnvIDThemesGetExecute(r ApiV1EnvironmentsEnvIDThemesGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsBrandingBrandingThemesApiService.V1EnvironmentsEnvIDThemesGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/themes"
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

type ApiV1EnvironmentsEnvIDThemesPostRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsBrandingBrandingThemesApiService
	envID string
	authorization *string
	contentType *string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDThemesPostRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDThemesPostRequest {
	r.authorization = &authorization
	return r
}
func (r ApiV1EnvironmentsEnvIDThemesPostRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDThemesPostRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDThemesPostRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDThemesPostRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDThemesPostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDThemesPostExecute(r)
}

/*
 * V1EnvironmentsEnvIDThemesPost CREATE Branding Theme
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @return ApiV1EnvironmentsEnvIDThemesPostRequest
 */
func (a *ManagementAPIsBrandingBrandingThemesApiService) V1EnvironmentsEnvIDThemesPost(ctx _context.Context, envID string) ApiV1EnvironmentsEnvIDThemesPostRequest {
	return ApiV1EnvironmentsEnvIDThemesPostRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsBrandingBrandingThemesApiService) V1EnvironmentsEnvIDThemesPostExecute(r ApiV1EnvironmentsEnvIDThemesPostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsBrandingBrandingThemesApiService.V1EnvironmentsEnvIDThemesPost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/themes"
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
	if r.authorization != nil {
		localVarHeaderParams["Authorization"] = parameterToString(*r.authorization, "")
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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiV1EnvironmentsEnvIDThemesThemeIDDefaultGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsBrandingBrandingThemesApiService
	envID string
	themeID string
	authorization *string
}

func (r ApiV1EnvironmentsEnvIDThemesThemeIDDefaultGetRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDThemesThemeIDDefaultGetRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1EnvironmentsEnvIDThemesThemeIDDefaultGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDThemesThemeIDDefaultGetExecute(r)
}

/*
 * V1EnvironmentsEnvIDThemesThemeIDDefaultGet READ Branding Theme Default
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @param themeID
 * @return ApiV1EnvironmentsEnvIDThemesThemeIDDefaultGetRequest
 */
func (a *ManagementAPIsBrandingBrandingThemesApiService) V1EnvironmentsEnvIDThemesThemeIDDefaultGet(ctx _context.Context, envID string, themeID string) ApiV1EnvironmentsEnvIDThemesThemeIDDefaultGetRequest {
	return ApiV1EnvironmentsEnvIDThemesThemeIDDefaultGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		themeID: themeID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsBrandingBrandingThemesApiService) V1EnvironmentsEnvIDThemesThemeIDDefaultGetExecute(r ApiV1EnvironmentsEnvIDThemesThemeIDDefaultGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsBrandingBrandingThemesApiService.V1EnvironmentsEnvIDThemesThemeIDDefaultGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/themes/{themeID}/default"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"themeID"+"}", _neturl.PathEscape(parameterToString(r.themeID, "")), -1)

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

type ApiV1EnvironmentsEnvIDThemesThemeIDDefaultPutRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsBrandingBrandingThemesApiService
	envID string
	themeID string
	authorization *string
	contentType *string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDThemesThemeIDDefaultPutRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDThemesThemeIDDefaultPutRequest {
	r.authorization = &authorization
	return r
}
func (r ApiV1EnvironmentsEnvIDThemesThemeIDDefaultPutRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDThemesThemeIDDefaultPutRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDThemesThemeIDDefaultPutRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDThemesThemeIDDefaultPutRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDThemesThemeIDDefaultPutRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDThemesThemeIDDefaultPutExecute(r)
}

/*
 * V1EnvironmentsEnvIDThemesThemeIDDefaultPut UPDATE Branding Theme Default
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @param themeID
 * @return ApiV1EnvironmentsEnvIDThemesThemeIDDefaultPutRequest
 */
func (a *ManagementAPIsBrandingBrandingThemesApiService) V1EnvironmentsEnvIDThemesThemeIDDefaultPut(ctx _context.Context, envID string, themeID string) ApiV1EnvironmentsEnvIDThemesThemeIDDefaultPutRequest {
	return ApiV1EnvironmentsEnvIDThemesThemeIDDefaultPutRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		themeID: themeID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsBrandingBrandingThemesApiService) V1EnvironmentsEnvIDThemesThemeIDDefaultPutExecute(r ApiV1EnvironmentsEnvIDThemesThemeIDDefaultPutRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsBrandingBrandingThemesApiService.V1EnvironmentsEnvIDThemesThemeIDDefaultPut")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/themes/{themeID}/default"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"themeID"+"}", _neturl.PathEscape(parameterToString(r.themeID, "")), -1)

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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiV1EnvironmentsEnvIDThemesThemeIDDeleteRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsBrandingBrandingThemesApiService
	envID string
	themeID string
	authorization *string
}

func (r ApiV1EnvironmentsEnvIDThemesThemeIDDeleteRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDThemesThemeIDDeleteRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1EnvironmentsEnvIDThemesThemeIDDeleteRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDThemesThemeIDDeleteExecute(r)
}

/*
 * V1EnvironmentsEnvIDThemesThemeIDDelete DELETE Branding Theme
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @param themeID
 * @return ApiV1EnvironmentsEnvIDThemesThemeIDDeleteRequest
 */
func (a *ManagementAPIsBrandingBrandingThemesApiService) V1EnvironmentsEnvIDThemesThemeIDDelete(ctx _context.Context, envID string, themeID string) ApiV1EnvironmentsEnvIDThemesThemeIDDeleteRequest {
	return ApiV1EnvironmentsEnvIDThemesThemeIDDeleteRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		themeID: themeID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsBrandingBrandingThemesApiService) V1EnvironmentsEnvIDThemesThemeIDDeleteExecute(r ApiV1EnvironmentsEnvIDThemesThemeIDDeleteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsBrandingBrandingThemesApiService.V1EnvironmentsEnvIDThemesThemeIDDelete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/themes/{themeID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"themeID"+"}", _neturl.PathEscape(parameterToString(r.themeID, "")), -1)

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

type ApiV1EnvironmentsEnvIDThemesThemeIDGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsBrandingBrandingThemesApiService
	envID string
	themeID string
	authorization *string
}

func (r ApiV1EnvironmentsEnvIDThemesThemeIDGetRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDThemesThemeIDGetRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1EnvironmentsEnvIDThemesThemeIDGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDThemesThemeIDGetExecute(r)
}

/*
 * V1EnvironmentsEnvIDThemesThemeIDGet READ One Branding Theme
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @param themeID
 * @return ApiV1EnvironmentsEnvIDThemesThemeIDGetRequest
 */
func (a *ManagementAPIsBrandingBrandingThemesApiService) V1EnvironmentsEnvIDThemesThemeIDGet(ctx _context.Context, envID string, themeID string) ApiV1EnvironmentsEnvIDThemesThemeIDGetRequest {
	return ApiV1EnvironmentsEnvIDThemesThemeIDGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		themeID: themeID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsBrandingBrandingThemesApiService) V1EnvironmentsEnvIDThemesThemeIDGetExecute(r ApiV1EnvironmentsEnvIDThemesThemeIDGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsBrandingBrandingThemesApiService.V1EnvironmentsEnvIDThemesThemeIDGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/themes/{themeID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"themeID"+"}", _neturl.PathEscape(parameterToString(r.themeID, "")), -1)

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

type ApiV1EnvironmentsEnvIDThemesThemeIDPutRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsBrandingBrandingThemesApiService
	envID string
	themeID string
	authorization *string
	contentType *string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDThemesThemeIDPutRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDThemesThemeIDPutRequest {
	r.authorization = &authorization
	return r
}
func (r ApiV1EnvironmentsEnvIDThemesThemeIDPutRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDThemesThemeIDPutRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDThemesThemeIDPutRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDThemesThemeIDPutRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDThemesThemeIDPutRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDThemesThemeIDPutExecute(r)
}

/*
 * V1EnvironmentsEnvIDThemesThemeIDPut UPDATE Branding Theme
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @param themeID
 * @return ApiV1EnvironmentsEnvIDThemesThemeIDPutRequest
 */
func (a *ManagementAPIsBrandingBrandingThemesApiService) V1EnvironmentsEnvIDThemesThemeIDPut(ctx _context.Context, envID string, themeID string) ApiV1EnvironmentsEnvIDThemesThemeIDPutRequest {
	return ApiV1EnvironmentsEnvIDThemesThemeIDPutRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		themeID: themeID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsBrandingBrandingThemesApiService) V1EnvironmentsEnvIDThemesThemeIDPutExecute(r ApiV1EnvironmentsEnvIDThemesThemeIDPutRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsBrandingBrandingThemesApiService.V1EnvironmentsEnvIDThemesThemeIDPut")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/themes/{themeID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"themeID"+"}", _neturl.PathEscape(parameterToString(r.themeID, "")), -1)

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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
