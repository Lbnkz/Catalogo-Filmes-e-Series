	package interfaces

	type Filters struct {
		Filters []Filter `json:"filters"`
	}

	type Filter struct {
		Field        string `json:"field"`
		Operator     string `json:"operator"`
		Value        string `json:"value"`
		JoinOperator string `json:"joinOperator"`
	}
