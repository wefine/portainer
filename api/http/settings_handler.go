package http

import (
	"github.com/portainer/portainer/api"

	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// SettingsHandler represents an HTTP API handler for managing settings.
type SettingsHandler struct {
	*mux.Router
	Logger   *log.Logger
	settings *portainer.Settings
}

// NewSettingsHandler returns a new instance of SettingsHandler.
func NewSettingsHandler(mw *middleWareService) *SettingsHandler {
	h := &SettingsHandler{
		Router: mux.NewRouter(),
		Logger: log.New(os.Stderr, "", log.LstdFlags),
	}
	h.Handle("/settings",
		mw.public(http.HandlerFunc(h.handleGetSettings)))

	return h
}

// handleGetSettings handles GET requests on /settings
func (handler *SettingsHandler) handleGetSettings(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		handleNotAllowed(w, []string{http.MethodGet})
		return
	}

	encodeJSON(w, handler.settings, handler.Logger)
}
