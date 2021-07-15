/*
 * Square
 *
 * Use Square APIs to manage and run business including payment, customer, product, inventory, and employee management.
 *
 * API version: 2.0
 * Contact: developers@squareup.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// SearchOrdersSortField : Specifies which timestamp to use to sort `SearchOrder` results.
type SearchOrdersSortField string

// List of SearchOrdersSortField
const (
	DO_NOT_USE_SearchOrdersSortField SearchOrdersSortField = "DO_NOT_USE"
	CREATED_AT_SearchOrdersSortField SearchOrdersSortField = "CREATED_AT"
	UPDATED_AT_SearchOrdersSortField SearchOrdersSortField = "UPDATED_AT"
	CLOSED_AT_SearchOrdersSortField  SearchOrdersSortField = "CLOSED_AT"
)
