# golang-interview

Simple HTTP server with a few bugs to solution during a working interview.

## Issue One, starting the server

There are several issues contained within this project. First the candidate should be prompted to start the server and issue a request to the [root](http://localhost:8080/simple). This will return a 404 status code until a handler is registered with the following snippet:

```
mux.Handle(simpleEndpoint, http.HandlerFunc(handlers.SimpleHandler))
```

## Issue Two, return JSON content from handler

The simple handler is initially configured to return a 'Hello World!' response. Prompt the candidate to update the simple handler to return the 'SimpleStruct' type with any string values in the two fields. Completion of this task should result in a properly formatted response with appropriate content type and error handling. Ensure candidate runs and corrects testing for the handler function.

```
simpleStruct := SimpleStruct{
	Reference:   "simple",
	HiddenValue: "hidden",
}

w.Header().Set("Content-Type", "application/json")
if err := json.NewEncoder(w).Encode(simpleStruct); err != nil {
	http.Error(w, fmt.Sprintf("failed to encode response: %s", err.Error()),
		http.StatusInternalServerError)
}
```

## Issue Three, block hidden value from response

The 'SimpleStruct' type is configured to return all fields in the JSON response. Update this struct to block the hdden value from returning to the end user. This should be an internal only field. Ensure candidate runs and corrects testing for the handler function.

```
type SimpleStruct struct {
	Reference   string `json:"reference"`
	HiddenValue string `json:"-"`
}
```