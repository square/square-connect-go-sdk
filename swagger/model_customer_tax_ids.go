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

// Represents the tax ID associated with a [customer profile](entity:Customer). The corresponding `tax_ids` field is available only for customers of sellers in EU countries or the United Kingdom.  For more information, see [Customer tax IDs](https://developer.squareup.com/docs/customers-api/what-it-does#customer-tax-ids).
type CustomerTaxIds struct {
	// The EU VAT identification number for the customer. For example, `IE3426675K`. The ID can contain alphanumeric characters only.
	EuVat string `json:"eu_vat,omitempty"`
}
