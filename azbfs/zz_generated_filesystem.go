package azbfs

// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
        "net/url"
    "github.com/Azure/azure-pipeline-go/pipeline"
        "net/url"
    "net/http"
        "net/url"
    "context"
        "net/url"
    "strconv"
        "net/url"
    "encoding/json"
        "net/url"
    "io/ioutil"
        "net/url"
    "io"
)

// filesystemClient is the azure Data Lake Storage provides storage for Hadoop and other big data workloads.
type filesystemClient struct {
    managementClient
}
// newFilesystemClient creates an instance of the filesystemClient client.
func newFilesystemClient(url url.URL, p pipeline.Pipeline) filesystemClient {
    return filesystemClient{newManagementClient(url, p)}
}

// Create create a filesystem rooted at the specified location. If the filesystem already exists, the operation fails.
// This operation does not support conditional HTTP requests.
//
// filesystem is the filesystem identifier.  The value must start and end with a letter or number and must contain only
// letters, numbers, and the dash (-) character.  Consecutive dashes are not permitted.  All letters must be lowercase.
// The value must have between 3 and 63 characters. xMsProperties is user-defined properties to be stored with the
// filesystem, in the format of a comma-separated list of name and value pairs "n1=v1, n2=v2, ...", where each value is
// a base64 encoded string. Note that the string may only contain ASCII characters in the ISO-8859-1 character set.
// xMsClientRequestID is a UUID recorded in the analytics logs for troubleshooting and correlation. timeout is an
// optional operation timeout value in seconds. The period begins when the request is received by the service. If the
// timeout value elapses before the operation completes, the operation fails. xMsDate is specifies the Coordinated
// Universal Time (UTC) for the request.  This is required when using shared key authorization.
func (client filesystemClient) Create(ctx context.Context, filesystem string, xMsProperties *string, xMsClientRequestID *string, timeout *int32, xMsDate *string) (*FilesystemCreateResponse, error) {
    if err := validate([]validation{
    { targetValue: filesystem,
     constraints: []constraint{	{target: "filesystem", name: maxLength, rule: 63, chain: nil },
    	{target: "filesystem", name: minLength, rule: 3, chain: nil },
    	{target: "filesystem", name: pattern, rule: `^[$a-z0-9][-a-z0-9]{1,61}[a-z0-9]$`, chain: nil }}},
    { targetValue: xMsClientRequestID,
     constraints: []constraint{	{target: "xMsClientRequestID", name: null, rule: false ,
    chain: []constraint{	{target: "xMsClientRequestID", name: pattern, rule: `^[{(]?[0-9a-f]{8}[-]?([0-9a-f]{4}[-]?){3}[0-9a-f]{12}[)}]?$`, chain: nil },
    }}}},
    { targetValue: timeout,
     constraints: []constraint{	{target: "timeout", name: null, rule: false ,
    chain: []constraint{	{target: "timeout", name: inclusiveMinimum, rule: 1, chain: nil },
    }}}}}); err != nil {
        return nil, err
    }
	req, err := client.createPreparer(filesystem, xMsProperties, xMsClientRequestID, timeout, xMsDate)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.createResponder}, req)
    if err != nil {
        return nil, err
    }
	return resp.(*FilesystemCreateResponse), err
}

// createPreparer prepares the Create request.
func (client filesystemClient) createPreparer(filesystem string, xMsProperties *string, xMsClientRequestID *string, timeout *int32, xMsDate *string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
    params := req.URL.Query()
    params.Set("resource", "filesystem")
    if timeout != nil {
        params.Set("timeout", strconv.FormatInt(int64(*timeout), 10))
    }
        req.URL.RawQuery = params.Encode()
    if xMsProperties != nil {
        req.Header.Set("x-ms-properties", *xMsProperties)
    }
    if xMsClientRequestID != nil {
        req.Header.Set("x-ms-client-request-id", *xMsClientRequestID)
    }
    if xMsDate != nil {
        req.Header.Set("x-ms-date", *xMsDate)
    }
    req.Header.Set("x-ms-version", ServiceVersion)
	return req, nil
}

