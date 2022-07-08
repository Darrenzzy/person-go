package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	urls "net/url"
	"reflect"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"

	"github.com/Darrenzzy/person-go/structures"
)

type Fn struct {
	A string
	B string
	C string
	D string
}
type baz struct {
	bar int
	foo int
}

type arrStruct []baz

func TestSliceInfo(t *testing.T) {
	sliceint := make([]int64, 45000)                                                     // 指向元素类型为int32的1000个元素的数组的切片
	fmt.Println("Size of []int32:", unsafe.Sizeof(sliceint))                             // 24
	fmt.Println("Size of [1000]int32:", unsafe.Sizeof([1000]int64{}))                    // 4000
	fmt.Println("Real size of s:", unsafe.Sizeof(sliceint)+unsafe.Sizeof([1000]int64{})) // 4024
}

func TestStringToByte(t *testing.T) {
	buf := make([]byte, 1)
	var r io.Reader

	r = &bytes.Buffer{}

	n, err := r.Read(buf)
	t.Log(n, err)

}

func TestShanValue(t *testing.T) {
	brokers := strings.Split("10.7.68.185:9092,10.7.68.186:9092,10.7.68.188:9092", ",")
	t.Log(brokers)
}

func TestSelectToGo(t *testing.T) {

	a := 2

dodo:
	println(0)

	switch {
	case a < 5:
		println(5)

	case 1 == a:
		println(1)
	case 2 == a:
		a = 1
		println(2)
		goto dodo
	default:
	}

}

func TestMapDel(t *testing.T) {
	var p map[string]string
	p = nil
	delete(p, "a")

	ii := int64(123_2123)
	println(ii)
}

// 洗牌算法
func TestRandRange(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	for i := 0; i < len(arr)-1; i++ {
		rand.Seed(time.Now().UnixNano())
		a := rand.Intn(len(arr) - i)
		arr[i], arr[a] = arr[a], arr[i]
	}

	t.Log(arr)
}

func TestUrlname(t *testing.T) {

	name := "%E6%98%A5%E5%A4%A9%E5%8F%AF%E7%9C%9F%E6%98%AF%E4%B8%AA%E5%B0%8F%E8%AE%A8%E5%8E%8C%E9%AC%BC%EF%BC%8C%E5%9C%A8%E6%88%91%E5%BF%83%E9%87%8C%E5%81%B7%E5%81%B7%E5%85%BB%E4%BA%86%E4%B8%80%E5%8F%AA%E5%B0%8F%E9%B9%BF%EF%BC%8C%E5%B0%B1%E6%92%92%E6%89%8B%E4%B8%8D%E7%AE%A1%E4%BA%86%E2%9C%A8%E2%9C%A8%E2%9C%A8"
	// escapeUrl := urls.QueryEscape(name)

	t.Log(urls.QueryUnescape(name))
	name = "darren is  iphone"
	t.Log(urls.QueryUnescape(name))

}

func TestBuInterface(t *testing.T) {
	arr := []Buiding{House{}, Shop{}, Toilet{}}

	for _, v := range arr {
		v.Builds()
	}
}

func TestTickerFor(t *testing.T) {

	tt := time.NewTicker(time.Second * 1)
	i := 1
	for ; true; <-tt.C {
		fmt.Println("tick", i)
		i++
	}

}

func TestByte(t *testing.T) {
	ss := "dHJ1ZQ=="
	ss = "d"
	var vv interface{}
	err := json.Unmarshal([]byte(ss), &vv)
	t.Log(vv, err)
}
func TestJsonByte(t *testing.T) {
	ss := NewA()
	b, _ := json.Marshal(ss)
	t.Log(b, string(b))
	t.Log(b, *(*string)(unsafe.Pointer(&b)))
	b = nil
	t.Log(b, 22, *(*string)(unsafe.Pointer(&b)))

}

// -gcflags '-m -l'
func TestAddPoint(t *testing.T) {
	println(addV1(1, 2))
	println(addV2(1, 2))
}
func addV2(a, b int) (add *int) {
	add = new(int)
	*add = a + b
	return add
}
func addV1(a, b int) *int {
	add := 0
	add = a + b
	return &add
}

func TestGoPanic(t *testing.T) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				t.Log("panic 11111", err)
			}
		}()
		i := 1
		for i < 10 {
			i++
			if i == 4 {
				panic(2)
			}
			time.Sleep(time.Second)
		}
		t.Log("end 1")

	}()

	go func() {
		defer func() {
			if err := recover(); err != nil {
				t.Log("panic 11111", err)
			}
		}()
		i := 1
		for i < 5 {
			i++
			t.Log(i)
			time.Sleep(time.Second)
		}
		t.Log("end 2")
	}()
	ch2 := make(chan int)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				t.Log("panic 11111", err)
			}
		}()
		i := 1
		for {
			ch2 <- i
			i++
			t.Log(333)
			time.Sleep(time.Second * 10)

		}
	}()
	go func() {
		defer func() {
			if err := recover(); err != nil {
				t.Log("panic 11111", err)
			}
		}()
		i := 1
		for {
			t.Log(<-ch2, i)
		}
	}()
	ch := make(chan int)
	ch <- 1

}

func TestTimeDaysAdd(t *testing.T) {
	t.Log(time.Now().String())
	t.Log(time.Now().GoString())
	now := time.Now().Unix()
	t.Log(now, now+3600*24)
	tt := time.Unix(1697839585, 0)
	t.Log(tt)
	ts := time.Now()
	tm1 := time.Date(ts.Year(), ts.Month(), ts.Day(), 0, 0, 0, 0, ts.Location())
	// tm2 := tm1.AddDate(0, 0, 1)
	t.Log(tm1, tm1.Unix())

	// 7天前0点
	day7 := ts.AddDate(0, 0, -7)
	before7Day := strconv.FormatInt(time.Date(day7.Year(), day7.Month(), day7.Day(), 0, 0, 0, 0, ts.Location()).Unix(), 10)
	t.Log(before7Day)

}

