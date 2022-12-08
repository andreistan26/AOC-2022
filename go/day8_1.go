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
    count := 0
    for idx_i := 0; idx_i < i; idx_i++{
        for idx_j := 0; idx_j < j; idx_j++{
            is_vis := true
            //up
            for k := idx_i - 1; k >= 0 ; k-- {
                if mat[k][idx_j] >= mat[idx_i][idx_j]{
                    is_vis = false
                    break
                }
            }
            if is_vis == true{
                count ++ 
                continue
            }
            is_vis = true
            //left
            for k := idx_j - 1; k >= 0 ; k-- {
                if mat[idx_i][k] >= mat[idx_i][idx_j]{
                    is_vis = false
                    break
                }
            }
            if is_vis == true{
                count ++ 
                continue
            }
            is_vis = true
            //down
            for k := idx_i + 1; k < i ; k++ {
                if mat[k][idx_j] >= mat[idx_i][idx_j]{
                    is_vis = false
                    break
                }
            }
            if is_vis == true{
                count ++ 
                continue
            }
            is_vis = true

            //right
            for k := idx_j + 1; k < j ; k++ {
                if mat[idx_i][k] >= mat[idx_i][idx_j]{
                    is_vis = false
                    break
                }
            }
            if is_vis == true{
                count ++ 
                continue
            }
        }
    }
    fmt.Println(count)
}
