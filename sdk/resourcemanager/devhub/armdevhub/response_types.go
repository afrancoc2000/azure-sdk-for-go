//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdevhub

// DeveloperHubServiceClientGitHubOAuthCallbackResponse contains the response from method DeveloperHubServiceClient.GitHubOAuthCallback.
type DeveloperHubServiceClientGitHubOAuthCallbackResponse struct {
	GitHubOAuthResponse
}

// DeveloperHubServiceClientGitHubOAuthResponse contains the response from method DeveloperHubServiceClient.GitHubOAuth.
type DeveloperHubServiceClientGitHubOAuthResponse struct {
	GitHubOAuthInfoResponse
}

// DeveloperHubServiceClientListGitHubOAuthResponse contains the response from method DeveloperHubServiceClient.ListGitHubOAuth.
type DeveloperHubServiceClientListGitHubOAuthResponse struct {
	GitHubOAuthListResponse
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}

// WorkflowClientCreateOrUpdateResponse contains the response from method WorkflowClient.CreateOrUpdate.
type WorkflowClientCreateOrUpdateResponse struct {
	Workflow
}

// WorkflowClientDeleteResponse contains the response from method WorkflowClient.Delete.
type WorkflowClientDeleteResponse struct {
	DeleteWorkflowResponse
}

// WorkflowClientGetResponse contains the response from method WorkflowClient.Get.
type WorkflowClientGetResponse struct {
	Workflow
}

// WorkflowClientListByResourceGroupResponse contains the response from method WorkflowClient.ListByResourceGroup.
type WorkflowClientListByResourceGroupResponse struct {
	WorkflowListResult
}

// WorkflowClientListResponse contains the response from method WorkflowClient.List.
type WorkflowClientListResponse struct {
	WorkflowListResult
}

// WorkflowClientUpdateTagsResponse contains the response from method WorkflowClient.UpdateTags.
type WorkflowClientUpdateTagsResponse struct {
	Workflow
}