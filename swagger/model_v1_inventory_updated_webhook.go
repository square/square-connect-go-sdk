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

// Published when the inventory quantity for a catalog item is updated.
type V1InventoryUpdatedWebhook struct {
	// The ID of the target merchant associated with the event.
	MerchantId string `json:"merchant_id,omitempty"`
	// The ID of the target merchant associated with the event.
	LocationId string `json:"location_id,omitempty"`
	// The type of event this represents, i.e. `INVENTORY_UPDATED`.
	EventType string `json:"event_type,omitempty"`
	// The ID of the V1 Item whose inventory was updated.
	EntityId string `json:"entity_id,omitempty"`
}
