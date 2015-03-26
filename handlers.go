package main

import "net/http"

func (s *Server) handleJoin(w http.ResponseWriter, r *http.Request) {
	// Join has an optional parameters:
	// - registered player ID
	// - specific game ID
	// playerID := r.FormValue("player")
	// gameID := r.FormValue("game")

	// First, generate or retrive player information (nickname)
}

func (s *Server) handleLeave(w http.ResponseWriter, r *http.Request) {
}

func (s *Server) handleResume(w http.ResponseWriter, r *http.Request) {
}

func (s *Server) handleWait(w http.ResponseWriter, r *http.Request) {
}

func (s *Server) handleBid(w http.ResponseWriter, r *http.Request) {
}

func (s *Server) handleCoinche(w http.ResponseWriter, r *http.Request) {
}

func (s *Server) handleMessage(w http.ResponseWriter, r *http.Request) {
}

func (s *Server) handlePlay(w http.ResponseWriter, r *http.Request) {
}

func (s *Server) prepareHandlers() {
	http.HandleFunc("join", s.handleJoin)
	http.HandleFunc("leave", s.handleLeave)
	http.HandleFunc("resume", s.handleResume)
	http.HandleFunc("wait", s.handleWait)
	http.HandleFunc("bid", s.handleBid)
	http.HandleFunc("coinche", s.handleCoinche)
	http.HandleFunc("play", s.handlePlay)
	http.HandleFunc("message", s.handleMessage)
}
