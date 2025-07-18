// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridcompute

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

// GatewaysClient contains the methods for the Gateways group.
// Don't use this type directly, use NewGatewaysClient() instead.
type GatewaysClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewGatewaysClient creates a new instance of GatewaysClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewGatewaysClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*GatewaysClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &GatewaysClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - The operation to create or update a gateway.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-19-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - gatewayName - The name of the Gateway.
//   - parameters - Parameters supplied to the Create gateway operation.
//   - options - GatewaysClientBeginCreateOrUpdateOptions contains the optional parameters for the GatewaysClient.BeginCreateOrUpdate
//     method.
func (client *GatewaysClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, gatewayName string, parameters Gateway, options *GatewaysClientBeginCreateOrUpdateOptions) (*runtime.Poller[GatewaysClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, gatewayName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[GatewaysClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[GatewaysClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - The operation to create or update a gateway.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-19-preview
func (client *GatewaysClient) createOrUpdate(ctx context.Context, resourceGroupName string, gatewayName string, parameters Gateway, options *GatewaysClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "GatewaysClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, gatewayName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *GatewaysClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, parameters Gateway, _ *GatewaysClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/gateways/{gatewayName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-19-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - The operation to delete a gateway.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-19-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - gatewayName - The name of the Gateway.
//   - options - GatewaysClientBeginDeleteOptions contains the optional parameters for the GatewaysClient.BeginDelete method.
func (client *GatewaysClient) BeginDelete(ctx context.Context, resourceGroupName string, gatewayName string, options *GatewaysClientBeginDeleteOptions) (*runtime.Poller[GatewaysClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, gatewayName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[GatewaysClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[GatewaysClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - The operation to delete a gateway.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-19-preview
func (client *GatewaysClient) deleteOperation(ctx context.Context, resourceGroupName string, gatewayName string, options *GatewaysClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "GatewaysClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, gatewayName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *GatewaysClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, _ *GatewaysClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/gateways/{gatewayName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-19-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieves information about the view of a gateway.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-19-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - gatewayName - The name of the Gateway.
//   - options - GatewaysClientGetOptions contains the optional parameters for the GatewaysClient.Get method.
func (client *GatewaysClient) Get(ctx context.Context, resourceGroupName string, gatewayName string, options *GatewaysClientGetOptions) (GatewaysClientGetResponse, error) {
	var err error
	const operationName = "GatewaysClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, gatewayName, options)
	if err != nil {
		return GatewaysClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GatewaysClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return GatewaysClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *GatewaysClient) getCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, _ *GatewaysClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/gateways/{gatewayName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-19-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *GatewaysClient) getHandleResponse(resp *http.Response) (GatewaysClientGetResponse, error) {
	result := GatewaysClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Gateway); err != nil {
		return GatewaysClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - The operation to get all gateways of a non-Azure machine
//
// Generated from API version 2025-02-19-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - GatewaysClientListByResourceGroupOptions contains the optional parameters for the GatewaysClient.NewListByResourceGroupPager
//     method.
func (client *GatewaysClient) NewListByResourceGroupPager(resourceGroupName string, options *GatewaysClientListByResourceGroupOptions) *runtime.Pager[GatewaysClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[GatewaysClientListByResourceGroupResponse]{
		More: func(page GatewaysClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *GatewaysClientListByResourceGroupResponse) (GatewaysClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "GatewaysClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return GatewaysClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *GatewaysClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, _ *GatewaysClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/gateways"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-19-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *GatewaysClient) listByResourceGroupHandleResponse(resp *http.Response) (GatewaysClientListByResourceGroupResponse, error) {
	result := GatewaysClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GatewaysListResult); err != nil {
		return GatewaysClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - The operation to get all gateways of a non-Azure machine
//
// Generated from API version 2025-02-19-preview
//   - options - GatewaysClientListBySubscriptionOptions contains the optional parameters for the GatewaysClient.NewListBySubscriptionPager
//     method.
func (client *GatewaysClient) NewListBySubscriptionPager(options *GatewaysClientListBySubscriptionOptions) *runtime.Pager[GatewaysClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[GatewaysClientListBySubscriptionResponse]{
		More: func(page GatewaysClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *GatewaysClientListBySubscriptionResponse) (GatewaysClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "GatewaysClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return GatewaysClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *GatewaysClient) listBySubscriptionCreateRequest(ctx context.Context, _ *GatewaysClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HybridCompute/gateways"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-19-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *GatewaysClient) listBySubscriptionHandleResponse(resp *http.Response) (GatewaysClientListBySubscriptionResponse, error) {
	result := GatewaysClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GatewaysListResult); err != nil {
		return GatewaysClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - The operation to update a gateway.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-19-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - gatewayName - The name of the Gateway.
//   - parameters - Parameters supplied to the Update gateway operation.
//   - options - GatewaysClientUpdateOptions contains the optional parameters for the GatewaysClient.Update method.
func (client *GatewaysClient) Update(ctx context.Context, resourceGroupName string, gatewayName string, parameters GatewayUpdate, options *GatewaysClientUpdateOptions) (GatewaysClientUpdateResponse, error) {
	var err error
	const operationName = "GatewaysClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, gatewayName, parameters, options)
	if err != nil {
		return GatewaysClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GatewaysClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return GatewaysClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *GatewaysClient) updateCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, parameters GatewayUpdate, _ *GatewaysClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/gateways/{gatewayName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-19-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *GatewaysClient) updateHandleResponse(resp *http.Response) (GatewaysClientUpdateResponse, error) {
	result := GatewaysClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Gateway); err != nil {
		return GatewaysClientUpdateResponse{}, err
	}
	return result, nil
}
