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

// LDAPOIDCRoleManagementApiService LDAPOIDCRoleManagementApi service
type LDAPOIDCRoleManagementApiService service

type LDAPOIDCRoleManagementApiApiDeleteOidcGroupMappingRequest struct {
	ctx _context.Context
	ApiService *LDAPOIDCRoleManagementApiService
	cUUID string
	groupName string
	request *interface{}
}

func (r LDAPOIDCRoleManagementApiApiDeleteOidcGroupMappingRequest) Request(request interface{}) LDAPOIDCRoleManagementApiApiDeleteOidcGroupMappingRequest {
	r.request = &request
	return r
}

func (r LDAPOIDCRoleManagementApiApiDeleteOidcGroupMappingRequest) Execute() (YBPSuccess, *_nethttp.Response, error) {
	return r.ApiService.DeleteOidcGroupMappingExecute(r)
}

/*
 * DeleteOidcGroupMapping Delete a OIDC group mapping
 * <b style="color:#ff0000">Deprecated since YBA version 2024.2.0.0.</b> Please use the v2 /auth/group-mappings/{groupUUID} DELETE endpoint instead
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID
 * @param groupName
 * @return LDAPOIDCRoleManagementApiApiDeleteOidcGroupMappingRequest
 */
func (a *LDAPOIDCRoleManagementApiService) DeleteOidcGroupMapping(ctx _context.Context, cUUID string, groupName string) LDAPOIDCRoleManagementApiApiDeleteOidcGroupMappingRequest {
	return LDAPOIDCRoleManagementApiApiDeleteOidcGroupMappingRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
		groupName: groupName,
	}
}

/*
 * Execute executes the request
 * @return YBPSuccess
 */
