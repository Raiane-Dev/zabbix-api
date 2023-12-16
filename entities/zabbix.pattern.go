package entities

type ZabbixPattern[T any] struct {
	ID             int    `json:"id"`
	Host           string `json:"-"`
	Method         string `json:"method"`
	Params         T      `json:"params"`
	JsonVersion    string `json:"jsonrpc" default:"2.0"`
	Authentication string `json:"auth,omitempty"`
}

type ZabbixCommun struct {
	CountOutput            bool                `json:"countOutput,omitempty"`
	Editable               bool                `json:"editable,omitempty"`
	ExcludeSearch          bool                `json:"excludeSearch,omitempty"`
	Filter                 []map[string]string `json:"filter,omitempty"`
	Limit                  int                 `json:"limit,omitempty"`
	Output                 []string            `json:"output,omitempty"`
	PreserveKeys           bool                `json:"preserveKeys,omitempty"`
	Search                 map[string]string   `json:"search,omitempty"`
	SearchByAny            bool                `json:"searchByAny,omitempty"`
	SearchWildCardsEnabled bool                `json:"searchWildCardsEnabled,omitempty"`
	SortOrder              []string            `json:"sortorder,omitempty"`
	SortField              []string            `json:"sortfield,omitempty"`
	StartSearch            bool                `json:"startSearch,omitempty"`
}

type LoginOutput struct {
	Result string `json:"result,omitempty"`
}

type IntegrationRPC struct {
	Host          string `json:"-"`
	Authorization string `json:"authorization"`
	Params        any    `json:"params"`
}

type Request struct {
	Body  any
	Query map[string]string
}

type Body[T any] struct {
	Return struct {
		Result []T   `json:"result,omitempty"`
		Error  Error `json:"error,omitempty"`
	} `json:"-"`
}

type Error struct {
	Code    int    `json:"code"`
	Data    string `json:"data"`
	Message string `json:"message"`
}
