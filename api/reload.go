package api

import (
	"context"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/go-chi/render"

	"github.com/stefanprodan/mgob/scheduler"
)

func reloadCtx(data *scheduler.Scheduler) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(context.WithValue(r.Context(), "scheduler", data))
			next.ServeHTTP(w, r)
		})
	}
}

func postReload(w http.ResponseWriter, r *http.Request) {
	sch := r.Context().Value("scheduler").(*scheduler.Scheduler)
	err := sch.Reload()
	if err != nil {
		render.Status(r, 500)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	logrus.Info("Plans reloaded")

	render.JSON(w, r, map[string]string{"status": "Plans reloaded"})
}
