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

type TaskService interface {
	GetAllTasks(ctx context.Context) ([]model.Task, error)
}

type TaskHandler struct {
	logger  *zap.SugaredLogger
	service TaskService
	config  *config.Config
}

func NewTaskHandler(logger *zap.SugaredLogger, service TaskService, cfg *config.Config) Handler {
	return &TaskHandler{
		logger:  logger,
		service: service,
		config:  cfg,
	}
}

func (th *TaskHandler) Register(router *mux.Router) {
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./templates/static/"))))
	router.HandleFunc("/", th.IndexHandler).Methods(http.MethodGet)
	router.HandleFunc("/todo", th.TaskHandler).Methods(http.MethodGet)
}

func (th *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	name := vars.Get("name")
	pass := vars.Get("pass")
	th.logger.Infof("task name: %s", name)
	th.logger.Infof("task password: %s", pass)
}

func (th *TaskHandler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	th.config.Session.Put(r.Context(), "message", "user exist")

	err := utils.RenderTemplate(w, r, "index.page.tmpl", &model.TemplateData{
		Title: "index page",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		th.logger.Fatalf("failed to execute html: %v", err)
	}
}

func (th *TaskHandler) TaskHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := th.service.GetAllTasks(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	data := model.TemplateData{
		Title: "Todo List",
		Tasks: tasks,
	}

	err = utils.RenderTemplate(w, r, "todo.page.tmpl", &data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		th.logger.Fatalf("failed to execute html: %v", err)
	}
}
