package entities

type TemplateObject struct {
	TemplateID     string `json:"templateid"`
	Host           string `json:"host"`
	Description    string `json:"description"`
	Name           string `json:"name"`
	UUID           string `json:"uuid"`
	VendorName     string `json:"vendor_name"`
	VendorVersionn string `json:"vendor_versionn"`
}

type TemplateTag struct {
	Tag   string `json:"tag"`
	Value string `json:"value"`
}
