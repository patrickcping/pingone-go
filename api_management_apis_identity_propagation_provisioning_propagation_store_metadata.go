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

// ManagementAPIsIdentityPropagationProvisioningPropagationStoreMetadataApiService ManagementAPIsIdentityPropagationProvisioningPropagationStoreMetadataApi service
type ManagementAPIsIdentityPropagationProvisioningPropagationStoreMetadataApiService service

type ApiV1EnvironmentsEnvIDPropagationStoreMetadataAqueraPostRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsIdentityPropagationProvisioningPropagationStoreMetadataApiService
	envID string
	contentType *string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDPropagationStoreMetadataAqueraPostRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDPropagationStoreMetadataAqueraPostRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDPropagationStoreMetadataAqueraPostRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDPropagationStoreMetadataAqueraPostRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDPropagationStoreMetadataAqueraPostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDPropagationStoreMetadataAqueraPostExecute(r)
}

/*
V1EnvironmentsEnvIDPropagationStoreMetadataAqueraPost Identity Propagation Store Metadata (Aquera)

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @return ApiV1EnvironmentsEnvIDPropagationStoreMetadataAqueraPostRequest
*/
func (a *ManagementAPIsIdentityPropagationProvisioningPropagationStoreMetadataApiService) V1EnvironmentsEnvIDPropagationStoreMetadataAqueraPost(ctx _context.Context, envID string) ApiV1EnvironmentsEnvIDPropagationStoreMetadataAqueraPostRequest {
	return ApiV1EnvironmentsEnvIDPropagationStoreMetadataAqueraPostRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
	}
}

// Execute executes the request
func (a *ManagementAPIsIdentityPropagationProvisioningPropagationStoreMetadataApiService) V1EnvironmentsEnvIDPropagationStoreMetadataAqueraPostExecute(r ApiV1EnvironmentsEnvIDPropagationStoreMetadataAqueraPostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsIdentityPropagationProvisioningPropagationStoreMetadataApiService.V1EnvironmentsEnvIDPropagationStoreMetadataAqueraPost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/propagation/storeMetadata/Aquera"
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

type ApiV1EnvironmentsEnvIDPropagationStoreMetadataSalesforceContactsPostRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsIdentityPropagationProvisioningPropagationStoreMetadataApiService
	envID string
	contentType *string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDPropagationStoreMetadataSalesforceContactsPostRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDPropagationStoreMetadataSalesforceContactsPostRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDPropagationStoreMetadataSalesforceContactsPostRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDPropagationStoreMetadataSalesforceContactsPostRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDPropagationStoreMetadataSalesforceContactsPostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDPropagationStoreMetadataSalesforceContactsPostExecute(r)
}

/*
V1EnvironmentsEnvIDPropagationStoreMetadataSalesforceContactsPost Identity Propagation Store Metadata (SalesforceContacts)

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @return ApiV1EnvironmentsEnvIDPropagationStoreMetadataSalesforceContactsPostRequest
*/
func (a *ManagementAPIsIdentityPropagationProvisioningPropagationStoreMetadataApiService) V1EnvironmentsEnvIDPropagationStoreMetadataSalesforceContactsPost(ctx _context.Context, envID string) ApiV1EnvironmentsEnvIDPropagationStoreMetadataSalesforceContactsPostRequest {
	return ApiV1EnvironmentsEnvIDPropagationStoreMetadataSalesforceContactsPostRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
	}
}

// Execute executes the request
func (a *ManagementAPIsIdentityPropagationProvisioningPropagationStoreMetadataApiService) V1EnvironmentsEnvIDPropagationStoreMetadataSalesforceContactsPostExecute(r ApiV1EnvironmentsEnvIDPropagationStoreMetadataSalesforceContactsPostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsIdentityPropagationProvisioningPropagationStoreMetadataApiService.V1EnvironmentsEnvIDPropagationStoreMetadataSalesforceContactsPost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/propagation/storeMetadata/SalesforceContacts"
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

type ApiV1EnvironmentsEnvIDPropagationStoreMetadataSalesforcePostRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsIdentityPropagationProvisioningPropagationStoreMetadataApiService
	envID string
	contentType *string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDPropagationStoreMetadataSalesforcePostRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDPropagationStoreMetadataSalesforcePostRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDPropagationStoreMetadataSalesforcePostRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDPropagationStoreMetadataSalesforcePostRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDPropagationStoreMetadataSalesforcePostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDPropagationStoreMetadataSalesforcePostExecute(r)
}

/*
V1EnvironmentsEnvIDPropagationStoreMetadataSalesforcePost Identity Propagation Store Metadata (Salesforce)

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @return ApiV1EnvironmentsEnvIDPropagationStoreMetadataSalesforcePostRequest
*/
func (a *ManagementAPIsIdentityPropagationProvisioningPropagationStoreMetadataApiService) V1EnvironmentsEnvIDPropagationStoreMetadataSalesforcePost(ctx _context.Context, envID string) ApiV1EnvironmentsEnvIDPropagationStoreMetadataSalesforcePostRequest {
	return ApiV1EnvironmentsEnvIDPropagationStoreMetadataSalesforcePostRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
	}
}

// Execute executes the request
func (a *ManagementAPIsIdentityPropagationProvisioningPropagationStoreMetadataApiService) V1EnvironmentsEnvIDPropagationStoreMetadataSalesforcePostExecute(r ApiV1EnvironmentsEnvIDPropagationStoreMetadataSalesforcePostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsIdentityPropagationProvisioningPropagationStoreMetadataApiService.V1EnvironmentsEnvIDPropagationStoreMetadataSalesforcePost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/propagation/storeMetadata/Salesforce"
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

type ApiV1EnvironmentsEnvIDPropagationStoreMetadataScimPostRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsIdentityPropagationProvisioningPropagationStoreMetadataApiService
	envID string
	contentType *string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDPropagationStoreMetadataScimPostRequest) ContentType(contentType string) ApiV1EnvironmentsEnvIDPropagationStoreMetadataScimPostRequest {
	r.contentType = &contentType
	return r
}
func (r ApiV1EnvironmentsEnvIDPropagationStoreMetadataScimPostRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDPropagationStoreMetadataScimPostRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDPropagationStoreMetadataScimPostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDPropagationStoreMetadataScimPostExecute(r)
}

/*
V1EnvironmentsEnvIDPropagationStoreMetadataScimPost Identity Propagation Store Metadata (SCIM)

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @return ApiV1EnvironmentsEnvIDPropagationStoreMetadataScimPostRequest
*/
func (a *ManagementAPIsIdentityPropagationProvisioningPropagationStoreMetadataApiService) V1EnvironmentsEnvIDPropagationStoreMetadataScimPost(ctx _context.Context, envID string) ApiV1EnvironmentsEnvIDPropagationStoreMetadataScimPostRequest {
	return ApiV1EnvironmentsEnvIDPropagationStoreMetadataScimPostRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
	}
}

// Execute executes the request
func (a *ManagementAPIsIdentityPropagationProvisioningPropagationStoreMetadataApiService) V1EnvironmentsEnvIDPropagationStoreMetadataScimPostExecute(r ApiV1EnvironmentsEnvIDPropagationStoreMetadataScimPostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsIdentityPropagationProvisioningPropagationStoreMetadataApiService.V1EnvironmentsEnvIDPropagationStoreMetadataScimPost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/propagation/storeMetadata/scim"
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
