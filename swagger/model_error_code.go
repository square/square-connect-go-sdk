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

// ErrorCode : Indicates the specific error that occurred during a request to a Square API.
type ErrorCode string

// List of ErrorCode
const (
	INTERNAL_SERVER_ERROR_ErrorCode                                             ErrorCode = "INTERNAL_SERVER_ERROR"
	UNAUTHORIZED_ErrorCode                                                      ErrorCode = "UNAUTHORIZED"
	ACCESS_TOKEN_EXPIRED_ErrorCode                                              ErrorCode = "ACCESS_TOKEN_EXPIRED"
	ACCESS_TOKEN_REVOKED_ErrorCode                                              ErrorCode = "ACCESS_TOKEN_REVOKED"
	CLIENT_DISABLED_ErrorCode                                                   ErrorCode = "CLIENT_DISABLED"
	FORBIDDEN_ErrorCode                                                         ErrorCode = "FORBIDDEN"
	INSUFFICIENT_SCOPES_ErrorCode                                               ErrorCode = "INSUFFICIENT_SCOPES"
	APPLICATION_DISABLED_ErrorCode                                              ErrorCode = "APPLICATION_DISABLED"
	V1_APPLICATION_ErrorCode                                                    ErrorCode = "V1_APPLICATION"
	V1_ACCESS_TOKEN_ErrorCode                                                   ErrorCode = "V1_ACCESS_TOKEN"
	CARD_PROCESSING_NOT_ENABLED_ErrorCode                                       ErrorCode = "CARD_PROCESSING_NOT_ENABLED"
	MERCHANT_SUBSCRIPTION_NOT_FOUND_ErrorCode                                   ErrorCode = "MERCHANT_SUBSCRIPTION_NOT_FOUND"
	BAD_REQUEST_ErrorCode                                                       ErrorCode = "BAD_REQUEST"
	MISSING_REQUIRED_PARAMETER_ErrorCode                                        ErrorCode = "MISSING_REQUIRED_PARAMETER"
	INCORRECT_TYPE_ErrorCode                                                    ErrorCode = "INCORRECT_TYPE"
	INVALID_TIME_ErrorCode                                                      ErrorCode = "INVALID_TIME"
	INVALID_TIME_RANGE_ErrorCode                                                ErrorCode = "INVALID_TIME_RANGE"
	INVALID_VALUE_ErrorCode                                                     ErrorCode = "INVALID_VALUE"
	INVALID_CURSOR_ErrorCode                                                    ErrorCode = "INVALID_CURSOR"
	UNKNOWN_QUERY_PARAMETER_ErrorCode                                           ErrorCode = "UNKNOWN_QUERY_PARAMETER"
	CONFLICTING_PARAMETERS_ErrorCode                                            ErrorCode = "CONFLICTING_PARAMETERS"
	EXPECTED_JSON_BODY_ErrorCode                                                ErrorCode = "EXPECTED_JSON_BODY"
	INVALID_SORT_ORDER_ErrorCode                                                ErrorCode = "INVALID_SORT_ORDER"
	VALUE_REGEX_MISMATCH_ErrorCode                                              ErrorCode = "VALUE_REGEX_MISMATCH"
	VALUE_TOO_SHORT_ErrorCode                                                   ErrorCode = "VALUE_TOO_SHORT"
	VALUE_TOO_LONG_ErrorCode                                                    ErrorCode = "VALUE_TOO_LONG"
	VALUE_TOO_LOW_ErrorCode                                                     ErrorCode = "VALUE_TOO_LOW"
	VALUE_TOO_HIGH_ErrorCode                                                    ErrorCode = "VALUE_TOO_HIGH"
	VALUE_EMPTY_ErrorCode                                                       ErrorCode = "VALUE_EMPTY"
	ARRAY_LENGTH_TOO_LONG_ErrorCode                                             ErrorCode = "ARRAY_LENGTH_TOO_LONG"
	ARRAY_LENGTH_TOO_SHORT_ErrorCode                                            ErrorCode = "ARRAY_LENGTH_TOO_SHORT"
	ARRAY_EMPTY_ErrorCode                                                       ErrorCode = "ARRAY_EMPTY"
	EXPECTED_BOOLEAN_ErrorCode                                                  ErrorCode = "EXPECTED_BOOLEAN"
	EXPECTED_INTEGER_ErrorCode                                                  ErrorCode = "EXPECTED_INTEGER"
	EXPECTED_FLOAT_ErrorCode                                                    ErrorCode = "EXPECTED_FLOAT"
	EXPECTED_STRING_ErrorCode                                                   ErrorCode = "EXPECTED_STRING"
	EXPECTED_OBJECT_ErrorCode                                                   ErrorCode = "EXPECTED_OBJECT"
	EXPECTED_ARRAY_ErrorCode                                                    ErrorCode = "EXPECTED_ARRAY"
	EXPECTED_MAP_ErrorCode                                                      ErrorCode = "EXPECTED_MAP"
	EXPECTED_BASE64_ENCODED_BYTE_ARRAY_ErrorCode                                ErrorCode = "EXPECTED_BASE64_ENCODED_BYTE_ARRAY"
	INVALID_ARRAY_VALUE_ErrorCode                                               ErrorCode = "INVALID_ARRAY_VALUE"
	INVALID_ENUM_VALUE_ErrorCode                                                ErrorCode = "INVALID_ENUM_VALUE"
	INVALID_CONTENT_TYPE_ErrorCode                                              ErrorCode = "INVALID_CONTENT_TYPE"
	INVALID_FORM_VALUE_ErrorCode                                                ErrorCode = "INVALID_FORM_VALUE"
	CUSTOMER_NOT_FOUND_ErrorCode                                                ErrorCode = "CUSTOMER_NOT_FOUND"
	BUYER_NOT_FOUND_ErrorCode                                                   ErrorCode = "BUYER_NOT_FOUND"
	ONE_INSTRUMENT_EXPECTED_ErrorCode                                           ErrorCode = "ONE_INSTRUMENT_EXPECTED"
	NO_FIELDS_SET_ErrorCode                                                     ErrorCode = "NO_FIELDS_SET"
	DEPRECATED_FIELD_SET_ErrorCode                                              ErrorCode = "DEPRECATED_FIELD_SET"
	RETIRED_FIELD_SET_ErrorCode                                                 ErrorCode = "RETIRED_FIELD_SET"
	TOO_MANY_MAP_ENTRIES_ErrorCode                                              ErrorCode = "TOO_MANY_MAP_ENTRIES"
	MAP_KEY_LENGTH_TOO_SHORT_ErrorCode                                          ErrorCode = "MAP_KEY_LENGTH_TOO_SHORT"
	MAP_KEY_LENGTH_TOO_LONG_ErrorCode                                           ErrorCode = "MAP_KEY_LENGTH_TOO_LONG"
	CUSTOMER_MISSING_NAME_ErrorCode                                             ErrorCode = "CUSTOMER_MISSING_NAME"
	CUSTOMER_MISSING_EMAIL_ErrorCode                                            ErrorCode = "CUSTOMER_MISSING_EMAIL"
	INVALID_PAUSE_LENGTH_ErrorCode                                              ErrorCode = "INVALID_PAUSE_LENGTH"
	INVALID_DATE_ErrorCode                                                      ErrorCode = "INVALID_DATE"
	JOB_TEMPLATE_NAME_TAKEN_ErrorCode                                           ErrorCode = "JOB_TEMPLATE_NAME_TAKEN"
	CLIENT_NOT_SUPPORTED_ErrorCode                                              ErrorCode = "CLIENT_NOT_SUPPORTED"
	CARD_EXPIRED_ErrorCode                                                      ErrorCode = "CARD_EXPIRED"
	INVALID_EXPIRATION_ErrorCode                                                ErrorCode = "INVALID_EXPIRATION"
	INVALID_EXPIRATION_YEAR_ErrorCode                                           ErrorCode = "INVALID_EXPIRATION_YEAR"
	INVALID_EXPIRATION_DATE_ErrorCode                                           ErrorCode = "INVALID_EXPIRATION_DATE"
	UNSUPPORTED_CARD_BRAND_ErrorCode                                            ErrorCode = "UNSUPPORTED_CARD_BRAND"
	UNSUPPORTED_ENTRY_METHOD_ErrorCode                                          ErrorCode = "UNSUPPORTED_ENTRY_METHOD"
	INVALID_ENCRYPTED_CARD_ErrorCode                                            ErrorCode = "INVALID_ENCRYPTED_CARD"
	INVALID_CARD_ErrorCode                                                      ErrorCode = "INVALID_CARD"
	GENERIC_DECLINE_ErrorCode                                                   ErrorCode = "GENERIC_DECLINE"
	CVV_FAILURE_ErrorCode                                                       ErrorCode = "CVV_FAILURE"
	ADDRESS_VERIFICATION_FAILURE_ErrorCode                                      ErrorCode = "ADDRESS_VERIFICATION_FAILURE"
	INVALID_ACCOUNT_ErrorCode                                                   ErrorCode = "INVALID_ACCOUNT"
	CURRENCY_MISMATCH_ErrorCode                                                 ErrorCode = "CURRENCY_MISMATCH"
	INSUFFICIENT_FUNDS_ErrorCode                                                ErrorCode = "INSUFFICIENT_FUNDS"
	INSUFFICIENT_PERMISSIONS_ErrorCode                                          ErrorCode = "INSUFFICIENT_PERMISSIONS"
	CARDHOLDER_INSUFFICIENT_PERMISSIONS_ErrorCode                               ErrorCode = "CARDHOLDER_INSUFFICIENT_PERMISSIONS"
	INVALID_LOCATION_ErrorCode                                                  ErrorCode = "INVALID_LOCATION"
	TRANSACTION_LIMIT_ErrorCode                                                 ErrorCode = "TRANSACTION_LIMIT"
	VOICE_FAILURE_ErrorCode                                                     ErrorCode = "VOICE_FAILURE"
	PAN_FAILURE_ErrorCode                                                       ErrorCode = "PAN_FAILURE"
	EXPIRATION_FAILURE_ErrorCode                                                ErrorCode = "EXPIRATION_FAILURE"
	CARD_NOT_SUPPORTED_ErrorCode                                                ErrorCode = "CARD_NOT_SUPPORTED"
	INVALID_PIN_ErrorCode                                                       ErrorCode = "INVALID_PIN"
	MISSING_PIN_ErrorCode                                                       ErrorCode = "MISSING_PIN"
	MISSING_ACCOUNT_TYPE_ErrorCode                                              ErrorCode = "MISSING_ACCOUNT_TYPE"
	INVALID_POSTAL_CODE_ErrorCode                                               ErrorCode = "INVALID_POSTAL_CODE"
	INVALID_FEES_ErrorCode                                                      ErrorCode = "INVALID_FEES"
	MANUALLY_ENTERED_PAYMENT_NOT_SUPPORTED_ErrorCode                            ErrorCode = "MANUALLY_ENTERED_PAYMENT_NOT_SUPPORTED"
	PAYMENT_LIMIT_EXCEEDED_ErrorCode                                            ErrorCode = "PAYMENT_LIMIT_EXCEEDED"
	GIFT_CARD_AVAILABLE_AMOUNT_ErrorCode                                        ErrorCode = "GIFT_CARD_AVAILABLE_AMOUNT"
	GIFT_CARD_BUYER_DAILY_LIMIT_REACHED_ErrorCode                               ErrorCode = "GIFT_CARD_BUYER_DAILY_LIMIT_REACHED"
	GIFT_CARD_INVALID_AMOUNT_ErrorCode                                          ErrorCode = "GIFT_CARD_INVALID_AMOUNT"
	GIFT_CARD_MAX_VALUE_REACHED_ErrorCode                                       ErrorCode = "GIFT_CARD_MAX_VALUE_REACHED"
	GIFT_CARD_MERCHANT_MAX_OUTSTANDING_BALANCE_REACHED_ErrorCode                ErrorCode = "GIFT_CARD_MERCHANT_MAX_OUTSTANDING_BALANCE_REACHED"
	GIFT_CARD_VALUE_ADDITION_LIMIT_REACHED_ErrorCode                            ErrorCode = "GIFT_CARD_VALUE_ADDITION_LIMIT_REACHED"
	ACCOUNT_UNUSABLE_ErrorCode                                                  ErrorCode = "ACCOUNT_UNUSABLE"
	BUYER_REFUSED_PAYMENT_ErrorCode                                             ErrorCode = "BUYER_REFUSED_PAYMENT"
	DELAYED_TRANSACTION_EXPIRED_ErrorCode                                       ErrorCode = "DELAYED_TRANSACTION_EXPIRED"
	DELAYED_TRANSACTION_CANCELED_ErrorCode                                      ErrorCode = "DELAYED_TRANSACTION_CANCELED"
	DELAYED_TRANSACTION_CAPTURED_ErrorCode                                      ErrorCode = "DELAYED_TRANSACTION_CAPTURED"
	DELAYED_TRANSACTION_FAILED_ErrorCode                                        ErrorCode = "DELAYED_TRANSACTION_FAILED"
	CARD_TOKEN_EXPIRED_ErrorCode                                                ErrorCode = "CARD_TOKEN_EXPIRED"
	CARD_TOKEN_USED_ErrorCode                                                   ErrorCode = "CARD_TOKEN_USED"
	AMOUNT_TOO_HIGH_ErrorCode                                                   ErrorCode = "AMOUNT_TOO_HIGH"
	UNSUPPORTED_INSTRUMENT_TYPE_ErrorCode                                       ErrorCode = "UNSUPPORTED_INSTRUMENT_TYPE"
	REFUND_AMOUNT_INVALID_ErrorCode                                             ErrorCode = "REFUND_AMOUNT_INVALID"
	REFUND_ALREADY_PENDING_ErrorCode                                            ErrorCode = "REFUND_ALREADY_PENDING"
	PAYMENT_NOT_REFUNDABLE_ErrorCode                                            ErrorCode = "PAYMENT_NOT_REFUNDABLE"
	REFUND_DECLINED_ErrorCode                                                   ErrorCode = "REFUND_DECLINED"
	INVALID_CARD_DATA_ErrorCode                                                 ErrorCode = "INVALID_CARD_DATA"
	SOURCE_USED_ErrorCode                                                       ErrorCode = "SOURCE_USED"
	SOURCE_EXPIRED_ErrorCode                                                    ErrorCode = "SOURCE_EXPIRED"
	UNSUPPORTED_LOYALTY_REWARD_TIER_ErrorCode                                   ErrorCode = "UNSUPPORTED_LOYALTY_REWARD_TIER"
	LOCATION_MISMATCH_ErrorCode                                                 ErrorCode = "LOCATION_MISMATCH"
	ORDER_EXPIRED_ErrorCode                                                     ErrorCode = "ORDER_EXPIRED"
	ORDER_ALREADY_USED_ErrorCode                                                ErrorCode = "ORDER_ALREADY_USED"
	ORDER_TOO_MANY_CATALOG_OBJECTS_ErrorCode                                    ErrorCode = "ORDER_TOO_MANY_CATALOG_OBJECTS"
	INSUFFICIENT_INVENTORY_ErrorCode                                            ErrorCode = "INSUFFICIENT_INVENTORY"
	PRICE_MISMATCH_ErrorCode                                                    ErrorCode = "PRICE_MISMATCH"
	VERSION_MISMATCH_ErrorCode                                                  ErrorCode = "VERSION_MISMATCH"
	IDEMPOTENCY_KEY_REUSED_ErrorCode                                            ErrorCode = "IDEMPOTENCY_KEY_REUSED"
	UNEXPECTED_VALUE_ErrorCode                                                  ErrorCode = "UNEXPECTED_VALUE"
	SANDBOX_NOT_SUPPORTED_ErrorCode                                             ErrorCode = "SANDBOX_NOT_SUPPORTED"
	INVALID_EMAIL_ADDRESS_ErrorCode                                             ErrorCode = "INVALID_EMAIL_ADDRESS"
	INVALID_PHONE_NUMBER_ErrorCode                                              ErrorCode = "INVALID_PHONE_NUMBER"
	CHECKOUT_EXPIRED_ErrorCode                                                  ErrorCode = "CHECKOUT_EXPIRED"
	BAD_CERTIFICATE_ErrorCode                                                   ErrorCode = "BAD_CERTIFICATE"
	INVALID_SQUARE_VERSION_FORMAT_ErrorCode                                     ErrorCode = "INVALID_SQUARE_VERSION_FORMAT"
	API_VERSION_INCOMPATIBLE_ErrorCode                                          ErrorCode = "API_VERSION_INCOMPATIBLE"
	INVALID_URL_ErrorCode                                                       ErrorCode = "INVALID_URL"
	HTTPS_ONLY_ErrorCode                                                        ErrorCode = "HTTPS_ONLY"
	UNREACHABLE_URL_ErrorCode                                                   ErrorCode = "UNREACHABLE_URL"
	SESSION_EXPIRED_ErrorCode                                                   ErrorCode = "SESSION_EXPIRED"
	CARD_DECLINED_ErrorCode                                                     ErrorCode = "CARD_DECLINED"
	VERIFY_CVV_FAILURE_ErrorCode                                                ErrorCode = "VERIFY_CVV_FAILURE"
	VERIFY_AVS_FAILURE_ErrorCode                                                ErrorCode = "VERIFY_AVS_FAILURE"
	CARD_DECLINED_CALL_ISSUER_ErrorCode                                         ErrorCode = "CARD_DECLINED_CALL_ISSUER"
	CARD_DECLINED_VERIFICATION_REQUIRED_ErrorCode                               ErrorCode = "CARD_DECLINED_VERIFICATION_REQUIRED"
	BAD_EXPIRATION_ErrorCode                                                    ErrorCode = "BAD_EXPIRATION"
	CHIP_INSERTION_REQUIRED_ErrorCode                                           ErrorCode = "CHIP_INSERTION_REQUIRED"
	ALLOWABLE_PIN_TRIES_EXCEEDED_ErrorCode                                      ErrorCode = "ALLOWABLE_PIN_TRIES_EXCEEDED"
	RESERVATION_DECLINED_ErrorCode                                              ErrorCode = "RESERVATION_DECLINED"
	BLOCKED_BY_BLOCKLIST_ErrorCode                                              ErrorCode = "BLOCKED_BY_BLOCKLIST"
	FULFILLMENT_PREFERENCES_RESTRICTED_DATE_NOT_UNIQUE_ErrorCode                ErrorCode = "FULFILLMENT_PREFERENCES_RESTRICTED_DATE_NOT_UNIQUE"
	FULFILLMENT_PREFERENCES_INVALID_ENABLEMENT_STATUS_WHEN_SCHEDULING_ErrorCode ErrorCode = "FULFILLMENT_PREFERENCES_INVALID_ENABLEMENT_STATUS_WHEN_SCHEDULING"
	FULFILLMENT_PREFERENCES_INVALID_ENABLEMENT_SCHEDULED_AT_ErrorCode           ErrorCode = "FULFILLMENT_PREFERENCES_INVALID_ENABLEMENT_SCHEDULED_AT"
	FULFILLMENT_PREFERENCES_FULFILLMENT_SCHEDULE_NOT_ALLOWED_ErrorCode          ErrorCode = "FULFILLMENT_PREFERENCES_FULFILLMENT_SCHEDULE_NOT_ALLOWED"
	FULFILLMENT_PREFERENCES_ASSIGNMENT_IS_IMMUTABLE_ErrorCode                   ErrorCode = "FULFILLMENT_PREFERENCES_ASSIGNMENT_IS_IMMUTABLE"
	INVALID_TIMEZONE_ErrorCode                                                  ErrorCode = "INVALID_TIMEZONE"
	UNKNOWN_BODY_PARAMETER_ErrorCode                                            ErrorCode = "UNKNOWN_BODY_PARAMETER"
	FULFILLMENT_PREFERENCES_CONFLICTING_ASSIGNMENT_TYPE_ErrorCode               ErrorCode = "FULFILLMENT_PREFERENCES_CONFLICTING_ASSIGNMENT_TYPE"
	NOT_FOUND_ErrorCode                                                         ErrorCode = "NOT_FOUND"
	APPLE_PAYMENT_PROCESSING_CERTIFICATE_HASH_NOT_FOUND_ErrorCode               ErrorCode = "APPLE_PAYMENT_PROCESSING_CERTIFICATE_HASH_NOT_FOUND"
	METHOD_NOT_ALLOWED_ErrorCode                                                ErrorCode = "METHOD_NOT_ALLOWED"
	NOT_ACCEPTABLE_ErrorCode                                                    ErrorCode = "NOT_ACCEPTABLE"
	REQUEST_TIMEOUT_ErrorCode                                                   ErrorCode = "REQUEST_TIMEOUT"
	CONFLICT_ErrorCode                                                          ErrorCode = "CONFLICT"
	GONE_ErrorCode                                                              ErrorCode = "GONE"
	REQUEST_ENTITY_TOO_LARGE_ErrorCode                                          ErrorCode = "REQUEST_ENTITY_TOO_LARGE"
	UNSUPPORTED_MEDIA_TYPE_ErrorCode                                            ErrorCode = "UNSUPPORTED_MEDIA_TYPE"
	UNPROCESSABLE_ENTITY_ErrorCode                                              ErrorCode = "UNPROCESSABLE_ENTITY"
	RATE_LIMITED_ErrorCode                                                      ErrorCode = "RATE_LIMITED"
	CLIENT_CLOSED_REQUEST_ErrorCode                                             ErrorCode = "CLIENT_CLOSED_REQUEST"
	NOT_IMPLEMENTED_ErrorCode                                                   ErrorCode = "NOT_IMPLEMENTED"
	BAD_GATEWAY_ErrorCode                                                       ErrorCode = "BAD_GATEWAY"
	SERVICE_UNAVAILABLE_ErrorCode                                               ErrorCode = "SERVICE_UNAVAILABLE"
	TEMPORARY_ERROR_ErrorCode                                                   ErrorCode = "TEMPORARY_ERROR"
	GATEWAY_TIMEOUT_ErrorCode                                                   ErrorCode = "GATEWAY_TIMEOUT"
)
