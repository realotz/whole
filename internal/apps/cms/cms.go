package cms

import (
	"github.com/google/wire"
	"github.com/realotz/whole/internal/apps/cms/biz"
	"github.com/realotz/whole/internal/apps/cms/data"
	"github.com/realotz/whole/internal/apps/cms/service"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	biz.ProviderSet,
	data.ProviderSet,
	service.NewCategoryServiceService,
)
