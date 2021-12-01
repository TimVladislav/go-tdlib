// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/TimVladislav/go-tdlib/tdlib"
)

// GetInlineQueryResults Sends an inline query to a bot and returns its results. Returns an error with code 502 if the bot fails to answer the query before the query timeout expires
// @param botUserID The identifier of the target bot
// @param chatID Identifier of the chat where the query was sent
// @param userLocation Location of the user, only if needed
// @param query Text of the query
// @param offset Offset of the first entry to return
func (client *Client) GetInlineQueryResults(botUserID int32, chatID int64, userLocation *tdlib.Location, query string, offset string) (*tdlib.InlineQueryResults, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":         "getInlineQueryResults",
		"bot_user_id":   botUserID,
		"chat_id":       chatID,
		"user_location": userLocation,
		"query":         query,
		"offset":        offset,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var inlineQueryResults tdlib.InlineQueryResults
	err = json.Unmarshal(result.Raw, &inlineQueryResults)
	return &inlineQueryResults, err

}
