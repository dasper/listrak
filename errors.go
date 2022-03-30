package listrak

import "errors"

const ERROR_INVALID_CREDENTIALS = "invalid client credentials, inactive integration, unauthorized IP address"
const ERROR_INVALID_PARAMETER = "null or invalid value for a parameter, non-unique value for a parameter (if the property is required to be unique for all resources in the collection)"
const ERROR_MALFORMED_REQUEST_BODY = "invalid JSON body was supplied"
const ERROR_RESOURCE_DEPENDENCY = "trying to delete a resource that is in use"
const ERROR_UNABLE_TO_LOCATE_RESOURCE = "valid route was supplied with invalid resource IDs"
const ERROR_UNAUTHORIZED = "invalid access token, inactive integration, insufficient access level, unauthorized IP address"
const ERROR_UNHANDLED_EXCEPTION = "unexpected error during execution of the request"
const ERROR_UNKNOWN_ROUTE = "invalid route was supplied"
const ERROR_UNSAFE_ROUTE = "invalid character was supplied in the route ('&' outside of query string, '<', '>', etc.)"
const ERROR_UNSUPPORTED_CONTENT_TYPE = "invalid Content-Type header was supplied (must be application/json), missing Content-Type header"
const ERROR_UNSUPPORTED_METHOD = "valid route supplied with an invalid HTTP method"
const ERROR_UNSUPPORTED_PROTOCOL = "request was not made over HTTPS"

var ErrInvalidCredentials = errors.New(ERROR_INVALID_CREDENTIALS)
var ErrInvalidParameter = errors.New(ERROR_INVALID_PARAMETER)
var ErrMalformedRequestBody = errors.New(ERROR_MALFORMED_REQUEST_BODY)
var ErrResourceDependency = errors.New(ERROR_RESOURCE_DEPENDENCY)
var ErrUnableToLocateResource = errors.New(ERROR_UNABLE_TO_LOCATE_RESOURCE)
var ErrUnauthorized = errors.New(ERROR_UNAUTHORIZED)
var ErrUhandledException = errors.New(ERROR_UNHANDLED_EXCEPTION)
var ErrUnknownRoute = errors.New(ERROR_UNKNOWN_ROUTE)
var ErrUnsafeRoute = errors.New(ERROR_UNSAFE_ROUTE)
var ErrUnsupportedContentType = errors.New(ERROR_UNSUPPORTED_CONTENT_TYPE)
var ErrUnsupportedMethod = errors.New(ERROR_UNSUPPORTED_METHOD)
var ErrUnsupportedProtocol = errors.New(ERROR_UNSUPPORTED_PROTOCOL)

// Contact
const ERROR_BANNED_EMAIL_ADDRESS = "subscribing a contact that has been banned"
const ERROR_BANNED_NEW_EMAIL_ADDRESS = "changing emailAddress to one that has been banned"
const ERROR_CHANGE_ADDRESS_CONTACT_UNSUBSCRIBED = "changing emailAddress to one that has unsubscribed from the list"
const ERROR_CHANGE_ADDRESS_WITH_EVENTS = "changing emailAddress with a list of events"
const ERROR_CHANGE_ADDRESS_WITH_SEGMENTATION_DATA = "changing emailAddress with profile data"
const ERROR_CONVERSION_ANALYTICS_NOT_ENABLED = "conversion analytics not enabled"
const ERROR_NOT_SUBSCRIBED_EMAIL_ADDRESS = "unsubscribing a contact that is not a subscriber for the list"
const ERROR_PENDING_EMAIL_ADDRESS = "already pending subscription through double opt-in"
const ERROR_PENDING_NEW_EMAIL_ADDRESS = "changing emailAddress to one that is already pending subscription through double opt-in"
const ERROR_SUBSCRIBED_EMAIL_ADDRESS = "subscribing a contact that has already been subscribed to the list"
const ERROR_SUBSCRIBED_NEW_EMAIL_ADDRESS = "changing emailAddress to one that is already subscribed to the list"
const ERROR_SUPPRESSED_EMAIL = "subscribing an emailAddress that has already been suppressed"
const ERROR_SUPPRESSED_NEW_EMAIL = "changing emailAddress to one that has already been suppressed"
const ERROR_TOO_MANY_SEGMENTATION_FIELDS = "number of requested profile fields exceeds the limit"
const ERROR_UNSUBSCRIBE_WITH_EVENTS = "unsubscribing a contact with a list of events"
const ERROR_UNSUBSCRIBE_WITH_SEGMENTATION_DATA = "unsubscribing a contact with profile data"
const ERROR_UNSUBSCRIBED_EMAIL_ADDRESS = "unsubscribing a contact that has already been unsubscribed from the list"
const ERROR_UNSUBSCRIBED_NEW_EMAIL_ADDRESS = "changing an emailAddress to one that has already been unsubscribed from the list"
const ERROR_UPDATE_UNSUBSCRIBE_WITH_EVENTS = "updating a contact that is unsubscribed with a list of events"

