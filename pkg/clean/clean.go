package clean

import (
	"strings"
)
//Limpa os dados, nesse caso o cnpj, tirando traço, barra e ponto.
func CleanCnpj(cnpj string) string{
	var withOutDot = strings.Replace(cnpj, ".", "", -1)
	var withOutBar = strings.Replace(withOutDot, "/", "", -1)
	var withOutDash = strings.Replace(withOutBar, "-", "", -1)
	return withOutDash
}
//Limpa os dados, nesse caso o cpf, tirando o traço e ponto.
func CleanCpf(cpf string) string{
	var withOutDot = strings.Replace(cpf, ".", "", -1)
	var withOutDash = strings.Replace(withOutDot, "-", "", -1)
	return withOutDash
}