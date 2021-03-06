package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/router"
	"github.com/kataras/iris/middleware/logger"
	"github.com/pantazheng/bci/config"
	"github.com/pantazheng/bci/controller"
	"strconv"
)

func init() {
	controller.Menu()
}

func newApp() (api *iris.Application) {
	api = iris.New()
	api.Use(logger.New())

	api.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		if _, err := ctx.Writef("404 not found"); err != nil {
			fmt.Printf("%v", err)
		}
	})

	api.RegisterView(iris.HTML("./view", ".html").Delims("[[", "]]"))
	api.PartyFunc("/vue", func(vue router.Party) {
		vue.Get("/index", func(ctx iris.Context) {
			_ = ctx.View("index.html")
		})
		vue.StaticWeb("/", "./view")
	})
	api.PartyFunc("/wechat", func(weChat router.Party) {
		weChat.Any("/", controller.WeChat)
	})

	api.PartyFunc("/user", func(user router.Party) {
		user.Post("/", controller.UserCreate)
		user.Get("/id/{id:uint}", controller.UserFindByID)
		user.Get("/telephone/{telephone:string}", controller.UserFindByTelephone)
		user.Get("/openid/{openid:string}", controller.UserFindByOpenID)
		user.Get("/level/{level:int}", controller.UsersFindByLevel)
		user.Get("/list", controller.UsersList)
		user.Put("/update", controller.UserUpdates)
		user.Put("/bind", controller.UserBind)
		user.Delete("/id/{id:uint}", controller.UserDeleteByID)
		user.Delete("/telephone/{telephone:string}", controller.UserDeleteByTelephone)
	})

	api.PartyFunc("/gain", func(gain router.Party) {
		gain.Post("/", controller.GainInsert)
		gain.Get("/id/{id:uint}", controller.GainFindByID)
		gain.Get("/mission/{id:uint}", controller.GainsFindByMissionID)
		gain.Get("/owner/{id:uint}", controller.GainsFindByOwnerID)
		gain.Get("/module/{id:uint}", controller.GainsFindByModuleID)
		gain.Get("/leader/{id:uint}", controller.GainsFindByLeaderID)
		gain.Get("/project/{id:uint}", controller.GainsFindByProjectID)
		gain.Get("/manager/{id:uint}", controller.GainsFindByManagerID)
		gain.Get("/all",
			controller.GainsFindAll)
		gain.Get("/file/{id:uint}", controller.GainDownFileByID)
		gain.Put("/", controller.GainUpdates)
		gain.Delete("/id/{id:uint}", controller.GainDeleteByID)
	})

	api.PartyFunc("/mission", func(mission router.Party) {
		mission.Post("/", controller.MissionInsert)
		mission.Get("/id/{id:uint}", controller.MissionFindByID)
		mission.Get("/owner/{id:uint}", controller.MissionsFindByOwnerID)
		mission.Get("/module/{id:uint}", controller.MissionsFindByModuleID)
		mission.Get("/leader/{id:uint}", controller.MissionsFindByLeaderID)
		mission.Get("/project/{id:uint}", controller.MissionsFindByProjectID)
		mission.Get("/manager/{id:uint}", controller.MissionsFindByManagerID)
		mission.Get("/all", controller.MissionsFindAll)
		mission.Put("/", controller.MissionUpdate)
		mission.Delete("/id/{id:uint}", controller.MissionDeleteByID)
	})

	api.PartyFunc("/module", func(module router.Party) {
		module.Post("/", controller.ModuleInsert)
		module.Get("/id/{id:uint}", controller.ModuleFindByID)
		module.Get("/leader/{id:uint}", controller.ModulesFindByLeaderID)
		module.Get("/project/{id:uint}", controller.ModulesFindByProjectID)
		module.Get("/manager/{id:uint}", controller.ModulesFindByManagerID)
		module.Get("/member/{id:uint}", controller.ModulesFindByMemberID)
		module.Get("/all", controller.ModulesFindAll)
		module.Put("/", controller.ModuleUpdate)
		module.Delete("/id/{id:uint}", controller.ModuleDeleteByID)
	})

	api.PartyFunc("/project", func(project router.Party) {
		project.Post("/", controller.ProjectInsert)
		project.Get("/id/{id:uint}", controller.ProjectFindByID)
		project.Get("/frame/{id:uint}", controller.ProjectFindFrame)
		project.Get("/manager/{id:uint}", controller.ProjectsFindByManagerID)
		project.Get("/member/{id:uint}", controller.ProjectsFindByMemberID)
		project.Get("/all", controller.ProjectsFindAll)
		project.Put("/", controller.ProjectUpdate)
		project.Delete("/id/{id:uint}", controller.ProjectDeleteByID)
	})
	return
}

func main() {
	app := newApp()
	_ = app.Run(iris.AutoTLS(":"+strconv.Itoa(config.Conf.APP.Port), "bci.renjiwulian.com", "pantazheng@vip.qq.com"))
}
