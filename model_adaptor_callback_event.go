/*
 * Uim channel adaptor api
 *
 * uim channel adaptor api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package channeladaptor

// AdaptorCallbackEvent struct for AdaptorCallbackEvent
type AdaptorCallbackEvent struct {
	Type         string                 `json:"type,omitempty"`
	ChannelKey   string                 `json:"channelKey,omitempty"`
	Token        string                 `json:"token,omitempty"`
	InnerEvent   map[string]interface{} `json:"innerEvent,omitempty"`
	EventId      string                 `json:"eventId,omitempty"`
	EventTime    int64                  `json:"eventTime,omitempty"`
	EventContext string                 `json:"eventContext,omitempty"`
}