// createResponder handles the response to the Create request.
func (client filesystemClient) createResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK,http.StatusCreated)
	if resp == nil {
		return nil, err
	}
    io.Copy(ioutil.Discard, resp.Response().Body)
    resp.Response().Body.Close()
        return &FilesystemCreateResponse{rawResponse: resp.Response()}, err
}

// Delete marks the filesystem for deletion.  When a filesystem is deleted, a filesystem with the same identifier
// cannot be created for at least 30 seconds. While the filesystem is being deleted, attempts to create a filesystem
// with the same identifier will fail with status code 409 (Conflict), with the service returning additional error
// information indicating that the filesystem is being deleted. All other operations, including operations on any files
// or directories within the filesystem, will fail with status code 404 (Not Found) while the filesystem is being
// deleted. This operation supports conditional HTTP requests.  For more information, see [Specifying Conditional
// Headers for Blob Service
// Operations](https://docs.microsoft.com/en-us/rest/api/storageservices/specifying-conditional-headers-for-blob-service-operations).
//
// filesystem is the filesystem identifier.  The value must start and end with a letter or number and must contain only
// letters, numbers, and the dash (-) character.  Consecutive dashes are not permitted.  All letters must be lowercase.
// The value must have between 3 and 63 characters. ifModifiedSince is optional. A date and time value. Specify this
// header to perform the operation only if the resource has been modified since the specified date and time.
// ifUnmodifiedSince is optional. A date and time value. Specify this header to perform the operation only if the
// resource has not been modified since the specified date and time. xMsClientRequestID is a UUID recorded in the
// analytics logs for troubleshooting and correlation. timeout is an optional operation timeout value in seconds. The
// period begins when the request is received by the service. If the timeout value elapses before the operation
// completes, the operation fails. xMsDate is specifies the Coordinated Universal Time (UTC) for the request.  This is
// required when using shared key authorization.
func (client filesystemClient) Delete(ctx context.Context, filesystem string, ifModifiedSince *string, ifUnmodifiedSince *string, xMsClientRequestID *string, timeout *int32, xMsDate *string) (*FilesystemDeleteResponse, error) {
    if err := validate([]validation{
    { targetValue: filesystem,
     constraints: []constraint{	{target: "filesystem", name: maxLength, rule: 63, chain: nil },
    	{target: "filesystem", name: minLength, rule: 3, chain: nil },
    	{target: "filesystem", name: pattern, rule: `^[$a-z0-9][-a-z0-9]{1,61}[a-z0-9]$`, chain: nil }}},
    { targetValue: xMsClientRequestID,
     constraints: []constraint{	{target: "xMsClientRequestID", name: null, rule: false ,
    chain: []constraint{	{target: "xMsClientRequestID", name: pattern, rule: `^[{(]?[0-9a-f]{8}[-]?([0-9a-f]{4}[-]?){3}[0-9a-f]{12}[)}]?$`, chain: nil },
    }}}},
    { targetValue: timeout,
     constraints: []constraint{	{target: "timeout", name: null, rule: false ,
    chain: []constraint{	{target: "timeout", name: inclusiveMinimum, rule: 1, chain: nil },
    }}}}}); err != nil {
        return nil, err
    }
	req, err := client.deletePreparer(filesystem, ifModifiedSince, ifUnmodifiedSince, xMsClientRequestID, timeout, xMsDate)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.deleteResponder}, req)
    if err != nil {
        return nil, err
    }
	return resp.(*FilesystemDeleteResponse), err
}

// deletePreparer prepares the Delete request.
func (client filesystemClient) deletePreparer(filesystem string, ifModifiedSince *string, ifUnmodifiedSince *string, xMsClientRequestID *string, timeout *int32, xMsDate *string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("DELETE", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
    params := req.URL.Query()
    params.Set("resource", "filesystem")
    if timeout != nil {
        params.Set("timeout", strconv.FormatInt(int64(*timeout), 10))
    }
        req.URL.RawQuery = params.Encode()
    if ifModifiedSince != nil {
        req.Header.Set("If-Modified-Since", *ifModifiedSince)
    }
    if ifUnmodifiedSince != nil {
        req.Header.Set("If-Unmodified-Since", *ifUnmodifiedSince)
    }
    if xMsClientRequestID != nil {
        req.Header.Set("x-ms-client-request-id", *xMsClientRequestID)
    }
    if xMsDate != nil {
        req.Header.Set("x-ms-date", *xMsDate)
    }
    req.Header.Set("x-ms-version", ServiceVersion)
	return req, nil
}

// deleteResponder handles the response to the Delete request.
func (client filesystemClient) deleteResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK,http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
    io.Copy(ioutil.Discard, resp.Response().Body)
    resp.Response().Body.Close()
        return &FilesystemDeleteResponse{rawResponse: resp.Response()}, err
}

