// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservices

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

// ReplicationUsagesClient contains the methods for the ReplicationUsages group.
// Don't use this type directly, use NewReplicationUsagesClient() instead.
type ReplicationUsagesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewReplicationUsagesClient creates a new instance of ReplicationUsagesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewReplicationUsagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ReplicationUsagesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ReplicationUsagesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Fetches the replication usages of the vault.
//
// Generated from API version 2025-02-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vaultName - The name of the recovery services vault.
//   - options - ReplicationUsagesClientListOptions contains the optional parameters for the ReplicationUsagesClient.NewListPager
//     method.
func (client *ReplicationUsagesClient) NewListPager(resourceGroupName string, vaultName string, options *ReplicationUsagesClientListOptions) *runtime.Pager[ReplicationUsagesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReplicationUsagesClientListResponse]{
		More: func(page ReplicationUsagesClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ReplicationUsagesClientListResponse) (ReplicationUsagesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ReplicationUsagesClient.NewListPager")
			req, err := client.listCreateRequest(ctx, resourceGroupName, vaultName, options)
			if err != nil {
				return ReplicationUsagesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ReplicationUsagesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReplicationUsagesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ReplicationUsagesClient) listCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, _ *ReplicationUsagesClientListOptions) (*policy.Request, error) {
	urlPath := "/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/replicationUsages"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ReplicationUsagesClient) listHandleResponse(resp *http.Response) (ReplicationUsagesClientListResponse, error) {
	result := ReplicationUsagesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReplicationUsageList); err != nil {
		return ReplicationUsagesClientListResponse{}, err
	}
	return result, nil
}
