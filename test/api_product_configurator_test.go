/*
Product Configurator Service

Testing ProductConfiguratorAPIService

*/

// Code generated by OpenAPI Generator (https://openapi_generator.tech);

package productconfigurator

import (
	"context"
	openapiclient "github.com/Gemini-Commerce/go_client_product_configurator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_product_configurator_ProductConfiguratorAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ProductConfiguratorAPIService ProductConfiguratorGetProperty", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tenantId string
		var propertyId string

		resp, httpRes, err := apiClient.ProductConfiguratorAPI.ProductConfiguratorGetProperty(context.Background(), tenantId, propertyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProductConfiguratorAPIService ProductConfiguratorListPropertiesByConfiguration", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tenantId string
		var configuratorId string
		var pageSize string

		resp, httpRes, err := apiClient.ProductConfiguratorAPI.ProductConfiguratorListPropertiesByConfiguration(context.Background(), tenantId, configuratorId, pageSize).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
