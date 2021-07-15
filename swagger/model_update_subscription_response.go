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

// Defines the fields that are included in the response from the [UpdateSubscription](api-endpoint:Subscriptions-UpdateSubscription) endpoint.
type UpdateSubscriptionResponse struct {
	// Information about errors encountered during the request.
	Errors       []ModelError  `json:"errors,omitempty"`
	Subscription *Subscription `json:"subscription,omitempty"`
}
