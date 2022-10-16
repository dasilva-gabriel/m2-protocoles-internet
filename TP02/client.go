package main

import (
	"fmt"
	//"httpclient"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func lastRequests(nb int, last string) string {

	resp, err := http.Get("https://jch.irif.fr:8082/chat/")

	if err != nil {
		fmt.Printf("Erreur dans la requête !")
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Erreur dans la requête !")
		os.Exit(1)
	}

	sb := string(body)

	fmt.Printf("======================!")

	splitted := strings.Split(sb, "\n")
	size := len(splitted)
	if size > nb && nb != -1 {
		size = nb
	}

	canPrint := false
	for i := len(splitted) - size; i < len(splitted)-1; i++ {

		if canPrint == false && (last == "" || splitted[i] == last) {
			fmt.Printf("YOUHOU")
			canPrint = true
		}

		request := fmt.Sprintf("https://jch.irif.fr:8082/chat/%s", splitted[i])
		res, err := http.Get(request)

		if err != nil {
			fmt.Printf("Erreur dans la requête ! 1")
			os.Exit(1)
		}

		body2, err := ioutil.ReadAll(res.Body)

		if err != nil {
			fmt.Printf("Erreur dans la requête ! 2")
			os.Exit(1)
		}

		sb := string(body2)
		if canPrint == true {
			fmt.Println(sb)
		}

	}

	if nb == -1 {
		fmt.Println(" ---------------------- FIN EXEC AVEC UN '", splitted[len(splitted)-1], "' EN ID")
		return splitted[len(splitted)-1]
	} else {
		fmt.Println(" ---------------------- FIN EXEC")
	}

	return ""

}

func main() {

	last := ""

	// Exercice 4
	//lastRequests(50, "")

	// Exercice 5
	for {
		last = lastRequests(-1, last)
		fmt.Println("")
		fmt.Println(last, "EST LE DERNIER MESSAGE")
		fmt.Println("")
	}

}
