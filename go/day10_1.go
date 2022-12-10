package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
    file, _ := os.Open("../res/input_10.txt")
    fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanLines)
    cycle := 0
    ctr := 1
    sum := 0
    save_state := []int{20, 60, 100, 140, 180, 220}
    for fileScanner.Scan(){
        if len(save_state) == 0{
            break
        }
        var cmd string
        var arg int 
        fmt.Sscanf(fileScanner.Text(), "%s %d", &cmd, &arg)
        if strings.Compare(cmd, "noop") == 0{
            cycle ++
            if save_state[0] <= cycle{
                sum += save_state[0] * ctr
                fmt.Printf("on cycle %d with stre %d\n", save_state[0], save_state[0] * ctr)
                save_state = append([]int{}, save_state[1:]...)
            }
            continue
        }
        if strings.Compare(cmd, "addx") == 0{
            cycle += 2
            if save_state[0] <= cycle{
                sum += save_state[0] * ctr
                fmt.Printf("on cycle %d with stre %d\n", save_state[0], save_state[0] * ctr)
                save_state = append([]int{}, save_state[1:]...)
            }
            ctr += arg
        }
    }
    fmt.Println(sum)
}
