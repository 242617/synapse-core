# synapse core

# Systemd unit

```
[Unit]
Description=Synapse core service
Requires=docker.service
After=docker.service

[Service]
User=synapse-core
Restart=always
ExecStart=/usr/bin/docker run --rm --name core -p 8080:8080 242617/synapse-core
ExecStop=/usr/bin/docker stop core

[Install]
WantedBy=local.target
```

# Implement

* gRPC
* Concourse deploy
* https://letsencrypt.org


# Vault

## Setup

```
vault auth enable github
vault write auth/github/config organization=synapse-service
vault login -method=github token={{token}}
```

## Fill

```
vault kv put synapse/core foo=world excited=true
vault kv get synapse/core
vault kv get -field=excited synapse/core

vault kv delete synapse/core
```