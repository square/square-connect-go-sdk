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

// A record representing an individual team member for a business.
type TeamMember struct {
	// The unique ID for the team member.
	Id string `json:"id,omitempty"`
	// A second ID used to associate the team member with an entity in another system.
	ReferenceId string `json:"reference_id,omitempty"`
	// Whether the team member is the owner of the Square account.
	IsOwner bool              `json:"is_owner,omitempty"`
	Status  *TeamMemberStatus `json:"status,omitempty"`
	// The given name (that is, the first name) associated with the team member.
	GivenName string `json:"given_name,omitempty"`
	// The family name (that is, the last name) associated with the team member.
	FamilyName string `json:"family_name,omitempty"`
	// The email address associated with the team member.
	EmailAddress string `json:"email_address,omitempty"`
	// The team member's phone number, in E.164 format. For example: +14155552671 - the country code is 1 for US +551155256325 - the country code is 55 for BR
	PhoneNumber string `json:"phone_number,omitempty"`
	// The timestamp, in RFC 3339 format, describing when the team member was created. For example, \"2018-10-04T04:00:00-07:00\" or \"2019-02-05T12:00:00Z\".
	CreatedAt string `json:"created_at,omitempty"`
	// The timestamp, in RFC 3339 format, describing when the team member was last updated. For example, \"2018-10-04T04:00:00-07:00\" or \"2019-02-05T12:00:00Z\".
	UpdatedAt         string                       `json:"updated_at,omitempty"`
	AssignedLocations *TeamMemberAssignedLocations `json:"assigned_locations,omitempty"`
}
