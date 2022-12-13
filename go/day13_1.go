package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type pa

func get_next(stream string) (res string, new_stream string, is_arr bool){
    start := 0
    end := 0
    if stream[0] == '['{
        par := 1
        start = 1
        for par != 0{
            end++
            if(stream[end] == ']'){
                par--
            }
            if(stream[end] == '['){
                par++
            }
        }
        is_arr = true
        end--
    }else{
        end = strings.Index(stream, "[")
        if end == -1{
            end = len(stream) - 1
        }else{
           end --
        }
        is_arr = false
    }

    new_stream_idx := strings.Index(stream[end+1:], ",")
    if new_stream_idx == -1 {
        new_stream_idx = end + 1
    }
    return stream[start:end+1], stream[new_stream_idx + 1:], is_arr
}

func main(){
    file, _ := os.Open("../res/tmp.txt")
    fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanLines)
    for fileScanner.Scan(){
        line1 := fileScanner.Text()[1:len(fileScanner.Text())-1]
        fileScanner.Scan()
        line2 := fileScanner.Text()[1:len(fileScanner.Text())-1]
        for true{
            res1, out1, is1 := get_next(line1)
            res2, out2, is2 := get_next(line2)
            line1 = out1
            line2 = out2
            fmt.Printf("%s|%s|%t\n" ,res1, out1, is1)
            fmt.Println(res2, out2, is2)
            if is1 == false || is2 == false{
                break
            }
        }
        fileScanner.Scan()
    }
}
