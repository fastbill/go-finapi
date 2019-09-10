# Category

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Category identifier.&lt;br/&gt;&lt;br/&gt;NOTE: Do NOT assume that the identifiers of the global finAPI categories are the same across different finAPI environments. In fact, the identifiers may change whenever a new finAPI version is released, even within the same environment. The identifiers are meant to be used for references within the finAPI services only, but not for hard-coding them in your application. If you need to hard-code the usage of a certain global category within your application, please instead refer to the category name. Also, please make sure to check the &#39;isCustom&#39; flag, which is false for all global categories (if you are not regarding this flag, you might end up referring to a user-specific category, and not the global category). | [default to null]
**Name** | **string** | Category name | [default to null]
**ParentId** | **int64** | Identifier of the parent category (if a parent category exists) | [optional] [default to null]
**ParentName** | **string** | Name of the parent category (if a parent category exists) | [optional] [default to null]
**IsCustom** | **bool** | Whether the category is a finAPI global category (in which case this field will be false), or the category was created by a user (in which case this field will be true) | [default to null]
**Children** | **[]int64** | List of sub-categories identifiers (if any exist) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


