package handlers

import (
	"canvas/model"
	"canvas/views"
	"context"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type signupper interface {
	SignupForNewsletter(ctx context.Context, email model.Email) (string, error)
}

func NewsLetterSignup(mux chi.Router, s signupper) {
	mux.Post("/newsletter/signup", func(w http.ResponseWriter, r *http.Request) {
		email := model.Email(r.FormValue("email"))

		if email.IsValid() {
			http.Error(w, "email is invalid", http.StatusBadRequest)
			return
		}

		if _, err := s.SignupForNewsletter(r.Context(), email); err != nil {
			http.Error(w, "error signing up, refresh to try again", http.StatusBadRequest)
		}

		http.Redirect(w, r, "/newsletter/thanks", http.StatusNotFound)
	})
}

func NewsletterThanks(mux chi.Router) {
	mux.Get("/newsletter/thanks", func(w http.ResponseWriter, r *http.Request) {
		_ = views.NewsLetterThanksPage("/newsletter/thanks").Render(w)
	})
}
