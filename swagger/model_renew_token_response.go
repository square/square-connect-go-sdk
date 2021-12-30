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

type RenewTokenResponse struct {
	// The renewed access token. This value might be different from the `access_token` you provided in your request. You provide this token in a header with every request to Connect API endpoints. See [Request and response headers](https://developer.squareup.com/docs/api/connect/v2/#requestandresponseheaders) for the format of this header.
	AccessToken string `json:"access_token,omitempty"`
	// This value is always _bearer_.
	TokenType string `json:"token_type,omitempty"`
	// The date when the `access_token` expires, in [ISO 8601](http://www.iso.org/iso/home/standards/iso8601.htm) format.
	ExpiresAt string `json:"expires_at,omitempty"`
	// The ID of the authorizing merchant's business.
	MerchantId string `json:"merchant_id,omitempty"`
	// __LEGACY FIELD__. The ID of the merchant subscription associated with the authorization. The ID is only present if the merchant signed up for a subscription during authorization.
	SubscriptionId string `json:"subscription_id,omitempty"`
	// __LEGACY FIELD__. The ID of the subscription plan the merchant signed up for. The ID is only present if the merchant signed up for a subscription plan during authorization.
	PlanId string `json:"plan_id,omitempty"`
}
