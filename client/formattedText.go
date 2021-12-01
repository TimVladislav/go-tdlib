// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/TimVladislav/go-tdlib/tdlib"
)

// ParseTextEntities Parses Bold, Italic, Underline, Strikethrough, Code, Pre, PreCode, TextUrl and MentionName entities contained in the text. Can be called synchronously
// @param text The text to parse
// @param parseMode Text parse mode
func (client *Client) ParseTextEntities(text string, parseMode tdlib.TextParseMode) (*tdlib.FormattedText, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":      "parseTextEntities",
		"text":       text,
		"parse_mode": parseMode,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var formattedText tdlib.FormattedText
	err = json.Unmarshal(result.Raw, &formattedText)
	return &formattedText, err

}

// ParseMarkdown Parses Markdown entities in a human-friendly format, ignoring markup errors. Can be called synchronously
// @param text The text to parse. For example, "__italic__ ~~strikethrough~~ **bold** `code` ```pre``` __[italic__ text_url](telegram.org) __italic**bold italic__bold**"
func (client *Client) ParseMarkdown(text *tdlib.FormattedText) (*tdlib.FormattedText, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "parseMarkdown",
		"text":  text,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var formattedText tdlib.FormattedText
	err = json.Unmarshal(result.Raw, &formattedText)
	return &formattedText, err

}

// GetMarkdownText Replaces text entities with Markdown formatting in a human-friendly format. Entities that can't be represented in Markdown unambiguously are kept as is. Can be called synchronously
// @param text The text
func (client *Client) GetMarkdownText(text *tdlib.FormattedText) (*tdlib.FormattedText, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "getMarkdownText",
		"text":  text,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var formattedText tdlib.FormattedText
	err = json.Unmarshal(result.Raw, &formattedText)
	return &formattedText, err

}