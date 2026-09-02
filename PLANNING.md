# DockOps - Project Planning & Roadmap

Este documento lista as funcionalidades planejadas para evoluir a arquitetura e as capacidades da API DockOps, tornando-a uma plataforma robusta para gerenciamento de containers.

## 1. Ciclo de Vida Completo (CRUD de Containers)
- [ ] `GET /api/v1/containers`: Listagem de todos os containers ativos (com paginação e filtros).
- [ ] `POST /api/v1/containers/:id/stop`: Pausar um container em execução.
- [ ] `POST /api/v1/containers/:id/start`: Iniciar um container parado.
- [ ] `POST /api/v1/containers/:id/restart`: Reiniciar um container.
- [ ] `DELETE /api/v1/containers/:id`: Destruir o container e limpar recursos.

## 2. Enriquecimento do Payload (Mais poder ao usuário)
- [ ] **Port Bindings:** Permitir mapeamento de portas (ex: `80:8080`).
- [ ] **Environment Variables:** Suporte à injeção de variáveis de ambiente (`ENV_VAR=value`).
- [ ] **Resource Limits:** Limitar uso de CPU e Memória dos containers.

## 3. Observabilidade e Logs (Experiência do Desenvolvedor)
- [ ] `GET /api/v1/containers/:id/logs`: Endpoint para resgatar os logs (stdout/stderr) do container.
- [ ] **Tempo Real:** Implementar WebSockets ou SSE (Server-Sent Events) para notificar mudanças de status automaticamente, sem necessidade de polling.

## 4. Isolamento e Multi-tenancy (Segurança)
- [ ] **Ownership (Dono do Container):** Vincular a criação do container ao `UserID` extraído do JWT.
- [ ] **Autorização Restrita (OPA):** Garantir que usuários só possam alterar ou deletar os próprios containers.
- [ ] `POST /api/v1/auth/register`: Rota para registro de novos usuários.
- [ ] `GET /api/v1/users/me`: Rota para visualização do próprio perfil.

## 5. Healthcheck de Infraestrutura
- [ ] `GET /health`: Endpoint público para verificar a disponibilidade do banco de dados (PostgreSQL) e do mensageiro (RabbitMQ).
