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

// TerminalActionActionType : Describes the type of this unit and indicates which field contains the unit information. This is an ‘open’ enum.
type TerminalActionActionType string

// List of TerminalActionActionType
const (
	INVALID_TYPE_TerminalActionActionType            TerminalActionActionType = "INVALID_TYPE"
	UNSUPPORTED_ACTION_TYPE_TerminalActionActionType TerminalActionActionType = "UNSUPPORTED_ACTION_TYPE"
	CHECKOUT_TerminalActionActionType                TerminalActionActionType = "CHECKOUT"
	REFUND_TerminalActionActionType                  TerminalActionActionType = "REFUND"
	QR_CODE_TerminalActionActionType                 TerminalActionActionType = "QR_CODE"
	PING_TerminalActionActionType                    TerminalActionActionType = "PING"
	SAVE_CARD_TerminalActionActionType               TerminalActionActionType = "SAVE_CARD"
	SIGNATURE_TerminalActionActionType               TerminalActionActionType = "SIGNATURE"
	CONFIRMATION_TerminalActionActionType            TerminalActionActionType = "CONFIRMATION"
	RECEIPT_TerminalActionActionType                 TerminalActionActionType = "RECEIPT"
	DATA_COLLECTION_TerminalActionActionType         TerminalActionActionType = "DATA_COLLECTION"
	SELECT__TerminalActionActionType                 TerminalActionActionType = "SELECT"
)
