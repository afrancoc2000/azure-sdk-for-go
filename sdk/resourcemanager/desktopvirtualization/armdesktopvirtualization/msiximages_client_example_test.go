//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdesktopvirtualization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/MsixImage_Expand_Post.json
func ExampleMsixImagesClient_NewExpandPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdesktopvirtualization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMsixImagesClient().NewExpandPager("resourceGroup1", "hostpool1", armdesktopvirtualization.MSIXImageURI{
		URI: to.Ptr("imagepath"),
	}, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ExpandMsixImageList = armdesktopvirtualization.ExpandMsixImageList{
		// 	Value: []*armdesktopvirtualization.ExpandMsixImage{
		// 		{
		// 			Name: to.Ptr("hostpool1/expandmsiximage"),
		// 			Type: to.Ptr("Microsoft.DesktopVirtualization/hostpools/expandmsiximage"),
		// 			ID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourcegroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/hostpools/hostpool1/expandmsiximage"),
		// 			Properties: &armdesktopvirtualization.ExpandMsixImageProperties{
		// 				DisplayName: to.Ptr("displayname"),
		// 				ImagePath: to.Ptr("imagepath"),
		// 				IsActive: to.Ptr(false),
		// 				IsRegularRegistration: to.Ptr(false),
		// 				LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2008-09-22T14:01:54.9571247Z"); return t}()),
		// 				PackageAlias: to.Ptr("msixpackagealias"),
		// 				PackageApplications: []*armdesktopvirtualization.MsixPackageApplications{
		// 					{
		// 						Description: to.Ptr("PackageApplicationDescription"),
		// 						AppID: to.Ptr("AppId"),
		// 						AppUserModelID: to.Ptr("AppUserModelId"),
		// 						FriendlyName: to.Ptr("FriendlyName"),
		// 						IconImageName: to.Ptr("Iconimagename"),
		// 						RawIcon: []byte("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
		// 						RawPNG: []byte("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
		// 				}},
		// 				PackageDependencies: []*armdesktopvirtualization.MsixPackageDependencies{
		// 				},
		// 				PackageFamilyName: to.Ptr("MsixPackage_FamilyName"),
		// 				PackageFullName: to.Ptr("MsixPackage_FullName"),
		// 				PackageName: to.Ptr("MsixPackageName"),
		// 				PackageRelativePath: to.Ptr("packagerelativepath"),
		// 				Version: to.Ptr("packageversion"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("hostpool1/expandmsiximage"),
		// 			Type: to.Ptr("Microsoft.DesktopVirtualization/hostpools/expandmsiximage"),
		// 			ID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourcegroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/hostpools/hostpool1/expandmsiximage"),
		// 			Properties: &armdesktopvirtualization.ExpandMsixImageProperties{
		// 				DisplayName: to.Ptr("displayname2"),
		// 				ImagePath: to.Ptr("imagepath"),
		// 				IsActive: to.Ptr(false),
		// 				IsRegularRegistration: to.Ptr(false),
		// 				LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2008-09-22T14:01:54.9571247Z"); return t}()),
		// 				PackageAlias: to.Ptr("msixpackagealias2"),
		// 				PackageApplications: []*armdesktopvirtualization.MsixPackageApplications{
		// 					{
		// 						Description: to.Ptr("PackageApplicationDescription1"),
		// 						AppID: to.Ptr("AppId1"),
		// 						AppUserModelID: to.Ptr("AppUserModelId1"),
		// 						FriendlyName: to.Ptr("FriendlyName1"),
		// 						IconImageName: to.Ptr("Iconimagename1"),
		// 						RawIcon: []byte("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
		// 						RawPNG: []byte("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
		// 					},
		// 					{
		// 						Description: to.Ptr("PackageApplicationDescription2"),
		// 						AppID: to.Ptr("AppId2"),
		// 						AppUserModelID: to.Ptr("AppUserModelId2"),
		// 						FriendlyName: to.Ptr("FriendlyName2"),
		// 						IconImageName: to.Ptr("Iconimagename2"),
		// 						RawIcon: []byte("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
		// 						RawPNG: []byte("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
		// 				}},
		// 				PackageDependencies: []*armdesktopvirtualization.MsixPackageDependencies{
		// 					{
		// 						DependencyName: to.Ptr("MsixPackageDependency1"),
		// 						MinVersion: to.Ptr("ver1"),
		// 						Publisher: to.Ptr("PublisherName1"),
		// 					},
		// 					{
		// 						DependencyName: to.Ptr("MsixPackageDependency2"),
		// 						MinVersion: to.Ptr("ver2"),
		// 						Publisher: to.Ptr("PublisherName2"),
		// 				}},
		// 				PackageFamilyName: to.Ptr("MsixPackage_FamilyName2"),
		// 				PackageFullName: to.Ptr("MsixPackage_FullName2"),
		// 				PackageName: to.Ptr("MsixPackageName2"),
		// 				PackageRelativePath: to.Ptr("packagerelativepath2"),
		// 				Version: to.Ptr("packageversion"),
		// 			},
		// 	}},
		// }
	}
}
