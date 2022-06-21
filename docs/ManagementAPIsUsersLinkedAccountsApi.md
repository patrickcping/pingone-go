# \ManagementAPIsUsersLinkedAccountsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDUsersUserIDLinkedAccountsGet**](ManagementAPIsUsersLinkedAccountsApi.md#V1EnvironmentsEnvIDUsersUserIDLinkedAccountsGet) | **Get** /v1/environments/{envID}/users/{userID}/linkedAccounts | READ Linked Accounts
[**V1EnvironmentsEnvIDUsersUserIDLinkedAccountsLinkedAccountIDDelete**](ManagementAPIsUsersLinkedAccountsApi.md#V1EnvironmentsEnvIDUsersUserIDLinkedAccountsLinkedAccountIDDelete) | **Delete** /v1/environments/{envID}/users/{userID}/linkedAccounts/{linkedAccountID} | DELETE Linked Account
[**V1EnvironmentsEnvIDUsersUserIDLinkedAccountsLinkedAccountIDGet**](ManagementAPIsUsersLinkedAccountsApi.md#V1EnvironmentsEnvIDUsersUserIDLinkedAccountsLinkedAccountIDGet) | **Get** /v1/environments/{envID}/users/{userID}/linkedAccounts/{linkedAccountID} | READ One Linked Account



## V1EnvironmentsEnvIDUsersUserIDLinkedAccountsGet

> V1EnvironmentsEnvIDUsersUserIDLinkedAccountsGet(ctx, envID, userID).Execute()

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
    envID := "envID_example" // string | 
    userID := "userID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsUsersLinkedAccountsApi.V1EnvironmentsEnvIDUsersUserIDLinkedAccountsGet(context.Background(), envID, userID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsUsersLinkedAccountsApi.V1EnvironmentsEnvIDUsersUserIDLinkedAccountsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUsersUserIDLinkedAccountsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDUsersUserIDLinkedAccountsLinkedAccountIDDelete

> V1EnvironmentsEnvIDUsersUserIDLinkedAccountsLinkedAccountIDDelete(ctx, envID, userID, linkedAccountID).Execute()

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
    envID := "envID_example" // string | 
    userID := "userID_example" // string | 
    linkedAccountID := "linkedAccountID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsUsersLinkedAccountsApi.V1EnvironmentsEnvIDUsersUserIDLinkedAccountsLinkedAccountIDDelete(context.Background(), envID, userID, linkedAccountID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsUsersLinkedAccountsApi.V1EnvironmentsEnvIDUsersUserIDLinkedAccountsLinkedAccountIDDelete``: %v\n", err)
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
**linkedAccountID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUsersUserIDLinkedAccountsLinkedAccountIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDUsersUserIDLinkedAccountsLinkedAccountIDGet

> V1EnvironmentsEnvIDUsersUserIDLinkedAccountsLinkedAccountIDGet(ctx, envID, userID, linkedAccountID).Execute()

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
    envID := "envID_example" // string | 
    userID := "userID_example" // string | 
    linkedAccountID := "linkedAccountID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsUsersLinkedAccountsApi.V1EnvironmentsEnvIDUsersUserIDLinkedAccountsLinkedAccountIDGet(context.Background(), envID, userID, linkedAccountID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsUsersLinkedAccountsApi.V1EnvironmentsEnvIDUsersUserIDLinkedAccountsLinkedAccountIDGet``: %v\n", err)
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
**linkedAccountID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUsersUserIDLinkedAccountsLinkedAccountIDGetRequest struct via the builder pattern


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

