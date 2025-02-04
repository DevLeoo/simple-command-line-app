# Simple Command Line App

Este é um aplicativo de linha de comando simples desenvolvido em **Go**, com o objetivo de aprendizado e experimentação com a biblioteca [urfave/cli](https://github.com/urfave/cli).

## 📌 Funcionalidades

- Buscar **endereços IP** de um domínio.
- Buscar **servidores de nome (NS)** de um domínio.

## 🚀 Pré-requisitos

Antes de rodar o projeto, é necessário ter instalado:

- Go (https://go.dev/doc/install) (versão 1.16 ou superior)

## 🔧 Como Instalar

1. Clone o repositório:

   git clone https://github.com/seu-usuario/seu-repositorio.git
   cd seu-repositorio

2. Instale as dependências:

   go mod tidy

## ⚡ Como Usar

O programa possui dois comandos principais:

### 🔍 Buscar IPs de um domínio

   go run main.go ip --host=exemplo.com

### 🔍 Buscar servidores de nome (NS) de um domínio

   go run main.go server --host=exemplo.com

Caso nenhum host seja especificado, ele utilizará amazon.com.br como padrão.

## 🛠 Estrutura do Projeto

/ simple-command-line-app
│── /app
│   ├── app.go  # Implementação da lógica dos comandos
│── main.go      # Entrada do programa
│── go.mod       # Gerenciamento de dependências
│── go.sum       # Controle de versões das dependências

## 📝 Notas

- Este projeto foi feito apenas para fins de estudo e aprendizado da linguagem Go e do desenvolvimento de CLI Apps.
- A biblioteca utilizada para construção da CLI é urfave/cli (https://github.com/urfave/cli).

## 📜 Licença

Este projeto é apenas para fins educacionais e não possui uma licença formal.

---

👨‍💻 Autor: [Leonardo Candido](https://github.com/DevLeoo)

