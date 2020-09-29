# V1EmployeeRole

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The role&#x27;s unique ID, Can only be set by Square. | [optional] [default to null]
**Name** | **string** | The role&#x27;s merchant-defined name. | [default to null]
**Permissions** | [**[]V1EmployeeRolePermissions**](V1EmployeeRolePermissions.md) | The role&#x27;s permissions. See [V1EmployeeRolePermissions](#type-v1employeerolepermissions) for possible values | [default to null]
**IsOwner** | **bool** | If true, employees with this role have all permissions, regardless of the values indicated in permissions. | [optional] [default to null]
**CreatedAt** | **string** | The time when the employee entity was created, in ISO 8601 format. Is set by Square when the Role is created. | [optional] [default to null]
**UpdatedAt** | **string** | The time when the employee entity was most recently updated, in ISO 8601 format. Is set by Square when the Role updated. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

