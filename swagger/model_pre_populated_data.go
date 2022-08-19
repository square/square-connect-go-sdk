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

// Describes buyer data to prepopulate in the payment form.  For more information,  see [Optional Checkout Configurations](https://developer.squareup.com/docs/checkout-api/optional-checkout-configurations).
type PrePopulatedData struct {
	// The buyer email to prepopulate in the payment form.
	BuyerEmail string `json:"buyer_email,omitempty"`
	// The buyer phone number to prepopulate in the payment form.
	BuyerPhoneNumber string   `json:"buyer_phone_number,omitempty"`
	BuyerAddress     *Address `json:"buyer_address,omitempty"`
}
