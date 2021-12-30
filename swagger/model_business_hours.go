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

// The hours of operation for a location.
type BusinessHours struct {
	// The list of time periods during which the business is open. There may be at most 10 periods per day.
	Periods []BusinessHoursPeriod `json:"periods,omitempty"`
}
