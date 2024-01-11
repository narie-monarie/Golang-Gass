# Golang Concurrency

- Parallelism - The task execute simultaneously and independent. Requires multiple CPUs
- Concurrency - seem to happen simultaneously but they are being swapped between very quickly

![[golang runtime]](./images/Pasted%20image%2020240111145020.png)

- Go routines are independent executing functions which run on a set of threads sometimes referred to as light weight threads
	- Run on top of threads
    - go routines are many in magnitude compared to threads
	- Go runtime chooses when the go routines are run
    
    ![[easy Explanation]](./images/goImages2.jpg)
    
    - The cars are the go routines[independent running tasks]
    - The road is the threads they are running on
    - Go runtime as the traffic controller directing traffic
    - Go routines like threads share the same address space

    ```go
    package main
    
    import "fmt"

    func main() {
	    hello()
	    goodbye()
    }

    func hello() {
	    fmt.Println("Hello World")
    }

    func goodbye() {
	    fmt.Println("GoodBye World")
    }
    ```

    1. There is allocation of memory for the program
    2. Start the main goroutine of our program
    3. Run the program on a thread
    4. Execute the code
    5. Shut down and clean up once the program completes