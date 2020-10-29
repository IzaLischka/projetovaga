package clean

import (
	"strings"
)

func CleanCnpj(cnpj string) string{
	var withOutDot = strings.Replace(cnpj, ".", "", -1)
	var withOutBar = strings.Replace(withOutDot, "/", "", -1)
	var withOutDash = strings.Replace(withOutBar, "-", "", -1)
	return withOutDash
}

func CleanCpf(cpf string) string{
	var withOutDot = strings.Replace(cpf, ".", "", -1)
	var withOutDash = strings.Replace(withOutDot, "-", "", -1)
	return withOutDash
}