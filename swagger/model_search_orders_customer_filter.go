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

// Filter based on Order `customer_id` and any Tender `customer_id` associated with the Order. Does not filter based on the [FulfillmentRecipient](#type-orderfulfillmentrecipient) `customer_id`.
type SearchOrdersCustomerFilter struct {
	// List of customer IDs to filter by.  Max: 10 customer IDs.
	CustomerIds []string `json:"customer_ids,omitempty"`
}
