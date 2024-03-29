package entities

type UserDirectoryObject struct {
	UserDirectoryID     string               `json:"userdirectoryid,omitempty"`
	IDPType             int                  `json:"idp_type,omitempty"`
	GroupName           string               `json:"group_name,omitempty"`
	UserUsername        string               `json:"user_username,omitempty"`
	UserLastname        string               `json:"user_lastname,omitempty"`
	ProvisionStatus     int                  `json:"provision_status,omitempty"`
	ProvisionGroups     []UserProvisionGroup `json:"provision_groups,omitempty"`
	ProvisionMedia      []UserProvisionMedia `json:"provision_media,omitempty"`
	Name                string               `json:"name,omitempty"`
	Host                string               `json:"host,omitempty"`
	Port                int                  `json:"port,omitempty"`
	BaseDN              string               `json:"base_dn,omitempty"`
	SearchAttribute     string               `json:"search_attribute,omitempty"`
	BindDN              string               `json:"bind_dn,omitempty"`
	BindPassword        string               `json:"bind_password,omitempty"`
	Description         string               `json:"description,omitempty"`
	GroupBaseDN         string               `json:"group_basedn,omitempty"`
	GroupFilter         string               `json:"group_filter,omitempty"`
	GroupMember         string               `json:"group_member,omitempty"`
	GroupMembership     string               `json:"group_membership,omitempty"`
	SearchFilter        string               `json:"search_filter,omitempty"`
	StartTLS            int                  `json:"start_tls,omitempty"`
	UserRefAttr         string               `json:"user_ref_attr,omitempty"`
	IDPEntityID         string               `json:"idp_entityid,omitempty"`
	SPEntityID          string               `json:"sp_entityid,omitempty"`
	UsernameAttribute   string               `json:"username_attribute,omitempty"`
	SSOURL              string               `json:"sso_url,omitempty"`
	SLOURL              string               `json:"slo_url,omitempty"`
	EncryptNameID       int                  `json:"encrypt_nameid,omitempty"`
	EncryptAssertions   int                  `json:"encrypt_assertions,omitempty"`
	NameIDFormat        string               `json:"nameid_format,omitempty"`
	SCIMStatus          int                  `json:"scim_status,omitempty"`
	SignAssertions      int                  `json:"sign_assertions,omitempty"`
	SignAuthnRequests   int                  `json:"sign_authn_requests,omitempty"`
	SignMessages        int                  `json:"sign_messages,omitempty"`
	SignLogoutRequests  int                  `json:"sign_logout_requests,omitempty"`
	SignLogoutResponses int                  `json:"sign_logout_responses,omitempty"`
}

type UserProvisionGroup struct {
	Name       string           `json:"name,omitempty"`
	RoleID     string           `json:"roleid,omitempty"`
	UserGroups []map[string]int `json:"user_groups,omitempty"`
}
type UserProvisionMedia struct {
	Name        string `json:"name,omitempty"`
	MediaTypeID string `json:"mediatypeid,omitempty"`
	Attribute   string `json:"attribute,omitempty"`
}

type UserDirectoryCreate struct {
	IDPType             int                  `json:"idp_type,omitempty"`
	GroupName           string               `json:"group_name,omitempty"`
	UserUsername        string               `json:"user_username,omitempty"`
	UserLastname        string               `json:"user_lastname,omitempty"`
	ProvisionStatus     int                  `json:"provision_status,omitempty"`
	ProvisionGroups     []UserProvisionGroup `json:"provision_groups,omitempty"`
	ProvisionMedia      []UserProvisionMedia `json:"provision_media,omitempty"`
	Name                string               `json:"name,omitempty"`
	Host                string               `json:"host,omitempty"`
	Port                int                  `json:"port,omitempty"`
	BaseDN              string               `json:"base_dn,omitempty"`
	SearchAttribute     string               `json:"search_attribute,omitempty"`
	BindDN              string               `json:"bind_dn,omitempty"`
	BindPassword        string               `json:"bind_password,omitempty"`
	Description         string               `json:"description,omitempty"`
	GroupBaseDN         string               `json:"group_basedn,omitempty"`
	GroupFilter         string               `json:"group_filter,omitempty"`
	GroupMember         string               `json:"group_member,omitempty"`
	GroupMembership     string               `json:"group_membership,omitempty"`
	SearchFilter        string               `json:"search_filter,omitempty"`
	StartTLS            int                  `json:"start_tls,omitempty"`
	UserRefAttr         string               `json:"user_ref_attr,omitempty"`
	IDPEntityID         string               `json:"idp_entityid,omitempty"`
	SPEntityID          string               `json:"sp_entityid,omitempty"`
	UsernameAttribute   string               `json:"username_attribute,omitempty"`
	SSOURL              string               `json:"sso_url,omitempty"`
	SLOURL              string               `json:"slo_url,omitempty"`
	EncryptNameID       int                  `json:"encrypt_nameid,omitempty"`
	EncryptAssertions   int                  `json:"encrypt_assertions,omitempty"`
	NameIDFormat        string               `json:"nameid_format,omitempty"`
	SCIMStatus          int                  `json:"scim_status,omitempty"`
	SignAssertions      int                  `json:"sign_assertions,omitempty"`
	SignAuthnRequests   int                  `json:"sign_authn_requests,omitempty"`
	SignMessages        int                  `json:"sign_messages,omitempty"`
	SignLogoutRequests  int                  `json:"sign_logout_requests,omitempty"`
	SignLogoutResponses int                  `json:"sign_logout_responses,omitempty"`
}

