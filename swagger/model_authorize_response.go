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

type AuthorizeResponse struct {
	// A valid authorization code. Authorization codes are exchanged for OAuth access tokens with the `ObtainToken` endpoint.
	Code string `json:"code,omitempty"`
	// The same value specified in the request.
	State string `json:"state,omitempty"`
}
