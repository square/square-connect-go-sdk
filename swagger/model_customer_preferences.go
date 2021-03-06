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

// Represents communication preferences for the customer profile.
type CustomerPreferences struct {
	// The customer has unsubscribed from receiving marketing campaign emails.
	EmailUnsubscribed bool `json:"email_unsubscribed,omitempty"`
}
