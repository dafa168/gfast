package boot

import (
	"fmt"
	"gfast/library/service"
	_ "gfast/swagger"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
)

var AdminGfToken *gtoken.GfToken

func init() {
	showLogo()
	g.Log().SetFlags(glog.F_ASYNC | glog.F_TIME_DATE | glog.F_TIME_TIME | glog.F_FILE_LONG)
	//g.Server().SetPort(8200)
	g.Server().AddStaticPath("/public", g.Cfg().GetString("server.ServerRoot"))
	//后台初始化配置
	initAdmin()
	//前台初始化
	initFront()
}

func initAdmin() {
	//无需验证权限的用户id
	service.NotCheckAuthAdminIds = g.Cfg().GetInts("adminInfo.notCheckAuthAdminIds")
	//后端分页长度配置
	service.AdminPageNum = g.Cfg().GetInt("adminInfo.pageNum")
	// 设置并启动后台gtoken处理
	initAdminGfToken()
}

func initFront() {
	gtoken := &gtoken.GfToken{
		LoginPath:       "/front/login",
		LoginBeforeFunc: service.FrontLogin,
		LogoutPath:      "/front/logout",
		AuthPaths:       g.SliceStr{"/front/*"},
		AuthAfterFunc:   service.AuthAfterFunc,
	}
	gtoken.Start()
}

func initAdminGfToken() {
	//多端登陆配置
	service.AdminMultiLogin = g.Cfg().GetBool("gToken.MultiLogin")
	AdminGfToken = &gtoken.GfToken{
		CacheMode:        g.Cfg().GetInt8("gToken.CacheMode"),
		CacheKey:         g.Cfg().GetString("gToken.CacheKey"),
		Timeout:          g.Cfg().GetInt("gToken.Timeout"),
		MaxRefresh:       g.Cfg().GetInt("gToken.MaxRefresh"),
		TokenDelimiter:   g.Cfg().GetString("gToken.TokenDelimiter"),
		EncryptKey:       g.Cfg().GetBytes("gToken.EncryptKey"),
		AuthFailMsg:      g.Cfg().GetString("gToken.AuthFailMsg"),
		MultiLogin:       service.AdminMultiLogin,
		LoginPath:        "/sysLogin/login",
		LoginBeforeFunc:  service.AdminLogin,
		LoginAfterFunc:   service.AdminLoginAfter,
		LogoutPath:       "/sysLogin/logout",
		AuthPaths:        g.SliceStr{"/system/*"},
		AuthAfterFunc:    service.AuthAfterFunc,
		LogoutBeforeFunc: service.AdminLoginOut,
	}
	AdminGfToken.Start()
}

func showLogo() {
	fmt.Println(" .----------------.  .----------------.  .----------------.  .----------------.  .----------------. \n| .--------------. || .--------------. || .--------------. || .--------------. || .--------------. |\n| |    ______    | || |  _________   | || |      __      | || |    _______   | || |  _________   | |\n| |  .' ___  |   | || | |_   ___  |  | || |     /  \\     | || |   /  ___  |  | || | |  _   _  |  | |\n| | / .'   \\_|   | || |   | |_  \\_|  | || |    / /\\ \\    | || |  |  (__ \\_|  | || | |_/ | | \\_|  | |\n| | | |    ____  | || |   |  _|      | || |   / ____ \\   | || |   '.___`-.   | || |     | |      | |\n| | \\ `.___]  _| | || |  _| |_       | || | _/ /    \\ \\_ | || |  |`\\____) |  | || |    _| |_     | |\n| |  `._____.'   | || | |_____|      | || ||____|  |____|| || |  |_______.'  | || |   |_____|    | |\n| |              | || |              | || |              | || |              | || |              | |\n| '--------------' || '--------------' || '--------------' || '--------------' || '--------------' |\n '----------------'  '----------------'  '----------------'  '----------------'  '----------------' ")
	fmt.Println("当前版本：", service.Version)
}
