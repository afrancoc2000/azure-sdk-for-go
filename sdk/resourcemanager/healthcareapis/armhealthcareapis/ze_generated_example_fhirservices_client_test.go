//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhealthcareapis_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis"
)

// x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/preview/2021-06-01-preview/examples/fhirservices/FhirServices_List.json
func ExampleFhirServicesClient_ListByWorkspace() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhealthcareapis.NewFhirServicesClient("<subscription-id>", cred, nil)
	pager := client.ListByWorkspace("<resource-group-name>",
		"<workspace-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("FhirService.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/preview/2021-06-01-preview/examples/fhirservices/FhirServices_Get.json
func ExampleFhirServicesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhealthcareapis.NewFhirServicesClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<fhir-service-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("FhirService.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/preview/2021-06-01-preview/examples/fhirservices/FhirServices_Create.json
func ExampleFhirServicesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhealthcareapis.NewFhirServicesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<fhir-service-name>",
		armhealthcareapis.FhirService{
			ServiceManagedIdentity: armhealthcareapis.ServiceManagedIdentity{
				Identity: &armhealthcareapis.ServiceManagedIdentityIdentity{
					Type: armhealthcareapis.ManagedServiceIdentityTypeSystemAssigned.ToPtr(),
				},
			},
			TaggedResource: armhealthcareapis.TaggedResource{
				LocationBasedResource: armhealthcareapis.LocationBasedResource{
					Location: to.StringPtr("<location>"),
				},
				ResourceTags: armhealthcareapis.ResourceTags{
					Tags: map[string]*string{
						"additionalProp1": to.StringPtr("string"),
						"additionalProp2": to.StringPtr("string"),
						"additionalProp3": to.StringPtr("string"),
					},
				},
			},
			Kind: armhealthcareapis.FhirServiceKindFhirR4.ToPtr(),
			Properties: &armhealthcareapis.FhirServiceProperties{
				AccessPolicies: []*armhealthcareapis.FhirServiceAccessPolicyEntry{
					{
						ObjectID: to.StringPtr("<object-id>"),
					},
					{
						ObjectID: to.StringPtr("<object-id>"),
					}},
				AcrConfiguration: &armhealthcareapis.FhirServiceAcrConfiguration{
					LoginServers: []*string{
						to.StringPtr("test1.azurecr.io")},
				},
				AuthenticationConfiguration: &armhealthcareapis.FhirServiceAuthenticationConfiguration{
					Audience:          to.StringPtr("<audience>"),
					Authority:         to.StringPtr("<authority>"),
					SmartProxyEnabled: to.BoolPtr(true),
				},
				CorsConfiguration: &armhealthcareapis.FhirServiceCorsConfiguration{
					AllowCredentials: to.BoolPtr(false),
					Headers: []*string{
						to.StringPtr("*")},
					MaxAge: to.Int32Ptr(1440),
					Methods: []*string{
						to.StringPtr("DELETE"),
						to.StringPtr("GET"),
						to.StringPtr("OPTIONS"),
						to.StringPtr("PATCH"),
						to.StringPtr("POST"),
						to.StringPtr("PUT")},
					Origins: []*string{
						to.StringPtr("*")},
				},
				ExportConfiguration: &armhealthcareapis.FhirServiceExportConfiguration{
					StorageAccountName: to.StringPtr("<storage-account-name>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("FhirService.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/preview/2021-06-01-preview/examples/fhirservices/FhirServices_Patch.json
func ExampleFhirServicesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhealthcareapis.NewFhirServicesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<fhir-service-name>",
		"<workspace-name>",
		armhealthcareapis.FhirServicePatchResource{
			ResourceTags: armhealthcareapis.ResourceTags{
				Tags: map[string]*string{
					"tagKey": to.StringPtr("tagValue"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("FhirService.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/preview/2021-06-01-preview/examples/fhirservices/FhirServices_Delete.json
func ExampleFhirServicesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhealthcareapis.NewFhirServicesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<fhir-service-name>",
		"<workspace-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}