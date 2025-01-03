/*
 * Numbers API
 *
 * The Numbers API enables you to manage your existing numbers and buy new virtual numbers for use with Nexmo's APIs. Further information is here: <https://developer.nexmo.com/numbers/overview>
 *
 * API version: 1.0.18
 * Contact: devrel@nexmo.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package number

// AvailableNumbers struct for AvailableNumbers
type AvailableNumbers struct {
	// The total amount of numbers available in the pool.
	Count int32 `json:"count,omitempty"`
	// A paginated array of available numbers and their details.
	Numbers []Availablenumber `json:"numbers,omitempty"`
}
