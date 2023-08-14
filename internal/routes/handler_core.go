package routes

import (
	"fmt"
	"github.com/timeforaninja/shortpaste/internal/middleware"
	"github.com/timeforaninja/shortpaste/internal/routes/handler_templates"
	"github.com/timeforaninja/shortpaste/internal/types"
	"log"
	"net/http"
)

func HandleRequests(app types.AppInf, bind string) {
	// Short links
	http.HandleFunc("/f/", middleware.EnforceGet(handler_templates.ResolveShortFile(app)))
	http.HandleFunc("/l/", middleware.EnforceGet(handler_templates.ResolveShortLink(app)))
	http.HandleFunc("/t/", middleware.EnforceGet(handler_templates.ResolveShortText(app)))

	// Admin API
	http.HandleFunc("/api/v1/", middleware.EnforceAuth(app, HandleRestAPI(app)))

	// Static file server
	fs := http.FileServer(http.Dir("./static"))
	http.HandleFunc("/", middleware.EnforceAuth(app, fs.ServeHTTP))

	// Listen on Port
	fmt.Printf("Server starting at %s\n", bind)
	log.Fatal(http.ListenAndServe(bind, nil))
}
