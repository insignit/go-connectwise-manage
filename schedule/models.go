package schedule

// HolidayListReference is a struct representing JSON data returned from ConnectWise.
type HolidayListReference struct {
	Id   int32             `json:"id"`
	Name string            `json:"string"`
	Info map[string]string `json:"_info"`
}

// Calendar is a struct representing JSON data returned from ConnectWise.
type Calendar struct {
	Id                 int32                `json:"id"`
	Name               string               `json:"name"`
	HolidayList        HolidayListReference `json:"holidayList"`
	MondayStartTime    string               `json:"mondayStartTime"`
	MondayEndTime      string               `json:"mondayEndTime"`
	TuesdayStartTime   string               `json:"tuesdayStartTime"`
	TuesdayEndTime     string               `json:"tuesdayEndTime"`
	WednesdayStartTime string               `json:"wednesdayStartTime"`
	WednesdayEndTime   string               `json:"wednesdayEndTime"`
	ThursdayStartTime  string               `json:"thursdayStartTime"`
	ThursdayEndTime    string               `json:"thursdayEndTime"`
	FridayStartTime    string               `json:"fridayStartTime"`
	FridayEndTime      string               `json:"fridayEndTime"`
	SaturdayStartTime  string               `json:"saturdayStartTime"`
	SaturdayEndTime    string               `json:"saturdayEndTime"`
	SundayStartTime    string               `json:"sundayStartTime"`
	SundayEndTime      string               `json:"sundayEndTime"`
	Info               map[string]string    `json:"_info"`
}
