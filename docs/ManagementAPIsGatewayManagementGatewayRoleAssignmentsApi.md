# \ManagementAPIsGatewayManagementGatewayRoleAssignmentsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDDelete**](ManagementAPIsGatewayManagementGatewayRoleAssignmentsApi.md#V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDDelete) | **Delete** /v1/environments/{envID}/gateways/{gatewayID}/roleAssignments/{gatewayRoleAssignmentsID} | DELETE Gateway Role Assignment
[**V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDGet**](ManagementAPIsGatewayManagementGatewayRoleAssignmentsApi.md#V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDGet) | **Get** /v1/environments/{envID}/gateways/{gatewayID}/roleAssignments/{gatewayRoleAssignmentsID} | READ One Gateway Role Assignment
[**V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDPut**](ManagementAPIsGatewayManagementGatewayRoleAssignmentsApi.md#V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDPut) | **Put** /v1/environments/{envID}/gateways/{gatewayID}/roleAssignments/{gatewayRoleAssignmentsID} | UPDATE Gateway Role Assignments
[**V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGet**](ManagementAPIsGatewayManagementGatewayRoleAssignmentsApi.md#V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGet) | **Get** /v1/environments/{envID}/gateways/{gatewayID}/roleAssignments | READ Gateway Role Assignments
[**V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsPost**](ManagementAPIsGatewayManagementGatewayRoleAssignmentsApi.md#V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsPost) | **Post** /v1/environments/{envID}/gateways/{gatewayID}/roleAssignments | CREATE Gateway Role Assignments



## V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDDelete

> V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDDelete(ctx, envID, gatewayID, gatewayRoleAssignmentsID).Execute()

DELETE Gateway Role Assignment



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
    gatewayID := "gatewayID_example" // string | 
    gatewayRoleAssignmentsID := "gatewayRoleAssignmentsID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsGatewayManagementGatewayRoleAssignmentsApi.V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDDelete(context.Background(), envID, gatewayID, gatewayRoleAssignmentsID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsGatewayManagementGatewayRoleAssignmentsApi.V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**gatewayID** | **string** |  | 
**gatewayRoleAssignmentsID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDGet

> V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDGet(ctx, envID, gatewayID, gatewayRoleAssignmentsID).Execute()

READ One Gateway Role Assignment



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
    gatewayID := "gatewayID_example" // string | 
    gatewayRoleAssignmentsID := "gatewayRoleAssignmentsID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsGatewayManagementGatewayRoleAssignmentsApi.V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDGet(context.Background(), envID, gatewayID, gatewayRoleAssignmentsID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsGatewayManagementGatewayRoleAssignmentsApi.V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**gatewayID** | **string** |  | 
**gatewayRoleAssignmentsID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDPut

> V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDPut(ctx, envID, gatewayID, gatewayRoleAssignmentsID).Body(body).Execute()

UPDATE Gateway Role Assignments



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
    gatewayID := "gatewayID_example" // string | 
    gatewayRoleAssignmentsID := "gatewayRoleAssignmentsID_example" // string | 
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsGatewayManagementGatewayRoleAssignmentsApi.V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDPut(context.Background(), envID, gatewayID, gatewayRoleAssignmentsID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsGatewayManagementGatewayRoleAssignmentsApi.V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**gatewayID** | **string** |  | 
**gatewayRoleAssignmentsID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGatewayRoleAssignmentsIDPutRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGet

> V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGet(ctx, envID, gatewayID).Execute()

READ Gateway Role Assignments



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
    gatewayID := "gatewayID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsGatewayManagementGatewayRoleAssignmentsApi.V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGet(context.Background(), envID, gatewayID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsGatewayManagementGatewayRoleAssignmentsApi.V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**gatewayID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsPost

> V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsPost(ctx, envID, gatewayID).Body(body).Execute()

CREATE Gateway Role Assignments



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
    gatewayID := "gatewayID_example" // string | 
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsGatewayManagementGatewayRoleAssignmentsApi.V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsPost(context.Background(), envID, gatewayID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsGatewayManagementGatewayRoleAssignmentsApi.V1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**gatewayID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDGatewaysGatewayIDRoleAssignmentsPostRequest struct via the builder pattern


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

