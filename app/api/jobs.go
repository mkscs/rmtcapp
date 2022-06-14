package api

type Jobs struct {
	Name         string
	Company      string
	ContractType string
	Salary       int
}

func (j Jobs) Public() interface{} {
	return map[string]interface{}{
		"name":    j.Name,
		"company": j.Company,
	}
}
