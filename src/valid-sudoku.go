package main

import "fmt"
import "math"

func main(){
  rows := []string{"53..7....", "6..195...", ".98....6.", "8...6...3", "4..8.3..1", "7...2...6", ".6....28.", "...419..5", "....8..79"}
  result := validSudoku(rows);
  fmt.Println(result);
}

func validSudoku(rows []string) int{
  // Get amount of cells in each sub-box, will be the same amount of rows or columns
  sizeOfBox := len(rows[0])
  // Since each sub-box is NxN size
  sizeOfSubBoxRow := int(math.Sqrt(float64(sizeOfBox)));
  boxes := make([]string,sizeOfBox)
  columns := make([]string,sizeOfBox)
  for i, row := range rows{
    // Check if row is valid
    if !isValid(row){ return 0;}
    rowValues := []rune(row)
    for j, rowValue := range rowValues{
      boxes[((j/sizeOfSubBoxRow)+(sizeOfSubBoxRow*(i/sizeOfSubBoxRow)))] += string(rowValue);
      columns[j] +=  string(rowValue);
    }
  }
  for _, column := range columns{
    if !isValid(column){ return 0; }
  }
  for _, box := range boxes{
    if !isValid(box){ return 0; }
  }
  return 1;
}

func isValid(values string) bool{
  valuesMap := map[rune]bool{};
  for _, val := range values{
    if val == '.'{ continue;}
    if valuesMap[val] == true {
      return false;
      }else{
        valuesMap[val] = true;
      }
    }
    return true;
  }
