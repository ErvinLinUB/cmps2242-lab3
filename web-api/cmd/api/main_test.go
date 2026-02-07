package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// Task 4: Write HTTP Handler Tests

// TestHomeHandler tests the home route handler
func TestHomeHandler(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid GET request",
			method:         "GET",
			expectedStatus: http.StatusOK,
			expectedBody:   "Welcome to the Shapes API",
		},
		{
			name:           "Invalid POST request",
			method:         "POST",
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Method not allowed\n",
		},
		{
			name:           "Invalid PUT request",
			method:         "PUT",
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Method not allowed\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/", nil)
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(homeHandler)

			handler.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("status code: got %v, expected %v", rr.Code, tt.expectedStatus)
			}

			if rr.Body.String() != tt.expectedBody {
				t.Errorf("response body: got %v, expected %v", rr.Body.String(), tt.expectedBody)
			}
		})
	}
}

// TestHealthHandler tests the health route handler
func TestHealthHandler(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid GET request",
			method:         "GET",
			expectedStatus: http.StatusOK,
			expectedBody:   "Server is running",
		},
		{
			name:           "Invalid POST request",
			method:         "POST",
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Method not allowed\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/health", nil)
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(healthHandler)

			handler.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("status code: got %v, expected %v", rr.Code, tt.expectedStatus)
			}

			if rr.Body.String() != tt.expectedBody {
				t.Errorf("response body: got %v, expected %v", rr.Body.String(), tt.expectedBody)
			}
		})
	}
}

// TestAboutHandler tests the about route handler
func TestAboutHandler(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid GET request",
			method:         "GET",
			expectedStatus: http.StatusOK,
			expectedBody:   "Ervin Lin",
		},
		{
			name:           "Invalid POST request",
			method:         "POST",
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Method not allowed\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/about", nil)
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(aboutHandler)

			handler.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("status code: got %v, expected %v", rr.Code, tt.expectedStatus)
			}

			if rr.Body.String() != tt.expectedBody {
				t.Errorf("response body: got %v, expected %v", rr.Body.String(), tt.expectedBody)
			}
		})
	}
}

// TestTimeHandler tests the time route handler
func TestTimeHandler(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		expectedStatus int
		shouldBeTime   bool
	}{
		{
			name:           "Valid GET request",
			method:         "GET",
			expectedStatus: http.StatusOK,
			shouldBeTime:   true,
		},
		{
			name:           "Invalid POST request",
			method:         "POST",
			expectedStatus: http.StatusMethodNotAllowed,
			shouldBeTime:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/time", nil)
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(timeHandler)

			handler.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("status code: got %v, expected %v", rr.Code, tt.expectedStatus)
			}

			if tt.shouldBeTime {
				// Test that the response is a valid time format
				body := rr.Body.String()
				_, err := time.Parse("2006-01-02 15:04:05", body)
				if err != nil {
					t.Errorf("response body is not valid time format: got %v, error: %v", body, err)
				}
			}
		})
	}
}

// TestGreetingHandler tests the custom greeting route handler
func TestGreetingHandler(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		path           string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid GET request - default name",
			method:         "GET",
			path:           "/greeting",
			expectedStatus: http.StatusOK,
			expectedBody:   "Hello, Guest! Welcome to the Shapes API!",
		},
		{
			name:           "Valid GET request - with name parameter",
			method:         "GET",
			path:           "/greeting?name=John",
			expectedStatus: http.StatusOK,
			expectedBody:   "Hello, John! Welcome to the Shapes API!",
		},
		{
			name:           "Valid GET request - with different name",
			method:         "GET",
			path:           "/greeting?name=Alice",
			expectedStatus: http.StatusOK,
			expectedBody:   "Hello, Alice! Welcome to the Shapes API!",
		},
		{
			name:           "Invalid POST request",
			method:         "POST",
			path:           "/greeting",
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Method not allowed\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, nil)
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(greetingHandler)

			handler.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("status code: got %v, expected %v", rr.Code, tt.expectedStatus)
			}

			body := rr.Body.String()
			if !strings.Contains(body, tt.expectedBody) && body != tt.expectedBody {
				t.Errorf("response body: got %v, expected to contain %v", body, tt.expectedBody)
			}
		})
	}
}

// TestAllHandlersWithTable tests all five handlers in a single table-driven test
func TestAllHandlersWithTable(t *testing.T) {
	tests := []struct {
		name           string
		handler        http.HandlerFunc
		method         string
		path           string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Home Handler - GET",
			handler:        homeHandler,
			method:         "GET",
			path:           "/",
			expectedStatus: http.StatusOK,
			expectedBody:   "Welcome to the Shapes API",
		},
		{
			name:           "Health Handler - GET",
			handler:        healthHandler,
			method:         "GET",
			path:           "/health",
			expectedStatus: http.StatusOK,
			expectedBody:   "Server is running",
		},
		{
			name:           "About Handler - GET",
			handler:        aboutHandler,
			method:         "GET",
			path:           "/about",
			expectedStatus: http.StatusOK,
			expectedBody:   "Ervin Lin",
		},
		{
			name:           "Greeting Handler - GET with default",
			handler:        greetingHandler,
			method:         "GET",
			path:           "/greeting",
			expectedStatus: http.StatusOK,
			expectedBody:   "Hello, Guest! Welcome to the Shapes API!",
		},
		{
			name:           "Greeting Handler - GET with name",
			handler:        greetingHandler,
			method:         "GET",
			path:           "/greeting?name=TestUser",
			expectedStatus: http.StatusOK,
			expectedBody:   "Hello, TestUser! Welcome to the Shapes API!",
		},
		{
			name:           "Home Handler - Invalid method",
			handler:        homeHandler,
			method:         "POST",
			path:           "/",
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Method not allowed\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, nil)
			rr := httptest.NewRecorder()

			tt.handler.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("status code: got %v, expected %v", rr.Code, tt.expectedStatus)
			}

			body := rr.Body.String()
			if !strings.Contains(body, tt.expectedBody) && body != tt.expectedBody {
				t.Errorf("response body: got %v, expected to contain %v", body, tt.expectedBody)
			}
		})
	}
}
