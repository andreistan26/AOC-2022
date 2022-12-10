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
    cycle := 1
    ctr := 1
    var disp [240]byte
    var val []int
    for fileScanner.Scan(){
        var cmd string
        var arg int 
        fmt.Sscanf(fileScanner.Text(), "%s %d", &cmd, &arg)

        if strings.Compare(cmd, "noop") == 0{
            val = append(val, ctr)
            cycle ++
            continue
        }
        if strings.Compare(cmd, "addx") == 0{
            cycle += 2
            val = append(val, ctr, ctr)
            ctr += arg
        }
    }
    for i := 0; i < len(val); i++{
            if  val[i] - 1 <= i % 40 && i % 40 <= val[i] + 1{
                disp[i % 240] = '#'
            }else{
                disp[i % 240] = '.'
            }
            continue
        }
    for i := 39; i < 240; i += 40{
        fmt.Printf("%c\n", disp[i - 39 : i])
    }
}
