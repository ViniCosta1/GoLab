# GoLab - Guia de Execução das Imagens Docker

## Arquivos Criados:
- `Dockerfile` - Para criar a imagem da aplicação Go
- `docker-compose.postgres.yml` - Para executar o PostgreSQL

## Como Usar:

### SERVIDOR 1 - PostgreSQL (Máquina que vai rodar o banco)

1. **Iniciar o PostgreSQL:**
   ```powershell
   docker-compose -f docker-compose.postgres.yml up -d
   ```

2. **Verificar se está rodando:**
   ```powershell
   docker-compose -f docker-compose.postgres.yml ps
   ```

3. **Ver logs do PostgreSQL:**
   ```powershell
   docker-compose -f docker-compose.postgres.yml logs -f
   ```

4. **Parar o PostgreSQL:**
   ```powershell
   docker-compose -f docker-compose.postgres.yml down
   ```

### SERVIDOR 2 - Aplicação Go (Máquina que vai rodar a aplicação)

1. **Criar a imagem da aplicação:**
   ```powershell
   docker build -t golab-app .
   ```

2. **Executar a aplicação (conectando ao PostgreSQL via Netbird):**
   ```powershell
   docker run -d \
     --name golab-app \
     -p 8080:8080 \
     -e DB_CONN="postgres://postgres:postgres@IP_DO_POSTGRES:5432/produtos_db?sslmode=disable" \
     golab-app
   ```

   **Substitua `IP_DO_POSTGRES` pelo IP real da máquina PostgreSQL no Netbird!**

3. **Ver logs da aplicação:**
   ```powershell
   docker logs -f golab-app
   ```

4. **Parar a aplicação:**
   ```powershell
   docker stop golab-app
   docker rm golab-app
   ```

## Configuração do Netbird:

1. **No Servidor PostgreSQL:**
   - Anote o IP que o Netbird atribuiu (exemplo: 100.64.0.10)
   
2. **No Servidor da Aplicação:**
   - Use esse IP no comando docker run na variável DB_CONN

## Testando a Aplicação:

Depois que ambos estiverem rodando:
```powershell
# Testar se a aplicação está respondendo
curl http://localhost:8080/produtos

# Criar um produto
curl -X POST http://localhost:8080/produtos -H "Content-Type: application/json" -d '{"nome":"Produto Teste","preco":29.99}'
```

## Comandos Úteis:

```powershell
# Ver todas as imagens
docker images

# Ver todos os containers rodando
docker ps

# Ver logs de um container específico
docker logs nome-do-container

# Entrar em um container
docker exec -it nome-do-container sh
```