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

// ManagementAPIsActiveIdentityCountsApiService ManagementAPIsActiveIdentityCountsApi service
type ManagementAPIsActiveIdentityCountsApiService service

type ApiV1EnvironmentsEnvIDActiveIdentityCountsGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsActiveIdentityCountsApiService
	envID string
	filter *string
	limit *int32
	order *string
}

func (r ApiV1EnvironmentsEnvIDActiveIdentityCountsGetRequest) Filter(filter string) ApiV1EnvironmentsEnvIDActiveIdentityCountsGetRequest {
	r.filter = &filter
	return r
}
func (r ApiV1EnvironmentsEnvIDActiveIdentityCountsGetRequest) Limit(limit int32) ApiV1EnvironmentsEnvIDActiveIdentityCountsGetRequest {
	r.limit = &limit
	return r
}
func (r ApiV1EnvironmentsEnvIDActiveIdentityCountsGetRequest) Order(order string) ApiV1EnvironmentsEnvIDActiveIdentityCountsGetRequest {
	r.order = &order
	return r
}

func (r ApiV1EnvironmentsEnvIDActiveIdentityCountsGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDActiveIdentityCountsGetExecute(r)
}

/*
V1EnvironmentsEnvIDActiveIdentityCountsGet READ Active Identity Counts (Deprecated)

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @return ApiV1EnvironmentsEnvIDActiveIdentityCountsGetRequest
*/
func (a *ManagementAPIsActiveIdentityCountsApiService) V1EnvironmentsEnvIDActiveIdentityCountsGet(ctx _context.Context, envID string) ApiV1EnvironmentsEnvIDActiveIdentityCountsGetRequest {
	return ApiV1EnvironmentsEnvIDActiveIdentityCountsGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
	}
}

// Execute executes the request
func (a *ManagementAPIsActiveIdentityCountsApiService) V1EnvironmentsEnvIDActiveIdentityCountsGetExecute(r ApiV1EnvironmentsEnvIDActiveIdentityCountsGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsActiveIdentityCountsApiService.V1EnvironmentsEnvIDActiveIdentityCountsGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/activeIdentityCounts"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.filter != nil {
		localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.order != nil {
		localVarQueryParams.Add("order", parameterToString(*r.order, ""))
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

type ApiV1EnvironmentsEnvIDMetricsActiveIdentityCountsGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsActiveIdentityCountsApiService
	envID string
	filter *string
	limit *int32
	order *string
	samplePeriod *string
}

func (r ApiV1EnvironmentsEnvIDMetricsActiveIdentityCountsGetRequest) Filter(filter string) ApiV1EnvironmentsEnvIDMetricsActiveIdentityCountsGetRequest {
	r.filter = &filter
	return r
}
func (r ApiV1EnvironmentsEnvIDMetricsActiveIdentityCountsGetRequest) Limit(limit int32) ApiV1EnvironmentsEnvIDMetricsActiveIdentityCountsGetRequest {
	r.limit = &limit
	return r
}
func (r ApiV1EnvironmentsEnvIDMetricsActiveIdentityCountsGetRequest) Order(order string) ApiV1EnvironmentsEnvIDMetricsActiveIdentityCountsGetRequest {
	r.order = &order
	return r
}
func (r ApiV1EnvironmentsEnvIDMetricsActiveIdentityCountsGetRequest) SamplePeriod(samplePeriod string) ApiV1EnvironmentsEnvIDMetricsActiveIdentityCountsGetRequest {
	r.samplePeriod = &samplePeriod
	return r
}

func (r ApiV1EnvironmentsEnvIDMetricsActiveIdentityCountsGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDMetricsActiveIdentityCountsGetExecute(r)
}

/*
V1EnvironmentsEnvIDMetricsActiveIdentityCountsGet READ Active Identity Counts by Date Range

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @return ApiV1EnvironmentsEnvIDMetricsActiveIdentityCountsGetRequest
*/
func (a *ManagementAPIsActiveIdentityCountsApiService) V1EnvironmentsEnvIDMetricsActiveIdentityCountsGet(ctx _context.Context, envID string) ApiV1EnvironmentsEnvIDMetricsActiveIdentityCountsGetRequest {
	return ApiV1EnvironmentsEnvIDMetricsActiveIdentityCountsGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
	}
}

