package main

import (
  "fmt"
  "strconv"
)

// 1       2       3       4       5       6        7          8         9       10      11      12
// □ ■ □   ■ □ □   ■ □ □   ■ □ □   ■ □ □   ■ □ □
// ■ ■ ■   ■ ■ □   ■ ■ ■   ■ ■ ■   ■ ■ ■   ■ □ □   ■ ■ □ □   ■ □ □ □   □ ■ □ □   ■ ■ □   ■ □ ■
// □ ■ □ , □ ■ ■ , □ ■ □ , □ □ ■ , ■ □ □ , ■ ■ ■ , □ ■ ■ ■ , ■ ■ ■ ■ , ■ ■ ■ ■ , ■ ■ ■ , ■ ■ ■ , ■ ■ ■ ■ ■

func main() {
  for {
    var sInput string
    if _, err := fmt.Scan(&sInput); err != nil {
      break
    }
    input, _ := strconv.Atoi(sInput)
    ids = append(ids, input)
  }
  length = len(ids)

  var board [5][]int
  for i := 0; i < 5; i++ {
    for j := 0; j < length; j++ {
      board[i] = append(board[i], 0)
    }
  }

  fillAddress()

  calcResult(0, board)
  fmt.Printf("解の個数: %d個", len(results))
  fmt.Println("")
}

// calcResult ブロックを当てはめる
func calcResult(n int, board [5][]int) {
  id := ids[n]
  isFinished := false
  for blockIndex, block := range blocks[id] {
    blockAddress := blockAddresses[id][blockIndex]
    width := len(block[0])
    height := len(block)
    var searchWidth, searchHeight int
    if n == 0 {
      searchWidth = 1+(length-width)/2
      searchHeight = 1+(5-height)/2
    } else {
      searchWidth = length-width+1
      searchHeight = 6-height
    }

    for i := 0; i < searchHeight; i++ {
      for j := 0; j < searchWidth; j++ {
        // ブロックが入るかのチェック
        isInsert := true
        for _, a := range blockAddress{
          if board[i+a[0]][j+a[1]] != 0 {
            isInsert = false
            break
          }
        }

        // 入る場合、ブロックを入れる
        if isInsert {
          for _, a := range blockAddress{
            board[i+a[0]][j+a[1]] = id
          }

          // 埋め尽くした場合、結果に追加。
          if n == len(ids)-1 {
            var result [5][]int
            for boardLineIndex, boardLine := range board {
              var resultLine []int
              for _, b := range boardLine {
                resultLine = append(resultLine, b)
              }
              result[boardLineIndex] = resultLine
            }
            results = append(results, result)
            printResult(result)
            isFinished = true
          } else {
            calcResult(n+1, board)
          }

          // ブロックを外す
          for _, a := range blockAddress{
            board[i+a[0]][j+a[1]] = 0
          }

          // 埋め尽くすパターンが１つ見つかった場合、このブロックをどんなに回転させようと別のパターンはない。
          if isFinished {
            return
          }
        }
      }
    }
  }
}

// fillAddress ブロックを座標表示に変換
func fillAddress() {
  // ループが多いので、変数名をa,b,c,dに。
  for aIndex, a := range blocks {
    for _, b := range a {
      var blockAddress [][2]int
      for cIndex, c := range b {
        for dIndex, d := range c {
          if d == 1 {
            address := [2]int{cIndex, dIndex}
            blockAddress = append(blockAddress, address)
          }
        }
      }
      blockAddresses[aIndex] = append(blockAddresses[aIndex], blockAddress)
    }
  }
}

// printResult 盤面を綺麗に表示
func printResult(result [5][]int) {
  for _, resultLine := range result {
    for _, r := range resultLine {
      if r < 10 {
        fmt.Printf(" %d ", r)
      } else {
        fmt.Printf("%d ", r)
      }
    }
    fmt.Println("")
  }
  fmt.Println("-----------")
}

