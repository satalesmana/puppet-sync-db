package sendmail

import (
	config "puppet-sync-db/internal/model/config"
)

type Repo struct {
	Cfg *config.Config
}

func NewRepoHandler(cfg *config.Config) Handler {
	return &Repo{
		Cfg: cfg,
	}
}

type Handler interface {
	SendEmail(to string)
}
