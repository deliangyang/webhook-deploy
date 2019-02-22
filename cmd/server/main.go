package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/webHook/51b5ff52540137c1e334aa2b114ba64c", webHook)
	http.ListenAndServe("127.0.0.1:5000", nil)
}

func webHook(w http.ResponseWriter, r *http.Request) {
	body, err := r.GetBody()
	if err != nil {
		fmt.Println(err)
		return
	}
	content, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(content)
	go func() {
		res, err := exec.Command("bash", "../../bin/deploy.sh").Output()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(string(res))
	}()
	w.Write([]byte("success"))
}
