package entities

type ProxyObject struct {
	ProxyID        string `json:"proxyid,omitempty"`
	Host           string `json:"host,omitempty"`
	Status         int    `json:"status,omitempty"`
	Description    string `json:"description,omitempty"`
	LastAccess     string `json:"lastaccess,omitempty"`
	TLSConnect     int    `json:"tls_connect,omitempty"`
	TLSAccept      int    `json:"tls_accept,omitempty"`
	TLSIssuer      string `json:"tls_issuer,omitempty"`
	TLSSubject     string `json:"tls_subject,omitempty"`
	TLSPSKIdentity string `json:"tls_psk_identity,omitempty"`
	TLSPsk         string `json:"tls_psk,omitempty"`
	ProxyAddress   string `json:"proxy_address,omitempty"`
	AutoCompress   int    `json:"auto_compress,omitempty"`
	Version        int    `json:"version,omitempty"`
	Compatibility  int    `json:"compatibility,omitempty"`
}

type ProxyCreate struct {
	Host           string                `json:"host,omitempty"`
	Status         int                   `json:"status,omitempty"`
	Description    string                `json:"description,omitempty"`
	LastAccess     string                `json:"lastaccess,omitempty"`
	TLSConnect     int                   `json:"tls_connect,omitempty"`
	TLSAccept      int                   `json:"tls_accept,omitempty"`
	TLSIssuer      string                `json:"tls_issuer,omitempty"`
	TLSSubject     string                `json:"tls_subject,omitempty"`
	TLSPSKIdentity string                `json:"tls_psk_identity,omitempty"`
	TLSPsk         string                `json:"tls_psk,omitempty"`
	ProxyAddress   string                `json:"proxy_address,omitempty"`
	AutoCompress   int                   `json:"auto_compress,omitempty"`
	Version        int                   `json:"version,omitempty"`
	Compatibility  int                   `json:"compatibility,omitempty"`
	Hosts          []HostObject          `json:"hosts,omitempty"`
	Interface      []HostInterfaceObject `json:"interface,omitempty"`
}

type ProxyUpdate struct {
	ProxyID        string                `json:"proxyid,omitempty"`
	Host           string                `json:"host,omitempty"`
	Status         int                   `json:"status,omitempty"`
	Description    string                `json:"description,omitempty"`
	LastAccess     string                `json:"lastaccess,omitempty"`
	TLSConnect     int                   `json:"tls_connect,omitempty"`
	TLSAccept      int                   `json:"tls_accept,omitempty"`
	TLSIssuer      string                `json:"tls_issuer,omitempty"`
	TLSSubject     string                `json:"tls_subject,omitempty"`
	TLSPSKIdentity string                `json:"tls_psk_identity,omitempty"`
	TLSPsk         string                `json:"tls_psk,omitempty"`
	ProxyAddress   string                `json:"proxy_address,omitempty"`
	AutoCompress   int                   `json:"auto_compress,omitempty"`
	Version        int                   `json:"version,omitempty"`
	Compatibility  int                   `json:"compatibility,omitempty"`
	Hosts          []HostObject          `json:"hosts,omitempty"`
	Interface      []HostInterfaceObject `json:"interface,omitempty"`
}

type ProxyInterface struct {
	DNS   string `json:"dns,omitempty"`
	IP    string `json:"ip,omitempty"`
	Port  string `json:"port,omitempty"`
	UseIP int    `json:"useip,omitempty"`
}

type ProxyGet struct {
	ProxyIDs        []string `json:"proxyids,omitempty"`
	SelectHosts     string   `json:"selectHosts,omitempty"`
	SelectInterface string   `json:"selectInterface,omitempty"`
	ZabbixCommun
}
