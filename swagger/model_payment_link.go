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

type PaymentLink struct {
	// The Square-assigned ID of the payment link.
	Id string `json:"id,omitempty"`
	// The Square-assigned version number, which is incremented each time an update is committed to the payment link.
	Version int32 `json:"version"`
	// The optional description of the `payment_link` object.  It is primarily for use by your application and is not used anywhere.
	Description string `json:"description,omitempty"`
	// The ID of the order associated with the payment link.
	OrderId          string            `json:"order_id,omitempty"`
	CheckoutOptions  *CheckoutOptions  `json:"checkout_options,omitempty"`
	PrePopulatedData *PrePopulatedData `json:"pre_populated_data,omitempty"`
	// The URL of the payment link.
	Url string `json:"url,omitempty"`
	// The timestamp when the payment link was created, in RFC 3339 format.
	CreatedAt string `json:"created_at,omitempty"`
	// The timestamp when the payment link was last updated, in RFC 3339 format.
	UpdatedAt string `json:"updated_at,omitempty"`
	// An optional note. After Square processes the payment, this note is added to the   resulting `Payment`.
	PaymentNote string `json:"payment_note,omitempty"`
}
