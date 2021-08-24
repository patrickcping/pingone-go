# \ManagementAPIsBrandingBrandingThemesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDThemesGet**](ManagementAPIsBrandingBrandingThemesApi.md#V1EnvironmentsEnvIDThemesGet) | **Get** /v1/environments/{envID}/themes | READ Branding Themes
[**V1EnvironmentsEnvIDThemesPost**](ManagementAPIsBrandingBrandingThemesApi.md#V1EnvironmentsEnvIDThemesPost) | **Post** /v1/environments/{envID}/themes | CREATE Branding Theme
[**V1EnvironmentsEnvIDThemesThemeIDDefaultGet**](ManagementAPIsBrandingBrandingThemesApi.md#V1EnvironmentsEnvIDThemesThemeIDDefaultGet) | **Get** /v1/environments/{envID}/themes/{themeID}/default | READ Branding Theme Default
[**V1EnvironmentsEnvIDThemesThemeIDDefaultPut**](ManagementAPIsBrandingBrandingThemesApi.md#V1EnvironmentsEnvIDThemesThemeIDDefaultPut) | **Put** /v1/environments/{envID}/themes/{themeID}/default | UPDATE Branding Theme Default
[**V1EnvironmentsEnvIDThemesThemeIDDelete**](ManagementAPIsBrandingBrandingThemesApi.md#V1EnvironmentsEnvIDThemesThemeIDDelete) | **Delete** /v1/environments/{envID}/themes/{themeID} | DELETE Branding Theme
[**V1EnvironmentsEnvIDThemesThemeIDGet**](ManagementAPIsBrandingBrandingThemesApi.md#V1EnvironmentsEnvIDThemesThemeIDGet) | **Get** /v1/environments/{envID}/themes/{themeID} | READ One Branding Theme
[**V1EnvironmentsEnvIDThemesThemeIDPut**](ManagementAPIsBrandingBrandingThemesApi.md#V1EnvironmentsEnvIDThemesThemeIDPut) | **Put** /v1/environments/{envID}/themes/{themeID} | UPDATE Branding Theme



## V1EnvironmentsEnvIDThemesGet

> V1EnvironmentsEnvIDThemesGet(ctx, envID).Authorization(authorization).Execute()

READ Branding Themes



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
    authorization := "{{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsBrandingBrandingThemesApi.V1EnvironmentsEnvIDThemesGet(context.Background(), envID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsBrandingBrandingThemesApi.V1EnvironmentsEnvIDThemesGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDThemesGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDThemesPost

> V1EnvironmentsEnvIDThemesPost(ctx, envID).Authorization(authorization).ContentType(contentType).Body(body).Execute()

CREATE Branding Theme



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
    authorization := "{{jwtToken}}" // string |  (optional)
    contentType := "application/json" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsBrandingBrandingThemesApi.V1EnvironmentsEnvIDThemesPost(context.Background(), envID).Authorization(authorization).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsBrandingBrandingThemesApi.V1EnvironmentsEnvIDThemesPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDThemesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 
 **contentType** | **string** |  | 
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


## V1EnvironmentsEnvIDThemesThemeIDDefaultGet

> V1EnvironmentsEnvIDThemesThemeIDDefaultGet(ctx, envID, themeID).Authorization(authorization).Execute()

READ Branding Theme Default



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
    themeID := "themeID_example" // string | 
    authorization := "{{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsBrandingBrandingThemesApi.V1EnvironmentsEnvIDThemesThemeIDDefaultGet(context.Background(), envID, themeID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsBrandingBrandingThemesApi.V1EnvironmentsEnvIDThemesThemeIDDefaultGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**themeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDThemesThemeIDDefaultGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDThemesThemeIDDefaultPut

> V1EnvironmentsEnvIDThemesThemeIDDefaultPut(ctx, envID, themeID).Authorization(authorization).ContentType(contentType).Body(body).Execute()

UPDATE Branding Theme Default



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
    themeID := "themeID_example" // string | 
    authorization := "{{jwtToken}}" // string |  (optional)
    contentType := "application/json" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsBrandingBrandingThemesApi.V1EnvironmentsEnvIDThemesThemeIDDefaultPut(context.Background(), envID, themeID).Authorization(authorization).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsBrandingBrandingThemesApi.V1EnvironmentsEnvIDThemesThemeIDDefaultPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**themeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDThemesThemeIDDefaultPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** |  | 
 **contentType** | **string** |  | 
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


## V1EnvironmentsEnvIDThemesThemeIDDelete

> V1EnvironmentsEnvIDThemesThemeIDDelete(ctx, envID, themeID).Authorization(authorization).Execute()

DELETE Branding Theme



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
    themeID := "themeID_example" // string | 
    authorization := "{{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsBrandingBrandingThemesApi.V1EnvironmentsEnvIDThemesThemeIDDelete(context.Background(), envID, themeID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsBrandingBrandingThemesApi.V1EnvironmentsEnvIDThemesThemeIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**themeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDThemesThemeIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDThemesThemeIDGet

> V1EnvironmentsEnvIDThemesThemeIDGet(ctx, envID, themeID).Authorization(authorization).Execute()

READ One Branding Theme



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
    themeID := "themeID_example" // string | 
    authorization := "{{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsBrandingBrandingThemesApi.V1EnvironmentsEnvIDThemesThemeIDGet(context.Background(), envID, themeID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsBrandingBrandingThemesApi.V1EnvironmentsEnvIDThemesThemeIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**themeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDThemesThemeIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDThemesThemeIDPut

> V1EnvironmentsEnvIDThemesThemeIDPut(ctx, envID, themeID).Authorization(authorization).ContentType(contentType).Body(body).Execute()

UPDATE Branding Theme



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
    themeID := "themeID_example" // string | 
    authorization := "{{jwtToken}}" // string |  (optional)
    contentType := "application/json" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsBrandingBrandingThemesApi.V1EnvironmentsEnvIDThemesThemeIDPut(context.Background(), envID, themeID).Authorization(authorization).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsBrandingBrandingThemesApi.V1EnvironmentsEnvIDThemesThemeIDPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**themeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDThemesThemeIDPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** |  | 
 **contentType** | **string** |  | 
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

