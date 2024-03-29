/*
Product Configurator Service

Testing ConfigurationAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package product-configurator

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/Gemini-Commerce/go-client-product-configurator"
)

func Test_product-configurator_ConfigurationAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ConfigurationAPIService ProductConfiguratorGetAvailableConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tenantId string
		var productId string

		resp, httpRes, err := apiClient.ConfigurationAPI.ProductConfiguratorGetAvailableConfiguration(context.Background(), tenantId, productId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigurationAPIService ProductConfiguratorGetAvailableConfiguration2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tenantId string
		var productId string

		resp, httpRes, err := apiClient.ConfigurationAPI.ProductConfiguratorGetAvailableConfiguration2(context.Background(), tenantId, productId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigurationAPIService ProductConfiguratorGetConfigurationFromSelections", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tenantId string
		var productId string

		resp, httpRes, err := apiClient.ConfigurationAPI.ProductConfiguratorGetConfigurationFromSelections(context.Background(), tenantId, productId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
