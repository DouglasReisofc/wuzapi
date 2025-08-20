# Referência da API Easyzap

A API utiliza duas formas de autenticação:

1. **Token de Usuário** – para endpoints comuns use o cabeçalho `Authorization` com o token do usuário.
2. **Token de Administrador** – endpoints `/admin/**` requerem o token definido na variável `EASYZAP_ADMIN_TOKEN`.

Todas as requisições devem ter `Content-Type: application/json` e incluir o cabeçalho de autenticação apropriado.

---

## Endpoints de Administração (Gestão de Usuários)

### Listar usuários
`GET /admin/users`

Retorna todos os usuários registrados.

### Adicionar usuário
`POST /admin/users`

Cria um novo usuário. Exemplo:

```bash
curl -X POST http://localhost:8080/admin/users \
  -H "Authorization: $EASYZAP_ADMIN_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name":"usuario2","token":"token2","webhook":"https://exemplo.com/webhook","events":"Message,ReadReceipt"}'
```

### Remover usuário
`DELETE /admin/users/{id}`

Remove o usuário informado.

---

## Endpoints de Sessão

### Conectar
`POST /session/connect`

Inicia a sessão do usuário e retorna QR Code se necessário.

### Desconectar
`POST /session/disconnect`

Encerra a sessão atual.

### Status
`GET /session/status`

Retorna o estado da conexão do usuário.

---

## Endpoints de Mensagens

### Enviar texto
`POST /messages/text`

Envia uma mensagem de texto para um número ou grupo.

### Enviar mídia
`POST /messages/media`

Envia imagens, áudios, documentos ou vídeos.

### Enviar enquete
`POST /chat/send/poll`

Cria uma enquete com opções e hashes retornados na resposta.

### Descriptografar voto de enquete
`POST /chat/decrypt/poll`

Recebe os dados criptografados de um voto e devolve as opções selecionadas.

---

## Endpoints de Webhook

### Definir webhook
`POST /webhook`

Configura a URL a ser chamada quando eventos forem recebidos.

### Consultar webhook
`GET /webhook`

Retorna a URL configurada.

---

## Formato do Webhook

Eventos são enviados como JSON no corpo do `POST` para a URL definida. Exemplos de eventos: mensagens recebidas, recibos de leitura e atualizações de enquete. Campos de mídia incluem chaves e URLs necessárias para baixar e descriptografar arquivos.

---

Para detalhes completos de cada endpoint consulte a documentação Swagger disponível em [/api](/api).

