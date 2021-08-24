# \ManagementAPIsSubscriptionsWebhooksApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDSubscriptionsGet**](ManagementAPIsSubscriptionsWebhooksApi.md#V1EnvironmentsEnvIDSubscriptionsGet) | **Get** /v1/environments/{envID}/subscriptions | READ All Subscriptions
[**V1EnvironmentsEnvIDSubscriptionsPost**](ManagementAPIsSubscriptionsWebhooksApi.md#V1EnvironmentsEnvIDSubscriptionsPost) | **Post** /v1/environments/{envID}/subscriptions | CREATE Subscriptions
[**V1EnvironmentsEnvIDSubscriptionsSubscriptionIDDelete**](ManagementAPIsSubscriptionsWebhooksApi.md#V1EnvironmentsEnvIDSubscriptionsSubscriptionIDDelete) | **Delete** /v1/environments/{envID}/subscriptions/{subscriptionID} | DELETE Subscription
[**V1EnvironmentsEnvIDSubscriptionsSubscriptionIDGet**](ManagementAPIsSubscriptionsWebhooksApi.md#V1EnvironmentsEnvIDSubscriptionsSubscriptionIDGet) | **Get** /v1/environments/{envID}/subscriptions/{subscriptionID} | READ One Subscription
[**V1EnvironmentsEnvIDSubscriptionsSubscriptionIDPut**](ManagementAPIsSubscriptionsWebhooksApi.md#V1EnvironmentsEnvIDSubscriptionsSubscriptionIDPut) | **Put** /v1/environments/{envID}/subscriptions/{subscriptionID} | UPDATE Subscription



## V1EnvironmentsEnvIDSubscriptionsGet

> V1EnvironmentsEnvIDSubscriptionsGet(ctx, envID).Authorization(authorization).Execute()

READ All Subscriptions



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
    resp, r, err := api_client.ManagementAPIsSubscriptionsWebhooksApi.V1EnvironmentsEnvIDSubscriptionsGet(context.Background(), envID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsSubscriptionsWebhooksApi.V1EnvironmentsEnvIDSubscriptionsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDSubscriptionsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDSubscriptionsPost

> V1EnvironmentsEnvIDSubscriptionsPost(ctx, envID).ContentType(contentType).Authorization(authorization).Body(body).Execute()

CREATE Subscriptions



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
    contentType := "application/json" // string |  (optional)
    authorization := "Bearer {{jwtToken}}" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsSubscriptionsWebhooksApi.V1EnvironmentsEnvIDSubscriptionsPost(context.Background(), envID).ContentType(contentType).Authorization(authorization).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsSubscriptionsWebhooksApi.V1EnvironmentsEnvIDSubscriptionsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDSubscriptionsPostRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDSubscriptionsSubscriptionIDDelete

> V1EnvironmentsEnvIDSubscriptionsSubscriptionIDDelete(ctx, envID, subscriptionID).Authorization(authorization).Execute()

DELETE Subscription



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
    subscriptionID := "subscriptionID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsSubscriptionsWebhooksApi.V1EnvironmentsEnvIDSubscriptionsSubscriptionIDDelete(context.Background(), envID, subscriptionID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsSubscriptionsWebhooksApi.V1EnvironmentsEnvIDSubscriptionsSubscriptionIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**subscriptionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDSubscriptionsSubscriptionIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDSubscriptionsSubscriptionIDGet

> V1EnvironmentsEnvIDSubscriptionsSubscriptionIDGet(ctx, envID, subscriptionID).Authorization(authorization).Execute()

READ One Subscription



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
    subscriptionID := "subscriptionID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsSubscriptionsWebhooksApi.V1EnvironmentsEnvIDSubscriptionsSubscriptionIDGet(context.Background(), envID, subscriptionID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsSubscriptionsWebhooksApi.V1EnvironmentsEnvIDSubscriptionsSubscriptionIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**subscriptionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDSubscriptionsSubscriptionIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDSubscriptionsSubscriptionIDPut

> V1EnvironmentsEnvIDSubscriptionsSubscriptionIDPut(ctx, envID, subscriptionID).ContentType(contentType).Authorization(authorization).Body(body).Execute()

UPDATE Subscription



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
    subscriptionID := "subscriptionID_example" // string | 
    contentType := "application/json" // string |  (optional)
    authorization := "Bearer {{jwtToken}}" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsSubscriptionsWebhooksApi.V1EnvironmentsEnvIDSubscriptionsSubscriptionIDPut(context.Background(), envID, subscriptionID).ContentType(contentType).Authorization(authorization).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsSubscriptionsWebhooksApi.V1EnvironmentsEnvIDSubscriptionsSubscriptionIDPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**subscriptionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDSubscriptionsSubscriptionIDPutRequest struct via the builder pattern


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

