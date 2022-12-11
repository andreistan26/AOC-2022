package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
    "sort"
)

type Operation int

const (
     Add Operation = iota
     Dif
     Mult   
     Div
)


type Monkey struct{
    op      Operation
    arg     int
    items   []int
    test    int
    if_cond [2]int
    count int
    is_old_arg bool
}

var op_map = map[string]Operation{
    "*" : Mult,
    "+" : Add,
    "-" : Dif, 
    "/" : Div,
}

func (monkey *Monkey) do_op(){
    if monkey.is_old_arg == true{
        monkey.arg = monkey.items[0]
    }
    switch monkey.op {
    case Mult:
        monkey.items[0] *= monkey.arg
    case Div:
        monkey.items[0] /= monkey.arg
    case Dif:
        monkey.items[0] -= monkey.arg
    case Add:
        monkey.items[0] += monkey.arg
    }
}

func main(){
    file, _ := os.Open("../res/input_11.txt")
    fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanLines)
    var monkeys []Monkey
    lcm := 1
    for fileScanner.Scan(){
        var monkey Monkey
        fileScanner.Scan()
        splited := strings.Split(fileScanner.Text(), ":")
        nums := strings.Split(splited[1], ",")
        for _, num := range nums{
            i, _ := strconv.Atoi(num[1:])
            monkey.items = append(monkey.items, i)
        }
        fileScanner.Scan()
        op_split := strings.Split(fileScanner.Text(), " ")
        if strings.Compare(op_split[7], "old") == 0{
            monkey.is_old_arg = true
        }else{
            monkey.arg, _ = strconv.Atoi(op_split[7])
        }
        monkey.op = op_map[op_split[6]]
        fileScanner.Scan()
        split_div := strings.Split(fileScanner.Text(), " ")
        monkey.test, _ = strconv.Atoi(split_div[5])
        for i := 0; i < 2; i++{
            fileScanner.Scan()
            test_split := strings.Split(fileScanner.Text(), " ")
            monkey.if_cond[1-i], _ = strconv.Atoi(test_split[9])
        }
        lcm *= monkey.test
        monkey.count = 0
        monkeys = append(monkeys, monkey)
        fileScanner.Scan()
    }

    for _, monkey := range monkeys{
        fmt.Println(monkey)
    }

    rounds := 10000
    for j := 0; j < rounds; j++{
        for idx := 0; idx < len(monkeys); idx++{
            for len(monkeys[idx].items) != 0{
                (&monkeys[idx]).do_op()
                is_t := 0
                if monkeys[idx].items[0] % monkeys[idx].test == 0{
                    is_t = 1
                }
                monkeys[monkeys[idx].if_cond[is_t]].items = append(monkeys[monkeys[idx].if_cond[is_t]].items, monkeys[idx].items[0] % lcm)
                monkeys[idx].items = append([]int{}, monkeys[idx].items[1:]...)
                monkeys[idx].count ++ 
            }
        }
    }
    for _, m := range monkeys{
        fmt.Println(m.count)
    }
    
    sort.Slice(monkeys, func(i, j int) bool {
      return monkeys[i].count > monkeys[j].count
    })
    fmt.Println(monkeys[1].count * monkeys[0].count)





}
