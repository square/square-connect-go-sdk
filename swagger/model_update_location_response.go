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

// The response object returned by the [UpdateLocation](api-endpoint:Locations-UpdateLocation) endpoint.
type UpdateLocationResponse struct {
	// Information about errors encountered during the request.
	Errors   []ModelError `json:"errors,omitempty"`
	Location *Location    `json:"location,omitempty"`
}
