# DockOps

<div align="center">
  <img src="https://img.shields.io/badge/go-1.25-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white" alt="Go" />
  <img src="https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white" alt="Docker" />
  <img src="https://img.shields.io/badge/Rabbitmq-FF6600?style=for-the-badge&logo=rabbitmq&logoColor=white" alt="RabbitMQ" />
  <img src="https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white" alt="PostgreSQL" />
  <img src="https://img.shields.io/badge/OPA-512BD4?style=for-the-badge&logo=openpolicyagent&logoColor=white" alt="OPA" />
</div>

## Descrição

O **DockOps** é uma plataforma de orquestração de containers moderna, construída sob os princípios de Clean Architecture, processamento assíncrono e segurança **Zero Trust**. O sistema permite a criação e gerenciamento de containers Docker de forma segura e escalável.

## Principais Recursos

*   **Orquestração Assíncrona:** A API delega tarefas de infraestrutura para workers especializados através de mensageria RabbitMQ.
*   **Integração Nativa com Docker:** Gerenciamento real de containers via Docker SDK.
*   **Segurança com OPA:** Autorização granular baseada em políticas (Open Policy Agent).
*   **Autenticação JWT:** Sistema de login com senhas criptografadas (bcrypt).
*   **Observabilidade:** Rastreabilidade total através de Event Sourcing no PostgreSQL.

## Tecnologias Utilizadas

*   **Linguagem:** Go
*   **Framework Web:** Gin
*   **Mensageria:** RabbitMQ
*   **Banco de Dados:** PostgreSQL (GORM)
*   **Segurança:** JWT, bcrypt e OPA
*   **Infraestrutura:** Docker e Docker Compose

## Como Executar o Projeto

1.  Configure o arquivo `.env`:
    ```env
    DB_USER=adm
    DB_PASSWORD=dockops123
    DB_NAME=dockops
    SECRET_KEY_JWT=SuaChaveSuperSecreta
    OPA_URL=http://opa:8181/v1/data/dockops/authz/allow
    RABBITMQ_URL=amqp://guest:guest@rabbitmq:5672/
    ```
2.  Inicialize os serviços:
    ```bash
    docker-compose up --build
    ```

## Guia de Testes (Postman / curl)

### 1. Obter Token (Login)
O sistema vem pré-configurado com os seguintes usuários:
*   **Admin:** `admin` / `admin123` (Role: admin)
*   **Developer:** `dev` / `dev123` (Role: developer)
*   **Viewer:** `viewer` / `viewer123` (Role: viewer)

**Endpoint:** `POST /api/v1/auth/login`
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username": "dev", "password": "dev123"}'
```
*Copie o campo `token` da resposta.*

### 2. Criar Container (Requer Role `developer` ou `admin`)
**Endpoint:** `POST /api/v1/containers/`
```bash
curl -X POST http://localhost:8080/api/v1/containers/ \
  -H "Authorization: Bearer <SEU_TOKEN>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "meu-nginx",
    "image": "nginx:alpine",
    "command": ["nginx", "-g", "daemon off;"]
  }'
```

### 3. Consultar Status (Requer qualquer Role válida)
**Endpoint:** `GET /api/v1/containers/<ID>`
```bash
curl -X GET http://localhost:8080/api/v1/containers/<ID_RETORNADO> \
  -H "Authorization: Bearer <SEU_TOKEN>"
```

### 4. Validação de Segurança
Tente realizar um `POST` utilizando o token do usuário `viewer`. O sistema deve retornar `403 Forbidden` via OPA.

## Testes Unitários
```bash
go test -v ./...
```
