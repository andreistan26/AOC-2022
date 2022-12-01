package main

import(
    "fmt"
    "strconv"
    "os"
    "bufio"
)

func main(){
    current, max := 0, 0
    readFile, _ := os.Open("input.txt")
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
    for fileScanner.Scan(){
        converted, err := strconv.Atoi(fileScanner.Text())
        if err == nil{
            current += converted
        }else{
            if max < current{
                max = current
            }
            current = 0
        }
    }
    if max < current{
        max = current
    }
    fmt.Println(max)
}
