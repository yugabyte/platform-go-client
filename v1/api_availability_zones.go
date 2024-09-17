/*
 * YugabyteDB Anywhere APIs
 *
 * ALPHA - NOT FOR EXTERNAL USE
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

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

// AvailabilityZonesApiService AvailabilityZonesApi service
type AvailabilityZonesApiService service

type AvailabilityZonesApiApiCreateAZRequest struct {
	ctx _context.Context
	ApiService *AvailabilityZonesApiService
	cUUID string
	pUUID string
	rUUID string
	azFormData *AvailabilityZoneFormData
	request *interface{}
}

func (r AvailabilityZonesApiApiCreateAZRequest) AzFormData(azFormData AvailabilityZoneFormData) AvailabilityZonesApiApiCreateAZRequest {
	r.azFormData = &azFormData
	return r
}
func (r AvailabilityZonesApiApiCreateAZRequest) Request(request interface{}) AvailabilityZonesApiApiCreateAZRequest {
	r.request = &request
	return r
}

func (r AvailabilityZonesApiApiCreateAZRequest) Execute() (map[string]AvailabilityZone, *_nethttp.Response, error) {
	return r.ApiService.CreateAZExecute(r)
}

/*
 * CreateAZ Create an availability zone - deprecated
 * <b style="color:#ff0000">Deprecated since YBA version 2.18.2.0.</b></p>Use /api/v1/customers/{cUUID}/provider/{pUUID}/provider_regions/:rUUID/region_zones instead
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID
 * @param pUUID
 * @param rUUID
 * @return AvailabilityZonesApiApiCreateAZRequest
 */
func (a *AvailabilityZonesApiService) CreateAZ(ctx _context.Context, cUUID string, pUUID string, rUUID string) AvailabilityZonesApiApiCreateAZRequest {
	return AvailabilityZonesApiApiCreateAZRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
		pUUID: pUUID,
		rUUID: rUUID,
	}
}

/*
 * Execute executes the request
 * @return map[string]AvailabilityZone
 */
