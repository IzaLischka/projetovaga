package validate

import (
	"github.com/Nhanderu/brdoc"
)
//Verifica se o cnpj é válido
func IsValidCnpj(cnpj string) bool {
	return brdoc.IsCNPJ(cnpj)
}
//Verifica se o cpf é válido
func IsValidCpf(cpf string) bool {
	return brdoc.IsCPF(cpf)
}