func TestSleepSpeed(t *testing.T) {
	t.Log("sleep speed test")
	start := time.Now()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 300)
		// time.Sleep(time.Millisecond * 2)
	}
	t.Log(time.Now().Sub(start))
}

func TestContains(t *testing.T) {
	// t.Log(strings.Contains(strings.ToLower("shenZhou"), "shenzhou"))
	// t.Log(strings.Contains("darren91231231i1892", "darren"))
	t.Log(getLatestIpAddr("127.0.0.1, 115.171.170.95, 10.5.12.212"))
}

func getLatestIpAddr(clientIP string) string {
	if index := strings.LastIndex(clientIP, ","); index >= 0 {
		clientIP = clientIP[index+1:]
	}
	clientIP = strings.TrimSpace(clientIP)
	if len(clientIP) > 0 {
		return clientIP
	}
	return ""
}

func (b arrStruct) Len() int {
	return len(b)
}

func (b arrStruct) Less(i, j int) bool {
	if b[i].bar < b[j].bar {
		return true
	}

	if b[i].bar == b[j].bar {
		return b[i].foo > b[j].foo
	}

	return false
}

func (b arrStruct) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func TestSortSlice(t *testing.T) {
	s := []baz{
		{5, 4},
		{6, 7},
		{2, 3},
		{6, 4},
	}
	sort.Sort(arrStruct(s))
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", s)
}

func TestWriteSlice(t *testing.T) {
	app := make([]int64, 0, 1000)
	var lock sync.RWMutex
	// wgs := sync.WaitGroup{}
	wg.Add(10)
	go func() {
		i := 0
		for {
			lock.Lock()
			app = append(app, rand.Int63())
			lock.Unlock()
			if i%10000000 == 0 {
				t.Log(len(app))
			}
			i++
		}
	}()

	go func() {
		tic := time.NewTicker(100 * time.Microsecond)
		for {
			select {
			case <-tic.C:
				println(len(app), 99999999)
				lock.Lock()
				app = make([]int64, 0, 1000)
				lock.Unlock()

				println(len(app), 9888888888)

			}
		}
	}()
	wg.Wait()

}

func TestChanCtx(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 1*time.Second)
	ctx = context.WithValue(ctx, "key", "value")
	defer cancel()

	go handle(ctx, 500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())

	}
	t.Log(ctx.Value("key"))
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}

func TestBoolSize(t *testing.T) {
	aa := make([]bool, 0, 10)
	aa = append(aa, true)
	aa = append(aa, true)
	println(len(aa), cap(aa), unsafe.Sizeof(aa), &aa)
	aa = aa[0:1]
	println(len(aa), cap(aa), unsafe.Sizeof(aa), &aa)
	aa = aa[0:2]
	println(len(aa), cap(aa), unsafe.Sizeof(aa), &aa)

	var b = make(map[int64]bool, 0)
	if unsafe.Sizeof(b) != 1 {
		t.Errorf("bool size is %d, want 1", unsafe.Sizeof(b))
	}
}

type CommBody interface {
	GetA() string
	keys() []string
}

func (fn Fn) GetA() string {
	return fn.A
}
func (fn Fn) keys() []string {
	return []string{"A", "B", "C", "D"}
}

func DefClass(param CommBody) {
	ms, _ := json.Marshal(param)
	fmt.Println(string(ms))

}

func TestClassDef(t *testing.T) {
	var fn = Fn{A: "a", B: "b", C: "c", D: "d"}
	DefClass(fn)
}

func Example_Print() {
	score := []int{1, 2, 3}
	fmt.Println(score)
	// Output: [1 2 3]
}

func Test_structChan(t *testing.T) {

	var v BigBar
	m := make(chan struct{})
	t.Log(unsafe.Sizeof(v))
	go func(a *BigBar) {
		_ = a
		m <- struct{}{}
	}(&v)
	<-m
}

func BenchmarkDirConcat(b *testing.B) {
	var s37 = []byte{36: 'x'} // len(s37) == 37
	var str string
	for i := 0; i < b.N; i++ {
		str = string(s37) + string(s37)
	}
	_ = str
}

func BenchmarkSplitedConcat(b *testing.B) {
	var s37 = []byte{36: 'x'} // len(s37) == 37
	var str string

	for i := 0; i < b.N; i++ {
		str = string(s37[:32]) +
			string(s37[32:]) +
			string(s37[:32]) +
			string(s37[32:])
	}
	_ = str
}

func TestFloatTostring(t *testing.T) {
	f := float64(23.434532)
	t.Logf("%f", f)
	t.Logf("%.2f", f)
}

func TestCompareInterface(t *testing.T) {
	var x interface{} = []int{1, 2}
	var y interface{} = map[string]int{"aa": 1, "bb": 1}
	var z interface{} = func() {}

	// The lines all print false.
	println(x == y)
	println(x == z)
	println(x == nil)
	t.Log(x, y, z)

	// Each of these line could produce a panic.
	// println(x == x)
	// println(y == y)
	// println(z == z)
}

func TestSortInt(t *testing.T) {

	arr := []int{1, 44, 2, 77, 3, 4, 5}

	s := sort.SearchInts(arr, 22)
	t.Log(s, arr)

}

func TestSliss(t *testing.T) {
	arr := []int{1}

	go func() {
		for i := 100; i > 0; i-- {
			println(len(arr))
			arr = append(arr, 2)
		}
	}()
	time.Sleep(time.Second)
	for len(arr) > 0 {
		l := 10
		if l > len(arr) {
			l = len(arr)
		}
		_ = arr[:l]
		t.Log(len(arr), l)
		arr = arr[l:]

	}

	t.Log(arr)

}

