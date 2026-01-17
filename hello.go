package main
import "fmt"
import "rsc.io/quote"
func main() {
	fmt.Println("Hello, World")
	fmt.Println(quote.Go())
	fmt.Println(Hello("Alice"))
}

func Hello(name string) string {
		message := fmt.Sprintf("\nHello, %s!", name)
		fmt.Printf("Hello, %s!", name)
		// message3 := fmt.Println("Hello, %s!", name)
		return message
}
