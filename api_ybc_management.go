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

// YbcManagementApiService YbcManagementApi service
type YbcManagementApiService service

type YbcManagementApiApiDisableRequest struct {
	ctx _context.Context
	ApiService *YbcManagementApiService
	cUUID string
	uniUUID string
}


func (r YbcManagementApiApiDisableRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DisableExecute(r)
}

/*
 * Disable Method for Disable
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID
 * @param uniUUID
 * @return YbcManagementApiApiDisableRequest
 */
func (a *YbcManagementApiService) Disable(ctx _context.Context, cUUID string, uniUUID string) YbcManagementApiApiDisableRequest {
	return YbcManagementApiApiDisableRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
		uniUUID: uniUUID,
	}
}

/*
 * Execute executes the request
 */
func (a *YbcManagementApiService) DisableExecute(r YbcManagementApiApiDisableRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "YbcManagementApiService.Disable")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/customers/{cUUID}/universes/{uniUUID}/ybc/disable"
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
	localVarHTTPHeaderAccepts := []string{}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type YbcManagementApiApiInstallRequest struct {
	ctx _context.Context
	ApiService *YbcManagementApiService
	cUUID string
	uniUUID string
	ybcVersion *string
}

func (r YbcManagementApiApiInstallRequest) YbcVersion(ybcVersion string) YbcManagementApiApiInstallRequest {
	r.ybcVersion = &ybcVersion
	return r
}

func (r YbcManagementApiApiInstallRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.InstallExecute(r)
}

/*
 * Install Method for Install
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID
 * @param uniUUID
 * @return YbcManagementApiApiInstallRequest
 */
func (a *YbcManagementApiService) Install(ctx _context.Context, cUUID string, uniUUID string) YbcManagementApiApiInstallRequest {
	return YbcManagementApiApiInstallRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
		uniUUID: uniUUID,
	}
}

/*
 * Execute executes the request
 */
func (a *YbcManagementApiService) InstallExecute(r YbcManagementApiApiInstallRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "YbcManagementApiService.Install")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/customers/{cUUID}/universes/{uniUUID}/ybc/install"
	localVarPath = strings.Replace(localVarPath, "{"+"cUUID"+"}", _neturl.PathEscape(parameterToString(r.cUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"uniUUID"+"}", _neturl.PathEscape(parameterToString(r.uniUUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.ybcVersion != nil {
		localVarQueryParams.Add("ybcVersion", parameterToString(*r.ybcVersion, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type YbcManagementApiApiUpgradeRequest struct {
	ctx _context.Context
	ApiService *YbcManagementApiService
	cUUID string
	uniUUID string
	ybcVersion *string
}

func (r YbcManagementApiApiUpgradeRequest) YbcVersion(ybcVersion string) YbcManagementApiApiUpgradeRequest {
	r.ybcVersion = &ybcVersion
	return r
}

func (r YbcManagementApiApiUpgradeRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.UpgradeExecute(r)
}

/*
 * Upgrade Method for Upgrade
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID
 * @param uniUUID
 * @return YbcManagementApiApiUpgradeRequest
 */
func (a *YbcManagementApiService) Upgrade(ctx _context.Context, cUUID string, uniUUID string) YbcManagementApiApiUpgradeRequest {
	return YbcManagementApiApiUpgradeRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
		uniUUID: uniUUID,
	}
}

/*
 * Execute executes the request
 */
func (a *YbcManagementApiService) UpgradeExecute(r YbcManagementApiApiUpgradeRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "YbcManagementApiService.Upgrade")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/customers/{cUUID}/universes/{uniUUID}/ybc/upgrade"
	localVarPath = strings.Replace(localVarPath, "{"+"cUUID"+"}", _neturl.PathEscape(parameterToString(r.cUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"uniUUID"+"}", _neturl.PathEscape(parameterToString(r.uniUUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.ybcVersion != nil {
		localVarQueryParams.Add("ybcVersion", parameterToString(*r.ybcVersion, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}