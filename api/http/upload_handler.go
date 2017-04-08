package http

import (
	"github.com/portainer/portainer/api"

	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

// UploadHandler represents an HTTP API handler for managing file uploads.
type UploadHandler struct {
	*mux.Router
	Logger      *log.Logger
	FileService portainer.FileService
}

// NewUploadHandler returns a new instance of UploadHandler.
func NewUploadHandler(mw *middleWareService) *UploadHandler {
	h := &UploadHandler{
		Router: mux.NewRouter(),
		Logger: log.New(os.Stderr, "", log.LstdFlags),
	}
	h.Handle("/upload/tls/{endpointID}/{certificate:(?:ca|cert|key)}",
		mw.authenticated(http.HandlerFunc(h.handlePostUploadTLS)))
	return h
}

func (handler *UploadHandler) handlePostUploadTLS(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		handleNotAllowed(w, []string{http.MethodPost})
		return
	}

	vars := mux.Vars(r)
	endpointID := vars["endpointID"]
	certificate := vars["certificate"]
	ID, err := strconv.Atoi(endpointID)
	if err != nil {
		Error(w, err, http.StatusInternalServerError, handler.Logger)
		return
	}

	file, _, err := r.FormFile("file")
	defer file.Close()
	if err != nil {
		Error(w, err, http.StatusInternalServerError, handler.Logger)
		return
	}

	var fileType portainer.TLSFileType
	switch certificate {
	case "ca":
		fileType = portainer.TLSFileCA
	case "cert":
		fileType = portainer.TLSFileCert
	case "key":
		fileType = portainer.TLSFileKey
	default:
		Error(w, portainer.ErrUndefinedTLSFileType, http.StatusInternalServerError, handler.Logger)
		return
	}

	err = handler.FileService.StoreTLSFile(portainer.EndpointID(ID), fileType, file)
	if err != nil {
		Error(w, err, http.StatusInternalServerError, handler.Logger)
	}
}
