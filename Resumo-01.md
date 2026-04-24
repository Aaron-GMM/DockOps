# Resumo de Desenvolvimento - DockOps (Fase 1)

Este documento resume as implementações, desafios e correções realizadas durante a **Fase 1: O Núcleo Plug-and-Play** do projeto DockOps.

## O que foi Feito

### 1. Arquitetura e Estrutura
- **Clean Architecture:** Implementação de uma estrutura de pastas "flat" e idiomática em Go, separando o coração da aplicação (`internal/core`) das implementações externas (`internal/api`, `internal/messaging`, `internal/worker`).
- **Contratos (Interfaces):** Definição de interfaces claras no core para `MessagePublisher` e `ContainerProvider`, garantindo o desacoplamento tecnológico.

### 2. Fluxo Assíncrono (API -> Fila)
- **API Handler:** Criação do endpoint `POST /api/v1/containers` utilizando o framework Gin.
- **Mensageria:** Implementação do `RabbitPublisher` para despacho de tarefas de orquestração.
- **Status HTTP 202 (Accepted):** A API agora retorna corretamente o status de "aceito para processamento", refletindo o design assíncrono.

### 3. Worker (Consumidor)
- **Consumer:** Implementação da lógica que escuta a fila do RabbitMQ e aciona o provedor de containers.
- **TDD (Test-Driven Development):** Desenvolvimento guiado por testes para garantir que as mensagens sejam corretamente deserializadas e processadas.

### 4. Observabilidade e Auditoria
- **Logging Estruturado:** Distribuição de logs por todo o ciclo de vida da requisição (Handler -> Publisher -> Consumer).
- **Logger Customizado:** Implementação de um wrapper sobre o pacote `log` do Go com prefixos para cada módulo.

## 🛠 Erros Enfrentados e Soluções

### 1. Quebra de Contrato de Interface
- **Problema:** A implementação do `RabbitPublisher` não coincidia com a interface `MessagePublisher` definida no core (assinatura de métodos diferente e falta do `context.Context`).
- **Solução:** Refatoração do publisher para aceitar `context.Context` e adequação dos parâmetros de envio para o RabbitMQ.

### 2. Roteamento Incorreto no RabbitMQ
- **Problema:** Tentativa de envio de mensagens usando uma variável `routingKey` inexistente no escopo da função `Publish`.
- **Solução:** Ajuste para usar a *default exchange* (vazia) e o nome da fila como a `routingKey`, garantindo que a mensagem chegue ao destino correto.

### 3. Bug de Formatação no Logger (Variadic Arguments)
- **Problema:** Logs exibindo `%!s(MISSING)` ou `%!d(MISSING)` devido ao repasse incorreto de fatias (slices) de argumentos para funções de formatação.
- **Solução:** Uso do operador de descompactação `...` (ex: `Printf(format, v...)`) nas funções do logger customizado.

---

##  Estado Atual

- **Testes Unitários:** Todos os testes em `handler_test.go` e `consumer_test.go` estão passando (**GREEN**).
- **Compilação:** O projeto compila sem erros de contrato.
- **Visibilidade:** O sistema agora possui uma trilha de auditoria clara no console, permitindo rastrear o fluxo desde a entrada na API até o processamento no Worker.

##  Observações
- A decisão de usar **GORM** com `AutoMigrate` facilitou a inicialização do banco de dados PostgreSQL após o descarte do Ariga Atlas.
- O uso de **Mocks** (Testify) permitiu validar a lógica de mensageria e execução sem depender de infraestrutura real durante o desenvolvimento.

---
**Próximo Objetivo:** Iniciar a **Fase 2: Event Sourcing**, substituindo o estado transacional por um histórico imutável de eventos na tabela `events`.
