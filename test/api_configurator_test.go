/*
Product Configurator Service

Testing ConfiguratorAPIService

*/

// Code generated by OpenAPI Generator (https://openapi_generator.tech);

package product_configurator

import (
	"context"
	openapiclient "github.com/Gemini_Commerce/go_client_product_configurator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_product_configurator_ConfiguratorAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ConfiguratorAPIService ProductConfiguratorCopyConfigurator", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tenantId string
		var sourceConfiguratorId string

		resp, httpRes, err := apiClient.ConfiguratorAPI.ProductConfiguratorCopyConfigurator(context.Background(), tenantId, sourceConfiguratorId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfiguratorAPIService ProductConfiguratorCreateConfigurator", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tenantId string
		var productId string

		resp, httpRes, err := apiClient.ConfiguratorAPI.ProductConfiguratorCreateConfigurator(context.Background(), tenantId, productId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfiguratorAPIService ProductConfiguratorDeleteConfigurator", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tenantId string
		var configuratorId string

		resp, httpRes, err := apiClient.ConfiguratorAPI.ProductConfiguratorDeleteConfigurator(context.Background(), tenantId, configuratorId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfiguratorAPIService ProductConfiguratorGetConfiguratorByProductId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tenantId string
		var productId string

		resp, httpRes, err := apiClient.ConfiguratorAPI.ProductConfiguratorGetConfiguratorByProductId(context.Background(), tenantId, productId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfiguratorAPIService ProductConfiguratorGetConfiguratorByProductId2", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tenantId string
		var productId string
		var status string

		resp, httpRes, err := apiClient.ConfiguratorAPI.ProductConfiguratorGetConfiguratorByProductId2(context.Background(), tenantId, productId, status).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfiguratorAPIService ProductConfiguratorListConfigurators", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tenantId string
		var productId string
		var pageSize int64

		resp, httpRes, err := apiClient.ConfiguratorAPI.ProductConfiguratorListConfigurators(context.Background(), tenantId, productId, pageSize).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfiguratorAPIService ProductConfiguratorUpdateConfigurator", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tenantId string
		var configuratorId string

		resp, httpRes, err := apiClient.ConfiguratorAPI.ProductConfiguratorUpdateConfigurator(context.Background(), tenantId, configuratorId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
