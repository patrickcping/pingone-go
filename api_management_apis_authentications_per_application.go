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

// ManagementAPIsAuthenticationsPerApplicationApiService ManagementAPIsAuthenticationsPerApplicationApi service
type ManagementAPIsAuthenticationsPerApplicationApiService service

type ApiV1EnvironmentsEnvIDApplicationSignonsGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsAuthenticationsPerApplicationApiService
	envID string
	limit *int32
	samplePeriod *int32
	samplePeriodCount *int32
	filter *string
}

func (r ApiV1EnvironmentsEnvIDApplicationSignonsGetRequest) Limit(limit int32) ApiV1EnvironmentsEnvIDApplicationSignonsGetRequest {
	r.limit = &limit
	return r
}
func (r ApiV1EnvironmentsEnvIDApplicationSignonsGetRequest) SamplePeriod(samplePeriod int32) ApiV1EnvironmentsEnvIDApplicationSignonsGetRequest {
	r.samplePeriod = &samplePeriod
	return r
}
func (r ApiV1EnvironmentsEnvIDApplicationSignonsGetRequest) SamplePeriodCount(samplePeriodCount int32) ApiV1EnvironmentsEnvIDApplicationSignonsGetRequest {
	r.samplePeriodCount = &samplePeriodCount
	return r
}
func (r ApiV1EnvironmentsEnvIDApplicationSignonsGetRequest) Filter(filter string) ApiV1EnvironmentsEnvIDApplicationSignonsGetRequest {
	r.filter = &filter
	return r
}

func (r ApiV1EnvironmentsEnvIDApplicationSignonsGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDApplicationSignonsGetExecute(r)
}

/*
V1EnvironmentsEnvIDApplicationSignonsGet READ Authentications Per Application (Partial)

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @return ApiV1EnvironmentsEnvIDApplicationSignonsGetRequest
*/
func (a *ManagementAPIsAuthenticationsPerApplicationApiService) V1EnvironmentsEnvIDApplicationSignonsGet(ctx _context.Context, envID string) ApiV1EnvironmentsEnvIDApplicationSignonsGetRequest {
	return ApiV1EnvironmentsEnvIDApplicationSignonsGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
	}
}

// Execute executes the request
func (a *ManagementAPIsAuthenticationsPerApplicationApiService) V1EnvironmentsEnvIDApplicationSignonsGetExecute(r ApiV1EnvironmentsEnvIDApplicationSignonsGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsAuthenticationsPerApplicationApiService.V1EnvironmentsEnvIDApplicationSignonsGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/applicationSignons"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.samplePeriod != nil {
		localVarQueryParams.Add("samplePeriod", parameterToString(*r.samplePeriod, ""))
	}
	if r.samplePeriodCount != nil {
		localVarQueryParams.Add("samplePeriodCount", parameterToString(*r.samplePeriodCount, ""))
	}
	if r.filter != nil {
		localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
	}
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
