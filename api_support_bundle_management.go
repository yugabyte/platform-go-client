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

// SupportBundleManagementApiService SupportBundleManagementApi service
type SupportBundleManagementApiService service

type SupportBundleManagementApiApiCreateSupportBundleRequest struct {
	ctx _context.Context
	ApiService *SupportBundleManagementApiService
	cUUID string
	uniUUID string
	supportBundle *SupportBundleFormData
	request *interface{}
}

func (r SupportBundleManagementApiApiCreateSupportBundleRequest) SupportBundle(supportBundle SupportBundleFormData) SupportBundleManagementApiApiCreateSupportBundleRequest {
	r.supportBundle = &supportBundle
	return r
}
func (r SupportBundleManagementApiApiCreateSupportBundleRequest) Request(request interface{}) SupportBundleManagementApiApiCreateSupportBundleRequest {
	r.request = &request
	return r
}

func (r SupportBundleManagementApiApiCreateSupportBundleRequest) Execute() (YBPTask, *_nethttp.Response, error) {
	return r.ApiService.CreateSupportBundleExecute(r)
}

/*
 * CreateSupportBundle Create support bundle for specific universe
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID
 * @param uniUUID
 * @return SupportBundleManagementApiApiCreateSupportBundleRequest
 */
func (a *SupportBundleManagementApiService) CreateSupportBundle(ctx _context.Context, cUUID string, uniUUID string) SupportBundleManagementApiApiCreateSupportBundleRequest {
	return SupportBundleManagementApiApiCreateSupportBundleRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
		uniUUID: uniUUID,
	}
}

/*
 * Execute executes the request
 * @return YBPTask
 */
