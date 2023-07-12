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

type LDAPRoleManagementApiApiListLdapDnToYbaRolesRequest struct {
	ctx _context.Context
	ApiService *LDAPRoleManagementApiService
	cUUID string
}


func (r LDAPRoleManagementApiApiListLdapDnToYbaRolesRequest) Execute() (LdapDnToYbaRoleData, *_nethttp.Response, error) {
	return r.ApiService.ListLdapDnToYbaRolesExecute(r)
}

/*
 * ListLdapDnToYbaRoles List LDAP Mappings
 * List LDAP Mappings
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID
 * @return LDAPRoleManagementApiApiListLdapDnToYbaRolesRequest
 */
func (a *LDAPRoleManagementApiService) ListLdapDnToYbaRoles(ctx _context.Context, cUUID string) LDAPRoleManagementApiApiListLdapDnToYbaRolesRequest {
	return LDAPRoleManagementApiApiListLdapDnToYbaRolesRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
	}
}

/*
 * Execute executes the request
 * @return LdapDnToYbaRoleData
 */
func (a *LDAPRoleManagementApiService) ListLdapDnToYbaRolesExecute(r LDAPRoleManagementApiApiListLdapDnToYbaRolesRequest) (LdapDnToYbaRoleData, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  LdapDnToYbaRoleData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LDAPRoleManagementApiService.ListLdapDnToYbaRoles")
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

type LDAPRoleManagementApiApiSetLdapDnToYbaRolesRequest struct {
	ctx _context.Context
	ApiService *LDAPRoleManagementApiService
	cUUID string
	ldapMappings *LdapDnToYbaRoleData
	request *interface{}
}

func (r LDAPRoleManagementApiApiSetLdapDnToYbaRolesRequest) LdapMappings(ldapMappings LdapDnToYbaRoleData) LDAPRoleManagementApiApiSetLdapDnToYbaRolesRequest {
	r.ldapMappings = &ldapMappings
	return r
}
func (r LDAPRoleManagementApiApiSetLdapDnToYbaRolesRequest) Request(request interface{}) LDAPRoleManagementApiApiSetLdapDnToYbaRolesRequest {
	r.request = &request
	return r
}

func (r LDAPRoleManagementApiApiSetLdapDnToYbaRolesRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.SetLdapDnToYbaRolesExecute(r)
}

/*
 * SetLdapDnToYbaRoles Set LDAP Mappings
 * Set LDAP Mappings
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID
 * @return LDAPRoleManagementApiApiSetLdapDnToYbaRolesRequest
 */
func (a *LDAPRoleManagementApiService) SetLdapDnToYbaRoles(ctx _context.Context, cUUID string) LDAPRoleManagementApiApiSetLdapDnToYbaRolesRequest {
	return LDAPRoleManagementApiApiSetLdapDnToYbaRolesRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
	}
}

/*
 * Execute executes the request
 */
func (a *LDAPRoleManagementApiService) SetLdapDnToYbaRolesExecute(r LDAPRoleManagementApiApiSetLdapDnToYbaRolesRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LDAPRoleManagementApiService.SetLdapDnToYbaRoles")
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
