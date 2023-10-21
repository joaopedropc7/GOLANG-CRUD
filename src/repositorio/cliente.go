package repositorio

import (
	"academia/src/modelos"
	"database/sql"
	"fmt"
)

type Cliente struct {
	db *sql.DB
}

func NovoRepositorioCliente(db *sql.DB) *Cliente {
	return &Cliente{db}
}

func (repositorio Cliente) CriarCliente(cliente modelos.Cliente) (uint64, error) {
	statement, erro := repositorio.db.Prepare(`
	insert into cliente (nome, email, numerotelefone) values (?, ?, ?)`)

	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(cliente.Nome, cliente.Email, cliente.NumeroTelefone)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

func (repositorio *Cliente) BuscarClientes() ([]modelos.Cliente, error) {

	linhas, erro := repositorio.db.Query(`
	select * from cliente`)
	if erro != nil {
		return []modelos.Cliente{}, erro
	}
	defer linhas.Close()

	var clientes []modelos.Cliente
	for linhas.Next() {
		var cliente modelos.Cliente

		if erro = linhas.Scan(
			&cliente.ID,
			&cliente.Nome,
			&cliente.Email,
			&cliente.NumeroTelefone,
			&cliente.CriadoEm,
		); erro != nil {
			return []modelos.Cliente{}, erro
		}

		clientes = append(clientes, cliente)
	}

	return clientes, nil
}

func (repositorio Cliente) BuscarCliente(clienteID uint) (modelos.Cliente, error) {
	linhas, erro := repositorio.db.Query(`
		select * from cliente where id = ?
	`, clienteID)
	if erro != nil {
		return modelos.Cliente{}, erro
	}
	defer linhas.Close()

	var cliente modelos.Cliente
	if linhas.Next() {
		if erro = linhas.Scan(
			&cliente.ID,
			&cliente.Nome,
			&cliente.Email,
			&cliente.NumeroTelefone,
			&cliente.CriadoEm,
		); erro != nil {
			return modelos.Cliente{}, erro
		}
	}

	return cliente, nil
}

func (repositorio Cliente) AtualizarCliente(clienteID uint64, cliente modelos.Cliente) error {

	statement, erro := repositorio.db.Prepare(`
	update cliente set nome = ?, email = ?, numerotelefone = ? where id = ?
	`)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(cliente.Nome, cliente.Email, cliente.NumeroTelefone, clienteID); erro != nil {
		return erro
	}

	return nil

}

func (repositorio Cliente) DeletarCliente(clienteID uint64) error {

	statement, erro := repositorio.db.Prepare(`
		delete from cliente where id = ?
	`)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(clienteID); erro != nil {
		return erro
	}

	return nil

}

func (repositorio Cliente) BuscaPorNome(nome string) ([]modelos.Cliente, error) {
	nome = fmt.Sprintf("%%%s%%", nome)

	fmt.Println(nome)

	linhas, erro := repositorio.db.Query("SELECT * FROM cliente WHERE nome LIKE ?;", nome)
	if erro != nil {
		return []modelos.Cliente{}, erro
	}
	defer linhas.Close()

	var clientes []modelos.Cliente
	for linhas.Next() {
		var cliente modelos.Cliente

		if erro = linhas.Scan(
			&cliente.ID,
			&cliente.Nome,
			&cliente.Email,
			&cliente.NumeroTelefone,
			&cliente.CriadoEm,
		); erro != nil {
			return []modelos.Cliente{}, erro
		}

		clientes = append(clientes, cliente)

	}

	return clientes, nil
}
