package main

import(
    "fmt"
    "os"
    "bufio"
    "strings"
)

var to_win = map[byte]int{
    'A':2,
    'B':3,
    'C':1,
}

var decision = map[byte]int{
    'Z':6,
    'Y':3,
    'X':0,
}

var values = map[byte]int{
    'A':1,
    'B':2,
    'C':3,
}

var to_lose = map[byte]int{
    'A':3,
    'B':1,
    'C':2,
}

func main(){
    readFile, _ := os.Open("../res/input_2.txt")
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
    sum := 0
    for fileScanner.Scan(){
        s := strings.Split(fileScanner.Text(), " ")
        opp, res := s[0][0], s[1][0]
        if res == 'Z'{
            sum += to_win[opp]
        }else if res == 'X'{
            sum += to_lose[opp]
        }else{
            sum += values[opp]
        }
        sum += decision[res]
    }
    fmt.Println(sum)
}

