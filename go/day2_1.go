package main

import(
    "fmt"
    "os"
    "bufio"
    "strings"
)

func main(){
    readFile, _ := os.Open("../res/input_2.txt")
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
    sum := 0
    for fileScanner.Scan(){
        s := strings.Split(fileScanner.Text(), " ")
        opp, me := s[0][0], s[1][0]
        piece_value := me - byte('W')
        me -=  uint8('X' - 'A')
        if opp == me{
            sum += 3
        }else{
            if opp == 'A' && me == 'B' || opp == 'B' && me == 'C' || opp == 'C' && me == 'A'{
                sum += 6
            }
        }
        sum += int(piece_value)
    }
    fmt.Println(sum)
}

