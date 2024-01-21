package main

import "fmt"

func main() {
    myTutors := [4]string{"Kirsty", "Mishell", "Jose", "Neil"}
    fmt.Println(myTutors)
    myTutorsVersion := myTutors[:]
    subjects :=[]string{"Go", "Java", "Python"}
    fmt.Println(myTutorsVersion)
    fmt.Println(subjects)
}