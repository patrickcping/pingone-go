# \ManagementAPIsSchemasApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAttribute**](ManagementAPIsSchemasApi.md#CreateAttribute) | **Post** /v1/environments/{envID}/schemas/{schemaID}/attributes | CREATE Attribute
[**DeleteAttribute**](ManagementAPIsSchemasApi.md#DeleteAttribute) | **Delete** /v1/environments/{envID}/schemas/{schemaID}/attributes/{attributeID} | DELETE Attribute
[**ReadAllSchemaAttributes**](ManagementAPIsSchemasApi.md#ReadAllSchemaAttributes) | **Get** /v1/environments/{envID}/schemas/{schemaID}/attributes | READ All (Schema) Attributes
[**ReadAllSchemas**](ManagementAPIsSchemasApi.md#ReadAllSchemas) | **Get** /v1/environments/{envID}/schemas | READ All Schemas
[**ReadOneAttribute**](ManagementAPIsSchemasApi.md#ReadOneAttribute) | **Get** /v1/environments/{envID}/schemas/{schemaID}/attributes/{attributeID} | READ One Attribute
[**ReadOneSchema**](ManagementAPIsSchemasApi.md#ReadOneSchema) | **Get** /v1/environments/{envID}/schemas/{schemaID} | READ One Schema
[**UpdateAttributePatch**](ManagementAPIsSchemasApi.md#UpdateAttributePatch) | **Patch** /v1/environments/{envID}/schemas/{schemaID}/attributes/{attributeID} | UPDATE Attribute (Patch)
[**UpdateAttributePut**](ManagementAPIsSchemasApi.md#UpdateAttributePut) | **Put** /v1/environments/{envID}/schemas/{schemaID}/attributes/{attributeID} | UPDATE Attribute (Put)



## CreateAttribute

> SchemaAttribute CreateAttribute(ctx, envID, schemaID).ContentType(contentType).SchemaAttribute(schemaAttribute).Execute()

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
    schemaAttribute := *openapiclient.NewSchemaAttribute() // SchemaAttribute |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsSchemasApi.CreateAttribute(context.Background(), envID, schemaID).ContentType(contentType).SchemaAttribute(schemaAttribute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsSchemasApi.CreateAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAttribute`: SchemaAttribute
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsSchemasApi.CreateAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**schemaID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contentType** | **string** |  | 
 **schemaAttribute** | [**SchemaAttribute**](SchemaAttribute.md) |  | 

### Return type

[**SchemaAttribute**](SchemaAttribute.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAttribute

> DeleteAttribute(ctx, envID, schemaID, attributeID).ContentType(contentType).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsSchemasApi.DeleteAttribute(context.Background(), envID, schemaID, attributeID).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsSchemasApi.DeleteAttribute``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAttributeRequest struct via the builder pattern


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


## ReadAllSchemaAttributes

> EntityArray ReadAllSchemaAttributes(ctx, envID, schemaID).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsSchemasApi.ReadAllSchemaAttributes(context.Background(), envID, schemaID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsSchemasApi.ReadAllSchemaAttributes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllSchemaAttributes`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsSchemasApi.ReadAllSchemaAttributes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**schemaID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllSchemaAttributesRequest struct via the builder pattern


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


## ReadAllSchemas

> EntityArray ReadAllSchemas(ctx, envID).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsSchemasApi.ReadAllSchemas(context.Background(), envID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsSchemasApi.ReadAllSchemas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllSchemas`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsSchemasApi.ReadAllSchemas`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllSchemasRequest struct via the builder pattern


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


## ReadOneAttribute

> SchemaAttribute ReadOneAttribute(ctx, envID, schemaID, attributeID).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsSchemasApi.ReadOneAttribute(context.Background(), envID, schemaID, attributeID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsSchemasApi.ReadOneAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneAttribute`: SchemaAttribute
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsSchemasApi.ReadOneAttribute`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiReadOneAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**SchemaAttribute**](SchemaAttribute.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadOneSchema

> Schema ReadOneSchema(ctx, envID, schemaID).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsSchemasApi.ReadOneSchema(context.Background(), envID, schemaID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsSchemasApi.ReadOneSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneSchema`: Schema
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsSchemasApi.ReadOneSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**schemaID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Schema**](Schema.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAttributePatch

> UpdateAttributePatch(ctx, envID, schemaID, attributeID).ContentType(contentType).SchemaAttribute(schemaAttribute).Execute()

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
    schemaAttribute := *openapiclient.NewSchemaAttribute() // SchemaAttribute |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsSchemasApi.UpdateAttributePatch(context.Background(), envID, schemaID, attributeID).ContentType(contentType).SchemaAttribute(schemaAttribute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsSchemasApi.UpdateAttributePatch``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateAttributePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **contentType** | **string** |  | 
 **schemaAttribute** | [**SchemaAttribute**](SchemaAttribute.md) |  | 

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


## UpdateAttributePut

> UpdateAttributePut(ctx, envID, schemaID, attributeID).ContentType(contentType).SchemaAttribute(schemaAttribute).Execute()

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
    schemaAttribute := *openapiclient.NewSchemaAttribute() // SchemaAttribute |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsSchemasApi.UpdateAttributePut(context.Background(), envID, schemaID, attributeID).ContentType(contentType).SchemaAttribute(schemaAttribute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsSchemasApi.UpdateAttributePut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateAttributePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **contentType** | **string** |  | 
 **schemaAttribute** | [**SchemaAttribute**](SchemaAttribute.md) |  | 

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

