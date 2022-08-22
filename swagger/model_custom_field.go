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

// Describes a custom form field to add to the checkout page to collect more information from buyers during checkout. For more information, see [Specify checkout options](https://developer.squareup.com/docs/checkout-api/optional-checkout-configurations#specify-checkout-options-1).
type CustomField struct {
	// The title of the custom field.
	Title string `json:"title"`
}
