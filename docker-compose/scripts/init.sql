CREATE TABLE PRODUTOS(
	id serial primary key,
	nome varchar,
	descricao varchar,
	preco decimal,
	quantidade integer
);
INSERT INTO PRODUTOS (nome,descricao,preco,quantidade) values ('Camiseta','preta',19.99,18),('Fone','Muito Bom',99.80,5);
