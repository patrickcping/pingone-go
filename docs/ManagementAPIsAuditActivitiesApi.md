# \ManagementAPIsAuditActivitiesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDActivitiesActivityIDGet**](ManagementAPIsAuditActivitiesApi.md#V1EnvironmentsEnvIDActivitiesActivityIDGet) | **Get** /v1/environments/{envID}/activities/{activityID} | GET One User Activity
[**V1EnvironmentsEnvIDActivitiesGet**](ManagementAPIsAuditActivitiesApi.md#V1EnvironmentsEnvIDActivitiesGet) | **Get** /v1/environments/{envID}/activities | GET User Activities
[**V1EnvironmentsEnvIDActivitiesPost**](ManagementAPIsAuditActivitiesApi.md#V1EnvironmentsEnvIDActivitiesPost) | **Post** /v1/environments/{envID}/activities | GET User Activities



## V1EnvironmentsEnvIDActivitiesActivityIDGet

> V1EnvironmentsEnvIDActivitiesActivityIDGet(ctx, envID, activityID).Execute()

GET One User Activity



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
    activityID := "activityID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsAuditActivitiesApi.V1EnvironmentsEnvIDActivitiesActivityIDGet(context.Background(), envID, activityID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsAuditActivitiesApi.V1EnvironmentsEnvIDActivitiesActivityIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**activityID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDActivitiesActivityIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnvironmentsEnvIDActivitiesGet

> V1EnvironmentsEnvIDActivitiesGet(ctx, envID).Filter(filter).Execute()

GET User Activities



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
    filter := "recordedat gt "2018-08-20T00:00:00Z" AND recordedat lt "2018-08-22T23:59:00Z"" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsAuditActivitiesApi.V1EnvironmentsEnvIDActivitiesGet(context.Background(), envID).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsAuditActivitiesApi.V1EnvironmentsEnvIDActivitiesGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDActivitiesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnvironmentsEnvIDActivitiesPost

> V1EnvironmentsEnvIDActivitiesPost(ctx, envID).Body(body).Execute()

GET User Activities



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
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsAuditActivitiesApi.V1EnvironmentsEnvIDActivitiesPost(context.Background(), envID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsAuditActivitiesApi.V1EnvironmentsEnvIDActivitiesPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDActivitiesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

