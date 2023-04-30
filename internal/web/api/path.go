package api

// GET shorts data
// func PathEndpoint(router *mux.Router) {
// 	router.HandleFunc("/{path}", func(rw http.ResponseWriter, r *http.Request) {
// 		vars := mux.Vars(r)
// 		shortPath := vars["path"]

// 		// Get all paths
// 		all, err := db.GetAllURLS()
// 		if err != nil {
// 			log.Println("Error getting keys")
// 		}

// 		// Check for paths
// 		for k, v := range all {
// 			if v == shortPath {
// 				longUrl, err := db.GetURL(shortPath)
// 				if err != nil {
// 					continue
// 				}
// 				timestamp, err := db.GetTime(shortPath)
// 				if err != nil {
// 					log.Println(err)
// 					timestamp = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
// 				}
// 				rw.WriteHeader(http.StatusOK)
// 				err = json.NewEncoder(rw).Encode(map[string]string{"path": shortPath, "url": longUrl, "time": timestamp.String()})
// 				if err != nil {
// 					return
// 				}
// 				break
// 			}
// 			if k >= len(all)-1 {
// 				rw.WriteHeader(http.StatusNotFound)
// 				_, err := rw.Write([]byte("this path does not exist"))
// 				if err != nil {
// 					return
// 				}
// 			}
// 		}
// 	}).Methods("GET")
// }
