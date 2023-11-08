# Serviço de Notificações

Este projeto é um serviço de notificações simples construído com Go, utilizando a biblioteca chi para roteamento HTTP. Ele permite criar notificações, recuperá-las por ID e listar todas as notificações existentes.

## Começando

Para começar a usar o Serviço de Notificações, você precisa ter Go instalado em sua máquina. Consulte [a documentação oficial do Go](https://golang.org/doc/install) para obter instruções sobre como instalar Go.

## Instalação

Clone o repositório para a sua máquina local usando:

```bash
    git clone https://github.com/esdrasspessoa/boNotification.git
    cd boNotification
```

Instale todas as dependências necessárias:
```bash
    go mod download
```

## Executando o serviço

Para iniciar o serviço, execute:

```bash
    go run ./cmd/api
```

Alternativamente, para um desenvolvimento mais ágil com recarregamento automático, considere usar ferramentas como [Air](https://github.com/cosmtrek/air).

Por padrão, o servidor é iniciado na porta 8080.

## Uso 

A API está acessível via HTTP com as seguintes rotas:

- POST /notifications - Cria uma notificação.
- GET /notifications - Lista todas as notificações.
- GET /notifications/{id} - Recupera uma notificação específica pelo ID.

Utilize ferramentas como cURL, Postman ou uma extensão de cliente REST em seu editor de texto para interagir com a API. Exemplos de requisições estão disponíveis no diretório docs.

## Contribuindo

Se você quiser contribuir para o projeto, por favor, faça um fork do repositório e use uma feature branch. Pull requests são bem-vindos.

## Contato
- Esdras Pessoa - [@esdrasspessoa](esdrassantos41@gmail.com)
- Link do Projeto: https://github.com/esdrasspessoa/boNotification

