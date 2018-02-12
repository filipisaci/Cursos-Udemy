# Trabalhando com Go Lang

## Instalações
* go get -v github.com/golang/dep/cmd/dep
* go get -v upper.io/db.v3/mysql
* go get -v github.com/flosch/pongo2
* pip install httpie mycli
* export GOPATH=/home/filipi/go/
* export PATH=$PATH:$GOPATH/bin

## Get Started
- Para iniciar o projeto com o *dep init* é preciso que o diretório esteja dentro do $GOPATH/src/meu-projeto e que o diretório esteja vazio
- *dep ensure* vai resolver as dependências, fazendo download da lib *echo*
- Para simular metodos posts através de formularios na CLI use: *http --form http://localhost:3000/v1/insert nome=Filipi email=filipi@localhost*
- 