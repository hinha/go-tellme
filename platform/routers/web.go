package routers

import (
	"context"
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Router struct {
	Method        string
	Path          string
	Handler       gin.HandlerFunc
	Policy        gin.HandlerFunc
	MiddlewareOne gin.HandlerFunc
}

type routing struct {
	host         string
	domain       string
	service      string
	secretCookie string
}

type handlerInfo struct {
	method string
	path   string
	time   time.Time
}

// Routers contains the functions of http handler to clean payloads and pass it the service
type Routers interface {
	Serve()
}

var handlers []*Router

// Initialize is for initialize the handler
func Initialize(host string, routers []*Router, domain string, service string, secretCookie string) Routers {
	handlers = routers
	return &routing{
		host:         host,
		domain:       domain,
		service:      service,
		secretCookie: secretCookie,
	}
}

func (us *routing) Serve() {
	server := gin.Default()

	server.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	switch us.service {
	case "client":
		server.Static("/css", "views/assets/css")
		server.Static("/js", "views/assets/js")
		server.Static("/fonts", "views/assets/fonts")
		server.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
			Root:      "views/frontend",
			Extension: ".html",
			Partials:  []string{"partials/navbar"},
			Funcs: template.FuncMap{
				"copy": func() string {
					return time.Now().Format("2006")
				},
			},
			DisableCache: true,
		})

		for _, r := range handlers {
			server.Handle(r.Method, r.Path, r.MiddlewareOne, r.Handler)
		}

		break

	}

	logrus.WithFields(logrus.Fields{
		"host":   us.host,
		"domain": us.domain,
	}).Info("Starts Serving on HTTP")

	s := &http.Server{
		Addr:           us.host,
		Handler:        server,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		// service connections
		server.Run(us.host)
	}()

	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

}
