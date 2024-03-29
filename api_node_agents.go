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

// NodeAgentsApiService NodeAgentsApi service
type NodeAgentsApiService service

type NodeAgentsApiApiDownloadNodeAgentInstallerRequest struct {
	ctx _context.Context
	ApiService *NodeAgentsApiService
	downloadType *string
	os *string
	arch *string
}

func (r NodeAgentsApiApiDownloadNodeAgentInstallerRequest) DownloadType(downloadType string) NodeAgentsApiApiDownloadNodeAgentInstallerRequest {
	r.downloadType = &downloadType
	return r
}
func (r NodeAgentsApiApiDownloadNodeAgentInstallerRequest) Os(os string) NodeAgentsApiApiDownloadNodeAgentInstallerRequest {
	r.os = &os
	return r
}
func (r NodeAgentsApiApiDownloadNodeAgentInstallerRequest) Arch(arch string) NodeAgentsApiApiDownloadNodeAgentInstallerRequest {
	r.arch = &arch
	return r
}

func (r NodeAgentsApiApiDownloadNodeAgentInstallerRequest) Execute() (string, *_nethttp.Response, error) {
	return r.ApiService.DownloadNodeAgentInstallerExecute(r)
}

/*
 * DownloadNodeAgentInstaller Download Node Agent Installer or Package
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return NodeAgentsApiApiDownloadNodeAgentInstallerRequest
 */
