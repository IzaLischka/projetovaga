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
    DATABASE = "mybb"
    USER     = "mybb"
    PASSWORD = "changeme"
)

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
//faz a conexão com o banco usando as credenciais passadas
func OpenDB(database string) (*sql.DB) {
    // Initialize connection string.
     var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", HOST, USER, PASSWORD, DATABASE)

    // Verifica se a conexão foi feita com sucesso ou não.
    db, err := sql.Open("postgres", connectionString)
    checkError(err)

    err = db.Ping()
    checkError(err)
    fmt.Println("Successfully created connection to database")

    //Verifica se já existe outra base com o mesmo nome caso exista faz drop.
    _, err = db.Exec("DROP TABLE IF EXISTS " + database + ";")
    checkError(err)
    fmt.Println("Finished dropping table (if existed)")

    // Criação da tabela.
    _, err = db.Exec("CREATE TABLE " + database + ` 
        (
            id serial PRIMARY KEY, 
            cpf VARCHAR(50), 
            private BOOLEAN,
            incomplete BOOLEAN,
            lastBuyDate VARCHAR(50),
            averageTicket VARCHAR(50),
            lastBuyTicket VARCHAR(50),
            oftenStore VARCHAR(50),
            lastBuyStore VARCHAR(50));`)
    checkError(err)
    fmt.Println("Finished creating table: " + database)

    _, err = db.Exec("CREATE INDEX index_lastbuystore ON " + database + " USING hash(lastbuystore);")
    checkError(err)
    fmt.Println("Created index index_lastbuystore")

    _, err = db.Exec("CREATE INDEX index_oftenstore ON " + database + " USING hash(oftenstore);")
    checkError(err)
    fmt.Println("Created index index_oftenstore")

    return db
}
// insert na tabela os valores das colunas e seus valores.
func Insert(db *sql.DB , newItem extract.Item){
     sql_statement := `INSERT INTO report 
        (cpf, private, incomplete, lastBuyDate, averageTicket, lastBuyTicket, oftenStore, lastBuyStore) 
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`
     db.Exec(
            sql_statement,
            newItem.Cpf,
            newItem.Private,
            newItem.Incomplete,
            newItem.LastBuyDate,
            newItem.AverageTicket,
            newItem.LastBuyTicket,
            newItem.OftenStore,
            newItem.LastBuyStore,
        )
}