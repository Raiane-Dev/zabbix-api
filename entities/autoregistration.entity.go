package entities

type AutoregistrationGet struct {
	Output string `json:"output,omitempty"`
}

type AutoregistrationObject struct {
	TLSAccept int `json:"tls_accept,omitempty"`
}

type AutoregistrationUpdate struct {
	TLSAccept      int    `json:"tls_accept,omitempty"`
	TLSPSKIdentify string `json:"tls_psk_identify,omitempty"`
	TLSPSK         string `json:"tls_psk,omitempty"`
}