// GetProperties all system and user-defined filesystem properties are specified in the response headers.
//
// filesystem is the filesystem identifier.  The value must start and end with a letter or number and must contain only
// letters, numbers, and the dash (-) character.  Consecutive dashes are not permitted.  All letters must be lowercase.
// The value must have between 3 and 63 characters. xMsClientRequestID is a UUID recorded in the analytics logs for
// troubleshooting and correlation. timeout is an optional operation timeout value in seconds. The period begins when
// the request is received by the service. If the timeout value elapses before the operation completes, the operation
// fails. xMsDate is specifies the Coordinated Universal Time (UTC) for the request.  This is required when using
// shared key authorization.
func (client filesystemClient) GetProperties(ctx context.Context, filesystem string, xMsClientRequestID *string, timeout *int32, xMsDate *string) (*FilesystemGetPropertiesResponse, error) {
    if err := validate([]validation{
    { targetValue: filesystem,
     constraints: []constraint{	{target: "filesystem", name: maxLength, rule: 63, chain: nil },
    	{target: "filesystem", name: minLength, rule: 3, chain: nil },
    	{target: "filesystem", name: pattern, rule: `^[$a-z0-9][-a-z0-9]{1,61}[a-z0-9]$`, chain: nil }}},
    { targetValue: xMsClientRequestID,
     constraints: []constraint{	{target: "xMsClientRequestID", name: null, rule: false ,
    chain: []constraint{	{target: "xMsClientRequestID", name: pattern, rule: `^[{(]?[0-9a-f]{8}[-]?([0-9a-f]{4}[-]?){3}[0-9a-f]{12}[)}]?$`, chain: nil },
    }}}},
    { targetValue: timeout,
     constraints: []constraint{	{target: "timeout", name: null, rule: false ,
    chain: []constraint{	{target: "timeout", name: inclusiveMinimum, rule: 1, chain: nil },
    }}}}}); err != nil {
        return nil, err
    }
	req, err := client.getPropertiesPreparer(filesystem, xMsClientRequestID, timeout, xMsDate)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getPropertiesResponder}, req)
    if err != nil {
        return nil, err
    }
	return resp.(*FilesystemGetPropertiesResponse), err
}

// getPropertiesPreparer prepares the GetProperties request.
func (client filesystemClient) getPropertiesPreparer(filesystem string, xMsClientRequestID *string, timeout *int32, xMsDate *string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("HEAD", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
    params := req.URL.Query()
    params.Set("resource", "filesystem")
    if timeout != nil {
        params.Set("timeout", strconv.FormatInt(int64(*timeout), 10))
    }
        req.URL.RawQuery = params.Encode()
    if xMsClientRequestID != nil {
        req.Header.Set("x-ms-client-request-id", *xMsClientRequestID)
    }
    if xMsDate != nil {
        req.Header.Set("x-ms-date", *xMsDate)
    }
    req.Header.Set("x-ms-version", ServiceVersion)
	return req, nil
}

// getPropertiesResponder handles the response to the GetProperties request.
func (client filesystemClient) getPropertiesResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
    io.Copy(ioutil.Discard, resp.Response().Body)
    resp.Response().Body.Close()
        return &FilesystemGetPropertiesResponse{rawResponse: resp.Response()}, err
}

