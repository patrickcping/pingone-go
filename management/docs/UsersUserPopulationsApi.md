# \UsersUserPopulationsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDUsersUserIDPopulationGet**](UsersUserPopulationsApi.md#V1EnvironmentsEnvironmentIDUsersUserIDPopulationGet) | **Get** /v1/environments/{environmentID}/users/{userID}/population | READ User Population
[**V1EnvironmentsEnvironmentIDUsersUserIDPopulationPut**](UsersUserPopulationsApi.md#V1EnvironmentsEnvironmentIDUsersUserIDPopulationPut) | **Put** /v1/environments/{environmentID}/users/{userID}/population | UPDATE User Population



## V1EnvironmentsEnvironmentIDUsersUserIDPopulationGet

> V1EnvironmentsEnvironmentIDUsersUserIDPopulationGet(ctx, environmentID, userID).Execute()

READ User Population

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
    environmentID := "environmentID_example" // string | 
    userID := "userID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersUserPopulationsApi.V1EnvironmentsEnvironmentIDUsersUserIDPopulationGet(context.Background(), environmentID, userID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersUserPopulationsApi.V1EnvironmentsEnvironmentIDUsersUserIDPopulationGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDUsersUserIDPopulationGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDUsersUserIDPopulationPut

> V1EnvironmentsEnvironmentIDUsersUserIDPopulationPut(ctx, environmentID, userID).Body(body).Execute()

UPDATE User Population

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
    environmentID := "environmentID_example" // string | 
    userID := "userID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersUserPopulationsApi.V1EnvironmentsEnvironmentIDUsersUserIDPopulationPut(context.Background(), environmentID, userID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersUserPopulationsApi.V1EnvironmentsEnvironmentIDUsersUserIDPopulationPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDUsersUserIDPopulationPutRequest struct via the builder pattern


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

