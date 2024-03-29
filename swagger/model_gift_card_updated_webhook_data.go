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

// The data associated with a `gift_card.updated` webhook event.
type GiftCardUpdatedWebhookData struct {
	// The type of object affected by the event. For this event, the value is `gift_card`.
	Type_ string `json:"type,omitempty"`
	// The ID of the updated gift card.
	Id     string                        `json:"id,omitempty"`
	Object *GiftCardUpdatedWebhookObject `json:"object,omitempty"`
}
