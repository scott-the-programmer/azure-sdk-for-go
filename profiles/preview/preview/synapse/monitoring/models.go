//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package monitoring

import original "github.com/Azure/azure-sdk-for-go/services/preview/synapse/2019-11-01-preview/monitoring"

type BaseClient = original.BaseClient
type Client = original.Client
type SQLQueryStringDataModel = original.SQLQueryStringDataModel
type SparkJob = original.SparkJob
type SparkJobListViewResponse = original.SparkJobListViewResponse

func New(endpoint string) BaseClient {
	return original.New(endpoint)
}
func NewClient(endpoint string) Client {
	return original.NewClient(endpoint)
}
func NewWithoutDefaults(endpoint string) BaseClient {
	return original.NewWithoutDefaults(endpoint)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
