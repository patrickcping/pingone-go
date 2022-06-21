# \ManagementAPIsGatewayManagementGatewayRoleAssignmentsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGatewayRoleAssignment**](ManagementAPIsGatewayManagementGatewayRoleAssignmentsApi.md#CreateGatewayRoleAssignment) | **Post** /v1/environments/{envID}/gateways/{gatewayID}/roleAssignments | CREATE Gateway Role Assignments
[**DeleteGatewayRoleAssignment**](ManagementAPIsGatewayManagementGatewayRoleAssignmentsApi.md#DeleteGatewayRoleAssignment) | **Delete** /v1/environments/{envID}/gateways/{gatewayID}/roleAssignments/{gatewayRoleAssignmentID} | DELETE Gateway Role Assignment
[**ReadGatewayRoleAssignments**](ManagementAPIsGatewayManagementGatewayRoleAssignmentsApi.md#ReadGatewayRoleAssignments) | **Get** /v1/environments/{envID}/gateways/{gatewayID}/roleAssignments | READ Gateway Role Assignments
[**ReadOneGatewayRoleAssignment**](ManagementAPIsGatewayManagementGatewayRoleAssignmentsApi.md#ReadOneGatewayRoleAssignment) | **Get** /v1/environments/{envID}/gateways/{gatewayID}/roleAssignments/{gatewayRoleAssignmentID} | READ One Gateway Role Assignment
[**UpdateGatewayRoleAssignment**](ManagementAPIsGatewayManagementGatewayRoleAssignmentsApi.md#UpdateGatewayRoleAssignment) | **Put** /v1/environments/{envID}/gateways/{gatewayID}/roleAssignments/{gatewayRoleAssignmentID} | UPDATE Gateway Role Assignments



## CreateGatewayRoleAssignment

> RoleAssignment CreateGatewayRoleAssignment(ctx, envID, gatewayID).RoleAssignment(roleAssignment).Execute()

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
    roleAssignment := *openapiclient.NewRoleAssignment(*openapiclient.NewRoleAssignmentRole("Id_example"), *openapiclient.NewRoleAssignmentScope("Id_example", "Type_example")) // RoleAssignment |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsGatewayManagementGatewayRoleAssignmentsApi.CreateGatewayRoleAssignment(context.Background(), envID, gatewayID).RoleAssignment(roleAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsGatewayManagementGatewayRoleAssignmentsApi.CreateGatewayRoleAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGatewayRoleAssignment`: RoleAssignment
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsGatewayManagementGatewayRoleAssignmentsApi.CreateGatewayRoleAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**gatewayID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGatewayRoleAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **roleAssignment** | [**RoleAssignment**](RoleAssignment.md) |  | 

### Return type

[**RoleAssignment**](RoleAssignment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGatewayRoleAssignment

> DeleteGatewayRoleAssignment(ctx, envID, gatewayID, gatewayRoleAssignmentID).Execute()

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
    gatewayRoleAssignmentID := "gatewayRoleAssignmentID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsGatewayManagementGatewayRoleAssignmentsApi.DeleteGatewayRoleAssignment(context.Background(), envID, gatewayID, gatewayRoleAssignmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsGatewayManagementGatewayRoleAssignmentsApi.DeleteGatewayRoleAssignment``: %v\n", err)
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
**gatewayRoleAssignmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGatewayRoleAssignmentRequest struct via the builder pattern


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


## ReadGatewayRoleAssignments

> EntityArray ReadGatewayRoleAssignments(ctx, envID, gatewayID).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsGatewayManagementGatewayRoleAssignmentsApi.ReadGatewayRoleAssignments(context.Background(), envID, gatewayID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsGatewayManagementGatewayRoleAssignmentsApi.ReadGatewayRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadGatewayRoleAssignments`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsGatewayManagementGatewayRoleAssignmentsApi.ReadGatewayRoleAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**gatewayID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadGatewayRoleAssignmentsRequest struct via the builder pattern


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


## ReadOneGatewayRoleAssignment

> RoleAssignment ReadOneGatewayRoleAssignment(ctx, envID, gatewayID, gatewayRoleAssignmentID).Execute()

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
    gatewayRoleAssignmentID := "gatewayRoleAssignmentID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsGatewayManagementGatewayRoleAssignmentsApi.ReadOneGatewayRoleAssignment(context.Background(), envID, gatewayID, gatewayRoleAssignmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsGatewayManagementGatewayRoleAssignmentsApi.ReadOneGatewayRoleAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneGatewayRoleAssignment`: RoleAssignment
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsGatewayManagementGatewayRoleAssignmentsApi.ReadOneGatewayRoleAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**gatewayID** | **string** |  | 
**gatewayRoleAssignmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneGatewayRoleAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RoleAssignment**](RoleAssignment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGatewayRoleAssignment

> RoleAssignment UpdateGatewayRoleAssignment(ctx, envID, gatewayID, gatewayRoleAssignmentID).Body(body).Execute()

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
    gatewayRoleAssignmentID := "gatewayRoleAssignmentID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsGatewayManagementGatewayRoleAssignmentsApi.UpdateGatewayRoleAssignment(context.Background(), envID, gatewayID, gatewayRoleAssignmentID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsGatewayManagementGatewayRoleAssignmentsApi.UpdateGatewayRoleAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGatewayRoleAssignment`: RoleAssignment
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsGatewayManagementGatewayRoleAssignmentsApi.UpdateGatewayRoleAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**gatewayID** | **string** |  | 
**gatewayRoleAssignmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGatewayRoleAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **map[string]interface{}** |  | 

### Return type

[**RoleAssignment**](RoleAssignment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

