package models

import (
	"rest-api/packages/db"
)

type Produto struct{
	Id int
	Nome string
	Descricao string
	Preco float64
	Quantidade int
}

func BuscaTodos()  []Produto{
	db := db.ConectaComBancoDeDados()
	meuSelect,err := db.Query("Select * from produtos")
	if err != nil{
		panic(err.Error())
	}
	p := Produto{}
	produtos := []Produto{}

	for meuSelect.Next(){
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = meuSelect.Scan(&id,&nome,&descricao,&preco,&quantidade)
		if err != nil {
			panic(err.Error())	
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

	    produtos = append(produtos,p)
	}
	defer db.Close()
	return produtos;
}

func CriarNovoProduto(nome, desc string, preco float64, quantidade int){
	db := db.ConectaComBancoDeDados()
    preparoDeDados, err := db.Prepare("Insert into produtos (nome, descricao, preco, quantidade) values ($1,$2,$3,$4)")
    if err!= nil{
        panic(err.Error())
    }
	preparoDeDados.Exec(nome,desc,preco,quantidade)
    defer db.Close()
}

func Deletar(id string){
	db := db.ConectaComBancoDeDados()

	deletarOProduto, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletarOProduto.Exec(id)
	defer db.Close()
}

func EditaProduto(id string) Produto{
	db := db.ConectaComBancoDeDados()
	produto, err := db.Query("select * from produtos where id=$1",id)
	
	if err!=nil{
		panic(err.Error())
	}

	produtoRetorno := Produto{}
	
	for produto.Next() {
		var id, quantidade int
		var nome, descricao	string
		var preco float64

		err = produto.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err!=nil{
			panic(err.Error())
		}
		
		produtoRetorno.Id = id
		produtoRetorno.Nome = nome
		produtoRetorno.Descricao = descricao
		produtoRetorno.Preco = preco
		produtoRetorno.Quantidade = quantidade
	}
	defer db.Close()
	return produtoRetorno
}

func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()

	produto, err := db.Prepare("update produtos set nome=$1, descricao=$2, quantidade=$3, preco=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	produto.Exec(nome,descricao,quantidade,preco,id)	
	defer db.Close()
}