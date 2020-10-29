package extract

import (
    "strings"
    "strconv"
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

func sanitize(data []string) Item {
    return Item{
        Cpf: data[0],
        IsValidCpf: false,
        Private: convertStringToBool(data[1]),
        Incomplete: convertStringToBool(data[2]),
        LastBuyDate: data[3],
        AverageTicket: data[4],
        LastBuyTicket: data[5],
        OftenStore: data[6],
        IsValidOftenStore: false,
        LastBuyStore: data[7],
        IsValidLastStore: false,
    }
}

func Extract(value string) Item {
    var data = standardizeSpaces(value)
    return sanitize(data)
}