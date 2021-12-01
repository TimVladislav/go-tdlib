// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/TimVladislav/go-tdlib/tdlib"
)

// GetLanguagePackStrings Returns strings from a language pack in the current localization target by their keys. Can be called before authorization
// @param languagePackID Language pack identifier of the strings to be returned
// @param keys Language pack keys of the strings to be returned; leave empty to request all available strings
func (client *Client) GetLanguagePackStrings(languagePackID string, keys []string) (*tdlib.LanguagePackStrings, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":            "getLanguagePackStrings",
		"language_pack_id": languagePackID,
		"keys":             keys,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var languagePackStrings tdlib.LanguagePackStrings
	err = json.Unmarshal(result.Raw, &languagePackStrings)
	return &languagePackStrings, err

}
