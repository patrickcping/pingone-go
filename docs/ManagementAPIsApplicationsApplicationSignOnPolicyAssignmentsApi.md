# \ManagementAPIsApplicationsApplicationSignOnPolicyAssignmentsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDApplicationsAppIDSignOnPolicyAssignmentsGet**](ManagementAPIsApplicationsApplicationSignOnPolicyAssignmentsApi.md#V1EnvironmentsEnvIDApplicationsAppIDSignOnPolicyAssignmentsGet) | **Get** /v1/environments/{envID}/applications/{appID}/signOnPolicyAssignments | READ All SOP Assignments
[**V1EnvironmentsEnvIDApplicationsAppIDSignOnPolicyAssignmentsPost**](ManagementAPIsApplicationsApplicationSignOnPolicyAssignmentsApi.md#V1EnvironmentsEnvIDApplicationsAppIDSignOnPolicyAssignmentsPost) | **Post** /v1/environments/{envID}/applications/{appID}/signOnPolicyAssignments | CREATE SOP Assignment
[**V1EnvironmentsEnvIDApplicationsAppIDSignOnPolicyAssignmentsSOPAssignmentIDDelete**](ManagementAPIsApplicationsApplicationSignOnPolicyAssignmentsApi.md#V1EnvironmentsEnvIDApplicationsAppIDSignOnPolicyAssignmentsSOPAssignmentIDDelete) | **Delete** /v1/environments/{envID}/applications/{appID}/signOnPolicyAssignments/{SOPAssignmentID} | DELETE SOP Assignment
[**V1EnvironmentsEnvIDApplicationsAppIDSignOnPolicyAssignmentsSOPAssignmentIDGet**](ManagementAPIsApplicationsApplicationSignOnPolicyAssignmentsApi.md#V1EnvironmentsEnvIDApplicationsAppIDSignOnPolicyAssignmentsSOPAssignmentIDGet) | **Get** /v1/environments/{envID}/applications/{appID}/signOnPolicyAssignments/{SOPAssignmentID} | READ One SOP Assignment
[**V1EnvironmentsEnvIDApplicationsAppIDSignOnPolicyAssignmentsSOPAssignmentIDPut**](ManagementAPIsApplicationsApplicationSignOnPolicyAssignmentsApi.md#V1EnvironmentsEnvIDApplicationsAppIDSignOnPolicyAssignmentsSOPAssignmentIDPut) | **Put** /v1/environments/{envID}/applications/{appID}/signOnPolicyAssignments/{SOPAssignmentID} | UPDATE SOP Assignment



## V1EnvironmentsEnvIDApplicationsAppIDSignOnPolicyAssignmentsGet

> V1EnvironmentsEnvIDApplicationsAppIDSignOnPolicyAssignmentsGet(ctx, envID, appID).Execute()

READ All SOP Assignments



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
    resp, r, err := apiClient.ManagementAPIsApplicationsApplicationSignOnPolicyAssignmentsApi.V1EnvironmentsEnvIDApplicationsAppIDSignOnPolicyAssignmentsGet(context.Background(), envID, appID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationSignOnPolicyAssignmentsApi.V1EnvironmentsEnvIDApplicationsAppIDSignOnPolicyAssignmentsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDApplicationsAppIDSignOnPolicyAssignmentsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDApplicationsAppIDSignOnPolicyAssignmentsPost

> V1EnvironmentsEnvIDApplicationsAppIDSignOnPolicyAssignmentsPost(ctx, envID, appID).Body(body).Execute()

CREATE SOP Assignment



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
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsApplicationsApplicationSignOnPolicyAssignmentsApi.V1EnvironmentsEnvIDApplicationsAppIDSignOnPolicyAssignmentsPost(context.Background(), envID, appID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationSignOnPolicyAssignmentsApi.V1EnvironmentsEnvIDApplicationsAppIDSignOnPolicyAssignmentsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDApplicationsAppIDSignOnPolicyAssignmentsPostRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDApplicationsAppIDSignOnPolicyAssignmentsSOPAssignmentIDDelete

> V1EnvironmentsEnvIDApplicationsAppIDSignOnPolicyAssignmentsSOPAssignmentIDDelete(ctx, envID, appID, sOPAssignmentID).Execute()

DELETE SOP Assignment



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
    sOPAssignmentID := "sOPAssignmentID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsApplicationsApplicationSignOnPolicyAssignmentsApi.V1EnvironmentsEnvIDApplicationsAppIDSignOnPolicyAssignmentsSOPAssignmentIDDelete(context.Background(), envID, appID, sOPAssignmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationSignOnPolicyAssignmentsApi.V1EnvironmentsEnvIDApplicationsAppIDSignOnPolicyAssignmentsSOPAssignmentIDDelete``: %v\n", err)
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
**sOPAssignmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDApplicationsAppIDSignOnPolicyAssignmentsSOPAssignmentIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDApplicationsAppIDSignOnPolicyAssignmentsSOPAssignmentIDGet

> V1EnvironmentsEnvIDApplicationsAppIDSignOnPolicyAssignmentsSOPAssignmentIDGet(ctx, envID, appID, sOPAssignmentID).Execute()

READ One SOP Assignment



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
    sOPAssignmentID := "sOPAssignmentID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsApplicationsApplicationSignOnPolicyAssignmentsApi.V1EnvironmentsEnvIDApplicationsAppIDSignOnPolicyAssignmentsSOPAssignmentIDGet(context.Background(), envID, appID, sOPAssignmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationSignOnPolicyAssignmentsApi.V1EnvironmentsEnvIDApplicationsAppIDSignOnPolicyAssignmentsSOPAssignmentIDGet``: %v\n", err)
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
**sOPAssignmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDApplicationsAppIDSignOnPolicyAssignmentsSOPAssignmentIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDApplicationsAppIDSignOnPolicyAssignmentsSOPAssignmentIDPut

> V1EnvironmentsEnvIDApplicationsAppIDSignOnPolicyAssignmentsSOPAssignmentIDPut(ctx, envID, appID, sOPAssignmentID).Body(body).Execute()

UPDATE SOP Assignment



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
    sOPAssignmentID := "sOPAssignmentID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsApplicationsApplicationSignOnPolicyAssignmentsApi.V1EnvironmentsEnvIDApplicationsAppIDSignOnPolicyAssignmentsSOPAssignmentIDPut(context.Background(), envID, appID, sOPAssignmentID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationSignOnPolicyAssignmentsApi.V1EnvironmentsEnvIDApplicationsAppIDSignOnPolicyAssignmentsSOPAssignmentIDPut``: %v\n", err)
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
**sOPAssignmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDApplicationsAppIDSignOnPolicyAssignmentsSOPAssignmentIDPutRequest struct via the builder pattern


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