func (a *SupportBundleManagementApiService) CreateSupportBundleExecute(r SupportBundleManagementApiApiCreateSupportBundleRequest) (YBPTask, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  YBPTask
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SupportBundleManagementApiService.CreateSupportBundle")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/customers/{cUUID}/universes/{uniUUID}/support_bundle"
	localVarPath = strings.Replace(localVarPath, "{"+"cUUID"+"}", _neturl.PathEscape(parameterToString(r.cUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"uniUUID"+"}", _neturl.PathEscape(parameterToString(r.uniUUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.supportBundle == nil {
		return localVarReturnValue, nil, reportError("supportBundle is required and must be specified")
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
	localVarPostBody = r.supportBundle
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

type SupportBundleManagementApiApiDeleteSupportBundleRequest struct {
	ctx _context.Context
	ApiService *SupportBundleManagementApiService
	cUUID string
	uniUUID string
	sbUUID string
	request *interface{}
}

func (r SupportBundleManagementApiApiDeleteSupportBundleRequest) Request(request interface{}) SupportBundleManagementApiApiDeleteSupportBundleRequest {
	r.request = &request
	return r
}

func (r SupportBundleManagementApiApiDeleteSupportBundleRequest) Execute() (YBPSuccess, *_nethttp.Response, error) {
	return r.ApiService.DeleteSupportBundleExecute(r)
}

/*
 * DeleteSupportBundle Delete a support bundle
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID
 * @param uniUUID
 * @param sbUUID
 * @return SupportBundleManagementApiApiDeleteSupportBundleRequest
 */
func (a *SupportBundleManagementApiService) DeleteSupportBundle(ctx _context.Context, cUUID string, uniUUID string, sbUUID string) SupportBundleManagementApiApiDeleteSupportBundleRequest {
	return SupportBundleManagementApiApiDeleteSupportBundleRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
		uniUUID: uniUUID,
		sbUUID: sbUUID,
	}
}

/*
 * Execute executes the request
 * @return YBPSuccess
 */
func (a *SupportBundleManagementApiService) DeleteSupportBundleExecute(r SupportBundleManagementApiApiDeleteSupportBundleRequest) (YBPSuccess, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  YBPSuccess
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SupportBundleManagementApiService.DeleteSupportBundle")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/customers/{cUUID}/universes/{uniUUID}/support_bundle/{sbUUID}"
	localVarPath = strings.Replace(localVarPath, "{"+"cUUID"+"}", _neturl.PathEscape(parameterToString(r.cUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"uniUUID"+"}", _neturl.PathEscape(parameterToString(r.uniUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sbUUID"+"}", _neturl.PathEscape(parameterToString(r.sbUUID, "")), -1)

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

type SupportBundleManagementApiApiDownloadSupportBundleRequest struct {
	ctx _context.Context
	ApiService *SupportBundleManagementApiService
	cUUID string
	uniUUID string
	sbUUID string
}


func (r SupportBundleManagementApiApiDownloadSupportBundleRequest) Execute() (string, *_nethttp.Response, error) {
	return r.ApiService.DownloadSupportBundleExecute(r)
}

/*
 * DownloadSupportBundle Download support bundle
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID
 * @param uniUUID
 * @param sbUUID
 * @return SupportBundleManagementApiApiDownloadSupportBundleRequest
 */
func (a *SupportBundleManagementApiService) DownloadSupportBundle(ctx _context.Context, cUUID string, uniUUID string, sbUUID string) SupportBundleManagementApiApiDownloadSupportBundleRequest {
	return SupportBundleManagementApiApiDownloadSupportBundleRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
		uniUUID: uniUUID,
		sbUUID: sbUUID,
	}
}

/*
 * Execute executes the request
 * @return string
 */
func (a *SupportBundleManagementApiService) DownloadSupportBundleExecute(r SupportBundleManagementApiApiDownloadSupportBundleRequest) (string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SupportBundleManagementApiService.DownloadSupportBundle")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/customers/{cUUID}/universes/{uniUUID}/support_bundle/{sbUUID}/download"
	localVarPath = strings.Replace(localVarPath, "{"+"cUUID"+"}", _neturl.PathEscape(parameterToString(r.cUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"uniUUID"+"}", _neturl.PathEscape(parameterToString(r.uniUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sbUUID"+"}", _neturl.PathEscape(parameterToString(r.sbUUID, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/x-compressed"}

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

type SupportBundleManagementApiApiGetSupportBundleRequest struct {
	ctx _context.Context
	ApiService *SupportBundleManagementApiService
	cUUID string
	uniUUID string
	sbUUID string
}


func (r SupportBundleManagementApiApiGetSupportBundleRequest) Execute() (SupportBundle, *_nethttp.Response, error) {
	return r.ApiService.GetSupportBundleExecute(r)
}

/*
 * GetSupportBundle Get a support bundle from a universe
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID
 * @param uniUUID
 * @param sbUUID
 * @return SupportBundleManagementApiApiGetSupportBundleRequest
 */
func (a *SupportBundleManagementApiService) GetSupportBundle(ctx _context.Context, cUUID string, uniUUID string, sbUUID string) SupportBundleManagementApiApiGetSupportBundleRequest {
	return SupportBundleManagementApiApiGetSupportBundleRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
		uniUUID: uniUUID,
		sbUUID: sbUUID,
	}
}

/*
 * Execute executes the request
 * @return SupportBundle
 */
func (a *SupportBundleManagementApiService) GetSupportBundleExecute(r SupportBundleManagementApiApiGetSupportBundleRequest) (SupportBundle, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SupportBundle
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SupportBundleManagementApiService.GetSupportBundle")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/customers/{cUUID}/universes/{uniUUID}/support_bundle/{sbUUID}"
	localVarPath = strings.Replace(localVarPath, "{"+"cUUID"+"}", _neturl.PathEscape(parameterToString(r.cUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"uniUUID"+"}", _neturl.PathEscape(parameterToString(r.uniUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sbUUID"+"}", _neturl.PathEscape(parameterToString(r.sbUUID, "")), -1)

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

type SupportBundleManagementApiApiListSupportBundleRequest struct {
	ctx _context.Context
	ApiService *SupportBundleManagementApiService
	cUUID string
	uniUUID string
}


func (r SupportBundleManagementApiApiListSupportBundleRequest) Execute() ([]SupportBundle, *_nethttp.Response, error) {
	return r.ApiService.ListSupportBundleExecute(r)
}

/*
 * ListSupportBundle List all support bundles from a universe
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID
 * @param uniUUID
 * @return SupportBundleManagementApiApiListSupportBundleRequest
 */
func (a *SupportBundleManagementApiService) ListSupportBundle(ctx _context.Context, cUUID string, uniUUID string) SupportBundleManagementApiApiListSupportBundleRequest {
	return SupportBundleManagementApiApiListSupportBundleRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
		uniUUID: uniUUID,
	}
}

/*
 * Execute executes the request
 * @return []SupportBundle
 */
func (a *SupportBundleManagementApiService) ListSupportBundleExecute(r SupportBundleManagementApiApiListSupportBundleRequest) ([]SupportBundle, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []SupportBundle
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SupportBundleManagementApiService.ListSupportBundle")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/customers/{cUUID}/universes/{uniUUID}/support_bundle"
	localVarPath = strings.Replace(localVarPath, "{"+"cUUID"+"}", _neturl.PathEscape(parameterToString(r.cUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"uniUUID"+"}", _neturl.PathEscape(parameterToString(r.uniUUID, "")), -1)

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

type SupportBundleManagementApiApiListSupportBundleComponentsRequest struct {
	ctx _context.Context
	ApiService *SupportBundleManagementApiService
	cUUID string
}


func (r SupportBundleManagementApiApiListSupportBundleComponentsRequest) Execute() ([]string, *_nethttp.Response, error) {
	return r.ApiService.ListSupportBundleComponentsExecute(r)
}

/*
 * ListSupportBundleComponents List all components available in support bundle
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID
 * @return SupportBundleManagementApiApiListSupportBundleComponentsRequest
 */
func (a *SupportBundleManagementApiService) ListSupportBundleComponents(ctx _context.Context, cUUID string) SupportBundleManagementApiApiListSupportBundleComponentsRequest {
	return SupportBundleManagementApiApiListSupportBundleComponentsRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
	}
}

/*
 * Execute executes the request
 * @return []string
 */
func (a *SupportBundleManagementApiService) ListSupportBundleComponentsExecute(r SupportBundleManagementApiApiListSupportBundleComponentsRequest) ([]string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SupportBundleManagementApiService.ListSupportBundleComponents")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/customers/{cUUID}/support_bundle/components"
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
