package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// Mutex
func increment(mu *sync.Mutex, counter *int) {
	mu.Lock()
	*counter++
	mu.Unlock()
}

func MutexExample() {
	var mu sync.Mutex
	var counter int
	for i := 0; i < 5; i++ {
		go increment(&mu, &counter)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Final counter:", counter)
}

// RWMutex
func readData(rwmu *sync.RWMutex, data *map[string]string, key string) {
	rwmu.RLock()
	defer rwmu.RUnlock()
	fmt.Println("Reading:", (*data)[key])
}

func writeData(rwmu *sync.RWMutex, data *map[string]string, key, value string) {
	rwmu.Lock()
	defer rwmu.Unlock()
	(*data)[key] = value
	fmt.Println("Writing:", key, value)
	time.Sleep(500 * time.Millisecond)
}

func RWLockExample() {
	var rwmu sync.RWMutex
	var data = make(map[string]string)
	go writeData(&rwmu, &data, "key1", "value1")
	time.Sleep(time.Millisecond)
	go readData(&rwmu, &data, "key1")
	time.Sleep(1 * time.Second)
}

// wait group
func worker(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Println("Worker", id, "started")
	time.Sleep(1 * time.Second)
	fmt.Println("Worker", id, "done")
}

func WaitGroupExample() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(&wg, i)
	}
	wg.Wait()
	fmt.Println("All workers finished")
}

// once
func initialize() {
	fmt.Println("Initialize only once")
}

func OnceExample() {
	var once sync.Once
	for i := 0; i < 3; i++ {
		go once.Do(initialize)
	}
	time.Sleep(1 * time.Second)
}

// cond
func workerCond(cond *sync.Cond, ready *bool) {
	cond.L.Lock()
	for !*ready {
		cond.Wait()
	}
	fmt.Println("Worker running")
	cond.L.Unlock()
}

func CondExample() {
	var mu sync.Mutex
	var cond = sync.NewCond(&mu)
	var ready = false

	go workerCond(cond, &ready)
	time.Sleep(1 * time.Second)

	cond.L.Lock()
	ready = true
	cond.Signal()
	cond.L.Unlock()

	time.Sleep(1 * time.Second)
}

func SyncMapExample() {
	var m sync.Map

	m.Store("key1", "value1")
	m.Store("key2", "value2")

	value, ok := m.Load("key1")
	if ok {
		fmt.Println("Loaded value:", value)
	} else {
		fmt.Println("Key not found")
	}

	actual, loaded := m.LoadOrStore("key3", "value3")
	if loaded {
		fmt.Println("Key already existed with value:", actual)
	} else {
		fmt.Println("New value stored:", actual)
	}

	m.Range(func(key, value interface{}) bool {
		fmt.Println("Key:", key, "Value:", value)
		return true
	})

	m.Delete("key2")

	m.Range(func(key, value interface{}) bool {
		fmt.Println("After delete - Key:", key, "Value:", value)
		return true
	})
}

func ObjectBufferPool() {
	var bytePool = sync.Pool{
		New: func() interface{} {
			return make([]byte, 1024)
		},
	}

	data := bytePool.Get().([]byte)
	fmt.Println("Got byte slice of length:", len(data))

	bytePool.Put(data)

	data1 := bytePool.Get().([]byte)
	fmt.Println("Reused byte slice of length:", len(data1))
	data2 := bytePool.Get().([]byte)
	fmt.Println("Reused byte slice of length:", len(data2))
	bytePool.Put(data1)
	bytePool.Put(data2)
}

func goroutineWorker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d processing job %d\n", id, job)
	}
}

func GoroutinePool() {
	var wg sync.WaitGroup
	jobs := make(chan int, 100)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go goroutineWorker(i, jobs, &wg)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait()
}

func producer(ch chan<- int, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Printf("Producer %d produced: %d\n", id, i)
		ch <- i
		time.Sleep(time.Millisecond * 500)
	}
}

func consumer(ch <-chan int, id int) {
	for i := range ch {
		fmt.Printf("Consumer %d consumed: %d\n", id, i)
		time.Sleep(time.Millisecond * 1000)
	}
}

func ChannelExample() {
	ch := make(chan int, 3)
	var wg sync.WaitGroup

	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go producer(ch, &wg, i)
	}

	for i := 1; i <= 2; i++ {
		go consumer(ch, i)
	}
	wg.Wait()
	close(ch)
	time.Sleep(time.Second * 5)
}
