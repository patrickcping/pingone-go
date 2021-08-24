# \ManagementAPIsActiveIdentityCountsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDActiveIdentityCountsGet**](ManagementAPIsActiveIdentityCountsApi.md#V1EnvironmentsEnvIDActiveIdentityCountsGet) | **Get** /v1/environments/{envID}/activeIdentityCounts | READ Active Identity Counts (Deprecated)
[**V1EnvironmentsEnvIDMetricsActiveIdentityCountsGet**](ManagementAPIsActiveIdentityCountsApi.md#V1EnvironmentsEnvIDMetricsActiveIdentityCountsGet) | **Get** /v1/environments/{envID}/metrics/activeIdentityCounts | READ Active Identity Counts by Date Range
[**V1OrganizationsOrgIDLicensesLicenseIDMetricsActiveIdentityCountsGet**](ManagementAPIsActiveIdentityCountsApi.md#V1OrganizationsOrgIDLicensesLicenseIDMetricsActiveIdentityCountsGet) | **Get** /v1/organizations/{orgID}/licenses/{licenseID}/metrics/activeIdentityCounts | READ Active Identity Counts by License



## V1EnvironmentsEnvIDActiveIdentityCountsGet

> V1EnvironmentsEnvIDActiveIdentityCountsGet(ctx, envID).Authorization(authorization).Filter(filter).Limit(limit).Order(order).Execute()

READ Active Identity Counts (Deprecated)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    envID := "envID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)
    filter := "startDate ge "2019-05-01T19:00:00Z" and samplingPeriod eq "10"" // string |  (optional)
    limit := int32(10) // int32 |  (optional)
    order := "-startDate" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsActiveIdentityCountsApi.V1EnvironmentsEnvIDActiveIdentityCountsGet(context.Background(), envID).Authorization(authorization).Filter(filter).Limit(limit).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsActiveIdentityCountsApi.V1EnvironmentsEnvIDActiveIdentityCountsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDActiveIdentityCountsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 
 **filter** | **string** |  | 
 **limit** | **int32** |  | 
 **order** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnvironmentsEnvIDMetricsActiveIdentityCountsGet

> V1EnvironmentsEnvIDMetricsActiveIdentityCountsGet(ctx, envID).Authorization(authorization).Filter(filter).Limit(limit).Order(order).SamplePeriod(samplePeriod).Execute()

READ Active Identity Counts by Date Range



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    envID := "envID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)
    filter := "startDate ge "2020-05-01T19:00:00Z"" // string |  (optional)
    limit := int32(10) // int32 |  (optional)
    order := "-startDate" // string |  (optional)
    samplePeriod := "MONTH" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsActiveIdentityCountsApi.V1EnvironmentsEnvIDMetricsActiveIdentityCountsGet(context.Background(), envID).Authorization(authorization).Filter(filter).Limit(limit).Order(order).SamplePeriod(samplePeriod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsActiveIdentityCountsApi.V1EnvironmentsEnvIDMetricsActiveIdentityCountsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDMetricsActiveIdentityCountsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 
 **filter** | **string** |  | 
 **limit** | **int32** |  | 
 **order** | **string** |  | 
 **samplePeriod** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OrganizationsOrgIDLicensesLicenseIDMetricsActiveIdentityCountsGet

> V1OrganizationsOrgIDLicensesLicenseIDMetricsActiveIdentityCountsGet(ctx, orgID, licenseID).ERRORUNKNOWN(eRRORUNKNOWN).AggregatedBy(aggregatedBy).Limit(limit).Order(order).Execute()

READ Active Identity Counts by License



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgID := "orgID_example" // string | 
    licenseID := "licenseID_example" // string | 
    eRRORUNKNOWN := "eRRORUNKNOWN_example" // string |  (optional)
    aggregatedBy := "calendarMonth" // string |  (optional)
    limit := int32(20) // int32 |  (optional)
    order := "-startDate" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsActiveIdentityCountsApi.V1OrganizationsOrgIDLicensesLicenseIDMetricsActiveIdentityCountsGet(context.Background(), orgID, licenseID).ERRORUNKNOWN(eRRORUNKNOWN).AggregatedBy(aggregatedBy).Limit(limit).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsActiveIdentityCountsApi.V1OrganizationsOrgIDLicensesLicenseIDMetricsActiveIdentityCountsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgID** | **string** |  | 
**licenseID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1OrganizationsOrgIDLicensesLicenseIDMetricsActiveIdentityCountsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **eRRORUNKNOWN** | **string** |  | 
 **aggregatedBy** | **string** |  | 
 **limit** | **int32** |  | 
 **order** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

