package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s type: %d) is required", name, typ)
}

// CreateTag

type CreateTagRequest struct {
	Code  uint64  `json:"code"`
	Tag   string  `json:"tag"`
	Date  string  `json:"date"`
	Hour  string  `json:"hour"`
	Value float64 `json:"value"`
}

func (r *CreateTagRequest) validate() error {
	if r.Code <= 0 && r.Date == "" && r.Hour == "" && r.Tag == "" && r.Value <= 0 {
		return fmt.Errorf("request body is empty")
	}
	if r.Code <= 0 {
		return errParamIsRequired("role", "uint64")
	}
	if r.Tag == "" {
		return errParamIsRequired("Tag", "string")
	}
	if r.Date == "" {
		return errParamIsRequired("Date", "string")
	}
	if r.Hour == "" {
		return errParamIsRequired("Hour", "string")
	}
	if r.Code <= 0.0 {
		return errParamIsRequired("Code", "float64")
	}

	return nil
}
