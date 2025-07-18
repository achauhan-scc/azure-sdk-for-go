// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armprivatelinkscopes

// ClientCreateOrUpdateResponse contains the response from method Client.CreateOrUpdate.
type ClientCreateOrUpdateResponse struct {
	// An Azure Arc PrivateLinkScope definition.
	KubernetesConfigurationPrivateLinkScope
}

// ClientDeleteResponse contains the response from method Client.Delete.
type ClientDeleteResponse struct {
	// placeholder for future response values
}

// ClientGetResponse contains the response from method Client.Get.
type ClientGetResponse struct {
	// An Azure Arc PrivateLinkScope definition.
	KubernetesConfigurationPrivateLinkScope
}

// ClientListByResourceGroupResponse contains the response from method Client.NewListByResourceGroupPager.
type ClientListByResourceGroupResponse struct {
	// Describes the list of Azure Arc PrivateLinkScope resources.
	KubernetesConfigurationPrivateLinkScopeListResult
}

// ClientListResponse contains the response from method Client.NewListPager.
type ClientListResponse struct {
	// Describes the list of Azure Arc PrivateLinkScope resources.
	KubernetesConfigurationPrivateLinkScopeListResult
}

// ClientUpdateTagsResponse contains the response from method Client.UpdateTags.
type ClientUpdateTagsResponse struct {
	// An Azure Arc PrivateLinkScope definition.
	KubernetesConfigurationPrivateLinkScope
}

// PrivateEndpointConnectionsClientCreateOrUpdateResponse contains the response from method PrivateEndpointConnectionsClient.BeginCreateOrUpdate.
type PrivateEndpointConnectionsClientCreateOrUpdateResponse struct {
	// The Private Endpoint Connection resource.
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.Delete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	// The Private Endpoint Connection resource.
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListByPrivateLinkScopeResponse contains the response from method PrivateEndpointConnectionsClient.ListByPrivateLinkScope.
type PrivateEndpointConnectionsClientListByPrivateLinkScopeResponse struct {
	// List of private endpoint connection associated with the specified storage account
	PrivateEndpointConnectionListResult
}

// PrivateLinkResourcesClientGetResponse contains the response from method PrivateLinkResourcesClient.Get.
type PrivateLinkResourcesClientGetResponse struct {
	// A private link resource
	PrivateLinkResource
}

// PrivateLinkResourcesClientListByPrivateLinkScopeResponse contains the response from method PrivateLinkResourcesClient.ListByPrivateLinkScope.
type PrivateLinkResourcesClientListByPrivateLinkScopeResponse struct {
	// A list of private link resources
	PrivateLinkResourceListResult
}
