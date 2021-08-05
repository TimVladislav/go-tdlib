// AUTOGENERATED - DO NOT EDIT

package tdlib

import (
	"encoding/json"
	"fmt"
)

// MessageLink Contains an HTTPS link to a message in a supergroup or channel
type MessageLink struct {
	tdCommon
	Link     string `json:"link"`      // Message link
	IsPublic bool   `json:"is_public"` // True, if the link will work for non-members of the chat
}

// MessageType return the string telegram-type of MessageLink
func (messageLink *MessageLink) MessageType() string {
	return "messageLink"
}

// NewMessageLink creates a new MessageLink
//
// @param link Message link
// @param isPublic True, if the link will work for non-members of the chat
func NewMessageLink(link string, isPublic bool) *MessageLink {
	messageLinkTemp := MessageLink{
		tdCommon: tdCommon{Type: "messageLink"},
		Link:     link,
		IsPublic: isPublic,
	}

	return &messageLinkTemp
}

// GetMessageLink Returns an HTTPS link to a message in a chat. Available only for already sent messages in supergroups and channels. This is an offline request
// @param chatID Identifier of the chat to which the message belongs
// @param messageID Identifier of the message
// @param forAlbum Pass true to create a link for the whole media album
// @param forComment Pass true to create a link to the message as a channel post comment, or from a message thread
func (client *Client) GetMessageLink(chatID int64, messageID int64, forAlbum bool, forComment bool) (*MessageLink, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "getMessageLink",
		"chat_id":     chatID,
		"message_id":  messageID,
		"for_album":   forAlbum,
		"for_comment": forComment,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageLink MessageLink
	err = json.Unmarshal(result.Raw, &messageLink)
	return &messageLink, err

}