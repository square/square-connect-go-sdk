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

// Represents the quantity of an item variation that is physically present at a specific location, verified by a seller or a seller's employee. For example, a physical count might come from an employee counting the item variations on hand or from syncing with an external system.
type InventoryPhysicalCount struct {
	// A unique ID generated by Square for the [InventoryPhysicalCount](#type-inventoryphysicalcount).
	Id string `json:"id,omitempty"`
	// An optional ID provided by the application to tie the [InventoryPhysicalCount](#type-inventoryphysicalcount) to an external system.
	ReferenceId string `json:"reference_id,omitempty"`
	// The Square generated ID of the `CatalogObject` being tracked.
	CatalogObjectId string `json:"catalog_object_id,omitempty"`
	// The `CatalogObjectType` of the `CatalogObject` being tracked. Tracking is only supported for the `ITEM_VARIATION` type.
	CatalogObjectType string          `json:"catalog_object_type,omitempty"`
	Status            *InventoryState `json:"status,omitempty"`
	State             *InventoryState `json:"state,omitempty"`
	// The Square ID of the [Location](#type-location) where the related quantity of items are being tracked.
	LocationId string `json:"location_id,omitempty"`
	// The number of items affected by the physical count as a decimal string. Can support up to 5 digits after the decimal point.
	Quantity string             `json:"quantity,omitempty"`
	Source   *SourceApplication `json:"source,omitempty"`
	// The Square ID of the [Employee](#type-employee) responsible for the physical count.
	EmployeeId string `json:"employee_id,omitempty"`
	// A client-generated timestamp in RFC 3339 format that indicates when the physical count took place. For write actions, the `occurred_at` timestamp cannot be older than 24 hours or in the future relative to the time of the request.
	OccurredAt string `json:"occurred_at,omitempty"`
	// A read-only timestamp in RFC 3339 format that indicates when Square received the physical count.
	CreatedAt string `json:"created_at,omitempty"`
}
