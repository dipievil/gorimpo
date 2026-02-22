package ports

import "github.com/LXSCA7/gorimpo/internal/config"

type ConfigManager interface {
	Get() *config.Config
}
