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

List the stacks on given board

```
$ deckctl list stacks --board foo
```

List the cards on given stack and board

```
$ deckctl list cards --board foo --stack bar
```

### New

Create new board

```
$ deckctl new board --title foo --color '#000000'
```

Create new stack on given board

```
$ deckctl new stack --title foo --board bar
```

Create new card  on given board and stack

```
$ deckctl new card --title foo --board bar --stack bar
```

## Configuration

You can create a config file ~/.deckctl.yaml (or another path with --config).

```
---
endpoint: https://nextcloud.local
user: your_username
password: your_password
```

## Test with a Nextcloud container

```
podman run -d -p 8080:80 -v nextcloud:/var/www/html -v  apps:/var/www/html/custom_apps -v config:/var/www/html/config -v data:/var/www/html/data --name nextcloud nextcloud:stable-apache
```

Install the NextCloud Deck app and point the endpoint in your config file to `localhost:8080`.

## Inspiration

Got a lot of inspiration from the inuits/12to8 CLI application on how to build a Golang CLI app.
