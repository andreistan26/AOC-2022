package main

import(
    "fmt"
    "os"
    "bufio"
)

func main(){
    readFile, _ := os.Open("../res/input_4.txt")
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
    sum := 0
    for fileScanner.Scan(){
        var g1 , g2[3]int
        fmt.Sscanf(fileScanner.Text(), "%d-%d,%d-%d", &g1[0], &g1[1], &g2[0], &g2[1])
        if ((g1[1] >= g2[0] && g1[0] <= g2[1]) || (g2[1] >= g1[0] && g2[0] <= g1[1])){
            sum ++;
        }
    }
    fmt.Println(sum)
}
