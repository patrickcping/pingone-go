# \ManagementAPIsNotificationsNotificationsTemplatesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDTemplatesGet**](ManagementAPIsNotificationsNotificationsTemplatesApi.md#V1EnvironmentsEnvIDTemplatesGet) | **Get** /v1/environments/{envID}/templates | READ All Templates
[**V1EnvironmentsEnvIDTemplatesTemplateNameContentsContentIDDelete**](ManagementAPIsNotificationsNotificationsTemplatesApi.md#V1EnvironmentsEnvIDTemplatesTemplateNameContentsContentIDDelete) | **Delete** /v1/environments/{envID}/templates/{templateName}/contents/{contentID} | DELETE Content
[**V1EnvironmentsEnvIDTemplatesTemplateNameContentsContentIDGet**](ManagementAPIsNotificationsNotificationsTemplatesApi.md#V1EnvironmentsEnvIDTemplatesTemplateNameContentsContentIDGet) | **Get** /v1/environments/{envID}/templates/{templateName}/contents/{contentID} | READ One Content
[**V1EnvironmentsEnvIDTemplatesTemplateNameContentsContentIDPut**](ManagementAPIsNotificationsNotificationsTemplatesApi.md#V1EnvironmentsEnvIDTemplatesTemplateNameContentsContentIDPut) | **Put** /v1/environments/{envID}/templates/{templateName}/contents/{contentID} | UPDATE Push Content
[**V1EnvironmentsEnvIDTemplatesTemplateNameContentsDelete**](ManagementAPIsNotificationsNotificationsTemplatesApi.md#V1EnvironmentsEnvIDTemplatesTemplateNameContentsDelete) | **Delete** /v1/environments/{envID}/templates/{templateName}/contents | DELETE Bulk Variant Contents
[**V1EnvironmentsEnvIDTemplatesTemplateNameContentsGet**](ManagementAPIsNotificationsNotificationsTemplatesApi.md#V1EnvironmentsEnvIDTemplatesTemplateNameContentsGet) | **Get** /v1/environments/{envID}/templates/{templateName}/contents | READ All Contents
[**V1EnvironmentsEnvIDTemplatesTemplateNameContentsPatch**](ManagementAPIsNotificationsNotificationsTemplatesApi.md#V1EnvironmentsEnvIDTemplatesTemplateNameContentsPatch) | **Patch** /v1/environments/{envID}/templates/{templateName}/contents | PATCH Bulk Variant Contents
[**V1EnvironmentsEnvIDTemplatesTemplateNameContentsPost**](ManagementAPIsNotificationsNotificationsTemplatesApi.md#V1EnvironmentsEnvIDTemplatesTemplateNameContentsPost) | **Post** /v1/environments/{envID}/templates/{templateName}/contents | CREATE Push Content
[**V1EnvironmentsEnvIDTemplatesTemplateNameGet**](ManagementAPIsNotificationsNotificationsTemplatesApi.md#V1EnvironmentsEnvIDTemplatesTemplateNameGet) | **Get** /v1/environments/{envID}/templates/{templateName} | READ One Template



## V1EnvironmentsEnvIDTemplatesGet

> V1EnvironmentsEnvIDTemplatesGet(ctx, envID).Authorization(authorization).Execute()

READ All Templates



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
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsNotificationsNotificationsTemplatesApi.V1EnvironmentsEnvIDTemplatesGet(context.Background(), envID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsNotificationsNotificationsTemplatesApi.V1EnvironmentsEnvIDTemplatesGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDTemplatesGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDTemplatesTemplateNameContentsContentIDDelete

> V1EnvironmentsEnvIDTemplatesTemplateNameContentsContentIDDelete(ctx, envID, templateName, contentID).Authorization(authorization).Execute()

DELETE Content



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
    templateName := "templateName_example" // string | 
    contentID := "contentID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsNotificationsNotificationsTemplatesApi.V1EnvironmentsEnvIDTemplatesTemplateNameContentsContentIDDelete(context.Background(), envID, templateName, contentID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsNotificationsNotificationsTemplatesApi.V1EnvironmentsEnvIDTemplatesTemplateNameContentsContentIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**templateName** | **string** |  | 
**contentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDTemplatesTemplateNameContentsContentIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDTemplatesTemplateNameContentsContentIDGet

> V1EnvironmentsEnvIDTemplatesTemplateNameContentsContentIDGet(ctx, envID, templateName, contentID).ContentType(contentType).Authorization(authorization).Execute()

READ One Content



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
    templateName := "templateName_example" // string | 
    contentID := "contentID_example" // string | 
    contentType := "application/json" // string |  (optional)
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsNotificationsNotificationsTemplatesApi.V1EnvironmentsEnvIDTemplatesTemplateNameContentsContentIDGet(context.Background(), envID, templateName, contentID).ContentType(contentType).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsNotificationsNotificationsTemplatesApi.V1EnvironmentsEnvIDTemplatesTemplateNameContentsContentIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**templateName** | **string** |  | 
**contentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDTemplatesTemplateNameContentsContentIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **contentType** | **string** |  | 
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


## V1EnvironmentsEnvIDTemplatesTemplateNameContentsContentIDPut

> V1EnvironmentsEnvIDTemplatesTemplateNameContentsContentIDPut(ctx, envID, templateName, contentID).ContentType(contentType).Authorization(authorization).Body(body).Execute()

UPDATE Push Content



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
    templateName := "templateName_example" // string | 
    contentID := "contentID_example" // string | 
    contentType := "application/json" // string |  (optional)
    authorization := "Bearer {{jwtToken}}" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsNotificationsNotificationsTemplatesApi.V1EnvironmentsEnvIDTemplatesTemplateNameContentsContentIDPut(context.Background(), envID, templateName, contentID).ContentType(contentType).Authorization(authorization).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsNotificationsNotificationsTemplatesApi.V1EnvironmentsEnvIDTemplatesTemplateNameContentsContentIDPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**templateName** | **string** |  | 
**contentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDTemplatesTemplateNameContentsContentIDPutRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDTemplatesTemplateNameContentsDelete

> V1EnvironmentsEnvIDTemplatesTemplateNameContentsDelete(ctx, envID, templateName).Authorization(authorization).Filter(filter).Execute()

DELETE Bulk Variant Contents



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
    templateName := "templateName_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)
    filter := "variant eq {{variantName}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsNotificationsNotificationsTemplatesApi.V1EnvironmentsEnvIDTemplatesTemplateNameContentsDelete(context.Background(), envID, templateName).Authorization(authorization).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsNotificationsNotificationsTemplatesApi.V1EnvironmentsEnvIDTemplatesTemplateNameContentsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**templateName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDTemplatesTemplateNameContentsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** |  | 
 **filter** | **string** |  | 

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


## V1EnvironmentsEnvIDTemplatesTemplateNameContentsGet

> V1EnvironmentsEnvIDTemplatesTemplateNameContentsGet(ctx, envID, templateName).Authorization(authorization).Execute()

READ All Contents



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
    templateName := "templateName_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsNotificationsNotificationsTemplatesApi.V1EnvironmentsEnvIDTemplatesTemplateNameContentsGet(context.Background(), envID, templateName).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsNotificationsNotificationsTemplatesApi.V1EnvironmentsEnvIDTemplatesTemplateNameContentsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**templateName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDTemplatesTemplateNameContentsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDTemplatesTemplateNameContentsPatch

> V1EnvironmentsEnvIDTemplatesTemplateNameContentsPatch(ctx, envID, templateName).Authorization(authorization).Filter(filter).Body(body).Execute()

PATCH Bulk Variant Contents



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
    templateName := "templateName_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)
    filter := "variant eq {{variantName}}" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsNotificationsNotificationsTemplatesApi.V1EnvironmentsEnvIDTemplatesTemplateNameContentsPatch(context.Background(), envID, templateName).Authorization(authorization).Filter(filter).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsNotificationsNotificationsTemplatesApi.V1EnvironmentsEnvIDTemplatesTemplateNameContentsPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**templateName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDTemplatesTemplateNameContentsPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** |  | 
 **filter** | **string** |  | 
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


## V1EnvironmentsEnvIDTemplatesTemplateNameContentsPost

> V1EnvironmentsEnvIDTemplatesTemplateNameContentsPost(ctx, envID, templateName).ContentType(contentType).Authorization(authorization).Body(body).Execute()

CREATE Push Content



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
    templateName := "templateName_example" // string | 
    contentType := "application/json" // string |  (optional)
    authorization := "Bearer {{jwtToken}}" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsNotificationsNotificationsTemplatesApi.V1EnvironmentsEnvIDTemplatesTemplateNameContentsPost(context.Background(), envID, templateName).ContentType(contentType).Authorization(authorization).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsNotificationsNotificationsTemplatesApi.V1EnvironmentsEnvIDTemplatesTemplateNameContentsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**templateName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDTemplatesTemplateNameContentsPostRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDTemplatesTemplateNameGet

> V1EnvironmentsEnvIDTemplatesTemplateNameGet(ctx, envID, templateName).Authorization(authorization).Execute()

READ One Template



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
    templateName := "templateName_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsNotificationsNotificationsTemplatesApi.V1EnvironmentsEnvIDTemplatesTemplateNameGet(context.Background(), envID, templateName).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsNotificationsNotificationsTemplatesApi.V1EnvironmentsEnvIDTemplatesTemplateNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**templateName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDTemplatesTemplateNameGetRequest struct via the builder pattern


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

