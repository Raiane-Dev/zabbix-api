package entities

type ProxyObject struct {
	ProxyID        string `json:"proxyid"`
	Host           string `json:"host"`
	Status         int    `json:"status"`
	Description    string `json:"description"`
	LastAccess     string `json:"lastaccess"`
	TLSConnect     int    `json:"tls_connect"`
	TLSAccept      int    `json:"tls_accept"`
	TLSIssuer      string `json:"tls_issuer"`
	TLSSubject     string `json:"tls_subject"`
	TLSPSKIdentity string `json:"tls_psk_identity"`
	TLSPsk         string `json:"tls_psk"`
	ProxyAddress   string `json:"proxy_address"`
	AutoCompress   int    `json:"auto_compress"`
	Version        int    `json:"version"`
	Compatibility  int    `json:"compatibility"`
}

type ProxyCreate struct {
	Host           string                `json:"host"`
	Status         int                   `json:"status"`
	Description    string                `json:"description"`
	LastAccess     string                `json:"lastaccess"`
	TLSConnect     int                   `json:"tls_connect"`
	TLSAccept      int                   `json:"tls_accept"`
	TLSIssuer      string                `json:"tls_issuer"`
	TLSSubject     string                `json:"tls_subject"`
	TLSPSKIdentity string                `json:"tls_psk_identity"`
	TLSPsk         string                `json:"tls_psk"`
	ProxyAddress   string                `json:"proxy_address"`
	AutoCompress   int                   `json:"auto_compress"`
	Version        int                   `json:"version"`
	Compatibility  int                   `json:"compatibility"`
	Hosts          []HostObject          `json:"hosts"`
	Interface      []HostInterfaceObject `json:"interface"`
}

type ProxyUpdate struct {
	ProxyID        string                `json:"proxyid"`
	Host           string                `json:"host"`
	Status         int                   `json:"status"`
	Description    string                `json:"description"`
	LastAccess     string                `json:"lastaccess"`
	TLSConnect     int                   `json:"tls_connect"`
	TLSAccept      int                   `json:"tls_accept"`
	TLSIssuer      string                `json:"tls_issuer"`
	TLSSubject     string                `json:"tls_subject"`
	TLSPSKIdentity string                `json:"tls_psk_identity"`
	TLSPsk         string                `json:"tls_psk"`
	ProxyAddress   string                `json:"proxy_address"`
	AutoCompress   int                   `json:"auto_compress"`
	Version        int                   `json:"version"`
	Compatibility  int                   `json:"compatibility"`
	Hosts          []HostObject          `json:"hosts"`
	Interface      []HostInterfaceObject `json:"interface"`
}

type ProxyInterface struct {
	DNS   string `json:"dns"`
	IP    string `json:"ip"`
	Port  string `json:"port"`
	UseIP int    `json:"useip"`
}

type ProxyGet struct {
	ProxyIDs        []string `json:"proxyids"`
	SelectHosts     string   `json:"selectHosts"`
	SelectInterface string   `json:"selectInterface"`
	ZabbixCommun
}
