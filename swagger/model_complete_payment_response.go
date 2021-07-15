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

// Defines the response returned by[CompletePayment](api-endpoint:Payments-CompletePayment).
type CompletePaymentResponse struct {
	// Information about errors encountered during the request.
	Errors  []ModelError `json:"errors,omitempty"`
	Payment *Payment     `json:"payment,omitempty"`
}
