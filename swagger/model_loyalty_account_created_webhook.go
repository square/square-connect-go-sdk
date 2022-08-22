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

// Published when a [loyalty account](entity:LoyaltyAccount) is created.
type LoyaltyAccountCreatedWebhook struct {
	// The ID of the Square seller associated with the event.
	MerchantId string `json:"merchant_id,omitempty"`
	// The type of event. For this event, the value is `loyalty.account.created`.
	Type_ string `json:"type,omitempty"`
	// The unique ID for the event, which is used for [idempotency support](https://developer.squareup.com/docs/webhooks/step4manage#webhooks-best-practices).
	EventId string `json:"event_id,omitempty"`
	// The timestamp of when the webhook event was created, in RFC 3339 format.
	CreatedAt string                            `json:"created_at,omitempty"`
	Data      *LoyaltyAccountCreatedWebhookData `json:"data,omitempty"`
}
