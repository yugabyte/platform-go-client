# \BackupAndRestoreAPI

All URIs are relative to *http://localhost:9000/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListYbcGflagsMetadata**](BackupAndRestoreAPI.md#ListYbcGflagsMetadata) | **Get** /ybc/gflags-metadata | List YBC Gflags metadata



## ListYbcGflagsMetadata

> []GflagMetadata ListYbcGflagsMetadata(ctx).Execute()

List YBC Gflags metadata



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupAndRestoreAPI.ListYbcGflagsMetadata(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAndRestoreAPI.ListYbcGflagsMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListYbcGflagsMetadata`: []GflagMetadata
	fmt.Fprintf(os.Stdout, "Response from `BackupAndRestoreAPI.ListYbcGflagsMetadata`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListYbcGflagsMetadataRequest struct via the builder pattern


### Return type

[**[]GflagMetadata**](GflagMetadata.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

