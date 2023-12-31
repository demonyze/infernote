package model

type TuningName string

const (
	Standard TuningName = "Standard"
)

type Tuning struct {
	Name        TuningName `json:"name"`
	Alternative string     `json:"alternative"`
}