func (a *NodeAgentsApiService) DownloadNodeAgentInstaller(ctx _context.Context) NodeAgentsApiApiDownloadNodeAgentInstallerRequest {
	return NodeAgentsApiApiDownloadNodeAgentInstallerRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return string
 */
func (a *NodeAgentsApiService) DownloadNodeAgentInstallerExecute(r NodeAgentsApiApiDownloadNodeAgentInstallerRequest) (string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NodeAgentsApiService.DownloadNodeAgentInstaller")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/node_agents/download"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.downloadType != nil {
		localVarQueryParams.Add("downloadType", parameterToString(*r.downloadType, ""))
	}
	if r.os != nil {
		localVarQueryParams.Add("os", parameterToString(*r.os, ""))
	}
	if r.arch != nil {
		localVarQueryParams.Add("arch", parameterToString(*r.arch, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/gzip", "application/x-sh"}

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

type NodeAgentsApiApiGetNodeAgentRequest struct {
	ctx _context.Context
	ApiService *NodeAgentsApiService
	cUUID string
	nUUID string
}


func (r NodeAgentsApiApiGetNodeAgentRequest) Execute() (NodeAgentResp, *_nethttp.Response, error) {
	return r.ApiService.GetNodeAgentExecute(r)
}

/*
 * GetNodeAgent Get Node Agent
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID
 * @param nUUID
 * @return NodeAgentsApiApiGetNodeAgentRequest
 */
func (a *NodeAgentsApiService) GetNodeAgent(ctx _context.Context, cUUID string, nUUID string) NodeAgentsApiApiGetNodeAgentRequest {
	return NodeAgentsApiApiGetNodeAgentRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
		nUUID: nUUID,
	}
}

/*
 * Execute executes the request
 * @return NodeAgentResp
 */
func (a *NodeAgentsApiService) GetNodeAgentExecute(r NodeAgentsApiApiGetNodeAgentRequest) (NodeAgentResp, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  NodeAgentResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NodeAgentsApiService.GetNodeAgent")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/customers/{cUUID}/node_agents/{nUUID}"
	localVarPath = strings.Replace(localVarPath, "{"+"cUUID"+"}", _neturl.PathEscape(parameterToString(r.cUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nUUID"+"}", _neturl.PathEscape(parameterToString(r.nUUID, "")), -1)

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

type NodeAgentsApiApiListNodeAgentsRequest struct {
	ctx _context.Context
	ApiService *NodeAgentsApiService
	cUUID string
	nodeIp *string
}

func (r NodeAgentsApiApiListNodeAgentsRequest) NodeIp(nodeIp string) NodeAgentsApiApiListNodeAgentsRequest {
	r.nodeIp = &nodeIp
	return r
}

func (r NodeAgentsApiApiListNodeAgentsRequest) Execute() ([]NodeAgentResp, *_nethttp.Response, error) {
	return r.ApiService.ListNodeAgentsExecute(r)
}

/*
 * ListNodeAgents List Node Agents
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID
 * @return NodeAgentsApiApiListNodeAgentsRequest
 */
func (a *NodeAgentsApiService) ListNodeAgents(ctx _context.Context, cUUID string) NodeAgentsApiApiListNodeAgentsRequest {
	return NodeAgentsApiApiListNodeAgentsRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
	}
}

/*
 * Execute executes the request
 * @return []NodeAgentResp
 */
func (a *NodeAgentsApiService) ListNodeAgentsExecute(r NodeAgentsApiApiListNodeAgentsRequest) ([]NodeAgentResp, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []NodeAgentResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NodeAgentsApiService.ListNodeAgents")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/customers/{cUUID}/node_agents"
	localVarPath = strings.Replace(localVarPath, "{"+"cUUID"+"}", _neturl.PathEscape(parameterToString(r.cUUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.nodeIp != nil {
		localVarQueryParams.Add("nodeIp", parameterToString(*r.nodeIp, ""))
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

type NodeAgentsApiApiPageListNodeAgentsRequest struct {
	ctx _context.Context
	ApiService *NodeAgentsApiService
	cUUID string
	pageNodeAgentRequest *NodeAgentPagedApiQuery
	request *interface{}
}

func (r NodeAgentsApiApiPageListNodeAgentsRequest) PageNodeAgentRequest(pageNodeAgentRequest NodeAgentPagedApiQuery) NodeAgentsApiApiPageListNodeAgentsRequest {
	r.pageNodeAgentRequest = &pageNodeAgentRequest
	return r
}
func (r NodeAgentsApiApiPageListNodeAgentsRequest) Request(request interface{}) NodeAgentsApiApiPageListNodeAgentsRequest {
	r.request = &request
	return r
}

func (r NodeAgentsApiApiPageListNodeAgentsRequest) Execute() (NodeAgentPagedApiResponse, *_nethttp.Response, error) {
	return r.ApiService.PageListNodeAgentsExecute(r)
}

/*
 * PageListNodeAgents List Node Agents (paginated)
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID
 * @return NodeAgentsApiApiPageListNodeAgentsRequest
 */
func (a *NodeAgentsApiService) PageListNodeAgents(ctx _context.Context, cUUID string) NodeAgentsApiApiPageListNodeAgentsRequest {
	return NodeAgentsApiApiPageListNodeAgentsRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
	}
}

/*
 * Execute executes the request
 * @return NodeAgentPagedApiResponse
 */
func (a *NodeAgentsApiService) PageListNodeAgentsExecute(r NodeAgentsApiApiPageListNodeAgentsRequest) (NodeAgentPagedApiResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  NodeAgentPagedApiResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NodeAgentsApiService.PageListNodeAgents")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/customers/{cUUID}/node_agents/page"
	localVarPath = strings.Replace(localVarPath, "{"+"cUUID"+"}", _neturl.PathEscape(parameterToString(r.cUUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.pageNodeAgentRequest == nil {
		return localVarReturnValue, nil, reportError("pageNodeAgentRequest is required and must be specified")
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
	localVarPostBody = r.pageNodeAgentRequest
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

type NodeAgentsApiApiReinstallNodeAgentRequest struct {
	ctx _context.Context
	ApiService *NodeAgentsApiService
	cUUID string
	uniUUID string
	reinstallNodeAgentForm *ReinstallNodeAgentForm
	request *interface{}
}

func (r NodeAgentsApiApiReinstallNodeAgentRequest) ReinstallNodeAgentForm(reinstallNodeAgentForm ReinstallNodeAgentForm) NodeAgentsApiApiReinstallNodeAgentRequest {
	r.reinstallNodeAgentForm = &reinstallNodeAgentForm
	return r
}
func (r NodeAgentsApiApiReinstallNodeAgentRequest) Request(request interface{}) NodeAgentsApiApiReinstallNodeAgentRequest {
	r.request = &request
	return r
}

func (r NodeAgentsApiApiReinstallNodeAgentRequest) Execute() (NodeAgent, *_nethttp.Response, error) {
	return r.ApiService.ReinstallNodeAgentExecute(r)
}

/*
 * ReinstallNodeAgent Reinstall Node Agent
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID
 * @param uniUUID
 * @return NodeAgentsApiApiReinstallNodeAgentRequest
 */
func (a *NodeAgentsApiService) ReinstallNodeAgent(ctx _context.Context, cUUID string, uniUUID string) NodeAgentsApiApiReinstallNodeAgentRequest {
	return NodeAgentsApiApiReinstallNodeAgentRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
		uniUUID: uniUUID,
	}
}

/*
 * Execute executes the request
 * @return NodeAgent
 */
func (a *NodeAgentsApiService) ReinstallNodeAgentExecute(r NodeAgentsApiApiReinstallNodeAgentRequest) (NodeAgent, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  NodeAgent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NodeAgentsApiService.ReinstallNodeAgent")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/customers/{cUUID}/universes/{uniUUID}/node_agents"
	localVarPath = strings.Replace(localVarPath, "{"+"cUUID"+"}", _neturl.PathEscape(parameterToString(r.cUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"uniUUID"+"}", _neturl.PathEscape(parameterToString(r.uniUUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.reinstallNodeAgentForm == nil {
		return localVarReturnValue, nil, reportError("reinstallNodeAgentForm is required and must be specified")
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
	localVarPostBody = r.reinstallNodeAgentForm
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
