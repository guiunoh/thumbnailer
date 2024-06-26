[![github action](https://github.com/guiunoh/thumbnailer/actions/workflows/go.yml/badge.svg)](https://github.com/guiunoh/thumbnailer/actions/workflows/go.yml)
![coverage](https://img.shields.io/endpoint?url=https://gist.githubusercontent.com/guiunoh/cb32648fb86009af712ddf269c3a49c8/raw/thumbnailer-coverage-badge.json)


* run
```base
# run redis
docker compose up -d

# run api
go run ./cmd/thumbnailer
```

* api test
```bash
http -f post http://localhost:8080/thumbnailer/api/v1/thumbnails file@$PWD/test/sample.jpg rate=RATE50 | jq -r '.id'

http -f post http://localhost:8080/thumbnailer/api/v1/thumbnails file@$PWD/test/sample.jpg rate=RATE50 \
  | jq -r '.id' \
  | xargs -I {} sh -c 'http http://localhost:8080/thumbnailer/api/v1/thumbnails/{} -o {}.png'

http http://localhost:8080$(http -f post http://localhost:8080/thumbnailer/api/v1/thumbnails file@$PWD/test/sample.jpg rate=RATE50 --headers \
  | grep -i '^Location:' \
  | awk '{print $2}' | tr -d '\r' \
) -o output.png

```
