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

// Defines the fields that are included in the response body of a request to the [ListLocations](api-endpoint:Locations-ListLocations) endpoint.  Either `errors` or `locations` is present in a given response (never both).
type ListLocationsResponse struct {
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
	// The business locations.
	Locations []Location `json:"locations,omitempty"`
}
