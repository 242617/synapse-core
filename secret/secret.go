package secret

import (
	"fmt"
	"os"

	"github.com/hashicorp/vault/api"

	"github.com/242617/synapse-core/config"
)

var (
	SentryDSN string
)

func Init() error {

	client, err := api.NewClient(&api.Config{Address: config.Cfg.Services.Vault.Address})
	if err != nil {
		return err
	}
	client.SetToken(os.Getenv("TOKEN"))

	secret, err := client.Logical().Read(config.Cfg.Services.Vault.Path)
	if err != nil {
		return err
	}

	data, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("cannot convert data")
	}

	SentryDSN, ok = data["sentry_dsn"].(string)
	if !ok {
		return fmt.Errorf("cannot convert item")
	}

	return nil

}
