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

// UserManagementApiService UserManagementApi service
type UserManagementApiService service

type UserManagementApiApiCreateUserRequest struct {
	ctx _context.Context
	ApiService *UserManagementApiService
	cUUID string
	user *UserRegistrationData
	request *interface{}
}

func (r UserManagementApiApiCreateUserRequest) User(user UserRegistrationData) UserManagementApiApiCreateUserRequest {
	r.user = &user
	return r
}
func (r UserManagementApiApiCreateUserRequest) Request(request interface{}) UserManagementApiApiCreateUserRequest {
	r.request = &request
	return r
}

func (r UserManagementApiApiCreateUserRequest) Execute() (UserWithFeatures, *_nethttp.Response, error) {
	return r.ApiService.CreateUserExecute(r)
}

/*
 * CreateUser Create a user
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID
 * @return UserManagementApiApiCreateUserRequest
 */
func (a *UserManagementApiService) CreateUser(ctx _context.Context, cUUID string) UserManagementApiApiCreateUserRequest {
	return UserManagementApiApiCreateUserRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
	}
}

/*
 * Execute executes the request
 * @return UserWithFeatures
 */
func (a *UserManagementApiService) CreateUserExecute(r UserManagementApiApiCreateUserRequest) (UserWithFeatures, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  UserWithFeatures
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserManagementApiService.CreateUser")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/customers/{cUUID}/users"
	localVarPath = strings.Replace(localVarPath, "{"+"cUUID"+"}", _neturl.PathEscape(parameterToString(r.cUUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.user == nil {
		return localVarReturnValue, nil, reportError("user is required and must be specified")
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
	localVarPostBody = r.user
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

type UserManagementApiApiDeleteUserRequest struct {
	ctx _context.Context
	ApiService *UserManagementApiService
	cUUID string
	uUUID string
	request *interface{}
}

func (r UserManagementApiApiDeleteUserRequest) Request(request interface{}) UserManagementApiApiDeleteUserRequest {
	r.request = &request
	return r
}

func (r UserManagementApiApiDeleteUserRequest) Execute() (YBPSuccess, *_nethttp.Response, error) {
	return r.ApiService.DeleteUserExecute(r)
}

/*
 * DeleteUser Delete a user
 * Deletes the specified user. Note that you can't delete a customer's primary user.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID
 * @param uUUID
 * @return UserManagementApiApiDeleteUserRequest
 */
func (a *UserManagementApiService) DeleteUser(ctx _context.Context, cUUID string, uUUID string) UserManagementApiApiDeleteUserRequest {
	return UserManagementApiApiDeleteUserRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
		uUUID: uUUID,
	}
}

/*
 * Execute executes the request
 * @return YBPSuccess
 */
func (a *UserManagementApiService) DeleteUserExecute(r UserManagementApiApiDeleteUserRequest) (YBPSuccess, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  YBPSuccess
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserManagementApiService.DeleteUser")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/customers/{cUUID}/users/{uUUID}"
	localVarPath = strings.Replace(localVarPath, "{"+"cUUID"+"}", _neturl.PathEscape(parameterToString(r.cUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"uUUID"+"}", _neturl.PathEscape(parameterToString(r.uUUID, "")), -1)

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

type UserManagementApiApiGetUserDetailsRequest struct {
	ctx _context.Context
	ApiService *UserManagementApiService
	cUUID string
	uUUID string
}


func (r UserManagementApiApiGetUserDetailsRequest) Execute() (UserWithFeatures, *_nethttp.Response, error) {
	return r.ApiService.GetUserDetailsExecute(r)
}

/*
 * GetUserDetails Get a user's details
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID
 * @param uUUID
 * @return UserManagementApiApiGetUserDetailsRequest
 */
func (a *UserManagementApiService) GetUserDetails(ctx _context.Context, cUUID string, uUUID string) UserManagementApiApiGetUserDetailsRequest {
	return UserManagementApiApiGetUserDetailsRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
		uUUID: uUUID,
	}
}

/*
 * Execute executes the request
 * @return UserWithFeatures
 */
func (a *UserManagementApiService) GetUserDetailsExecute(r UserManagementApiApiGetUserDetailsRequest) (UserWithFeatures, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  UserWithFeatures
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserManagementApiService.GetUserDetails")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/customers/{cUUID}/users/{uUUID}"
	localVarPath = strings.Replace(localVarPath, "{"+"cUUID"+"}", _neturl.PathEscape(parameterToString(r.cUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"uUUID"+"}", _neturl.PathEscape(parameterToString(r.uUUID, "")), -1)

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

type UserManagementApiApiListUsersRequest struct {
	ctx _context.Context
	ApiService *UserManagementApiService
	cUUID string
}


func (r UserManagementApiApiListUsersRequest) Execute() ([]UserWithFeatures, *_nethttp.Response, error) {
	return r.ApiService.ListUsersExecute(r)
}

/*
 * ListUsers List all users
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID
 * @return UserManagementApiApiListUsersRequest
 */
func (a *UserManagementApiService) ListUsers(ctx _context.Context, cUUID string) UserManagementApiApiListUsersRequest {
	return UserManagementApiApiListUsersRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
	}
}

/*
 * Execute executes the request
 * @return []UserWithFeatures
 */
func (a *UserManagementApiService) ListUsersExecute(r UserManagementApiApiListUsersRequest) ([]UserWithFeatures, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []UserWithFeatures
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserManagementApiService.ListUsers")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/customers/{cUUID}/users"
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

type UserManagementApiApiUpdateUserPasswordRequest struct {
	ctx _context.Context
	ApiService *UserManagementApiService
	cUUID string
	uUUID string
	users *UserRegistrationData
	request *interface{}
}

func (r UserManagementApiApiUpdateUserPasswordRequest) Users(users UserRegistrationData) UserManagementApiApiUpdateUserPasswordRequest {
	r.users = &users
	return r
}
func (r UserManagementApiApiUpdateUserPasswordRequest) Request(request interface{}) UserManagementApiApiUpdateUserPasswordRequest {
	r.request = &request
	return r
}

func (r UserManagementApiApiUpdateUserPasswordRequest) Execute() (YBPSuccess, *_nethttp.Response, error) {
	return r.ApiService.UpdateUserPasswordExecute(r)
}

/*
 * UpdateUserPassword Change a user's password
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID
 * @param uUUID
 * @return UserManagementApiApiUpdateUserPasswordRequest
 */
func (a *UserManagementApiService) UpdateUserPassword(ctx _context.Context, cUUID string, uUUID string) UserManagementApiApiUpdateUserPasswordRequest {
	return UserManagementApiApiUpdateUserPasswordRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
		uUUID: uUUID,
	}
}

/*
 * Execute executes the request
 * @return YBPSuccess
 */
func (a *UserManagementApiService) UpdateUserPasswordExecute(r UserManagementApiApiUpdateUserPasswordRequest) (YBPSuccess, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  YBPSuccess
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserManagementApiService.UpdateUserPassword")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/customers/{cUUID}/users/{uUUID}/change_password"
	localVarPath = strings.Replace(localVarPath, "{"+"cUUID"+"}", _neturl.PathEscape(parameterToString(r.cUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"uUUID"+"}", _neturl.PathEscape(parameterToString(r.uUUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.users == nil {
		return localVarReturnValue, nil, reportError("users is required and must be specified")
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
	localVarPostBody = r.users
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

type UserManagementApiApiUpdateUserProfileRequest struct {
	ctx _context.Context
	ApiService *UserManagementApiService
	cUUID string
	uUUID string
	users *UserProfileData
	request *interface{}
}

func (r UserManagementApiApiUpdateUserProfileRequest) Users(users UserProfileData) UserManagementApiApiUpdateUserProfileRequest {
	r.users = &users
	return r
}
func (r UserManagementApiApiUpdateUserProfileRequest) Request(request interface{}) UserManagementApiApiUpdateUserProfileRequest {
	r.request = &request
	return r
}

func (r UserManagementApiApiUpdateUserProfileRequest) Execute() (Users, *_nethttp.Response, error) {
	return r.ApiService.UpdateUserProfileExecute(r)
}

/*
 * UpdateUserProfile Update a user's profile
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID
 * @param uUUID
 * @return UserManagementApiApiUpdateUserProfileRequest
 */
func (a *UserManagementApiService) UpdateUserProfile(ctx _context.Context, cUUID string, uUUID string) UserManagementApiApiUpdateUserProfileRequest {
	return UserManagementApiApiUpdateUserProfileRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
		uUUID: uUUID,
	}
}

/*
 * Execute executes the request
 * @return Users
 */
func (a *UserManagementApiService) UpdateUserProfileExecute(r UserManagementApiApiUpdateUserProfileRequest) (Users, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Users
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserManagementApiService.UpdateUserProfile")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/customers/{cUUID}/users/{uUUID}/update_profile"
	localVarPath = strings.Replace(localVarPath, "{"+"cUUID"+"}", _neturl.PathEscape(parameterToString(r.cUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"uUUID"+"}", _neturl.PathEscape(parameterToString(r.uUUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.users == nil {
		return localVarReturnValue, nil, reportError("users is required and must be specified")
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
	localVarPostBody = r.users
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

type UserManagementApiApiUpdateUserRoleRequest struct {
	ctx _context.Context
	ApiService *UserManagementApiService
	cUUID string
	uUUID string
	role *string
	request *interface{}
}

func (r UserManagementApiApiUpdateUserRoleRequest) Role(role string) UserManagementApiApiUpdateUserRoleRequest {
	r.role = &role
	return r
}
func (r UserManagementApiApiUpdateUserRoleRequest) Request(request interface{}) UserManagementApiApiUpdateUserRoleRequest {
	r.request = &request
	return r
}

func (r UserManagementApiApiUpdateUserRoleRequest) Execute() (YBPSuccess, *_nethttp.Response, error) {
	return r.ApiService.UpdateUserRoleExecute(r)
}

/*
 * UpdateUserRole Change a user's role
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID
 * @param uUUID
 * @return UserManagementApiApiUpdateUserRoleRequest
 */
func (a *UserManagementApiService) UpdateUserRole(ctx _context.Context, cUUID string, uUUID string) UserManagementApiApiUpdateUserRoleRequest {
	return UserManagementApiApiUpdateUserRoleRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
		uUUID: uUUID,
	}
}

/*
 * Execute executes the request
 * @return YBPSuccess
 */
func (a *UserManagementApiService) UpdateUserRoleExecute(r UserManagementApiApiUpdateUserRoleRequest) (YBPSuccess, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  YBPSuccess
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserManagementApiService.UpdateUserRole")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/customers/{cUUID}/users/{uUUID}"
	localVarPath = strings.Replace(localVarPath, "{"+"cUUID"+"}", _neturl.PathEscape(parameterToString(r.cUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"uUUID"+"}", _neturl.PathEscape(parameterToString(r.uUUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.role != nil {
		localVarQueryParams.Add("role", parameterToString(*r.role, ""))
	}
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
