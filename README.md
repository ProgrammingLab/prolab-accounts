# prolab-accounts
[![Build Status](https://travis-ci.com/ProgrammingLab/prolab-accounts.svg?branch=master)](https://travis-ci.com/ProgrammingLab/prolab-accounts)
[![Go Report Card](https://goreportcard.com/badge/github.com/ProgrammingLab/prolab-accounts)](https://goreportcard.com/report/github.com/ProgrammingLab/prolab-accounts)

## Requirements

- Docker

## Usage

### Configuration

```
$ cp .env.sample .env
$ cp sqlboiler.toml.sample sqlboiler.toml
```

### Schema migration for database

```
$ docker-compose up -d mysql
$ scripts/setup-db
```

### Run the backend server

```
$ docker-compose up
```

### OAuth 2.0 Authorization Code Flow
```
$ docker-compose exec hydra \
    hydra clients create \
    --endpoint http://localhost:4445 \
    --id prolab-test-client \
    --name テストクライアント \
    --secret secret \
    --grant-types authorization_code,refresh_token \
    --response-types code,id_token \
    --scope read_profile,write_profile \
    --callbacks http://127.0.0.1:5555/callback
$ docker-compose exec hydra \
    hydra token user \
    --client-id prolab-test-client \
    --client-secret secret \
    --endpoint http://localhost:4444/ \
    --port 5555 \
    --scope read_profile,write_profile
```
Go to http://127.0.0.1:5555/
