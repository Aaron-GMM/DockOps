# Resumo de Desenvolvimento - DockOps (Fase 2)

Este documento resume as implementações, decisões arquiteturais e mudanças metodológicas realizadas durante a **Fase 2: Event Sourcing e Estado Imutável**.

## 🚀 O que foi Feito

### 1. Event Sourcing e Estado Imutável
- **Abandono do CRUD Tradicional:** O sistema deixou de utilizar atualizações diretas de estado (UPDATEs) para adotar um log de eventos imutáveis.
- **Reconstrução de Estado:** Implementação da função pura `DetermineContainerState` no core, capaz de reduzir uma lista de eventos para determinar o status atual do container (ex: Pending, Running, Stopped).
- **Persistência de Eventos:** O `EventRepository` foi integrado tanto na API quanto no Worker para garantir a rastreabilidade total do ciclo de vida.

### 2. Fluxo de Eventos na Arquitetura
- **Handler (API):** Agora gera um ID de container e regista o evento `ContainerCreated` no PostgreSQL antes de enviar para a fila.
- **Worker (Consumer):** Após a execução bem-sucedida pelo provedor, o worker regista o evento `ContainerStarted`.

### 3. Geração de IDs Nativos
- **Simplificação de Dependências:** Substituição da biblioteca externa de UUID por uma função nativa no core (`GenerateID`) utilizando o pacote `crypto/rand` do Go, gerando IDs hexadecimais de 16 caracteres.

### 4. Nova Metodologia de Testes
- **Code-First com Cobertura Imediata:** Transição do TDD estrito para uma abordagem onde a funcionalidade é escrita e, imediatamente, são criados os testes unitários.
- **Padrão AAA (Arrange, Act, Assert):** Todos os novos testes seguem esta estrutura para maior clareza.
- **Nomenclatura Descritiva:** Adoção do padrão `Test{NomeFunção}_{Cenário}_{Resultado}` para facilitar a identificação de falhas.
- **Testes de Cenários de Erro:** Implementação de testes específicos para falhas de binding de JSON, erros de banco de dados e indisponibilidade do broker de mensagens.

---

## 🛠️ Desafios Enfrentados e Soluções

### 1. Rigor de Tipagem em Mocks
- **Problema:** Erro de compilação nos testes devido à incompatibilidade entre ponteiros (`*core.Event`) e valores (`core.Event`) na interface do repositório.
- **Solução:** Ajuste da assinatura nos mocks para respeitar estritamente o contrato da interface.

### 2. Validação de Argumentos no Logger
- **Problema:** Falha no `go test` devido ao uso de placeholders de formatação (`%v`) sem a passagem das variáveis correspondentes.
- **Solução:** Correção das chamadas de log para incluir os argumentos necessários, prevenindo pânicos em tempo de execução.

### 3. Falha de Correspondência no Mock (Unexpected Method Call)
- **Problema:** Pânico nos testes do Worker porque o `ResourceID` estava a chegar vazio ao mock do banco de dados.
- **Solução:** Inicialização correta dos dados de teste (payload) para garantir que os IDs simulados coincidam com as expectativas do mock.

---

## 📊 Estado Atual

- **Conformidade:** A lógica de negócio está isolada e protegida por testes que cobrem tanto caminhos de sucesso quanto de falha.
- **Versionamento:** Criada a branch `refactor/test-aaa-pattern` para isolar a mudança estrutural dos testes com commits atômicos.
- **Infraestrutura:** O arquivo `docker-compose.yml` e o `.env` estão configurados para suportar o ambiente de desenvolvimento com Postgres e RabbitMQ.

---
**Próximo Objetivo:** Iniciar a **Fase 3: Segurança Desacoplada**, focando em Middlewares, JWT e integração com OPA (Open Policy Agent).