package validate

import (
	"github.com/Nhanderu/brdoc"
)

func IsValidCnpj(cnpj string) bool {
	return brdoc.IsCNPJ(cnpj)
}

func IsValidCpf(cpf string) bool {
	return brdoc.IsCPF(cpf)
}