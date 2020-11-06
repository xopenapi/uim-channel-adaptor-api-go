/*
 * Uim channel adaptor api
 *
 * uim channel adaptor api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package channeladaptor

// ContactMessageEvent struct for ContactMessageEvent
type ContactMessageEvent struct {
	// 运营号渠道Id
	ProfilePlatformUid string `json:"profilePlatformUid,omitempty"`
	// 联系人渠道Id
	ContactPlatformUid string `json:"contactPlatformUid,omitempty"`
	// 消息方向 1 接收 2 发送
	Direction int32 `json:"direction,omitempty"`
	// 消息类型
	MessageType int32 `json:"messageType,omitempty"`
	// 消息内容
	Content string `json:"content,omitempty"`
	// 渠道Id
	ServiceID     int32  `json:"serviceID,omitempty"`
	PlatformMsgId string `json:"platformMsgId,omitempty"`
	MessageTime   int64  `json:"messageTime,omitempty"`
}