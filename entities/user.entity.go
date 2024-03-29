package entities

import "time"

type UserObject struct {
	UserID          string    `json:"userid,omitempty"`
	Username        string    `json:"username,omitempty"`
	Password        string    `json:"passwd,omitempty"`
	RoleID          string    `json:"roleid,omitempty"`
	AttemptClock    time.Time `json:"attempt_clock,omitempty"`
	AttemptFailed   int       `json:"attempt_failed,omitempty"`
	AttemptIP       string    `json:"attempt_ip,omitempty"`
	AutoLogin       int       `json:"autologin,omitempty"`
	AutoLogout      string    `json:"autologout,omitempty"`
	Lang            string    `json:"lang,omitempty"`
	Name            string    `json:"name,omitempty"`
	Refresh         string    `json:"refresh,omitempty"`
	RowsPerPage     int       `json:"rows_per_page,omitempty"`
	Surname         string    `json:"surname,omitempty"`
	Theme           string    `json:"theme,omitempty"`
	TSProvisioned   time.Time `json:"ts_provisioned,omitempty"`
	URL             string    `json:"url,omitempty"`
	UserDirectoryID string    `json:"userdirectoryid,omitempty"`
	TimeZone        string    `json:"timezone,omitempty"`
}

type UserMediaTypeObject struct {
	MediaTypeID string      `json:"mediatypeid,omitempty"`
	SendTo      interface{} `json:"sendto,omitempty"`
	Active      int         `json:"active,omitempty"`
	Severity    int         `json:"severity,omitempty"`
	Period      string      `json:"period,omitempty"`
}

type UserCheckAuthentication struct {
	Extend    bool   `json:"extend,omitempty"`
	SessionID string `json:"sessionid,omitempty"`
	Secret    string `json:"secret,omitempty"`
	Token     string `json:"token,omitempty"`
}

type UserCheckAuthenticationObject struct {
	AuthType      int    `json:"auth_type,omitempty"`
	DebugMode     int    `json:"debug_mode,omitempty"`
	Deprovisioned bool   `json:"deprovisioned,omitempty"`
	GUIAccess     string `json:"gui_access,omitempty"`
	Secret        string `json:"secret,omitempty"`
	SessionID     string `json:"sessionid,omitempty"`
	Type          int    `json:"type,omitempty"`
	UserIP        string `json:"userip,omitempty"`
}

type UserLogin struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	UserData string `json:"userdata,omitempty"`
}

type UserCreate struct {
	Username        string                `json:"username,omitempty"`
	Password        string                `json:"passwd,omitempty"`
	RoleID          string                `json:"roleid,omitempty"`
	AttemptClock    time.Time             `json:"attempt_clock,omitempty"`
	AttemptFailed   int                   `json:"attempt_failed,omitempty"`
	AttemptIP       string                `json:"attempt_ip,omitempty"`
	AutoLogin       int                   `json:"autologin,omitempty"`
	AutoLogout      string                `json:"autologout,omitempty"`
	Lang            string                `json:"lang,omitempty"`
	Name            string                `json:"name,omitempty"`
	Refresh         string                `json:"refresh,omitempty"`
	RowsPerPage     int                   `json:"rows_per_page,omitempty"`
	Surname         string                `json:"surname,omitempty"`
	Theme           string                `json:"theme,omitempty"`
	TSProvisioned   time.Time             `json:"ts_provisioned,omitempty"`
	URL             string                `json:"url,omitempty"`
	UserDirectoryID string                `json:"userdirectoryid,omitempty"`
	TimeZone        string                `json:"timezone,omitempty"`
	UsrGrps         []map[string]string   `json:"usrgrps,omitempty"`
	Medias          []UserMediaTypeObject `json:"medias,omitempty"`
}

type UserUpdate struct {
	UserID          string                `json:"userid,omitempty"`
	Username        string                `json:"username,omitempty"`
	Password        string                `json:"passwd,omitempty"`
	RoleID          string                `json:"roleid,omitempty"`
	AttemptClock    time.Time             `json:"attempt_clock,omitempty"`
	AttemptFailed   int                   `json:"attempt_failed,omitempty"`
	AttemptIP       string                `json:"attempt_ip,omitempty"`
	AutoLogin       int                   `json:"autologin,omitempty"`
	AutoLogout      string                `json:"autologout,omitempty"`
	Lang            string                `json:"lang,omitempty"`
	Name            string                `json:"name,omitempty"`
	Refresh         string                `json:"refresh,omitempty"`
	RowsPerPage     int                   `json:"rows_per_page,omitempty"`
	Surname         string                `json:"surname,omitempty"`
	Theme           string                `json:"theme,omitempty"`
	TSProvisioned   time.Time             `json:"ts_provisioned,omitempty"`
	URL             string                `json:"url,omitempty"`
	UserDirectoryID string                `json:"userdirectoryid,omitempty"`
	TimeZone        string                `json:"timezone,omitempty"`
	UsrGrps         []map[string]string   `json:"usrgrps,omitempty"`
	Medias          []UserMediaTypeObject `json:"medias,omitempty"`
	CurrentPasswd   string                `json:"current_passwd,omitempty"`
}

type UserResponse struct {
	UserIDs []string `json:"userids,omitempty"`
}

type UserGet struct {
	MediaIDs         []string `json:"mediaids,omitempty"`
	MediaTypeIDs     []string `json:"mediatypeids,omitempty"`
	UserIDs          []string `json:"userids,omitempty"`
	UserGroupIDs     []string `json:"usrgrpids,omitempty"`
	GetAccess        bool     `json:"getAccess,omitempty"`
	SelectMedias     bool     `json:"selectMedias,omitempty"`
	SelectMediaTypes bool     `json:"selectMediatypes,omitempty"`
	SelectUserGroups bool     `json:"selectUsrgrps,omitempty"`
	SelectRole       bool     `json:"selectRole,omitempty"`
}
