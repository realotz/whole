package users

import (
	"github.com/google/wire"
	"github.com/realotz/whole/internal/apps/users/biz"
	"github.com/realotz/whole/internal/apps/users/data"
	"github.com/realotz/whole/internal/apps/users/service"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	data.ProviderSet,
	biz.ProviderSet,
	service.NewMessageService,
	service.NewMemberService,
)
