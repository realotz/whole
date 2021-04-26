package cms

import (
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	cmsV1 "github.com/realotz/whole/api/cms/v1"
	"github.com/realotz/whole/internal/server"
	"github.com/realotz/whole/internal/services/cms/biz"
	"github.com/realotz/whole/internal/services/cms/data"
	"github.com/realotz/whole/internal/services/cms/service"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	biz.ProviderSet,
	data.ProviderSet,
	service.NewCategoryService,
	NewCmsApp,
)

type Cms struct {
	*server.App
}

func NewCmsApp(hs *http.Server, gs *grpc.Server, m middleware.Middleware, category cmsV1.CategoryServiceServer) *Cms {
	// cms 文章cms服务组
	cmsV1.RegisterCategoryServiceServer(gs, category) //分类

	// cms 文章cms服务组
	hs.HandlePrefix("/cms/category", cmsV1.NewCategoryServiceHandler(category, http.Middleware(m))) // 分类
	return &Cms{server.NewApp(hs, gs)}
}
