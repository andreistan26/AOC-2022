package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)


func main(){
    file, _ := os.Open("../res/input_5.txt")
    fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanLines)
    var stacks [30][]byte
    for fileScanner.Scan(){
        if strings.Contains(fileScanner.Text(), "["){
            repl := strings.ReplaceAll(fileScanner.Text(), "   ", "[ ]")
            size := len(repl)
            ctr := 0
            for i := 1; i < size; i+=4{
                if (repl[i] != ' ') && (repl[i] != ']') && (repl[i] != '['){
                    stacks[ctr] = append([]byte{repl[i]}, stacks[ctr]...)
                }
                ctr++
            }
        }else{
            var move, from, to int
            val , _ := fmt.Sscanf(fileScanner.Text(), "move %d from %d to %d\n", &move, &from, &to)
            to --;
            from --;
            if val == 3{
                part := make([]byte, move)
                copy(part, stacks[from][ len(stacks[from]) - move : len(stacks[from])]);
                stacks[to] = append(stacks[to], part...)
                stacks[from] = stacks[from][:len(stacks[from]) - move]
            }
        }

    }
    for i := 0; cap(stacks[i]) != 0; i++{
        fmt.Printf("%c", stacks[i][len(stacks[i]) - 1])
    }
}
