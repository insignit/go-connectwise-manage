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

func (s ScheduleAPI) PostScheduleCalendarByIdCopy(id string, options ...cwm.CWRequestOption) (calendars Calendar, err error) {
	j, err := s.client.Post(fmt.Sprintf("/schedule/calendars/%s/copy", id), []byte{}, options...)
	if err != nil {
		return calendars, fmt.Errorf("Cannot copying schedule calendars: %v", err)
	}
	err = json.Unmarshal([]byte(j), &calendars)
	if err != nil {
		return calendars, fmt.Errorf("Can't get unmarshal data %v", err)
	}
	return
}

func (s ScheduleAPI) GetScheduleCalendarsByIdInfo(id string, options ...cwm.CWRequestOption) (calendarInfo CalendarInfo, err error) {
	j, err := s.client.Get(fmt.Sprintf("/schedule/calendars/%s/info", id), options...)
	if err != nil {
		return calendarInfo, fmt.Errorf("Cannot retrieve schedule calendar info: %v", err)
	}
	err = json.Unmarshal(j, &calendarInfo)
	if err != nil {
		return calendarInfo, fmt.Errorf("Can't get unmarshal data %v", err)
	}
	return
}

func (s ScheduleAPI) GetScheduleCalendarsByIdUsages(id string, options ...cwm.CWRequestOption) (usage Usage, err error) {
	j, err := s.client.Get(fmt.Sprintf("/schedule/calendars/%s/usages", id), options...)
	if err != nil {
		return usage, fmt.Errorf("Cannot retrieve schedule calendar usages: %v", err)
	}
	err = json.Unmarshal(j, &usage)
	if err != nil {
		return usage, fmt.Errorf("Can't get unmarshal data %v", err)
	}
	return
}

func (s ScheduleAPI) GetScheduleCalendarsByIdUsagesList(id string, options ...cwm.CWRequestOption) (usage []Usage, err error) {
	j, err := s.client.Get(fmt.Sprintf("/schedule/calendars/%s/usages/list", id), options...)
	if err != nil {
		return usage, fmt.Errorf("Cannot retrieve schedule calendar usages list: %v", err)
	}
	err = json.Unmarshal(j, &usage)
	if err != nil {
		return usage, fmt.Errorf("Can't get unmarshal data %v", err)
	}
	return
}

func (s ScheduleAPI) GetScheduleCalendarsCount(options ...cwm.CWRequestOption) (count Count, err error) {
	j, err := s.client.Get("/schedule/calendars/count", options...)
	if err != nil {
		return count, fmt.Errorf("Cannot retrieve schedule calendars count: %v", err)
	}
	err = json.Unmarshal(j, &count)
	if err != nil {
		return count, fmt.Errorf("Can't get unmarshal data %v", err)
	}
	return
}

func (s ScheduleAPI) GetScheduleCalendarsInfo(options ...cwm.CWRequestOption) (calendarInfo []CalendarInfo, err error) {
	j, err := s.client.Get("/schedule/calendars/info", options...)
	if err != nil {
		return calendarInfo, fmt.Errorf("Cannot retrieve schedule calendar info: %v", err)
	}
	err = json.Unmarshal(j, &calendarInfo)
	if err != nil {
		return calendarInfo, fmt.Errorf("Can't get unmarshal data %v", err)
	}
	return
}

func (s ScheduleAPI) GetScheduleCalendarsInfoCount(options ...cwm.CWRequestOption) (count Count, err error) {
	j, err := s.client.Get("/schedule/calendars/info/count", options...)
	if err != nil {
		return count, fmt.Errorf("Cannot retrieve schedule calendar info: %v", err)
	}
	err = json.Unmarshal(j, &count)
	if err != nil {
		return count, fmt.Errorf("Can't get unmarshal data %v", err)
	}
	return
}

//* ScheduleColors

func (s ScheduleAPI) GetScheduleColors(options ...cwm.CWRequestOption) (color []ScheduleColor, err error) {
	j, err := s.client.Get("/schedule/colors", options...)
	if err != nil {
		return color, fmt.Errorf("Cannot retrieve schedule colors: %v", err)
	}
	err = json.Unmarshal(j, &color)
	if err != nil {
		return color, fmt.Errorf("Can't get unmarshal data %v", err)
	}
	return
}

func (s ScheduleAPI) GetScheduleColorsById(id string, options ...cwm.CWRequestOption) (color ScheduleColor, err error) {
	j, err := s.client.Get(fmt.Sprintf("/schedule/colors/%s", id), options...)
	if err != nil {
		return color, fmt.Errorf("Cannot retrieve schedule color by id: %v", err)
	}
	err = json.Unmarshal(j, &color)
	if err != nil {
		return color, fmt.Errorf("Can't get unmarshal data %v", err)
	}
	return
}

func (s ScheduleAPI) PostScheduleColorsByIdClear(id string, options ...cwm.CWRequestOption) (color ScheduleColor, err error) {
	j, err := s.client.Post(fmt.Sprintf("/schedule/colors/%s/clear", id), []byte{}, options...)
	if err != nil {
		return color, fmt.Errorf("Cannot clearing schedule color: %v", err)
	}
	err = json.Unmarshal([]byte(j), &color)
	if err != nil {
		return color, fmt.Errorf("Can't get unmarshal data %v", err)
	}
	return
}

//* ScheduleEntries

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

func (s ScheduleAPI) PatchScheduleEntriesById(id string, operations []cwm.PatchOperation, options ...cwm.CWRequestOption) (returnEntry ScheduleEntry, err error) {
	r, err := s.client.Patch(fmt.Sprintf("/schedule/entries/%s", id), operations, options...)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(r), &returnEntry)
	if err != nil {
		return returnEntry, fmt.Errorf("Can't get unmarshal data %v", err)
	}
	return
}
