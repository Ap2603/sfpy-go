# \DefaultAPI

All URIs are relative to *https://dev.api.getsafepay.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAPIKey**](DefaultAPI.md#CreateAPIKey) | **Post** /client/api-settings/v1/ | Create API Key
[**CreateGuest**](DefaultAPI.md#CreateGuest) | **Post** /user/v1/guest/ | Create Guest JWT
[**CreateJWT**](DefaultAPI.md#CreateJWT) | **Post** /auth/v1/company/authenticate | Create JWT Token
[**GenerateTimeBasedToken**](DefaultAPI.md#GenerateTimeBasedToken) | **Post** /client/passport/v1/token | Generate Time Based Token



## CreateAPIKey

> CreateAPIKeyResponse CreateAPIKey(ctx).Execute()

Create API Key

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateAPIKey(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateAPIKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAPIKey`: CreateAPIKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateAPIKey`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAPIKeyRequest struct via the builder pattern


### Return type

[**CreateAPIKeyResponse**](CreateAPIKeyResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGuest

> CreateGuestResponse CreateGuest(ctx).CreateGuestRequest(createGuestRequest).Execute()

Create Guest JWT

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	createGuestRequest := *openapiclient.NewCreateGuestRequest("Email_example", "Phone_example", "Country_example") // CreateGuestRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateGuest(context.Background()).CreateGuestRequest(createGuestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateGuest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGuest`: CreateGuestResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateGuest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGuestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createGuestRequest** | [**CreateGuestRequest**](CreateGuestRequest.md) |  | 

### Return type

[**CreateGuestResponse**](CreateGuestResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [merchantSecret](../README.md#merchantSecret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateJWT

> CreateJWTResponse CreateJWT(ctx).CreateJWTRequest(createJWTRequest).Execute()

Create JWT Token

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	createJWTRequest := *openapiclient.NewCreateJWTRequest("Client_example", "Email_example", "Password_example") // CreateJWTRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateJWT(context.Background()).CreateJWTRequest(createJWTRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateJWT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateJWT`: CreateJWTResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateJWT`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateJWTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createJWTRequest** | [**CreateJWTRequest**](CreateJWTRequest.md) |  | 

### Return type

[**CreateJWTResponse**](CreateJWTResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [merchantSecret](../README.md#merchantSecret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateTimeBasedToken

> GenerateTimeBasedTokenResponse GenerateTimeBasedToken(ctx).Execute()

Generate Time Based Token

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GenerateTimeBasedToken(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GenerateTimeBasedToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateTimeBasedToken`: GenerateTimeBasedTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GenerateTimeBasedToken`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateTimeBasedTokenRequest struct via the builder pattern


### Return type

[**GenerateTimeBasedTokenResponse**](GenerateTimeBasedTokenResponse.md)

### Authorization

[merchantSecret](../README.md#merchantSecret)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

