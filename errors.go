package listrak

import "errors"

var ErrNotFound = errors.New("file or directory not found")
var ErrUnhandledStatusCode = errors.New("unhandled StatusCode")

var errorMap = make(map[string]error)

func init() {
	errorMap["ERROR_INVALID_CREDENTIALS"] = ErrInvalidCredentials
	errorMap["ERROR_INVALID_PARAMETER"] = ErrInvalidParameter
	errorMap["ERROR_MALFORMED_REQUEST_BODY"] = ErrMalformedRequestBody
	errorMap["ERROR_RESOURCE_DEPENDENCY"] = ErrResourceDependency
	errorMap["ERROR_UNABLE_TO_LOCATE_RESOURCE"] = ErrUnableToLocateResource
	errorMap["ERROR_UNAUTHORIZED"] = ErrUnauthorized
	errorMap["ERROR_UNHANDLED_EXCEPTION"] = ErrUhandledException
	errorMap["ERROR_UNKNOWN_ROUTE"] = ErrUnknownRoute
	errorMap["ERROR_UNSAFE_ROUTE"] = ErrUnsafeRoute
	errorMap["ERROR_UNSUPPORTED_CONTENT_TYPE"] = ErrUnsupportedContentType
	errorMap["ERROR_UNSUPPORTED_METHOD"] = ErrUnsupportedMethod
	errorMap["ERROR_UNSUPPORTED_PROTOCOL"] = ErrUnsupportedProtocol
	// Contact
	errorMap["ERROR_BANNED_EMAIL_ADDRESS"] = ErrBannedEmailAddress
	errorMap["ERROR_BANNED_NEW_EMAIL_ADDRESS"] = ErrBannedNewEmailAddress
	errorMap["ERROR_CHANGE_ADDRESS_CONTACT_UNSUBSCRIBED"] = ErrChangeAddressContactUnsubscribed
	errorMap["ERROR_CHANGE_ADDRESS_WITH_EVENTS"] = ErrChangeAddressWithEvents
	errorMap["ERROR_CHANGE_ADDRESS_WITH_SEGMENTATION_DATA"] = ErrChangeAddressWithSegmentationData
	errorMap["ERROR_CONVERSION_ANALYTICS_NOT_ENABLED"] = ErrConversionAnalyticsNotEnabled
	errorMap["ERROR_NOT_SUBSCRIBED_EMAIL_ADDRESS"] = ErrNotSubscribedEmailAddress
	errorMap["ERROR_PENDING_EMAIL_ADDRESS"] = ErrPendingEmailAddress
	errorMap["ERROR_PENDING_NEW_EMAIL_ADDRESS"] = ErrPendingNewEmailAddress
	errorMap["ERROR_SUBSCRIBED_EMAIL_ADDRESS"] = ErrSubscribedEmailAddress
	errorMap["ERROR_SUBSCRIBED_NEW_EMAIL_ADDRESS"] = ErrSubscribedNewEmailAddress
	errorMap["ERROR_SUPPRESSED_EMAIL"] = ErrSuppressedEmail
	errorMap["ERROR_SUPPRESSED_NEW_EMAIL"] = ErrSuppressedNewEmail
	errorMap["ERROR_TOO_MANY_SEGMENTATION_FIELDS"] = ErrTooManSegmentationFields
	errorMap["ERROR_UNSUBSCRIBE_WITH_EVENTS"] = ErrUnsubscribeWithEvents
	errorMap["ERROR_UNSUBSCRIBE_WITH_SEGMENTATION_DATA"] = ErrUnsubscribeWithSegmentationData
	errorMap["ERROR_UNSUBSCRIBED_EMAIL_ADDRESS"] = ErrUnsubscribeEemailAddress
	errorMap["ERROR_UNSUBSCRIBED_NEW_EMAIL_ADDRESS"] = ErrUnsubscribedNewEmailAddress
	errorMap["ERROR_UPDATE_UNSUBSCRIBE_WITH_EVENTS"] = ErrUpdateUnsubscribeWithEvents
	// List
	errorMap["ERROR_GOOGLE_TRACKING_DOMAIN_LIMIT_MET"] = ErrGoogleTrackingDomainLimitMet
	errorMap["ERROR_DOMAIN_ALIAS_INVALID"] = ErrDomainAliasInvalid
	errorMap["ERROR_DOMAIN_ALIAS_MISCONFIGURED_DNS"] = ErrDomainAliasMisconfiguredDNS
	errorMap["ERROR_DOMAIN_ALIAS_INVALID_HSTS"] = ErrDomainAliasInvalidHSTS
	errorMap["ERROR_DOMAIN_ALIAS_PROXY_CERTIFICATE_GENEREATION_FAILED"] = ErrDomainAliasProxyCertificateGenerationFailed
	errorMap["ERROR_DOMAIN_ALIAS_PROXY_CERTIFICATE_GENERATION_FAILED"] = ErrDomainAliasProxyCertificateGenerationFailed
	// Message
	errorMap["ERROR_REVIEW_FLAG_AND_SEND_DATE_SET"] = ErrReviewFlagAndSendDateSet
	errorMap["ERROR_TEST_FLAG_AND_REVIEW_FLAG_SET"] = ErrTestFlagAndReviewFlagSet
	errorMap["ERROR_TEST_FLAG_AND_SEND_DATE_SET"] = ErrTestFlagAndSendDateSet
	// Segmentation Field Group
	errorMap["ERROR_SEGMENTATION_FIELD_GROUP_LIMIT_MET"] = ErrSegmentationFieldGroupLimitMet
	// Transactional Message
	errorMap["ERROR_SEGMENTATION_FIELD_DEFINED_TWICE"] = ErrSegmentationFieldDefinedTwice
	errorMap["ERROR_TRANSACTIONAL_MESSAGE_EXTERNAL_CONTENT"] = ErrTransactionalMessageExternalContent
	errorMap["ERROR_TRANSACTIONAL_MESSAGE_SYSTEM_LINK"] = ErrTransactionalMessageSYSTEMLink
}