func (a *AvailabilityZonesApiService) CreateAZExecute(r AvailabilityZonesApiApiCreateAZRequest) (map[string]AvailabilityZone, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]AvailabilityZone
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AvailabilityZonesApiService.CreateAZ")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/customers/{cUUID}/providers/{pUUID}/regions/{rUUID}/zones"
	localVarPath = strings.Replace(localVarPath, "{"+"cUUID"+"}", _neturl.PathEscape(parameterToString(r.cUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pUUID"+"}", _neturl.PathEscape(parameterToString(r.pUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rUUID"+"}", _neturl.PathEscape(parameterToString(r.rUUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.azFormData == nil {
		return localVarReturnValue, nil, reportError("azFormData is required and must be specified")
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
	localVarPostBody = r.azFormData
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

type AvailabilityZonesApiApiCreateZoneRequest struct {
	ctx _context.Context
	ApiService *AvailabilityZonesApiService
	cUUID string
	pUUID string
	rUUID string
	azData *AvailabilityZone
	request *interface{}
}

func (r AvailabilityZonesApiApiCreateZoneRequest) AzData(azData AvailabilityZone) AvailabilityZonesApiApiCreateZoneRequest {
	r.azData = &azData
	return r
}
func (r AvailabilityZonesApiApiCreateZoneRequest) Request(request interface{}) AvailabilityZonesApiApiCreateZoneRequest {
	r.request = &request
	return r
}

func (r AvailabilityZonesApiApiCreateZoneRequest) Execute() (map[string]AvailabilityZone, *_nethttp.Response, error) {
	return r.ApiService.CreateZoneExecute(r)
}

/*
 * CreateZone Create an availability zone
 * WARNING: This is a preview API that could change.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID
 * @param pUUID
 * @param rUUID
 * @return AvailabilityZonesApiApiCreateZoneRequest
 */
func (a *AvailabilityZonesApiService) CreateZone(ctx _context.Context, cUUID string, pUUID string, rUUID string) AvailabilityZonesApiApiCreateZoneRequest {
	return AvailabilityZonesApiApiCreateZoneRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
		pUUID: pUUID,
		rUUID: rUUID,
	}
}

/*
 * Execute executes the request
 * @return map[string]AvailabilityZone
 */
func (a *AvailabilityZonesApiService) CreateZoneExecute(r AvailabilityZonesApiApiCreateZoneRequest) (map[string]AvailabilityZone, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]AvailabilityZone
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AvailabilityZonesApiService.CreateZone")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/customers/{cUUID}/providers/{pUUID}/provider_regions/{rUUID}/region_zones"
	localVarPath = strings.Replace(localVarPath, "{"+"cUUID"+"}", _neturl.PathEscape(parameterToString(r.cUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pUUID"+"}", _neturl.PathEscape(parameterToString(r.pUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rUUID"+"}", _neturl.PathEscape(parameterToString(r.rUUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.azData == nil {
		return localVarReturnValue, nil, reportError("azData is required and must be specified")
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
	localVarPostBody = r.azData
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

type AvailabilityZonesApiApiDeleteAZRequest struct {
	ctx _context.Context
	ApiService *AvailabilityZonesApiService
	cUUID string
	pUUID string
	rUUID string
	azUUID string
	request *interface{}
}

func (r AvailabilityZonesApiApiDeleteAZRequest) Request(request interface{}) AvailabilityZonesApiApiDeleteAZRequest {
	r.request = &request
	return r
}

func (r AvailabilityZonesApiApiDeleteAZRequest) Execute() (YBPSuccess, *_nethttp.Response, error) {
	return r.ApiService.DeleteAZExecute(r)
}

/*
 * DeleteAZ Delete an availability zone
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID
 * @param pUUID
 * @param rUUID
 * @param azUUID
 * @return AvailabilityZonesApiApiDeleteAZRequest
 */
func (a *AvailabilityZonesApiService) DeleteAZ(ctx _context.Context, cUUID string, pUUID string, rUUID string, azUUID string) AvailabilityZonesApiApiDeleteAZRequest {
	return AvailabilityZonesApiApiDeleteAZRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
		pUUID: pUUID,
		rUUID: rUUID,
		azUUID: azUUID,
	}
}

/*
 * Execute executes the request
 * @return YBPSuccess
 */
func (a *AvailabilityZonesApiService) DeleteAZExecute(r AvailabilityZonesApiApiDeleteAZRequest) (YBPSuccess, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  YBPSuccess
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AvailabilityZonesApiService.DeleteAZ")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/customers/{cUUID}/providers/{pUUID}/regions/{rUUID}/zones/{azUUID}"
	localVarPath = strings.Replace(localVarPath, "{"+"cUUID"+"}", _neturl.PathEscape(parameterToString(r.cUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pUUID"+"}", _neturl.PathEscape(parameterToString(r.pUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rUUID"+"}", _neturl.PathEscape(parameterToString(r.rUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"azUUID"+"}", _neturl.PathEscape(parameterToString(r.azUUID, "")), -1)

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

type AvailabilityZonesApiApiEditAZRequest struct {
	ctx _context.Context
	ApiService *AvailabilityZonesApiService
	cUUID string
	pUUID string
	rUUID string
	azUUID string
	azFormData *AvailabilityZoneFormData
	request *interface{}
}

func (r AvailabilityZonesApiApiEditAZRequest) AzFormData(azFormData AvailabilityZoneFormData) AvailabilityZonesApiApiEditAZRequest {
	r.azFormData = &azFormData
	return r
}
func (r AvailabilityZonesApiApiEditAZRequest) Request(request interface{}) AvailabilityZonesApiApiEditAZRequest {
	r.request = &request
	return r
}

func (r AvailabilityZonesApiApiEditAZRequest) Execute() (AvailabilityZone, *_nethttp.Response, error) {
	return r.ApiService.EditAZExecute(r)
}

/*
 * EditAZ Edit an Availabilty Zone - deprecated
 * <b style="color:#ff0000">Deprecated since YBA version 2.18.2.0.</b></p>Use /api/v1/customers/{cUUID}/provider/{pUUID}/provider_regions/:rUUID/region_zones/:zUUID instead
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID
 * @param pUUID
 * @param rUUID
 * @param azUUID
 * @return AvailabilityZonesApiApiEditAZRequest
 */
func (a *AvailabilityZonesApiService) EditAZ(ctx _context.Context, cUUID string, pUUID string, rUUID string, azUUID string) AvailabilityZonesApiApiEditAZRequest {
	return AvailabilityZonesApiApiEditAZRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
		pUUID: pUUID,
		rUUID: rUUID,
		azUUID: azUUID,
	}
}

/*
 * Execute executes the request
 * @return AvailabilityZone
 */
func (a *AvailabilityZonesApiService) EditAZExecute(r AvailabilityZonesApiApiEditAZRequest) (AvailabilityZone, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AvailabilityZone
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AvailabilityZonesApiService.EditAZ")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/customers/{cUUID}/providers/{pUUID}/regions/{rUUID}/zones/{azUUID}"
	localVarPath = strings.Replace(localVarPath, "{"+"cUUID"+"}", _neturl.PathEscape(parameterToString(r.cUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pUUID"+"}", _neturl.PathEscape(parameterToString(r.pUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rUUID"+"}", _neturl.PathEscape(parameterToString(r.rUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"azUUID"+"}", _neturl.PathEscape(parameterToString(r.azUUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.azFormData == nil {
		return localVarReturnValue, nil, reportError("azFormData is required and must be specified")
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
	localVarPostBody = r.azFormData
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

type AvailabilityZonesApiApiEditZoneRequest struct {
	ctx _context.Context
	ApiService *AvailabilityZonesApiService
	cUUID string
	pUUID string
	rUUID string
	azUUID string
	azData *AvailabilityZone
	request *interface{}
}

func (r AvailabilityZonesApiApiEditZoneRequest) AzData(azData AvailabilityZone) AvailabilityZonesApiApiEditZoneRequest {
	r.azData = &azData
	return r
}
func (r AvailabilityZonesApiApiEditZoneRequest) Request(request interface{}) AvailabilityZonesApiApiEditZoneRequest {
	r.request = &request
	return r
}

func (r AvailabilityZonesApiApiEditZoneRequest) Execute() (AvailabilityZone, *_nethttp.Response, error) {
	return r.ApiService.EditZoneExecute(r)
}

/*
 * EditZone Modify an availability zone
 * WARNING: This is a preview API that could change.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID
 * @param pUUID
 * @param rUUID
 * @param azUUID
 * @return AvailabilityZonesApiApiEditZoneRequest
 */
func (a *AvailabilityZonesApiService) EditZone(ctx _context.Context, cUUID string, pUUID string, rUUID string, azUUID string) AvailabilityZonesApiApiEditZoneRequest {
	return AvailabilityZonesApiApiEditZoneRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
		pUUID: pUUID,
		rUUID: rUUID,
		azUUID: azUUID,
	}
}

/*
 * Execute executes the request
 * @return AvailabilityZone
 */
func (a *AvailabilityZonesApiService) EditZoneExecute(r AvailabilityZonesApiApiEditZoneRequest) (AvailabilityZone, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AvailabilityZone
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AvailabilityZonesApiService.EditZone")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/customers/{cUUID}/providers/{pUUID}/provider_regions/{rUUID}/region_zones/{azUUID}"
	localVarPath = strings.Replace(localVarPath, "{"+"cUUID"+"}", _neturl.PathEscape(parameterToString(r.cUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pUUID"+"}", _neturl.PathEscape(parameterToString(r.pUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rUUID"+"}", _neturl.PathEscape(parameterToString(r.rUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"azUUID"+"}", _neturl.PathEscape(parameterToString(r.azUUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.azData == nil {
		return localVarReturnValue, nil, reportError("azData is required and must be specified")
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
	localVarPostBody = r.azData
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

type AvailabilityZonesApiApiListOfAZRequest struct {
	ctx _context.Context
	ApiService *AvailabilityZonesApiService
	cUUID string
	pUUID string
	rUUID string
}


func (r AvailabilityZonesApiApiListOfAZRequest) Execute() ([]AvailabilityZone, *_nethttp.Response, error) {
	return r.ApiService.ListOfAZExecute(r)
}

/*
 * ListOfAZ List availability zones
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID
 * @param pUUID
 * @param rUUID
 * @return AvailabilityZonesApiApiListOfAZRequest
 */
func (a *AvailabilityZonesApiService) ListOfAZ(ctx _context.Context, cUUID string, pUUID string, rUUID string) AvailabilityZonesApiApiListOfAZRequest {
	return AvailabilityZonesApiApiListOfAZRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
		pUUID: pUUID,
		rUUID: rUUID,
	}
}

/*
 * Execute executes the request
 * @return []AvailabilityZone
 */
func (a *AvailabilityZonesApiService) ListOfAZExecute(r AvailabilityZonesApiApiListOfAZRequest) ([]AvailabilityZone, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []AvailabilityZone
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AvailabilityZonesApiService.ListOfAZ")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/customers/{cUUID}/providers/{pUUID}/regions/{rUUID}/zones"
	localVarPath = strings.Replace(localVarPath, "{"+"cUUID"+"}", _neturl.PathEscape(parameterToString(r.cUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pUUID"+"}", _neturl.PathEscape(parameterToString(r.pUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rUUID"+"}", _neturl.PathEscape(parameterToString(r.rUUID, "")), -1)

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
