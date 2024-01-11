# Golang Concurrency

- Parallelism - The task execute simultaneously and independent. Requires multiple CPUs
- Concurrency - seem to happen simultaneously but they are being swapped between very quickly

![[golang runtime]](./images/Pasted%20image%2020240111145020.png)

- Go routines are independent executing functions which run on a set of threads sometimes referred to as light weight threads
	- Run on top of threads
    - go routines are many in magnitude compared to threads
	- Go runtime chooses when the go routines are run
    
    ![[easy Explanation]](./images/goImages2.jpg)
    
    - The cars are the go routines [independent running tasks]
    - The road is the threads they are running on
    - Go runtime as the traffic controller directing traffic
    - Go routines like threads share the same address space

- Common golang execution 
    1. There is allocation of memory for the program
    2. Start the main goroutine of our program
    3. Run the program on a thread
    4. Execute the code
    5. Shut down and clean up once the program completes

    ```go
    package main
    
    import "fmt"

    func main() {
	    go hello()
	    goodbye()
    }

    func hello() {
	    fmt.Println("Hello World")
    }

    func goodbye() {
	    fmt.Println("GoodBye World")
    }
    ```
    - What is currently happening in the above code
    ![[Go routine]](./images/goImages3.jpg)
    
    - Hello is inconsistent because the main goroutine

    ```go
    x func main(){ // -> main goroutine has started at time x
    x1    go hello() //-> main goroutine invokes the hello function and another Hello goroutine is created between time x1 and x2
    x2    goodbye() // -> the main goroutine invokes this function and then exits between time x2 and x3  
    x3 }
    ```
    - by the time the hello goroutine wants to print "Hello world", the main goroutine had already exited however if it prints, theres is an out of order messages. In order to change this, we can add a sleep timer to prevent the main goroutine from exiting early
    
    ```go
    x func main(){ // -> main goroutine has started at time x
    x1    go hello() //-> main goroutine invokes the hello function and another Hello goroutine is created between time x1 and x2
          time.Sleep(1 * time.Second) // -> the main goroutine invokes the time function to sleep and the hello goroutine now executes
    x2    goodbye() // -> the main goroutine invokes this function and then exits between time x2 and x3  
    x3 }
    ```
    - but sleep is not good in production

## The sync.WaitGroup
- used to wait for goroutines to finish
- under the hood it uses a simple counter and an inner lock
- the zero value of the wait group is ready to be used 
- var sync.waitGroup

```go
func (wg *WaitGroup) Add(delta int) //->add(number of goroutines we wish to wait for and panics if the inner counter is negative)
func (wg *WaitGroup) Done() //-> decrements innner counter by 1 an should be used when a g routine finishes the work assigned
func (wg *WaitGroup) Wait() //-> blocks the routine in which it is invoked until the counter reaches 0
```


### Race Conditions
- Occur when multiple goroutine read and write shared data without any synchronization techniques

## Channels
![[Channels]](./images/image3.jpg)

- channels is like a tunned to send data between go routines
![[channel Directions]](./images/image4.jpg)

- sends are receives are blockin until the send and receive are succesful

*Unbuffered Channels*
- they are synchronous
![[unbuffered]](./images/image5.jpg)

*Buffered Channels*
- They are asynchronous
![[buffered Channels]](./images/image6.jpg)

*Channel Directions*
![[Chennel Directions]](./images/image7.jpg)