func TestAppends(t *testing.T) {

	arr := []int{}
	ch := make(chan struct{}, 0)
	isClose = false
	i := int32(0)
	maxCount := int32(5000)
	println(runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
	go func(i int32) {
		for {
			if i > maxCount {
				ch <- struct{}{}
				return
			}
			atomic.AddInt32(&i, 1)
			arr = append(arr, int(i))
			if len(arr)%100 == 0 {
				fmt.Println(i)
			}
		}
	}(i)
	go func(i int32) {
		for {
			if i > maxCount {
				ch <- struct{}{}
				return
			}
			atomic.AddInt32(&i, 1)
			arr = append(arr, int(i))
			if len(arr)%100 == 0 {
				fmt.Println(i)
			}
		}
	}(i)
	go func(i int32) {
		for {
			if i > maxCount {
				ch <- struct{}{}
				return
			}
			atomic.AddInt32(&i, 1)
			arr = append(arr, int(i))
			if len(arr)%100 == 0 {
				fmt.Println(i)
			}
		}
	}(i)
	for {
		select {
		case <-ch:
			isClose = true
		default:
			if len(arr) > 20 {
				_ = arr[:20]
				// fmt.Println(ss)
				arr = arr[20:]
			}

		}

		if isClose && len(arr) == 0 {
			return
		} else if isClose && len(arr) < 20 {
			ss := arr[:]
			fmt.Println("***", ss)
			arr = arr[len(arr):]
		}

	}

}

func TestWgAdd(t *testing.T) {
	var wgs sync.WaitGroup

	k := uint64(1)
	md := make([]int, 1000)

	for i := int(1); i < 1000; i++ {
		md[i] = i
		wgs.Add(1)
		go func(i int) {
			md[int(k)] = int(i)
			atomic.AddUint64(&k, 1)
			wgs.Done()

		}(i)
	}

	wgs.Wait()
	t.Log(md, len(md))
	md = md[:k]
	t.Log(md, len(md))
}

func TestMapRead(t *testing.T) {
	m := map[string]int{}
	m1 := map[string]int{}
	go func() {
		for {
			m1["key1"] = 1
			m = m1
		}
	}()
	for {
		_ = m["key2"]
	}
}

func TestParams(t *testing.T) {

	B := Fn{}
	B.A = "qqqq"
	P(B)
	B.B = "qqqq"
	t.Log(B.A, B)

}
func P(f Fn) {
	f.A = "aaa"
	f.B = "aaa"

}

func TestDeferSort(t *testing.T) {
	x := 1
	y := AddD(&x)
	println(*y, x)

	x1 := 1
	y1 := AddE(&x1)
	println(x1, y1)

}
func AddD(a *int) *int {

	println(a, *a, 111)
	defer func() {
		*a = *a + 1
	}()
	println(a, *a, 333)

	return a
}

func AddE(a *int) *int {

	println(a, *a, 111)
	defer func() {
		*a = *a + 1
	}()
	println(a, *a, 333)

	return nil
}

func TestSliceEqual(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 3, 2, 4}
	c := []int{1, 2, 3, 4}
	d := []byte{1, 2, 3, 4}
	fmt.Println(reflect.DeepEqual(a, b))
	fmt.Println(reflect.DeepEqual(a, c))
	fmt.Println(reflect.DeepEqual(a, d))

}

func TestStructNil(t *testing.T) {
	D := new(PayWay)
	t.Log(D.Test)
	t.Log(D.Test2)
	t.Log(D.Test.S)
	t.Log(D.Test.S.Alias)
	// 指针 就是nil  会panic！
	if D.Test2 != nil {
		t.Log(D.Test2.Alias)
	}

}

func TestFloat32To64(t *testing.T) {

	f := float64(44.532424234234)
	t.Log(float32(f))

	f = float64(445324242342.34)
	t.Log(float32(f))

	f = float64(117.11166743741192)
	t.Log(float32(f))

	f = float64(99999.532424234234)
	t.Log(float32(f))
}

func TestSliceSplit(t *testing.T) {

	aa := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// ss := aa[2:3000]
	ss := aa[2:2]

	var name string
	name = "asSSd"
	aaaa := strings.Count(name, "")

	t.Log(strings.ToUpper(name))
	t.Log(aaaa)
	t.Log(ss, 2222)
}

func TestAaa(t *testing.T) {
	// // target := 8
	// target := 6
	// // target := 1
	// // arrStruct := []int{1, 2, 3, 3, 3, 3, 3, 3, 9, 10}
	// arrStruct := []int{5, 7, 7, 8, 8, 10}
	// // arrStruct := []int{1}
	// fmt.Println(searchRange(arrStruct, target))
	// beginningOfTime := time.Unix(time.Now().Unix(), 0)
	beginningOfTime := time.Unix(99999123123, 0)
	fmt.Println(beginningOfTime.Unix())
}

const url = "https://github.com/EDDYCJY"

func TestForRange(t *testing.T) {

	i := 0
	for {
		t.Log(time.Now())
		timer := time.NewTimer(time.Second)
		go func() {
			defer func() {
				if r := recover(); r != nil {
					t.Error("Recovered in f", r)
				}
			}()
			select {
			case <-timer.C:
				t.Log(time.Now(), "###", i)
				i++
			}

			if i > 500 {
				panic("444")
			}
		}()

	}

}

func TestTimeMicr(t *testing.T) {
	t.Log(strconv.FormatInt(time.Now().Unix(), 10))
	t.Log(time.Now().Unix() * 1000)
}

func TestAdd(t *testing.T) {
	s := Add(url)
	if s == "" {
		t.Errorf("Test.Add error!")
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(url)
	}
}

func TestSpilt(t *testing.T) {
	msg := "aaskjdhakshdlkhsada审"
	t.Log(len(strings.Split(msg, "审批于")))
	t.Log(strings.Contains("15618802115 18658852500", "18658852500"))

}

