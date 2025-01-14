//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v4"
	"net/http"
	"net/url"
	"regexp"
)

// TrustedAccessRolesServer is a fake server for instances of the armcontainerservice.TrustedAccessRolesClient type.
type TrustedAccessRolesServer struct {
	// NewListPager is the fake for method TrustedAccessRolesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(location string, options *armcontainerservice.TrustedAccessRolesClientListOptions) (resp azfake.PagerResponder[armcontainerservice.TrustedAccessRolesClientListResponse])
}

// NewTrustedAccessRolesServerTransport creates a new instance of TrustedAccessRolesServerTransport with the provided implementation.
// The returned TrustedAccessRolesServerTransport instance is connected to an instance of armcontainerservice.TrustedAccessRolesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewTrustedAccessRolesServerTransport(srv *TrustedAccessRolesServer) *TrustedAccessRolesServerTransport {
	return &TrustedAccessRolesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armcontainerservice.TrustedAccessRolesClientListResponse]](),
	}
}

// TrustedAccessRolesServerTransport connects instances of armcontainerservice.TrustedAccessRolesClient to instances of TrustedAccessRolesServer.
// Don't use this type directly, use NewTrustedAccessRolesServerTransport instead.
type TrustedAccessRolesServerTransport struct {
	srv          *TrustedAccessRolesServer
	newListPager *tracker[azfake.PagerResponder[armcontainerservice.TrustedAccessRolesClientListResponse]]
}

// Do implements the policy.Transporter interface for TrustedAccessRolesServerTransport.
func (t *TrustedAccessRolesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "TrustedAccessRolesClient.NewListPager":
		resp, err = t.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (t *TrustedAccessRolesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if t.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := t.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.ContainerService/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/trustedAccessRoles`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		locationUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
		if err != nil {
			return nil, err
		}
		resp := t.srv.NewListPager(locationUnescaped, nil)
		newListPager = &resp
		t.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armcontainerservice.TrustedAccessRolesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		t.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		t.newListPager.remove(req)
	}
	return resp, nil
}
