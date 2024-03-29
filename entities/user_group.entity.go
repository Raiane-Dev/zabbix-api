package entities

type UserGroupObject struct {
	UsrGrpID        string `json:"usrgrpid,omitempty"`
	Name            string `json:"name,omitempty"`
	DebugMode       int    `json:"debug_mode,omitempty"`
	GuiAccess       int    `json:"gui_access,omitempty"`
	UsersStatus     int    `json:"users_status,omitempty"`
	UserDirectoryID string `json:"userdirectoryid,omitempty"`
}

type PermissionsObject struct {
	ID         string `json:"id,omitempty"`
	Permission int    `json:"permission,omitempty"`
}
type TagBasedPermission struct {
	GroupID string `json:"groupid,omitempty"`
	Tag     string `json:"tag,omitempty"`
	Value   string `json:"value,omitempty"`
}
type UserGroupCreate struct {
	Name                string              `json:"name,omitempty"`
	DebugMode           int                 `json:"debug_mode,omitempty"`
	GuiAccess           int                 `json:"gui_access,omitempty"`
	UsersStatus         int                 `json:"users_status,omitempty"`
	UserDirectoryID     string              `json:"userdirectoryid,omitempty"`
	HostGroupRights     []PermissionsObject `json:"hostgroup_rights,omitempty"`
	TemplateGroupRights []PermissionsObject `json:"templategroup_rights,omitempty"`
	TagFilters          TagBasedPermission  `json:"tag_filters,omitempty"`
	Users               []UserCreate        `json:"users,omitempty"`
}

type UserGroupUpdate struct {
	UsrGrpID            string              `json:"usrgrpid,omitempty"`
	Name                string              `json:"name,omitempty"`
	DebugMode           int                 `json:"debug_mode,omitempty"`
	GuiAccess           int                 `json:"gui_access,omitempty"`
	UsersStatus         int                 `json:"users_status,omitempty"`
	UserDirectoryID     string              `json:"userdirectoryid,omitempty"`
	HostGroupRights     []PermissionsObject `json:"hostgroup_rights,omitempty"`
	TemplateGroupRights []PermissionsObject `json:"templategroup_rights,omitempty"`
	TagFilters          TagBasedPermission  `json:"tag_filters,omitempty"`
	Users               []UserCreate        `json:"users,omitempty"`
}

type UserGroupResponse struct {
	UsrGrpIDs []string `json:"usrgrpids,omitempty"`
}

type UserGroupGet struct {
	Status                    int      `json:"status,omitempty"`
	UserIDs                   []string `json:"userids,omitempty"`
	UserGroupIDs              []string `json:"usrgrpids,omitempty"`
	SelectTagFilters          string   `json:"selectTagFilters,omitempty"`
	SelectUsers               bool     `json:"selectUsers,omitempty"`
	SelectHostGroupRights     string   `json:"selectHostGroupRights,omitempty"`
	SelectTemplateGroupRights string   `json:"selectTemplateGroupRights,omitempty"`
	LimitSelects              int      `json:"limitSelects,omitempty"`
}
