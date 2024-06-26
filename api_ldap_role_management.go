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

// LDAPRoleManagementApiService LDAPRoleManagementApi service
type LDAPRoleManagementApiService service

type LDAPRoleManagementApiApiSyncLdapUniverseRequest struct {
	ctx _context.Context
	ApiService *LDAPRoleManagementApiService
	cUUID string
	univUUID string
	syncLdapUniverse *LdapUnivSyncFormData
	request *interface{}
}

func (r LDAPRoleManagementApiApiSyncLdapUniverseRequest) SyncLdapUniverse(syncLdapUniverse LdapUnivSyncFormData) LDAPRoleManagementApiApiSyncLdapUniverseRequest {
	r.syncLdapUniverse = &syncLdapUniverse
	return r
}
func (r LDAPRoleManagementApiApiSyncLdapUniverseRequest) Request(request interface{}) LDAPRoleManagementApiApiSyncLdapUniverseRequest {
	r.request = &request
	return r
}

func (r LDAPRoleManagementApiApiSyncLdapUniverseRequest) Execute() (YBPTask, *_nethttp.Response, error) {
	return r.ApiService.SyncLdapUniverseExecute(r)
}

/*
 * SyncLdapUniverse Perform an LDAP users sync on the universe
 * WARNING: This is a preview API that could change.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID
 * @param univUUID
 * @return LDAPRoleManagementApiApiSyncLdapUniverseRequest
 */
func (a *LDAPRoleManagementApiService) SyncLdapUniverse(ctx _context.Context, cUUID string, univUUID string) LDAPRoleManagementApiApiSyncLdapUniverseRequest {
	return LDAPRoleManagementApiApiSyncLdapUniverseRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
		univUUID: univUUID,
	}
}

/*
 * Execute executes the request
 * @return YBPTask
 */
func (a *LDAPRoleManagementApiService) SyncLdapUniverseExecute(r LDAPRoleManagementApiApiSyncLdapUniverseRequest) (YBPTask, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  YBPTask
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LDAPRoleManagementApiService.SyncLdapUniverse")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/customers/{cUUID}/universes/{univUUID}/ldap_roles_sync"
	localVarPath = strings.Replace(localVarPath, "{"+"cUUID"+"}", _neturl.PathEscape(parameterToString(r.cUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"univUUID"+"}", _neturl.PathEscape(parameterToString(r.univUUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.syncLdapUniverse == nil {
		return localVarReturnValue, nil, reportError("syncLdapUniverse is required and must be specified")
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
	localVarPostBody = r.syncLdapUniverse
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
