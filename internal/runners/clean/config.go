package clean

import (
	"errors"
	"os"

	"github.com/ActiveState/cli/internal/config"
	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/logging"
	"github.com/ActiveState/cli/internal/output"
)

type Config struct {
	output  output.Outputer
	confirm confirmAble
	path    string
}

type ConfigParams struct {
	Force bool
}

func NewConfig(prime primeable) *Config {
	return newConfig(prime.Output(), prime.Prompt())
}

func newConfig(out output.Outputer, confirm confirmAble) *Config {
	return &Config{
		output:  out,
		confirm: confirm,
		path:    config.ConfigPath(),
	}
}

func (c *Config) Run(params *ConfigParams) error {
	if os.Getenv(constants.ActivatedStateEnvVarName) != "" {
		return errors.New(locale.T("err_clean_cache_activated"))
	}

	if !params.Force {
		ok, fail := c.confirm.Confirm(locale.T("confirm"), locale.T("clean_config_confirm"), false)
		if fail != nil {
			return fail.ToError()
		}
		if !ok {
			return nil
		}
	}

	logging.Debug("Removing config directory: %s", c.path)
	return removeConfig(c.path)
}
