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

// MemberReference is a struct representing JSON data returned from ConnectWise.
type MemberReference struct {
	Id         int32             `json:"id"`
	Identifier string            `json:"identifier"`
	Name       string            `json:"name"`
	Info       map[string]string `json:"_info"`
}

// ServiceLocationReference is a struct representing JSON data returned from ConnectWise.
type ServiceLocationReference struct {
	Id   int32             `json:"id"`
	Name string            `json:"name"`
	Info map[string]string `json:"_info"`
}

// ReminderReference is a struct representing JSON data returned from ConnectWise.
type ReminderReference struct {
	Id   int32             `json:"id"`
	Name string            `json:"name"`
	Info map[string]string `json:"_info"`
}

// ScheduleStatusReference is a struct representing JSON data returned from ConnectWise.
type ScheduleStatusReference struct {
	Id   int32             `json:"id"`
	Name string            `json:"name"`
	Info map[string]string `json:"_info"`
}

// ScheduleTypeReference is a struct representing JSON data returned from ConnectWise.
type ScheduleTypeReference struct {
	Id   int32             `json:"id"`
	Name string            `json:"name"`
	Info map[string]string `json:"_info"`
}

// ScheduleSpanReference is a struct representing JSON data returned from ConnectWise.
type ScheduleSpanReference struct {
	Id         int32             `json:"id"`
	Identifier string            `json:"identifier"`
	Info       map[string]string `json:"_info"`
}

// ScheduleEntry is a struct representing JSON data returned from ConnectWise.
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
