/*
 * Uim channel adaptor api
 *
 * uim channel adaptor api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package channeladaptor

// ApiResponse struct for ApiResponse
type ApiResponse struct {
	Code int32                  `json:"code,omitempty"`
	Msg  string                 `json:"msg,omitempty"`
	Data map[string]interface{} `json:"data,omitempty"`
}
