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

// Represents a contact of a [Vendor](entity:Vendor).
type VendorContact struct {
	// A unique Square-generated ID for the [VendorContact](entity:VendorContact). This field is required when attempting to update a [VendorContact](entity:VendorContact).
	Id string `json:"id,omitempty"`
	// The name of the [VendorContact](entity:VendorContact). This field is required when attempting to create a [Vendor](entity:Vendor).
	Name string `json:"name,omitempty"`
	// The email address of the [VendorContact](entity:VendorContact).
	EmailAddress string `json:"email_address,omitempty"`
	// The phone number of the [VendorContact](entity:VendorContact).
	PhoneNumber string `json:"phone_number,omitempty"`
	// The state of the [VendorContact](entity:VendorContact).
	Removed bool `json:"removed,omitempty"`
	// The ordinal of the [VendorContact](entity:VendorContact).
	Ordinal int32 `json:"ordinal"`
}