func TestPanicV4(t *testing.T) {

	type R struct {
		S *int64
		K string
	}

	w := int64(2)
	aa := R{
		S: &w,
		K: "123",
	}
	if aa.S == nil || *aa.S == 0 {
		println(123)
	}
	aa.S = nil
	bb, err := aa, errors.New("123")
	if err == nil || bb.K == "123" {
		println("ppppp")
	}

}

func TestRangeNil(t *testing.T) {
	obj := make([]string, 0)
	obj = nil
	// obj = append(obj, "123")
	for v := range obj {
		println(v)
	}
}

func TestMapv2(t *testing.T) {
	m := make(map[string][]int)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%p---%v\n", m, m)
	m["test"] = s
	fmt.Printf("%p---%v\n", s, s)
	fmt.Printf("%p---%v\n", m["test"], m["test"])
	fmt.Printf("%p---%v\n", m, m)

}

func TestSliceV2(t *testing.T) {
	s := make([]int, 1)
	s[0], s, s[0] = 333, []int{1, 2, 3}, 222
	t.Log(s)
}

func TestFnLoop(t *testing.T) {

	aa := new(Fn)
	aa.Geta()
	aa.Getb()
	aa.Getc()
	aa.Getd()
	aa.Geta()

}
func (f *Fn) Geta() string {
	return f.A
}

func (f *Fn) Getb() string {
	return f.B
}
func (f *Fn) Getc() string {
	return f.C
}
func (f *Fn) Getd() string {
	return f.D
}

func TestNilFun(t *testing.T) {

	a := NewA()
	c := context.TODO()
	fmt.Println(a.GetName(&c, "222"))

}

type A struct {
	name  string
	Alias string
}

func NewA() *A {
	return &A{
		name: "111",
	}
}

func (a *A) GetName(ctx *context.Context, name string) string {
	a.name = name
	return a.name
}

func WhichIsBest() int {
	a, b, c, d := rand.Intn(10), rand.Intn(10), rand.Intn(10), 0
	switch {
	case a == 1:
		d = 1
	case b == 1:
		d = 2
	case c == 1:
		d = 3
	default:
		d = 4
	}
	return d
}

func WhichIsBestV2() int {
	a, b, c, d := rand.Intn(10), rand.Intn(10), rand.Intn(10), 0
	switch {
	case a == 1:
		d = 1
		return d
	case b == 1:
		d = 2
		return d
	case c == 1:
		d = 3
		return d
	}
	return d
}

type SR string

func TestPoint(t *testing.T) {
	var vv = SR("初始值")
	d := vv
	d.myVal()

	d.Get1()
	d.myVal()

	vv = "2次"
	d = vv
	d.Get2()
	d.myVal()

	d.Get4()
	d.myVal()

	d.Get3()
	d.myVal()
}

func (s *SR) Get1() {
	// s = nil
	// s.myVal()
}

func (s SR) Get2() {
	s = SR("期望的值2")
	s.myVal()

}
func (s *SR) Get3() {
	v := SR("期望的值3")
	s = &v
	s.myVal()
}
func (s *SR) Get4() {
	v := SR("期望的值4")
	*s = v
	s.myVal()
}

func (s *SR) myVal() {
	fmt.Printf("the val : %p %s \n", s, *s)
}

func TestAffirmation(t *testing.T) {
	var a = uint8(90)
	println(int64(a))
	println(int8(a))
	var m interface{}
	m = a

	if s, ok := m.(int64); !ok {
		println(s)
	}

}

func TestSwitchs(t *testing.T) {
	aa := structures.Interval{Start: 123, End: 333}
	switch {
	case aa.End == 333:
		println(2)
	case aa.Start == 23:
		println(3)
	case aa.End == 123:
		println(4)
	default:
		println(5)

	}
}

func regexMatch(regex string, operation string) bool {
	r, err := regexp.Compile(regex)
	if err != nil {
		return false
	}
	return r.FindString(operation) == operation
}

func TestRegex(t *testing.T) {

	println(regexMatch(`^.*login.*$`, "asdkalogin/hsj"))
	println(regexMatch(`^.*login.*$`, "1qweqwi"))
	println(regexMatch(`.*2014.*$`, "1qwe[2014]qwi"))
	println(regexMatch(`^.*TenantSso/Login/.*$`, "/helloworld/aaa?asdasdjk"))
}

func TestArrayGroup(t *testing.T) {
	// 	原来 arrStruct[ "qwe","weq","wqe","abc","cba"]
	// 	期望 arrStruct[["qwe","weq","wqe"],["abc","cba"]]

	arr := []string{"qwe", "weq", "wqe", "abc", "cba"}

	// pp := make(map[int32][][26]int, 10)

	println(arr, 'a'-97)

	for _, v := range arr {
		sum := int32(0)
		m := [26]int{}
		for _, vv := range v {
			sum += vv
			m[vv-97]++
		}

	}

	// fmt.Println(m, pp)
	// p := make(map[byte][]map[byte]bool, 0)
	//
	// for _, v := range arrStruct {
	// 	m := make(map[byte]bool, 0)
	// 	var s byte
	// 	for k := range v {
	// 		m[v[k]] = true
	// 		s += v[k]
	// 	}
	// 	if vv, ok := p[s]; ok {
	// 		vv = append(vv, m)
	// 		p[s] = vv
	// 	} else {
	// 		p[s] = []map[byte]bool{m}
	// 		println(s)
	// 	}
	// }
	// fmt.Println(p)

}

