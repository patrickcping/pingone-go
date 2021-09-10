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

// ManagementAPIsGatewayManagementGatewayRoleAssignmentsApiService ManagementAPIsGatewayManagementGatewayRoleAssignmentsApi service
type ManagementAPIsGatewayManagementGatewayRoleAssignmentsApiService service

type ApiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDDeleteRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsGatewayManagementGatewayRoleAssignmentsApiService
	envID string
	gatewayID string
	gatewayRoleAssignmentsID string
}


func (r ApiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDDeleteRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDDeleteExecute(r)
}

/*
V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDDelete DELETE Gateway Role Assignment

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param gatewayID
 @param gatewayRoleAssignmentsID
 @return ApiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDDeleteRequest
*/
func (a *ManagementAPIsGatewayManagementGatewayRoleAssignmentsApiService) V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDDelete(ctx _context.Context, envID string, gatewayID string, gatewayRoleAssignmentsID string) ApiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDDeleteRequest {
	return ApiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDDeleteRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		gatewayID: gatewayID,
		gatewayRoleAssignmentsID: gatewayRoleAssignmentsID,
	}
}

// Execute executes the request
func (a *ManagementAPIsGatewayManagementGatewayRoleAssignmentsApiService) V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDDeleteExecute(r ApiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDDeleteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsGatewayManagementGatewayRoleAssignmentsApiService.V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDDelete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/gateways/{gatewayID}/roleAssignments/{gatewayRoleAssignmentsID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"gatewayID"+"}", _neturl.PathEscape(parameterToString(r.gatewayID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"gatewayRoleAssignmentsID"+"}", _neturl.PathEscape(parameterToString(r.gatewayRoleAssignmentsID, "")), -1)

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

type ApiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsGatewayManagementGatewayRoleAssignmentsApiService
	envID string
	gatewayID string
	gatewayRoleAssignmentsID string
}


func (r ApiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDGetExecute(r)
}

/*
V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDGet READ One Gateway Role Assignment

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param gatewayID
 @param gatewayRoleAssignmentsID
 @return ApiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDGetRequest
*/
func (a *ManagementAPIsGatewayManagementGatewayRoleAssignmentsApiService) V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDGet(ctx _context.Context, envID string, gatewayID string, gatewayRoleAssignmentsID string) ApiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDGetRequest {
	return ApiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		gatewayID: gatewayID,
		gatewayRoleAssignmentsID: gatewayRoleAssignmentsID,
	}
}

// Execute executes the request
func (a *ManagementAPIsGatewayManagementGatewayRoleAssignmentsApiService) V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDGetExecute(r ApiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsGatewayManagementGatewayRoleAssignmentsApiService.V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/gateways/{gatewayID}/roleAssignments/{gatewayRoleAssignmentsID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"gatewayID"+"}", _neturl.PathEscape(parameterToString(r.gatewayID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"gatewayRoleAssignmentsID"+"}", _neturl.PathEscape(parameterToString(r.gatewayRoleAssignmentsID, "")), -1)

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

type ApiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDPutRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsGatewayManagementGatewayRoleAssignmentsApiService
	envID string
	gatewayID string
	gatewayRoleAssignmentsID string
	contentType *string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDPutRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDPutRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDPutRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDPutRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDPutRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDPutExecute(r)
}

/*
V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDPut UPDATE Gateway Role Assignments

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param gatewayID
 @param gatewayRoleAssignmentsID
 @return ApiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDPutRequest
*/
func (a *ManagementAPIsGatewayManagementGatewayRoleAssignmentsApiService) V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDPut(ctx _context.Context, envID string, gatewayID string, gatewayRoleAssignmentsID string) ApiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDPutRequest {
	return ApiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDPutRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		gatewayID: gatewayID,
		gatewayRoleAssignmentsID: gatewayRoleAssignmentsID,
	}
}

// Execute executes the request
func (a *ManagementAPIsGatewayManagementGatewayRoleAssignmentsApiService) V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDPutExecute(r ApiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDPutRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsGatewayManagementGatewayRoleAssignmentsApiService.V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDPut")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/gateways/{gatewayID}/roleAssignments/{gatewayRoleAssignmentsID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"gatewayID"+"}", _neturl.PathEscape(parameterToString(r.gatewayID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"gatewayRoleAssignmentsID"+"}", _neturl.PathEscape(parameterToString(r.gatewayRoleAssignmentsID, "")), -1)

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

type ApiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsGatewayManagementGatewayRoleAssignmentsApiService
	envID string
	gatewayID string
}


func (r ApiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGetExecute(r)
}

/*
V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGet READ Gateway Role Assignments

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param gatewayID
 @return ApiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGetRequest
*/
func (a *ManagementAPIsGatewayManagementGatewayRoleAssignmentsApiService) V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGet(ctx _context.Context, envID string, gatewayID string) ApiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGetRequest {
	return ApiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		gatewayID: gatewayID,
	}
}

// Execute executes the request
func (a *ManagementAPIsGatewayManagementGatewayRoleAssignmentsApiService) V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGetExecute(r ApiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsGatewayManagementGatewayRoleAssignmentsApiService.V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/gateways/{gatewayID}/roleAssignments"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"gatewayID"+"}", _neturl.PathEscape(parameterToString(r.gatewayID, "")), -1)

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

type ApiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsPostRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsGatewayManagementGatewayRoleAssignmentsApiService
	envID string
	gatewayID string
	contentType *string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsPostRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsPostRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsPostRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsPostRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsPostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsPostExecute(r)
}

/*
V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsPost CREATE Gateway Role Assignments

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param gatewayID
 @return ApiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsPostRequest
*/
func (a *ManagementAPIsGatewayManagementGatewayRoleAssignmentsApiService) V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsPost(ctx _context.Context, envID string, gatewayID string) ApiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsPostRequest {
	return ApiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsPostRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		gatewayID: gatewayID,
	}
}

// Execute executes the request
func (a *ManagementAPIsGatewayManagementGatewayRoleAssignmentsApiService) V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsPostExecute(r ApiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsPostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsGatewayManagementGatewayRoleAssignmentsApiService.V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsPost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/gateways/{gatewayID}/roleAssignments"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"gatewayID"+"}", _neturl.PathEscape(parameterToString(r.gatewayID, "")), -1)

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
