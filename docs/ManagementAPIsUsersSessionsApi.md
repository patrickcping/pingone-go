# \ManagementAPIsUsersSessionsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDUsersUserIDSessionsGet**](ManagementAPIsUsersSessionsApi.md#V1EnvironmentsEnvIDUsersUserIDSessionsGet) | **Get** /v1/environments/{envID}/users/{userID}/sessions | READ All Sessions
[**V1EnvironmentsEnvIDUsersUserIDSessionssessionIDDelete**](ManagementAPIsUsersSessionsApi.md#V1EnvironmentsEnvIDUsersUserIDSessionssessionIDDelete) | **Delete** /v1/environments/{envID}/users/{userID}/sessions{sessionID} | DELETE Session
[**V1EnvironmentsEnvIDUsersUserIDSessionssessionIDGet**](ManagementAPIsUsersSessionsApi.md#V1EnvironmentsEnvIDUsersUserIDSessionssessionIDGet) | **Get** /v1/environments/{envID}/users/{userID}/sessions{sessionID} | READ One Session



## V1EnvironmentsEnvIDUsersUserIDSessionsGet

> V1EnvironmentsEnvIDUsersUserIDSessionsGet(ctx, envID, userID).Authorization(authorization).Execute()

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
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsUsersSessionsApi.V1EnvironmentsEnvIDUsersUserIDSessionsGet(context.Background(), envID, userID).Authorization(authorization).Execute()
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


 **authorization** | **string** |  | 

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


## V1EnvironmentsEnvIDUsersUserIDSessionssessionIDDelete

> V1EnvironmentsEnvIDUsersUserIDSessionssessionIDDelete(ctx, envID, userID, sessionID).Authorization(authorization).Execute()

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
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsUsersSessionsApi.V1EnvironmentsEnvIDUsersUserIDSessionssessionIDDelete(context.Background(), envID, userID, sessionID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsUsersSessionsApi.V1EnvironmentsEnvIDUsersUserIDSessionssessionIDDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUsersUserIDSessionssessionIDDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** |  | 

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


## V1EnvironmentsEnvIDUsersUserIDSessionssessionIDGet

> V1EnvironmentsEnvIDUsersUserIDSessionssessionIDGet(ctx, envID, userID, sessionID).Authorization(authorization).Execute()

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
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsUsersSessionsApi.V1EnvironmentsEnvIDUsersUserIDSessionssessionIDGet(context.Background(), envID, userID, sessionID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsUsersSessionsApi.V1EnvironmentsEnvIDUsersUserIDSessionssessionIDGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUsersUserIDSessionssessionIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** |  | 

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

