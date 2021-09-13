# \ManagementAPIsUsersGroupMembershipApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadAllGroupIDsForUser**](ManagementAPIsUsersGroupMembershipApi.md#ReadAllGroupIDsForUser) | **Get** /v1/environments/{envID}/users/{userID} | READ All Group IDs for User
[**V1EnvironmentsEnvIDUsersUserIDMemberOfGroupsGet**](ManagementAPIsUsersGroupMembershipApi.md#V1EnvironmentsEnvIDUsersUserIDMemberOfGroupsGet) | **Get** /v1/environments/{envID}/users/{userID}/memberOfGroups | READ All Group Memberships for User
[**V1EnvironmentsEnvIDUsersUserIDMemberOfGroupsGroupIDDelete**](ManagementAPIsUsersGroupMembershipApi.md#V1EnvironmentsEnvIDUsersUserIDMemberOfGroupsGroupIDDelete) | **Delete** /v1/environments/{envID}/users/{userID}/memberOfGroups/{groupID} | REMOVE User from Group
[**V1EnvironmentsEnvIDUsersUserIDMemberOfGroupsGroupIDGet**](ManagementAPIsUsersGroupMembershipApi.md#V1EnvironmentsEnvIDUsersUserIDMemberOfGroupsGroupIDGet) | **Get** /v1/environments/{envID}/users/{userID}/memberOfGroups/{groupID} | READ One Group Membership for User
[**V1EnvironmentsEnvIDUsersUserIDMemberOfGroupsPost**](ManagementAPIsUsersGroupMembershipApi.md#V1EnvironmentsEnvIDUsersUserIDMemberOfGroupsPost) | **Post** /v1/environments/{envID}/users/{userID}/memberOfGroups | ADD User to Group



## ReadAllGroupIDsForUser

> ReadAllGroupIDsForUser(ctx, envID, userID).Include(include).Execute()

READ All Group IDs for User



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
    include := "memberOfGroupIDs" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsUsersGroupMembershipApi.ReadAllGroupIDsForUser(context.Background(), envID, userID).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsUsersGroupMembershipApi.ReadAllGroupIDsForUser``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReadAllGroupIDsForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **include** | **string** |  | 

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


## V1EnvironmentsEnvIDUsersUserIDMemberOfGroupsGet

> V1EnvironmentsEnvIDUsersUserIDMemberOfGroupsGet(ctx, envID, userID).Expand(expand).Limit(limit).Filter(filter).Execute()

READ All Group Memberships for User



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
    expand := "group" // string |  (optional)
    limit := int32(100) // int32 |  (optional)
    filter := "type eq "DIRECT"" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsUsersGroupMembershipApi.V1EnvironmentsEnvIDUsersUserIDMemberOfGroupsGet(context.Background(), envID, userID).Expand(expand).Limit(limit).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsUsersGroupMembershipApi.V1EnvironmentsEnvIDUsersUserIDMemberOfGroupsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUsersUserIDMemberOfGroupsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **string** |  | 
 **limit** | **int32** |  | 
 **filter** | **string** |  | 

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


## V1EnvironmentsEnvIDUsersUserIDMemberOfGroupsGroupIDDelete

> V1EnvironmentsEnvIDUsersUserIDMemberOfGroupsGroupIDDelete(ctx, envID, userID, groupID).Execute()

REMOVE User from Group



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
    groupID := "groupID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsUsersGroupMembershipApi.V1EnvironmentsEnvIDUsersUserIDMemberOfGroupsGroupIDDelete(context.Background(), envID, userID, groupID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsUsersGroupMembershipApi.V1EnvironmentsEnvIDUsersUserIDMemberOfGroupsGroupIDDelete``: %v\n", err)
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
**groupID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUsersUserIDMemberOfGroupsGroupIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDUsersUserIDMemberOfGroupsGroupIDGet

> V1EnvironmentsEnvIDUsersUserIDMemberOfGroupsGroupIDGet(ctx, envID, userID, groupID).Expand(expand).Execute()

READ One Group Membership for User



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
    groupID := "groupID_example" // string | 
    expand := "group" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsUsersGroupMembershipApi.V1EnvironmentsEnvIDUsersUserIDMemberOfGroupsGroupIDGet(context.Background(), envID, userID, groupID).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsUsersGroupMembershipApi.V1EnvironmentsEnvIDUsersUserIDMemberOfGroupsGroupIDGet``: %v\n", err)
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
**groupID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUsersUserIDMemberOfGroupsGroupIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **expand** | **string** |  | 

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


## V1EnvironmentsEnvIDUsersUserIDMemberOfGroupsPost

> V1EnvironmentsEnvIDUsersUserIDMemberOfGroupsPost(ctx, envID, userID).Body(body).Execute()

ADD User to Group



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
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsUsersGroupMembershipApi.V1EnvironmentsEnvIDUsersUserIDMemberOfGroupsPost(context.Background(), envID, userID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsUsersGroupMembershipApi.V1EnvironmentsEnvIDUsersUserIDMemberOfGroupsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUsersUserIDMemberOfGroupsPostRequest struct via the builder pattern


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