var ErrBannedEmailAddress = errors.New(ERROR_BANNED_EMAIL_ADDRESS)
var ErrBannedNewEmailAddress = errors.New(ERROR_BANNED_NEW_EMAIL_ADDRESS)
var ErrChangeAddressContactUnsubscribed = errors.New(ERROR_CHANGE_ADDRESS_CONTACT_UNSUBSCRIBED)
var ErrChangeAddressWithEvents = errors.New(ERROR_CHANGE_ADDRESS_WITH_EVENTS)
var ErrChangeAddressWithSegmentationData = errors.New(ERROR_CHANGE_ADDRESS_WITH_SEGMENTATION_DATA)
var ErrConversionAnalyticsNotEnabled = errors.New(ERROR_CONVERSION_ANALYTICS_NOT_ENABLED)
var ErrNotSubscribedEmailAddress = errors.New(ERROR_NOT_SUBSCRIBED_EMAIL_ADDRESS)
var ErrPendingEmailAddress = errors.New(ERROR_PENDING_EMAIL_ADDRESS)
var ErrPendingNewEmailAddress = errors.New(ERROR_PENDING_NEW_EMAIL_ADDRESS)
var ErrSubscribedEmailAddress = errors.New(ERROR_SUBSCRIBED_EMAIL_ADDRESS)
var ErrSubscribedNewEmailAddress = errors.New(ERROR_SUBSCRIBED_NEW_EMAIL_ADDRESS)
var ErrSuppressedEmail = errors.New(ERROR_SUPPRESSED_EMAIL)
var ErrSuppressedNewEmail = errors.New(ERROR_SUPPRESSED_NEW_EMAIL)
var ErrTooManSegmentationFields = errors.New(ERROR_TOO_MANY_SEGMENTATION_FIELDS)
var ErrUnsubscribeWithEvents = errors.New(ERROR_UNSUBSCRIBE_WITH_EVENTS)
var ErrUnsubscribeWithSegmentationData = errors.New(ERROR_UNSUBSCRIBE_WITH_SEGMENTATION_DATA)
var ErrUnsubscribeEemailAddress = errors.New(ERROR_UNSUBSCRIBED_EMAIL_ADDRESS)
var ErrUnsubscribedNewEmailAddress = errors.New(ERROR_UNSUBSCRIBED_NEW_EMAIL_ADDRESS)
var ErrUpdateUnsubscribeWithEvents = errors.New(ERROR_UPDATE_UNSUBSCRIBE_WITH_EVENTS)

// List
const ERROR_GOOGLE_TRACKING_DOMAIN_LIMIT_MET = "Maximum number of Google tracking domains have already been set for the list"
const ERROR_DOMAIN_ALIAS_INVALID = "The provided domain alias is not a valid domain"
const ERROR_DOMAIN_ALIAS_MISCONFIGURED_DNS = "The DNS record for the provided domain alias does not point to the correct endpoint"
const ERROR_DOMAIN_ALIAS_INVALID_HSTS = "The provided domain alias has HSTS enabled for either all sub-domains, or a domain alias type we are unable to secure"
const ERROR_DOMAIN_ALIAS_PROXY_CERTIFICATE_GENEREATION_FAILED = "We were unable to secure the specfied domain alias"

// Message
const ERROR_REVIEW_FLAG_AND_SEND_DATE_SET = "sendReviewMessage set to true and sendDate is set"
const ERROR_TEST_FLAG_AND_REVIEW_FLAG_SET = "sendReviewMessage and sendTestMessage set to true"
const ERROR_TEST_FLAG_AND_SEND_DATE_SET = "sendTestMessage set to true and sendDate is set"

// Segmentation Field Group
const ERROR_SEGMENTATION_FIELD_GROUP_LIMIT_MET = "Maximum number of profile field groups have already been created for the list"

// Transactional Message
const ERROR_SEGMENTATION_FIELD_DEFINED_TWICE = "Two or more values were supplied for the same segmentationFieldId"
const ERROR_TRANSACTIONAL_MESSAGE_EXTERNAL_CONTENT = "Contains external content tags"
const ERROR_TRANSACTIONAL_MESSAGE_SYSTEM_LINK = "Contains system link tags"
