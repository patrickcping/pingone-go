# \ManagementAPIsLanguageManagementLanguagesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDLanguagesGet**](ManagementAPIsLanguageManagementLanguagesApi.md#V1EnvironmentsEnvIDLanguagesGet) | **Get** /v1/environments/{envID}/languages | READ Languages
[**V1EnvironmentsEnvIDLanguagesLanguageIDDelete**](ManagementAPIsLanguageManagementLanguagesApi.md#V1EnvironmentsEnvIDLanguagesLanguageIDDelete) | **Delete** /v1/environments/{envID}/languages/{languageID} | DELETE Language
[**V1EnvironmentsEnvIDLanguagesLanguageIDGet**](ManagementAPIsLanguageManagementLanguagesApi.md#V1EnvironmentsEnvIDLanguagesLanguageIDGet) | **Get** /v1/environments/{envID}/languages/{languageID} | READ One Language
[**V1EnvironmentsEnvIDLanguagesLanguageIDPut**](ManagementAPIsLanguageManagementLanguagesApi.md#V1EnvironmentsEnvIDLanguagesLanguageIDPut) | **Put** /v1/environments/{envID}/languages/{languageID} | UPDATE Language 
[**V1EnvironmentsEnvIDLanguagesPost**](ManagementAPIsLanguageManagementLanguagesApi.md#V1EnvironmentsEnvIDLanguagesPost) | **Post** /v1/environments/{envID}/languages/ | CREATE Language



## V1EnvironmentsEnvIDLanguagesGet

> V1EnvironmentsEnvIDLanguagesGet(ctx, envID).Execute()

READ Languages



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsLanguageManagementLanguagesApi.V1EnvironmentsEnvIDLanguagesGet(context.Background(), envID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsLanguageManagementLanguagesApi.V1EnvironmentsEnvIDLanguagesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDLanguagesGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDLanguagesLanguageIDDelete

> V1EnvironmentsEnvIDLanguagesLanguageIDDelete(ctx, envID, languageID).Execute()

DELETE Language



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsLanguageManagementLanguagesApi.V1EnvironmentsEnvIDLanguagesLanguageIDDelete(context.Background(), envID, languageID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsLanguageManagementLanguagesApi.V1EnvironmentsEnvIDLanguagesLanguageIDDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDLanguagesLanguageIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDLanguagesLanguageIDGet

> V1EnvironmentsEnvIDLanguagesLanguageIDGet(ctx, envID, languageID).Execute()

READ One Language



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsLanguageManagementLanguagesApi.V1EnvironmentsEnvIDLanguagesLanguageIDGet(context.Background(), envID, languageID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsLanguageManagementLanguagesApi.V1EnvironmentsEnvIDLanguagesLanguageIDGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDLanguagesLanguageIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDLanguagesLanguageIDPut

> V1EnvironmentsEnvIDLanguagesLanguageIDPut(ctx, envID, languageID).Body(body).Execute()

UPDATE Language 



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
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsLanguageManagementLanguagesApi.V1EnvironmentsEnvIDLanguagesLanguageIDPut(context.Background(), envID, languageID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsLanguageManagementLanguagesApi.V1EnvironmentsEnvIDLanguagesLanguageIDPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDLanguagesLanguageIDPutRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDLanguagesPost

> V1EnvironmentsEnvIDLanguagesPost(ctx, envID).Body(body).Execute()

CREATE Language



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
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsLanguageManagementLanguagesApi.V1EnvironmentsEnvIDLanguagesPost(context.Background(), envID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsLanguageManagementLanguagesApi.V1EnvironmentsEnvIDLanguagesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDLanguagesPostRequest struct via the builder pattern


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

