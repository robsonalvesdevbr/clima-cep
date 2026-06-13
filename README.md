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

Na pasta test temos o clima-cep.http com testes
