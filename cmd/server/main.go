package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/webHook/51b5ff52540137c1e334aa2b114ba64c", webHook)
	http.ListenAndServe("127.0.0.1:5000", nil)
}

func webHook(w http.ResponseWriter, r *http.Request) {
	go func() {
		res, err := exec.Command("bash", "../../bin/deploy.sh").Output()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(string(res))
	}()
	w.Write([]byte("success"))
}
