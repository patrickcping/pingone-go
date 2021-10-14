# \ManagementAPIsUsersUsersApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUser**](ManagementAPIsUsersUsersApi.md#CreateUser) | **Post** /v1/environments/{envID}/users | CREATE User
[**DeleteUser**](ManagementAPIsUsersUsersApi.md#DeleteUser) | **Delete** /v1/environments/{envID}/users/{userID} | DELETE User
[**ReadAllUsers**](ManagementAPIsUsersUsersApi.md#ReadAllUsers) | **Get** /v1/environments/{envID}/users | READ All Users
[**ReadUser**](ManagementAPIsUsersUsersApi.md#ReadUser) | **Get** /v1/environments/{envID}/users/{userID} | READ User
[**UpdateUserPatch**](ManagementAPIsUsersUsersApi.md#UpdateUserPatch) | **Patch** /v1/environments/{envID}/users/{userID} | UPDATE User (Patch)
[**UpdateUserPut**](ManagementAPIsUsersUsersApi.md#UpdateUserPut) | **Put** /v1/environments/{envID}/users/{userID} | UPDATE User (Put)
[**V1EnvironmentsEnvIDUsersUserIDIdentityProviderGet**](ManagementAPIsUsersUsersApi.md#V1EnvironmentsEnvIDUsersUserIDIdentityProviderGet) | **Get** /v1/environments/{envID}/users/{userID}/identityProvider | READ User Identity Provider
[**V1EnvironmentsEnvIDUsersUserIDIdentityProviderPut**](ManagementAPIsUsersUsersApi.md#V1EnvironmentsEnvIDUsersUserIDIdentityProviderPut) | **Put** /v1/environments/{envID}/users/{userID}/identityProvider | UPDATE User Identity Provider
[**V1EnvironmentsEnvIDUsersUserIDVerifyStatusGet**](ManagementAPIsUsersUsersApi.md#V1EnvironmentsEnvIDUsersUserIDVerifyStatusGet) | **Get** /v1/environments/{envID}/users/{userID}/verifyStatus | READ user verification status
[**V1EnvironmentsEnvIDUsersUserIDVerifyStatusPut**](ManagementAPIsUsersUsersApi.md#V1EnvironmentsEnvIDUsersUserIDVerifyStatusPut) | **Put** /v1/environments/{envID}/users/{userID}/verifyStatus | UPDATE user verification status



## CreateUser

> User CreateUser(ctx, envID).ContentType(contentType).User(user).Execute()

CREATE User



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
    contentType := "application/vnd.pingidentity.user.import+json" // string |  (optional)
    user := *openapiclient.NewUser("Email_example", *openapiclient.NewUserPopulation("Id_example"), "Username_example") // User |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsUsersUsersApi.CreateUser(context.Background(), envID).ContentType(contentType).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsUsersUsersApi.CreateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUser`: User
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsUsersUsersApi.CreateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | 
 **user** | [**User**](User.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> DeleteUser(ctx, envID, userID).Execute()

DELETE User



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
    resp, r, err := api_client.ManagementAPIsUsersUsersApi.DeleteUser(context.Background(), envID, userID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsUsersUsersApi.DeleteUser``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


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


## ReadAllUsers

> EntityArray ReadAllUsers(ctx, envID).Filter(filter).Limit(limit).Execute()

READ All Users



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
    filter := "memberOfGroups[id eq "{{groupID}}"] and name.family eq "demo"" // string |  (optional)
    limit := int32(100) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsUsersUsersApi.ReadAllUsers(context.Background(), envID).Filter(filter).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsUsersUsersApi.ReadAllUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllUsers`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsUsersUsersApi.ReadAllUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** |  | 
 **limit** | **int32** |  | 

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


## ReadUser

> User ReadUser(ctx, envID, userID).Include(include).Execute()

READ User



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
    resp, r, err := api_client.ManagementAPIsUsersUsersApi.ReadUser(context.Background(), envID, userID).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsUsersUsersApi.ReadUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadUser`: User
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsUsersUsersApi.ReadUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**userID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **include** | **string** |  | 

### Return type

[**User**](User.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserPatch

> User UpdateUserPatch(ctx, envID, userID).User(user).Execute()

UPDATE User (Patch)



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
    user := *openapiclient.NewUser("Email_example", *openapiclient.NewUserPopulation("Id_example"), "Username_example") // User |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsUsersUsersApi.UpdateUserPatch(context.Background(), envID, userID).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsUsersUsersApi.UpdateUserPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserPatch`: User
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsUsersUsersApi.UpdateUserPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**userID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **user** | [**User**](User.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserPut

> User UpdateUserPut(ctx, envID, userID).User(user).Execute()

UPDATE User (Put)



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
    user := *openapiclient.NewUser("Email_example", *openapiclient.NewUserPopulation("Id_example"), "Username_example") // User |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsUsersUsersApi.UpdateUserPut(context.Background(), envID, userID).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsUsersUsersApi.UpdateUserPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserPut`: User
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsUsersUsersApi.UpdateUserPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**userID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **user** | [**User**](User.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnvironmentsEnvIDUsersUserIDIdentityProviderGet

> V1EnvironmentsEnvIDUsersUserIDIdentityProviderGet(ctx, envID, userID).Execute()

READ User Identity Provider



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
    resp, r, err := api_client.ManagementAPIsUsersUsersApi.V1EnvironmentsEnvIDUsersUserIDIdentityProviderGet(context.Background(), envID, userID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsUsersUsersApi.V1EnvironmentsEnvIDUsersUserIDIdentityProviderGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUsersUserIDIdentityProviderGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDUsersUserIDIdentityProviderPut

> V1EnvironmentsEnvIDUsersUserIDIdentityProviderPut(ctx, envID, userID).Body(body).Execute()

UPDATE User Identity Provider



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
    resp, r, err := api_client.ManagementAPIsUsersUsersApi.V1EnvironmentsEnvIDUsersUserIDIdentityProviderPut(context.Background(), envID, userID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsUsersUsersApi.V1EnvironmentsEnvIDUsersUserIDIdentityProviderPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUsersUserIDIdentityProviderPutRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDUsersUserIDVerifyStatusGet

> V1EnvironmentsEnvIDUsersUserIDVerifyStatusGet(ctx, envID, userID).Execute()

READ user verification status



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
    resp, r, err := api_client.ManagementAPIsUsersUsersApi.V1EnvironmentsEnvIDUsersUserIDVerifyStatusGet(context.Background(), envID, userID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsUsersUsersApi.V1EnvironmentsEnvIDUsersUserIDVerifyStatusGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUsersUserIDVerifyStatusGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDUsersUserIDVerifyStatusPut

> V1EnvironmentsEnvIDUsersUserIDVerifyStatusPut(ctx, envID, userID).Body(body).Execute()

UPDATE user verification status



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
    resp, r, err := api_client.ManagementAPIsUsersUsersApi.V1EnvironmentsEnvIDUsersUserIDVerifyStatusPut(context.Background(), envID, userID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsUsersUsersApi.V1EnvironmentsEnvIDUsersUserIDVerifyStatusPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUsersUserIDVerifyStatusPutRequest struct via the builder pattern


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

