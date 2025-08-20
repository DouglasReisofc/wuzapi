# Easyzap

<img src="static/favicon.ico" width="30"> Easyzap é uma implementação da biblioteca [@tulir/whatsmeow](https://github.com/tulir/whatsmeow) como um serviço de API REST simples, com suporte a múltiplos aparelhos e sessões simultâneas.

Whatsmeow não utiliza Puppeteer nem emuladores Android. A comunicação é direta com os servidores do WhatsApp via WebSocket, garantindo maior velocidade e menor consumo de recursos. Alterações futuras no protocolo podem exigir atualizações da biblioteca.

## :warning: Aviso

**O uso deste software em desacordo com os Termos de Serviço do WhatsApp pode resultar no banimento do número.** Utilize por sua conta e risco e evite enviar SPAM. Para usos comerciais, procure um provedor oficial da API Business.

## Endpoints disponíveis

* **Sessão:** conectar, desconectar e fazer logout do WhatsApp. Obter status e QR Codes.
* **Mensagens:** enviar textos, imagens, áudios, documentos, templates, vídeos, figurinhas, localização, contatos e enquetes.
* **Usuários:** verificar se números possuem WhatsApp, obter informações e avatares, e listar contatos.
* **Chat:** definir presença (digitando, gravando), marcar mensagens como lidas, baixar mídias e enviar reações.
* **Grupos:** criar, remover e listar grupos, alterar nome, foto, participantes e obter link de convite.
* **Webhooks:** definir e consultar webhooks para receber eventos e mensagens.

## Pré-requisitos

**Obrigatório:**
* Go (Linguagem Go)

**Opcional:**
* Docker (para conteinerização)

## Atualizando dependências

Este projeto usa a biblioteca whatsmeow. Para atualizá-la:

```bash
go get -u go.mau.fi/whatsmeow@latest
go mod tidy
```

## Compilação

```bash
go build .
```

## Execução

Por padrão o serviço REST inicia na porta 8080. Parâmetros principais:

* `-admintoken` : token de autenticação para endpoints de administração (padrão: variável de ambiente).
* `-address` : endereço IP para escutar (padrão `0.0.0.0`).
* `-port` : porta (padrão `8080`).
* `-logtype` : formato de log (`console` ou `json`).
* `-color` : habilita cores nos logs de console.
* `-osname` : nome do dispositivo na conexão com o WhatsApp.
* `-skipmedia` : ignora download de mídias recebidas.
* `-wadebug` : nível de debug do whatsmeow (`INFO` ou `DEBUG`).
* `-sslcertificate` : arquivo de certificado SSL.
* `-sslprivatekey` : chave privada SSL.

Exemplos:

```bash
./easyzap -logtype=console -color=true
```

```bash
./easyzap -logtype json
```

Com fuso horário:

```bash
TZ=America/Sao_Paulo ./easyzap ...
```

## Configuração

Easyzap utiliza um arquivo `.env` para configuração. Exemplo para PostgreSQL:

```env
EASYZAP_ADMIN_TOKEN=seu_token_aqui
DB_USER=easyzap
DB_PASSWORD=easyzap
DB_NAME=easyzap
DB_HOST=localhost
DB_PORT=5432
TZ=America/Sao_Paulo
WEBHOOK_FORMAT=json
SESSION_DEVICE_NAME=Easyzap
```

Para SQLite:

```env
EASYZAP_ADMIN_TOKEN=seu_token_aqui
TZ=America/Sao_Paulo
```

Principais opções:

* `EASYZAP_ADMIN_TOKEN`: token de administração obrigatório.
* `TZ`: fuso horário (padrão: UTC).
* Parâmetros do PostgreSQL: necessários apenas ao utilizar esse banco.

## Uso

Todas as requisições devem incluir o cabeçalho `Authorization` com o token do usuário. É possível operar múltiplos usuários (números de WhatsApp) no mesmo servidor.

* Referência Swagger da API em [/api](/api)
* Página de login e leitura de QR Code em [/login](/login) (use `?token=seu_token`)

## Licença

[Licença MIT](LICENSE)

