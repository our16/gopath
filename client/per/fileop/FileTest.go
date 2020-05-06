package fileop

import (
	"log"
	"os"
)

type FileOp struct {

}

var (
	file *os.File
	err error
)

func ( fi *FileOp)Test(){
	file,err = os.Create("test.txt")
	if err != nil{
		log.Fatal(err)
	}
	file.Close()
}

type Goods struct {
	Name string
	Price int
}
type Book struct{
	Goods
}

func (g Goods)GetName() string  {
	return g.Name
}

func (b Book)GetName() string  {
	return "ok"
}