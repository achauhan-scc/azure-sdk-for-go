//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.
// Code generated by @autorest/go. DO NOT EDIT.

package armelastic

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// OpenAIClient contains the methods for the OpenAI group.
// Don't use this type directly, use NewOpenAIClient() instead.
type OpenAIClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewOpenAIClient creates a new instance of OpenAIClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewOpenAIClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*OpenAIClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &OpenAIClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update a OpenAI integration rule for a given monitor resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - monitorName - Monitor resource name
//   - integrationName - OpenAI Integration name
//   - options - OpenAIClientCreateOrUpdateOptions contains the optional parameters for the OpenAIClient.CreateOrUpdate method.
func (client *OpenAIClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, monitorName string, integrationName string, options *OpenAIClientCreateOrUpdateOptions) (OpenAIClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "OpenAIClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, monitorName, integrationName, options)
	if err != nil {
		return OpenAIClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OpenAIClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return OpenAIClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *OpenAIClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, integrationName string, options *OpenAIClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Elastic/monitors/{monitorName}/openAIIntegrations/{integrationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	if integrationName == "" {
		return nil, errors.New("parameter integrationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationName}", url.PathEscape(integrationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Body != nil {
		if err := runtime.MarshalAsJSON(req, *options.Body); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *OpenAIClient) createOrUpdateHandleResponse(resp *http.Response) (OpenAIClientCreateOrUpdateResponse, error) {
	result := OpenAIClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OpenAIIntegrationRPModel); err != nil {
		return OpenAIClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete OpenAI integration rule for a given monitor resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - monitorName - Monitor resource name
//   - integrationName - OpenAI Integration name
//   - options - OpenAIClientDeleteOptions contains the optional parameters for the OpenAIClient.Delete method.
func (client *OpenAIClient) Delete(ctx context.Context, resourceGroupName string, monitorName string, integrationName string, options *OpenAIClientDeleteOptions) (OpenAIClientDeleteResponse, error) {
	var err error
	const operationName = "OpenAIClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, monitorName, integrationName, options)
	if err != nil {
		return OpenAIClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OpenAIClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OpenAIClientDeleteResponse{}, err
	}
	return OpenAIClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *OpenAIClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, integrationName string, options *OpenAIClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Elastic/monitors/{monitorName}/openAIIntegrations/{integrationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	if integrationName == "" {
		return nil, errors.New("parameter integrationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationName}", url.PathEscape(integrationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get OpenAI integration rule for a given monitor resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - monitorName - Monitor resource name
//   - integrationName - OpenAI Integration name
//   - options - OpenAIClientGetOptions contains the optional parameters for the OpenAIClient.Get method.
func (client *OpenAIClient) Get(ctx context.Context, resourceGroupName string, monitorName string, integrationName string, options *OpenAIClientGetOptions) (OpenAIClientGetResponse, error) {
	var err error
	const operationName = "OpenAIClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, monitorName, integrationName, options)
	if err != nil {
		return OpenAIClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OpenAIClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OpenAIClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *OpenAIClient) getCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, integrationName string, options *OpenAIClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Elastic/monitors/{monitorName}/openAIIntegrations/{integrationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	if integrationName == "" {
		return nil, errors.New("parameter integrationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationName}", url.PathEscape(integrationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *OpenAIClient) getHandleResponse(resp *http.Response) (OpenAIClientGetResponse, error) {
	result := OpenAIClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OpenAIIntegrationRPModel); err != nil {
		return OpenAIClientGetResponse{}, err
	}
	return result, nil
}

// GetStatus - Get OpenAI integration status for a given integration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - monitorName - Monitor resource name
//   - integrationName - OpenAI Integration name
//   - options - OpenAIClientGetStatusOptions contains the optional parameters for the OpenAIClient.GetStatus method.
func (client *OpenAIClient) GetStatus(ctx context.Context, resourceGroupName string, monitorName string, integrationName string, options *OpenAIClientGetStatusOptions) (OpenAIClientGetStatusResponse, error) {
	var err error
	const operationName = "OpenAIClient.GetStatus"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getStatusCreateRequest(ctx, resourceGroupName, monitorName, integrationName, options)
	if err != nil {
		return OpenAIClientGetStatusResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OpenAIClientGetStatusResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OpenAIClientGetStatusResponse{}, err
	}
	resp, err := client.getStatusHandleResponse(httpResp)
	return resp, err
}

// getStatusCreateRequest creates the GetStatus request.
func (client *OpenAIClient) getStatusCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, integrationName string, options *OpenAIClientGetStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Elastic/monitors/{monitorName}/openAIIntegrations/{integrationName}/getStatus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	if integrationName == "" {
		return nil, errors.New("parameter integrationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationName}", url.PathEscape(integrationName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getStatusHandleResponse handles the GetStatus response.
func (client *OpenAIClient) getStatusHandleResponse(resp *http.Response) (OpenAIClientGetStatusResponse, error) {
	result := OpenAIClientGetStatusResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OpenAIIntegrationStatusResponse); err != nil {
		return OpenAIClientGetStatusResponse{}, err
	}
	return result, nil
}

// NewListPager - List OpenAI integration rule for a given monitor resource.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - monitorName - Monitor resource name
//   - options - OpenAIClientListOptions contains the optional parameters for the OpenAIClient.NewListPager method.
func (client *OpenAIClient) NewListPager(resourceGroupName string, monitorName string, options *OpenAIClientListOptions) *runtime.Pager[OpenAIClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[OpenAIClientListResponse]{
		More: func(page OpenAIClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OpenAIClientListResponse) (OpenAIClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "OpenAIClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, monitorName, options)
			}, nil)
			if err != nil {
				return OpenAIClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *OpenAIClient) listCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, options *OpenAIClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Elastic/monitors/{monitorName}/openAIIntegrations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *OpenAIClient) listHandleResponse(resp *http.Response) (OpenAIClientListResponse, error) {
	result := OpenAIClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OpenAIIntegrationRPModelListResponse); err != nil {
		return OpenAIClientListResponse{}, err
	}
	return result, nil
}
