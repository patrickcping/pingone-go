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

// ManagementAPIsBrandingBrandingSettingsApiService ManagementAPIsBrandingBrandingSettingsApi service
type ManagementAPIsBrandingBrandingSettingsApiService service

type ApiV1EnvironmentsEnvIDBrandingSettingsGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsBrandingBrandingSettingsApiService
	envID string
	authorization *string
}

func (r ApiV1EnvironmentsEnvIDBrandingSettingsGetRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDBrandingSettingsGetRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1EnvironmentsEnvIDBrandingSettingsGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDBrandingSettingsGetExecute(r)
}

/*
V1EnvironmentsEnvIDBrandingSettingsGet READ Branding Settings

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @return ApiV1EnvironmentsEnvIDBrandingSettingsGetRequest
*/
func (a *ManagementAPIsBrandingBrandingSettingsApiService) V1EnvironmentsEnvIDBrandingSettingsGet(ctx _context.Context, envID string) ApiV1EnvironmentsEnvIDBrandingSettingsGetRequest {
	return ApiV1EnvironmentsEnvIDBrandingSettingsGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
	}
}

// Execute executes the request
func (a *ManagementAPIsBrandingBrandingSettingsApiService) V1EnvironmentsEnvIDBrandingSettingsGetExecute(r ApiV1EnvironmentsEnvIDBrandingSettingsGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsBrandingBrandingSettingsApiService.V1EnvironmentsEnvIDBrandingSettingsGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/brandingSettings"
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

type ApiV1EnvironmentsEnvIDBrandingSettingsPutRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsBrandingBrandingSettingsApiService
	envID string
	contentType *string
	authorization *string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDBrandingSettingsPutRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDBrandingSettingsPutRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDBrandingSettingsPutRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDBrandingSettingsPutRequest {
	r.authorization = &authorization
	return r
}
func (r ApiV1EnvironmentsEnvIDBrandingSettingsPutRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDBrandingSettingsPutRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDBrandingSettingsPutRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDBrandingSettingsPutExecute(r)
}

/*
V1EnvironmentsEnvIDBrandingSettingsPut UPDATE Branding Settings

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @return ApiV1EnvironmentsEnvIDBrandingSettingsPutRequest
*/
func (a *ManagementAPIsBrandingBrandingSettingsApiService) V1EnvironmentsEnvIDBrandingSettingsPut(ctx _context.Context, envID string) ApiV1EnvironmentsEnvIDBrandingSettingsPutRequest {
	return ApiV1EnvironmentsEnvIDBrandingSettingsPutRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
	}
}

// Execute executes the request
func (a *ManagementAPIsBrandingBrandingSettingsApiService) V1EnvironmentsEnvIDBrandingSettingsPutExecute(r ApiV1EnvironmentsEnvIDBrandingSettingsPutRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsBrandingBrandingSettingsApiService.V1EnvironmentsEnvIDBrandingSettingsPut")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/brandingSettings"
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
