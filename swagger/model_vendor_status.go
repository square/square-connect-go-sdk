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

// VendorStatus : The status of the [Vendor](entity:Vendor), whether a [Vendor](entity:Vendor) is active or inactive.
type VendorStatus string

// List of VendorStatus
const (
	DO_NOT_USE_VendorStatus VendorStatus = "DO_NOT_USE"
	ACTIVE_VendorStatus     VendorStatus = "ACTIVE"
	INACTIVE_VendorStatus   VendorStatus = "INACTIVE"
)
