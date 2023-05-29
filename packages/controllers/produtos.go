package controllers

import (
	"html/template"
	"log"
	"net/http"
	"rest-api/packages/models"
	"strconv"
)

var temp =  template.Must(template.ParseGlob("templates/*.html"))


func Index(w http.ResponseWriter, r *http.Request) {	
	produtos := models.BuscaTodos()
	temp.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {	
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST"{
		nome := r.FormValue("nome")
		preco := r.FormValue("preco")
		descricao := r.FormValue("descricao")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco,64)
		if err!=nil{
			log.Println("Erro na conversão de preço: ",preco)
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err!=nil{
					log.Println("Erro na conversão de quantidade: ",quantidade)
				}
				models.CriarNovoProduto(nome,descricao,precoConvertido,quantidadeConvertida)
	}
	http.Redirect(w,r,"/",301)
}

func Delete(w http.ResponseWriter, r *http.Request){
	produtoId := r.URL.Query().Get("id")
	models.Deletar(produtoId)
	http.Redirect(w,r,"/",301)
}

func Edit(w http.ResponseWriter, r *http.Request){
	produtoId := r.URL.Query().Get("id")
	produto := models.EditaProduto(produtoId)
	temp.ExecuteTemplate(w,"Edit",produto)
}

func Update(w http.ResponseWriter, r *http.Request){
	if r.Method =="POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")


		preco,err :=strconv.ParseFloat(r.FormValue("preco"),64)
		if err != nil{
			log.Println("Erro ao converter de String para Float:",err)
		}
		id, err :=strconv.Atoi(r.FormValue("id"))
		if err != nil{
			log.Println("Erro ao converter de String para Integer:",err)	
		}		
		quantidade, err :=strconv.Atoi(r.FormValue("quantidade"))
		if err != nil{
			log.Println("Erro ao converter de String para Integer:",err)	
		}

		models.AtualizaProduto(id,nome,descricao,preco, quantidade)
	}
	http.Redirect(w,r, "/",301)
}