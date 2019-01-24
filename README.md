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
