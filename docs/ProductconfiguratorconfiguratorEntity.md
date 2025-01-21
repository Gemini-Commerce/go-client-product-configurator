# # ProductconfiguratorconfiguratorEntity
Represents a configurator entity, including its unique identifier, global resource name (GRN), product association, and a list of configuration steps.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **string** | The unique identifier for the configurator.  | [optional]
**Grn**| **string** | The global resource name associated with the configurator.  | [optional]
**ProductId**| **string** | The identifier of the product linked to the configurator.  | [optional]
**Label**| **string** | A descriptive label for the configurator.  | [optional]
**Status**| [**ProductconfiguratorconfiguratorStatus**](ProductconfiguratorconfiguratorStatus.md) |  for more information please, see Model/ProductconfiguratorconfiguratorStatus.php  | [optional] [default to PRODUCTCONFIGURATORCONFIGURATORSTATUS_UNKNOWN]
**Steps**| [**[]ProductconfiguratorstepEntity**](ProductconfiguratorstepEntity.md) | A list of steps that define the configurator&#39;s structure.  | [optional]
**CreatedAt**| [**time.Time**](time.Time.md) |   | [optional]
**UpdatedAt**| [**time.Time**](time.Time.md) |   | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

