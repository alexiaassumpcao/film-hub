# Hub de filmes

## API
- https://medium.com/swlh/building-a-restful-api-with-go-and-postgresql-494819f51810
- https://paulovitorweb.wordpress.com/2017/09/09/criando-tabelas-importando-dados-e-executando-consultas-no-postgresql/#:~:text=J%C3%A1%20no%20pgAdmin%2C%20vamos%20criar,%3E%20Databases%20%3E%20postgres%20%3E%20Schemas%20.
- https://www.youtube.com/watch?v=XAL05B5cCX4
- https://www.postgresql.org/docs/9.1/arrays.html


## exemplos de querys
CREATE TABLE users (
	user_id SERIAL ,
	user_name varchar(100)
	user_last_name varchar(100),
	user_document varchar(100)
	user_email varchar(100)
	user_password varchar(100)
	PRIMARY KEY (user_id), UNIQUE (user_document)
)

SELECT * FROM users;

INSERT INT users (<campos>) VALUES (<valores>)

CREATE TABLE users_addresses (
	addr_id SERIAL
	user_id INTERGER REFERENCES users (user_id)
	addr_street varchar(100)
	addr_number varchar(100)
	addr_complement varchar(100)
	addr_district varchar(100)
	addr_city varchar(100)
	addr_state varchar(2)
	addr_country varchar(100)
	PRIMARY KEY (addr_id)
)

SELECT * FROM users_addresses;

INSERT INTO users_addresses (user_id) VALUES 
(1);