var count int = 0
var ids []int
var length int
var results [][5][]int
var blockAddresses [13][][][2]int
var blocks = [13][][][]int{
  {}, // id: 0 は使わない
  {
    {
      {0, 1, 0},
      {1, 1, 1},
      {0, 1, 0},
    },
  },
  {
    {
      {1, 0, 0},
      {1, 1, 0},
      {0, 1, 1},
    },
    {
      {0, 1, 1},
      {1, 1, 0},
      {1, 0, 0},
    },
    {
      {1, 1, 0},
      {0, 1, 1},
      {0, 0, 1},
    },
    {
      {0, 0, 1},
      {0, 1, 1},
      {1, 1, 0},
    },
  },
  {
    {
      {1, 0, 0},
      {1, 1, 1},
      {0, 1, 0},
    },
    {
      {0, 1, 1},
      {1, 1, 0},
      {0, 1, 0},
    },
    {
      {0, 1, 0},
      {1, 1, 1},
      {0, 0, 1},
    },
    {
      {0, 1, 0},
      {0, 1, 1},
      {1, 1, 0},
    },
    {
      {0, 0, 1},
      {1, 1, 1},
      {0, 1, 0},
    },
    {
      {0, 1, 0},
      {1, 1, 0},
      {0, 1, 1},
    },
    {
      {0, 1, 0},
      {1, 1, 1},
      {1, 0, 0},
    },
    {
      {1, 1, 0},
      {0, 1, 1},
      {0, 1, 0},
    },
  },
  {
    {
      {1, 0, 0},
      {1, 1, 1},
      {0, 0, 1},
    },
    {
      {0, 1, 1},
      {0, 1, 0},
      {1, 1, 0},
    },
    {
      {0, 0, 1},
      {1, 1, 1},
      {1, 0, 0},
    },
    {
      {1, 1, 0},
      {0, 1, 0},
      {0, 1, 1},
    },
  },
  {
    {
      {1, 0, 0},
      {1, 1, 1},
      {1, 0, 0},
    },
    {
      {1, 1, 1},
      {0, 1, 0},
      {0, 1, 0},
    },
    {
      {0, 0, 1},
      {1, 1, 1},
      {0, 0, 1},
    },
    {
      {0, 1, 0},
      {0, 1, 0},
      {1, 1, 1},
    },
  },
  {
    {
      {1, 0, 0},
      {1, 0, 0},
      {1, 1, 1},
    },
    {
      {1, 1, 1},
      {1, 0, 0},
      {1, 0, 0},
    },
    {
      {1, 1, 1},
      {0, 0, 1},
      {0, 0, 1},
    },
    {
      {0, 0, 1},
      {0, 0, 1},
      {1, 1, 1},
    },
  },
  {
    {
      {1, 1, 0, 0},
      {0, 1, 1, 1},
    },
    {
      {0, 1},
      {1, 1},
      {1, 0},
      {1, 0},
    },
    {
      {1, 1, 1, 0},
      {0, 0, 1, 1},
    },
    {
      {0, 1},
      {0, 1},
      {1, 1},
      {1, 0},
    },
    {
      {0, 0, 1, 1},
      {1, 1, 1, 0},
    },
    {
      {1, 0},
      {1, 0},
      {1, 1},
      {0, 1},
    },
    {
      {0, 1, 1, 1},
      {1, 1, 0, 0},
    },
    {
      {1, 0},
      {1, 1},
      {0, 1},
      {0, 1},
    },
  },
  {
    {
      {1, 0, 0, 0},
      {1, 1, 1, 1},
    },
    {
      {1, 1},
      {1, 0},
      {1, 0},
      {1, 0},
    },
    {
      {1, 1, 1, 1},
      {0, 0, 0, 1},
    },
    {
      {0, 1},
      {0, 1},
      {0, 1},
      {1, 1},
    },
    {
      {0, 0, 0, 1},
      {1, 1, 1, 1},
    },
    {
      {1, 0},
      {1, 0},
      {1, 0},
      {1, 1},
    },
    {
      {1, 1, 1, 1},
      {1, 0, 0, 0},
    },
    {
      {1, 1},
      {0, 1},
      {0, 1},
      {0, 1},
    },
  },
  {
    {
      {0, 1, 0, 0},
      {1, 1, 1, 1},
    },
    {
      {1, 0},
      {1, 1},
      {1, 0},
      {1, 0},
    },
    {
      {1, 1, 1, 1},
      {0, 0, 1, 0},
    },
    {
      {0, 1},
      {0, 1},
      {1, 1},
      {0, 1},
    },
    {
      {0, 0, 1, 0},
      {1, 1, 1, 1},
    },
    {
      {1, 0},
      {1, 0},
      {1, 1},
      {1, 0},
    },
    {
      {1, 1, 1, 1},
      {0, 1, 0, 0},
    },
    {
      {0, 1},
      {1, 1},
      {0, 1},
      {0, 1},
    },
  },
  {
    {
      {1, 1, 0},
      {1, 1, 1},
    },
    {
      {1, 1},
      {1, 1},
      {1, 0},
    },
    {
      {1, 1, 1},
      {0, 1, 1},
    },
    {
      {0, 1},
      {1, 1},
      {1, 1},
    },
    {
      {0, 1, 1},
      {1, 1, 1},
    },
    {
      {1, 0},
      {1, 1},
      {1, 1},
    },
    {
      {1, 1, 1},
      {1, 1, 0},
    },
    {
      {1, 1},
      {1, 1},
      {0, 1},
    },
  },
  {
    {
      {1, 0, 1},
      {1, 1, 1},
    },
    {
      {1, 1},
      {1, 0},
      {1, 1},
    },
    {
      {1, 1, 1},
      {1, 0, 1},
    },
    {
      {1, 1},
      {0, 1},
      {1, 1},
    },
  },
  {
    {
      {1, 1, 1, 1, 1},
    },
    {
      {1},
      {1},
      {1},
      {1},
      {1},
    },
  },
}
