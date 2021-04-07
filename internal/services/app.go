package services

import (
	"github.com/google/wire"
	"github.com/realotz/whole/internal/server"
	"github.com/realotz/whole/internal/services/cms"
	"github.com/realotz/whole/internal/services/systems"
	"github.com/realotz/whole/internal/services/users"
)

// ProviderSet is apps providers.
var ProviderSet = wire.NewSet(
	cms.ProviderSet,
	systems.ProviderSet,
	users.ProviderSet,
	NewApps,
)

func NewApps(app *cms.Cms,
	_ *systems.Systems,
	_ *users.Users,
) *server.App {
	return app.App
}
