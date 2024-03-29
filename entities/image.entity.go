package entities

type ImageObject struct {
	ImageID   string `json:"imageid"`
	Name      string `json:"name"`
	ImageType int    `json:"imagetype"`
	Image     string `json:"image"`
}

type ImageGet struct {
	ImageIDs    []string `json:"imageids"`
	SysmapIDs   []string `json:"sysmapids"`
	SelectImage string   `json:"select_image"`
	ZabbixCommun
}

type ImageUpdate struct {
	ImageID   string `json:"imageid"`
	Name      string `json:"name"`
	ImageType int    `json:"imagetype"`
	Image     string `json:"image"`
}

type ImageCreate struct {
	Name      string `json:"name"`
	ImageType int    `json:"imagetype"`
	Image     string `json:"image"`
}

type ImageResponse struct {
	ImageIDs []string `json:"imageids"`
}
