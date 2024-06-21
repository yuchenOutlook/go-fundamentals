package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type User struct {
	ID			int
	Username	string
}

type Server struct {
	db map[int]*User
	cache map[int]*User
	dbhit int
}

func NewServer() *Server {
	db := make(map[int]*User)

	for i := 0; i < 100; i++ {
		db[i + 1] = &User {
			ID: i + 1,
			Username: fmt.Sprintf("user%d", i + 1),
		} 
	}

	return &Server{
		db: db,
		cache: make(map[int]*User),
	}
}

func (s *Server) tryCache(id int) (*User, bool) {
	user, ok := s.cache[id]
	if ok {
		return user, true
	}
	return nil, false
}

func (s *Server) handleGetUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	// try to get user from cache
	user, ok := s.tryCache(id)
	if ok {
		json.NewEncoder(w).Encode(user)
		return
	}

	// if user found in db, increment dbhit
	user, ok = s.db[id]
	if !ok {
		panic("user not found")
	}
	s.dbhit++

	// insert user into cache
	s.cache[id] = user
	
	json.NewEncoder(w).Encode(user)
}

func main() {
	
}