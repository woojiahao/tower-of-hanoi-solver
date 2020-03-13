package solver

import (
  "fmt"
  "math"
)

type (
  Disk struct {
    number int
    width  int
  }
  Rod struct {
    disks []Disk
  }
)

func Solve(disksCount int) {
  moves := 0

  puzzle := loadPuzzle(disksCount)
  fmt.Println(puzzle)

  ideal := int(math.Pow(2, float64(disksCount))) - 1
  if moves != ideal {
    panic(-1)
  }
}

func loadPuzzle(disksCount int) []Rod {
  disks := make([]Disk, disksCount, disksCount)
  for i := 0; i < disksCount; i++ {
    disks[i] = Disk{
      number: i,
      width:  i,
    }
  }
  rods := []Rod{
    {disks},
    {make([]Disk, 0, disksCount)},
    {make([]Disk, 0, disksCount)},
  }
  return rods
}
