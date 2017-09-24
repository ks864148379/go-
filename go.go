import fmt

func func_name() {
	
}
func GetName() (firstName, lastName, nickName string) {
	return "M", "R", "A"
}
_, _, nickName := GetName()
var v1 int = 10
var v2 = 10
v3 := 10
i, j = j, i
const Pi float64 = 3.141592653
const zero = 0.0
const (
	size int64 = 1024
	eof = -1
)
const u, v float32 = 0, 3
const mask = 1 << 3
const Home = os.GetEnv("HOME")
const (  // iota = 0
	c0 = iota // c0 = 0
	c1 = iota // c1 = 1
	c2 = iota // c2 = 2
)

const (
	a = 1 << iota // a = 1
	b = 1 << iota // b = 2
	c = 1 << iota // c = 4
)
const X = iota  //0
const Y = iota  //0
var c rune = 'a'
var p int64 = 3423
var v1 bool = true
var v2 bool = false
v3 := (1 == 2)

var b1 uint32 = 34234
var value1 int64
value3 := 64
//value1 = value3 error!
value1 = int64(value3)
i, j = 1, 2
if i == j {
	fmt.Println("equal")
}
import "math"

func IsEqual(f1, f2, p float64) bool {
	return math.Abs(f1-f2) < p
}

var str string
str = "Hello World"
ch := str[0]
str := "Hello"
n := len(str)
for i := 0; i < n; i++ {
	ch := str[i]
	fmt.Println(i, ch)
}
[2*N] struct { x, y int32 }
[1000]*float64 
arrLength := len(arr)
for i := 0; i < len(arr); i++ {
	fmt.Println()
}
for i, v := range arr {

}

package main

import "fmt"

func main() {
	var myArray [10]int = [10]int{1, 2, 3}
	var mySlice []int = myArray[:5]
	mySlice = myArray[:]
}

mySlice := make([]int, 5, 10)
for i, v := range mySlice {

}
fmt.Println(cap(mySlice))
mySlice = append(mySlice, 1, 2, 3)
mySlice = append(mySlice, mySlice2...)
oldSlice := []int{1, 2, 3}
newSlice := oldSlice[:3]

type PersonInfo struct {
	ID string
	Name string
	Address string
}

func main() {
	var personDB map[string] PersonInfo
	personDB = make(map[string] PersonInfo)

	personDB["123"] = PersonInfo{"123", "T", "R"}
	personDB["1"] = PersonInfo{"1", "J", "R"}

	person, ok := personDB["1234"]

	if ok {
		fmt.Println("F", person.Name, "1123")
	} else {
		fmt.Println("Didnt find 1123")
	}
}

var myMap map[string] PersonInfo

myMap = make(map[string] PersonInfo, 100)
myMap = map[string] PersonInfo {
	"123": PersonInfo {"1", "J", "R"},
}

delete(myMap, "1234")

value, ok := myMap["1234"]
if ok {

}

if a < 5 {
	return 0
} else {
	return 1
}

switch i {
case 0:
	fmt.Println("0")
default:
	fmt.Println("default")
}
sum := 0
for i := 0; i < 10; i++ {
	sum += i
}
sum := 0
for {
	sum++
	if sum > 100 {
		break
	}
}
a := []int{1, 2, 3}
for i, j := 0, len(a) - 1; i < j; i, j = i + 1, j -1 {
	a[i], a[j] = a[j], a[i]
}
for j := 0; j < 5; j++ {
	for i := 0; i < 10; i++ {
		if i > 5 {
			break JLoop
		}
	}
}
JLoop:

func myfunc() {
	i := 0
	HERE:
	fmt.Println(i)
	i++
	if i < 10 {
		goto HERE
	}
}

package mymath

import "errors"

func Add(a int, b int) (ret int, err error) {
	if a < 0 || b < 0 {
		err = error.New("ass")
		return
	}
	return a + b, nil
}

import "mymath"

c := mymath.Add(1, 2)

func myfunc(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}
myfunc(2, 3, 4)
myfunc(1, 3, 7, 13)

func myfunc(arg ...int) {
	myfunc(arg...)
	myfunc(arg[1:]...)
}

func Printf(format string, args ...interface{}) {
	
}

package main

import "fmt"

func MyPrintf(arg ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "ss")
		case string:
			fmt.Println(arg, "ff")
		case int64:
			fmt.Println(arg, "fa")
		}
	}
}

func main() {
	var v1 int = 1
	var v2 int64 = 234
	var v3 string = "hello"
	var v4 float32 = 1.123
	MyPrintf(v1, v2, v3, v4)
}

n, _ := f.Read(buf)

f := func(a, b int, z float64) bool {
	return a * b < int(z)
}
func(ch chan int) {
	ch <- ACK
} (reply_chan)

package main

import (
	"fmt"
)

func main() {
	var j int = 5

	a := func() (func()) {
		var i int = 10
		return func() {
			fmt.Println(i, j)
		}
	}()
	a()
	j *= 2
	a()
}

type error interface {
	Error() string
}

func Foo(param int) (n int, err error) {

}
n, err := Foo(0)

