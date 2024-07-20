# Godog Demo
  

## Descrição  

Esse projeto consiste em uma demonstração do uso da biblioteca Godog para validação de respostas de uma api.

## Como Executar

  Primeiramente o usuário deverá executar a aplicação contida na pasta src do projeto. Essa aplicação consiste em uma api simples a qual subirá na porta 8080 e possui apenas um endpoint de versionamento o qual poderá ser chamado através do endpoint Get /version.
	  
Após subir a aplicação	basta ir na pasta component_tests e executar o teste contido no arquivo main_test.go. O godog executa os testes utilizando a suit nativa de testes da linguagem Go.