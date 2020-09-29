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

// OrderLineItemTaxType : Indicates how the tax is applied to the associated line item or order.
type OrderLineItemTaxType string

// List of OrderLineItemTaxType
const (
	UNKNOWN_TAX_OrderLineItemTaxType OrderLineItemTaxType = "UNKNOWN_TAX"
	ADDITIVE_OrderLineItemTaxType    OrderLineItemTaxType = "ADDITIVE"
	INCLUSIVE_OrderLineItemTaxType   OrderLineItemTaxType = "INCLUSIVE"
)