type UserDirectoryUpdate struct {
	UserDirectoryID     string               `json:"userdirectoryid,omitempty"`
	IDPType             int                  `json:"idp_type,omitempty"`
	GroupName           string               `json:"group_name,omitempty"`
	UserUsername        string               `json:"user_username,omitempty"`
	UserLastname        string               `json:"user_lastname,omitempty"`
	ProvisionStatus     int                  `json:"provision_status,omitempty"`
	ProvisionGroups     []UserProvisionGroup `json:"provision_groups,omitempty"`
	ProvisionMedia      []UserProvisionMedia `json:"provision_media,omitempty"`
	Name                string               `json:"name,omitempty"`
	Host                string               `json:"host,omitempty"`
	Port                int                  `json:"port,omitempty"`
	BaseDN              string               `json:"base_dn,omitempty"`
	SearchAttribute     string               `json:"search_attribute,omitempty"`
	BindDN              string               `json:"bind_dn,omitempty"`
	BindPassword        string               `json:"bind_password,omitempty"`
	Description         string               `json:"description,omitempty"`
	GroupBaseDN         string               `json:"group_basedn,omitempty"`
	GroupFilter         string               `json:"group_filter,omitempty"`
	GroupMember         string               `json:"group_member,omitempty"`
	GroupMembership     string               `json:"group_membership,omitempty"`
	SearchFilter        string               `json:"search_filter,omitempty"`
	StartTLS            int                  `json:"start_tls,omitempty"`
	UserRefAttr         string               `json:"user_ref_attr,omitempty"`
	IDPEntityID         string               `json:"idp_entityid,omitempty"`
	SPEntityID          string               `json:"sp_entityid,omitempty"`
	UsernameAttribute   string               `json:"username_attribute,omitempty"`
	SSOURL              string               `json:"sso_url,omitempty"`
	SLOURL              string               `json:"slo_url,omitempty"`
	EncryptNameID       int                  `json:"encrypt_nameid,omitempty"`
	EncryptAssertions   int                  `json:"encrypt_assertions,omitempty"`
	NameIDFormat        string               `json:"nameid_format,omitempty"`
	SCIMStatus          int                  `json:"scim_status,omitempty"`
	SignAssertions      int                  `json:"sign_assertions,omitempty"`
	SignAuthnRequests   int                  `json:"sign_authn_requests,omitempty"`
	SignMessages        int                  `json:"sign_messages,omitempty"`
	SignLogoutRequests  int                  `json:"sign_logout_requests,omitempty"`
	SignLogoutResponses int                  `json:"sign_logout_responses,omitempty"`
}

type UserDirectoryResponse struct {
	UserDirectoryIDs []string `json:"userdirectoryids,omitempty"`
}

type UserDirectoryGet struct {
	UserDirectoryIDs      []string `json:"userdirectoryids,omitempty"`
	SelectUsrgrps         bool     `json:"selectUsrgrps,omitempty"`
	SelectProvisionMedia  bool     `json:"selectProvisionMedia,omitempty"`
	SelectProvisionGroups bool     `json:"selectProvisionGroups,omitempty"`
	ZabbixCommun
}

type UserDirectoryTest struct {
	UserDirectoryID     string               `json:"userdirectoryid,omitempty"`
	IDPType             int                  `json:"idp_type,omitempty"`
	GroupName           string               `json:"group_name,omitempty"`
	UserUsername        string               `json:"user_username,omitempty"`
	UserLastname        string               `json:"user_lastname,omitempty"`
	ProvisionStatus     int                  `json:"provision_status,omitempty"`
	ProvisionGroups     []UserProvisionGroup `json:"provision_groups,omitempty"`
	ProvisionMedia      []UserProvisionMedia `json:"provision_media,omitempty"`
	Name                string               `json:"name,omitempty"`
	Host                string               `json:"host,omitempty"`
	Port                int                  `json:"port,omitempty"`
	BaseDN              string               `json:"base_dn,omitempty"`
	SearchAttribute     string               `json:"search_attribute,omitempty"`
	BindDN              string               `json:"bind_dn,omitempty"`
	BindPassword        string               `json:"bind_password,omitempty"`
	Description         string               `json:"description,omitempty"`
	GroupBaseDN         string               `json:"group_basedn,omitempty"`
	GroupFilter         string               `json:"group_filter,omitempty"`
	GroupMember         string               `json:"group_member,omitempty"`
	GroupMembership     string               `json:"group_membership,omitempty"`
	SearchFilter        string               `json:"search_filter,omitempty"`
	StartTLS            int                  `json:"start_tls,omitempty"`
	UserRefAttr         string               `json:"user_ref_attr,omitempty"`
	IDPEntityID         string               `json:"idp_entityid,omitempty"`
	SPEntityID          string               `json:"sp_entityid,omitempty"`
	UsernameAttribute   string               `json:"username_attribute,omitempty"`
	SSOURL              string               `json:"sso_url,omitempty"`
	SLOURL              string               `json:"slo_url,omitempty"`
	EncryptNameID       int                  `json:"encrypt_nameid,omitempty"`
	EncryptAssertions   int                  `json:"encrypt_assertions,omitempty"`
	NameIDFormat        string               `json:"nameid_format,omitempty"`
	SCIMStatus          int                  `json:"scim_status,omitempty"`
	SignAssertions      int                  `json:"sign_assertions,omitempty"`
	SignAuthnRequests   int                  `json:"sign_authn_requests,omitempty"`
	SignMessages        int                  `json:"sign_messages,omitempty"`
	SignLogoutRequests  int                  `json:"sign_logout_requests,omitempty"`
	SignLogoutResponses int                  `json:"sign_logout_responses,omitempty"`
	TestUsername        string               `json:"test_username,omitempty"`
	TestPassword        string               `json:"test_password,omitempty"`
}