var ErrInvalidCredentials = errors.New("invalid client credentials, inactive integration, unauthorized IP address")
var ErrInvalidParameter = errors.New("null or invalid value for a parameter, non-unique value for a parameter (if the property is required to be unique for all resources in the collection)")
var ErrMalformedRequestBody = errors.New("invalid JSON body was supplied")
var ErrResourceDependency = errors.New("trying to delete a resource that is in use")
var ErrUnableToLocateResource = errors.New("valid route was supplied with invalid resource IDs")
var ErrUnauthorized = errors.New("invalid access token, inactive integration, insufficient access level, unauthorized IP address")
var ErrUhandledException = errors.New("unexpected error during execution of the request")
var ErrUnknownRoute = errors.New("invalid route was supplied")
var ErrUnsafeRoute = errors.New("invalid character was supplied in the route ('&' outside of query string, '<', '>', etc.)")
var ErrUnsupportedContentType = errors.New("invalid Content-Type header was supplied (must be application/json), missing Content-Type header")
var ErrUnsupportedMethod = errors.New("valid route supplied with an invalid HTTP method")
var ErrUnsupportedProtocol = errors.New("request was not made over HTTPS")

var ErrBannedEmailAddress = errors.New("subscribing a contact that has been banned")
var ErrBannedNewEmailAddress = errors.New("changing emailAddress to one that has been banned")
var ErrChangeAddressContactUnsubscribed = errors.New("changing emailAddress to one that has unsubscribed from the list")
var ErrChangeAddressWithEvents = errors.New("changing emailAddress with a list of events")
var ErrChangeAddressWithSegmentationData = errors.New("changing emailAddress with profile data")
var ErrConversionAnalyticsNotEnabled = errors.New("conversion analytics not enabled")
var ErrNotSubscribedEmailAddress = errors.New("unsubscribing a contact that is not a subscriber for the list")
var ErrPendingEmailAddress = errors.New("already pending subscription through double opt-in")
var ErrPendingNewEmailAddress = errors.New("changing emailAddress to one that is already pending subscription through double opt-in")
var ErrSubscribedEmailAddress = errors.New("subscribing a contact that has already been subscribed to the list")
var ErrSubscribedNewEmailAddress = errors.New("changing emailAddress to one that is already subscribed to the list")
var ErrSuppressedEmail = errors.New("subscribing an emailAddress that has already been suppressed")
var ErrSuppressedNewEmail = errors.New("changing emailAddress to one that has already been suppressed")
var ErrTooManSegmentationFields = errors.New("number of requested profile fields exceeds the limit")
var ErrUnsubscribeWithEvents = errors.New("unsubscribing a contact with a list of events")
var ErrUnsubscribeWithSegmentationData = errors.New("unsubscribing a contact with profile data")
var ErrUnsubscribeEemailAddress = errors.New("unsubscribing a contact that has already been unsubscribed from the list")
var ErrUnsubscribedNewEmailAddress = errors.New("changing an emailAddress to one that has already been unsubscribed from the list")
var ErrUpdateUnsubscribeWithEvents = errors.New("updating a contact that is unsubscribed with a list of events")

var ErrGoogleTrackingDomainLimitMet = errors.New("maximum number of Google tracking domains have already been set for the list")
var ErrDomainAliasInvalid = errors.New("the provided domain alias is not a valid domain")
var ErrDomainAliasMisconfiguredDNS = errors.New("the DNS record for the provided domain alias does not point to the correct endpoint")
var ErrDomainAliasInvalidHSTS = errors.New("the provided domain alias has HSTS enabled for either all sub-domains, or a domain alias type we are unable to secure")
var ErrDomainAliasProxyCertificateGenerationFailed = errors.New("we were unable to secure the specified domain alias")

var ErrReviewFlagAndSendDateSet = errors.New("sendReviewMessage set to true and sendDate is set")
var ErrTestFlagAndReviewFlagSet = errors.New("sendReviewMessage and sendTestMessage set to true")
var ErrTestFlagAndSendDateSet = errors.New("sendTestMessage set to true and sendDate is set")

var ErrSegmentationFieldGroupLimitMet = errors.New("maximum number of profile field groups have already been created for the list")

// Transactional Message
var ErrSegmentationFieldDefinedTwice = errors.New("two or more values were supplied for the same segmentationFieldId")
var ErrTransactionalMessageExternalContent = errors.New("contains external content tags")
var ErrTransactionalMessageSYSTEMLink = errors.New("contains system link tags")
