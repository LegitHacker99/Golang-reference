package app

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func App() error {
	http.HandleFunc("/", handlebaseurl)

	var wg sync.WaitGroup
	wg.Add(1)

	mainRouter := http.NewServeMux()
	userRouter := SetupUserRoutes()
	adminRouter := SetupAdminRoutes()

	mainRouter.Handle("/api/", http.StripPrefix("/api", userRouter))
	mainRouter.Handle("/admin/", http.StripPrefix("/admin", adminRouter))

	MiddlewareStack := CreateMiddlewareStack(
		Logger,
	)

	srv := http.Server{
		Addr:    ":8080",
		Handler: MiddlewareStack(mainRouter),
	}

	errChan := make(chan error)

	go func() {
		defer wg.Done()
		if err := srv.ListenAndServe(); err != nil {
			errChan <- err
		}
		close(errChan)
	}()

	log.Println("SERVER UP")

	wg.Wait()
	if err := <-errChan; err != nil {
		log.Println("SERVER STARTUP ERROR")
		return err
	}

	return nil
}

func handlebaseurl(w http.ResponseWriter, r *http.Request) {
	fmt.Println("working index endpoint")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(`{
		"msg": "success"
	}`))
}
