/*
 * Voice API
 *
 * The Voice API lets you create outbound calls, control in-progress calls and get information about historical calls. More information about the Voice API can be found at <https://developer.nexmo.com/voice/voice-api/overview>.
 *
 * API version: 1.3.2
 * Contact: devrel@nexmo.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package voice

// EndpointPhoneFrom Connect to a Phone (PSTN) number
type EndpointPhoneFrom struct {
	// The type of connection. Must be `phone`
	Type string `json:"type"`
	// The phone number to connect to
	Number string `json:"number"`
}