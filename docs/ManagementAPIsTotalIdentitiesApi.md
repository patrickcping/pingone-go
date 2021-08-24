# \ManagementAPIsTotalIdentitiesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDTotalIdentitiesGet**](ManagementAPIsTotalIdentitiesApi.md#V1EnvironmentsEnvIDTotalIdentitiesGet) | **Get** /v1/environments/{envID}/totalIdentities | READ Total Identity Counts



## V1EnvironmentsEnvIDTotalIdentitiesGet

> V1EnvironmentsEnvIDTotalIdentitiesGet(ctx, envID).Authorization(authorization).Filter(filter).Execute()

READ Total Identity Counts



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
    filter := "startDate eq "2019-05-01T19:00:00Z" and endDate eq "2019-05-31T19:00:00Z"" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsTotalIdentitiesApi.V1EnvironmentsEnvIDTotalIdentitiesGet(context.Background(), envID).Authorization(authorization).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsTotalIdentitiesApi.V1EnvironmentsEnvIDTotalIdentitiesGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDTotalIdentitiesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 
 **filter** | **string** |  | 

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

