# Deckctl

Deckctl is a CLI application to manage the NextCloud Deck app (https://github.com/nextcloud/deck).

## How to install

```
$ go get github.com/tcassaert/deckctl
```

## Usage

### List

List your boards

```
$ deckctl list boards
```

## Configuration

You can create a config file ~/.deckctl.yaml (or another path with --config).

```
---
endpoint: https://nextcloud.local
user: your_username
password: your_password
```

## Inspiration

Got a lot of inspiration from the inuits/12to8 CLI application on how to build a Golang CLI app.
