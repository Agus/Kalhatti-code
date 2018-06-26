package main

import "fmt"
import "math"
import "os"

func main(){
  rows := []string{"53..7....", "6..195...", ".98....6.", "8...6...3", "4..8.3..1", "7...2...6", ".6....28.", "...419..5", "....8..79"}
  sizeOfBox := len(rows[0]) // Get amount of cells in each sub-box, will be the same amount of rows/columns
  sizeOfSubBoxRow := int(math.Sqrt(float64(sizeOfBox))); // Since each sub-box is NxN size
  boxes := make([]string,sizeOfBox)
  columns := make([]string,sizeOfBox)
  for i, row := range rows{

    if !isValid(row) { // Check if row is valid
      os.Exit(1);
    }
    rowValues := []rune(row)
    for j, rowValue := range rowValues{
      boxes[((j/sizeOfSubBoxRow)+(sizeOfSubBoxRow*(i/sizeOfSubBoxRow)))] += string(rowValue);
      columns[j] +=  string(rowValue);
    }
  }
  for _, column := range columns{
    if !isValid(column){ os.Exit(1); }
  }
  for _, box := range boxes{
    if !isValid(box){ os.Exit(1); }
  }
  fmt.Println("Valid sudoku!!!");
}

func isValid(values string) bool{
  valuesMap := map[rune]bool{};
  for _, val := range values{
    if val == '.'{ continue;}
    if valuesMap[val] == true {
      fmt.Println("Invalid sudoku");
      return false;
      }else{
        valuesMap[val] = true;
      }
    }
    return true;
  }
