package main

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"os/exec"
)

func goVet(url string) string {
	// url = "github.com/gostaticanalysis/skeleton/v2"
	// const url = "github.com/tenntenn/greeting/tree/main/v2"

	list_result, err := exec.Command("go", "list", "-json", "-m", "-versions", url).CombinedOutput()

	if err != nil {
		fmt.Println("list_err", err)
		return err.Error()
	}
	fmt.Println(string(list_result))

	type List struct {
		Dir      string
		Versions []string
	}

	var list List
	json.Unmarshal(list_result, &list)
	fmt.Println(list.Versions)

	var responseData string

	for i := 0; i < len(list.Versions); i++ {
		urlWithVersion := url + "@" + list.Versions[i]

		_, err := exec.Command("go", "get", urlWithVersion).CombinedOutput()

		if err != nil {
			fmt.Println("get err:", err)
			return err.Error()
		}

		list_result_of_version, err := exec.Command("go", "list", "-json", "-m", urlWithVersion).CombinedOutput()

		if err != nil {
			fmt.Println("list_result_of_version err: ", err)
			return err.Error()
		}

		type Dir struct {
			Dir string
		}

		var path Dir
		json.Unmarshal(list_result_of_version, &path)
		fmt.Println("path_dir: ", path.Dir)

		cmd := exec.Command("go", "vet", "-json", "./...")

		cmd.Dir = path.Dir
		vet_result, err := cmd.CombinedOutput()

		if err != nil {
			fmt.Println("vet_result err: ", err)
			return err.Error()
		}
		// fmt.Println(string(vet_result))
		responseData += urlWithVersion + "\n"
		responseData += string(vet_result) + "\n"
	}
	println(responseData)
	return responseData
}

func main() {

	urlMap := make(map[string]string, 2)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "Hello, %s\n", html.EscapeString(r.URL.Path))

		if urlMap[r.URL.Path[1:]] == "" {
			// fmt.Fprintln(w, "1回目")
			var response = goVet(r.URL.Path[1:])
			println("fuga",response)
			urlMap[r.URL.Path[1:]] = response
			fmt.Fprintln(w, "hoge")
		} else {
			fmt.Fprintln(w, urlMap[r.URL.Path[1:]])
			// fmt.Fprintln(w, "すでに実行済みです")
		}

	})

	http.ListenAndServe(":8000", nil)

}
