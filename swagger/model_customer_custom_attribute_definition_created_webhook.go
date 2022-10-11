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

// Published when a customer [custom attribute definition](entity:CustomAttributeDefinition)  is created by the subscribing application.  This event is replaced by  [customer.custom_attribute_definition.owned.created](webhook:customer.custom_attribute_definition.owned.created).
type CustomerCustomAttributeDefinitionCreatedWebhook struct {
	// The ID of the seller associated with the event that triggered the event notification.
	MerchantId string `json:"merchant_id,omitempty"`
	// The type of this event. The value is `\"customer.custom_attribute_definition.created\"`.
	Type_ string `json:"type,omitempty"`
	// A unique ID for the event notification.
	EventId string `json:"event_id,omitempty"`
	// The timestamp that indicates when the event notification was created, in RFC 3339 format.
	CreatedAt string                                `json:"created_at,omitempty"`
	Data      *CustomAttributeDefinitionWebhookData `json:"data,omitempty"`
}
