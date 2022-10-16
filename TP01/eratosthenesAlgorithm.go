package main

import (
	"fmt"
)

/*
* Le crible d’Ératosthène
* Prenons un tableau contenant les entiers de 2 a n que l’on suppose tous premiers.
* Ensuite il suffit, pour chaque élément du tableau, de supprimer tous ses multiples.
 */

func eratosthenesAlgorithm(n int) {
	a := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		a[i] = true
	}

	for i := 2; i*i <= n; i++ {

		if a[i] == true {

			for j := i; i*j <= n; j++ {
				a[i*j] = false
			}
		}

	}

	for i := 2; i <= n; i++ {
		if a[i] == true {
			fmt.Println("Le nombre ", (i), "est premier.")
		}
	}

}

func main() {
	eratosthenesAlgorithm(1000)
}
