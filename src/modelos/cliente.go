package modelos

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type Cliente struct {
	ID             uint64    `json:"id,omitempty"`
	Nome           string    `json:"nome,omitempty"`
	Email          string    `json:"email,omitempty"`
	NumeroTelefone string    `json:"numeroTelefone,omitempty"`
	CriadoEm       time.Time `json:"criadoEm,omitempty"`
}

func (cliente *Cliente) Preparar(etapa string) error {
	if erro := cliente.validar(etapa); erro != nil {
		return erro
	}

	if erro := cliente.formatar(etapa); erro != nil {
		return erro
	}

	return nil
}

func (cliente *Cliente) validar(etapa string) error {
	if cliente.Nome == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}

	if cliente.Email == "" {
		return errors.New("O email é obrigatório e não pode estar em branco")
	}

	if cliente.NumeroTelefone == "" {
		return errors.New("O numero de telefone é obrigatório e não pode estar em branco")
	}

	if erro := checkmail.ValidateFormat(cliente.Email); erro != nil {
		return errors.New("O e-mail inserido é inválido")
	}

	return nil
}

func (cliente *Cliente) formatar(etapa string) error {
	cliente.Nome = strings.TrimSpace(cliente.Nome)
	cliente.Email = strings.TrimSpace(cliente.Email)
	cliente.NumeroTelefone = strings.TrimSpace(cliente.NumeroTelefone)

	return nil
}