// Execute executes the request
func (a *ManagementAPIsActiveIdentityCountsApiService) V1EnvironmentsEnvIDMetricsActiveIdentityCountsGetExecute(r ApiV1EnvironmentsEnvIDMetricsActiveIdentityCountsGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsActiveIdentityCountsApiService.V1EnvironmentsEnvIDMetricsActiveIdentityCountsGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/metrics/activeIdentityCounts"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", _neturl.PathEscape(parameterToString(r.envID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.filter != nil {
		localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.order != nil {
		localVarQueryParams.Add("order", parameterToString(*r.order, ""))
	}
	if r.samplePeriod != nil {
		localVarQueryParams.Add("samplePeriod", parameterToString(*r.samplePeriod, ""))
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

type ApiV1OrganizationsOrgIDLicensesLicenseIDMetricsActiveIdentityCountsGetRequest struct {
	ctx _context.Context
	ApiService *ManagementAPIsActiveIdentityCountsApiService
	orgID string
	licenseID string
	aggregatedBy *string
	limit *int32
	order *string
}

func (r ApiV1OrganizationsOrgIDLicensesLicenseIDMetricsActiveIdentityCountsGetRequest) AggregatedBy(aggregatedBy string) ApiV1OrganizationsOrgIDLicensesLicenseIDMetricsActiveIdentityCountsGetRequest {
	r.aggregatedBy = &aggregatedBy
	return r
}
func (r ApiV1OrganizationsOrgIDLicensesLicenseIDMetricsActiveIdentityCountsGetRequest) Limit(limit int32) ApiV1OrganizationsOrgIDLicensesLicenseIDMetricsActiveIdentityCountsGetRequest {
	r.limit = &limit
	return r
}
func (r ApiV1OrganizationsOrgIDLicensesLicenseIDMetricsActiveIdentityCountsGetRequest) Order(order string) ApiV1OrganizationsOrgIDLicensesLicenseIDMetricsActiveIdentityCountsGetRequest {
	r.order = &order
	return r
}

func (r ApiV1OrganizationsOrgIDLicensesLicenseIDMetricsActiveIdentityCountsGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.V1OrganizationsOrgIDLicensesLicenseIDMetricsActiveIdentityCountsGetExecute(r)
}

/*
V1OrganizationsOrgIDLicensesLicenseIDMetricsActiveIdentityCountsGet READ Active Identity Counts by License

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgID
 @param licenseID
 @return ApiV1OrganizationsOrgIDLicensesLicenseIDMetricsActiveIdentityCountsGetRequest
*/
func (a *ManagementAPIsActiveIdentityCountsApiService) V1OrganizationsOrgIDLicensesLicenseIDMetricsActiveIdentityCountsGet(ctx _context.Context, orgID string, licenseID string) ApiV1OrganizationsOrgIDLicensesLicenseIDMetricsActiveIdentityCountsGetRequest {
	return ApiV1OrganizationsOrgIDLicensesLicenseIDMetricsActiveIdentityCountsGetRequest{
		ApiService: a,
		ctx: ctx,
		orgID: orgID,
		licenseID: licenseID,
	}
}

// Execute executes the request
func (a *ManagementAPIsActiveIdentityCountsApiService) V1OrganizationsOrgIDLicensesLicenseIDMetricsActiveIdentityCountsGetExecute(r ApiV1OrganizationsOrgIDLicensesLicenseIDMetricsActiveIdentityCountsGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManagementAPIsActiveIdentityCountsApiService.V1OrganizationsOrgIDLicensesLicenseIDMetricsActiveIdentityCountsGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/organizations/{orgID}/licenses/{licenseID}/metrics/activeIdentityCounts"
	localVarPath = strings.Replace(localVarPath, "{"+"orgID"+"}", _neturl.PathEscape(parameterToString(r.orgID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"licenseID"+"}", _neturl.PathEscape(parameterToString(r.licenseID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.aggregatedBy != nil {
		localVarQueryParams.Add("aggregatedBy", parameterToString(*r.aggregatedBy, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.order != nil {
		localVarQueryParams.Add("order", parameterToString(*r.order, ""))
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
