package service

import (
	"net/http"
	"pql/pkg/utils/tool"
)

type FileService struct {
	*ServiceContext
}

func NewFileServer(sc *ServiceContext) *FileService {
	return &FileService{
		ServiceContext: sc,
	}
}

func (f *FileService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dir := tool.Flag(f.App.Env.Info().Debug, "bin/.PQL", ".PQL")
	http.FileServer(http.Dir(dir)).ServeHTTP(w, r)
}
