package middelware

// type Middleware func(HandlerFunc) HandlerFunc

// func MiddlewareChain(h HandlerFunc, middlewares ...Middleware) HandlerFunc {
// 	for _, m := range middlewares {
// 		h = m(h)
// 	}

// 	return h
// }

// func WithCors(h HandlerFunc) HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) error {

// 		allowedOrigin := os.Getenv("APP_BASE_URL")

// 		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

// 		if r.Method == "OPTIONS" {
// 			w.WriteHeader(http.StatusOK)
// 			return nil
// 		}

// 		return h(w, r)
// 	}
// }
