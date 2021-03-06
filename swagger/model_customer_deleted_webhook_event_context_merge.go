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

// Information about a merge operation, which creates a new customer using aggregated properties from two or more existing customers.
type CustomerDeletedWebhookEventContextMerge struct {
	// The IDs of the existing customers that were merged and then deleted.
	FromCustomerIds []string `json:"from_customer_ids,omitempty"`
	// The ID of the new customer created by the merge.
	ToCustomerId string `json:"to_customer_id,omitempty"`
}
