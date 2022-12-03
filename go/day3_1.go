package main

import(
    "fmt"
    "os"
    "bufio"
)

func contains(s []byte, char byte) bool {
    for _, a := range s {
        if a == char {
            return true
        }
    }
    return false
}

func main(){
    readFile, _ := os.Open("../res/input_3.txt")
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
    sum := 0
    for fileScanner.Scan(){
        rune_ar := []byte(fileScanner.Text())
        first := rune_ar[0 : len(rune_ar) / 2]
        sec := rune_ar[len(rune_ar)/2 : len(rune_ar) ]
        for _, a := range sec{
            if contains(first, a){
                if a <= 90{
                    sum += int(a - 'A' + 1) + 26
                    break
                }else{
                    sum += int(a - 'a' + 1)
                    break
                }
            }
        }
    }
    fmt.Println(sum)
}
