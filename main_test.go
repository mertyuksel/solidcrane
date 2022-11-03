package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestConvertUpperCase(t *testing.T) {

	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.POST("/uppercase", convertUpperCase)

	body := strings.NewReader(`{"message":"Just testing"}`)

	req, err := http.NewRequest(http.MethodPost, "/uppercase", body)

	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	fmt.Println(w.Body)

	var response text

	error := json.NewDecoder(w.Body).Decode(&response)
	if error != nil {
		t.Errorf("Unexpected error: %s", error)
	}

	// Check to see if the response was what you expected
	if response.Message == "JUST TESTING" {
		t.Logf("Expected to get string \"%s\" is same as \"%s\" \n", "JUST TESTING", response.Message)
	} else {
		t.Fatalf("Expected to get status \"%s\" but instead got \"%s\"\n", "JUST TESTING", response.Message)
	}

	// Check to see if the status was what you expected
	if w.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same as %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}
