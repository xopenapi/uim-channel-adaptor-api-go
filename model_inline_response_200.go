/*
 * Uim channel adaptor api
 *
 * uim channel adaptor api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package channeladaptor

// InlineResponse200 struct for InlineResponse200
type InlineResponse200 struct {
	Code int32                  `json:"code,omitempty"`
	Msg  string                 `json:"msg,omitempty"`
	Data map[string]interface{} `json:"data,omitempty"`
}
