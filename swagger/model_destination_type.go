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

// DestinationType : List of possible destinations against which a payout can be made.
type DestinationType string

// List of DestinationType
const (
	UNKNOWN_DESTINATION_TYPE_DO_NOT_USE_DestinationType DestinationType = "UNKNOWN_DESTINATION_TYPE_DO_NOT_USE"
	BANK_ACCOUNT_DestinationType                        DestinationType = "BANK_ACCOUNT"
	CARD_DestinationType                                DestinationType = "CARD"
	SQUARE_BALANCE_DestinationType                      DestinationType = "SQUARE_BALANCE"
	SQUARE_STORED_BALANCE_DestinationType               DestinationType = "SQUARE_STORED_BALANCE"
)
