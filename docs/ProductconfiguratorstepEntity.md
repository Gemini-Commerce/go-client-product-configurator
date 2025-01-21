# # ProductconfiguratorstepEntity
Represents a step within a configurator, including its unique identifier, localized label, description, and associated options.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **string** | The unique identifier for the step.  | [optional]
**Grn**| **string** |   | [optional]
**Label**| [**LocalisationLocalizedText**](LocalisationLocalizedText.md) |   | [optional]
**Description**| [**LocalisationLocalizedText**](LocalisationLocalizedText.md) |   | [optional]
**SubjectToStepId**| **string** |   | [optional]
**Position**| **string** |   | [optional]
**IsRequired**| **bool** |   | [optional]
**Options**| [**[]ProductconfiguratoroptionEntity**](ProductconfiguratoroptionEntity.md) |   | [optional]
**HasMultipleSelection**| **bool** |   | [optional]
**OptionsHaveQuantity**| **bool** |   | [optional]
**CreatedAt**| [**time.Time**](time.Time.md) |   | [optional]
**UpdatedAt**| [**time.Time**](time.Time.md) |   | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

