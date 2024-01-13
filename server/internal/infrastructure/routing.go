package infrastructure

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/takeuchi-shogo/ticket-api/config"
	"github.com/takeuchi-shogo/ticket-api/internal/adapters/controllers"
	"github.com/takeuchi-shogo/ticket-api/internal/infrastructure/middleware"
	"github.com/takeuchi-shogo/ticket-api/pkg/token"
)

const (
	ApiVersion = "/v1"

	HealthCheckRoot = "/healthChecks"

	// API Group
	ProductGroup = "/product"

	// API Routing
	UserRoot = "/users"
)

type Routing struct {
	// db   *DB
	Gin  *gin.Engine
	Port string
}

type Controllers struct {
	admin         controllers.AdministratorsController
	artist        controllers.ArtistsController
	auth          controllers.AuthController
	buy           controllers.BuyController
	creditCard    controllers.CreditCardsController
	event         controllers.EventController
	me            controllers.MeController
	organizer     controllers.OrganizersController
	ticket        controllers.TicketsController
	ticketItem    controllers.TicketItemsController
	ticketHasItem controllers.TicketHasItemsController
	user          controllers.UserController
}

func NewControllers(
	ad controllers.AdministratorsController,
	ar controllers.ArtistsController,
	a controllers.AuthController,
	b controllers.BuyController,
	c controllers.CreditCardsController,
	e controllers.EventController,
	me controllers.MeController,
	o controllers.OrganizersController,
	t controllers.TicketsController,
	ti controllers.TicketItemsController,
	thi controllers.TicketHasItemsController,
	u controllers.UserController,
) Controllers {
	return Controllers{
		admin:         ad,
		artist:        ar,
		auth:          a,
		creditCard:    c,
		event:         e,
		buy:           b,
		me:            me,
		organizer:     o,
		ticket:        t,
		ticketItem:    ti,
		ticketHasItem: thi,
		user:          u,
	}
}

func NewRouting(config config.ServerConfig, c Controllers) *Routing {

	if config.AppEnvironment == "development" {
		gin.SetMode(gin.DebugMode)
		fmt.Println("=========== development ===============")
	}

	r := &Routing{
		Gin:  gin.Default(),
		Port: fmt.Sprintf(":%d", config.Port),
	}

	r.cors(&config)

	r.Gin.Use(r.auth())

	// ハンドラーをGinに登録する
	r.Gin.GET("/artists", func(ctx *gin.Context) { c.artist.GetList(ctx) })
	r.Gin.POST("/artists", func(ctx *gin.Context) { c.artist.Post(ctx) })
	r.Gin.GET("/artists/:id", func(ctx *gin.Context) { c.artist.Get(ctx) })

	r.Gin.POST("/registerEmails", func(ctx *gin.Context) { c.auth.RegisterEmail(ctx) })
	r.Gin.POST("/verifyCode", func(ctx *gin.Context) { c.auth.VerifyCode(ctx) })
	r.Gin.POST("/signup", func(ctx *gin.Context) { c.auth.Signup(ctx) })
	r.Gin.POST("/signin", func(ctx *gin.Context) { c.auth.Signin(ctx) })
	r.Gin.POST("/logout", func(ctx *gin.Context) { c.auth.Logout(ctx) })

	r.Gin.POST("/buy", func(ctx *gin.Context) { c.buy.Post(ctx) })

	r.Gin.GET("/events", func(ctx *gin.Context) { c.event.GetList(ctx) })
	r.Gin.POST("/events", func(ctx *gin.Context) { c.event.Post(ctx) })
	r.Gin.GET("/events/:id", func(ctx *gin.Context) { c.event.Get(ctx) })
	r.Gin.GET("/events/artists/:artistID", func(ctx *gin.Context) { c.event.GetListByArtistID(ctx) })

	r.Gin.POST("/me", func(ctx *gin.Context) { c.me.Post(ctx) })

	r.Gin.POST("/organizers", func(ctx *gin.Context) { c.organizer.Post(ctx) })
	r.Gin.GET("/organizers/:id", func(ctx *gin.Context) { c.organizer.Get(ctx) })

	r.Gin.GET("/tickets", func(ctx *gin.Context) { c.ticket.GetList(ctx) })
	r.Gin.POST("/tickets", func(ctx *gin.Context) { c.ticket.Post(ctx) })
	r.Gin.GET("/tickets/:id", func(ctx *gin.Context) { c.ticket.Get(ctx) })

	r.Gin.POST("/ticketItems", func(ctx *gin.Context) { c.ticketItem.Post(ctx) })

	r.Gin.POST("/ticketHasItems", func(ctx *gin.Context) { c.ticketHasItem.Post(ctx) })

	r.Gin.GET("/users/:id", func(ctx *gin.Context) { c.user.Get(ctx) })

	v1Auth := r.Gin.Use(middleware.JwtAuthMiddleware(token.NewJwtMaker(config)))
	{
		v1Auth.GET("/credit_cards", func(ctx *gin.Context) { c.creditCard.Get(ctx) })
		v1Auth.POST("/credit_cards", func(ctx *gin.Context) { c.creditCard.Post(ctx) })

		v1Auth.GET("/me", func(ctx *gin.Context) { c.me.GetMe(ctx) })
		v1Auth.PATCH("/me", func(ctx *gin.Context) { c.me.Patch(ctx) })
	}

	v1Admin := r.Gin.Group("/admin")
	{
		v1Admin.GET("/administrators/:id", func(ctx *gin.Context) { c.admin.Get(ctx) })
	}

	return r
}

// CORS 対応
func (r *Routing) cors(config *config.ServerConfig) {
	corsConfig := cors.DefaultConfig()
	// リクエストの送信元の指定
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	// 資格情報（Cookie、認証ヘッダー、TLSクライアント証明書）の送信をOKするか
	// 実装予定

	// リクエスト間に使用できるHTTPヘッダーを指定
	corsConfig.AllowHeaders = []string{
		"Access-Control-Allow-Credentials",
		"Access-Control-Allow-Headers",
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"Authorization",
		"X-API-Key",
	}
	// ヘッダー名を羅列して、レスポンスの一部として開示するものを指定
	// 既定のセーフリストは7つだけだから
	corsConfig.ExposeHeaders = append(corsConfig.ExposeHeaders, "Authorization")
	r.Gin.Use(cors.New(corsConfig))
}

func (r *Routing) auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.GetHeader("X-API-KEY") != "tacketmaster-api-key" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "API Key is invalid",
			})
			return
		}
	}
}

func (r *Routing) Run() {
	s := &http.Server{
		Addr:           r.Port,
		Handler:        r.Gin,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

	fmt.Printf("Running server: %s\n", r.Port)
}