// List list filesystems and their properties in given account.
//
// prefix is filters results to filesystems within the specified prefix. continuation is the number of filesystems
// returned with each invocation is limited. If the number of filesystems to be returned exceeds this limit, a
// continuation token is returned in the response header x-ms-continuation. When a continuation token is  returned in
// the response, it must be specified in a subsequent invocation of the list operation to continue listing the
// filesystems. maxResults is an optional value that specifies the maximum number of items to return. If omitted or
// greater than 5,000, the response will include up to 5,000 items. xMsClientRequestID is a UUID recorded in the
// analytics logs for troubleshooting and correlation. timeout is an optional operation timeout value in seconds. The
// period begins when the request is received by the service. If the timeout value elapses before the operation
// completes, the operation fails. xMsDate is specifies the Coordinated Universal Time (UTC) for the request.  This is
// required when using shared key authorization.
func (client filesystemClient) List(ctx context.Context, prefix *string, continuation *string, maxResults *int32, xMsClientRequestID *string, timeout *int32, xMsDate *string) (*FilesystemList, error) {
    if err := validate([]validation{
    { targetValue: maxResults,
     constraints: []constraint{	{target: "maxResults", name: null, rule: false ,
    chain: []constraint{	{target: "maxResults", name: inclusiveMinimum, rule: 1, chain: nil },
    }}}},
    { targetValue: xMsClientRequestID,
     constraints: []constraint{	{target: "xMsClientRequestID", name: null, rule: false ,
    chain: []constraint{	{target: "xMsClientRequestID", name: pattern, rule: `^[{(]?[0-9a-f]{8}[-]?([0-9a-f]{4}[-]?){3}[0-9a-f]{12}[)}]?$`, chain: nil },
    }}}},
    { targetValue: timeout,
     constraints: []constraint{	{target: "timeout", name: null, rule: false ,
    chain: []constraint{	{target: "timeout", name: inclusiveMinimum, rule: 1, chain: nil },
    }}}}}); err != nil {
        return nil, err
    }
	req, err := client.listPreparer(prefix, continuation, maxResults, xMsClientRequestID, timeout, xMsDate)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listResponder}, req)
    if err != nil {
        return nil, err
    }
	return resp.(*FilesystemList), err
}

// listPreparer prepares the List request.
func (client filesystemClient) listPreparer(prefix *string, continuation *string, maxResults *int32, xMsClientRequestID *string, timeout *int32, xMsDate *string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("GET", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
    params := req.URL.Query()
    params.Set("resource", "account")
    if prefix != nil && len(*prefix) > 0 {
        params.Set("prefix", *prefix)
    }
    if continuation != nil && len(*continuation) > 0 {
        params.Set("continuation", *continuation)
    }
    if maxResults != nil {
        params.Set("maxResults", strconv.FormatInt(int64(*maxResults), 10))
    }
    if timeout != nil {
        params.Set("timeout", strconv.FormatInt(int64(*timeout), 10))
    }
        req.URL.RawQuery = params.Encode()
    if xMsClientRequestID != nil {
        req.Header.Set("x-ms-client-request-id", *xMsClientRequestID)
    }
    if xMsDate != nil {
        req.Header.Set("x-ms-date", *xMsDate)
    }
    req.Header.Set("x-ms-version", ServiceVersion)
	return req, nil
}

// listResponder handles the response to the List request.
func (client filesystemClient) listResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
    result:= &FilesystemList{rawResponse: resp.Response()}
    if err != nil {
        return result, err
    }
    defer resp.Response().Body.Close()
    b, err:= ioutil.ReadAll(resp.Response().Body)
    if err != nil {
        return result, err
    }
    if len(b) > 0 {
        b = removeBOM(b)
        err = json.Unmarshal(b, result)
        if err != nil {
            return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
        }
    }
    return result, nil
}