// 当前用例
// 	期望： 1,8，2，7，3，6,4，5
func TestReverseV3(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	list := structures.Ints2List(arr)
	f := list
	s := list
	// 快慢指针取到中点
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
	}
	// 翻转慢 list
	mm := structures.Reverse(s)
	aaa := &structures.ListNode{}
	res := aaa
	i := 0
	// 交替 append 到新链表中，直到完成
	for list != nil && mm != nil {
		if i&1 == 0 {
			res.Next = &structures.ListNode{Val: list.Val}
			list = list.Next
		} else {
			res.Next = &structures.ListNode{Val: mm.Val}
			mm = mm.Next
		}
		res = res.Next
		i++
	}
	structures.Travel(aaa.Next)
}
func TestZigzagLevelOrderV2(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	tree := structures.Ints2TreeNode(arr)

	stack := []*structures.TreeNode{tree}

	aas := make([][]int, 0)
	isSS := false
	level := 0
	for len(stack) != 0 {
		i := 0

		if len(aas) <= level {
			aas = append(aas, []int{})
		}
		l := len(stack)
		for l > i {
			tt := stack[0]
			if !isSS {
				aas[level] = append(aas[level], tt.Val)
			} else {
				aas[level] = append([]int{tt.Val}, aas[level]...)
			}

			stack = stack[1:]
			if tt.Left != nil {
				stack = append(stack, tt.Left)
			}
			if tt.Right != nil {
				stack = append(stack, tt.Right)
			}
			i++
		}
		if isSS {
			isSS = false
		} else {
			isSS = true
		}
		level++
	}
	fmt.Println(aas)
}

func TestArraySum(t *testing.T) {
	A := []int{1, 2, 3, 4, 5, 0, 7}
	B := []int{6, 7, 0}

	la := len(A) - 1
	lb := len(B) - 1
	if la > lb {

	}
	x := 0
	for k := range A {
		if lb-k >= 0 {
			A[la-k] += B[lb-k]
		}
		A[la-k] = A[lb-k] + x
		x = 0
		if A[la-k] >= 10 {
			A[la-k] %= 10
			x = 1
		}
	}
	if x > 0 {
		for i := la; i <= lb; i++ {
			B[lb-i] += x
			x = 0
			if B[lb-i] >= 10 {
				B[lb-i] %= 10
				x = 1
			}
		}
	}

	if x > 0 {

	}
}

func TestPractice(t *testing.T) {
	// HeapSoft([]int{1, 888, 11, 2, 44, 3, 777, 4, 55, 5, 67})
	singleNumber([]int{1, 2, 3, 4, 3, 2, 4, 1})
}
func singleNumber(nums []int) int {
	// bit:=
	a := 0
	for v := range nums {
		a ^= nums[v]
	}
	println(a)
	return a
}

func HeapSoft(arr []int) {
	l := len(arr)
	fmt.Println(arr)

	for i := l / 2; i >= 0; i-- {
		BuildHeapV2(arr, i, l)
	}

	l--
	for i := l; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		BuildHeapV2(arr, 0, l)
		l--
	}
	fmt.Println(arr)
}

func BuildHeapV2(arr []int, n, lens int) {
	k := n
	for n < lens {
		i := n*2 + 1
		j := i + 1
		if i < lens && arr[k] < arr[i] {
			k = i
		}
		if j < lens && arr[k] < arr[j] {
			k = j
		}
		if k != n {
			arr[k], arr[n] = arr[n], arr[k]
			n = k
		} else {
			// n = n * 2
			break
		}
		// println(n, j, i)
	}

}

func xRuntime() {
	runtime.Gosched()                                    // 切换任务
	fmt.Println("cpus:", runtime.NumCPU())               // 返回当前系统的CPU核数量
	fmt.Println("goroot:", runtime.GOROOT())             //
	fmt.Println("NumGoroutine:", runtime.NumGoroutine()) // 返回真该执行和排队的任务总数
	fmt.Println("archive:", runtime.GOOS)                // 目标操作系统
}

func smallestDistancePair(nums []int, k int) int {
	keys := make(map[int]int, 0)
	arr := make([]int, 0)
	l := len(nums)
	for i := range nums {
		for j := i + 1; j < l; j++ {
			diff := nums[j] - nums[i]
			if diff < 0 {
				diff = ^diff + 1
			}
			keys[diff]++
			// arrStruct = mergeappend(arrStruct, diff)

		}
	}
	fmt.Println(arr, k, keys)
	for m := range keys {
		arr = mergeappend(arr, m)
	}

	if len(arr) > 0 {

	}

	return 0
}

func mergeappend(arr []int, r int) []int {
	for i := 0; i < len(arr); i++ {
		if arr[i] >= r {
			arr = append(arr[:i], append([]int{r}, arr[i:]...)...)
			return arr
		}
	}
	arr = append(arr, r)

	return arr
}

func TestCase(t *testing.T) {

	x := 24
	a := 14
	b := -10
	// s:=b^b
	println(a|b, x^a^b, ^b+1)
	fmt.Printf("%b \n", x)
	fmt.Printf("%b \n", a)
	fmt.Printf("%b \n", b)
	fmt.Printf("%b \n", a&b)

}

func TestTwoSum(t *testing.T) {
	res := twoSum([]int{20, 70, 20, 150}, 220)
	t.Log(res)
}

func TestMaxLeng(t *testing.T) {
	// t.Log(maxLength([]int{2, 3, 4, 1, 5}))
	// t.Log(maxLength([]int{2, 2, 3, 4, 1, 5}))
	t.Log(maxLength([]int{1, 1, 1, 2, 2, 3, 4, 5, 6, 7, 7, 8, 9}))
}

func TestIsValid(t *testing.T) {
	// t.Log(isValid("[](){}"))
	// t.Log(isValid("[]({}[]{{{}}}){}"))
	// t.Log(isValid("{[}]"))
	// t.Log(isValid("]"))
	// t.Log(Fibonacci(4))
	// t.Log(FibonacciV2(4))
	// t.Log(Fibonacci(10))
	// t.Log(FibonacciV2(30))
	// t.Log(search([]int{1, 2, 2, 3, 3, 6, 8, 8, 8, 9, 9, 9}, 6))
	// t.Log(search([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 6))
	// t.Log(search([]int{1, 1, 1, 1, 6, 6, 6, 6, 6, 7, 8, 9, 10, 10, 101}, 101))
	// t.Log(search([]int{-2, 1, 2}, -2))
	t.Log(LRU([][]int{{1, 1, 1}, {1, 2, 2}, {1, 3, 2}, {2, 1}, {1, 4, 4}, {2, 2}}, 3))
}

