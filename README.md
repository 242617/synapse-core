# Synapse core

Synapse is a service for distrubuted computation.

Contains of two units – one core and several crawlers. Crawlers pulls down tasks from core and do their jobs, then sends results back.

Services, technologies and frameworks may be excessive, developing for experiencing new features:

| Technology | What for?                         | Description        |
|------------|-----------------------------------|--------------------|
| Vault      | Manage secrets                    |                    |
| gRPC       | Serializing transferring data     |                    |
| OpenSSL    | Authenificate crawlers (mainly)   | Own rootCA and PKI |
| Concourse  | Build and deploy                  | CI/CD              |
| Sentry     | Alerts                            |                    |

# Setup

## Systemd unit

```
sudo adduser synapse-core
sudo adduser synapse-core docker
```

```
[Unit]
Description=Synapse core service
Requires=docker.service
After=docker.service

[Service]
User=synapse-core
EnvironmentFile=/etc/synapse/core.env
Restart=always
ExecStart=/usr/bin/docker run --rm --name core -p 50051:50051 -e TOKEN=${TOKEN} 242617/synapse-core
ExecStop=/usr/bin/docker stop core

[Install]
WantedBy=local.target
```