# \ManagementAPIsUsersSessionsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDUsersUserIDSessionsGet**](ManagementAPIsUsersSessionsApi.md#V1EnvironmentsEnvIDUsersUserIDSessionsGet) | **Get** /v1/environments/{envID}/users/{userID}/sessions | READ All Sessions
[**V1EnvironmentsEnvIDUsersUserIDSessionsSessionIDDelete**](ManagementAPIsUsersSessionsApi.md#V1EnvironmentsEnvIDUsersUserIDSessionsSessionIDDelete) | **Delete** /v1/environments/{envID}/users/{userID}/sessions/{sessionID} | DELETE Session
[**V1EnvironmentsEnvIDUsersUserIDSessionsSessionIDGet**](ManagementAPIsUsersSessionsApi.md#V1EnvironmentsEnvIDUsersUserIDSessionsSessionIDGet) | **Get** /v1/environments/{envID}/users/{userID}/sessions/{sessionID} | READ One Session



## V1EnvironmentsEnvIDUsersUserIDSessionsGet

> V1EnvironmentsEnvIDUsersUserIDSessionsGet(ctx, envID, userID).Execute()

READ All Sessions



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
    userID := "userID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsUsersSessionsApi.V1EnvironmentsEnvIDUsersUserIDSessionsGet(context.Background(), envID, userID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsUsersSessionsApi.V1EnvironmentsEnvIDUsersUserIDSessionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**userID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUsersUserIDSessionsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDUsersUserIDSessionsSessionIDDelete

> V1EnvironmentsEnvIDUsersUserIDSessionsSessionIDDelete(ctx, envID, userID, sessionID).Execute()

DELETE Session



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
    userID := "userID_example" // string | 
    sessionID := "sessionID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsUsersSessionsApi.V1EnvironmentsEnvIDUsersUserIDSessionsSessionIDDelete(context.Background(), envID, userID, sessionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsUsersSessionsApi.V1EnvironmentsEnvIDUsersUserIDSessionsSessionIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**userID** | **string** |  | 
**sessionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUsersUserIDSessionsSessionIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDUsersUserIDSessionsSessionIDGet

> V1EnvironmentsEnvIDUsersUserIDSessionsSessionIDGet(ctx, envID, userID, sessionID).Execute()

READ One Session



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
    userID := "userID_example" // string | 
    sessionID := "sessionID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsUsersSessionsApi.V1EnvironmentsEnvIDUsersUserIDSessionsSessionIDGet(context.Background(), envID, userID, sessionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsUsersSessionsApi.V1EnvironmentsEnvIDUsersUserIDSessionsSessionIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**userID** | **string** |  | 
**sessionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUsersUserIDSessionsSessionIDGetRequest struct via the builder pattern


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

