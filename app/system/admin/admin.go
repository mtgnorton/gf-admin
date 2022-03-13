package cmd

import (
	"context"
	"gf-admin/app/shared"
	"gf-admin/app/system/admin/internal/controller"
	"gf-admin/app/system/admin/internal/service"
	"gf-admin/utility/response"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gsession"
	"github.com/gogf/gf/v2/protocol/goai"
	"github.com/gogf/gf/v2/util/gmode"
)

func Run(ctx context.Context) {
	var (
		s   = g.Server()
		oai = s.GetOpenApi()
	)

	// OpenApi自定义信息
	oai.Info.Title = `API Reference`
	oai.Config.CommonResponse = response.JsonRes{}
	oai.Config.CommonResponseDataField = `Data`

	// 静态目录设置
	uploadPath := g.Cfg().MustGet(ctx, "upload.path").String()
	if uploadPath == "" {
		g.Log().Fatal(ctx, "文件上传配置路径不能为空")
	}
	if !gfile.IsDir(uploadPath) {
		err := gfile.Mkdir(uploadPath)
		if err != nil {
			return
		}
	}
	s.AddStaticPath("/upload", uploadPath)

	// HOOK, 开发阶段禁止浏览器缓存,方便调试
	if gmode.IsDevelop() {
		s.BindHookHandler("/*", ghttp.HookBeforeServe, func(r *ghttp.Request) {
			r.Response.Header().Set("Cache-Control", "no-store")
		})
	}

	prefix, err := g.Cfg().Get(ctx, "server.prefix")
	if err != nil {
		g.Log().Fatalf(ctx, "get server admin prefix error,error info following : %s", err)
	}

	service.AdminTokenInstance.Init(ctx)
	// 前台系统路由注册
	s.Group(prefix.String(), func(group *ghttp.RouterGroup) {

		group.ALL("/ws", controller.Ws.Ws)

		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(
				shared.Middleware.Ctx,
				service.Middleware.ResponseHandler,
			)

			group.Bind(controller.NoAuth)

			group.Group("/", func(group *ghttp.RouterGroup) {
				//登录验证路由
				service.AdminTokenInstance.Middleware(group)

				group.Middleware(service.Middleware.Auth)

				group.Bind(controller.Personal)
				group.Bind(controller.Global)

				// 权限验证路由
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware.Permission)

					group.Bind(
						controller.Administrator,
						controller.Role,
						controller.Menu,
					)
				})
			})
		})

	})
	// 自定义丰富文档
	enhanceOpenAPIDoc(s)
	sessionConfig(s)
	service.Enforcer.Register(ctx)

	s.BindHookHandlerByMap("/*", map[string]ghttp.HandlerFunc{
		ghttp.HookBeforeServe: func(r *ghttp.Request) {
			glog.Println(ghttp.HookBeforeServe)
			r.SetParam("key1", "v11")
			r.GetRequest("key1")
		},
	})

	// 启动Http Server
	s.Run()
	return

}
func sessionConfig(s *ghttp.Server) {

	err := s.SetConfigWithMap(g.Map{
		"SessionStorage": gsession.NewStorageRedis(g.Redis("session")),
	})
	if err != nil {
		g.Log().Fatalf(gctx.New(), "init session driver error, %s", err)
	}
}
func enhanceOpenAPIDoc(s *ghttp.Server) {
	openapi := s.GetOpenApi()
	openapi.Config.CommonResponse = ghttp.DefaultHandlerResponse{}
	openapi.Config.CommonResponseDataField = `Data`

	openapi.Components = goai.Components{
		SecuritySchemes: goai.SecuritySchemes{
			"ApiKeyAuth": goai.SecuritySchemeRef{
				Ref: "", // 暂时还不知道该值是干什么用的
				Value: &goai.SecurityScheme{
					Type: "apiKey",
					In:   "header",
					Name: "Authorization",
				},
			},
		},
	}
	openapi.Security = &goai.SecurityRequirements{
		goai.SecurityRequirement{"ApiKeyAuth": []string{}},
	}
	// API description.
	openapi.Info.Title = `gf-admin`
	openapi.Info.Description = `后台接口文档`
}
