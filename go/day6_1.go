package main

import(
    "fmt"
    "os"
    "bufio"
)

func uniq (str string) bool {
    m := make(map[rune]int)
    for _, c := range str{
        m[c] ++ 
        if m[c] != 1{
            return false
        }
    }
    return true
}

func main(){
    file, _ := os.Open("../res/input_6.txt")
    fileReader := bufio.NewScanner(file)
    fileReader.Split(bufio.ScanLines)
    fileReader.Scan()
    str := fileReader.Text()
    for i, _:= range str[:len(str) - 15] {
        if uniq(str[i : i + 14]) == true{
            fmt.Println(i + 14)
            break
        }
    }

}
