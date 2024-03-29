package entities

type ImageObject struct {
	ImageID   string `json:"imageid,omitempty"`
	Name      string `json:"name,omitempty"`
	ImageType int    `json:"imagetype,omitempty"`
	Image     string `json:"image,omitempty"`
}

type ImageGet struct {
	ImageIDs    []string `json:"imageids,omitempty"`
	SysmapIDs   []string `json:"sysmapids,omitempty"`
	SelectImage string   `json:"select_image,omitempty"`
	ZabbixCommun
}

type ImageUpdate struct {
	ImageID   string `json:"imageid,omitempty"`
	Name      string `json:"name,omitempty"`
	ImageType int    `json:"imagetype,omitempty"`
	Image     string `json:"image,omitempty"`
}

type ImageCreate struct {
	Name      string `json:"name,omitempty"`
	ImageType int    `json:"imagetype,omitempty"`
	Image     string `json:"image,omitempty"`
}

type ImageResponse struct {
	ImageIDs []string `json:"imageids,omitempty"`
}
