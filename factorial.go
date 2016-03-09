package main
import "fmt"

func main(){
	var input int
	fmt.Printf("Give us a number for which you want to compute a factorial:")
	fmt.Scan(&input)

	var factorial int
	factorial=1
	for i := input; i >0; i-- {
		factorial=factorial*i

	}

	fmt.Printf("Factorial is:",factorial)

}
		
		