type Lru struct {
	md   map[int]*node
	buf  int
	max  int
	head *node
	tail *node
}

type node struct {
	pre, next *node
	val, key  int
}

func LRU(operators [][]int, k int) []int {
	// write code here
	lru := initLru(k)
	res := []int{}
	for i := range operators {
		if operators[i][0] == 1 {
			lru.set(operators[i][1], operators[i][2])

		} else {
			res = append(res, lru.get(operators[i][1]))
		}
	}
	return res
}
func initLru(k int) Lru {
	return Lru{
		md:  make(map[int]*node),
		max: k,
	}
}

func (this *Lru) get(k int) int {
	if v, ok := this.md[k]; ok {
		this.remove(v)
		this.add(v)
		return v.val
	}
	return -1
}
func (this *Lru) set(k, x int) {
	if v, ok := this.md[k]; ok {
		this.remove(v)
		this.add(v)
		return
	} else {
		n := &node{val: x, key: k}
		this.md[k] = n
		this.add(n)

	}

	if len(this.md) > this.max {
		delete(this.md, this.tail.key)
		this.remove(this.tail)
	}

}

func (this *Lru) remove(n *node) {
	if this.head == n {
		this.head = n.next
		this.head.pre = nil
		return
	}

	if this.tail == n {
		this.tail = n.pre
		n.pre.next = nil
		n.pre = nil
		return

	}

	n.pre.next = n.next
	n.next.pre = n.pre

	return

}

func (this *Lru) add(n *node) {
	n.next = this.head
	if this.head != nil {
		this.head.pre = n
	}
	this.head = n
	if this.tail == nil {
		this.tail = n
		this.tail.next = nil
	}

	return
}

func search(nums []int, target int) int {
	// write code here
	l := len(nums) - 1
	i := 0
	mid := 0
	for i <= l {
		mid = int(uint(i+l) >> 1)
		// fmt.Println(mid, i, l)
		if nums[mid] == target {
			for mid > 1 && nums[mid-1] == target {
				mid--
			}
			return mid
		}
		if nums[mid] < target {
			i = mid + 1
		} else {
			l = mid - 1

		}
	}

	return -1
}

func FibonacciV2(n int) int {

	if n > 40 || n < 0 {
		return n
	}

	arr := [40]int{}
	arr[0] = 0
	arr[1] = 1

	for i := 2; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[n]

}
func Fibonacci(n int) int {
	if n >= 2 {
		return Fibonacci(n-1) + Fibonacci(n-2)
	}
	if n == 1 {
		return 1
	}
	if n == 0 {
		return 0
	}
	return 0
	// write code here
}

// 二分查找
func TestSearch(t *testing.T) {
	arr := []int{1, 2, 3, 3, 3, 3, 5, 6, 7, 8, 9, 9, 9, 9, 99}
	target := 9
	low, fast := 0, len(arr)-1
	for low <= fast {
		mid := len(arr) - (fast-low)>>1
		if target > arr[mid] {
			low = mid
		} else if target < arr[mid] {
			fast = mid
		} else {
			for mid < len(arr)-2 && arr[mid+1] == target {
				mid++
			}
			println(arr[mid], mid)
			break
		}
	}
}

