package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Pair struct{
    a, b int
}
var i, j int

func get_nb(coord Pair) (nbs []Pair){
    if coord.a != 0{
        nbs = append(nbs, Pair{coord.a - 1, coord.b})
    } 
    
    if coord.a != i - 1{
        nbs = append(nbs, Pair{coord.a + 1, coord.b})
    } 
    
    if coord.b != 0{
        nbs = append(nbs, Pair{coord.a, coord.b - 1})
    } 

    if coord.b != j - 1{
        nbs = append(nbs, Pair{coord.a, coord.b + 1})
    } 
    return nbs
}

func main(){
    file, _ := os.Open("../res/input_12.txt")
    fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanBytes)
    line := 0
    var end Pair
    var start Pair
    var mat [200][200]int
    for fileScanner.Scan(){
        if fileScanner.Bytes()[0] == 'E'{
            end = Pair{i, j}
        }
        if fileScanner.Bytes()[0] == 'S'{
            start = Pair{i, j}
        }
        if fileScanner.Bytes()[0] == '\n'{
            i++
            line = j
            j = 0
        }else{
            if fileScanner.Bytes()[0] == 'S'{
                mat[i][j] = 0
            }else if fileScanner.Bytes()[0] == 'E'{
                mat[i][j] = 26
            }else{
                mat[i][j] = int(fileScanner.Bytes()[0]) - int('a') + 1
            }
            j++
        }
    }
    j = line

    var queue []Pair
    var vis  [200][200]bool
    var dist  [200][200]int
    for k := 0; k < i; k++{
        for q := 0; q < j; q++{
            dist[k][q] = math.MaxInt64
            vis[k][q] = false
        }
    }
    fmt.Printf("%d %d \n", i, j)
    vis[start.a][start.b] = true
    dist[start.a][start.b] = 0
    queue = append(queue, Pair{start.a, start.b})
    for len(queue) > 0 {
        u := queue[0]
        queue = queue[1:]

        for _, n := range get_nb(u){
            is_grt := mat[n.a][n.b] > mat[u.a][u.b]
            if((mat[n.a][n.b] - mat[u.a][u.b] > 1 && is_grt)  || vis[n.a][n.b] == true){
                continue
            }
            vis[n.a][n.b] = true
            dist[n.a][n.b] = dist[u.a][u.b] + 1
            queue = append(queue, n)

            if(n.a == end.a && n.b == end.b){
                fmt.Println(dist[n.a][n.b])
            }

        }
    }
}

