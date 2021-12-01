// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/TimVladislav/go-tdlib/tdlib"
)

// GetBasicGroup Returns information about a basic group by its identifier. This is an offline request if the current user is not a bot
// @param basicGroupID Basic group identifier
func (client *Client) GetBasicGroup(basicGroupID int32) (*tdlib.BasicGroup, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":          "getBasicGroup",
		"basic_group_id": basicGroupID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var basicGroupDummy tdlib.BasicGroup
	err = json.Unmarshal(result.Raw, &basicGroupDummy)
	return &basicGroupDummy, err

}
