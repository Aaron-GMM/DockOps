# DockOps

<div align="center">
  <img src="https://img.shields.io/badge/go-1.25-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white" alt="Go" />
  <img src="https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white" alt="Docker" />
  <img src="https://img.shields.io/badge/Rabbitmq-FF6600?style=for-the-badge&logo=rabbitmq&logoColor=white" alt="RabbitMQ" />
  <img src="https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white" alt="PostgreSQL" />
</div>

## Descrição

O **DockOps** é uma plataforma de orquestração de containers construída sob os princípios de Clean Architecture e processamento assíncrono. O objetivo central é fornecer um núcleo plug-and-play onde a API delega tarefas complexas de infraestrutura para workers especializados através de mensageria. Isso garante um alto nível de desacoplamento, confiabilidade e facilidade de manutenção.

## Tecnologias Utilizadas

* **Linguagem:** Go 1.25
* **Framework Web:** Gin
* **Mensageria:** RabbitMQ
* **Banco de Dados:** PostgreSQL (com GORM)
* **Infraestrutura:** Docker e Docker Compose
* **Testes:** Testify (Mocks e Asserts)

## Arquitetura

A estrutura do projeto segue o padrão de Clean Architecture, isolando a regra de negócio (`core`) de detalhes de implementação (`api`, `messaging`, `storage`).

```text
DockOps/ 
├── cmd/ 
│   └── main.go                    # Ponto de entrada da aplicação (Wire-up) 
├── internal/ 
│   ├── api/ 
│   │   ├── handler/               # Controladores (Rotas HTTP) 
│   │   └── security/              # Middlewares (ex: OPA/JWT) e Tokens 
│   ├── config/ 
│   │   ├── config.go              # Leitura de variáveis de ambiente (.env) 
│   │   └── logger/                # Wrapper customizado para logs estruturados 
│   ├── core/                      # O CORAÇÃO: Entidades, Eventos e Interfaces puras 
│   ├── messaging/ 
│   │   ├── events/                # Estruturas de publicação de fila 
│   │   └── rabbitmq/              # Conexão, Producer e Consumer do RabbitMQ 
│   ├── provider/ 
│   │   └── docker/                # Implementação real (Plugin) do Docker Client SDK 
│   ├── storage/ 
│   │   └── postgres/              # Implementação do banco: Models (GORM) e Repositório 
│   ├── telemetry/ 
│   │   └── tracer.go              # Setup de Observabilidade (OpenTelemetry/Jaeger) 
│   └── worker/ 
│       ├── consumer.go            # Goroutine que escuta a fila 
│       └── factory.go             # Roteador que decide qual Provider instanciar 
├── docker-compose.yml             # Infraestrutura local (Postgres, RabbitMQ, App) 
├── Dockerfile                     # Receita de build do container da aplicação 
├── go.mod / go.sum                # Gerenciamento de dependências 
└── .env                           # Variáveis secretas (não versionado)
```

## Como Executar o Projeto

### Pré-requisitos
* Docker e Docker Compose instalados.
* Go 1.25 (caso deseje rodar a aplicação fora dos containers).

### Passos

1. Clone o repositório para sua máquina local.
2. Crie um arquivo `.env` na raiz do projeto contendo as seguintes variáveis:
   ```env
   DB_USER=adm
   DB_PASSWORD=dockops123
   DB_NAME=dockops
   SECRET_KEY_JWT=SuaChaveSuperSecreta
   DB_URL=postgres://adm:dockops123@db:5432/dockops?sslmode=disable
   userRbMQ=guest
   passRbMq=guest
   RABBITMQ_URL=amqp://guest:guest@rabbitmq:5672/
   ```
3. Inicialize os serviços com o Docker Compose:
   ```bash
   docker-compose up --build -d
   ```
4. A API estará disponível em `http://localhost:8080`.

## Como Rodar os Testes

Para executar a suíte de testes unitários do projeto, basta utilizar a ferramenta padrão da linguagem:

```bash
go test -v ./...
```
