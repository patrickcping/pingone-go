# \ManagementAPIsApplicationsApplicationRoleAssignmentsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDApplicationsAppIDRoleAssignmentsGet**](ManagementAPIsApplicationsApplicationRoleAssignmentsApi.md#V1EnvironmentsEnvIDApplicationsAppIDRoleAssignmentsGet) | **Get** /v1/environments/{envID}/applications/{appID}/roleAssignments | READ Application Role Assignments
[**V1EnvironmentsEnvIDApplicationsAppIDRoleAssignmentsPost**](ManagementAPIsApplicationsApplicationRoleAssignmentsApi.md#V1EnvironmentsEnvIDApplicationsAppIDRoleAssignmentsPost) | **Post** /v1/environments/{envID}/applications/{appID}/roleAssignments | CREATE Application Role Assignments
[**V1EnvironmentsEnvIDApplicationsAppIDRoleAssignmentsRoleAssignmentIDDelete**](ManagementAPIsApplicationsApplicationRoleAssignmentsApi.md#V1EnvironmentsEnvIDApplicationsAppIDRoleAssignmentsRoleAssignmentIDDelete) | **Delete** /v1/environments/{envID}/applications/{appID}/roleAssignments/{roleAssignmentID} | DELETE Application Role Assignment
[**V1EnvironmentsEnvIDApplicationsAppIDRoleAssignmentsRoleAssignmentIDGet**](ManagementAPIsApplicationsApplicationRoleAssignmentsApi.md#V1EnvironmentsEnvIDApplicationsAppIDRoleAssignmentsRoleAssignmentIDGet) | **Get** /v1/environments/{envID}/applications/{appID}/roleAssignments/{roleAssignmentID} | READ One Application Role Assignment



## V1EnvironmentsEnvIDApplicationsAppIDRoleAssignmentsGet

> V1EnvironmentsEnvIDApplicationsAppIDRoleAssignmentsGet(ctx, envID, appID).Authorization(authorization).Execute()

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
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsApplicationsApplicationRoleAssignmentsApi.V1EnvironmentsEnvIDApplicationsAppIDRoleAssignmentsGet(context.Background(), envID, appID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationRoleAssignmentsApi.V1EnvironmentsEnvIDApplicationsAppIDRoleAssignmentsGet``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDApplicationsAppIDRoleAssignmentsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDApplicationsAppIDRoleAssignmentsPost

> V1EnvironmentsEnvIDApplicationsAppIDRoleAssignmentsPost(ctx, envID, appID).ContentType(contentType).Authorization(authorization).Body(body).Execute()

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
    contentType := "application/json" // string |  (optional)
    authorization := "Bearer {{jwtToken}}" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsApplicationsApplicationRoleAssignmentsApi.V1EnvironmentsEnvIDApplicationsAppIDRoleAssignmentsPost(context.Background(), envID, appID).ContentType(contentType).Authorization(authorization).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationRoleAssignmentsApi.V1EnvironmentsEnvIDApplicationsAppIDRoleAssignmentsPost``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDApplicationsAppIDRoleAssignmentsPostRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDApplicationsAppIDRoleAssignmentsRoleAssignmentIDDelete

> V1EnvironmentsEnvIDApplicationsAppIDRoleAssignmentsRoleAssignmentIDDelete(ctx, envID, appID, roleAssignmentID).Authorization(authorization).Execute()

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
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsApplicationsApplicationRoleAssignmentsApi.V1EnvironmentsEnvIDApplicationsAppIDRoleAssignmentsRoleAssignmentIDDelete(context.Background(), envID, appID, roleAssignmentID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationRoleAssignmentsApi.V1EnvironmentsEnvIDApplicationsAppIDRoleAssignmentsRoleAssignmentIDDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDApplicationsAppIDRoleAssignmentsRoleAssignmentIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDApplicationsAppIDRoleAssignmentsRoleAssignmentIDGet

> V1EnvironmentsEnvIDApplicationsAppIDRoleAssignmentsRoleAssignmentIDGet(ctx, envID, appID, roleAssignmentID).Authorization(authorization).Execute()

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
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsApplicationsApplicationRoleAssignmentsApi.V1EnvironmentsEnvIDApplicationsAppIDRoleAssignmentsRoleAssignmentIDGet(context.Background(), envID, appID, roleAssignmentID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationRoleAssignmentsApi.V1EnvironmentsEnvIDApplicationsAppIDRoleAssignmentsRoleAssignmentIDGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDApplicationsAppIDRoleAssignmentsRoleAssignmentIDGetRequest struct via the builder pattern


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

