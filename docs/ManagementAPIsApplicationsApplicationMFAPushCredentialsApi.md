# \ManagementAPIsApplicationsApplicationMFAPushCredentialsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMFAPushCredential**](ManagementAPIsApplicationsApplicationMFAPushCredentialsApi.md#CreateMFAPushCredential) | **Post** /v1/environments/{envID}/applications/{appID}/pushCredentials | CREATE MFA Push Credential
[**DeleteMFAPushCredential**](ManagementAPIsApplicationsApplicationMFAPushCredentialsApi.md#DeleteMFAPushCredential) | **Delete** /v1/environments/{envID}/applications/{appID}/pushCredentials/{pushCredID} | DELETE MFA Push Credential
[**ReadAllMFAPushCredentials**](ManagementAPIsApplicationsApplicationMFAPushCredentialsApi.md#ReadAllMFAPushCredentials) | **Get** /v1/environments/{envID}/applications/{appID}/pushCredentials | READ All MFA Push Credentials
[**ReadOneMFAPushCredential**](ManagementAPIsApplicationsApplicationMFAPushCredentialsApi.md#ReadOneMFAPushCredential) | **Get** /v1/environments/{envID}/applications/{appID}/pushCredentials/{pushCredID} | READ One MFA Push Credential
[**UpdateMFAPushCredential**](ManagementAPIsApplicationsApplicationMFAPushCredentialsApi.md#UpdateMFAPushCredential) | **Put** /v1/environments/{envID}/applications/{appID}/pushCredentials/{pushCredID} | UPDATE MFA Push Credential



## CreateMFAPushCredential

> OneOfMFAPushCredentialAPNSMFAPushCredential CreateMFAPushCredential(ctx, envID, appID).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

CREATE MFA Push Credential



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
    appID := "appID_example" // string | 
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsApplicationsApplicationMFAPushCredentialsApi.CreateMFAPushCredential(context.Background(), envID, appID).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationMFAPushCredentialsApi.CreateMFAPushCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMFAPushCredential`: OneOfMFAPushCredentialAPNSMFAPushCredential
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsApplicationsApplicationMFAPushCredentialsApi.CreateMFAPushCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**appID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMFAPushCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) |  | 

### Return type

[**OneOfMFAPushCredentialAPNSMFAPushCredential**](oneOf&lt;MFAPushCredentialAPNS,MFAPushCredential&gt;.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMFAPushCredential

> DeleteMFAPushCredential(ctx, envID, appID, pushCredID).Authorization(authorization).Execute()

DELETE MFA Push Credential



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
    appID := "appID_example" // string | 
    pushCredID := "pushCredID_example" // string | 
    authorization := "Bearer {{accessToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsApplicationsApplicationMFAPushCredentialsApi.DeleteMFAPushCredential(context.Background(), envID, appID, pushCredID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationMFAPushCredentialsApi.DeleteMFAPushCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**appID** | **string** |  | 
**pushCredID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMFAPushCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** |  | 

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


## ReadAllMFAPushCredentials

> EntityArray ReadAllMFAPushCredentials(ctx, envID, appID).Execute()

READ All MFA Push Credentials



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
    appID := "appID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsApplicationsApplicationMFAPushCredentialsApi.ReadAllMFAPushCredentials(context.Background(), envID, appID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationMFAPushCredentialsApi.ReadAllMFAPushCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllMFAPushCredentials`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsApplicationsApplicationMFAPushCredentialsApi.ReadAllMFAPushCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**appID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllMFAPushCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EntityArray**](EntityArray.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadOneMFAPushCredential

> OneOfMFAPushCredentialAPNSMFAPushCredential ReadOneMFAPushCredential(ctx, envID, appID, pushCredID).Execute()

READ One MFA Push Credential



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
    appID := "appID_example" // string | 
    pushCredID := "pushCredID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsApplicationsApplicationMFAPushCredentialsApi.ReadOneMFAPushCredential(context.Background(), envID, appID, pushCredID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationMFAPushCredentialsApi.ReadOneMFAPushCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneMFAPushCredential`: OneOfMFAPushCredentialAPNSMFAPushCredential
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsApplicationsApplicationMFAPushCredentialsApi.ReadOneMFAPushCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**appID** | **string** |  | 
**pushCredID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneMFAPushCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OneOfMFAPushCredentialAPNSMFAPushCredential**](oneOf&lt;MFAPushCredentialAPNS,MFAPushCredential&gt;.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMFAPushCredential

> OneOfMFAPushCredentialAPNSMFAPushCredential UpdateMFAPushCredential(ctx, envID, appID, pushCredID).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

UPDATE MFA Push Credential



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
    appID := "appID_example" // string | 
    pushCredID := "pushCredID_example" // string | 
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsApplicationsApplicationMFAPushCredentialsApi.UpdateMFAPushCredential(context.Background(), envID, appID, pushCredID).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationMFAPushCredentialsApi.UpdateMFAPushCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMFAPushCredential`: OneOfMFAPushCredentialAPNSMFAPushCredential
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsApplicationsApplicationMFAPushCredentialsApi.UpdateMFAPushCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**appID** | **string** |  | 
**pushCredID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMFAPushCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) |  | 

### Return type

[**OneOfMFAPushCredentialAPNSMFAPushCredential**](oneOf&lt;MFAPushCredentialAPNS,MFAPushCredential&gt;.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

