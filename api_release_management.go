/*
 * Yugabyte Platform APIs
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

// ReleaseManagementApiService ReleaseManagementApi service
type ReleaseManagementApiService service

type ReleaseManagementApiApiCreateReleaseRequest struct {
	ctx _context.Context
	ApiService *ReleaseManagementApiService
	cUUID string
	release *ReleaseFormData
}

func (r ReleaseManagementApiApiCreateReleaseRequest) Release(release ReleaseFormData) ReleaseManagementApiApiCreateReleaseRequest {
	r.release = &release
	return r
}

func (r ReleaseManagementApiApiCreateReleaseRequest) Execute() (YBPSuccess, *_nethttp.Response, error) {
	return r.ApiService.CreateReleaseExecute(r)
}

/*
 * CreateRelease Create a release
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID
 * @return ReleaseManagementApiApiCreateReleaseRequest
 */
func (a *ReleaseManagementApiService) CreateRelease(ctx _context.Context, cUUID string) ReleaseManagementApiApiCreateReleaseRequest {
	return ReleaseManagementApiApiCreateReleaseRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
	}
}

/*
 * Execute executes the request
 * @return YBPSuccess
 */
func (a *ReleaseManagementApiService) CreateReleaseExecute(r ReleaseManagementApiApiCreateReleaseRequest) (YBPSuccess, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  YBPSuccess
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReleaseManagementApiService.CreateRelease")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/customers/{cUUID}/releases"
	localVarPath = strings.Replace(localVarPath, "{"+"cUUID"+"}", _neturl.PathEscape(parameterToString(r.cUUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.release == nil {
		return localVarReturnValue, nil, reportError("release is required and must be specified")
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
	localVarPostBody = r.release
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

type ReleaseManagementApiApiDeleteReleaseRequest struct {
	ctx _context.Context
	ApiService *ReleaseManagementApiService
	cUUID string
	name string
}


func (r ReleaseManagementApiApiDeleteReleaseRequest) Execute() (ReleaseMetadata, *_nethttp.Response, error) {
	return r.ApiService.DeleteReleaseExecute(r)
}

/*
 * DeleteRelease Delete a release
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID
 * @param name
 * @return ReleaseManagementApiApiDeleteReleaseRequest
 */
func (a *ReleaseManagementApiService) DeleteRelease(ctx _context.Context, cUUID string, name string) ReleaseManagementApiApiDeleteReleaseRequest {
	return ReleaseManagementApiApiDeleteReleaseRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
		name: name,
	}
}

/*
 * Execute executes the request
 * @return ReleaseMetadata
 */
func (a *ReleaseManagementApiService) DeleteReleaseExecute(r ReleaseManagementApiApiDeleteReleaseRequest) (ReleaseMetadata, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ReleaseMetadata
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReleaseManagementApiService.DeleteRelease")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/customers/{cUUID}/releases/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"cUUID"+"}", _neturl.PathEscape(parameterToString(r.cUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", _neturl.PathEscape(parameterToString(r.name, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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

type ReleaseManagementApiApiGetListOfReleasesRequest struct {
	ctx _context.Context
	ApiService *ReleaseManagementApiService
	cUUID string
	includeMetadata *bool
}

func (r ReleaseManagementApiApiGetListOfReleasesRequest) IncludeMetadata(includeMetadata bool) ReleaseManagementApiApiGetListOfReleasesRequest {
	r.includeMetadata = &includeMetadata
	return r
}

func (r ReleaseManagementApiApiGetListOfReleasesRequest) Execute() (map[string]map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetListOfReleasesExecute(r)
}

/*
 * GetListOfReleases List all releases
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID
 * @return ReleaseManagementApiApiGetListOfReleasesRequest
 */
func (a *ReleaseManagementApiService) GetListOfReleases(ctx _context.Context, cUUID string) ReleaseManagementApiApiGetListOfReleasesRequest {
	return ReleaseManagementApiApiGetListOfReleasesRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
	}
}

/*
 * Execute executes the request
 * @return map[string]map[string]interface{}
 */
func (a *ReleaseManagementApiService) GetListOfReleasesExecute(r ReleaseManagementApiApiGetListOfReleasesRequest) (map[string]map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReleaseManagementApiService.GetListOfReleases")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/customers/{cUUID}/releases"
	localVarPath = strings.Replace(localVarPath, "{"+"cUUID"+"}", _neturl.PathEscape(parameterToString(r.cUUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.includeMetadata != nil {
		localVarQueryParams.Add("includeMetadata", parameterToString(*r.includeMetadata, ""))
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

type ReleaseManagementApiApiRefreshRequest struct {
	ctx _context.Context
	ApiService *ReleaseManagementApiService
	cUUID string
}


func (r ReleaseManagementApiApiRefreshRequest) Execute() (YBPSuccess, *_nethttp.Response, error) {
	return r.ApiService.RefreshExecute(r)
}

/*
 * Refresh Refresh a release
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID
 * @return ReleaseManagementApiApiRefreshRequest
 */
func (a *ReleaseManagementApiService) Refresh(ctx _context.Context, cUUID string) ReleaseManagementApiApiRefreshRequest {
	return ReleaseManagementApiApiRefreshRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
	}
}

/*
 * Execute executes the request
 * @return YBPSuccess
 */
func (a *ReleaseManagementApiService) RefreshExecute(r ReleaseManagementApiApiRefreshRequest) (YBPSuccess, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  YBPSuccess
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReleaseManagementApiService.Refresh")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/customers/{cUUID}/releases"
	localVarPath = strings.Replace(localVarPath, "{"+"cUUID"+"}", _neturl.PathEscape(parameterToString(r.cUUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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

type ReleaseManagementApiApiUpdateReleaseRequest struct {
	ctx _context.Context
	ApiService *ReleaseManagementApiService
	cUUID string
	name string
	release *map[string]interface{}
}

func (r ReleaseManagementApiApiUpdateReleaseRequest) Release(release map[string]interface{}) ReleaseManagementApiApiUpdateReleaseRequest {
	r.release = &release
	return r
}

func (r ReleaseManagementApiApiUpdateReleaseRequest) Execute() (ReleaseMetadata, *_nethttp.Response, error) {
	return r.ApiService.UpdateReleaseExecute(r)
}

/*
 * UpdateRelease Update a release
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID
 * @param name
 * @return ReleaseManagementApiApiUpdateReleaseRequest
 */
func (a *ReleaseManagementApiService) UpdateRelease(ctx _context.Context, cUUID string, name string) ReleaseManagementApiApiUpdateReleaseRequest {
	return ReleaseManagementApiApiUpdateReleaseRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
		name: name,
	}
}

/*
 * Execute executes the request
 * @return ReleaseMetadata
 */
func (a *ReleaseManagementApiService) UpdateReleaseExecute(r ReleaseManagementApiApiUpdateReleaseRequest) (ReleaseMetadata, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ReleaseMetadata
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReleaseManagementApiService.UpdateRelease")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/customers/{cUUID}/releases/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"cUUID"+"}", _neturl.PathEscape(parameterToString(r.cUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", _neturl.PathEscape(parameterToString(r.name, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.release == nil {
		return localVarReturnValue, nil, reportError("release is required and must be specified")
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
	localVarPostBody = r.release
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
