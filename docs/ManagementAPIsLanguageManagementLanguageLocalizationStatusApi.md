# \ManagementAPIsLanguageManagementLanguageLocalizationStatusApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDLanguagesLanguageIDStatusGet**](ManagementAPIsLanguageManagementLanguageLocalizationStatusApi.md#V1EnvironmentsEnvIDLanguagesLanguageIDStatusGet) | **Get** /v1/environments/{envID}/languages/{languageID}/status | READ Language Localization Status
[**V1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDDelete**](ManagementAPIsLanguageManagementLanguageLocalizationStatusApi.md#V1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDDelete) | **Delete** /v1/environments/{envID}/languages/{languageID}/status/{l10nStatusID} | DELETE Language Localization Status
[**V1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDGet**](ManagementAPIsLanguageManagementLanguageLocalizationStatusApi.md#V1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDGet) | **Get** /v1/environments/{envID}/languages/{languageID}/status/{l10nStatusID} | READ One Language Localization Status
[**V1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDPut**](ManagementAPIsLanguageManagementLanguageLocalizationStatusApi.md#V1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDPut) | **Put** /v1/environments/{envID}/languages/{languageID}/status/{l10nStatusID} | CREATE Language Localization Status
[**V1EnvironmentsEnvIDLanguagesLanguageIDStatusPost**](ManagementAPIsLanguageManagementLanguageLocalizationStatusApi.md#V1EnvironmentsEnvIDLanguagesLanguageIDStatusPost) | **Post** /v1/environments/{envID}/languages/{languageID}/status | CREATE Language Localization Status



## V1EnvironmentsEnvIDLanguagesLanguageIDStatusGet

> V1EnvironmentsEnvIDLanguagesLanguageIDStatusGet(ctx, envID, languageID).Authorization(authorization).Execute()

READ Language Localization Status



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
    languageID := "languageID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsLanguageManagementLanguageLocalizationStatusApi.V1EnvironmentsEnvIDLanguagesLanguageIDStatusGet(context.Background(), envID, languageID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsLanguageManagementLanguageLocalizationStatusApi.V1EnvironmentsEnvIDLanguagesLanguageIDStatusGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**languageID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDLanguagesLanguageIDStatusGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDDelete

> V1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDDelete(ctx, envID, languageID, l10nStatusID).Authorization(authorization).Execute()

DELETE Language Localization Status



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
    languageID := "languageID_example" // string | 
    l10nStatusID := "l10nStatusID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsLanguageManagementLanguageLocalizationStatusApi.V1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDDelete(context.Background(), envID, languageID, l10nStatusID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsLanguageManagementLanguageLocalizationStatusApi.V1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**languageID** | **string** |  | 
**l10nStatusID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDGet

> V1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDGet(ctx, envID, languageID, l10nStatusID).Authorization(authorization).Execute()

READ One Language Localization Status



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
    languageID := "languageID_example" // string | 
    l10nStatusID := "l10nStatusID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsLanguageManagementLanguageLocalizationStatusApi.V1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDGet(context.Background(), envID, languageID, l10nStatusID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsLanguageManagementLanguageLocalizationStatusApi.V1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**languageID** | **string** |  | 
**l10nStatusID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDPut

> V1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDPut(ctx, envID, languageID, l10nStatusID).ContentType(contentType).Authorization(authorization).Body(body).Execute()

CREATE Language Localization Status



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
    languageID := "languageID_example" // string | 
    l10nStatusID := "l10nStatusID_example" // string | 
    contentType := "application/json" // string |  (optional)
    authorization := "Bearer {{jwtToken}}" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsLanguageManagementLanguageLocalizationStatusApi.V1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDPut(context.Background(), envID, languageID, l10nStatusID).ContentType(contentType).Authorization(authorization).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsLanguageManagementLanguageLocalizationStatusApi.V1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**languageID** | **string** |  | 
**l10nStatusID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDLanguagesLanguageIDStatusL10nStatusIDPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **contentType** | **string** |  | 
 **authorization** | **string** |  | 
 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnvironmentsEnvIDLanguagesLanguageIDStatusPost

> V1EnvironmentsEnvIDLanguagesLanguageIDStatusPost(ctx, envID, languageID).ContentType(contentType).Authorization(authorization).Body(body).Execute()

CREATE Language Localization Status



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
    languageID := "languageID_example" // string | 
    contentType := "application/json" // string |  (optional)
    authorization := "Bearer {{jwtToken}}" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsLanguageManagementLanguageLocalizationStatusApi.V1EnvironmentsEnvIDLanguagesLanguageIDStatusPost(context.Background(), envID, languageID).ContentType(contentType).Authorization(authorization).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsLanguageManagementLanguageLocalizationStatusApi.V1EnvironmentsEnvIDLanguagesLanguageIDStatusPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**languageID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDLanguagesLanguageIDStatusPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contentType** | **string** |  | 
 **authorization** | **string** |  | 
 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

