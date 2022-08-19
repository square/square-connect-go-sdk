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

type PaymentBalanceActivityAutomaticSavingsReversedDetail struct {
	// The ID of the payment associated with this activity.
	PaymentId string `json:"payment_id,omitempty"`
	// The ID of the payout associated with this activity.
	PayoutId string `json:"payout_id,omitempty"`
}
