# slides-gospel

Escolha as músicas e gere os slides para a missa.

## Get started

To start the service exec:

```
make up
```

To create an apresentation exec:

```bash
curl -X POST localhost:3000/apresentation \
  -d '{
      "title": "Missa de domingo",
      "prayer": "1",
      "songs": [
        "tu-es-o-centro-frei-gilson"
      ]
    }'
```

```
curl localhost:3000/songs/tu-es-o-centro-frei-gilson
```

## Deploy

To deploy the service to heroku, exec:

```bash
make deploy
```

## Eucharistic Prayers

https://www.catolicoorante.com.br/oeucaristicas.html
