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

// Describes a `UpdateInvoice` request.
type UpdateInvoiceRequest struct {
	Invoice *Invoice `json:"invoice"`
	// A unique string that identifies the `UpdateInvoice` request. If you do not provide `idempotency_key` (or provide an empty string as the value), the endpoint treats each request as independent.  For more information, see [Idempotency](https://developer.squareup.com/docs/working-with-apis/idempotency).
	IdempotencyKey string `json:"idempotency_key,omitempty"`
	// The list of fields to clear. For examples, see [Update an Invoice](https://developer.squareup.com/docs/invoices-api/update-invoices).
	FieldsToClear []string `json:"fields_to_clear,omitempty"`
}
