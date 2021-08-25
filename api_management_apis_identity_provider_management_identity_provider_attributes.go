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

// ManagementAPIsIdentityProviderManagementIdentityProviderAttributesApiService ManagementAPIsIdentityProviderManagementIdentityProviderAttributesApi service
type ManagementAPIsIdentityProviderManagementIdentityProviderAttributesApiService service

type ApiV11EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDDeleteRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsIdentityProviderManagementIdentityProviderAttributesApiService
	envID string
	providerID string
	idpAttrID string
	authorization *string
}

func (r ApiV11EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDDeleteRequest) Authorization(authorization string) ApiV11EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDDeleteRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV11EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDDeleteRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V11EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDDeleteExecute(r)
}

/*
 * V11EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDDelete DELETE Identity Provider Attribute
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @param providerID
 * @param idpAttrID
 * @return ApiV11EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDDeleteRequest
 */
func (a *ManagementAPIsIdentityProviderManagementIdentityProviderAttributesApiService) V11EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDDelete(ctx _context.Context, envID string, providerID string, idpAttrID string) ApiV11EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDDeleteRequest {
	return ApiV11EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDDeleteRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		providerID: providerID,
		idpAttrID: idpAttrID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsIdentityProviderManagementIdentityProviderAttributesApiService) V11EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDDeleteExecute(r ApiV11EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDDeleteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsIdentityProviderManagementIdentityProviderAttributesApiService.V11EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDDelete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v11/environments/{envID}/identityProviders/{providerID}/attributes/{idpAttrID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"providerID"+"}", _neturl.PathEscape(parameterToString(r.providerID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"idpAttrID"+"}", _neturl.PathEscape(parameterToString(r.idpAttrID, "")), -1)

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

type ApiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsIdentityProviderManagementIdentityProviderAttributesApiService
	envID string
	providerID string
	authorization *string
}

func (r ApiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesGetRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesGetRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesGetExecute(r)
}

/*
 * V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesGet READ All Identity Provider Attributes
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @param providerID
 * @return ApiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesGetRequest
 */
func (a *ManagementAPIsIdentityProviderManagementIdentityProviderAttributesApiService) V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesGet(ctx _context.Context, envID string, providerID string) ApiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesGetRequest {
	return ApiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		providerID: providerID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsIdentityProviderManagementIdentityProviderAttributesApiService) V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesGetExecute(r ApiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsIdentityProviderManagementIdentityProviderAttributesApiService.V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/identityProviders/{providerID}/attributes"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"providerID"+"}", _neturl.PathEscape(parameterToString(r.providerID, "")), -1)

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

type ApiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsIdentityProviderManagementIdentityProviderAttributesApiService
	envID string
	providerID string
	idpAttrID string
	authorization *string
}

func (r ApiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDGetRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDGetRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDGetExecute(r)
}

/*
 * V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDGet READ One Identity Provider Attribute
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @param providerID
 * @param idpAttrID
 * @return ApiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDGetRequest
 */
func (a *ManagementAPIsIdentityProviderManagementIdentityProviderAttributesApiService) V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDGet(ctx _context.Context, envID string, providerID string, idpAttrID string) ApiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDGetRequest {
	return ApiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		providerID: providerID,
		idpAttrID: idpAttrID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsIdentityProviderManagementIdentityProviderAttributesApiService) V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDGetExecute(r ApiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsIdentityProviderManagementIdentityProviderAttributesApiService.V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/identityProviders/{providerID}/attributes/{idpAttrID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"providerID"+"}", _neturl.PathEscape(parameterToString(r.providerID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"idpAttrID"+"}", _neturl.PathEscape(parameterToString(r.idpAttrID, "")), -1)

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

type ApiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDPutRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsIdentityProviderManagementIdentityProviderAttributesApiService
	envID string
	providerID string
	idpAttrID string
	contentType *string
	authorization *string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDPutRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDPutRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDPutRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDPutRequest {
	r.authorization = &authorization
	return r
}
func (r ApiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDPutRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDPutRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDPutRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDPutExecute(r)
}

/*
 * V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDPut UPDATE Identity Provider Attribute
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @param providerID
 * @param idpAttrID
 * @return ApiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDPutRequest
 */
func (a *ManagementAPIsIdentityProviderManagementIdentityProviderAttributesApiService) V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDPut(ctx _context.Context, envID string, providerID string, idpAttrID string) ApiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDPutRequest {
	return ApiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDPutRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		providerID: providerID,
		idpAttrID: idpAttrID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsIdentityProviderManagementIdentityProviderAttributesApiService) V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDPutExecute(r ApiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDPutRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsIdentityProviderManagementIdentityProviderAttributesApiService.V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDPut")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/identityProviders/{providerID}/attributes/{idpAttrID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"providerID"+"}", _neturl.PathEscape(parameterToString(r.providerID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"idpAttrID"+"}", _neturl.PathEscape(parameterToString(r.idpAttrID, "")), -1)

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

type ApiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesPostRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsIdentityProviderManagementIdentityProviderAttributesApiService
	envID string
	providerID string
	contentType *string
	authorization *string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesPostRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesPostRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesPostRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesPostRequest {
	r.authorization = &authorization
	return r
}
func (r ApiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesPostRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesPostRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesPostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesPostExecute(r)
}

/*
 * V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesPost CREATE Identity Provider Attribute (SAML)
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @param providerID
 * @return ApiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesPostRequest
 */
func (a *ManagementAPIsIdentityProviderManagementIdentityProviderAttributesApiService) V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesPost(ctx _context.Context, envID string, providerID string) ApiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesPostRequest {
	return ApiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesPostRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		providerID: providerID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsIdentityProviderManagementIdentityProviderAttributesApiService) V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesPostExecute(r ApiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesPostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsIdentityProviderManagementIdentityProviderAttributesApiService.V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesPost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/identityProviders/{providerID}/attributes"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"providerID"+"}", _neturl.PathEscape(parameterToString(r.providerID, "")), -1)

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
