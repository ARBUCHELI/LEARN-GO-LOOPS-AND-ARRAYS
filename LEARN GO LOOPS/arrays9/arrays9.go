package main
import "fmt"

func main() {
    myTutors := []string{"Kirsty", "Mishell", "Jose", "Neil"}
    changeLastElement(myTutors, "Bobby")
}

func changeLastElement(slice []string, newName string){
  if (len(slice) > 0) {
    slice[len(slice) - 1] = newName
  }
  fmt.Println(slice)
}