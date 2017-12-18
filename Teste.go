package main
import(
	"fmt"
	"math"
	"os"
	"bufio"
	"strings"
	"strconv"
)

type Tiro struct{
	x, y int
	z float64
}

func define(x int)[]Tiro{
	return make([]Tiro,x)
}
func defineInt(x int) []int{
	return make([]int,x)
}
func defz(x int,y int)float64{
	var b int
	b = x*x+y*y
	return math.Sqrt(float64(b))
}
var (
	newFile *os.File
	err error
)
func main() {
	file,_ := os.Open("teste.txt")

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	var linha = strings.Split(scanner.Text()," ")

	fmt.Println(len(linha))
	var circulos,_ = strconv.Atoi(linha[0])
	var tiros,_ = strconv.Atoi(linha[1])
	fmt.Println(circulos,tiros)
	var circ = defineInt(circulos)
	var acertos = define(tiros)
	for i:=0;i<circulos;i++{
		scanner.Scan()
		line := scanner.Text()
		circ[i],_= strconv.Atoi(line)

	}
	for i:=0;i<tiros;i++{
		scanner.Scan()
		line := strings.Split(scanner.Text()," ")
		acertos[i].y,_ =strconv.Atoi(line[0])
		acertos[i].x,_=strconv.Atoi(line[1])
		acertos[i].z=defz(acertos[i].x,acertos[i].y)
		fmt.Println(acertos[i].z)
		fmt.Printf("i = %f -> %d",acertos[i].z,i)
	}


	var contador int
	for i:=0;i<tiros;i++{
		for j :=0;j<circulos;j++{
			if acertos[i].z <= float64(circ[j]){
				contador++
			}
		}
	}
	fmt.Println(contador)
}
