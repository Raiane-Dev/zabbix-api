package entities

type AuthenticationGet struct {
	Output string `json:"output"`
}

type AuthenticationOutput struct {
	AuthenticationType   string `json:"authentication_type,omitempty"`
	HTTPAuthEnabled      string `json:"http_auth_enabled,omitempty"`
	HTTPLoginForm        string `json:"http_login_form,omitempty"`
	HTTPStripDomains     string `json:"http_strip_domains,omitempty"`
	HTTPCaseSensitive    string `json:"http_case_sensitive,omitempty"`
	LDAPAuthEnabled      string `json:"ldap_auth_enabled,omitempty"`
	LDAPCaseSensitive    string `json:"ldap_case_sensitive,omitempty"`
	LDAPUserDirectoryID  string `json:"ldap_userdirectoryid,omitempty"`
	SAMLAuthEnabled      string `json:"saml_auth_enabled,omitempty"`
	SAMLCaseSensitive    string `json:"saml_case_sensitive,omitempty"`
	PasswordMinLength    string `json:"passwd_min_length,omitempty"`
	PasswordCheckRules   string `json:"passwd_check_rules,omitempty"`
	JITProvisionInterval string `json:"jit_provision_interval,omitempty"`
	SAMLJITStatus        string `json:"saml_jit_status,omitempty"`
	LDAPJITStatus        string `json:"ldap_jit_status,omitempty"`
	DisabledUserGroupID  string `json:"disabled_usrgrpid,omitempty"`
}

type AuthenticationUpdate struct {
	AuthenticationType  string `json:"authentication_type,omitempty"`
	HTTPAuthEnabled     string `json:"http_auth_enabled,omitempty"`
	HTTPLoginForm       string `json:"http_login_form,omitempty"`
	HTTPStripDomains    string `json:"http_strip_domains,omitempty"`
	HTTPCaseSensitive   string `json:"http_case_sensitive,omitempty"`
	LDAPAuthEnabled     string `json:"ldap_auth_enabled,omitempty"`
	LDAPCaseSensitive   string `json:"ldap_case_sensitive,omitempty"`
	LDAPUserDirectoryID string `json:"ldap_userdirectoryid,omitempty"`
	SAMLAuthEnabled     string `json:"saml_auth_enabled,omitempty"`
	SAMLCaseSensitive   string `json:"saml_case_sensitive,omitempty"`
	PasswordMinLength   string `json:"passwd_min_length,omitempty"`
	PasswordCheckRules  string `json:"passwd_check_rules,omitempty"`
	SAMLJITStatus       string `json:"saml_jit_status,omitempty"`
	LDAPJITStatus       string `json:"ldap_jit_status,omitempty"`
	DisabledUserGroupID string `json:"disabled_usrgrpid,omitempty"`
}
