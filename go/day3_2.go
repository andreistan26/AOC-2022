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
    stop := false
    for {
        var b [3][]byte
        for i := 0; i < 3; i++ {
            if !fileScanner.Scan(){
                stop = true
                break
            }
            b[i] = []byte(fileScanner.Text())
        } 
        if stop == true{
            break
        }
        for _, a := range b[0]{
            if contains(b[1], a) && contains(b[2], a){
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
