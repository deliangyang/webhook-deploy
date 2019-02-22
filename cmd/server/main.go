package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
)

func main() {
	http.HandleFunc("/webHook/51b5ff52540137c1e334aa2b114ba64c", webHook)
	http.ListenAndServe("127.0.0.1:5000", nil)
}

func webHook(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	if strings.Contains(string(content), "chore: build dev prod") {
		fmt.Println("chore: build dev prod")
		return
	}
	fmt.Println(string(content))
	go func() {
		res, err := exec.Command("bash", "../../bin/deploy.sh").Output()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(string(res))
	}()
	w.Write([]byte("success"))
}