func (a *LDAPOIDCRoleManagementApiService) DeleteOidcGroupMappingExecute(r LDAPOIDCRoleManagementApiApiDeleteOidcGroupMappingRequest) (YBPSuccess, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  YBPSuccess
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LDAPOIDCRoleManagementApiService.DeleteOidcGroupMapping")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/customers/{cUUID}/oidc_mappings/{groupName}"
	localVarPath = strings.Replace(localVarPath, "{"+"cUUID"+"}", _neturl.PathEscape(parameterToString(r.cUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"groupName"+"}", _neturl.PathEscape(parameterToString(r.groupName, "")), -1)

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

type LDAPOIDCRoleManagementApiApiListLdapDnToYbaRolesRequest struct {
	ctx _context.Context
	ApiService *LDAPOIDCRoleManagementApiService
	cUUID string
}


func (r LDAPOIDCRoleManagementApiApiListLdapDnToYbaRolesRequest) Execute() (LdapDnToYbaRoleData, *_nethttp.Response, error) {
	return r.ApiService.ListLdapDnToYbaRolesExecute(r)
}

/*
 * ListLdapDnToYbaRoles List LDAP Mappings
 * <b style="color:#ff0000">Deprecated since YBA version 2024.2.0.0.</b> Please use the v2 /auth/group-mappings GET endpoint instead. Note that this API will not return the custom roles assigned to groups via the new /api/v2/customers/{cUUID}/auth/group-mappings API.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID
 * @return LDAPOIDCRoleManagementApiApiListLdapDnToYbaRolesRequest
 */
func (a *LDAPOIDCRoleManagementApiService) ListLdapDnToYbaRoles(ctx _context.Context, cUUID string) LDAPOIDCRoleManagementApiApiListLdapDnToYbaRolesRequest {
	return LDAPOIDCRoleManagementApiApiListLdapDnToYbaRolesRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
	}
}

/*
 * Execute executes the request
 * @return LdapDnToYbaRoleData
 */
func (a *LDAPOIDCRoleManagementApiService) ListLdapDnToYbaRolesExecute(r LDAPOIDCRoleManagementApiApiListLdapDnToYbaRolesRequest) (LdapDnToYbaRoleData, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  LdapDnToYbaRoleData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LDAPOIDCRoleManagementApiService.ListLdapDnToYbaRoles")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/customers/{cUUID}/ldap_mappings"
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

type LDAPOIDCRoleManagementApiApiListOidcGroupToYbaRolesRequest struct {
	ctx _context.Context
	ApiService *LDAPOIDCRoleManagementApiService
	cUUID string
}


func (r LDAPOIDCRoleManagementApiApiListOidcGroupToYbaRolesRequest) Execute() (OidcGroupToYbaRolesData, *_nethttp.Response, error) {
	return r.ApiService.ListOidcGroupToYbaRolesExecute(r)
}

/*
 * ListOidcGroupToYbaRoles List OIDC Group Mappings
 * <b style="color:#ff0000">Deprecated since YBA version 2024.2.0.0.</b> Please use the v2 /auth/group-mappings GET endpoint instead. Note that this API will not return the custom roles assigned to groups via the new /api/v2/customers/{cUUID}/auth/group-mappings API.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID
 * @return LDAPOIDCRoleManagementApiApiListOidcGroupToYbaRolesRequest
 */
func (a *LDAPOIDCRoleManagementApiService) ListOidcGroupToYbaRoles(ctx _context.Context, cUUID string) LDAPOIDCRoleManagementApiApiListOidcGroupToYbaRolesRequest {
	return LDAPOIDCRoleManagementApiApiListOidcGroupToYbaRolesRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
	}
}

/*
 * Execute executes the request
 * @return OidcGroupToYbaRolesData
 */
func (a *LDAPOIDCRoleManagementApiService) ListOidcGroupToYbaRolesExecute(r LDAPOIDCRoleManagementApiApiListOidcGroupToYbaRolesRequest) (OidcGroupToYbaRolesData, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  OidcGroupToYbaRolesData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LDAPOIDCRoleManagementApiService.ListOidcGroupToYbaRoles")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/customers/{cUUID}/oidc_mappings"
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

type LDAPOIDCRoleManagementApiApiMapOidcGroupToYbaRolesRequest struct {
	ctx _context.Context
	ApiService *LDAPOIDCRoleManagementApiService
	cUUID string
	oidcMappings *OidcGroupToYbaRolesData
	request *interface{}
}

func (r LDAPOIDCRoleManagementApiApiMapOidcGroupToYbaRolesRequest) OidcMappings(oidcMappings OidcGroupToYbaRolesData) LDAPOIDCRoleManagementApiApiMapOidcGroupToYbaRolesRequest {
	r.oidcMappings = &oidcMappings
	return r
}
func (r LDAPOIDCRoleManagementApiApiMapOidcGroupToYbaRolesRequest) Request(request interface{}) LDAPOIDCRoleManagementApiApiMapOidcGroupToYbaRolesRequest {
	r.request = &request
	return r
}

func (r LDAPOIDCRoleManagementApiApiMapOidcGroupToYbaRolesRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.MapOidcGroupToYbaRolesExecute(r)
}

/*
 * MapOidcGroupToYbaRoles Set OIDC Mappings
 * <b style="color:#ff0000">Deprecated since YBA version 2024.2.0.0.</b> Please use the v2 /auth/group-mappings PUT endpoint instead
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID
 * @return LDAPOIDCRoleManagementApiApiMapOidcGroupToYbaRolesRequest
 */
func (a *LDAPOIDCRoleManagementApiService) MapOidcGroupToYbaRoles(ctx _context.Context, cUUID string) LDAPOIDCRoleManagementApiApiMapOidcGroupToYbaRolesRequest {
	return LDAPOIDCRoleManagementApiApiMapOidcGroupToYbaRolesRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
	}
}

/*
 * Execute executes the request
 */
func (a *LDAPOIDCRoleManagementApiService) MapOidcGroupToYbaRolesExecute(r LDAPOIDCRoleManagementApiApiMapOidcGroupToYbaRolesRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LDAPOIDCRoleManagementApiService.MapOidcGroupToYbaRoles")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/customers/{cUUID}/oidc_mappings"
	localVarPath = strings.Replace(localVarPath, "{"+"cUUID"+"}", _neturl.PathEscape(parameterToString(r.cUUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.oidcMappings == nil {
		return nil, reportError("oidcMappings is required and must be specified")
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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.oidcMappings
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

type LDAPOIDCRoleManagementApiApiSetLdapDnToYbaRolesRequest struct {
	ctx _context.Context
	ApiService *LDAPOIDCRoleManagementApiService
	cUUID string
	ldapMappings *LdapDnToYbaRoleData
	request *interface{}
}

func (r LDAPOIDCRoleManagementApiApiSetLdapDnToYbaRolesRequest) LdapMappings(ldapMappings LdapDnToYbaRoleData) LDAPOIDCRoleManagementApiApiSetLdapDnToYbaRolesRequest {
	r.ldapMappings = &ldapMappings
	return r
}
func (r LDAPOIDCRoleManagementApiApiSetLdapDnToYbaRolesRequest) Request(request interface{}) LDAPOIDCRoleManagementApiApiSetLdapDnToYbaRolesRequest {
	r.request = &request
	return r
}

func (r LDAPOIDCRoleManagementApiApiSetLdapDnToYbaRolesRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.SetLdapDnToYbaRolesExecute(r)
}

/*
 * SetLdapDnToYbaRoles Set LDAP Mappings
 * <b style="color:#ff0000">Deprecated since YBA version 2024.2.0.0.</b> Please use the v2 /auth/group-mappings PUT endpoint instead
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID
 * @return LDAPOIDCRoleManagementApiApiSetLdapDnToYbaRolesRequest
 */
func (a *LDAPOIDCRoleManagementApiService) SetLdapDnToYbaRoles(ctx _context.Context, cUUID string) LDAPOIDCRoleManagementApiApiSetLdapDnToYbaRolesRequest {
	return LDAPOIDCRoleManagementApiApiSetLdapDnToYbaRolesRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
	}
}

/*
 * Execute executes the request
 */
func (a *LDAPOIDCRoleManagementApiService) SetLdapDnToYbaRolesExecute(r LDAPOIDCRoleManagementApiApiSetLdapDnToYbaRolesRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LDAPOIDCRoleManagementApiService.SetLdapDnToYbaRoles")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/customers/{cUUID}/ldap_mappings"
	localVarPath = strings.Replace(localVarPath, "{"+"cUUID"+"}", _neturl.PathEscape(parameterToString(r.cUUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.ldapMappings == nil {
		return nil, reportError("ldapMappings is required and must be specified")
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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.ldapMappings
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
