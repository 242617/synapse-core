package secret

import (
	"fmt"
	"os"

	"github.com/hashicorp/vault/api"

	"github.com/242617/synapse-core/config"
)

type secret struct {
	SentryDSN string
}

func Init(config config.VaultConfig) (*secret, error) {

	client, err := api.NewClient(&api.Config{Address: config.Address})
	if err != nil {
		return nil, err
	}
	client.SetToken(os.Getenv("TOKEN"))

	secrets, err := client.Logical().Read(config.Path)
	if err != nil {
		return nil, err
	}

	data, ok := secrets.Data["data"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("cannot convert data")
	}

	var scrt secret
	scrt.SentryDSN, ok = data["sentry_dsn"].(string)
	if !ok {
		return nil, fmt.Errorf("cannot convert item")
	}

	return &scrt, nil

}
