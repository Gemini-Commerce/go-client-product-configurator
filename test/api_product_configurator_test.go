/*
Product Configurator Service

Testing ProductConfiguratorApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package product_configurator

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_product_configurator_ProductConfiguratorApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test ProductConfiguratorApiService ProductConfiguratorBulkCreateOptions", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string
        var stepId string

        resp, httpRes, err := apiClient.ProductConfiguratorApi.ProductConfiguratorBulkCreateOptions(context.Background(), tenantId, stepId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProductConfiguratorApiService ProductConfiguratorBulkCreateProperties", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string

        resp, httpRes, err := apiClient.ProductConfiguratorApi.ProductConfiguratorBulkCreateProperties(context.Background(), tenantId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProductConfiguratorApiService ProductConfiguratorBulkCreateSteps", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string
        var configuratorId string

        resp, httpRes, err := apiClient.ProductConfiguratorApi.ProductConfiguratorBulkCreateSteps(context.Background(), tenantId, configuratorId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProductConfiguratorApiService ProductConfiguratorBulkDeleteOptions", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string

        resp, httpRes, err := apiClient.ProductConfiguratorApi.ProductConfiguratorBulkDeleteOptions(context.Background(), tenantId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProductConfiguratorApiService ProductConfiguratorBulkDeleteSteps", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string

        resp, httpRes, err := apiClient.ProductConfiguratorApi.ProductConfiguratorBulkDeleteSteps(context.Background(), tenantId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProductConfiguratorApiService ProductConfiguratorBulkUpdateOptions", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string

        resp, httpRes, err := apiClient.ProductConfiguratorApi.ProductConfiguratorBulkUpdateOptions(context.Background(), tenantId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProductConfiguratorApiService ProductConfiguratorBulkUpdateProperties", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string

        resp, httpRes, err := apiClient.ProductConfiguratorApi.ProductConfiguratorBulkUpdateProperties(context.Background(), tenantId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProductConfiguratorApiService ProductConfiguratorCopyConfigurator2", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string
        var sourceProductId string

        resp, httpRes, err := apiClient.ProductConfiguratorApi.ProductConfiguratorCopyConfigurator2(context.Background(), tenantId, sourceProductId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProductConfiguratorApiService ProductConfiguratorCopyOption", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string
        var sourceOptionId string

        resp, httpRes, err := apiClient.ProductConfiguratorApi.ProductConfiguratorCopyOption(context.Background(), tenantId, sourceOptionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProductConfiguratorApiService ProductConfiguratorCopyStep", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string
        var sourceStepId string

        resp, httpRes, err := apiClient.ProductConfiguratorApi.ProductConfiguratorCopyStep(context.Background(), tenantId, sourceStepId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProductConfiguratorApiService ProductConfiguratorCreateConfigurator", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string
        var productId string

        resp, httpRes, err := apiClient.ProductConfiguratorApi.ProductConfiguratorCreateConfigurator(context.Background(), tenantId, productId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProductConfiguratorApiService ProductConfiguratorCreateDependency", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string
        var stepId string

        resp, httpRes, err := apiClient.ProductConfiguratorApi.ProductConfiguratorCreateDependency(context.Background(), tenantId, stepId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProductConfiguratorApiService ProductConfiguratorCreateMatrix", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string

        resp, httpRes, err := apiClient.ProductConfiguratorApi.ProductConfiguratorCreateMatrix(context.Background(), tenantId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProductConfiguratorApiService ProductConfiguratorCreateOption", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string
        var stepId string

        resp, httpRes, err := apiClient.ProductConfiguratorApi.ProductConfiguratorCreateOption(context.Background(), tenantId, stepId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProductConfiguratorApiService ProductConfiguratorCreateProperty", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string

        resp, httpRes, err := apiClient.ProductConfiguratorApi.ProductConfiguratorCreateProperty(context.Background(), tenantId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProductConfiguratorApiService ProductConfiguratorCreateStep", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string
        var configuratorId string

        resp, httpRes, err := apiClient.ProductConfiguratorApi.ProductConfiguratorCreateStep(context.Background(), tenantId, configuratorId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProductConfiguratorApiService ProductConfiguratorDeleteConfigurator", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string
        var configuratorId string

        resp, httpRes, err := apiClient.ProductConfiguratorApi.ProductConfiguratorDeleteConfigurator(context.Background(), tenantId, configuratorId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProductConfiguratorApiService ProductConfiguratorDeleteDependency", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string
        var dependencyId string

        resp, httpRes, err := apiClient.ProductConfiguratorApi.ProductConfiguratorDeleteDependency(context.Background(), tenantId, dependencyId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProductConfiguratorApiService ProductConfiguratorDeleteMatrix", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string
        var matrixId string

        resp, httpRes, err := apiClient.ProductConfiguratorApi.ProductConfiguratorDeleteMatrix(context.Background(), tenantId, matrixId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProductConfiguratorApiService ProductConfiguratorDeleteOption", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string
        var optionId string

        resp, httpRes, err := apiClient.ProductConfiguratorApi.ProductConfiguratorDeleteOption(context.Background(), tenantId, optionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProductConfiguratorApiService ProductConfiguratorDeleteStep", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string
        var stepId string

        resp, httpRes, err := apiClient.ProductConfiguratorApi.ProductConfiguratorDeleteStep(context.Background(), tenantId, stepId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProductConfiguratorApiService ProductConfiguratorGetAvailableConfiguration", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string
        var productId string

        resp, httpRes, err := apiClient.ProductConfiguratorApi.ProductConfiguratorGetAvailableConfiguration(context.Background(), tenantId, productId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProductConfiguratorApiService ProductConfiguratorGetAvailableConfiguration2", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string
        var productId string

        resp, httpRes, err := apiClient.ProductConfiguratorApi.ProductConfiguratorGetAvailableConfiguration2(context.Background(), tenantId, productId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProductConfiguratorApiService ProductConfiguratorGetConfiguratorByProductId", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string
        var productId string

        resp, httpRes, err := apiClient.ProductConfiguratorApi.ProductConfiguratorGetConfiguratorByProductId(context.Background(), tenantId, productId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProductConfiguratorApiService ProductConfiguratorGetConfiguratorByProductId2", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string
        var productId string
        var status string

        resp, httpRes, err := apiClient.ProductConfiguratorApi.ProductConfiguratorGetConfiguratorByProductId2(context.Background(), tenantId, productId, status).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProductConfiguratorApiService ProductConfiguratorGetMatrix", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string
        var matrixId string

        resp, httpRes, err := apiClient.ProductConfiguratorApi.ProductConfiguratorGetMatrix(context.Background(), tenantId, matrixId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProductConfiguratorApiService ProductConfiguratorListConfigurators", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string
        var productId string
        var pageSize int64

        resp, httpRes, err := apiClient.ProductConfiguratorApi.ProductConfiguratorListConfigurators(context.Background(), tenantId, productId, pageSize).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProductConfiguratorApiService ProductConfiguratorListDependencies", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string
        var pageSize int64

        resp, httpRes, err := apiClient.ProductConfiguratorApi.ProductConfiguratorListDependencies(context.Background(), tenantId, pageSize).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProductConfiguratorApiService ProductConfiguratorListMatrices", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string
        var configuratorId string
        var pageSize string

        resp, httpRes, err := apiClient.ProductConfiguratorApi.ProductConfiguratorListMatrices(context.Background(), tenantId, configuratorId, pageSize).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProductConfiguratorApiService ProductConfiguratorListOptions", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string
        var stepId string
        var pageSize int64

        resp, httpRes, err := apiClient.ProductConfiguratorApi.ProductConfiguratorListOptions(context.Background(), tenantId, stepId, pageSize).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProductConfiguratorApiService ProductConfiguratorListProperties", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string
        var matrixId string
        var pageSize string

        resp, httpRes, err := apiClient.ProductConfiguratorApi.ProductConfiguratorListProperties(context.Background(), tenantId, matrixId, pageSize).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProductConfiguratorApiService ProductConfiguratorUpdateConfigurator2", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string
        var configuratorId string

        resp, httpRes, err := apiClient.ProductConfiguratorApi.ProductConfiguratorUpdateConfigurator2(context.Background(), tenantId, configuratorId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProductConfiguratorApiService ProductConfiguratorUpdateDependency", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string
        var dependencyId string

        resp, httpRes, err := apiClient.ProductConfiguratorApi.ProductConfiguratorUpdateDependency(context.Background(), tenantId, dependencyId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProductConfiguratorApiService ProductConfiguratorUpdateMatrix", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string
        var matrixId string

        resp, httpRes, err := apiClient.ProductConfiguratorApi.ProductConfiguratorUpdateMatrix(context.Background(), tenantId, matrixId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProductConfiguratorApiService ProductConfiguratorUpdateOption", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string
        var optionId string

        resp, httpRes, err := apiClient.ProductConfiguratorApi.ProductConfiguratorUpdateOption(context.Background(), tenantId, optionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProductConfiguratorApiService ProductConfiguratorUpdateProperty", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string
        var propertyId string

        resp, httpRes, err := apiClient.ProductConfiguratorApi.ProductConfiguratorUpdateProperty(context.Background(), tenantId, propertyId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProductConfiguratorApiService ProductConfiguratorUpdateStep", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string
        var stepId string

        resp, httpRes, err := apiClient.ProductConfiguratorApi.ProductConfiguratorUpdateStep(context.Background(), tenantId, stepId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}