/*
PingOne Platform API - Management

A bare-bones collection for the PingOne API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pingone

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// ManagementAPIsAgreementManagementAgreementLanguagesResourcesApiService ManagementAPIsAgreementManagementAgreementLanguagesResourcesApi service
type ManagementAPIsAgreementManagementAgreementLanguagesResourcesApiService service

type ApiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesGetRequest struct {
	ctx context.Context
	ApiService *ManagementAPIsAgreementManagementAgreementLanguagesResourcesApiService
	envID string
	agreementID string
}

func (r ApiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesGetExecute(r)
}

/*
V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesGet READ All Languages

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param agreementID
 @return ApiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesGetRequest
*/
func (a *ManagementAPIsAgreementManagementAgreementLanguagesResourcesApiService) V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesGet(ctx context.Context, envID string, agreementID string) ApiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesGetRequest {
	return ApiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		agreementID: agreementID,
	}
}

// Execute executes the request
func (a *ManagementAPIsAgreementManagementAgreementLanguagesResourcesApiService) V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesGetExecute(r ApiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsAgreementManagementAgreementLanguagesResourcesApiService.V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/agreements/{agreementID}/languages"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", url.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"agreementID"+"}", url.PathEscape(parameterToString(r.agreementID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
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

type ApiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDDeleteRequest struct {
	ctx context.Context
	ApiService *ManagementAPIsAgreementManagementAgreementLanguagesResourcesApiService
	envID string
	agreementID string
	languageID string
}

func (r ApiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDDeleteExecute(r)
}

/*
V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDDelete DELETE Language

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param agreementID
 @param languageID
 @return ApiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDDeleteRequest
*/
func (a *ManagementAPIsAgreementManagementAgreementLanguagesResourcesApiService) V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDDelete(ctx context.Context, envID string, agreementID string, languageID string) ApiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDDeleteRequest {
	return ApiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDDeleteRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		agreementID: agreementID,
		languageID: languageID,
	}
}

// Execute executes the request
func (a *ManagementAPIsAgreementManagementAgreementLanguagesResourcesApiService) V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDDeleteExecute(r ApiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsAgreementManagementAgreementLanguagesResourcesApiService.V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/agreements/{agreementID}/languages/{languageID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", url.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"agreementID"+"}", url.PathEscape(parameterToString(r.agreementID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"languageID"+"}", url.PathEscape(parameterToString(r.languageID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
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

type ApiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDGetRequest struct {
	ctx context.Context
	ApiService *ManagementAPIsAgreementManagementAgreementLanguagesResourcesApiService
	envID string
	agreementID string
	languageID string
}

func (r ApiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDGetExecute(r)
}

/*
V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDGet READ One Language

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param agreementID
 @param languageID
 @return ApiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDGetRequest
*/
func (a *ManagementAPIsAgreementManagementAgreementLanguagesResourcesApiService) V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDGet(ctx context.Context, envID string, agreementID string, languageID string) ApiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDGetRequest {
	return ApiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		agreementID: agreementID,
		languageID: languageID,
	}
}

// Execute executes the request
func (a *ManagementAPIsAgreementManagementAgreementLanguagesResourcesApiService) V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDGetExecute(r ApiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsAgreementManagementAgreementLanguagesResourcesApiService.V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/agreements/{agreementID}/languages/{languageID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", url.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"agreementID"+"}", url.PathEscape(parameterToString(r.agreementID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"languageID"+"}", url.PathEscape(parameterToString(r.languageID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
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

type ApiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDPutRequest struct {
	ctx context.Context
	ApiService *ManagementAPIsAgreementManagementAgreementLanguagesResourcesApiService
	envID string
	agreementID string
	languageID string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDPutRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDPutRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDPutRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDPutExecute(r)
}

/*
V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDPut UPDATE Language

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param agreementID
 @param languageID
 @return ApiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDPutRequest
*/
func (a *ManagementAPIsAgreementManagementAgreementLanguagesResourcesApiService) V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDPut(ctx context.Context, envID string, agreementID string, languageID string) ApiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDPutRequest {
	return ApiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDPutRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		agreementID: agreementID,
		languageID: languageID,
	}
}

// Execute executes the request
func (a *ManagementAPIsAgreementManagementAgreementLanguagesResourcesApiService) V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDPutExecute(r ApiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDPutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsAgreementManagementAgreementLanguagesResourcesApiService.V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDPut")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/agreements/{agreementID}/languages/{languageID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", url.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"agreementID"+"}", url.PathEscape(parameterToString(r.agreementID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"languageID"+"}", url.PathEscape(parameterToString(r.languageID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
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

type ApiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesPostRequest struct {
	ctx context.Context
	ApiService *ManagementAPIsAgreementManagementAgreementLanguagesResourcesApiService
	envID string
	agreementID string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesPostRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesPostRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesPostExecute(r)
}

/*
V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesPost CREATE Language

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param agreementID
 @return ApiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesPostRequest
*/
func (a *ManagementAPIsAgreementManagementAgreementLanguagesResourcesApiService) V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesPost(ctx context.Context, envID string, agreementID string) ApiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesPostRequest {
	return ApiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesPostRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		agreementID: agreementID,
	}
}

// Execute executes the request
func (a *ManagementAPIsAgreementManagementAgreementLanguagesResourcesApiService) V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesPostExecute(r ApiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsAgreementManagementAgreementLanguagesResourcesApiService.V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/agreements/{agreementID}/languages"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", url.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"agreementID"+"}", url.PathEscape(parameterToString(r.agreementID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
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
