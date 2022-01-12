//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourcemover

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"net/http"
	"time"
)

// MoveCollectionsBulkRemovePollerResponse contains the response from method MoveCollections.BulkRemove.
type MoveCollectionsBulkRemovePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *MoveCollectionsBulkRemovePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l MoveCollectionsBulkRemovePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (MoveCollectionsBulkRemoveResponse, error) {
	respType := MoveCollectionsBulkRemoveResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.OperationStatus)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a MoveCollectionsBulkRemovePollerResponse from the provided client and resume token.
func (l *MoveCollectionsBulkRemovePollerResponse) Resume(ctx context.Context, client *MoveCollectionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("MoveCollectionsClient.BulkRemove", token, client.pl, client.bulkRemoveHandleError)
	if err != nil {
		return err
	}
	poller := &MoveCollectionsBulkRemovePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// MoveCollectionsBulkRemoveResponse contains the response from method MoveCollections.BulkRemove.
type MoveCollectionsBulkRemoveResponse struct {
	MoveCollectionsBulkRemoveResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MoveCollectionsBulkRemoveResult contains the result from method MoveCollections.BulkRemove.
type MoveCollectionsBulkRemoveResult struct {
	OperationStatus
}

// MoveCollectionsCommitPollerResponse contains the response from method MoveCollections.Commit.
type MoveCollectionsCommitPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *MoveCollectionsCommitPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l MoveCollectionsCommitPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (MoveCollectionsCommitResponse, error) {
	respType := MoveCollectionsCommitResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.OperationStatus)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a MoveCollectionsCommitPollerResponse from the provided client and resume token.
func (l *MoveCollectionsCommitPollerResponse) Resume(ctx context.Context, client *MoveCollectionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("MoveCollectionsClient.Commit", token, client.pl, client.commitHandleError)
	if err != nil {
		return err
	}
	poller := &MoveCollectionsCommitPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// MoveCollectionsCommitResponse contains the response from method MoveCollections.Commit.
type MoveCollectionsCommitResponse struct {
	MoveCollectionsCommitResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MoveCollectionsCommitResult contains the result from method MoveCollections.Commit.
type MoveCollectionsCommitResult struct {
	OperationStatus
}

// MoveCollectionsCreateResponse contains the response from method MoveCollections.Create.
type MoveCollectionsCreateResponse struct {
	MoveCollectionsCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MoveCollectionsCreateResult contains the result from method MoveCollections.Create.
type MoveCollectionsCreateResult struct {
	MoveCollection
}

// MoveCollectionsDeletePollerResponse contains the response from method MoveCollections.Delete.
type MoveCollectionsDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *MoveCollectionsDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l MoveCollectionsDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (MoveCollectionsDeleteResponse, error) {
	respType := MoveCollectionsDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.OperationStatus)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a MoveCollectionsDeletePollerResponse from the provided client and resume token.
func (l *MoveCollectionsDeletePollerResponse) Resume(ctx context.Context, client *MoveCollectionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("MoveCollectionsClient.Delete", token, client.pl, client.deleteHandleError)
	if err != nil {
		return err
	}
	poller := &MoveCollectionsDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// MoveCollectionsDeleteResponse contains the response from method MoveCollections.Delete.
type MoveCollectionsDeleteResponse struct {
	MoveCollectionsDeleteResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MoveCollectionsDeleteResult contains the result from method MoveCollections.Delete.
type MoveCollectionsDeleteResult struct {
	OperationStatus
}

// MoveCollectionsDiscardPollerResponse contains the response from method MoveCollections.Discard.
type MoveCollectionsDiscardPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *MoveCollectionsDiscardPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l MoveCollectionsDiscardPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (MoveCollectionsDiscardResponse, error) {
	respType := MoveCollectionsDiscardResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.OperationStatus)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a MoveCollectionsDiscardPollerResponse from the provided client and resume token.
func (l *MoveCollectionsDiscardPollerResponse) Resume(ctx context.Context, client *MoveCollectionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("MoveCollectionsClient.Discard", token, client.pl, client.discardHandleError)
	if err != nil {
		return err
	}
	poller := &MoveCollectionsDiscardPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// MoveCollectionsDiscardResponse contains the response from method MoveCollections.Discard.
type MoveCollectionsDiscardResponse struct {
	MoveCollectionsDiscardResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MoveCollectionsDiscardResult contains the result from method MoveCollections.Discard.
type MoveCollectionsDiscardResult struct {
	OperationStatus
}

// MoveCollectionsGetResponse contains the response from method MoveCollections.Get.
type MoveCollectionsGetResponse struct {
	MoveCollectionsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MoveCollectionsGetResult contains the result from method MoveCollections.Get.
type MoveCollectionsGetResult struct {
	MoveCollection
}

// MoveCollectionsInitiateMovePollerResponse contains the response from method MoveCollections.InitiateMove.
type MoveCollectionsInitiateMovePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *MoveCollectionsInitiateMovePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l MoveCollectionsInitiateMovePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (MoveCollectionsInitiateMoveResponse, error) {
	respType := MoveCollectionsInitiateMoveResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.OperationStatus)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a MoveCollectionsInitiateMovePollerResponse from the provided client and resume token.
func (l *MoveCollectionsInitiateMovePollerResponse) Resume(ctx context.Context, client *MoveCollectionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("MoveCollectionsClient.InitiateMove", token, client.pl, client.initiateMoveHandleError)
	if err != nil {
		return err
	}
	poller := &MoveCollectionsInitiateMovePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// MoveCollectionsInitiateMoveResponse contains the response from method MoveCollections.InitiateMove.
type MoveCollectionsInitiateMoveResponse struct {
	MoveCollectionsInitiateMoveResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MoveCollectionsInitiateMoveResult contains the result from method MoveCollections.InitiateMove.
type MoveCollectionsInitiateMoveResult struct {
	OperationStatus
}

// MoveCollectionsListMoveCollectionsByResourceGroupResponse contains the response from method MoveCollections.ListMoveCollectionsByResourceGroup.
type MoveCollectionsListMoveCollectionsByResourceGroupResponse struct {
	MoveCollectionsListMoveCollectionsByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MoveCollectionsListMoveCollectionsByResourceGroupResult contains the result from method MoveCollections.ListMoveCollectionsByResourceGroup.
type MoveCollectionsListMoveCollectionsByResourceGroupResult struct {
	MoveCollectionResultList
}

// MoveCollectionsListMoveCollectionsBySubscriptionResponse contains the response from method MoveCollections.ListMoveCollectionsBySubscription.
type MoveCollectionsListMoveCollectionsBySubscriptionResponse struct {
	MoveCollectionsListMoveCollectionsBySubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MoveCollectionsListMoveCollectionsBySubscriptionResult contains the result from method MoveCollections.ListMoveCollectionsBySubscription.
type MoveCollectionsListMoveCollectionsBySubscriptionResult struct {
	MoveCollectionResultList
}

// MoveCollectionsListRequiredForResponse contains the response from method MoveCollections.ListRequiredFor.
type MoveCollectionsListRequiredForResponse struct {
	MoveCollectionsListRequiredForResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MoveCollectionsListRequiredForResult contains the result from method MoveCollections.ListRequiredFor.
type MoveCollectionsListRequiredForResult struct {
	RequiredForResourcesCollection
}

// MoveCollectionsPreparePollerResponse contains the response from method MoveCollections.Prepare.
type MoveCollectionsPreparePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *MoveCollectionsPreparePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l MoveCollectionsPreparePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (MoveCollectionsPrepareResponse, error) {
	respType := MoveCollectionsPrepareResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.OperationStatus)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a MoveCollectionsPreparePollerResponse from the provided client and resume token.
func (l *MoveCollectionsPreparePollerResponse) Resume(ctx context.Context, client *MoveCollectionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("MoveCollectionsClient.Prepare", token, client.pl, client.prepareHandleError)
	if err != nil {
		return err
	}
	poller := &MoveCollectionsPreparePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// MoveCollectionsPrepareResponse contains the response from method MoveCollections.Prepare.
type MoveCollectionsPrepareResponse struct {
	MoveCollectionsPrepareResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MoveCollectionsPrepareResult contains the result from method MoveCollections.Prepare.
type MoveCollectionsPrepareResult struct {
	OperationStatus
}

// MoveCollectionsResolveDependenciesPollerResponse contains the response from method MoveCollections.ResolveDependencies.
type MoveCollectionsResolveDependenciesPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *MoveCollectionsResolveDependenciesPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l MoveCollectionsResolveDependenciesPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (MoveCollectionsResolveDependenciesResponse, error) {
	respType := MoveCollectionsResolveDependenciesResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.OperationStatus)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a MoveCollectionsResolveDependenciesPollerResponse from the provided client and resume token.
func (l *MoveCollectionsResolveDependenciesPollerResponse) Resume(ctx context.Context, client *MoveCollectionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("MoveCollectionsClient.ResolveDependencies", token, client.pl, client.resolveDependenciesHandleError)
	if err != nil {
		return err
	}
	poller := &MoveCollectionsResolveDependenciesPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// MoveCollectionsResolveDependenciesResponse contains the response from method MoveCollections.ResolveDependencies.
type MoveCollectionsResolveDependenciesResponse struct {
	MoveCollectionsResolveDependenciesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MoveCollectionsResolveDependenciesResult contains the result from method MoveCollections.ResolveDependencies.
type MoveCollectionsResolveDependenciesResult struct {
	OperationStatus
}

// MoveCollectionsUpdateResponse contains the response from method MoveCollections.Update.
type MoveCollectionsUpdateResponse struct {
	MoveCollectionsUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MoveCollectionsUpdateResult contains the result from method MoveCollections.Update.
type MoveCollectionsUpdateResult struct {
	MoveCollection
}

// MoveResourcesCreatePollerResponse contains the response from method MoveResources.Create.
type MoveResourcesCreatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *MoveResourcesCreatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l MoveResourcesCreatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (MoveResourcesCreateResponse, error) {
	respType := MoveResourcesCreateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.MoveResource)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a MoveResourcesCreatePollerResponse from the provided client and resume token.
func (l *MoveResourcesCreatePollerResponse) Resume(ctx context.Context, client *MoveResourcesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("MoveResourcesClient.Create", token, client.pl, client.createHandleError)
	if err != nil {
		return err
	}
	poller := &MoveResourcesCreatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// MoveResourcesCreateResponse contains the response from method MoveResources.Create.
type MoveResourcesCreateResponse struct {
	MoveResourcesCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MoveResourcesCreateResult contains the result from method MoveResources.Create.
type MoveResourcesCreateResult struct {
	MoveResource
}

// MoveResourcesDeletePollerResponse contains the response from method MoveResources.Delete.
type MoveResourcesDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *MoveResourcesDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l MoveResourcesDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (MoveResourcesDeleteResponse, error) {
	respType := MoveResourcesDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.OperationStatus)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a MoveResourcesDeletePollerResponse from the provided client and resume token.
func (l *MoveResourcesDeletePollerResponse) Resume(ctx context.Context, client *MoveResourcesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("MoveResourcesClient.Delete", token, client.pl, client.deleteHandleError)
	if err != nil {
		return err
	}
	poller := &MoveResourcesDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// MoveResourcesDeleteResponse contains the response from method MoveResources.Delete.
type MoveResourcesDeleteResponse struct {
	MoveResourcesDeleteResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MoveResourcesDeleteResult contains the result from method MoveResources.Delete.
type MoveResourcesDeleteResult struct {
	OperationStatus
}

// MoveResourcesGetResponse contains the response from method MoveResources.Get.
type MoveResourcesGetResponse struct {
	MoveResourcesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MoveResourcesGetResult contains the result from method MoveResources.Get.
type MoveResourcesGetResult struct {
	MoveResource
}

// MoveResourcesListResponse contains the response from method MoveResources.List.
type MoveResourcesListResponse struct {
	MoveResourcesListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MoveResourcesListResult contains the result from method MoveResources.List.
type MoveResourcesListResult struct {
	MoveResourceCollection
}

// OperationsDiscoveryGetResponse contains the response from method OperationsDiscovery.Get.
type OperationsDiscoveryGetResponse struct {
	OperationsDiscoveryGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsDiscoveryGetResult contains the result from method OperationsDiscovery.Get.
type OperationsDiscoveryGetResult struct {
	OperationsDiscoveryCollection
}

// UnresolvedDependenciesGetResponse contains the response from method UnresolvedDependencies.Get.
type UnresolvedDependenciesGetResponse struct {
	UnresolvedDependenciesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// UnresolvedDependenciesGetResult contains the result from method UnresolvedDependencies.Get.
type UnresolvedDependenciesGetResult struct {
	UnresolvedDependencyCollection
}