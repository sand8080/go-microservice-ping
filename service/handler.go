package service

import (
	"encoding/json"
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	msg := map[string]string{"message": "pong"}
	resp, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func Serve(context *cli.Context) error {
	http.HandleFunc("/", handler)
	addr := fmt.Sprintf("%s:%d", context.String("address"), context.Int("port"))
	fmt.Printf("Starting serve at %s\n", addr)
	http.ListenAndServe(addr, nil)
	return nil
}
