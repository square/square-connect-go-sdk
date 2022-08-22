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

// An object that contains the gift card and customer ID associated with a  `gift_card.customer_linked` webhook event.
type GiftCardCustomerLinkedWebhookObject struct {
	GiftCard *GiftCard `json:"gift_card,omitempty"`
	// The ID of the linked [customer](entity:Customer).
	LinkedCustomerId string `json:"linked_customer_id,omitempty"`
}
