package secret

import (
	"fmt"
	"os"

	"github.com/hashicorp/vault/api"

	"github.com/242617/synapse-core/config"
)

type secret struct {
	DBConnection string
	SentryDSN    string
}

func (s *secret) apply(data map[string]interface{}) error {
	for key, value := range map[string]*string{
		"sentry_dsn":    &s.SentryDSN,
		"db_connection": &s.DBConnection,
	} {
		var ok bool
		*value, ok = data[key].(string)
		if !ok {
			return fmt.Errorf("cannot convert item %q", key)
		}
	}
	return nil
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
	if err := scrt.apply(data); err != nil {
		return nil, err
	}

	return &scrt, nil
}
