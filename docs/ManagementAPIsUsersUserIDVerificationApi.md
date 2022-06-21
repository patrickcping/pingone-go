# \ManagementAPIsUsersUserIDVerificationApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsGet**](ManagementAPIsUsersUserIDVerificationApi.md#V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsGet) | **Get** /v1/environments/{envID}/users/{userID}/verifyTransactions | READ All ID Verification Transaction Records for a User
[**V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsPost**](ManagementAPIsUsersUserIDVerificationApi.md#V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsPost) | **Post** /v1/environments/{envID}/users/{userID}/verifyTransactions | CREATE ID Verification Transaction Record for a User
[**V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDDelete**](ManagementAPIsUsersUserIDVerificationApi.md#V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDDelete) | **Delete** /v1/environments/{envID}/users/{userID}/verifyTransactions/{transactionID} | DELETE ID Verification Transaction Record for a User
[**V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDGet**](ManagementAPIsUsersUserIDVerificationApi.md#V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDGet) | **Get** /v1/environments/{envID}/users/{userID}/verifyTransactions/{transactionID} | READ ID Verification Transaction Record for a User
[**V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDPut**](ManagementAPIsUsersUserIDVerificationApi.md#V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDPut) | **Put** /v1/environments/{envID}/users/{userID}/verifyTransactions/{transactionID} | UPDATE ID Verification Transaction Record for a User



## V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsGet

> V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsGet(ctx, envID, userID).Execute()

READ All ID Verification Transaction Records for a User



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
    resp, r, err := apiClient.ManagementAPIsUsersUserIDVerificationApi.V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsGet(context.Background(), envID, userID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsUsersUserIDVerificationApi.V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsPost

> V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsPost(ctx, envID, userID).Execute()

CREATE ID Verification Transaction Record for a User



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
    resp, r, err := apiClient.ManagementAPIsUsersUserIDVerificationApi.V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsPost(context.Background(), envID, userID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsUsersUserIDVerificationApi.V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsPostRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDDelete

> V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDDelete(ctx, envID, userID, transactionID).Execute()

DELETE ID Verification Transaction Record for a User



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
    transactionID := "transactionID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsUsersUserIDVerificationApi.V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDDelete(context.Background(), envID, userID, transactionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsUsersUserIDVerificationApi.V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDDelete``: %v\n", err)
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
**transactionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDGet

> V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDGet(ctx, envID, userID, transactionID).Execute()

READ ID Verification Transaction Record for a User



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
    transactionID := "transactionID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsUsersUserIDVerificationApi.V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDGet(context.Background(), envID, userID, transactionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsUsersUserIDVerificationApi.V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDGet``: %v\n", err)
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
**transactionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDPut

> V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDPut(ctx, envID, userID, transactionID).Body(body).Execute()

UPDATE ID Verification Transaction Record for a User



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
    transactionID := "transactionID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsUsersUserIDVerificationApi.V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDPut(context.Background(), envID, userID, transactionID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsUsersUserIDVerificationApi.V1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDPut``: %v\n", err)
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
**transactionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUsersUserIDVerifyTransactionsTransactionIDPutRequest struct via the builder pattern


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

