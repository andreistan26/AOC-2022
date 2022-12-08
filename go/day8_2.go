package main

import(
    "fmt"
    "strconv"
    "os"
    "bufio"
    "strings"
)


func main(){
    file, _ := os.Open("../res/input_8.txt")
    fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanLines)
    var mat     [100][100]int
    i := 0
    j := 0
    for fileScanner.Scan(){
        arr := strings.Split(fileScanner.Text(), "")
        j = 0
        for _, el := range arr{
            mat[i][j], _ = strconv.Atoi(el)
            j++
        }
        i++
    }
    max := 0
    for idx_i := 0; idx_i < i; idx_i++{
        for idx_j := 0; idx_j < j; idx_j++{
            scenic := 1
            //up
            part := 0
            for k := idx_i - 1; k >= 0 ; k-- {
                part++
                if mat[k][idx_j] >= mat[idx_i][idx_j]{
                    break
                }
            }
            scenic *= part
            part = 0
            for k := idx_j - 1; k >= 0 ; k-- {
                part ++
                if mat[idx_i][k] >= mat[idx_i][idx_j]{
                    break
                }
            }
            scenic *= part
            part = 0
            for k := idx_i + 1; k < i ; k++ {
                part ++
                if mat[k][idx_j] >= mat[idx_i][idx_j]{
                    break
                }
            }
            scenic *= part
            part = 0
            for k := idx_j + 1; k < j ; k++ {
                part ++
                if mat[idx_i][k] >= mat[idx_i][idx_j]{
                    break
                }
            }
            scenic *= part
            if scenic > max{
                max = scenic
            }
        }
    }
    fmt.Println(max)
}
