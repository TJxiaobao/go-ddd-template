package po

type TestModel struct {
	BaseModel
	Context  string `json:"context"`
	TestName string `json:"test_name"`
}

func (m *TestModel) TableName() string {
	return "test"
}
