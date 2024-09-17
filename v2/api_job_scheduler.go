/*
 * YugabyteDB Anywhere V2 APIs
 *
 * An improved set of APIs for managing YugabyteDB Anywhere
 *
 * API version: v2
 * Contact: support@yugabyte.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

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

// JobSchedulerApiService JobSchedulerApi service
type JobSchedulerApiService service

type JobSchedulerApiApiDeleteJobScheduleRequest struct {
	ctx _context.Context
	ApiService *JobSchedulerApiService
	cUUID string
	jUUID string
}


func (r JobSchedulerApiApiDeleteJobScheduleRequest) Execute() (JobSchedule, *_nethttp.Response, error) {
	return r.ApiService.DeleteJobScheduleExecute(r)
}

/*
 * DeleteJobSchedule Delete Job Schedule
 * WARNING: This is a preview API that could change.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID Customer UUID
 * @param jUUID Job Schedule UUID
 * @return JobSchedulerApiApiDeleteJobScheduleRequest
 */
func (a *JobSchedulerApiService) DeleteJobSchedule(ctx _context.Context, cUUID string, jUUID string) JobSchedulerApiApiDeleteJobScheduleRequest {
	return JobSchedulerApiApiDeleteJobScheduleRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
		jUUID: jUUID,
	}
}

/*
 * Execute executes the request
 * @return JobSchedule
 */
