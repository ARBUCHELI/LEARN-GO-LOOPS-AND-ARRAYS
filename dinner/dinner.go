package main

import (
  "fmt"
)

func main() {

  menu := []string{"Hamburgers", "Cheeseburgers", "Fries"}

  fmt.Println("The menu:")
  // WRITE LOOP GOING THROUGH MENU BELOW THIS LINE
  for index, value := range menu {
    fmt.Println("Number:", index, "Item:", value)
  }

  orders := map[string]string{
    "John": "Cheeseburgers",
    "Janet": "Hamburgers",
    "Jordan": "Fries",
  }
  // A late order
  orders["James"] = "Chicken Sandwich"
  
    fmt.Println("\nThe friend's orders:")
  // WRITE LOOP GOING THROUGH ORDERS BELOW THIS LINE
  for index, value := range orders {
    fmt.Println("Name:", index, "Order:", value)
  }
} 