package main

import(
    "fmt"
    "os"
    "bufio"
    "math"
)

type Point struct{
    x int
    y int
}

func (p1 *Point) add(p2 Point){
    p1.x += p2.x
    p1.y += p2.y
}

func (p1 Point) is_nearby(p2 Point) (b bool){
    if math.Abs(float64(p1.x - p2.x)) <= 1 && math.Abs(float64(p1.y - p2.y)) <= 1{
        return true
    }
    return false
}

func (p1 Point) is_diag(p2 Point) (b bool){
    if math.Abs(float64(p1.x - p2.x)) == 1 && math.Abs(float64(p1.y - p2.y)) == 1{
        return true
    }
    return false
}

func (p1 Point) negate(){
    p1.x = - p1.x
    p1.y = - p1.y
}

// D and U inverted
var dir = map[byte]Point{
    'U' : Point{ x: -1 , y: 0, },
    'D' : Point{ x: 1 , y:  0, },
    'L' : Point{ x: 0 , y:  -1, },
    'R' : Point{ x: 0 , y:  1, },
}

func main(){
    file, _ := os.Open("../res/input_9.txt")
    fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanLines)

    var mat [1000][1000]int 
    mat[500][500] = 1
    counter := 1
    head := Point{500, 500}
    tail := Point{500, 500}

    for fileScanner.Scan() {
        var dir_str byte
        var dist    int
        fmt.Sscanf(fileScanner.Text(), "%c %d ", &dir_str, &dist)
        for i := 0; i < dist; i++{
            prev := head
            (&head).add(dir[dir_str])
            if !tail.is_nearby(head){
                if tail.is_diag(prev){
                    tail = prev
                }else{
                    (&tail).add(dir[dir_str])
                }
            }
            if mat[tail.x][tail.y] == 0{
                counter ++
                mat[tail.x][tail.y] = 1
            }
        }
    }
    fmt.Println(counter)
}
