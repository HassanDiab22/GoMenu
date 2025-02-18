package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

// ValidationMiddleware dynamically checks for missing fields in the request body
func ValidationMiddleware(dto interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Read the request body
		jsonData, err := io.ReadAll(c.Request.Body)
		fmt.Println("jsonData:", string(jsonData))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			c.Abort()
			return
		}

		// Reset request body so it can be read again later
		c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonData))

		// Parse JSON into a map
		var requestBody map[string]interface{}
		if err := json.Unmarshal(jsonData, &requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
			c.Abort()
			return
		}

		// Check for missing required fields
		missingFields := getMissingFields(dto, requestBody)
		if len(missingFields) > 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":          "Missing required fields",
				"missing_fields": missingFields,
			})
			c.Abort()
			return
		}

		// Store parsed request body in context for the next handlers
		c.Set("validatedData", requestBody)
		c.Next()
	}
}

// getMissingFields checks which required fields are missing from the request
func getMissingFields(dto interface{}, requestBody map[string]interface{}) []string {
	missingFields := []string{}

	// Use reflection to inspect the struct fields
	dtoType := reflect.TypeOf(dto)

	// Ensure we get the actual type if a pointer is passed
	if dtoType.Kind() == reflect.Ptr {
		dtoType = dtoType.Elem()
	}

	for i := 0; i < dtoType.NumField(); i++ {
		field := dtoType.Field(i)

		// Check for the "binding" tag with "required"
		if tag, exists := field.Tag.Lookup("binding"); exists && tag == "required" {
			jsonKey := field.Tag.Get("json") // Get JSON field name
			if jsonKey == "" {
				jsonKey = field.Name
			}

			// Check if the field is missing in the request
			if _, exists := requestBody[jsonKey]; !exists {
				missingFields = append(missingFields, jsonKey)
			}
		}
	}

	return missingFields
}
