# Clima por CEP — Go Expert (Cloud Run)

Sistema em Go que recebe um CEP, identifica a cidade via [ViaCEP](https://viacep.com.br/)
e retorna a temperatura atual em **Celsius, Fahrenheit e Kelvin** consultando a
[WeatherAPI](https://www.weatherapi.com/).

## URL no Google Cloud Run

```
https://<SUA-URL>.run.app
```

> Substitua pelo endereço gerado no deploy. Exemplo de chamada:
> `https://<SUA-URL>.run.app/clima/temp?cep=80050250`

## Contrato da API

Endpoint principal: **`GET /clima/temp?cep={cep}`**

### Sucesso — `200 OK`

```json
{
  "temp_C": 28.5,
  "temp_F": 83.3,
  "temp_K": 301.65
}
```

### Cenários de falha

| Cenário          | Condição                                   | Status | Mensagem               |
|------------------|--------------------------------------------|--------|------------------------|
| Formato inválido | CEP sem 8 dígitos ou com caracteres inválidos | `422`  | `invalid zipcode`      |
| Não encontrado   | CEP com formato correto, mas inexistente   | `404`  | `can not find zipcode` |

### Fórmulas de conversão

- Celsius → Fahrenheit: `F = C × 1.8 + 32`
- Celsius → Kelvin: `K = C + 273.15` *(usa 273.15 para bater com o exemplo oficial 28.5 → 301.65)*

## Endpoints

| Método | Rota                  | Descrição                                            |
|--------|-----------------------|------------------------------------------------------|
| `GET`  | `/clima/temp?cep=`    | **Principal** — retorna `temp_C`, `temp_F`, `temp_K` |
| `GET`  | `/clima?cep=`         | Retorna o payload bruto da WeatherAPI                |
| `GET`  | `/cep?cep=`           | Retorna os dados do ViaCEP (depuração)               |
| `GET`  | `/hello`              | Health check simples                                 |

## Variáveis de ambiente

| Variável          | Obrigatória | Default | Descrição                                  |
|-------------------|-------------|---------|--------------------------------------------|
| `WEATHER_API_KEY` | Não\*       | embutida | Chave da WeatherAPI                         |
| `PORT`            | Não         | `8080`  | Porta do servidor (Cloud Run injeta esta)  |

\* Há uma chave de fallback embutida para facilitar a execução local, mas recomenda-se
definir a sua própria via `WEATHER_API_KEY`.

## Rodando localmente (Go)

```bash
export WEATHER_API_KEY=sua_chave_aqui
go run .
# Server running on port 8080
```

## Rodando localmente (Docker)

```bash
# Build da imagem
docker build -t clima-cep:latest .

# Executar o container
docker run --rm -p 8080:8080 -e WEATHER_API_KEY=sua_chave_aqui clima-cep:latest
```

## Testando a API

Com a aplicação rodando em `http://localhost:8080`:

```bash
# Sucesso — 200
curl "http://localhost:8080/clima/temp?cep=80050250"
# {"temp_C":...,"temp_F":...,"temp_K":...}

# Formato inválido — 422 invalid zipcode
curl -i "http://localhost:8080/clima/temp?cep=9995025"

# Não encontrado — 404 can not find zipcode
curl -i "http://localhost:8080/clima/temp?cep=99950250"
```

Há também o arquivo [`test/clima-cep.http`](test/clima-cep.http) com requisições prontas
para usar com a extensão *REST Client* do VS Code.

## Rodando os testes automatizados

```bash
go test ./...

# com verbose e cobertura
go test ./... -v -cover
```

Os testes cobrem:

- **Conversões** de temperatura (`pkg/temp_convert_test.go`), incluindo o exemplo oficial.
- **Validação de CEP** (`internal/entity/viacep_test.go`).
- **Handlers** via `httptest` + mocks das interfaces (`internal/handlers/clima_handlers_test.go`),
  cobrindo os cenários `200`, `422` e `404`.

## Deploy no Google Cloud Run

A partir do código-fonte (o Cloud Run constrói a imagem com o `Dockerfile`):

```bash
gcloud run deploy clima-cep \
  --source . \
  --region us-central1 \
  --allow-unauthenticated \
  --set-env-vars WEATHER_API_KEY=sua_chave_aqui
```

Ao final, o `gcloud` exibe a **Service URL** — cole-a na seção
[URL no Google Cloud Run](#url-no-google-cloud-run) acima.
