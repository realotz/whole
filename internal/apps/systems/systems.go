package systems

import (
	"github.com/google/wire"
	"github.com/realotz/whole/internal/apps/systems/biz"
	"github.com/realotz/whole/internal/apps/systems/data"
	"github.com/realotz/whole/internal/apps/systems/service"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	data.ProviderSet,
	biz.ProviderSet,
	service.NewFileServiceService,
)
