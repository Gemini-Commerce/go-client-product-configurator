/*
Product Configurator Service

Testing MatrixAPIService

*/

// Code generated by OpenAPI Generator (https://openapi_generator.tech);

package productconfigurator

import (
	"context"
	openapiclient "github.com/Gemini_Commerce/go_client_product_configurator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_product_configurator_MatrixAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MatrixAPIService ProductConfiguratorCreateMatrix", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tenantId string

		resp, httpRes, err := apiClient.MatrixAPI.ProductConfiguratorCreateMatrix(context.Background(), tenantId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MatrixAPIService ProductConfiguratorDeleteMatrix", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tenantId string
		var matrixId string

		resp, httpRes, err := apiClient.MatrixAPI.ProductConfiguratorDeleteMatrix(context.Background(), tenantId, matrixId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MatrixAPIService ProductConfiguratorGetMatrix", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tenantId string
		var matrixId string

		resp, httpRes, err := apiClient.MatrixAPI.ProductConfiguratorGetMatrix(context.Background(), tenantId, matrixId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MatrixAPIService ProductConfiguratorListMatrices", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tenantId string
		var configuratorId string
		var pageSize string

		resp, httpRes, err := apiClient.MatrixAPI.ProductConfiguratorListMatrices(context.Background(), tenantId, configuratorId, pageSize).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MatrixAPIService ProductConfiguratorRemovePricelistFromMatrix", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tenantId string
		var matrixId string
		var pricelistGrn string

		resp, httpRes, err := apiClient.MatrixAPI.ProductConfiguratorRemovePricelistFromMatrix(context.Background(), tenantId, matrixId, pricelistGrn).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MatrixAPIService ProductConfiguratorUpdateMatrix", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tenantId string
		var matrixId string

		resp, httpRes, err := apiClient.MatrixAPI.ProductConfiguratorUpdateMatrix(context.Background(), tenantId, matrixId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
