package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

// Implementiamo il metodo Error per ErrNegativeSqrt
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// La funzione Sqrt restituisce la radice quadrata di x, o un errore se x è negativo
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

  // Implementazione del metodo di Newton per l'approssimazione della radice quadrata
	z := x
  // Iteriamo fino a convergenza o per un numero fisso di iterazioni
	for i := 0; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))   // Dovrebbe stampare un'approssimazione di sqrt(2)
	fmt.Println(Sqrt(-2))  // Dovrebbe stampare un errore
}
