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

// A query filter to search for buyer-accessible appointment segments by.
type SegmentFilter struct {
	// The ID of the [CatalogItemVariation](entity:CatalogItemVariation) object representing the service booked in this segment.
	ServiceVariationId string       `json:"service_variation_id"`
	TeamMemberIdFilter *FilterValue `json:"team_member_id_filter,omitempty"`
}