func isValid(s string) bool {
	mp := make(map[uint8]uint8, 3)
	mp['['] = ']'
	mp['{'] = '}'
	mp['('] = ')'
	stack := make([]uint8, 0)

	for i := range s {
		if v, ok := mp[s[i]]; ok {
			stack = append(stack, v)
		} else {
			if len(stack) > 0 && stack[len(stack)-1] == s[i] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
		fmt.Println(stack)
	}

	return len(stack) == 0
}
func maxLength(arr []int) int {
	long, i, r := 0, 0, 0
	l := len(arr)
	as := [256]byte{}
	for i < l {
		if r == l {
			return long
		}
		if as[arr[r]] == 0 {
			as[arr[r]]++
			r++
		} else {
			as[arr[i]]--
			i++
		}
		long = max(long, r-i)
		fmt.Println(i, r, long)
	}

	return long
}
func twoSum(numbers []int, target int) []int {

	l := len(numbers)
	for k := range numbers {
		for j := 1; j < l; j++ {
			if k != j && target == numbers[k]+numbers[j] {
				return []int{k + 1, j + 1}
			}
		}

	}
	return []int{}

	// write code here
}

// 归并排序 不用额外空间，改变原来数组
func merge(A []int, m int, B []int, n int) {
	var a = m - 1
	var b = n - 1
	var i int
	for i = m + n - 1; a >= 0 && b >= 0; i-- {
		if A[a] >= B[b] {
			A[i] = A[a]
			a--
		} else {
			A[i] = B[b]
			b--
		}
	}
	if a < 0 {
		for ; i >= 0; i-- {
			A[i] = B[b]
			b--
		}
	}
	fmt.Println(A)
}

// 最小路径和
func minPathSum(matrix [][]int) int {
	n := len(matrix)
	m := len(matrix[0])

	dp := make([][]int, n)
	for k := range matrix {
		if dp[k] == nil {
			dp[k] = make([]int, m)
			dp[0][0] = matrix[0][0]
		}
		if k < 1 {
			continue
		}
		dp[k][0] = matrix[k][0] + dp[k-1][0]
	}

	for k := range matrix[0] {
		if k > 0 {
			dp[0][k] = matrix[0][k] + dp[0][k-1]
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + matrix[i][j]
		}
	}
	fmt.Println(matrix)
	fmt.Println(dp)
	return dp[n-1][m-1]
}

// 最小路径和 用原来数组不需要创建
func minPathSumV2(matrix [][]int) int {
	n := len(matrix)
	m := len(matrix[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if j == 0 {
				if i > 0 {
					matrix[i][j] += matrix[i-1][j]
				}
				continue
			}
			if i == 0 {
				matrix[i][j] += matrix[i][j-1]
				continue
			}
			matrix[i][j] = min(matrix[i-1][j], matrix[i][j-1]) + matrix[i][j]
		}
	}
	fmt.Println(matrix)
	return matrix[n-1][m-1]
}

func TestLeetcode(t *testing.T) {

	// k := getLongestPalindrome("ab1234321abcvbnmmnbvcba1", 24)
	// k := minPathSum([][]int{[]int{1, 3, 5, 9}, []int{8, 1, 3, 4}, []int{5, 0, 6, 1}, []int{8, 8, 4, 0}})
	k := minPathSumV2([][]int{{1, 3, 5, 9}, {8, 1, 3, 4}, {5, 0, 6, 1}, {8, 8, 4, 0}})
	// k := minPathSum([][]int{[]int{1, 1, 5, 9}, []int{8, 1, 3, 4}, []int{5, 0, 6, 1}, []int{8, 1, 1, 0}})
	// k := minPathSumV2([][]int{[]int{1, 1, 5, 9}, []int{8, 1, 3, 4}, []int{5, 0, 6, 1}, []int{8, 1, 1, 0}})

	println(k)
}
func getLongestPalindrome(A string, n int) int {

	k := 0
	for i := 0; i < n; i++ {
		// 两种情况： 一种是 aba  一种是：aa 所以用 2 个 for 循环
		for j := 0; i-j >= 0 && j+i < n; j++ {
			if A[i-j] != A[i+j] {
				break
			}
			k = max(k, 2*j+1)

		}
		for j := 0; i-j >= 0 && j+i+1 < n; j++ {
			if A[i-j] != A[i+j+1] {
				break
			}
			k = max(k, 2*(j+1))
		}
	}
	return max(k, 0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 小练归并排序
func TestMergeArr(t *testing.T) {
	arr1 := []int{1, 3, 5, 7, 9}
	arr2 := []int{2, 4, 6, 8, 10}
	i, j := 0, 0
	for j < len(arr2) && i < len(arr1) {
		if arr1[i] <= arr2[j] {
			i++
		} else {
			arr1 = append(arr1[:i], append([]int{arr2[j]}, arr1[i:]...)...)
			i++
			j++
		}
	}
	if j < len(arr2) {
		arr1 = append(arr1, arr2[j])
		j++
	}

	fmt.Println(arr1)
}

func TestChanV2(t *testing.T) {

	ob := &S{
		cl:     make(chan struct{}),
		notity: make(chan int, 1),
	}
	ob.wg.Add(1)

	// 写线程
	go func() {
		ob.wg.Add(1)

		i := 0
		for {
			select {
			case <-ob.cl:
				ob.wg.Done()
				return
			case ob.notity <- i:

			}
			ob.num++
			i++
			time.Sleep(time.Second)
		}
	}()

	// 写线程
	go func() {
		ob.wg.Add(1)

		i := 500

		for {
			select {
			case <-ob.cl:
				ob.wg.Done()
				return
			case ob.notity <- i:

			}
			ob.num--
			i--
			time.Sleep(time.Second)
		}
	}()

	// 写线程
	go func() {
		ob.wg.Add(1)

		i := 200

		for {
			select {
			case <-ob.cl:
				ob.wg.Done()
				fmt.Println("推出 chan")
				return
			case ob.notity <- i:

			}
			ob.num--
			i--
			time.Sleep(time.Second)
		}
	}()

	// 读线程
	go func() {
		for v := range ob.notity {
			fmt.Println("读取", v, ob.num)
		}
		// for {
		// 	select {
		// 	case v, ok := <-ob.notity:
		// 		if ok {
		// 			fmt.Println(v, ob.num)
		// 		} else {
		// 			fmt.Println(999)
		// 			return
		// 		}
		// 	case <-ob.cl:
		// 		return
		// 	}
		// }
	}()

	go func() {
		time.Sleep(time.Second * 10)
		ob.wg.Done()
		close(ob.cl)
	}()

	println(1222)
	ob.wg.Wait()
	println(333)

	time.Sleep(time.Hour)
	return

}

func TestQuickSoft(t *testing.T) {
	arr := []int{4, 3, 5, 1, 2, 6, 33, 12, 1, 55, 3, 2, 111, 57, 7, 5}
	fmt.Println(arr)
	QuickSoft(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

// 堆排序小练
func TestHeapSort(t *testing.T) {
	// arrStruct := []int{4, 3, 5, 1, 2, 6, 7}
	arr := []int{1, 4, 3, 2, 6, 5, 8, 7, 9, 0}

	fmt.Println(arr)

	BuildHeap(arr, len(arr))
	fmt.Println(arr)
}

// 随便练一下 二叉树排序 =》堆排序
func Test2TreeSoft(t *testing.T) {
	// arrStruct := []int{1, 2, 3, 4, 5, 6, 7, 8}
	arr := []int{4, 3, 5, 1, 2, 6, 7}

	ts := []int{0, 0, 0, 1}
	ss := copy(ts, arr)
	fmt.Println(len(ts), cap(ts), arr, ts, ss)
	fmt.Printf("%v,%p,%p,", ss, ts, arr)

	node := &structures.TreeNode{Val: 4}
	for v := range arr[:6] {
		node = CreateTree(node, arr[v+1])
	}
	fmt.Println("begin")
	Travel(node)
}

func TestBlocking(t *testing.T) {
	ch := make(chan struct{})

	// var x interface{} = nil
	// var y *int = nil
	// interfaceIsNil(x)
	// interfaceIsNil(y)

	aa1 := "aaa" + "222你好"
	var aa2 strings.Builder
	aa2.WriteString(aa1)
	aa2.WriteString("24444")
	fmt.Println(aa2.String())
	go func() {
		time.Sleep(time.Hour)
		ch <- struct{}{}
	}()
	<-ch

}

// 无缓冲 buf chan
func TestChanNoBuf(t *testing.T) {

	ch := make(chan int)
	timeout := make(chan struct{})
	go func() {
		i := 0
		for {
			i++
			select {
			case <-timeout:
				return
			default:
				ch <- i
			}
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		i := 1
		for {
			aa, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println(aa, ok, i)
			i++
			if i == 3 {
				timeout <- struct{}{}
				close(ch)
				return
			}

		}
	}()

	ww := sync.WaitGroup{}
	// ww.Add(10)
	go func() {
		for {
			time.Sleep(10 * time.Second)
			ww.Done()
		}
	}()

	ww.Wait()

}

func TestBfs(t *testing.T) {
	// 初始化树
	tree := structures.Ints2TreeNode([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	// BfsTree(tree)
	DfsTreeV2(tree)
	// DfsTree(tree)
	BfsTree(tree)
}

// 深度遍历
func DfsTree(tree *structures.TreeNode) {
	if tree == nil {
		return
	}
	fmt.Println(tree.Val)
	if tree.Left != nil {
		DfsTree(tree.Left)
	}
	if tree.Right != nil {
		DfsTree(tree.Right)
	}
}

// 深度遍历 压栈处理
func DfsTreeV2(tree *structures.TreeNode) {

}

// 广度遍历 队列实现
func BfsTree(tree *structures.TreeNode) {
	if tree == nil {
		return
	}
	var node []*structures.TreeNode
	node = []*structures.TreeNode{tree}
	for len(node) != 0 {

		t := node[0]
		fmt.Println(t.Val, len(node))
		node = node[1:]

		l := t.Left
		if l != nil {
			node = append(node, l)
		}
		r := t.Right
		if r != nil {
			node = append(node, r)
		}

	}

}

func TestWeiyi(t *testing.T) {
	// 000011
	aa := 3
	// 0000100
	bb := 4
	t.Log(aa >> 1)
	t.Log(aa << 10)
	// 取相反数
	t.Log(^99)
	// 异位或
	t.Log(15 | 20)
	t.Log(99 | 91)
	// 判断奇偶
	// 异位与
	t.Log(bb & 1)
	//
	t.Log(aa ^ bb)
	t.Log(aa | bb)

}

func TestArrEq(t *testing.T) {
	aa := []byte{1, 2, 3}
	bb := []byte{1, 2, 3}
	cc := []byte{1, 3, 2}
	dd := []int{1, 3, 2}

	println(bytes.Equal(aa, bb))
	println(reflect.DeepEqual(aa, cc))
	println(reflect.DeepEqual(dd, cc))
	println(reflect.DeepEqual(aa, bb))
}

func TestSliceRange(t *testing.T) {
	t.Helper()
	aa := []*PayWay{}

	aa = append(aa, &PayWay{
		Id:  123,
		Ids: 123,
	})
	aa = append(aa, &PayWay{
		Id:  222,
		Ids: 222,
	})
	aa = append(aa, &PayWay{
		Id:  333,
		Ids: 333,
	})

	for k, v := range aa {
		fmt.Println(v, k)
	}

}

func TestPanicdefer(t *testing.T) {
	t.Log(DeferTest())
}

func DeferTest() int {
	a := 1
	b := 2
	defer calc(a, b, "1")
	// defer calc(a, calc(a, b, "0"), "1")
	// a = 0
	// defer calc(a, calc(a, b, "3"), "2")
	return a + 1
}

func calc(x, y int, s string) int {
	fmt.Println(s)
	fmt.Println(x, y, x+y)
	return x + y
}

func TestZhengzebiaoda(t *testing.T) {
	text := "fff${LastDateOfMonth(3)}ffff aa2021年02月30日aaa${LastDateOfMonth(123)}aaa     "
	mach := "\\$\\{LastDateOfMonth.([0-9]+.)\\}"
	re, _ := regexp.Compile(mach)

	// 取出所有符合规则日期
	list := re.FindAllString(text, -1)
	re1, _ := regexp.Compile("[0-9]+")
	t.Log("替换前：", text, "\n")

	// 遍历替换不同日期
	for _, v := range list {
		dayString := re1.Find([]byte(v))
		days, _ := strconv.Atoi(string(dayString))
		// 获取目标日期
		targetDate := LastDateOfMonth(days, time.Now())
		// 整合当前替换规则
		curDate := "\\$\\{LastDateOfMonth.(" + string(dayString) + ".)\\}"
		// 生成当前替换规则
		re1, _ := regexp.Compile(curDate)
		// 执行替换
		text = re1.ReplaceAllString(text, targetDate)
	}
	t.Log("替换后：", text, "\n")
	ts := time.Now()
	tm1 := time.Date(ts.Year(), ts.Month(), ts.Day()+1, 0, 0, 0, 0, ts.Location())
	tm2 := tm1.AddDate(0, 0, 1)
	t.Log(tm1, tm2)

}

// param: days 为多少天以后
// return: 今天+days 天之后的日期,所在月的最后一天, 按"2006年01月02日"格式化
func LastDateOfMonth(days int, ct time.Time) string {
	d := ct.AddDate(0, 0, days)              // time.Now()可以换成支持测试环境调时间的方法
	firstDate := d.AddDate(0, 0, -d.Day()+1) // 当月的第一天
	lastDate := firstDate.AddDate(0, 2, -1)
	// lastDate.Unix()
	// 当月的最后一天
	return lastDate.Format("2006年01月02日")
}

type PayWay struct {
	//    支付id
	Id  int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Ids int64 `protobuf:"varint,2,opt,name=id,proto3" json:"ids,omitempty"`
	// 支付名称
	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Test  B
	Test2 *A
}

type B struct {
	S A
}

type S struct {
	cl     chan struct{}
	num    int
	notity chan int
	wg     sync.WaitGroup
	sync.Mutex
}
