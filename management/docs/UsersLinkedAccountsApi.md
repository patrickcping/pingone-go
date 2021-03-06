# \UsersLinkedAccountsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsGet**](UsersLinkedAccountsApi.md#V1EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsGet) | **Get** /v1/environments/{environmentID}/users/{userID}/linkedAccounts | READ Linked Accounts
[**V1EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDDelete**](UsersLinkedAccountsApi.md#V1EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDDelete) | **Delete** /v1/environments/{environmentID}/users/{userID}/linkedAccounts/{linkedAccountID} | DELETE Linked Account
[**V1EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDGet**](UsersLinkedAccountsApi.md#V1EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDGet) | **Get** /v1/environments/{environmentID}/users/{userID}/linkedAccounts/{linkedAccountID} | READ One Linked Account



## V1EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsGet

> V1EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsGet(ctx, environmentID, userID).Execute()

READ Linked Accounts

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
    resp, r, err := apiClient.UsersLinkedAccountsApi.V1EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsGet(context.Background(), environmentID, userID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersLinkedAccountsApi.V1EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDDelete

> V1EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDDelete(ctx, environmentID, userID, linkedAccountID).Execute()

DELETE Linked Account

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
    linkedAccountID := "linkedAccountID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersLinkedAccountsApi.V1EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDDelete(context.Background(), environmentID, userID, linkedAccountID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersLinkedAccountsApi.V1EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDDelete``: %v\n", err)
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
**linkedAccountID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDGet

> V1EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDGet(ctx, environmentID, userID, linkedAccountID).Execute()

READ One Linked Account

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
    linkedAccountID := "linkedAccountID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersLinkedAccountsApi.V1EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDGet(context.Background(), environmentID, userID, linkedAccountID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersLinkedAccountsApi.V1EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDGet``: %v\n", err)
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
**linkedAccountID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDGetRequest struct via the builder pattern


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

