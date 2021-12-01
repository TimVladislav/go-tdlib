// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/TimVladislav/go-tdlib/tdlib"
)

// GetMessageLinkInfo Returns information about a public or private message link
// @param uRL The message link in the format "https://t.me/c/...", or "tg://privatepost?...", or "https://t.me/username/...", or "tg://resolve?..."
func (client *Client) GetMessageLinkInfo(uRL string) (*tdlib.MessageLinkInfo, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "getMessageLinkInfo",
		"url":   uRL,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageLinkInfo tdlib.MessageLinkInfo
	err = json.Unmarshal(result.Raw, &messageLinkInfo)
	return &messageLinkInfo, err

}
