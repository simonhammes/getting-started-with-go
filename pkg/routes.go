package pkg

func (s *server) routes() {
	s.router.HandleFunc("/cache", s.handleGetCacheItem()).Methods("GET")
	s.router.HandleFunc("/cache", s.handleSetCacheItem()).Methods("POST")
}
