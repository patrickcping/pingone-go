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

// ManagementAPIsIdentityPropagationProvisioningPropagationRevisionsApiService ManagementAPIsIdentityPropagationProvisioningPropagationRevisionsApi service
type ManagementAPIsIdentityPropagationProvisioningPropagationRevisionsApiService service

type ApiV1EnvironmentsEnvIDPropagationRevisionsIdlatestGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsIdentityPropagationProvisioningPropagationRevisionsApiService
	envID string
	accept *string
	authorization *string
}

func (r ApiV1EnvironmentsEnvIDPropagationRevisionsIdlatestGetRequest) Accept(accept string) ApiV1EnvironmentsEnvIDPropagationRevisionsIdlatestGetRequest {
	r.accept = &accept
	return r
}
func (r ApiV1EnvironmentsEnvIDPropagationRevisionsIdlatestGetRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDPropagationRevisionsIdlatestGetRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1EnvironmentsEnvIDPropagationRevisionsIdlatestGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDPropagationRevisionsIdlatestGetExecute(r)
}

/*
 * V1EnvironmentsEnvIDPropagationRevisionsIdlatestGet READ Latest Revision
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @return ApiV1EnvironmentsEnvIDPropagationRevisionsIdlatestGetRequest
 */
func (a *ManagementAPIsIdentityPropagationProvisioningPropagationRevisionsApiService) V1EnvironmentsEnvIDPropagationRevisionsIdlatestGet(ctx _context.Context, envID string) ApiV1EnvironmentsEnvIDPropagationRevisionsIdlatestGetRequest {
	return ApiV1EnvironmentsEnvIDPropagationRevisionsIdlatestGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsIdentityPropagationProvisioningPropagationRevisionsApiService) V1EnvironmentsEnvIDPropagationRevisionsIdlatestGetExecute(r ApiV1EnvironmentsEnvIDPropagationRevisionsIdlatestGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsIdentityPropagationProvisioningPropagationRevisionsApiService.V1EnvironmentsEnvIDPropagationRevisionsIdlatestGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/propagation/revisions/id:latest"
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
	if r.accept != nil {
		localVarHeaderParams["Accept"] = parameterToString(*r.accept, "")
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

type ApiV1EnvironmentsEnvIDPropagationRevisionsPostRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsIdentityPropagationProvisioningPropagationRevisionsApiService
	envID string
	contentType *string
	authorization *string
}

func (r ApiV1EnvironmentsEnvIDPropagationRevisionsPostRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDPropagationRevisionsPostRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDPropagationRevisionsPostRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDPropagationRevisionsPostRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1EnvironmentsEnvIDPropagationRevisionsPostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDPropagationRevisionsPostExecute(r)
}

/*
 * V1EnvironmentsEnvIDPropagationRevisionsPost CREATE Revision
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @return ApiV1EnvironmentsEnvIDPropagationRevisionsPostRequest
 */
func (a *ManagementAPIsIdentityPropagationProvisioningPropagationRevisionsApiService) V1EnvironmentsEnvIDPropagationRevisionsPost(ctx _context.Context, envID string) ApiV1EnvironmentsEnvIDPropagationRevisionsPostRequest {
	return ApiV1EnvironmentsEnvIDPropagationRevisionsPostRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsIdentityPropagationProvisioningPropagationRevisionsApiService) V1EnvironmentsEnvIDPropagationRevisionsPostExecute(r ApiV1EnvironmentsEnvIDPropagationRevisionsPostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsIdentityPropagationProvisioningPropagationRevisionsApiService.V1EnvironmentsEnvIDPropagationRevisionsPost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/propagation/revisions"
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

type ApiV1EnvironmentsEnvIDPropagationRevisionsPreviousRevisionIDGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsIdentityPropagationProvisioningPropagationRevisionsApiService
	envID string
	previousRevisionID string
	accept *string
	authorization *string
}

func (r ApiV1EnvironmentsEnvIDPropagationRevisionsPreviousRevisionIDGetRequest) Accept(accept string) ApiV1EnvironmentsEnvIDPropagationRevisionsPreviousRevisionIDGetRequest {
	r.accept = &accept
	return r
}
func (r ApiV1EnvironmentsEnvIDPropagationRevisionsPreviousRevisionIDGetRequest) Authorization(authorization string) ApiV1EnvironmentsEnvIDPropagationRevisionsPreviousRevisionIDGetRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1EnvironmentsEnvIDPropagationRevisionsPreviousRevisionIDGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDPropagationRevisionsPreviousRevisionIDGetExecute(r)
}

/*
 * V1EnvironmentsEnvIDPropagationRevisionsPreviousRevisionIDGet READ Previous Revision
 * By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envID
 * @param previousRevisionID
 * @return ApiV1EnvironmentsEnvIDPropagationRevisionsPreviousRevisionIDGetRequest
 */
func (a *ManagementAPIsIdentityPropagationProvisioningPropagationRevisionsApiService) V1EnvironmentsEnvIDPropagationRevisionsPreviousRevisionIDGet(ctx _context.Context, envID string, previousRevisionID string) ApiV1EnvironmentsEnvIDPropagationRevisionsPreviousRevisionIDGetRequest {
	return ApiV1EnvironmentsEnvIDPropagationRevisionsPreviousRevisionIDGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		previousRevisionID: previousRevisionID,
	}
}

/*
 * Execute executes the request
 */
func (a *ManagementAPIsIdentityPropagationProvisioningPropagationRevisionsApiService) V1EnvironmentsEnvIDPropagationRevisionsPreviousRevisionIDGetExecute(r ApiV1EnvironmentsEnvIDPropagationRevisionsPreviousRevisionIDGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsIdentityPropagationProvisioningPropagationRevisionsApiService.V1EnvironmentsEnvIDPropagationRevisionsPreviousRevisionIDGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/propagation/revisions/{previousRevisionID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"previousRevisionID"+"}", _neturl.PathEscape(parameterToString(r.previousRevisionID, "")), -1)

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
	if r.accept != nil {
		localVarHeaderParams["Accept"] = parameterToString(*r.accept, "")
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
