/*
 * Application API
 *
 * Nexmo provides an Application API to allow management of your Nexmo Applications.  This API is backwards compatible with version 1. Applications created using version 1 of the API can also be managed using version 2 (this version) of the API.
 *
 * API version: 2.0.5
 * Contact: devrel@nexmo.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package application

// ApplicationResponseCapabilitiesVoiceWebhooksFallbackAnswerUrl struct for ApplicationResponseCapabilitiesVoiceWebhooksFallbackAnswerUrl
type ApplicationResponseCapabilitiesVoiceWebhooksFallbackAnswerUrl struct {
	// If your `answer_url` is offline or returns a HTTP error code, Nexmo will make a request to a `fallback_answer_url` if it is set. This URL must return an NCCO.
	Address string `json:"address,omitempty"`
	// The HTTP method used to fetch your NCCO from your `answer_url`
	HttpMethod string `json:"http_method,omitempty"`
}
