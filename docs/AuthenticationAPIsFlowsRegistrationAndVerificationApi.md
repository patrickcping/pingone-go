# \AuthenticationAPIsFlowsRegistrationAndVerificationApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnvIDFlowsFlowIDPost**](AuthenticationAPIsFlowsRegistrationAndVerificationApi.md#EnvIDFlowsFlowIDPost) | **Post** /{envID}/flows/{flowID} | Verify User



## EnvIDFlowsFlowIDPost

> EnvIDFlowsFlowIDPost(ctx, envID, flowID).ContentType(contentType).Cookie(cookie).Body(body).Execute()

Verify User



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
    flowID := "flowID_example" // string | 
    contentType := "application/vnd.pingidentity.user.verify+json" // string |  (optional)
    cookie := "ST={{sessionToken}}" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationAPIsFlowsRegistrationAndVerificationApi.EnvIDFlowsFlowIDPost(context.Background(), envID, flowID).ContentType(contentType).Cookie(cookie).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPIsFlowsRegistrationAndVerificationApi.EnvIDFlowsFlowIDPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**flowID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvIDFlowsFlowIDPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contentType** | **string** |  | 
 **cookie** | **string** |  | 
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

