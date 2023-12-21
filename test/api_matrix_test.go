/*
Product Configurator Service

Testing MatrixAPIService

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

func Test_product_configurator_MatrixAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MatrixAPIService ProductConfiguratorCreateMatrix", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tenantId string

		resp, httpRes, err := apiClient.MatrixAPI.ProductConfiguratorCreateMatrix(context.Background(), tenantId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MatrixAPIService ProductConfiguratorDeleteMatrix", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tenantId string
		var matrixId string

		resp, httpRes, err := apiClient.MatrixAPI.ProductConfiguratorDeleteMatrix(context.Background(), tenantId, matrixId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MatrixAPIService ProductConfiguratorGetMatrix", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tenantId string
		var matrixId string

		resp, httpRes, err := apiClient.MatrixAPI.ProductConfiguratorGetMatrix(context.Background(), tenantId, matrixId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MatrixAPIService ProductConfiguratorListMatrices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tenantId string
		var configuratorId string
		var pageSize string

		resp, httpRes, err := apiClient.MatrixAPI.ProductConfiguratorListMatrices(context.Background(), tenantId, configuratorId, pageSize).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MatrixAPIService ProductConfiguratorRemovePricelistFromMatrix", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tenantId string
		var matrixId string
		var pricelistGrn string

		resp, httpRes, err := apiClient.MatrixAPI.ProductConfiguratorRemovePricelistFromMatrix(context.Background(), tenantId, matrixId, pricelistGrn).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MatrixAPIService ProductConfiguratorUpdateMatrix", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tenantId string
		var matrixId string

		resp, httpRes, err := apiClient.MatrixAPI.ProductConfiguratorUpdateMatrix(context.Background(), tenantId, matrixId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}