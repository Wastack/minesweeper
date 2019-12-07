package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
    "github.com/minesweeper/mine"
)

func display_map(sweeper mine.MineSweeper, w int, h int) {
    for i:=0;i<h;i++ {
        for j:=0;j<w;j++{
            field, _ := sweeper.Get(mine.Position{j,i})
            fmt.Printf("%s ", string(field.Display(sweeper)))
        }
        fmt.Print("\n")
    }
}

func _debug_print(sweeper mine.MineSweeper, w int, h int) {
    for i:=0;i<h;i++ {
        for j:=0;j<w;j++{
            field, _ := sweeper.Get(mine.Position{j,i})
            field_copy := *field
            field_copy.IsRevealed = true
            fmt.Printf("%s ", string(field_copy.Display(sweeper)))
        }
        fmt.Print("\n")
    }
}

func main() {
    var sweeper mine.MineSweeper
    const WIDTH, HEIGHT = 5, 5
    sweeper = mine.CreateMatrix(WIDTH, HEIGHT, 0.25)

    display_map(sweeper, WIDTH, HEIGHT)
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Print("Enter command: (e.g. 0 4)\n> ")
    for scanner.Scan() {
        splits := strings.Split(scanner.Text(), " ")
        if len(splits) < 2 {
            fmt.Println("Invalid user input")
            fmt.Print("Enter command: (e.g. 0 4)\n> ")
            continue
        }
        x, err := strconv.Atoi(splits[0])
        if err != nil  {
            fmt.Println("Invalid user input")
            fmt.Print("Enter command: (e.g. 0 4)\n> ")
            continue
        }
        y, err := strconv.Atoi(splits[1])
        if err != nil {
            fmt.Println("Invalid user input")
            fmt.Print("Enter command: (e.g. 0 4)\n> ")
            continue
        }
        f, ok := sweeper.Get(mine.Position{x,y})
        if !ok {
            fmt.Println("Invalid user input")
            fmt.Print("Enter command: (e.g. 0 4)\n> ")
            continue
        }
        f.IsRevealed = true
        display_map(sweeper, WIDTH, HEIGHT)
        if f.IsMine {
            fmt.Println("Game over. You suck..")
            break
        }
        if sweeper.AllNotMineRevealed() {
            fmt.Println("Victory!")
        }
        fmt.Print("Enter command: (e.g. 0 4)\n> ")
    }
}
