package main

import "fmt"

func main() {
  // Create an array for our 4 player scores.
  var playerScores [4]int
  // Set the first player's score to 43
  playerScores[0] = 43
  // Set the second player's score to 7
  playerScores[1] = 7
  // Set the third player's score to 32
  playerScores[2] = 32
  // Set the fourth player's score to 65
  playerScores[3] = 65
  // Print the scores of all the players!
  fmt.Println(playerScores)
}