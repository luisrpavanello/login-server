# OpenTibiaBR Login Server

Login Server e o servico HTTP/gRPC responsavel por autenticar o client e entregar a lista de mundos/personagens para conexao ao Canary. Ele foi escrito em Go e deve usar o mesmo banco de dados do servidor.

Rotas HTTP suportadas:

- `/`
- `/login`
- `/login.php`

## Estrutura Importante

| Caminho | Descricao |
| --- | --- |
| `login-server` | Binario local ja compilado. |
| `.env.example` | Modelo de variaveis de ambiente. |
| `.env` | Configuracao local usada em ambiente de desenvolvimento. |
| `src/` | Codigo-fonte Go. |
| `docker/` | Dockerfile e compose para execucao em container. |
| `logs/` | Saida padrao de logs quando `ENV_LOG_FILE` estiver configurado. |

## Requisitos

Para executar o binario:

- MySQL/MariaDB com o schema do Canary.
- Variaveis de ambiente configuradas em `.env`.

Para desenvolver ou recompilar:

- Go compativel com o `go.mod`.
- Acesso ao banco de dados usado pelo Canary.

## Configuracao

Crie o arquivo `.env` a partir do modelo:

```bash
cp .env.example .env
```

Edite os valores principais:

```env
SERVER_PATH=/Users/luispavanello/Dev/ProjectOT/canary

MYSQL_DBNAME=canary_db
MYSQL_HOST=127.0.0.1
MYSQL_PORT=3306
MYSQL_USER=canary
MYSQL_PASS=canary123

LOGIN_IP=127.0.0.1
LOGIN_HTTP_PORT=8088
LOGIN_GRPC_PORT=9090

SERVER_IP=127.0.0.1
SERVER_NAME=OTServBR-Global
SERVER_PORT=7172
SERVER_LOCATION=BRA
```

Campos importantes:

| Variavel | Uso |
| --- | --- |
| `MYSQL_*` | Credenciais do banco compartilhado com o Canary. |
| `LOGIN_HTTP_PORT` | Porta HTTP usada pelo OTClient para login web. |
| `LOGIN_GRPC_PORT` | Porta interna gRPC. |
| `SERVER_IP` | IP enviado ao client para conectar no mundo. |
| `SERVER_PORT` | Porta do game server, normalmente `7172`. |
| `SERVER_NAME` | Nome do mundo exibido/retornado no login. |
| `RATE_LIMITER_RATE` e `RATE_LIMITER_BURST` | Controle de requisicoes por usuario/IP. |

## Como Iniciar

Na raiz deste diretorio:

```bash
chmod +x ./login-server
./login-server
```

Com a configuracao acima, o endpoint local fica em:

```text
http://127.0.0.1:8088/login
```

O OTClient deste workspace ja possui uma entrada local em `init.lua` apontando para `127.0.0.1` na porta `8088`.

## Executar pelo Codigo-Fonte

```bash
go run ./src
```

## Build

```bash
go build -o login-server ./src
```

## Testes

```bash
go test ./...
```

## Docker

Para construir e iniciar pelo compose do projeto:

```bash
cd docker
docker compose up -d --build
```

O compose local expoe:

| Porta | Uso |
| --- | --- |
| `80` | HTTP login. |
| `9090` | gRPC. |

Tambem existe imagem publica:

```bash
docker pull opentibiabr/login-server:latest
```

## Integracao com Canary e OTClient

Fluxo recomendado:

1. Inicie o banco e o Canary.
2. Confirme que o Canary esta aceitando conexoes na porta `7172`.
3. Configure o `login-server` com o mesmo banco.
4. Inicie `./login-server`.
5. No OTClient, use o host `127.0.0.1` e a porta HTTP configurada, por exemplo `8088`.

## Solucao de Problemas

| Problema | Verificacao |
| --- | --- |
| Login retorna erro de banco | Confira `MYSQL_HOST`, `MYSQL_DBNAME`, usuario, senha e permissoes. |
| Client autentica mas nao entra no mundo | Confira `SERVER_IP`, `SERVER_PORT` e se o Canary esta online. |
| Porta ocupada | Altere `LOGIN_HTTP_PORT` ou finalize o processo usando a porta. |
| `.env` nao carregado | Execute o binario a partir da raiz do diretorio `login-server`. |

## Licenca

Consulte `LICENSE` para detalhes de licenciamento.
