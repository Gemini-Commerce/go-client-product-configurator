/*
Product Configurator Service

Testing DependencyAPIService

*/

// Code generated by OpenAPI Generator (https://openapi_generator.tech);

package productconfigurator

import (
	"context"
	openapiclient "github.com/Gemini-Commerce/go_client_product-configurator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_product_configurator_DependencyAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DependencyAPIService ProductConfiguratorCreateDependency", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tenantId string
		var stepId string

		resp, httpRes, err := apiClient.DependencyAPI.ProductConfiguratorCreateDependency(context.Background(), tenantId, stepId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DependencyAPIService ProductConfiguratorDeleteDependency", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tenantId string
		var dependencyId string

		resp, httpRes, err := apiClient.DependencyAPI.ProductConfiguratorDeleteDependency(context.Background(), tenantId, dependencyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DependencyAPIService ProductConfiguratorListDependencies", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tenantId string
		var pageSize int64

		resp, httpRes, err := apiClient.DependencyAPI.ProductConfiguratorListDependencies(context.Background(), tenantId, pageSize).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DependencyAPIService ProductConfiguratorUpdateDependency", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tenantId string
		var dependencyId string

		resp, httpRes, err := apiClient.DependencyAPI.ProductConfiguratorUpdateDependency(context.Background(), tenantId, dependencyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
