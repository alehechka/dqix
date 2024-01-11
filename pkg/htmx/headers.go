package htmx

// htmx Header Reference: https://htmx.org/reference/#headers

// htmx Request Headers: https://htmx.org/reference/#request_headers
const (
	// RequestHeaderBoosted	indicates that the request is via an element using hx-boost
	RequestHeaderBoosted string = "Hx-Boosted"
	// RequestHeaderCurrentURL represents the current URL of the browser
	RequestHeaderCurrentURL string = "Hx-Current-URL"
	// RequestHeaderHistoryRestoreRequest is `true` if the request is for history restoration after a miss in the local history cache
	RequestHeaderHistoryRestoreRequest string = "Hx-History-Restore-Request"
	// RequestHeaderPrompt represents the user response to an hx-prompt
	RequestHeaderPrompt string = "Hx-Prompt"
	// RequestHeaderRequest is always `true` if request is sent via htmx
	RequestHeaderRequest string = "Hx-Request"
	// RequestHeaderTarget represents the id of the target element if it exists
	RequestHeaderTarget string = "Hx-Target"
	// RequestHeaderTriggerName represents the name of the triggered element if it exists
	RequestHeaderTriggerName string = "Hx-Trigger-Name"
	// RequestHeaderTrigger represents the id of the triggered element if it exists
	RequestHeaderTrigger string = "Hx-Trigger"
)

// htmx Response Headers: https://htmx.org/reference/#response_headers
const (
	// ResponseHeaderLocation Allows you to do a client-side redirect that does not do a full page reload
	ResponseHeaderLocation string = "Hx-Location"
	// ResponseHeaderPushUrl pushes a new url into the history stack
	ResponseHeaderPushUrl string = "Hx-Push-Url"
	// ResponseHeaderRedirect can be used to do a client-side redirect to a new location
	ResponseHeaderRedirect string = "Hx-Redirect"
	// ResponseHeaderRefresh if set to “true” the client side will do a a full refresh of the page
	ResponseHeaderRefresh string = "Hx-Refresh"
	// ResponseHeaderReplaceUrl replaces the current URL in the location bar
	ResponseHeaderReplaceUrl string = "Hx-Replace-Url"
	// ResponseHeaderReswap Allows you to specify how the response will be swapped. See hx-swap for possible values
	ResponseHeaderReswap string = "Hx-Reswap"
	// ResponseHeaderRetarget A CSS selector that updates the target of the content update to a different element on the page
	ResponseHeaderRetarget string = "Hx-Retarget"
	// ResponseHeaderReselect A CSS selector that allows you to choose which part of the response is used to be swapped in. Overrides an existing hx-select on the triggering element
	ResponseHeaderReselect string = "Hx-Reselect"
	// ResponseHeaderTrigger allows you to trigger client side events, see the documentation for more info: https://htmx.org/headers/hx-trigger/
	ResponseHeaderTrigger string = "Hx-Trigger"
	// ResponseHeaderTriggerAfterSettle allows you to trigger client side events, see the documentation for more info: https://htmx.org/headers/hx-trigger/
	ResponseHeaderTriggerAfterSettle string = "Hx-Trigger-After-Settle"
	// ResponseHeaderTriggerAfterSwap allows you to trigger client side events, see the documentation for more info: https://htmx.org/headers/hx-trigger/
	ResponseHeaderTriggerAfterSwap string = "Hx-Trigger-After-Swap"
)

// htmx custom extension headers
const (
	// ResponseHeaderTitle is associated with a custom htmx extension, "title-header", that intercepts this response header and updates the document's title to the new value when present.
	ResponseHeaderTitle string = "Hx-Title"
)
