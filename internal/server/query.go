package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"joshuapare.com/fxdb/internal/engine"
)

// QueryHandler is an http.Handler that runs a query based
// on the contents
type QueryHandler struct {
	CE *engine.CollectionEngine
}

// QueryResponse is a response back to the client on a query
type QueryResponse struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func NewQueryHandler(ce *engine.CollectionEngine) *QueryHandler {
	return &QueryHandler{
		CE: ce,
	}
}

// ServeHTTP handles an HTTP request to the /echo endpoint.
func (h *QueryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := h.ParseQuery(r.Body)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// ParseQuery parses out the tokens of a query to be executed by
// the various query handlers
func (h *QueryHandler) ParseQuery(query io.ReadCloser) *QueryResponse {
	var response *QueryResponse

	// Parse out the string from the
	buf := new(strings.Builder)
	_, err := io.Copy(buf, query)
	if err != nil {
		fmt.Println(err)
		response = createErrorResponse("RUNTIME", err.Error())
		return response
	}
	queryString := buf.String()
	fmt.Printf("QUERY: %s\n", queryString)

	// Parse out the operation token to execute, passing the remaining portion
	// of the query to be handled by the operation parser
	queryParts := strings.SplitN(queryString, " ", 2)
	switch strings.ToLower(queryParts[0]) {
	case "list":
		response = h.ParseListQuery(queryParts[1])
	case "create":
		response = h.ParseCreateQuery(queryParts[1])
	case "delete":
		response = h.ParseDeleteQuery(queryParts[1])
	}

	return response
}

// ParseCreateQuery parses and executes a CREATE command
func (h *QueryHandler) ParseCreateQuery(query string) *QueryResponse {
	parts := strings.Split(query, " ")
	if len(parts) > 1 {
		return createErrorResponse("SYNTAX", "too many arguments supplied to CREATE")
	}

	if err := h.CE.CreateCollection(parts[0]); err != nil {
		return createErrorResponse("RUNTIME", err.Error())
	}

	return createSuccessResponse("collection created", nil)
}

// ParseListQuery parses and executes a LIST command
func (h *QueryHandler) ParseListQuery(query string) *QueryResponse {
	parts := strings.Split(query, " ")
	if len(parts) > 1 {
		return createErrorResponse("SYNTAX", "too many arguments supplied to LIST")
	}

	switch strings.ToLower(parts[0]) {
	case "collections":
		collectionList := h.CE.GetCollections()
		responseData := make(map[string]interface{})
		responseData["collections"] = collectionList
		return createSuccessResponse("collections retrieved", responseData)
	default:
		return createErrorResponse("SYNTAX", fmt.Sprintf("invalid operand '%s' passed to LIST command", parts[0]))
	}

}

// ParseDeleteQuery parses and executes a DELETE command
func (h *QueryHandler) ParseDeleteQuery(query string) *QueryResponse {
	parts := strings.Split(query, " ")
	if len(parts) > 1 {
		return createErrorResponse("SYNTAX", "too many arguments supplied to DELETE")
	}

	if err := h.CE.DeleteCollection(parts[0]); err != nil {
		return createErrorResponse("RUNTIME", err.Error())
	}

	return createSuccessResponse(fmt.Sprintf("collection '%s' deleted", parts[0]), nil)
}

func createSuccessResponse(message string, data map[string](interface{})) *QueryResponse {
	return &QueryResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
}

func createErrorResponse(category string, reason string) *QueryResponse {
	err := createQueryError(category, reason)

	return &QueryResponse{
		Message: err.Error(),
		Success: false,
		Data:    nil,
	}
}

func createQueryError(category string, reason string) error {
	var categoryMessage string

	switch category {
	case "SYNTAX":
		categoryMessage = "you have an error in your syntax"
	case "RUNTIME":
		categoryMessage = "there was an error running your query"
	default:
		categoryMessage = "hit an unexpected error"
	}

	return fmt.Errorf("%s: %s", strings.ToUpper(categoryMessage), reason)
}
