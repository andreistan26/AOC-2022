package main

import(
    "fmt"
    "strconv"
    "os"
    "bufio"
)

func update_max(max []int, current int) {
    for i := 0; i < 3; i++{
        if current > max[i]{
            for j := 2; j > i ; j--{
                max[j - 1], max[j] = max[j], max[j - 1]
            }
            max[i] = current
            break;
        }
    }
}

func main(){
    current := 0
    max := make([]int, 3, 3)
    readFile, _ := os.Open("input.txt")
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
    for fileScanner.Scan(){
        converted, err := strconv.Atoi(fileScanner.Text())
        if err == nil{
            current += converted
        }else{
            update_max(max, current)
            current = 0
        }
    }
    update_max(max, current)
    sum := 0
    for i := 0; i < 3; i++{
        sum += max[i]
    }
    fmt.Println(sum)
}
