/*
 * YugabyteDB Anywhere APIs
 *
 * ALPHA - NOT FOR EXTERNAL USE
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ywclient

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// ExtractMetadataFromRemoteTarballApiService ExtractMetadataFromRemoteTarballApi service
type ExtractMetadataFromRemoteTarballApiService service

type ExtractMetadataFromRemoteTarballApiApiExtractMetadataRequest struct {
	ctx _context.Context
	ApiService *ExtractMetadataFromRemoteTarballApiService
	cUUID string
	releaseURL *ExtractMetadata
	request *interface{}
}

func (r ExtractMetadataFromRemoteTarballApiApiExtractMetadataRequest) ReleaseURL(releaseURL ExtractMetadata) ExtractMetadataFromRemoteTarballApiApiExtractMetadataRequest {
	r.releaseURL = &releaseURL
	return r
}
func (r ExtractMetadataFromRemoteTarballApiApiExtractMetadataRequest) Request(request interface{}) ExtractMetadataFromRemoteTarballApiApiExtractMetadataRequest {
	r.request = &request
	return r
}

func (r ExtractMetadataFromRemoteTarballApiApiExtractMetadataRequest) Execute() (YBPCreateSuccess, *_nethttp.Response, error) {
	return r.ApiService.ExtractMetadataExecute(r)
}

/*
 * ExtractMetadata helper to extract release metadata from a remote tarball
 * WARNING: This is a preview API that could change: start extracting metadata from a remote tgz url
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID
 * @return ExtractMetadataFromRemoteTarballApiApiExtractMetadataRequest
 */
func (a *ExtractMetadataFromRemoteTarballApiService) ExtractMetadata(ctx _context.Context, cUUID string) ExtractMetadataFromRemoteTarballApiApiExtractMetadataRequest {
	return ExtractMetadataFromRemoteTarballApiApiExtractMetadataRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
	}
}

/*
 * Execute executes the request
 * @return YBPCreateSuccess
 */
func (a *ExtractMetadataFromRemoteTarballApiService) ExtractMetadataExecute(r ExtractMetadataFromRemoteTarballApiApiExtractMetadataRequest) (YBPCreateSuccess, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  YBPCreateSuccess
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExtractMetadataFromRemoteTarballApiService.ExtractMetadata")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/customers/{cUUID}/ybdb_release/extract_metadata"
	localVarPath = strings.Replace(localVarPath, "{"+"cUUID"+"}", _neturl.PathEscape(parameterToString(r.cUUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.releaseURL == nil {
		return localVarReturnValue, nil, reportError("releaseURL is required and must be specified")
	}

	if r.request != nil {
		localVarQueryParams.Add("request", parameterToString(*r.request, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.releaseURL
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-AUTH-YW-API-TOKEN"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ExtractMetadataFromRemoteTarballApiApiExtractMetadata_0Request struct {
	ctx _context.Context
	ApiService *ExtractMetadataFromRemoteTarballApiService
	cUUID string
	rUUID string
	request *interface{}
}

func (r ExtractMetadataFromRemoteTarballApiApiExtractMetadata_0Request) Request(request interface{}) ExtractMetadataFromRemoteTarballApiApiExtractMetadata_0Request {
	r.request = &request
	return r
}

func (r ExtractMetadataFromRemoteTarballApiApiExtractMetadata_0Request) Execute() (ResponseExtractMetadata, *_nethttp.Response, error) {
	return r.ApiService.ExtractMetadata_1Execute(r)
}

/*
 * ExtractMetadata_0 get the extract release metadata from a remote tarball
 * WARNING: This is a preview API that could change: Get extract metadata and its progress.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID
 * @param rUUID
 * @return ExtractMetadataFromRemoteTarballApiApiExtractMetadata_0Request
 */
func (a *ExtractMetadataFromRemoteTarballApiService) ExtractMetadata_1(ctx _context.Context, cUUID string, rUUID string) ExtractMetadataFromRemoteTarballApiApiExtractMetadata_0Request {
	return ExtractMetadataFromRemoteTarballApiApiExtractMetadata_0Request{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
		rUUID: rUUID,
	}
}

/*
 * Execute executes the request
 * @return ResponseExtractMetadata
 */
func (a *ExtractMetadataFromRemoteTarballApiService) ExtractMetadata_1Execute(r ExtractMetadataFromRemoteTarballApiApiExtractMetadata_0Request) (ResponseExtractMetadata, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ResponseExtractMetadata
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExtractMetadataFromRemoteTarballApiService.ExtractMetadata_1")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/customers/{cUUID}/ybdb_release/extract_metadata/{rUUID}"
	localVarPath = strings.Replace(localVarPath, "{"+"cUUID"+"}", _neturl.PathEscape(parameterToString(r.cUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rUUID"+"}", _neturl.PathEscape(parameterToString(r.rUUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.request != nil {
		localVarQueryParams.Add("request", parameterToString(*r.request, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-AUTH-YW-API-TOKEN"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
