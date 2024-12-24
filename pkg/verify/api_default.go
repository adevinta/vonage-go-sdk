/*
 * Nexmo Verify API
 *
 * The Verify API helps you to implement 2FA (two-factor authentication) in your applications. This is useful for:  * Protecting against spam, by preventing spammers from creating multiple accounts * Monitoring suspicious activity, by forcing an account user to verify ownership of a number * Ensuring that you can reach your users at any time because you have their correct phone number More information is available at <https://developer.nexmo.com/verify>
 *
 * API version: 1.1.3
 * Contact: devrel@nexmo.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package verify

import (
	"bytes"
	_context "context"
	"io"
	_nethttp "net/http"
	_neturl "net/url"
	"reflect"
	"strings"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ _context.Context
)

// DefaultApiService DefaultApi service
type DefaultApiService service

// VerifyCheckOpts Optional parameters for the method 'VerifyCheck'
type VerifyCheckOpts struct {
	IpAddress optional.String
}

/*
VerifyCheck Verify Check
Use Verify check to confirm that the PIN you received from your user matches the one sent by Nexmo in your Verify request.  1. Send the verification &#x60;code&#x60; that your user supplied, with the corresponding &#x60;request_id&#x60; from the Verify request. 2. Check the &#x60;status&#x60; of the response to determine if the code the user supplied matches the one sent by Nexmo.  *Note that this endpoint is available by &#x60;GET&#x60; request as well as &#x60;POST&#x60;.*
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param format The response format.
  - @param apiKey You can find your API key in your [account dashboard](https://dashboard.nexmo.com)
  - @param apiSecret You can find your API secret in your [account dashboard](https://dashboard.nexmo.com)
  - @param requestId The Verify request to check. This is the `request_id` you received in the response to the Verify request.
  - @param code The verification code entered by your user.
  - @param optional nil or *VerifyCheckOpts - Optional Parameters:
  - @param "IpAddress" (optional.String) -  (This field is no longer used)

@return CheckResponse
*/
func (a *DefaultApiService) VerifyCheck(ctx _context.Context, format string, apiKey string, apiSecret string, requestId string, code string, localVarOptionals *VerifyCheckOpts) (CheckResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CheckResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/check/{format}"
	localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", _neturl.QueryEscape(parameterToString(format, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if strlen(requestId) > 32 {
		return localVarReturnValue, nil, reportError("requestId must have less than 32 elements")
	}
	if strlen(code) < 4 {
		return localVarReturnValue, nil, reportError("code must have at least 4 elements")
	}
	if strlen(code) > 6 {
		return localVarReturnValue, nil, reportError("code must have less than 6 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarFormParams.Add("api_key", parameterToString(apiKey, ""))
	localVarFormParams.Add("api_secret", parameterToString(apiSecret, ""))
	localVarFormParams.Add("request_id", parameterToString(requestId, ""))
	localVarFormParams.Add("code", parameterToString(code, ""))
	if localVarOptionals != nil && localVarOptionals.IpAddress.IsSet() {
		localVarFormParams.Add("ip_address", parameterToString(localVarOptionals.IpAddress.Value(), ""))
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()

	// hack to reinstate the body in case we need it
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))

	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
VerifyControl Verify Control
Control the progress of your Verify requests. To cancel an existing Verify request, or to trigger the next verification event:   1. Send a Verify control request with the appropriate command (&#x60;cmd&#x60;) for what you want to achieve.  2. Check the &#x60;status&#x60; in the response.   *Note that this endpoint is available by &#x60;GET&#x60; request as well as &#x60;POST&#x60;.*
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param format The response format.
  - @param apiKey You can find your API key in your [account dashboard](https://dashboard.nexmo.com)
  - @param apiSecret You can find your API secret in your [account dashboard](https://dashboard.nexmo.com)
  - @param requestId The `request_id` you received in the response to the Verify request.
  - @param cmd The possible commands are `cancel` to request cancellation of the verification process, or `trigger_next_event` to advance  to the next verification event (if any). Cancellation is only possible 30 seconds after the start of the verification request and before the second event (either TTS or SMS) has taken place.

@return ControlResponse
*/
func (a *DefaultApiService) VerifyControl(ctx _context.Context, format string, apiKey string, apiSecret string, requestId string, cmd string) (ControlResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ControlResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/control/{format}"
	localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", _neturl.QueryEscape(parameterToString(format, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarFormParams.Add("api_key", parameterToString(apiKey, ""))
	localVarFormParams.Add("api_secret", parameterToString(apiSecret, ""))
	localVarFormParams.Add("request_id", parameterToString(requestId, ""))
	localVarFormParams.Add("cmd", parameterToString(cmd, ""))
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()

	// hack to reinstate the body in case we need it
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))

	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// VerifyRequestOpts Optional parameters for the method 'VerifyRequest'
type VerifyRequestOpts struct {
	Country       optional.String
	SenderId      optional.String
	CodeLength    optional.Int32
	Lg            optional.String
	PinExpiry     optional.Int32
	NextEventWait optional.Int32
	WorkflowId    optional.Int32
}

/*
VerifyRequest Request a Verification
Use Verify request to generate and send a PIN to your user:  1. Create a request to send a verification code to your user.  2. Check the &#x60;status&#x60; field in the response to ensure that your request was successful (zero is success).  3. Use the &#x60;request_id&#x60; field in the response for the Verify check.  *Note that this endpoint is available by &#x60;GET&#x60; request as well as &#x60;POST&#x60;.*
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param format The response format.
  - @param apiKey You can find your API key in your [account dashboard](https://dashboard.nexmo.com)
  - @param apiSecret You can find your API secret in your [account dashboard](https://dashboard.nexmo.com)
  - @param number The mobile or landline phone number to verify. Unless you are setting `country` explicitly, this number must be in [E.164](https://en.wikipedia.org/wiki/E.164) format.
  - @param brand An 18-character alphanumeric string you can use to personalize the verification request SMS body, to help users identify your company or application name. For example: \\\"Your `Acme Inc` PIN is ...\\\"
  - @param optional nil or *VerifyRequestOpts - Optional Parameters:
  - @param "Country" (optional.String) -  If you do not provide `number` in international format or you are not sure if `number` is correctly formatted, specify the two-character country code in `country`. Verify will then format the number for you.
  - @param "SenderId" (optional.String) -  An 11-character alphanumeric string that represents the [identity of the sender](https://developer.nexmo.com/messaging/sms/guides/custom-sender-id) of the verification request. Depending on the destination of the phone number you are sending the verification SMS to, restrictions might apply.
  - @param "CodeLength" (optional.Int32) -  The length of the verification code.
  - @param "Lg" (optional.String) -  By default, the SMS or text-to-speech (TTS) message is generated in the locale that matches the `number`. For example, the text message or TTS message for a `33*` number is sent in French. Use this parameter to explicitly control the language used for the Verify request. A list of languages is available: <https://developer.nexmo.com/verify/guides/verify-languages>
  - @param "PinExpiry" (optional.Int32) -  How long the generated verification code is valid for, in seconds. When you specify both `pin_expiry` and `next_event_wait` then `pin_expiry` must be an integer multiple of `next_event_wait` otherwise `pin_expiry` is defaulted to equal next_event_wait. See [changing the event timings](https://developer.nexmo.com/verify/guides/changing-default-timings).
  - @param "NextEventWait" (optional.Int32) -  Specifies the wait time in seconds between attempts to deliver the verification code.
  - @param "WorkflowId" (optional.Int32) -  Selects the predefined sequence of SMS and TTS (Text To Speech) actions to use in order to convey the PIN to your user. For example, an id of 1 identifies the workflow SMS - TTS - TTS. For a list of all workflows and their associated ids, please visit the [developer portal](https://developer.nexmo.com/verify/guides/workflows-and-events).

@return RequestResponse
*/
func (a *DefaultApiService) VerifyRequest(ctx _context.Context, format string, apiKey string, apiSecret string, number string, brand string, localVarOptionals *VerifyRequestOpts) (RequestResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  RequestResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/{format}"
	localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", _neturl.QueryEscape(parameterToString(format, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if strlen(brand) > 18 {
		return localVarReturnValue, nil, reportError("brand must have less than 18 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarFormParams.Add("api_key", parameterToString(apiKey, ""))
	localVarFormParams.Add("api_secret", parameterToString(apiSecret, ""))
	localVarFormParams.Add("number", parameterToString(number, ""))
	if localVarOptionals != nil && localVarOptionals.Country.IsSet() {
		localVarFormParams.Add("country", parameterToString(localVarOptionals.Country.Value(), ""))
	}
	localVarFormParams.Add("brand", parameterToString(brand, ""))
	if localVarOptionals != nil && localVarOptionals.SenderId.IsSet() {
		localVarFormParams.Add("sender_id", parameterToString(localVarOptionals.SenderId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CodeLength.IsSet() {
		localVarFormParams.Add("code_length", parameterToString(localVarOptionals.CodeLength.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Lg.IsSet() {
		localVarFormParams.Add("lg", parameterToString(localVarOptionals.Lg.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PinExpiry.IsSet() {
		localVarFormParams.Add("pin_expiry", parameterToString(localVarOptionals.PinExpiry.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.NextEventWait.IsSet() {
		localVarFormParams.Add("next_event_wait", parameterToString(localVarOptionals.NextEventWait.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.WorkflowId.IsSet() {
		localVarFormParams.Add("workflow_id", parameterToString(localVarOptionals.WorkflowId.Value(), ""))
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// VerifyRequestWithPSD2Opts Optional parameters for the method 'VerifyRequestWithPSD2'
type VerifyRequestWithPSD2Opts struct {
	Country       optional.String
	CodeLength    optional.Int32
	Lg            optional.String
	PinExpiry     optional.Int32
	NextEventWait optional.Int32
	WorkflowId    optional.Int32
}

/*
VerifyRequestWithPSD2 PSD2 (Payment Services Directive 2) Request
Use Verify request to generate and send a PIN to your user to authorize a payment: 1. Create a request to send a verification code to your user. 2. Check the &#x60;status&#x60; field in the response to ensure that your request was successful (zero is success). 3. Use the &#x60;request_id&#x60; field in the response for the Verify check. (Please note that XML format is not supported for the Payment Services Directive endpoint at this time.)
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param format The response format.
  - @param apiKey You can find your API key in your [account dashboard](https://dashboard.nexmo.com)
  - @param apiSecret You can find your API secret in your [account dashboard](https://dashboard.nexmo.com)
  - @param number The mobile or landline phone number to verify. Unless you are setting `country` explicitly, this number must be in [E.164](https://en.wikipedia.org/wiki/E.164) format.
  - @param payee An alphanumeric string to indicate to the user the name of the recipient that they are confirming a payment to.
  - @param amount The decimal amount of the payment to be confirmed, in Euros
  - @param optional nil or *VerifyRequestWithPSD2Opts - Optional Parameters:
  - @param "Country" (optional.String) -  If you do not provide `number` in international format or you are not sure if `number` is correctly formatted, specify the two-character country code in `country`. Verify will then format the number for you.
  - @param "CodeLength" (optional.Int32) -  The length of the verification code.
  - @param "Lg" (optional.String) -  By default, the SMS or text-to-speech (TTS) message is generated in the locale that matches the `number`. For example, the text message or TTS message for a `33*` number is sent in French. Use this parameter to explicitly control the language used. *Note: Voice calls in English for `bg-bg`, `ee-et`, `ga-ie`, `lv-lv`, `lt-lt`, `mt-mt`, `sk-sk`, `sk-si`
  - @param "PinExpiry" (optional.Int32) -  How long the generated verification code is valid for, in seconds. When you specify both `pin_expiry` and `next_event_wait` then `pin_expiry` must be an integer multiple of `next_event_wait` otherwise `pin_expiry` is defaulted to equal next_event_wait. See [changing the event timings](https://developer.nexmo.com/verify/guides/changing-default-timings).
  - @param "NextEventWait" (optional.Int32) -  Specifies the wait time in seconds between attempts to deliver the verification code.
  - @param "WorkflowId" (optional.Int32) -  Selects the predefined sequence of SMS and TTS (Text To Speech) actions to use in order to convey the PIN to your user. For example, an id of 1 identifies the workflow SMS - TTS - TTS. For a list of all workflows and their associated ids, please visit the [developer portal](https://developer.nexmo.com/verify/guides/workflows-and-events).

@return RequestResponse
*/
func (a *DefaultApiService) VerifyRequestWithPSD2(ctx _context.Context, format string, apiKey string, apiSecret string, number string, payee string, amount float32, localVarOptionals *VerifyRequestWithPSD2Opts) (RequestResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  RequestResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/psd2/{format}"
	localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", _neturl.QueryEscape(parameterToString(format, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if strlen(payee) > 18 {
		return localVarReturnValue, nil, reportError("payee must have less than 18 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarFormParams.Add("api_key", parameterToString(apiKey, ""))
	localVarFormParams.Add("api_secret", parameterToString(apiSecret, ""))
	localVarFormParams.Add("number", parameterToString(number, ""))
	if localVarOptionals != nil && localVarOptionals.Country.IsSet() {
		localVarFormParams.Add("country", parameterToString(localVarOptionals.Country.Value(), ""))
	}
	localVarFormParams.Add("payee", parameterToString(payee, ""))
	localVarFormParams.Add("amount", parameterToString(amount, ""))
	if localVarOptionals != nil && localVarOptionals.CodeLength.IsSet() {
		localVarFormParams.Add("code_length", parameterToString(localVarOptionals.CodeLength.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Lg.IsSet() {
		localVarFormParams.Add("lg", parameterToString(localVarOptionals.Lg.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PinExpiry.IsSet() {
		localVarFormParams.Add("pin_expiry", parameterToString(localVarOptionals.PinExpiry.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.NextEventWait.IsSet() {
		localVarFormParams.Add("next_event_wait", parameterToString(localVarOptionals.NextEventWait.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.WorkflowId.IsSet() {
		localVarFormParams.Add("workflow_id", parameterToString(localVarOptionals.WorkflowId.Value(), ""))
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()

	// hack to reinstate the body in case we need it
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))

	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// VerifySearchOpts Optional parameters for the method 'VerifySearch'
type VerifySearchOpts struct {
	RequestId  optional.String
	RequestIds optional.Interface
}

/*
VerifySearch Verify Search
Use Verify search to check the status of past or current verification requests:  1. Send a Verify search request containing the &#x60;request_id&#x60;s of the verification requests you are interested in. 2. Use the &#x60;status&#x60; of each verification request in the &#x60;checks&#x60; array of the response object to determine the outcome.  *Note that this endpoint is available by &#x60;POST&#x60; request as well as &#x60;GET&#x60;.*
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param format The response format.
  - @param apiKey
  - @param apiSecret
  - @param optional nil or *VerifySearchOpts - Optional Parameters:
  - @param "RequestId" (optional.String) -  The `request_id` you received in the Verify Request Response.
  - @param "RequestIds" (optional.Interface of []string) -  More than one `request_id`. Each `request_id` is a new parameter in the Verify Search request.

@return SearchResponse
*/
func (a *DefaultApiService) VerifySearch(ctx _context.Context, format string, apiKey string, apiSecret string, localVarOptionals *VerifySearchOpts) (SearchResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SearchResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/search/{format}"
	localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", _neturl.QueryEscape(parameterToString(format, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	localVarQueryParams.Add("api_key", parameterToString(apiKey, ""))
	localVarQueryParams.Add("api_secret", parameterToString(apiSecret, ""))
	if localVarOptionals != nil && localVarOptionals.RequestId.IsSet() {
		localVarQueryParams.Add("request_id", parameterToString(localVarOptionals.RequestId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RequestIds.IsSet() {
		t := localVarOptionals.RequestIds.Value()
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("request_ids", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("request_ids", parameterToString(t, "multi"))
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()

	// hack to reinstate the body in case we need it
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))

	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
