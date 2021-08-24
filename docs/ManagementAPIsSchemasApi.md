# \ManagementAPIsSchemasApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDSchemasGet**](ManagementAPIsSchemasApi.md#V1EnvironmentsEnvIDSchemasGet) | **Get** /v1/environments/{envID}/schemas | READ All Schemas
[**V1EnvironmentsEnvIDSchemasSchemaIDAttributesAttributeIDDelete**](ManagementAPIsSchemasApi.md#V1EnvironmentsEnvIDSchemasSchemaIDAttributesAttributeIDDelete) | **Delete** /v1/environments/{envID}/schemas/{schemaID}/attributes/{attributeID} | DELETE Attribute
[**V1EnvironmentsEnvIDSchemasSchemaIDAttributesAttributeIDGet**](ManagementAPIsSchemasApi.md#V1EnvironmentsEnvIDSchemasSchemaIDAttributesAttributeIDGet) | **Get** /v1/environments/{envID}/schemas/{schemaID}/attributes/{attributeID} | READ One Attribute
[**V1EnvironmentsEnvIDSchemasSchemaIDAttributesAttributeIDPatch**](ManagementAPIsSchemasApi.md#V1EnvironmentsEnvIDSchemasSchemaIDAttributesAttributeIDPatch) | **Patch** /v1/environments/{envID}/schemas/{schemaID}/attributes/{attributeID} | UPDATE Attribute (Patch)
[**V1EnvironmentsEnvIDSchemasSchemaIDAttributesAttributeIDPut**](ManagementAPIsSchemasApi.md#V1EnvironmentsEnvIDSchemasSchemaIDAttributesAttributeIDPut) | **Put** /v1/environments/{envID}/schemas/{schemaID}/attributes/{attributeID} | UPDATE Attribute (Put)
[**V1EnvironmentsEnvIDSchemasSchemaIDAttributesGet**](ManagementAPIsSchemasApi.md#V1EnvironmentsEnvIDSchemasSchemaIDAttributesGet) | **Get** /v1/environments/{envID}/schemas/{schemaID}/attributes | READ All (Schema) Attributes
[**V1EnvironmentsEnvIDSchemasSchemaIDAttributesPost**](ManagementAPIsSchemasApi.md#V1EnvironmentsEnvIDSchemasSchemaIDAttributesPost) | **Post** /v1/environments/{envID}/schemas/{schemaID}/attributes | CREATE Attribute
[**V1EnvironmentsEnvIDSchemasSchemaIDGet**](ManagementAPIsSchemasApi.md#V1EnvironmentsEnvIDSchemasSchemaIDGet) | **Get** /v1/environments/{envID}/schemas/{schemaID} | READ One Schema



## V1EnvironmentsEnvIDSchemasGet

> V1EnvironmentsEnvIDSchemasGet(ctx, envID).Authorization(authorization).Execute()

READ All Schemas



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
    resp, r, err := api_client.ManagementAPIsSchemasApi.V1EnvironmentsEnvIDSchemasGet(context.Background(), envID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsSchemasApi.V1EnvironmentsEnvIDSchemasGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDSchemasGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDSchemasSchemaIDAttributesAttributeIDDelete

> V1EnvironmentsEnvIDSchemasSchemaIDAttributesAttributeIDDelete(ctx, envID, schemaID, attributeID).ContentType(contentType).Authorization(authorization).Execute()

DELETE Attribute



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
    schemaID := "schemaID_example" // string | 
    attributeID := "attributeID_example" // string | 
    contentType := "application/json" // string |  (optional)
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsSchemasApi.V1EnvironmentsEnvIDSchemasSchemaIDAttributesAttributeIDDelete(context.Background(), envID, schemaID, attributeID).ContentType(contentType).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsSchemasApi.V1EnvironmentsEnvIDSchemasSchemaIDAttributesAttributeIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**schemaID** | **string** |  | 
**attributeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDSchemasSchemaIDAttributesAttributeIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDSchemasSchemaIDAttributesAttributeIDGet

> V1EnvironmentsEnvIDSchemasSchemaIDAttributesAttributeIDGet(ctx, envID, schemaID, attributeID).Authorization(authorization).Execute()

READ One Attribute



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
    schemaID := "schemaID_example" // string | 
    attributeID := "attributeID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsSchemasApi.V1EnvironmentsEnvIDSchemasSchemaIDAttributesAttributeIDGet(context.Background(), envID, schemaID, attributeID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsSchemasApi.V1EnvironmentsEnvIDSchemasSchemaIDAttributesAttributeIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**schemaID** | **string** |  | 
**attributeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDSchemasSchemaIDAttributesAttributeIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDSchemasSchemaIDAttributesAttributeIDPatch

> V1EnvironmentsEnvIDSchemasSchemaIDAttributesAttributeIDPatch(ctx, envID, schemaID, attributeID).ContentType(contentType).Authorization(authorization).Body(body).Execute()

UPDATE Attribute (Patch)



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
    schemaID := "schemaID_example" // string | 
    attributeID := "attributeID_example" // string | 
    contentType := "application/json" // string |  (optional)
    authorization := "Bearer {{jwtToken}}" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsSchemasApi.V1EnvironmentsEnvIDSchemasSchemaIDAttributesAttributeIDPatch(context.Background(), envID, schemaID, attributeID).ContentType(contentType).Authorization(authorization).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsSchemasApi.V1EnvironmentsEnvIDSchemasSchemaIDAttributesAttributeIDPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**schemaID** | **string** |  | 
**attributeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDSchemasSchemaIDAttributesAttributeIDPatchRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDSchemasSchemaIDAttributesAttributeIDPut

> V1EnvironmentsEnvIDSchemasSchemaIDAttributesAttributeIDPut(ctx, envID, schemaID, attributeID).ContentType(contentType).Authorization(authorization).Body(body).Execute()

UPDATE Attribute (Put)



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
    schemaID := "schemaID_example" // string | 
    attributeID := "attributeID_example" // string | 
    contentType := "application/json" // string |  (optional)
    authorization := "Bearer {{jwtToken}}" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsSchemasApi.V1EnvironmentsEnvIDSchemasSchemaIDAttributesAttributeIDPut(context.Background(), envID, schemaID, attributeID).ContentType(contentType).Authorization(authorization).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsSchemasApi.V1EnvironmentsEnvIDSchemasSchemaIDAttributesAttributeIDPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**schemaID** | **string** |  | 
**attributeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDSchemasSchemaIDAttributesAttributeIDPutRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDSchemasSchemaIDAttributesGet

> V1EnvironmentsEnvIDSchemasSchemaIDAttributesGet(ctx, envID, schemaID).Authorization(authorization).Execute()

READ All (Schema) Attributes



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
    schemaID := "schemaID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsSchemasApi.V1EnvironmentsEnvIDSchemasSchemaIDAttributesGet(context.Background(), envID, schemaID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsSchemasApi.V1EnvironmentsEnvIDSchemasSchemaIDAttributesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**schemaID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDSchemasSchemaIDAttributesGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDSchemasSchemaIDAttributesPost

> V1EnvironmentsEnvIDSchemasSchemaIDAttributesPost(ctx, envID, schemaID).ContentType(contentType).Authorization(authorization).Body(body).Execute()

CREATE Attribute



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
    schemaID := "schemaID_example" // string | 
    contentType := "application/json" // string |  (optional)
    authorization := "Bearer {{jwtToken}}" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsSchemasApi.V1EnvironmentsEnvIDSchemasSchemaIDAttributesPost(context.Background(), envID, schemaID).ContentType(contentType).Authorization(authorization).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsSchemasApi.V1EnvironmentsEnvIDSchemasSchemaIDAttributesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**schemaID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDSchemasSchemaIDAttributesPostRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDSchemasSchemaIDGet

> V1EnvironmentsEnvIDSchemasSchemaIDGet(ctx, envID, schemaID).Authorization(authorization).Execute()

READ One Schema



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
    schemaID := "schemaID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsSchemasApi.V1EnvironmentsEnvIDSchemasSchemaIDGet(context.Background(), envID, schemaID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsSchemasApi.V1EnvironmentsEnvIDSchemasSchemaIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**schemaID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDSchemasSchemaIDGetRequest struct via the builder pattern


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

