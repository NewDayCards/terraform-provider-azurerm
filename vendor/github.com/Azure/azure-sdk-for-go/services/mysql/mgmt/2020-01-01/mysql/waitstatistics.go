package mysql

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// WaitStatisticsClient is the the Microsoft Azure management API provides create, read, update, and delete
// functionality for Azure MySQL resources including servers, databases, firewall rules, VNET rules, log files and
// configurations with new business model.
type WaitStatisticsClient struct {
	BaseClient
}

// NewWaitStatisticsClient creates an instance of the WaitStatisticsClient client.
func NewWaitStatisticsClient(subscriptionID string) WaitStatisticsClient {
	return NewWaitStatisticsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWaitStatisticsClientWithBaseURI creates an instance of the WaitStatisticsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWaitStatisticsClientWithBaseURI(baseURI string, subscriptionID string) WaitStatisticsClient {
	return WaitStatisticsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get retrieve wait statistics for specified identifier.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// serverName - the name of the server.
// waitStatisticsID - the Wait Statistic identifier.
func (client WaitStatisticsClient) Get(ctx context.Context, resourceGroupName string, serverName string, waitStatisticsID string) (result WaitStatistic, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WaitStatisticsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("mysql.WaitStatisticsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, serverName, waitStatisticsID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mysql.WaitStatisticsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "mysql.WaitStatisticsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mysql.WaitStatisticsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client WaitStatisticsClient) GetPreparer(ctx context.Context, resourceGroupName string, serverName string, waitStatisticsID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"waitStatisticsId":  autorest.Encode("path", waitStatisticsID),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/waitStatistics/{waitStatisticsId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client WaitStatisticsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client WaitStatisticsClient) GetResponder(resp *http.Response) (result WaitStatistic, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByServer retrieve wait statistics for specified aggregation window.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// serverName - the name of the server.
// parameters - the required parameters for retrieving wait statistics.
func (client WaitStatisticsClient) ListByServer(ctx context.Context, resourceGroupName string, serverName string, parameters WaitStatisticsInput) (result WaitStatisticsResultListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WaitStatisticsClient.ListByServer")
		defer func() {
			sc := -1
			if result.wsrl.Response.Response != nil {
				sc = result.wsrl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.WaitStatisticsInputProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.WaitStatisticsInputProperties.ObservationStartTime", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.WaitStatisticsInputProperties.ObservationEndTime", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.WaitStatisticsInputProperties.AggregationWindow", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("mysql.WaitStatisticsClient", "ListByServer", err.Error())
	}

	result.fn = client.listByServerNextResults
	req, err := client.ListByServerPreparer(ctx, resourceGroupName, serverName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mysql.WaitStatisticsClient", "ListByServer", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServerSender(req)
	if err != nil {
		result.wsrl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "mysql.WaitStatisticsClient", "ListByServer", resp, "Failure sending request")
		return
	}

	result.wsrl, err = client.ListByServerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mysql.WaitStatisticsClient", "ListByServer", resp, "Failure responding to request")
		return
	}
	if result.wsrl.hasNextLink() && result.wsrl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByServerPreparer prepares the ListByServer request.
func (client WaitStatisticsClient) ListByServerPreparer(ctx context.Context, resourceGroupName string, serverName string, parameters WaitStatisticsInput) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/waitStatistics", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByServerSender sends the ListByServer request. The method will close the
// http.Response Body if it receives an error.
func (client WaitStatisticsClient) ListByServerSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByServerResponder handles the response to the ListByServer request. The method always
// closes the http.Response Body.
func (client WaitStatisticsClient) ListByServerResponder(resp *http.Response) (result WaitStatisticsResultList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByServerNextResults retrieves the next set of results, if any.
func (client WaitStatisticsClient) listByServerNextResults(ctx context.Context, lastResults WaitStatisticsResultList) (result WaitStatisticsResultList, err error) {
	req, err := lastResults.waitStatisticsResultListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "mysql.WaitStatisticsClient", "listByServerNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByServerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "mysql.WaitStatisticsClient", "listByServerNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByServerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mysql.WaitStatisticsClient", "listByServerNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByServerComplete enumerates all values, automatically crossing page boundaries as required.
func (client WaitStatisticsClient) ListByServerComplete(ctx context.Context, resourceGroupName string, serverName string, parameters WaitStatisticsInput) (result WaitStatisticsResultListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WaitStatisticsClient.ListByServer")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByServer(ctx, resourceGroupName, serverName, parameters)
	return
}
