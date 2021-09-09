# \ManagementAPIsUsersUserAgreementConsentsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDGet**](ManagementAPIsUsersUserAgreementConsentsApi.md#V1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDGet) | **Get** /v1/environments/{envID}/users/{userID}/agreementConsents/{agreementID} | READ One User Agreement Consent
[**V1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDPost**](ManagementAPIsUsersUserAgreementConsentsApi.md#V1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDPost) | **Post** /v1/environments/{envID}/users/{userID}/agreementConsents/{agreementID} | Revoke Agreement
[**V1EnvironmentsEnvIDUsersUserIDAgreementConsentsGet**](ManagementAPIsUsersUserAgreementConsentsApi.md#V1EnvironmentsEnvIDUsersUserIDAgreementConsentsGet) | **Get** /v1/environments/{envID}/users/{userID}/agreementConsents | READ All User Agreement Consents



## V1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDGet

> V1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDGet(ctx, envID, userID, agreementID).Execute()

READ One User Agreement Consent



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
    agreementID := "agreementID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsUsersUserAgreementConsentsApi.V1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDGet(context.Background(), envID, userID, agreementID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsUsersUserAgreementConsentsApi.V1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDGet``: %v\n", err)
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
**agreementID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDPost

> V1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDPost(ctx, envID, userID, agreementID).ContentType(contentType).Execute()

Revoke Agreement



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
    agreementID := "agreementID_example" // string | 
    contentType := "application/vnd.pingidentity.consent.revoke+json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsUsersUserAgreementConsentsApi.V1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDPost(context.Background(), envID, userID, agreementID).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsUsersUserAgreementConsentsApi.V1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDPost``: %v\n", err)
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
**agreementID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUsersUserIDAgreementConsentsAgreementIDPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **contentType** | **string** |  | 

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


## V1EnvironmentsEnvIDUsersUserIDAgreementConsentsGet

> V1EnvironmentsEnvIDUsersUserIDAgreementConsentsGet(ctx, envID, userID).Execute()

READ All User Agreement Consents



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
    resp, r, err := api_client.ManagementAPIsUsersUserAgreementConsentsApi.V1EnvironmentsEnvIDUsersUserIDAgreementConsentsGet(context.Background(), envID, userID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsUsersUserAgreementConsentsApi.V1EnvironmentsEnvIDUsersUserIDAgreementConsentsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUsersUserIDAgreementConsentsGetRequest struct via the builder pattern


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

