package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/loginradius/lr-cli/request"
)

type HooksResponse struct {
	Data []struct {
		ID               string    `json:"Id"`
		Appid            int       `json:"AppId"`
		Createddate      time.Time `json:"CreatedDate"`
		Lastmodifieddate time.Time `json:"LastModifiedDate"`
		Targeturl        string    `json:"TargetUrl"`
		Event            string    `json:"Event"`
		Name             string    `json:"Name"`
	} `json:"Data"`
}

func GetHooks() (*HooksResponse, error) {
	hooks := conf.AdminConsoleAPIDomain + "/integrations/webhook?"
	resp, err := request.Rest(http.MethodGet, hooks, nil, "")
	if err != nil {
		return nil, err
	}
	var resultResp HooksResponse
	err = json.Unmarshal(resp, &resultResp)
	if err != nil {
		return nil, err
	}
	return &resultResp, nil
}
