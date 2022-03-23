package handlers

import (
	"context"
	"net/http"

	"github.com/Elren44/elren_todo/internal/config"
	"github.com/Elren44/elren_todo/internal/model"
	"github.com/Elren44/elren_todo/pkg/utils"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type UserService interface {
	CreateUser(ctx context.Context, user model.User) (string, error)
}

type UserHandler struct {
	logger  *zap.SugaredLogger
	service UserService
	config  *config.Config
}

func NewUserHandler(logger *zap.SugaredLogger, service UserService, cfg *config.Config) Handler {
	return &UserHandler{
		logger:  logger,
		service: service,
		config:  cfg,
	}
}

func (ah *UserHandler) Register(router *mux.Router) {
	router.HandleFunc("/signup", ah.SignupHandler).Methods(http.MethodGet)
	router.HandleFunc("/signup", ah.PostSignupHandler).Methods(http.MethodPost)
	router.HandleFunc("/registered", ah.Registered).Methods(http.MethodGet)
	router.HandleFunc("/login", ah.LoginHandler).Methods(http.MethodGet)
	router.HandleFunc("/login", ah.PostLoginHandler).Methods(http.MethodPost)

}

func (ah *UserHandler) PostSignupHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	var data = model.TemplateData{
		Title:     "Registration",
		ExistUser: false,
	}

	if verifyFormEmpty(r) {
		user.Email = r.Form.Get("email")
		user.Password = r.Form.Get("password")
		uuid, err := ah.service.CreateUser(r.Context(), user)
		if err != nil {
			if err.Error() == "user exists" {
				ah.logger.Errorf("failed to create user: %v", err)
				data.ExistUser = true
				err := utils.RenderTemplate(w, r, "registration.page.tmpl", &data)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					ah.logger.Fatalf("failed to execute html: %v", err)
				}
			}
		} else {
			data.Title = "Registered"
			user.UUID = uuid
			ah.logger.Infof("user successfully created, uuid: %s", uuid)
			ah.logger.Info(user)

			http.Redirect(w, r, "/registered", http.StatusMovedPermanently)
			//
			//err = utils.RenderTemplate(w, r, "registered.page.tmpl", &data)
			//if err != nil {
			//	http.Error(w, err.Error(), http.StatusInternalServerError)
			//	ah.logger.Fatalf("failed to execute html: %v", err)
			//}
		}
	}
}

func (ah *UserHandler) SignupHandler(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	message := ah.config.Session.GetString(r.Context(), "message")
	stringMap["message"] = message

	var data = model.TemplateData{
		Title:     "Registration",
		ExistUser: false,
		StringMap: stringMap,
	}

	err := utils.RenderTemplate(w, r, "registration.page.tmpl", &data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		ah.logger.Fatalf("failed to execute html: %v", err)
	}

}

func verifyFormEmpty(r *http.Request) bool {
	if r.FormValue("email") == "" || r.FormValue("password") == "" || r.FormValue("password2") == "" {
		return false
	}
	return true
}
func verifyLoginFormEmpty(r *http.Request) bool {
	if r.FormValue("email") == "" || r.FormValue("password") == "" {
		return false
	}
	return true
}

func (ah *UserHandler) Registered(w http.ResponseWriter, r *http.Request) {

	data := model.TemplateData{
		Title: "Registered",
	}

	err := utils.RenderTemplate(w, r, "registered.page.tmpl", &data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		ah.logger.Fatalf("failed to execute html: %v", err)
	}
}

func (ah *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	data := model.TemplateData{
		Title:     "Login",
		ExistUser: true,
	}

	err := utils.RenderTemplate(w, r, "login.page.tmpl", &data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		ah.logger.Fatalf("failed to execute html: %v", err)
	}
}

func (ah *UserHandler) PostLoginHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	var data = model.TemplateData{
		Title:     "Login",
		ExistUser: false,
	}

	if verifyLoginFormEmpty(r) {
		user.Email = r.Form.Get("email")
		user.Password = r.Form.Get("password")
		ah.logger.Infof("search user with email - %s and password - %s", user.Email, user.Password)
		err := utils.RenderTemplate(w, r, "index.page.tmpl", &data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			ah.logger.Fatalf("failed to execute html: %v", err)
		}
	}
}
