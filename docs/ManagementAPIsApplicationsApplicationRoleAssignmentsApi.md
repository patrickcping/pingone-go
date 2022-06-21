# \ManagementAPIsApplicationsApplicationRoleAssignmentsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplicationRoleAssignment**](ManagementAPIsApplicationsApplicationRoleAssignmentsApi.md#CreateApplicationRoleAssignment) | **Post** /v1/environments/{envID}/applications/{appID}/roleAssignments | CREATE Application Role Assignments
[**DeleteApplicationRoleAssignment**](ManagementAPIsApplicationsApplicationRoleAssignmentsApi.md#DeleteApplicationRoleAssignment) | **Delete** /v1/environments/{envID}/applications/{appID}/roleAssignments/{roleAssignmentID} | DELETE Application Role Assignment
[**ReadApplicationRoleAssignments**](ManagementAPIsApplicationsApplicationRoleAssignmentsApi.md#ReadApplicationRoleAssignments) | **Get** /v1/environments/{envID}/applications/{appID}/roleAssignments | READ Application Role Assignments
[**ReadOneApplicationRoleAssignment**](ManagementAPIsApplicationsApplicationRoleAssignmentsApi.md#ReadOneApplicationRoleAssignment) | **Get** /v1/environments/{envID}/applications/{appID}/roleAssignments/{roleAssignmentID} | READ One Application Role Assignment



## CreateApplicationRoleAssignment

> RoleAssignment CreateApplicationRoleAssignment(ctx, envID, appID).RoleAssignment(roleAssignment).Execute()

CREATE Application Role Assignments



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
    roleAssignment := *openapiclient.NewRoleAssignment(*openapiclient.NewRoleAssignmentRole("Id_example"), *openapiclient.NewRoleAssignmentScope("Id_example", "Type_example")) // RoleAssignment |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsApplicationsApplicationRoleAssignmentsApi.CreateApplicationRoleAssignment(context.Background(), envID, appID).RoleAssignment(roleAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationRoleAssignmentsApi.CreateApplicationRoleAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApplicationRoleAssignment`: RoleAssignment
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsApplicationsApplicationRoleAssignmentsApi.CreateApplicationRoleAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**appID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationRoleAssignmentRequest struct via the builder pattern


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


## DeleteApplicationRoleAssignment

> DeleteApplicationRoleAssignment(ctx, envID, appID, roleAssignmentID).Execute()

DELETE Application Role Assignment



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
    roleAssignmentID := "roleAssignmentID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsApplicationsApplicationRoleAssignmentsApi.DeleteApplicationRoleAssignment(context.Background(), envID, appID, roleAssignmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationRoleAssignmentsApi.DeleteApplicationRoleAssignment``: %v\n", err)
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
**roleAssignmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationRoleAssignmentRequest struct via the builder pattern


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


## ReadApplicationRoleAssignments

> EntityArray ReadApplicationRoleAssignments(ctx, envID, appID).Execute()

READ Application Role Assignments



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsApplicationsApplicationRoleAssignmentsApi.ReadApplicationRoleAssignments(context.Background(), envID, appID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationRoleAssignmentsApi.ReadApplicationRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadApplicationRoleAssignments`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsApplicationsApplicationRoleAssignmentsApi.ReadApplicationRoleAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**appID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadApplicationRoleAssignmentsRequest struct via the builder pattern


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


## ReadOneApplicationRoleAssignment

> RoleAssignment ReadOneApplicationRoleAssignment(ctx, envID, appID, roleAssignmentID).Execute()

READ One Application Role Assignment



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
    roleAssignmentID := "roleAssignmentID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsApplicationsApplicationRoleAssignmentsApi.ReadOneApplicationRoleAssignment(context.Background(), envID, appID, roleAssignmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationRoleAssignmentsApi.ReadOneApplicationRoleAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneApplicationRoleAssignment`: RoleAssignment
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsApplicationsApplicationRoleAssignmentsApi.ReadOneApplicationRoleAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**appID** | **string** |  | 
**roleAssignmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneApplicationRoleAssignmentRequest struct via the builder pattern


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

