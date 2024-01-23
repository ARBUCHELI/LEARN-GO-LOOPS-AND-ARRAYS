package main

import (
  "fmt" 
  "math/rand" 
  "time"
)

func getRandomElement(slice []string) string {
    length := len(slice)
    randomValue := rand.Intn(length)
    return slice[randomValue]
  }

func main() {
  rand.Seed(time.Now().UnixNano())
  guests := [8]string{"Tom", "Holly", "Ioki", "Kurt", "Christ", "Dave", "Billy", "Jimmy"}

  objects := [8]string{"A Toy Chest", "A Crate", "A Box", "A Drawer", "A TV", "A Shelf", "A Washing Machine", "A Desk"}

  fmt.Println("Guests: ", guests)
  fmt.Println("Objects: ", objects)

  
  guestsSliceVersion := guests[:]
  objectsSliceVersion := objects[:]

  theGuest := getRandomElement(guestsSliceVersion)
  theLocation := getRandomElement(objectsSliceVersion)

  fmt.Println(theGuest, "hid the cat by putting it in ", theLocation)
}