// SetProperties set properties for the filesystem.  This operation supports conditional HTTP requests.  For more
// information, see [Specifying Conditional Headers for Blob Service
// Operations](https://docs.microsoft.com/en-us/rest/api/storageservices/specifying-conditional-headers-for-blob-service-operations).
//
// filesystem is the filesystem identifier.  The value must start and end with a letter or number and must contain only
// letters, numbers, and the dash (-) character.  Consecutive dashes are not permitted.  All letters must be lowercase.
// The value must have between 3 and 63 characters. xMsProperties is optional. User-defined properties to be stored
// with the filesystem, in the format of a comma-separated list of name and value pairs "n1=v1, n2=v2, ...", where each
// value is a base64 encoded string. Note that the string may only contain ASCII characters in the ISO-8859-1 character
// set.  If the filesystem exists, any properties not included in the list will be removed.  All properties are removed
// if the header is omitted.  To merge new and existing properties, first get all existing properties and the current
// E-Tag, then make a conditional request with the E-Tag and include values for all properties. ifModifiedSince is
// optional. A date and time value. Specify this header to perform the operation only if the resource has been modified
// since the specified date and time. ifUnmodifiedSince is optional. A date and time value. Specify this header to
// perform the operation only if the resource has not been modified since the specified date and time.
// xMsClientRequestID is a UUID recorded in the analytics logs for troubleshooting and correlation. timeout is an
// optional operation timeout value in seconds. The period begins when the request is received by the service. If the
// timeout value elapses before the operation completes, the operation fails. xMsDate is specifies the Coordinated
// Universal Time (UTC) for the request.  This is required when using shared key authorization.
func (client filesystemClient) SetProperties(ctx context.Context, filesystem string, xMsProperties *string, ifModifiedSince *string, ifUnmodifiedSince *string, xMsClientRequestID *string, timeout *int32, xMsDate *string) (*FilesystemSetPropertiesResponse, error) {
    if err := validate([]validation{
    { targetValue: filesystem,
     constraints: []constraint{	{target: "filesystem", name: maxLength, rule: 63, chain: nil },
    	{target: "filesystem", name: minLength, rule: 3, chain: nil },
    	{target: "filesystem", name: pattern, rule: `^[$a-z0-9][-a-z0-9]{1,61}[a-z0-9]$`, chain: nil }}},
    { targetValue: xMsClientRequestID,
     constraints: []constraint{	{target: "xMsClientRequestID", name: null, rule: false ,
    chain: []constraint{	{target: "xMsClientRequestID", name: pattern, rule: `^[{(]?[0-9a-f]{8}[-]?([0-9a-f]{4}[-]?){3}[0-9a-f]{12}[)}]?$`, chain: nil },
    }}}},
    { targetValue: timeout,
     constraints: []constraint{	{target: "timeout", name: null, rule: false ,
    chain: []constraint{	{target: "timeout", name: inclusiveMinimum, rule: 1, chain: nil },
    }}}}}); err != nil {
        return nil, err
    }
	req, err := client.setPropertiesPreparer(filesystem, xMsProperties, ifModifiedSince, ifUnmodifiedSince, xMsClientRequestID, timeout, xMsDate)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.setPropertiesResponder}, req)
    if err != nil {
        return nil, err
    }
	return resp.(*FilesystemSetPropertiesResponse), err
}

// setPropertiesPreparer prepares the SetProperties request.
func (client filesystemClient) setPropertiesPreparer(filesystem string, xMsProperties *string, ifModifiedSince *string, ifUnmodifiedSince *string, xMsClientRequestID *string, timeout *int32, xMsDate *string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PATCH", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
    params := req.URL.Query()
    params.Set("resource", "filesystem")
    if timeout != nil {
        params.Set("timeout", strconv.FormatInt(int64(*timeout), 10))
    }
        req.URL.RawQuery = params.Encode()
    if xMsProperties != nil {
        req.Header.Set("x-ms-properties", *xMsProperties)
    }
    if ifModifiedSince != nil {
        req.Header.Set("If-Modified-Since", *ifModifiedSince)
    }
    if ifUnmodifiedSince != nil {
        req.Header.Set("If-Unmodified-Since", *ifUnmodifiedSince)
    }
    if xMsClientRequestID != nil {
        req.Header.Set("x-ms-client-request-id", *xMsClientRequestID)
    }
    if xMsDate != nil {
        req.Header.Set("x-ms-date", *xMsDate)
    }
    req.Header.Set("x-ms-version", ServiceVersion)
	return req, nil
}

// setPropertiesResponder handles the response to the SetProperties request.
func (client filesystemClient) setPropertiesResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
    io.Copy(ioutil.Discard, resp.Response().Body)
    resp.Response().Body.Close()
        return &FilesystemSetPropertiesResponse{rawResponse: resp.Response()}, err
}

