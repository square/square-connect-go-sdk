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

// Published when a customer [custom attribute definition](entity:CustomAttributeDefinition) that is visible to the subscribing app is created.
type CustomerCustomAttributeDefinitionCreatedPublicWebhook struct {
	// The ID of the target seller associated with the event.
	MerchantId string `json:"merchant_id,omitempty"`
	// The type of this event. The value is `\"customer.custom_attribute_definition.public.created\"`.
	Type_ string `json:"type,omitempty"`
	// A unique ID for the webhook event.
	EventId string `json:"event_id,omitempty"`
	// The timestamp of when the webhook event was created, in RFC 3339 format.
	CreatedAt string                                `json:"created_at,omitempty"`
	Data      *CustomAttributeDefinitionWebhookData `json:"data,omitempty"`
}
