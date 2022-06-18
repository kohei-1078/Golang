package main

import (
	"fmt"
	"sync"
	"time"
	"runtime"
)

func sayHello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello")
}

func memConsumed() uint64 {
	runtime.GC()

	var s runtime.MemStats

	runtime.ReadMemStats(&s)

	return s.Sys
}

func Hello(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Printf("Hello from %v!\n", id)
} 

type value struct {
	mu sync.Mutex
	value int
	name string
}

var onceA, onceB sync.Once

func A() {
	fmt.Println("A")
	onceB.Do(B)
}

func B() {
	fmt.Println("B")
	onceA.Do(A)
}

type Person struct {
	Name string
}

func main() {
	count := 0
	mypool := & sync.Pool{
		New: func() interface{} {
			count++
			fmt.Println("Creating...")
			return struct{}{}
			// fmt.Println("Creating...")
			// return new(Person)
			// fmt.Println("Create new instance")
			// return struct{}{}
		},
	}

	mypool.Put("manualy added: 1")
	mypool.Put("manualy added: 2")

	var wg sync.WaitGroup

	wg.Add(10000)

	for i := 0; i < 10000; i++ {
		time.Sleep(1 * time.Millisecond)
		go func() {
			defer wg.Done()
			instance := mypool.Get()
			mypool.Put(instance)
		}()
	}

	wg.Wait()

	fmt.Printf("created instance: %d", count)



	// mypool.Put(&Person{Name: "1"})
	// mypool.Put(&Person{Name: "2"})

	// instance1 := mypool.Get()
	// instance2 := mypool.Get()
	// instance3 := mypool.Get().(*Person)

	// fmt.Println(instance1, instance2, instance3)

	// instance3.Name = "3"

	// fmt.Println(instance1, instance2, instance3)

	// mypool.Put(instance1)
	// mypool.Put(instance2)
	// mypool.Put(instance3)

	// instnace4 := mypool.Get()
	// fmt.Println(instnace4)

	// instance := mypool.Get()
	// mypool.Put(instance)
	// mypool.Get()

	// onceA.Do(A)

	// count := 0

	// increment := func() {
	// 	count++
	// }

	// decrement := func() {
	// 	count--
	// }

	// var once sync.Once
	
	// once.Do(increment)
	// once.Do(decrement)

	// fmt.Println(count)


	// cond := sync.NewCond(&sync.Mutex{})

	// go func() {
	// 	for range time.Tick(1 * time.Second) {
	// 		cond.Broadcast()
	// 	}
	// }()

	// var flag [2]bool

	// takeStep := func() {
	// 	cond.L.Lock()
	// 	cond.Wait()
	// 	cond.L.Unlock()
	// }

	// var wg sync.WaitGroup

	// p0 := func() {
	// 	defer wg.Done()

	// 	flag[0] = true

	// 	takeStep()

	// 	for flag[1] {
	// 		takeStep()

	// 		flag[0] = false

	// 		takeStep()

	// 		if flag[0] != flag[1] {
	// 			break
	// 		}

	// 		takeStep()

	// 		flag[0] = true

	// 		takeStep()
	// 	}
	// }

	// p1 := func() {
	// 	defer wg.Done()

	// 	flag[1] = true

	// 	takeStep()

	// 	for flag[1] {
	// 		takeStep()

	// 		flag[1] = false

	// 		takeStep()

	// 		if flag[1] != flag[1] {
	// 			break
	// 		}

	// 		takeStep()

	// 		flag[1] = true

	// 		takeStep()
	// 	}
	// }

	// wg.Add(2)

	// go p0()
	// go p1()

	// wg.Wait()

	// var mutex sync.Mutex
	// cond := sync.NewCond(&mutex)

	// for _, name := range []string{"A","B","C"} {
	// 	go func(name string) {
	// 		mutex.Lock()
	// 		defer mutex.Unlock()

	// 		cond.Wait()
	// 		fmt.Println(name)
	// 	}(name)
	// }

	// fmt.Println("Ready...")
	// time.Sleep(time.Second)
	// fmt.Println("Go!")
	// for i := 0; i < 3; i++ {
	// 	time.Sleep(time.Second)
	// 	cond.Signal()
	// }
	// cond.Broadcast()

	// time.Sleep(time.Second)

	// fmt.Println("Done")

	// var count int 
	// var lock sync.RWMutex
	// var wg sync.WaitGroup

	// increment := func(wg *sync.WaitGroup, l sync.Locker) {
	// 	l.Lock()
	// 	defer l.Unlock()
	// 	defer wg.Done()

	// 	fmt.Println("increment")
	// 	count++
	// 	time.Sleep(1 * time.Second)
	// }

	// read := func(wg *sync.WaitGroup, l sync.Locker) {
	// 	l.Lock()
	// 	defer l.Unlock()
	// 	defer wg.Done()

	// 	fmt.Println("read")
	// 	fmt.Println(count)
	// 	time.Sleep(1 * time.Second)
	// }

	// start := time.Now()

	// for i := 0; i < 5; i++ {
	// 	wg.Add(1)

	// 	go increment(&wg, &lock)
	// }


	// for i := 0; i < 5; i++ {
	// 	wg.Add(1)

	// 	go read(&wg, lock.RLocker())
	// }

	// wg.Wait()

	// end := time.Now()

	// fmt.Println(end.Sub(start))

	// var wg sync.WaitGroup
	// var lock sync.Mutex

	// const timer = 1 * time.Millisecond

	// greedyWorker := func() {
	// 	defer wg.Done()

	// 	count := 0

	// 	begin := time.Now()

	// 	for time.Since(begin) <= timer {
	// 		lock.Lock()
	// 		time.Sleep(3 * time.Nanosecond)
	// 		lock.Unlock()
	// 		count++
	// 	}

	// 	fmt.Printf("greedyWorker: %v\n", count)
	// }

	// politeWocker := func() {
	// 	defer wg.Done()

	// 	count := 0

	// 	begin := time.Now()

	// 	for time.Since(begin) <= timer {
	// 		lock.Lock()
	// 		time.Sleep(1 * time.Nanosecond)
	// 		lock.Unlock()
	// 		lock.Lock()
	// 		time.Sleep(1 * time.Nanosecond)
	// 		lock.Unlock()
	// 		lock.Lock()
	// 		time.Sleep(1 * time.Nanosecond)
	// 		lock.Unlock()
	// 		count++
	// 	}

	// 	fmt.Printf("greedyWorker: %v\n", count)
	// }

	// wg.Add(2)

	// go greedyWorker()
	// go politeWocker()

	// wg.Wait()

	// var wg sync.WaitGroup

	// printSum := func(v1, v2 *value) {
	// 	defer wg.Done()
	// 	v1.mu.Lock()
	// 	fmt.Printf("%v がロックを取得しました\n", v1.name)
	// 	defer v1.mu.Unlock()

	// 	time.Sleep(2 * time.Second)

	// 	v2.mu.Lock()
	// 	fmt.Printf("%v がロックを取得しました\n", v2.name)
	// 	defer v2.mu.Unlock()

	// 	fmt.Println(v1.value + v2.value)
	// }

	// var a value = value{name: "a"}
	// var b value = value{name: "b"}

	// wg.Add(2)

	// go printSum(&a, &b)
	// go printSum(&b, &a)

	// wg.Wait()

	// var wg sync.WaitGroup
	// var memoryAccess sync.Mutex
	// var data int

	// wg.Add(1)

	// go func() {
	// 	defer wg.Done()
	// 	memoryAccess.Lock()
	// 	data++
	// 	memoryAccess.Unlock()
	// }()

	// wg.Wait()

	// memoryAccess.Lock()
	// if data == 0 {
	// 	fmt.Println(0)
	// } else {
	// 	fmt.Println(data)
	// }
	// memoryAccess.Unlock()


	// var wg sync.WaitGroup

	// var CPU int = runtime.NumCPU()

	// wg.Add(CPU)
	// for i := 1; i <= CPU; i++ {
	// 	go Hello(&wg, i)
	// }

	// wg.Wait()

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	fmt.Println("1st Groutine Start")
	// 	time.Sleep(1 * time.Second)
	// 	fmt.Println("1st Groutine Done")
	// }()

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	fmt.Println("2nd Groutine Start")
	// 	time.Sleep(1 * time.Second)
	// 	fmt.Println("2nd Groutine Done")
	// }()

	// wg.Wait()


	// var ch <-chan interface{}
	// var wg sync.WaitGroup

	// noop := func() {
	// 	wg.Done()
	// 	<-ch
	// }

	// const numGoroutines = 1000000

	// wg.Add(numGoroutines)

	// before := memConsumed()

	// for i := 0; i < numGoroutines; i++{
	// 	go noop()
	// }

	// wg.Wait()

	// after := memConsumed()

	// fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)

	// var wg sync.WaitGroup

	// tasks := []string{"A", "B", "C"}

	// for _, task := range tasks {
	// 	wg.Add(1)

	// 	go func(task string) {
	// 		defer wg.Done()
	// 		fmt.Println(task)
	// 	}(task)
	// }

	// wg.Wait()

	// say := "Hello"

	// wg.Add(1)

	// go func() {
	// 	defer wg.Done()
	// 	say = "GoodBye"
	// }()

	// wg.Wait()

	// fmt.Println(say)

	// var wg sync.WaitGroup

	// wg.Add(1)

	// go sayHello(&wg)

	// wg.Add(1)

	// go func() {
	// 	defer wg.Done()
	// 	fmt.Println("Hello")
	// }()
	
	// wg.Wait()
}