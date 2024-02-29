/*
Product Configurator Service

Testing PropertyAPIService

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

func Test_product-configurator_PropertyAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PropertyAPIService ProductConfiguratorBulkCreateProperties", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tenantId string

		resp, httpRes, err := apiClient.PropertyAPI.ProductConfiguratorBulkCreateProperties(context.Background(), tenantId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PropertyAPIService ProductConfiguratorBulkUpdateProperties", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tenantId string

		resp, httpRes, err := apiClient.PropertyAPI.ProductConfiguratorBulkUpdateProperties(context.Background(), tenantId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PropertyAPIService ProductConfiguratorCreateProperty", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tenantId string

		resp, httpRes, err := apiClient.PropertyAPI.ProductConfiguratorCreateProperty(context.Background(), tenantId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PropertyAPIService ProductConfiguratorListProperties", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tenantId string
		var matrixId string
		var pageSize string

		resp, httpRes, err := apiClient.PropertyAPI.ProductConfiguratorListProperties(context.Background(), tenantId, matrixId, pageSize).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PropertyAPIService ProductConfiguratorUpdateProperty", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tenantId string
		var propertyId string

		resp, httpRes, err := apiClient.PropertyAPI.ProductConfiguratorUpdateProperty(context.Background(), tenantId, propertyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
