package controllers

import (
	"net/http"
	"time"

	"github.com/caarlos0/watchub/config"
	"github.com/caarlos0/watchub/datastore"
	"github.com/caarlos0/watchub/shared/pages"
	"github.com/gorilla/sessions"
)

// Schedule ctrl
type Schedule struct {
	Base
	store datastore.Datastore
}

// NewSchedule ctrl
func NewSchedule(
	config config.Config,
	session sessions.Store,
	store datastore.Datastore,
) *Schedule {
	return &Schedule{
		Base: Base{
			config:  config,
			session: session,
		},
		store: store,
	}
}

// Handler handles /Schedule
func (ctrl *Schedule) Handler(w http.ResponseWriter, r *http.Request) {
	session, _ := ctrl.session.Get(r, ctrl.config.SessionName)
	id, _ := session.Values["user_id"].(int64)
	if session.IsNew || id == 0 {
		http.Error(w, "not logged in", http.StatusForbidden)
		return
	}
	if err := ctrl.store.Schedule(id, time.Now()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	pages.Render(w, "scheduled", ctrl.sessionData(w, r))
}
