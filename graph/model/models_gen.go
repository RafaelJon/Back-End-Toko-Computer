// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Computer struct {
	ID         string  `json:"id"`
	RAM        int     `json:"ram"`
	Battery    float64 `json:"battery"`
	Proccessor string  `json:"proccessor"`
	Gpu        string  `json:"gpu"`
}

type NewComputer struct {
	RAM        int     `json:"ram"`
	Battery    float64 `json:"battery"`
	Proccessor string  `json:"proccessor"`
	Gpu        string  `json:"gpu"`
}
