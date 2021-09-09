# \ManagementAPIsApplicationsApplicationMFAPushCredentialsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsGet**](ManagementAPIsApplicationsApplicationMFAPushCredentialsApi.md#V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsGet) | **Get** /v1/environments/{envID}/applications/{appID}/pushCredentials | READ All MFA Push Credentials
[**V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPost**](ManagementAPIsApplicationsApplicationMFAPushCredentialsApi.md#V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPost) | **Post** /v1/environments/{envID}/applications/{appID}/pushCredentials | CREATE MFA Push Credential (FCM)
[**V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDDelete**](ManagementAPIsApplicationsApplicationMFAPushCredentialsApi.md#V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDDelete) | **Delete** /v1/environments/{envID}/applications/{appID}/pushCredentials/{pushCredID} | DELETE MFA Push Credential
[**V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDGet**](ManagementAPIsApplicationsApplicationMFAPushCredentialsApi.md#V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDGet) | **Get** /v1/environments/{envID}/applications/{appID}/pushCredentials/{pushCredID} | READ One MFA Push Credential
[**V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDPut**](ManagementAPIsApplicationsApplicationMFAPushCredentialsApi.md#V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDPut) | **Put** /v1/environments/{envID}/applications/{appID}/pushCredentials/{pushCredID} | UPDATE MFA Push Credential



## V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsGet

> V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsGet(ctx, envID, appID).Execute()

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
    resp, r, err := api_client.ManagementAPIsApplicationsApplicationMFAPushCredentialsApi.V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsGet(context.Background(), envID, appID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationMFAPushCredentialsApi.V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsGet``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPost

> V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPost(ctx, envID, appID).ContentType(contentType).Body(body).Execute()

CREATE MFA Push Credential (FCM)



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
    contentType := "application/json" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsApplicationsApplicationMFAPushCredentialsApi.V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPost(context.Background(), envID, appID).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationMFAPushCredentialsApi.V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPost``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contentType** | **string** |  | 
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


## V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDDelete

> V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDDelete(ctx, envID, appID, pushCredID).Authorization(authorization).Execute()

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
    resp, r, err := api_client.ManagementAPIsApplicationsApplicationMFAPushCredentialsApi.V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDDelete(context.Background(), envID, appID, pushCredID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationMFAPushCredentialsApi.V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDGet

> V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDGet(ctx, envID, appID, pushCredID).Execute()

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
    resp, r, err := api_client.ManagementAPIsApplicationsApplicationMFAPushCredentialsApi.V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDGet(context.Background(), envID, appID, pushCredID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationMFAPushCredentialsApi.V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDPut

> V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDPut(ctx, envID, appID, pushCredID).ContentType(contentType).Body(body).Execute()

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
    contentType := "application/json" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsApplicationsApplicationMFAPushCredentialsApi.V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDPut(context.Background(), envID, appID, pushCredID).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationMFAPushCredentialsApi.V1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDApplicationsAppIDPushCredentialsPushCredIDPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **contentType** | **string** |  | 
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

