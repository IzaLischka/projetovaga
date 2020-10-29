# Teste Vaga

## Instruction

* In the terminal inside the folder where you downloaded the project you should run:

`docker-compose build && docker-compose up`

## Usage

Access database from pgadmin

localhost:5050

Steps 

* Access with credentials 

      user: testevaga@gmail.com
      password: testevaga

* Create new server

* Add infos on connection tab

      host: db_postgres
      port: 5432
      username: iza
      password:testevaga

* Execute query

`SELECT * FROM report;`

## Database documentation

* Create table

```sql
("CREATE TABLE " + database + `
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
            loja_mais_frequente_valida VARCHAR(50), 
            loja_ultima_compra VARCHAR(50),
            loja_ultima_compra_valida VARCHAR(50));`)
```

* Create index 

```sql
"CREATE INDEX index_loja_mais_frequente ON " + table + " USING hash(loja_mais_frequente);
"CREATE INDEX index_loja_ultima_compra ON " + table + " USING hash(loja_ultima_compra );
```

* Inserts

```sql
`INSERT INTO report 
  (
		cpf
		cpf_valido
		private
		incompleto
		data_ultima_compra
		ticket_medio
		ticket_ultima_compra
		loja_mais_frequente
		loja_mais_frequente_valida
		loja_ultima_compra
		loja_ultima_compra_valida);` 
```

## Libraries

* "github.com/lib/pq" - connection postgres
* "github.com/Nhanderu/brdoc" - cpf/cnpj validation