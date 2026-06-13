# Laboratório Clima-Cep

## Criar a imagem localmente para test

```bash
docker build -t robsonalves-fullcycle/clima-cep:latest .
```

## Executando o container

```bash
docker run --name clima-cep --rm -p 8080:8080 robsonalves-fullcycle/clima-cep:latest
```

## Testando a API

```bash
# Busca do cep
curl <http://localhost:8080/cep?cep=80050250>

# Busca do clima
curl http://localhost:8080/clima?cep=80050250
```
