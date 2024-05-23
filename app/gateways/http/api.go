package http

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	studygroups "github.com/mellotonio/desafiogo/app/domain/studyGroups"
	httpStudyGroup "github.com/mellotonio/desafiogo/app/gateways/studyGroups"
)

// Presentation layer depends on Account, Transfer, Auth services
type API struct {
	StudyGroupService studygroups.Service
}

func NewApi(StudyGroupService studygroups.Service) *API {
	return &API{
		StudyGroupService: StudyGroupService,
	}
}

func (api API) Start(host string, port string) {
	router := chi.NewMux()

	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Handlers - Account & Transfer
	httpStudyGroup.NewHandler(router, api.StudyGroupService)

	applicationPort := fmt.Sprintf("%s:%s", host, port)

	server := &http.Server{
		Handler:      router,
		Addr:         applicationPort,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	fmt.Println("Starting api...")
	log.Fatal(server.ListenAndServe())
}
