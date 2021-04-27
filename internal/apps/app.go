package apps

import (
	"github.com/google/wire"
	"github.com/realotz/whole/internal/apps/admin"
	"github.com/realotz/whole/internal/apps/cms"
	"github.com/realotz/whole/internal/apps/systems"
	"github.com/realotz/whole/internal/apps/users"
	"github.com/realotz/whole/internal/server"
)

// ProviderSet is apps providers.
var ProviderSet = wire.NewSet(
	cms.ProviderSet,
	systems.ProviderSet,
	users.ProviderSet,
	admin.ProviderSet,
	NewApps,
)

func NewApps(app *systems.Systems) *server.App {
	return app.App
}
