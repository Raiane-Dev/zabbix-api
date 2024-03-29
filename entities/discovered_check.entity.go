package entities

type DiscoveredCheckGet struct {
	DCheckIDs   []string `json:"dcheckids,omitempty"`
	DRuleIDs    []string `json:"druleids,omitempty"`
	DServiceIDs []string `json:"dserviceids,omitempty"`
	ZabbixCommun
}

type DiscoveredCheckOutput struct {
	DCheckID             string `json:"dcheckid,omitempty"`
	DRuleID              string `json:"druleid,omitempty"`
	Key                  string `json:"key_,omitempty"`
	Ports                string `json:"ports,omitempty"`
	SNMPCommunity        string `json:"snmp_community,omitempty"`
	SNMPv3AuthPassPhrase string `json:"snmpv3_authpassphrase,omitempty"`
	SNMPv3AuthProtocol   string `json:"snmpv3_authprotocol,omitempty"`
	SNMPv3ContextName    string `json:"snmpv3_contextname,omitempty"`
	SNMPv3PrivPassPhrase string `json:"snmpv3_privpassphrase,omitempty"`
	SNMPv3PrivProtocol   int    `json:"snmpv3_privprotocol,omitempty"`
	SNMPv3SecurityLevel  string `json:"snmpv3_securitylevel,omitempty"`
	SNMPv3SecurityName   string `json:"snmpv3_securityname,omitempty"`
	SNMPv3Type           int    `json:"snmpv3_type,omitempty"`
	Uniq                 int    `json:"uniq,omitempty"`
	HostSource           int    `json:"host_source,omitempty"`
	NameSource           int    `json:"name_source,omitempty"`
}
