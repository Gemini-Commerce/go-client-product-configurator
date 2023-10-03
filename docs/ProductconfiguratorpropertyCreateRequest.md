# # ProductconfiguratorpropertyCreateRequest


## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId**| **string** |   | [optional]
**MatrixId**| **string** |   | [optional]
**OptionIds**| **[]string** |   | [optional]
**GenericProperty**| [**ProductconfiguratorpropertyGenericProperty**](ProductconfiguratorpropertyGenericProperty.md) |   | [optional]
**PriceProperty**| **map[string]interface{}** |   | [optional]
**Coordinates**| **[]int64** | coordinates of the entity in the matrix. The order matters. Example: [1, 2] means that the entity is located at the first row, second column. Example: [1, 2, 3] means that the entity is located at the first row, second column, third layer. Example: [1, 2, 3, 4] means that the entity is located at the first row, second column, third layer, fourth depth.  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

