package extract

import (
    "strings"
    "strconv"
    "app/pkg/clean"
    "app/pkg/validate"
)
//informa o tipo do item 
type Item struct {
    Cpf string
    IsValidCpf bool
    Private bool
    Incomplete bool
    LastBuyDate string
    AverageTicket string
    LastBuyTicket string
    OftenStore string
    IsValidOftenStore bool
    LastBuyStore string
    IsValidLastStore bool
}

func convertStringToBool(s string) bool{
    result, _ := strconv.ParseBool(s)
    return result
}

//faz o split, covertendo o espaço para;
func standardizeSpaces(s string) []string {
    return strings.Split(strings.Join(strings.Fields(s), ";"), ";")
}

//faz sepação dos dados em propriedades correspondentes.
func sanitize(data []string) Item {
    return Item{
        Cpf: clean.CleanCpf(data[0]),
        IsValidCpf: validate.IsValidCpf(data[0]),
        Private: convertStringToBool(data[1]),
        Incomplete: convertStringToBool(data[2]),
        LastBuyDate: data[3],
        AverageTicket: data[4],
        LastBuyTicket: data[5],
        OftenStore: clean.CleanCnpj(data[6]),
        IsValidOftenStore: validate.IsValidCnpj(data[6]),
        LastBuyStore: clean.CleanCnpj(data[7]),
        IsValidLastStore: validate.IsValidCnpj(data[7]),
    }
}

func Extract(value string) Item {
    var data = standardizeSpaces(value)
    return sanitize(data)
}