/*
 * Uim channel adaptor api
 *
 * uim channel adaptor api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package channeladaptor

// InlineResponse2001 struct for InlineResponse2001
type InlineResponse2001 struct {
	Code int32     `json:"code,omitempty"`
	Msg  string    `json:"msg,omitempty"`
	Data []Contact `json:"data,omitempty"`
}
