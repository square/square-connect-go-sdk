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

// Defines output parameters in a response from the  [CancelSubscription](api-endpoint:Subscriptions-CancelSubscription) endpoint.
type CancelSubscriptionResponse struct {
	// Errors encountered during the request.
	Errors       []ModelError  `json:"errors,omitempty"`
	Subscription *Subscription `json:"subscription,omitempty"`
	// A list of a single `CANCEL` action scheduled for the subscription.
	Actions []SubscriptionAction `json:"actions,omitempty"`
}
