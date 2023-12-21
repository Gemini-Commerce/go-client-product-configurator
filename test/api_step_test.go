/*
Product Configurator Service

Testing StepAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package product_configurator

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/gemini-commerce/go-client-product_configurator"
)

func Test_product_configurator_StepAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test StepAPIService ProductConfiguratorBulkCreateSteps", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tenantId string
		var configuratorId string

		resp, httpRes, err := apiClient.StepAPI.ProductConfiguratorBulkCreateSteps(context.Background(), tenantId, configuratorId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StepAPIService ProductConfiguratorBulkDeleteSteps", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tenantId string

		resp, httpRes, err := apiClient.StepAPI.ProductConfiguratorBulkDeleteSteps(context.Background(), tenantId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StepAPIService ProductConfiguratorCopyStep", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tenantId string
		var sourceStepId string

		resp, httpRes, err := apiClient.StepAPI.ProductConfiguratorCopyStep(context.Background(), tenantId, sourceStepId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StepAPIService ProductConfiguratorCreateStep", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tenantId string
		var configuratorId string

		resp, httpRes, err := apiClient.StepAPI.ProductConfiguratorCreateStep(context.Background(), tenantId, configuratorId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StepAPIService ProductConfiguratorDeleteStep", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tenantId string
		var stepId string

		resp, httpRes, err := apiClient.StepAPI.ProductConfiguratorDeleteStep(context.Background(), tenantId, stepId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StepAPIService ProductConfiguratorUpdateStep", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tenantId string
		var stepId string

		resp, httpRes, err := apiClient.StepAPI.ProductConfiguratorUpdateStep(context.Background(), tenantId, stepId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}