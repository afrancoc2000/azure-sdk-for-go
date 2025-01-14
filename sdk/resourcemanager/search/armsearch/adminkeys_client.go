//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsearch

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

// AdminKeysClient contains the methods for the AdminKeys group.
// Don't use this type directly, use NewAdminKeysClient() instead.
type AdminKeysClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAdminKeysClient creates a new instance of AdminKeysClient with the specified values.
//   - subscriptionID - The unique identifier for a Microsoft Azure subscription. You can obtain this value from the Azure Resource
//     Manager API or the portal.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAdminKeysClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AdminKeysClient, error) {
	cl, err := arm.NewClient(moduleName+".AdminKeysClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AdminKeysClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets the primary and secondary admin API keys for the specified Azure Cognitive Search service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
//   - resourceGroupName - The name of the resource group within the current subscription. You can obtain this value from the
//     Azure Resource Manager API or the portal.
//   - searchServiceName - The name of the Azure Cognitive Search service associated with the specified resource group.
//   - SearchManagementRequestOptions - SearchManagementRequestOptions contains a group of parameters for the AdminKeysClient.Get
//     method.
//   - options - AdminKeysClientGetOptions contains the optional parameters for the AdminKeysClient.Get method.
func (client *AdminKeysClient) Get(ctx context.Context, resourceGroupName string, searchServiceName string, searchManagementRequestOptions *SearchManagementRequestOptions, options *AdminKeysClientGetOptions) (AdminKeysClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, searchServiceName, searchManagementRequestOptions, options)
	if err != nil {
		return AdminKeysClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AdminKeysClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AdminKeysClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AdminKeysClient) getCreateRequest(ctx context.Context, resourceGroupName string, searchServiceName string, searchManagementRequestOptions *SearchManagementRequestOptions, options *AdminKeysClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/listAdminKeys"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if searchServiceName == "" {
		return nil, errors.New("parameter searchServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{searchServiceName}", url.PathEscape(searchServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if searchManagementRequestOptions != nil && searchManagementRequestOptions.ClientRequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*searchManagementRequestOptions.ClientRequestID}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AdminKeysClient) getHandleResponse(resp *http.Response) (AdminKeysClientGetResponse, error) {
	result := AdminKeysClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AdminKeyResult); err != nil {
		return AdminKeysClientGetResponse{}, err
	}
	return result, nil
}

// Regenerate - Regenerates either the primary or secondary admin API key. You can only regenerate one key at a time.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
//   - resourceGroupName - The name of the resource group within the current subscription. You can obtain this value from the
//     Azure Resource Manager API or the portal.
//   - searchServiceName - The name of the Azure Cognitive Search service associated with the specified resource group.
//   - keyKind - Specifies which key to regenerate. Valid values include 'primary' and 'secondary'.
//   - SearchManagementRequestOptions - SearchManagementRequestOptions contains a group of parameters for the AdminKeysClient.Get
//     method.
//   - options - AdminKeysClientRegenerateOptions contains the optional parameters for the AdminKeysClient.Regenerate method.
func (client *AdminKeysClient) Regenerate(ctx context.Context, resourceGroupName string, searchServiceName string, keyKind AdminKeyKind, searchManagementRequestOptions *SearchManagementRequestOptions, options *AdminKeysClientRegenerateOptions) (AdminKeysClientRegenerateResponse, error) {
	var err error
	req, err := client.regenerateCreateRequest(ctx, resourceGroupName, searchServiceName, keyKind, searchManagementRequestOptions, options)
	if err != nil {
		return AdminKeysClientRegenerateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AdminKeysClientRegenerateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AdminKeysClientRegenerateResponse{}, err
	}
	resp, err := client.regenerateHandleResponse(httpResp)
	return resp, err
}

// regenerateCreateRequest creates the Regenerate request.
func (client *AdminKeysClient) regenerateCreateRequest(ctx context.Context, resourceGroupName string, searchServiceName string, keyKind AdminKeyKind, searchManagementRequestOptions *SearchManagementRequestOptions, options *AdminKeysClientRegenerateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/regenerateAdminKey/{keyKind}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if searchServiceName == "" {
		return nil, errors.New("parameter searchServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{searchServiceName}", url.PathEscape(searchServiceName))
	if keyKind == "" {
		return nil, errors.New("parameter keyKind cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{keyKind}", url.PathEscape(string(keyKind)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if searchManagementRequestOptions != nil && searchManagementRequestOptions.ClientRequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*searchManagementRequestOptions.ClientRequestID}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// regenerateHandleResponse handles the Regenerate response.
func (client *AdminKeysClient) regenerateHandleResponse(resp *http.Response) (AdminKeysClientRegenerateResponse, error) {
	result := AdminKeysClientRegenerateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AdminKeyResult); err != nil {
		return AdminKeysClientRegenerateResponse{}, err
	}
	return result, nil
}
