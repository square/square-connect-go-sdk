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

// CatalogItemProductType : The type of a CatalogItem. Connect V2 only allows the creation of `REGULAR` or `APPOINTMENTS_SERVICE` items.
type CatalogItemProductType string

// List of CatalogItemProductType
const (
	REGULAR_CatalogItemProductType              CatalogItemProductType = "REGULAR"
	GIFT_CARD_CatalogItemProductType            CatalogItemProductType = "GIFT_CARD"
	APPOINTMENTS_SERVICE_CatalogItemProductType CatalogItemProductType = "APPOINTMENTS_SERVICE"
	CLASS_TICKET_CatalogItemProductType         CatalogItemProductType = "CLASS_TICKET"
)
