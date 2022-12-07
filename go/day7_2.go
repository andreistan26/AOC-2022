package main

import(
    "os"
    "bufio"
    "fmt"
    "strings"
    "strconv"
)

type File struct{
    name        string
    size        int
}

type Node struct{
    name        string
    parent      *Node
    children    []*Node
    files       []File
    sum         int
    vis         bool
}


type FileSystem struct{ 
    root        *Node
    pwd         *Node
}

var global_sum int
var smallest int
func (fs *FileSystem) init(){
    fs.root = &Node{
        name : "/",
        parent: nil,
        sum: 0,
        vis: false,
    }
    fs.pwd = fs.root
}

func (fs *FileSystem) cd(dir string){
    if dir == ".."{
        fs.pwd = fs.pwd.parent
        return
    }

    for _, child := range fs.pwd.children{
        if child.name == dir{
            fs.pwd = child
            return
        }
    }
}

func (node *Node) update_util() (sum int){
    sum = 0
    for _, file := range node.files{
        sum += file.size
    }
    if len(node.children) == 0 && node.vis == true{
        return sum
    }
    for _, child := range node.children{
        sum += child.update_util()
    }
    node.sum = sum
    global_sum += sum
    node.vis = true
    return sum
}

func (fs *FileSystem) update_size(){
    fs.root.update_util()
}

func (node *Node) walk_dirs(req int) {
    for _, child := range node.children{
        if child.sum < smallest && child.sum > req{
            smallest = child.sum
        }
        child.walk_dirs(req)
    }
}


func (fs *FileSystem) ls_line(line string){
    arr := strings.Split(line, " ")
    if strings.Contains(line, "dir"){
        new := &Node{
            name : arr[1],
            parent : fs.pwd,
        }
        fs.pwd.children = append(fs.pwd.children, new)
        return 
    }
    dim, _ := strconv.Atoi(arr[0])
    fs.pwd.files = append(fs.pwd.files, File{
        name:arr[1],
        size:dim,
    })
    fs.pwd.sum += dim
}


func main(){
    file, _ := os.Open("../res/input_7.txt")
    fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanLines)
    var fs FileSystem
    fs.init()

    in_ls := false
    for fileScanner.Scan(){
        arr := strings.Split(fileScanner.Text(), " ")
        if arr[0][0] == '$'{
            if strings.Compare(arr[1], "ls") == 0{
                in_ls = true
            } else if strings.Compare(arr[1], "cd") == 0{
                in_ls = false
                fs.cd(arr[2])
                continue
            }
        }
        if in_ls == true{
            fs.ls_line(fileScanner.Text())
        }
    }
    fs.update_size()
    smallest = 30000000
    free_space := 70000000 - fs.root.sum
    to_free := 30000000 - free_space
    fs.root.walk_dirs(to_free)
    fmt.Println(smallest)
}
