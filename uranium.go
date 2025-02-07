package uranium

import (
	"sync"

	"github.com/gofiber/fiber/v2"

	"github.com/go-uranium/uranium/cache"
	"github.com/go-uranium/uranium/storage"
	"github.com/go-uranium/uranium/utils/sendmail"
)

type Config struct {
	SiteName string
	SiteURL  string
}

type Uranium struct {
	storage storage.Provider
	cache   cache.Cacher
	sender  sendmail.Sender
	config  Config
	lock    *sync.RWMutex
}

type Error struct {
	Code int    `json:"-"`
	Err  bool   `json:"err"`
	Msg  string `json:"msg"`
}

func (e *Error) Error() string {
	return e.Msg
}

func NewError(status int, msg string) *Error {
	return &Error{
		Code: status,
		Err:  true,
		Msg:  msg,
	}
}

func New(s storage.Provider, c cache.Cacher, mail sendmail.Sender, conf Config) (*Uranium, error) {
	err := s.Init()
	if err != nil {
		return &Uranium{}, err
	}
	return &Uranium{
		storage: s,
		cache:   c,
		sender:  mail,
		config:  conf,
		lock:    &sync.RWMutex{},
	}, nil
}

func (uranium *Uranium) RouteForFiber(app *fiber.App) {
	app.Get("/user/:uid/info", uranium.HandleUserInfoByUID)
	app.Get("/user/:uid/basic", uranium.HandleUserBasicByUID)
	app.Get("/user/:uid/profile", uranium.HandleUserProfileByUID)
	app.Get("/user/:uid/sudo/auth", uranium.HandleUserAuthByUID)
	//app.Get("/user/:user/auth_methods",uranium.HandleUserInfoByUID)
	//app.Get("/user/:uid/sudo/totp",uranium.HandleUserInfoByUID)
	//app.Get("/user/:uid/sudo/webauthn",uranium.HandleUserInfoByUID)
	//app.Get("/user/username/:username/info",uranium.HandleUserInfoByUID)
	//app.Get("/user/username/:username/basic",uranium.HandleUserInfoByUID)
	//app.Get("/user/username/:username/profile",uranium.HandleUserInfoByUID)
}

//type Config struct {
//	SiteName string
//	Sender   sendmail.Sender
//	PageSize int64
//}
//
//// call uranium.Lock.Lock() before exiting
//type Ushio struct {
//	Data   storage.Provider
//	Cache  cache.Cacher
//	Config *Config
//	Lock   *sync.RWMutex
//}
//
//func New(provider storage.Provider, config *Config) (*Ushio, error) {
//	if config.PageSize < 1 {
//		config.PageSize = 35
//	}
//	cc := memory.New(provider, config.PageSize)
//	err := cc.Init()
//	if err != nil {
//		return &Ushio{}, err
//	}
//	return &Ushio{
//		Data:   provider,
//		Cache:  cc,
//		Config: config,
//		Lock:   &sync.RWMutex{},
//	}, nil
//}
//
//func (ushio *Ushio) Configure(app *fiber.App) {
//	app.Get("/", func(ctx *fiber.Ctx) error {
//		return ctx.Redirect("/home", http.StatusTemporaryRedirect)
//	})
//	app.Get("/home", ushio.HandleHome)
//	app.Get("/u/:name", ushio.HandleUser)
//	app.Get("/u/:name/posts", func(ctx *fiber.Ctx) error {
//		return ctx.Redirect("/u/"+ctx.Params("name"), 302)
//	})
//	app.Get("/u/:name/comments", ushio.HandleUserComments)
//	app.Get("/p/:post", ushio.HandlePost)
//	app.Get("/c/:tname", ushio.HandleCategory)
//	app.Get("/login", ushio.HandleLogin)
//	app.Post("/login", ushio.HandlePOSTLogin)
//	app.Get("/sign_up", ushio.HandleSignUp)
//	app.Post("/sign_up", ushio.HandlePOSTSignUp)
//	app.Get("/compose", ushio.HandleCompose)
//	app.Post("/compose", ushio.HandlePOSTCompose)
//	app.Post("/vote/post", ushio.HandlePOSTVotePost)
//	app.Get("/logout", func(ctx *fiber.Ctx) error {
//		ctx.ClearCookie("token")
//		return ctx.Redirect("/", http.StatusTemporaryRedirect)
//	})
//
//	app.Post("/md_parse", func(ctx *fiber.Ctx) error {
//		html, e := mdparse.Parse(string(ctx.Body()))
//		if e != nil {
//			return e
//		}
//		return ctx.SendString(string(*html))
//	})
//}
