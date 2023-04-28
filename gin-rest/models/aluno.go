package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model
	Nome string `json:"nome" validate:"nonzero"`
	CPF  string `json:"cpf" validate:"len=11, regex=^[0-9]*$"`
	RG   string `json:"rg" validate:"len=9, regex=^[0-9]*$"`
}

func ValidaDados(aluno *Aluno) error {
	if err := validator.Validate(aluno); err != nil {
		return err
	}

	return nil
}