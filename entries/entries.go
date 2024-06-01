package entries 
import "fmt"
type File struct {
    Name string;
    Content string;
}


type Directory struct {
    Name string;
    Files map[string]*File;
    Directories map[string]*Directory;
}

func RootDir () *Directory {
    return &Directory{
        Name:"/",
        Directories: make(map[string]*Directory),
        Files:make(map[string]*File),
    }
}

func NewDir (name string) *Directory {
    return &Directory{
        Name:name,
        Directories: make(map[string]*Directory),
        Files:make(map[string]*File),
    }
}


func (d *Directory)Delete(name string) {
    _, ok := d.Directories[name]
    if !ok{
        fmt.Printf("Directory %s not foudn\n", name)
        return 
    }

    delete(d.Directories, name)
}


func (d *Directory) Create(name string) {
    _, ok := d.Directories[name]
    if ok{
        fmt.Println("Dir with the same name already exists")
        return 
    }

    newDir := &Directory{
        Name:name,
        Directories:make(map[string]*Directory),
        Files:make(map[string]*File), 
    }

    d.Directories[name] = newDir
    fmt.Printf("New dir named %s has been created\n", name)
}


func (d *Directory) CreateFile(name string){
    _, ok := d.Files[name]
    if ok{
        fmt.Println("A file with the same name already exists")
        return 
    }

    newFile := &File{
        Name:name,
        Content:"",
    }
    d.Files[name] = newFile 
    fmt.Printf("%s has been created\n", name)
}

func (d *Directory) DeleteFile(name string){
    _, ok := d.Files[name]
    if !ok{
        fmt.Println("File doesn't exit")
        return 
    }

    delete(d.Files, name)
    fmt.Println("File has been removed successfully")
    
}




func (d *Directory) UpdateDir(name string, newName string){
    dir, ok := d.Directories[name]
    if !ok{
        fmt.Println("Directory doesn't exit")
        return 
    }

    dir.Name = newName;
    d.Directories[newName] = dir
    delete(d.Directories, name)
}

func (d *Directory) UpdateFile(name string, newName string){
    file, ok := d.Files[name]
    if !ok{
        fmt.Println("File doesn't exit")
        return 
    }

    file.Name = newName;
    d.Files[newName] = file
    delete(d.Files, name)
} 
