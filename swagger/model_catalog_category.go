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

// A category to which a `CatalogItem` instance belongs.
type CatalogCategory struct {
	// The category name. This is a searchable attribute for use in applicable query filters, and its value length is of Unicode code points.
	Name string `json:"name,omitempty"`
	// The IDs of images associated with this `CatalogCategory` instance. Currently these images are not displayed by Square, but are free to be displayed in 3rd party applications.
	ImageIds []string `json:"image_ids,omitempty"`
}
