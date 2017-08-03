package service

import (
	"gopkg.in/urfave/cli.v1"
	"fmt"
	"net/http"
	"encoding/json"
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
	http.ListenAndServe(addr, nil)
	return nil
}
