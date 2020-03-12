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

# Deploy

* Build and export Docker image
* Transfer image to host
* Import image
* Restart service

# Implement

* https://letsencrypt.org
