# \ManagementAPIsUserActivitiesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDUserActivitiesGet**](ManagementAPIsUserActivitiesApi.md#V1EnvironmentsEnvIDUserActivitiesGet) | **Get** /v1/environments/{envID}/userActivities | READ User Activities



## V1EnvironmentsEnvIDUserActivitiesGet

> V1EnvironmentsEnvIDUserActivitiesGet(ctx, envID).Filter(filter).Execute()

READ User Activities



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
    filter := "startDate eq "2018-02-17T09:10:12-04:00" and endDate eq "2018-02-23T09:10:12-04:00"" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsUserActivitiesApi.V1EnvironmentsEnvIDUserActivitiesGet(context.Background(), envID).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsUserActivitiesApi.V1EnvironmentsEnvIDUserActivitiesGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUserActivitiesGetRequest struct via the builder pattern


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

