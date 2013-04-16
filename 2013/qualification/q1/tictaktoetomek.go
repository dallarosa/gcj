package main

import (
  "fmt"
  "os"
  "io/ioutil"
  "strings"
  "strconv"
)

const (
  X_WON = "X won"
  O_WON = "O won"
  DRAW = "Draw"
  NOT_COMPLETED = "Game has not completed"
  X = rune('X')
  O = rune('O')
  SPACE = rune('.')
  T = rune('T')
)
type Board [][]rune

func hasSymbolAtRow(symbol rune, row []rune) (result bool, count int) {
  result = false
  for _,element := range row {
    if element == symbol {
      result = true
      count++
    }
  }
  return
}

func hasSpaceAt(row []rune) bool {
  result,_ := hasSymbolAtRow(SPACE, row)
  return result
}

func hasWinAt(row []rune) (result bool, symbol rune) {
    hasT,_ := hasSymbolAtRow(T, row)
    if hasX, xCount := hasSymbolAtRow(X, row); 
    ( hasX && xCount == 4 ) || ( hasX && xCount == 3 && hasT ){
      result = true
      symbol = X
    } else if hasO, oCount := hasSymbolAtRow(O, row);
    ( hasO && oCount == 4 ) || ( hasO && oCount == 3 && hasT ) {
      result = true
      symbol = O
    }
    return
}

func main () {
  filename := os.Args[1]
  bContents,_ := ioutil.ReadFile(filename)
  sContents := string(bContents)
  aContents := strings.Split(sContents,"\n")
  caseNum,_ := strconv.Atoi(aContents[0])

  for cn := 0; cn < caseNum; cn++ {
    var hasWin bool
    var hasSpace bool
    var element rune

    board := Board{
      []rune(aContents[4*cn + cn + 1]),
      []rune(aContents[4*cn + cn + 2]),
      []rune(aContents[4*cn + cn + 3]),
      []rune(aContents[4*cn + cn + 4])}

    for i:= 0; !hasWin && i < 4; i++ {
      hasWin, element = hasWinAt(board[i])
      if hasSpaceAt(board[i]) {
        hasSpace = true
      }
    }
    for i:= 0; !hasWin && i < 4; i++ {
      column := []rune{
        board[0][i],
        board[1][i],
        board[2][i],
        board[3][i]}
      hasWin, element = hasWinAt(column)
      if hasSpaceAt(column) {
        hasSpace = true
      }
    }

    if !hasWin {
      diagonal := []rune{
        board[0][0],
        board[1][1],
        board[2][2],
        board[3][3]}
      hasWin, element = hasWinAt(diagonal)
    }
    if !hasWin {
      diagonal := []rune{
        board[0][3],
        board[1][2],
        board[2][1],
        board[3][0]}
      hasWin, element = hasWinAt(diagonal)
    }

    if hasWin {
      if element == O {
        fmt.Printf("Case #%d: %s\n", cn + 1, O_WON)
      } else {
        fmt.Printf("Case #%d: %s\n", cn + 1, X_WON)
      }
    } else {
      if hasSpace {
        fmt.Printf("Case #%d: %s\n", cn + 1, NOT_COMPLETED)
      } else {
        fmt.Printf("Case #%d: %s\n", cn + 1, DRAW)
      }
    }

  }
}
