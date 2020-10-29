package database

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    "app/pkg/extract"
)
// inicializa uma constante com as credenciais banco
const (
    HOST     = "db"
    DATABASE = "info"
    USER     = "iza"
    PASSWORD = "testevaga"
)

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
//faz a conexão com o banco usando as credenciais passadas
func OpenDB(table string) (*sql.DB) {
    // Initialize connection string.
     var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", HOST, USER, PASSWORD, DATABASE)

    // Verifica se a conexão foi feita com sucesso ou não.
    db, err := sql.Open("postgres", connectionString)
    checkError(err)

    err = db.Ping()
    checkError(err)
    fmt.Println("Successfully created connection to database")

    //Verifica se já existe outra base com o mesmo nome caso exista faz drop.
    _, err = db.Exec("DROP TABLE IF EXISTS " + table + ";")
    checkError(err)
    fmt.Println("Finished dropping table (if existed)")

    // Criação da tabela.
    _, err = db.Exec("CREATE TABLE " + table + ` 
        (
            id serial PRIMARY KEY, 
            cpf VARCHAR(50),
            cpf_valido BOOLEAN, 
            private BOOLEAN,
            incompleto BOOLEAN,
            data_ultima_compra VARCHAR(50),
            ticket_medio VARCHAR(50),
            ticket_ultima_compra VARCHAR(50),
            loja_mais_frequente VARCHAR(50),
            loja_mais_frequente_valida BOOLEAN,
            loja_ultima_compra VARCHAR(50),
            loja_ultima_compra_valida BOOLEAN);`)
    checkError(err)
    fmt.Println("Finished creating table: " + table)

    _, err = db.Exec("CREATE INDEX index_loja_ultima_compra ON " + table + " USING hash(loja_ultima_compra);")
    checkError(err)
    fmt.Println("Created index index_loja_ultima_compra")

    _, err = db.Exec("CREATE INDEX index_loja_mais_frequente ON " + table + " USING hash(loja_mais_frequente);")
    checkError(err)
    fmt.Println("Created index index_loja_mais_frequente")

    return db
}
// insert na tabela os valores das colunas e seus valores.
func Insert(db *sql.DB , newItem extract.Item){
     sql_statement := `INSERT INTO report 
        (cpf, cpf_valido, private, incompleto, data_ultima_compra, ticket_medio, ticket_ultima_compra, loja_mais_frequente, loja_mais_frequente_valida, loja_ultima_compra, loja_ultima_compra_valida) 
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);`
     db.Exec(
            sql_statement,
            newItem.Cpf,
            newItem.IsValidCpf,
            newItem.Private,
            newItem.Incomplete,
            newItem.LastBuyDate,
            newItem.AverageTicket,
            newItem.LastBuyTicket,
            newItem.OftenStore,
            newItem.IsValidOftenStore,
            newItem.LastBuyStore,
            newItem.IsValidLastStore,
        )
}