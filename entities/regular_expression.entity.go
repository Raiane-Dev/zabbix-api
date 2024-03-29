package entities

type RegularExpressionObject struct {
	RegexpID   string `json:"regexpid,omitempty"`
	Name       string `json:"name,omitempty"`
	TestString string `json:"test_string,omitempty"`
}

type RegularExpressionCreate struct {
	Name       string       `json:"name,omitempty"`
	TestString string       `json:"test_string,omitempty"`
	Expression []Expression `json:"expression,omitempty"`
}

type RegularExpressionUpdate struct {
	RegexpID   string       `json:"regexpid,omitempty"`
	Name       string       `json:"name,omitempty"`
	TestString string       `json:"test_string,omitempty"`
	Expression []Expression `json:"expression,omitempty"`
}

type RegularExpressionGet struct {
	RegexpIDs        string `json:"regexpids,omitempty"`
	SelectExpression string `json:"selectExpression,omitempty"`
	ZabbixCommun
}

type Expression struct {
	Expression     string `json:"expression,omitempty"`
	ExpressionType int    `json:"expression_type,omitempty"`
	ExpDelimiter   string `json:"exp_delimiter,omitempty"`
	CaseSensitive  int    `json:"case_sensitive,omitempty"`
}
