package entities

type RegularExpressionObject struct {
	RegexpID   string `json:"regexpid"`
	Name       string `json:"name"`
	TestString string `json:"test_string"`
}

type RegularExpressionCreate struct {
	Name       string       `json:"name"`
	TestString string       `json:"test_string"`
	Expression []Expression `json:"expression"`
}

type RegularExpressionUpdate struct {
	RegexpID   string       `json:"regexpid"`
	Name       string       `json:"name"`
	TestString string       `json:"test_string"`
	Expression []Expression `json:"expression"`
}

type RegularExpressionGet struct {
	RegexpIDs        string `json:"regexpids"`
	SelectExpression string `json:"selectExpression"`
	ZabbixCommun
}

type Expression struct {
	Expression     string `json:"expression"`
	ExpressionType int    `json:"expression_type"`
	ExpDelimiter   string `json:"exp_delimiter"`
	CaseSensitive  int    `json:"case_sensitive"`
}
