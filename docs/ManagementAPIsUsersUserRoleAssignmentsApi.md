# \ManagementAPIsUsersUserRoleAssignmentsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsGet**](ManagementAPIsUsersUserRoleAssignmentsApi.md#V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsGet) | **Get** /v1/environments/{envID}/users/{userID}/roleAssignments | READ Role Assignments
[**V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsPost**](ManagementAPIsUsersUserRoleAssignmentsApi.md#V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsPost) | **Post** /v1/environments/{envID}/users/{userID}/roleAssignments | CREATE User Role Assignment
[**V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsRoleAssignmentIDDelete**](ManagementAPIsUsersUserRoleAssignmentsApi.md#V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsRoleAssignmentIDDelete) | **Delete** /v1/environments/{envID}/users/{userID}/roleAssignments/{roleAssignmentID} | DELETE User&#39;s Role Assignment
[**V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsRoleAssignmentIDGet**](ManagementAPIsUsersUserRoleAssignmentsApi.md#V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsRoleAssignmentIDGet) | **Get** /v1/environments/{envID}/users/{userID}/roleAssignments/{roleAssignmentID} | READ One Role Assignment



## V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsGet

> V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsGet(ctx, envID, userID).Execute()

READ Role Assignments



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
    userID := "userID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsUsersUserRoleAssignmentsApi.V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsGet(context.Background(), envID, userID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsUsersUserRoleAssignmentsApi.V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**userID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUsersUserIDRoleAssignmentsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsPost

> V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsPost(ctx, envID, userID).ContentType(contentType).Body(body).Execute()

CREATE User Role Assignment



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
    userID := "userID_example" // string | 
    contentType := "application/json" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsUsersUserRoleAssignmentsApi.V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsPost(context.Background(), envID, userID).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsUsersUserRoleAssignmentsApi.V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**userID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUsersUserIDRoleAssignmentsPostRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsRoleAssignmentIDDelete

> V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsRoleAssignmentIDDelete(ctx, envID, userID, roleAssignmentID).Execute()

DELETE User's Role Assignment



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
    userID := "userID_example" // string | 
    roleAssignmentID := "roleAssignmentID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsUsersUserRoleAssignmentsApi.V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsRoleAssignmentIDDelete(context.Background(), envID, userID, roleAssignmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsUsersUserRoleAssignmentsApi.V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsRoleAssignmentIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**userID** | **string** |  | 
**roleAssignmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUsersUserIDRoleAssignmentsRoleAssignmentIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsRoleAssignmentIDGet

> V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsRoleAssignmentIDGet(ctx, envID, userID, roleAssignmentID).Execute()

READ One Role Assignment



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
    userID := "userID_example" // string | 
    roleAssignmentID := "roleAssignmentID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsUsersUserRoleAssignmentsApi.V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsRoleAssignmentIDGet(context.Background(), envID, userID, roleAssignmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsUsersUserRoleAssignmentsApi.V1EnvironmentsEnvIDUsersUserIDRoleAssignmentsRoleAssignmentIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**userID** | **string** |  | 
**roleAssignmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUsersUserIDRoleAssignmentsRoleAssignmentIDGetRequest struct via the builder pattern


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

