package schedule

type HolidayListReference struct {
	Id   int32             `json:"id"`
	Name string            `json:"string"`
	Info map[string]string `json:"_info"`
}

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

type CalendarInfo struct {
	Id   int32             `json:"id"`
	Name string            `json:"name"`
	Info map[string]string `json:"_info"`
}

type MemberReference struct {
	Id         int32             `json:"id"`
	Identifier string            `json:"identifier"`
	Name       string            `json:"name"`
	Info       map[string]string `json:"_info"`
}

type ServiceLocationReference struct {
	Id   int32             `json:"id"`
	Name string            `json:"name"`
	Info map[string]string `json:"_info"`
}

type ReminderReference struct {
	Id   int32             `json:"id"`
	Name string            `json:"name"`
	Info map[string]string `json:"_info"`
}

type ScheduleStatusReference struct {
	Id   int32             `json:"id"`
	Name string            `json:"name"`
	Info map[string]string `json:"_info"`
}

type ScheduleTypeReference struct {
	Id   int32             `json:"id"`
	Name string            `json:"name"`
	Info map[string]string `json:"_info"`
}

type ScheduleSpanReference struct {
	Id         int32             `json:"id"`
	Identifier string            `json:"identifier"`
	Info       map[string]string `json:"_info"`
}

type ScheduleEntry struct {
	Id                         int32                    `json:"id"`
	ObjectId                   int32                    `json:"objectId"`
	Name                       string                   `json:"name"`
	Member                     MemberReference          `json:"member"`
	Where                      ServiceLocationReference `json:"where"`
	DateStart                  string                   `json:"dateStart"`
	DateEnd                    string                   `json:"dateEnd"`
	Reminder                   ReminderReference        `json:"reminder"`
	Status                     ScheduleStatusReference  `json:"status"`
	Type                       ScheduleTypeReference    `json:"type"`
	Span                       ScheduleSpanReference    `json:"span"`
	DoneFlag                   bool                     `json:"doneFlag"`
	AcknowledgedFlag           bool                     `json:"acknowledgedFlag"`
	OwnerFlag                  bool                     `json:"ownerFlag"`
	MeetingFlag                bool                     `json:"meetingFlag"`
	AllowScheduleConflictsFlag bool                     `json:"allowScheduleConflictsFlag"`
	AddMemberToProjectFlag     bool                     `json:"addMemberToProjectFlag"`
	ProjectRoleId              int32                    `json:"projectRoleId"`
	MobileGuid                 string                   `json:"mobileGuid"`
	AcknowledgedDate           string                   `json:"acknowledgedDate"`
	CloseDate                  string                   `json:"closeDate"`
	Hours                      float32                  `json:"hours"`
	Info                       map[string]string        `json:"_info"`
}

type Usage struct {
	Type        string `json:"type"`
	Count       int32  `json:"count"`
	Id          int32  `json:"id"`
	Description string `json:"description"`
	Hyperlink   string `json:"hyperlink"`
	TypeKey     string `json:"typeKey"`
}

type ScheduleColor struct {
	Id           int32             `json:"id"`
	StartPercent int32             `json:"startPercent"`
	EndPercent   int32             `json:"enPercent"`
	Color        string            `json:"color"`
	Info         map[string]string `json:"_info"`
}

type Count struct {
	Count int32 `json:"count"`
}
