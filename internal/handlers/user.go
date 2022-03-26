package handlers

import (
	"context"
	"net/http"

	"github.com/Elren44/elren_todo/internal/config"
	"github.com/Elren44/elren_todo/internal/forms"
	"github.com/Elren44/elren_todo/internal/model"
	"github.com/Elren44/elren_todo/pkg/utils"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type UserService interface {
	CreateUser(ctx context.Context, userDTO model.UserDTO) (string, error)
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

	err := r.ParseForm()
	if err != nil {
		ah.logger.Errorf("failed to parse form, %v", err)
	}

	userData := model.UserDTO{
		Email:     r.Form.Get("email"),
		Password:  r.Form.Get("password"),
		Password2: r.Form.Get("password2"),
	}

	form := forms.NewForm(r.PostForm)

	form.Required("email", "password", "password2")
	form.MinLenght("password", 8, r)
	form.MinLenght("password2", 8, r)
	form.IsEmail("email")
	form.EqualPasswords(r)

	var data = model.TemplateData{
		Title: "Registration",
		Form:  form,
	}

	if !form.Valid() {
		formData := make(map[string]interface{})
		formData["user_data"] = userData
		data.Data = formData

		err := utils.RenderTemplate(w, r, "registration.page.tmpl", &data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			ah.logger.Fatalf("failed to execute html: %v", err)
		}
		return
	}

	_, err = ah.service.CreateUser(r.Context(), userData)
	if err != nil {
		if err.Error() == "user exists" {
			ah.logger.Errorf("failed to create user: %v", err)
			form.ExistUser()
			err := utils.RenderTemplate(w, r, "registration.page.tmpl", &data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				ah.logger.Fatalf("failed to execute html: %v", err)
			}
		}
	} else {
		ah.logger.Info(userData)

		http.Redirect(w, r, "/registered", http.StatusMovedPermanently)
	}

}

func (ah *UserHandler) SignupHandler(w http.ResponseWriter, r *http.Request) {

	// stringMap := make(map[string]string)
	// message := ah.config.Session.GetString(r.Context(), "message")
	// stringMap["message"] = message

	var emptyUserData model.UserDTO
	data := make(map[string]interface{})
	data["user_data"] = emptyUserData

	var templateData = model.TemplateData{
		Title: "Registration",
		Form:  forms.NewForm(nil),
		Data:  data,
	}

	err := utils.RenderTemplate(w, r, "registration.page.tmpl", &templateData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		ah.logger.Fatalf("failed to execute html: %v", err)
	}

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
		Title: "Login",
		Form:  forms.NewForm(nil),
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
		Title: "Login",
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
