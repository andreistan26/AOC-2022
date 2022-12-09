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
    'D' : Point{ x: -1 , y: 0, },
    'U' : Point{ x: 1 , y:  0, },
    'L' : Point{ x: 0 , y:  -1, },
    'R' : Point{ x: 0 , y:  1, },
}

func sign(n int) (s int){
    if n > 0{
        return 1
    }else{
        return -1
    }
}

func main(){
    file, _ := os.Open("../res/input_9.txt")
    fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanLines)

    var mat [1000][1000]int 
    mat[500][500] = 1
    counter := 1
    var arr [10]Point
    for i := 0; i < 10; i++{
        arr[i] = Point{500, 500}
    }
    for fileScanner.Scan() {
        var dir_str byte
        var dist    int
        fmt.Sscanf(fileScanner.Text(), "%c %d ", &dir_str, &dist)
        for i := 0; i < dist; i++{
            (&arr[0]).add(dir[dir_str])
            for j := 1; j < 10; j++{
                d_x := arr[j-1].x - arr[j].x
                d_y := arr[j-1].y - arr[j].y
                if !arr[j].is_nearby(arr[j-1]){
                    arr[j].add(Point{ x: sign(d_x), y : sign(d_y)})
                }
            }
            if mat[arr[9].x][arr[9].y] == 0{
                counter ++
                fmt.Println(arr[9])
                mat[arr[9].x][arr[9].y] = 1
            }
        }
    }
    
    fmt.Println(arr)
    fmt.Println(counter)
}
