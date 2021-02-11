# pager-email-client

##  Description

A simple SMTP client which can read emails and push notifications

## Usage

### Locally

```bash
export DISCOVERY_URL=https://authprovider/.well-known # Your authentication provider's [OIDC discovery URL](https://auth0.com/docs/protocols/configure-applications-with-oidc-discovery)
export CLIENT_ID=emailclient # The client ID representing this client in your authentication provider
export CLIENT_SECRET=supersecret # The client secret of the aforementioned client
export EVENTS_SERVICE_URL=https://events.mydomain.io # URL to your instance of the [events service](https://github.com/tuuturu/pager-event-service)
export IMAP_SERVER_URL=imap.gmail.com:993 # The URL and port to your IMAP server
export IMAP_USERNAME=johndoe@gmail.com # Your username for the IMAP server
export IMAP_PASSWORD=evensecretter # Your password for the IMAP server

pager-email-client -f filterconfig.yaml
```

### Docker

```bash
docker run \
    -e DISCOVERY_URL=https://authprovider/.well-known \
    -e CLIENT_ID=emailclient \
    -e CLIENT_SECRET=supersecret \
    -e EVENTS_SERVICE_URL=https://events.mydomain.io \
    -e IMAP_SERVER_URL=imap.gmail.com:993 \
    -e IMAP_USERNAME=johndoe@gmail.com \
    -e IMAP_PASSWORD=evensecretter \
    docker.pkg.github.com/tuuturu/pager-email-client/pager-email-client \
    -f filterconfig.yaml
```

## Filter config

The root element `on` contains a list of filters. A filter can react one or both of two fields:

1. Subject
2. From / sender

### Schema
```yaml
on:
  - subject: regex
  - from: regex
```

### Example

```yaml
on:
  - subject: '[dD]onate'
  - from: \w+@example.com
```