if err != nil {

} else {

}

if a < 5 {

} else {

}

type PathError struct {
	Op string
	Path string
	Err error
}
func (e *PathError) Error() string {

}
func Stat(name string) (fi FileInfo, err error) {
	var stat syscall.Stat_t

	err = syscall.Stat(name, &stat)

	if err != nil {
		return nil, 
	}
	return 
}

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

func main() {
	var a Integer = 1
	if a.Less(2) {
		fmt.Println(a, "Less 2")
	}
}

func (a *Integer) Add(b Integer) {
	*a += b
}

func main() {
	var a Integer = 1
	a.Add(2)
	fmt.Println("a =", a) // a = 3
}

a := 1
func (a Integer) Add(b Integer) {
	a += b
}
a.Add(2) // a = 1

type Header map[string][]string

func (h Header) Add(key, value string) {
	textproto.MIMEHeader(h).Add(key, value)
}

func (h Header) Set(key, value string) {
	textproto.MIMEHeader(h).Set(key, value)
}

b = a
b.Modeify()

type Rect struct {
	x, y float64
	width, heigth float64
}

func (r *Rect) Area() float64 {
	return r.width * r.heigth
}

rect1 := new(Rect)
rect2 := &Rect{}
rect3 := &Rect{0, 0, 100, 200}
rect4 := &Rect{width: 100, heigth: 200}

func NewRect(x, y, width, heigth float64) *Rect {
	return &Rect{x, y, width, heigth}
}

type Base struct {
	Name string
}

func (base *Base) Foo() {
	
}
func (base *Base) Bar() {
	
}

type Foo struct {
	Base
}

func (foo * Foo) Bar() {
	foo.Base.Bar
}

type Foo struct {
	*Base
}

type Job struct {
	Command string
	*log.Logger
}
func (job *Job) Start() {
	job.Log("start")
	job.Log("end")
}

type Rect struct {
	X, Y float64
	Width, Height float64
}
func (r *Rect) area() float64 {
	return r.Width * r.Height
}

goroutine
channel

type File struct {

}

func (f *File) Read(buf []byte) (n int, err error)
func (f *File) Write(buf []byte) (n int, err error)
func (f *File) Seek(buf []byte) (n int, err error)
func (f *File) Close(buf []byte) error

type IFile interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
	Seek(off int64, whence int) (pos int64, err error)
	Close() error
}

type IReader interface {
	Read(buf []byte) (n int, err error)
}

type IWriter interface {
	Write(buf []byte) (n int, err error)
}

type ICloser interface {
	Close() error
}

var file1 IFile = new(File)
var file2 IReader = new(File)
var file3 IWriter = new(File)
var file4 ICloser = new(File)

var v1 interface{} = 
switch v := v1.(type) {
	
}

func Add(x, y int) {
	z := x + y
	fmt.Println(z)
}

go Add(1, 1)

func main() {
	for i := 0; i < 10; i++ {
		go Add(i, i)
	}
}

func Count(lock *sync.Mutex) {
	lock.Lock()
	counter++
	fmt.Println(counter)
	lock.Unlock()
}
func main() {
	lock := &sync.Mutex{}

	for i := 0; i < 10; i++ {
		go Count(lock)
	}

	for {
		lock.Lock()

		c := counter

		lock.Unlock()

		runtime.Gosched()
		if c >= 10 {
			break
		}
	}
}


package main 

import "fmt"

func Count(ch chan int) {
	fmt.Println("counting")
	ch <- 1
}

func main() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}

	for _, ch := range(chs) {
		<-ch
	}
}

var chanName chan ElementType

var ch chan int //生命

var m map[string]  chan bool
ch := make(chan int) //定义

ch <- value
value := <-ch

ch := make(chan int, 1)
for {
	select {
		case ch <- 0 {

		}
		case ch <- 1 {

		}
	}
	i := <-ch
}

c := make(chan int, 1024)

for i := range c {
	fmt.Println("")
}

i := <-ch

timeout := make(chan bool, 1)
go func {
	time.Sleep(9)
	timeout <- true
}()

select {
	case <- ch {

	}
	case <- timeout {

	}
}

type PipeData struct {
	value int
	handler fun(int) int
	next chan int
}

func handle(queue chan *PipeData) {
	for data := range queue {
		data.next <- data.handler(data.value)
	}
}

var ch2 chan<- float64
var ch3 <-chan int

ch4 := make(chan int)

ch5 := <-chan int(ch4)
ch6 := chan<- int(ch4)

func Parse(ch <- chan int) {
	for value := range ch {
		fmt.Println("Parse", value)
	}
}

close(ch)

x, ok := <-ch
type Vector []float64 

func (v Vector)DoSome(i, n int, u Vector, c chan int) {
	for ; i < n; i++ {
		v[i] += u.Op(v[i])
	}
	c <- 1
}

const NCPU = 16

func (v Vector)DoAll(u Vector) {
	c := make(chan int, NCPU)
	for i := 0; i < NCPU; i++ {
		go v.DoSome(i*len(v)/NCPU, (i+1)*len(v)/NCPU, u, c)
	}
}


















