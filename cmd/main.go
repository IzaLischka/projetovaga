package main

import (
    "database/sql"
)

import "app/pkg/file"
import "app/pkg/extract"
import "app/pkg/database"

type callbackType func(string)

//Le uma linha por vez do arquivo  e grava no banco
func process(db *sql.DB) func(value string) {
    return func(value string) {      
        var extracted extract.Item
        extracted = extract.Extract(value)
        database.Insert(db, extracted)
    }
}
//Le o txt, e roda a func process
func main() {
    var db = database.OpenDB("report");
    file.Read("text.txt", process(db))
}