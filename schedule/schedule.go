package schedule

import (
	"encoding/json"
	"fmt"

	cwm "github.com/insignit/go-connectwise-manage"
)

type ScheduleAPI struct {
	client cwm.CWClient
}

// NewScheduleAPI create a new ConnectWise Schedule API client.
func NewScheduleAPI(site string, clientId string, company string, publicKey string, privateKey string) (api ScheduleAPI, err error) {
	cwmClient, err := cwm.NewCWClient(site, clientId, company, publicKey, privateKey)
	if err != nil {
		return
	}
	api = ScheduleAPI{
		client: cwmClient,
	}
	return
}

// GetScheduleCalendars return all schedule calendars from ConnectWise
func (s ScheduleAPI) GetScheduleCalendars(options ...cwm.CWRequestOption) (calendars []Calendar, err error) {
	j, err := s.client.Get("/schedule/calendars", options...)
	if err != nil {
		return calendars, fmt.Errorf("Cannot retrieve schedule calendars: %v", err)
	}
	err = json.Unmarshal(j, &calendars)
	if err != nil {
		return calendars, fmt.Errorf("Can't get unmarshal data %v", err)
	}
	return calendars, nil
}
