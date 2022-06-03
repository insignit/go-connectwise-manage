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

//* ScheduleCalendars

// GetScheduleCalendars returns all schedule calendars from ConnectWise
func (s ScheduleAPI) GetScheduleCalendars(options ...cwm.CWRequestOption) (calendars []Calendar, err error) {
	j, err := s.client.Get("/schedule/calendars", options...)
	if err != nil {
		return calendars, fmt.Errorf("Cannot retrieve schedule calendars: %v", err)
	}
	err = json.Unmarshal(j, &calendars)
	if err != nil {
		return calendars, fmt.Errorf("Can't get unmarshal data %v", err)
	}
	return
}

// PostScheduleCalendars creates a schedule calendar in ConnectWise
func (s ScheduleAPI) PostScheduleCalendars(calendar Calendar, options ...cwm.CWRequestOption) (returnCalendar Calendar, err error) {
	body, err := json.Marshal(calendar)
	if err != nil {
		return returnCalendar, fmt.Errorf("Can't marshal data %v", err)
	}
	r, err := s.client.Post("/schedule/calendars", body, options...)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(r), &returnCalendar)
	if err != nil {
		return returnCalendar, fmt.Errorf("Can't get unmarshal data %v", err)
	}
	return
}

// GetScheduleCalendarsById returns a specific schedule calendar from ConnectWise by id
func (s ScheduleAPI) GetScheduleCalendarsById(id string, options ...cwm.CWRequestOption) (calendars Calendar, err error) {
	j, err := s.client.Get(fmt.Sprintf("/schedule/calendars/%s", id), options...)
	if err != nil {
		return calendars, fmt.Errorf("Cannot retrieve schedule calendars: %v", err)
	}
	err = json.Unmarshal(j, &calendars)
	if err != nil {
		return calendars, fmt.Errorf("Can't get unmarshal data %v", err)
	}
	return
}

//* ScheduleEntries

// GetScheduleEntries returns all schedule entries from ConnectWise
func (s ScheduleAPI) GetScheduleEntries(options ...cwm.CWRequestOption) (entries []ScheduleEntry, err error) {
	j, err := s.client.Get("/schedule/entries", options...)
	if err != nil {
		return entries, fmt.Errorf("Cannot retrieve schedule entries: %v", err)
	}
	err = json.Unmarshal(j, &entries)
	if err != nil {
		return entries, fmt.Errorf("Can't get unmarshal data %v", err)
	}
	return
}

// PostScheduleEntries create a schedule entry within ConnectWise
func (s ScheduleAPI) PostScheduleEntries(entry ScheduleEntry, options ...cwm.CWRequestOption) (returnEntry ScheduleEntry, err error) {
	body, err := json.Marshal(entry)
	if err != nil {
		return returnEntry, fmt.Errorf("Can't marshal data %v", err)
	}
	r, err := s.client.Post("/schedule/entries", body, options...)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(r), &returnEntry)
	if err != nil {
		return returnEntry, fmt.Errorf("Can't get unmarshal data %v", err)
	}
	return
}

// GetScheduleEntriesById returns a specific schedule entry from ConnectWise by id
func (s ScheduleAPI) GetScheduleEntriesById(id string, options ...cwm.CWRequestOption) (entry ScheduleEntry, err error) {
	j, err := s.client.Get(fmt.Sprintf("/schedule/entries/%s", id), options...)
	if err != nil {
		return entry, fmt.Errorf("Cannot retrieve schedule calendars: %v", err)
	}
	err = json.Unmarshal(j, &entry)
	if err != nil {
		return entry, fmt.Errorf("Can't get unmarshal data %v", err)
	}
	return
}
