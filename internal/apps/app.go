package apps

import (
	"github.com/google/wire"
	"github.com/realotz/whole/internal/apps/cms"
	"github.com/realotz/whole/internal/apps/systems"
	"github.com/realotz/whole/internal/apps/users"
)

// ProviderSet is apps providers.
var ProviderSet = wire.NewSet(
	cms.ProviderSet,
	systems.ProviderSet,
	users.ProviderSet,
)
