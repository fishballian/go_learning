package go_learning

type readWriteChan struct {
	setChan chan string
	getChan chan chan<- string
	value   string
}

func NewReadWriteChan() *readWriteChan {
	r := &readWriteChan{
		setChan: make(chan string),
		getChan: make(chan chan<- string),
	}
	r.Init()
	return r
}

func (r *readWriteChan) Init() {
	go func() {
		for {
			r.loop()
		}
	}()
}

func (r *readWriteChan) loop() {
	select {
	case s := <-r.setChan:
		r.value = s
	case c := <-r.getChan:
		c <- r.value
	}
}

func (r *readWriteChan) Set(s string) {
	r.setChan <- s
}

func (r *readWriteChan) Get() string {
	c := make(chan string)
	r.getChan <- c
	select {
	case s := <-c:
		return s
	}
}
