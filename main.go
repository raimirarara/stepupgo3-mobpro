package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

func main() {

	const url = "github.com/gostaticanalysis/skeleton/v2"
	//const url = "github.com/tenntenn/greeting/tree/main/v2"

	list_result, err := exec.Command("go", "list", "-json", "-m", "-versions", url).Output()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(list_result))

	type List struct {
		Dir      string
		Versions []string
	}

	var list List
	json.Unmarshal(list_result, &list)
	fmt.Println(list.Versions)

	for i := 0; i < len(list.Versions); i++ {
		urlWithVersion := url + "@" + list.Versions[i]

		_, err := exec.Command("go", "get", urlWithVersion).Output()

		if err != nil {
			fmt.Println(err)
		}

		list_result_of_version, err := exec.Command("go", "list", "-json", "-m", urlWithVersion).Output()

		if err != nil {
			fmt.Println(err)
		}

		type Dir struct {
			Dir string
		}

		var path Dir
		json.Unmarshal(list_result_of_version, &path)
		fmt.Println(path.Dir)

		vet_result, err := exec.Command("go", "vet", "-json", path.Dir+"/...").CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(vet_result))
	}
}