func (a *JobSchedulerApiService) DeleteJobScheduleExecute(r JobSchedulerApiApiDeleteJobScheduleRequest) (JobSchedule, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  JobSchedule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobSchedulerApiService.DeleteJobSchedule")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customers/{cUUID}/job-schedules/{jUUID}"
	localVarPath = strings.Replace(localVarPath, "{"+"cUUID"+"}", _neturl.PathEscape(parameterToString(r.cUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"jUUID"+"}", _neturl.PathEscape(parameterToString(r.jUUID, "")), -1)

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

type JobSchedulerApiApiGetJobScheduleRequest struct {
	ctx _context.Context
	ApiService *JobSchedulerApiService
	cUUID string
	jUUID string
}


func (r JobSchedulerApiApiGetJobScheduleRequest) Execute() (JobSchedule, *_nethttp.Response, error) {
	return r.ApiService.GetJobScheduleExecute(r)
}

/*
 * GetJobSchedule Get Job Schedule
 * WARNING: This is a preview API that could change.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID Customer UUID
 * @param jUUID Job Schedule UUID
 * @return JobSchedulerApiApiGetJobScheduleRequest
 */
func (a *JobSchedulerApiService) GetJobSchedule(ctx _context.Context, cUUID string, jUUID string) JobSchedulerApiApiGetJobScheduleRequest {
	return JobSchedulerApiApiGetJobScheduleRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
		jUUID: jUUID,
	}
}

/*
 * Execute executes the request
 * @return JobSchedule
 */
func (a *JobSchedulerApiService) GetJobScheduleExecute(r JobSchedulerApiApiGetJobScheduleRequest) (JobSchedule, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  JobSchedule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobSchedulerApiService.GetJobSchedule")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customers/{cUUID}/job-schedules/{jUUID}"
	localVarPath = strings.Replace(localVarPath, "{"+"cUUID"+"}", _neturl.PathEscape(parameterToString(r.cUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"jUUID"+"}", _neturl.PathEscape(parameterToString(r.jUUID, "")), -1)

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

type JobSchedulerApiApiPageListJobInstancesRequest struct {
	ctx _context.Context
	ApiService *JobSchedulerApiService
	cUUID string
	jUUID string
	jobInstancePagedQuerySpec *JobInstancePagedQuerySpec
}

func (r JobSchedulerApiApiPageListJobInstancesRequest) JobInstancePagedQuerySpec(jobInstancePagedQuerySpec JobInstancePagedQuerySpec) JobSchedulerApiApiPageListJobInstancesRequest {
	r.jobInstancePagedQuerySpec = &jobInstancePagedQuerySpec
	return r
}

func (r JobSchedulerApiApiPageListJobInstancesRequest) Execute() (JobInstancePagedResp, *_nethttp.Response, error) {
	return r.ApiService.PageListJobInstancesExecute(r)
}

/*
 * PageListJobInstances List Job Instances (paginated)
 * WARNING: This is a preview API that could change.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID Customer UUID
 * @param jUUID Job Schedule UUID
 * @return JobSchedulerApiApiPageListJobInstancesRequest
 */
func (a *JobSchedulerApiService) PageListJobInstances(ctx _context.Context, cUUID string, jUUID string) JobSchedulerApiApiPageListJobInstancesRequest {
	return JobSchedulerApiApiPageListJobInstancesRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
		jUUID: jUUID,
	}
}

/*
 * Execute executes the request
 * @return JobInstancePagedResp
 */
func (a *JobSchedulerApiService) PageListJobInstancesExecute(r JobSchedulerApiApiPageListJobInstancesRequest) (JobInstancePagedResp, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  JobInstancePagedResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobSchedulerApiService.PageListJobInstances")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customers/{cUUID}/job-schedules/{jUUID}/job-instances/page"
	localVarPath = strings.Replace(localVarPath, "{"+"cUUID"+"}", _neturl.PathEscape(parameterToString(r.cUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"jUUID"+"}", _neturl.PathEscape(parameterToString(r.jUUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.jobInstancePagedQuerySpec == nil {
		return localVarReturnValue, nil, reportError("jobInstancePagedQuerySpec is required and must be specified")
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
	localVarPostBody = r.jobInstancePagedQuerySpec
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

type JobSchedulerApiApiPageListJobSchedulesRequest struct {
	ctx _context.Context
	ApiService *JobSchedulerApiService
	cUUID string
	jobSchedulePagedQuerySpec *JobSchedulePagedQuerySpec
}

func (r JobSchedulerApiApiPageListJobSchedulesRequest) JobSchedulePagedQuerySpec(jobSchedulePagedQuerySpec JobSchedulePagedQuerySpec) JobSchedulerApiApiPageListJobSchedulesRequest {
	r.jobSchedulePagedQuerySpec = &jobSchedulePagedQuerySpec
	return r
}

func (r JobSchedulerApiApiPageListJobSchedulesRequest) Execute() (JobSchedulePagedResp, *_nethttp.Response, error) {
	return r.ApiService.PageListJobSchedulesExecute(r)
}

/*
 * PageListJobSchedules List Job Schedules (paginated)
 * WARNING: This is a preview API that could change.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID Customer UUID
 * @return JobSchedulerApiApiPageListJobSchedulesRequest
 */
func (a *JobSchedulerApiService) PageListJobSchedules(ctx _context.Context, cUUID string) JobSchedulerApiApiPageListJobSchedulesRequest {
	return JobSchedulerApiApiPageListJobSchedulesRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
	}
}

/*
 * Execute executes the request
 * @return JobSchedulePagedResp
 */
func (a *JobSchedulerApiService) PageListJobSchedulesExecute(r JobSchedulerApiApiPageListJobSchedulesRequest) (JobSchedulePagedResp, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  JobSchedulePagedResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobSchedulerApiService.PageListJobSchedules")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customers/{cUUID}/job-schedules/page"
	localVarPath = strings.Replace(localVarPath, "{"+"cUUID"+"}", _neturl.PathEscape(parameterToString(r.cUUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.jobSchedulePagedQuerySpec == nil {
		return localVarReturnValue, nil, reportError("jobSchedulePagedQuerySpec is required and must be specified")
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
	localVarPostBody = r.jobSchedulePagedQuerySpec
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

type JobSchedulerApiApiSnoozeJobScheduleRequest struct {
	ctx _context.Context
	ApiService *JobSchedulerApiService
	cUUID string
	jUUID string
	jobScheduleSnoozeSpec *JobScheduleSnoozeSpec
}

func (r JobSchedulerApiApiSnoozeJobScheduleRequest) JobScheduleSnoozeSpec(jobScheduleSnoozeSpec JobScheduleSnoozeSpec) JobSchedulerApiApiSnoozeJobScheduleRequest {
	r.jobScheduleSnoozeSpec = &jobScheduleSnoozeSpec
	return r
}

func (r JobSchedulerApiApiSnoozeJobScheduleRequest) Execute() (JobSchedule, *_nethttp.Response, error) {
	return r.ApiService.SnoozeJobScheduleExecute(r)
}

/*
 * SnoozeJobSchedule Snooze Job Schedule
 * WARNING: This is a preview API that could change.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID Customer UUID
 * @param jUUID Job Schedule UUID
 * @return JobSchedulerApiApiSnoozeJobScheduleRequest
 */
func (a *JobSchedulerApiService) SnoozeJobSchedule(ctx _context.Context, cUUID string, jUUID string) JobSchedulerApiApiSnoozeJobScheduleRequest {
	return JobSchedulerApiApiSnoozeJobScheduleRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
		jUUID: jUUID,
	}
}

/*
 * Execute executes the request
 * @return JobSchedule
 */
func (a *JobSchedulerApiService) SnoozeJobScheduleExecute(r JobSchedulerApiApiSnoozeJobScheduleRequest) (JobSchedule, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  JobSchedule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobSchedulerApiService.SnoozeJobSchedule")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customers/{cUUID}/job-schedules/{jUUID}/snooze"
	localVarPath = strings.Replace(localVarPath, "{"+"cUUID"+"}", _neturl.PathEscape(parameterToString(r.cUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"jUUID"+"}", _neturl.PathEscape(parameterToString(r.jUUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.jobScheduleSnoozeSpec == nil {
		return localVarReturnValue, nil, reportError("jobScheduleSnoozeSpec is required and must be specified")
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
	localVarPostBody = r.jobScheduleSnoozeSpec
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

type JobSchedulerApiApiUpdateJobScheduleRequest struct {
	ctx _context.Context
	ApiService *JobSchedulerApiService
	cUUID string
	jUUID string
	jobScheduleUpdateSpec *JobScheduleUpdateSpec
}

func (r JobSchedulerApiApiUpdateJobScheduleRequest) JobScheduleUpdateSpec(jobScheduleUpdateSpec JobScheduleUpdateSpec) JobSchedulerApiApiUpdateJobScheduleRequest {
	r.jobScheduleUpdateSpec = &jobScheduleUpdateSpec
	return r
}

func (r JobSchedulerApiApiUpdateJobScheduleRequest) Execute() (JobSchedule, *_nethttp.Response, error) {
	return r.ApiService.UpdateJobScheduleExecute(r)
}

/*
 * UpdateJobSchedule Update Job Schedule
 * WARNING: This is a preview API that could change.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cUUID Customer UUID
 * @param jUUID Job Schedule UUID
 * @return JobSchedulerApiApiUpdateJobScheduleRequest
 */
func (a *JobSchedulerApiService) UpdateJobSchedule(ctx _context.Context, cUUID string, jUUID string) JobSchedulerApiApiUpdateJobScheduleRequest {
	return JobSchedulerApiApiUpdateJobScheduleRequest{
		ApiService: a,
		ctx: ctx,
		cUUID: cUUID,
		jUUID: jUUID,
	}
}

/*
 * Execute executes the request
 * @return JobSchedule
 */
func (a *JobSchedulerApiService) UpdateJobScheduleExecute(r JobSchedulerApiApiUpdateJobScheduleRequest) (JobSchedule, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  JobSchedule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobSchedulerApiService.UpdateJobSchedule")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customers/{cUUID}/job-schedules/{jUUID}"
	localVarPath = strings.Replace(localVarPath, "{"+"cUUID"+"}", _neturl.PathEscape(parameterToString(r.cUUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"jUUID"+"}", _neturl.PathEscape(parameterToString(r.jUUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.jobScheduleUpdateSpec == nil {
		return localVarReturnValue, nil, reportError("jobScheduleUpdateSpec is required and must be specified")
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
	localVarPostBody = r.jobScheduleUpdateSpec
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
