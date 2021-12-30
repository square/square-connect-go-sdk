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

// SubscriptionStatus : Supported subscription statuses.
type SubscriptionStatus string

// List of SubscriptionStatus
const (
	DEFAULT_SUBSCRIPTION_STATUS_DO_NOT_USE_SubscriptionStatus SubscriptionStatus = "DEFAULT_SUBSCRIPTION_STATUS_DO_NOT_USE"
	PENDING_SubscriptionStatus                                SubscriptionStatus = "PENDING"
	ACTIVE_SubscriptionStatus                                 SubscriptionStatus = "ACTIVE"
	CANCELED_SubscriptionStatus                               SubscriptionStatus = "CANCELED"
	DEACTIVATED_SubscriptionStatus                            SubscriptionStatus = "DEACTIVATED"
	PAUSED_SubscriptionStatus                                 SubscriptionStatus = "PAUSED"
)
