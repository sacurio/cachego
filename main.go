package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const (
	ErrUserNotFound = "user not found"
)

type (
	User struct {
		ID       int
		Username string
	}
	Server struct {
		db    map[int]*User
		cache map[int]*User
		dbhit int
	}
)

func NewServer() *Server {
	db := make(map[int]*User)

	for i := 0; i < 100; i++ {
		db[i+1] = &User{
			ID:       i + 1,
			Username: fmt.Sprintf("user_%d", i+1),
		}
	}

	return &Server{
		db:    db,
		cache: make(map[int]*User),
	}
}

func (s *Server) tryCache(id int) (*User, bool) {
	user, ok := s.cache[id]
	return user, ok
}

func (s *Server) handleGetUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	// handle error after
	id, _ := strconv.Atoi(idStr)

	// First try to hit the cache
	user, ok := s.tryCache(id)
	if ok {
		json.NewEncoder(w).Encode(user)
		return
	}

	// hit the database
	user, ok = s.db[id]
	if !ok {
		panic(ErrUserNotFound)
	}
	s.dbhit++

	// insert in cache
	s.cache[id] = user
	json.NewEncoder(w).Encode(user)
}

func main() {

}
