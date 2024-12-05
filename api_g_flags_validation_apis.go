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

// GFlagsValidationAPIsApiService GFlagsValidationAPIsApi service
type GFlagsValidationAPIsApiService service

type GFlagsValidationAPIsApiApiGetGFlagMetadataRequest struct {
	ctx _context.Context
	ApiService *GFlagsValidationAPIsApiService
	version string
	name *string
	server *string
}

func (r GFlagsValidationAPIsApiApiGetGFlagMetadataRequest) Name(name string) GFlagsValidationAPIsApiApiGetGFlagMetadataRequest {
	r.name = &name
	return r
}
func (r GFlagsValidationAPIsApiApiGetGFlagMetadataRequest) Server(server string) GFlagsValidationAPIsApiApiGetGFlagMetadataRequest {
	r.server = &server
	return r
}

func (r GFlagsValidationAPIsApiApiGetGFlagMetadataRequest) Execute() (GFlagDetails, *_nethttp.Response, error) {
	return r.ApiService.GetGFlagMetadataExecute(r)
}

/*
 * GetGFlagMetadata Get gflag metadata
 * WARNING: This is a preview API that could change.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param version
 * @return GFlagsValidationAPIsApiApiGetGFlagMetadataRequest
 */
func (a *GFlagsValidationAPIsApiService) GetGFlagMetadata(ctx _context.Context, version string) GFlagsValidationAPIsApiApiGetGFlagMetadataRequest {
	return GFlagsValidationAPIsApiApiGetGFlagMetadataRequest{
		ApiService: a,
		ctx: ctx,
		version: version,
	}
}

/*
 * Execute executes the request
 * @return GFlagDetails
 */
func (a *GFlagsValidationAPIsApiService) GetGFlagMetadataExecute(r GFlagsValidationAPIsApiApiGetGFlagMetadataRequest) (GFlagDetails, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GFlagDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GFlagsValidationAPIsApiService.GetGFlagMetadata")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/metadata/version/{version}/gflag"
	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", _neturl.PathEscape(parameterToString(r.version, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.name != nil {
		localVarQueryParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.server != nil {
		localVarQueryParams.Add("server", parameterToString(*r.server, ""))
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

type GFlagsValidationAPIsApiApiListGFlagsRequest struct {
	ctx _context.Context
	ApiService *GFlagsValidationAPIsApiService
	version string
	name *string
	server *string
	mostUsedGFlags *bool
}

func (r GFlagsValidationAPIsApiApiListGFlagsRequest) Name(name string) GFlagsValidationAPIsApiApiListGFlagsRequest {
	r.name = &name
	return r
}
func (r GFlagsValidationAPIsApiApiListGFlagsRequest) Server(server string) GFlagsValidationAPIsApiApiListGFlagsRequest {
	r.server = &server
	return r
}
func (r GFlagsValidationAPIsApiApiListGFlagsRequest) MostUsedGFlags(mostUsedGFlags bool) GFlagsValidationAPIsApiApiListGFlagsRequest {
	r.mostUsedGFlags = &mostUsedGFlags
	return r
}

func (r GFlagsValidationAPIsApiApiListGFlagsRequest) Execute() ([]GFlagDetails, *_nethttp.Response, error) {
	return r.ApiService.ListGFlagsExecute(r)
}

/*
 * ListGFlags List all gflags for a release
 * WARNING: This is a preview API that could change.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param version
 * @return GFlagsValidationAPIsApiApiListGFlagsRequest
 */
func (a *GFlagsValidationAPIsApiService) ListGFlags(ctx _context.Context, version string) GFlagsValidationAPIsApiApiListGFlagsRequest {
	return GFlagsValidationAPIsApiApiListGFlagsRequest{
		ApiService: a,
		ctx: ctx,
		version: version,
	}
}

/*
 * Execute executes the request
 * @return []GFlagDetails
 */
func (a *GFlagsValidationAPIsApiService) ListGFlagsExecute(r GFlagsValidationAPIsApiApiListGFlagsRequest) ([]GFlagDetails, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []GFlagDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GFlagsValidationAPIsApiService.ListGFlags")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/metadata/version/{version}/list_gflags"
	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", _neturl.PathEscape(parameterToString(r.version, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.name != nil {
		localVarQueryParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.server != nil {
		localVarQueryParams.Add("server", parameterToString(*r.server, ""))
	}
	if r.mostUsedGFlags != nil {
		localVarQueryParams.Add("mostUsedGFlags", parameterToString(*r.mostUsedGFlags, ""))
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

type GFlagsValidationAPIsApiApiValidateGFlagsRequest struct {
	ctx _context.Context
	ApiService *GFlagsValidationAPIsApiService
	version string
	gflagValidationFormData *GFlagsValidationFormData
	request *interface{}
}

func (r GFlagsValidationAPIsApiApiValidateGFlagsRequest) GflagValidationFormData(gflagValidationFormData GFlagsValidationFormData) GFlagsValidationAPIsApiApiValidateGFlagsRequest {
	r.gflagValidationFormData = &gflagValidationFormData
	return r
}
func (r GFlagsValidationAPIsApiApiValidateGFlagsRequest) Request(request interface{}) GFlagsValidationAPIsApiApiValidateGFlagsRequest {
	r.request = &request
	return r
}

func (r GFlagsValidationAPIsApiApiValidateGFlagsRequest) Execute() ([]GFlagsValidationResponse, *_nethttp.Response, error) {
	return r.ApiService.ValidateGFlagsExecute(r)
}

/*
 * ValidateGFlags Validate gflags
 * WARNING: This is a preview API that could change.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param version
 * @return GFlagsValidationAPIsApiApiValidateGFlagsRequest
 */
func (a *GFlagsValidationAPIsApiService) ValidateGFlags(ctx _context.Context, version string) GFlagsValidationAPIsApiApiValidateGFlagsRequest {
	return GFlagsValidationAPIsApiApiValidateGFlagsRequest{
		ApiService: a,
		ctx: ctx,
		version: version,
	}
}

/*
 * Execute executes the request
 * @return []GFlagsValidationResponse
 */
func (a *GFlagsValidationAPIsApiService) ValidateGFlagsExecute(r GFlagsValidationAPIsApiApiValidateGFlagsRequest) ([]GFlagsValidationResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []GFlagsValidationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GFlagsValidationAPIsApiService.ValidateGFlags")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/metadata/version/{version}/validate_gflags"
	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", _neturl.PathEscape(parameterToString(r.version, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.gflagValidationFormData == nil {
		return localVarReturnValue, nil, reportError("gflagValidationFormData is required and must be specified")
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
	localVarPostBody = r.gflagValidationFormData
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
