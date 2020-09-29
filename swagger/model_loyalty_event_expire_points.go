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

// Provides metadata when the event `type` is `EXPIRE_POINTS`.
type LoyaltyEventExpirePoints struct {
	// The Square-assigned ID of the [loyalty program](#type-LoyaltyProgram).
	LoyaltyProgramId string `json:"loyalty_program_id"`
	// The number of points expired.
	Points int32 `json:"points"`
}