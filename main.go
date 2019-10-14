package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"os/user"
	"strconv"
	"syscall"
	"time"

	"github.com/domgoodwin/go-api/app/bundles/devicebundle"
	"github.com/gorilla/mux"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Println("Running go-api as", user.Name, ":", user.Uid, "/", strconv.Itoa(os.Getpid()))
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	// Controllers declaration
	kc := &devicebundle.DeviceController{}
	go func() {
		sig := <-gracefulStop
		fmt.Printf("caught sig: %+v", sig)
		fmt.Println("Wait for 20 second to finish processing")
		time.Sleep(20 * time.Second)
		os.Exit(0)
	}()
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1/").Subrouter()
	// Routes handling
	s.HandleFunc("/prime", kc.NextPrime).Methods("GET")
	s.HandleFunc("/wait", kc.Wait).Methods("GET")
	s.HandleFunc("/device", kc.Index).Methods("GET")
	s.HandleFunc("/device", kc.Create).Methods("POST")
	s.HandleFunc("/db/tables", kc.ListTables).Methods("GET")
	s.HandleFunc("/r/records/{id}", kc.GetRecords).Methods("GET")
	s.HandleFunc("/r/records", kc.UpdateRecordSet).Methods("POST")
	s.HandleFunc("/bb/propen", kc.HandlePROpen).Methods("POST")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
