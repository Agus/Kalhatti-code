package main

import "fmt"
// import "strings"

func main(){
  grid := []string{"53..7....", "6..195...", ".98....6.", "8...6...3", "4..8.3..1", "7...2...6", ".6....28.", "...419..5", "....8..79"}
  fmt.Println(isValid(grid[0]))
}

func isValid(values string) bool{
  valuesMap := map[rune]bool{};
  for _, val := range values{
    if val == '.'{ continue;}
    if valuesMap[val] == true {
      return false;
      } else {
        valuesMap[val] = true
      }
    }

    return true;
  }
