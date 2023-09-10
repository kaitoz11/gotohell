package scanner

import (
	"fmt"
	"net"
)

func worker(target string, ports chan int, result chan int) {
	for port := range ports {
		address := fmt.Sprintf("%s:%d", target, port)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Printf("%s\n", err)
			result <- -1
			continue
		}
		conn.Close()
		result <- port
	}
}

// func ScanWG(target string, ports []int) []int {
// 	var wg sync.WaitGroup
// 	var result []int
// 	for _, p := range ports {
// 		wg.Add(1)
// 		go func(port int) {
// 			defer wg.Done()
// 			address := fmt.Sprintf("%s:%d", target, port)
// 			conn, err := net.Dial("tcp", address)
// 			if err != nil {
// 				fmt.Printf("%s\n", err)
// 				return
// 			}
// 			conn.Close()
// 			// fmt.Printf("Reached %s\n", address)
// 			result = append(result, port)
// 		}(p)
// 	}
// 	wg.Wait()
// 	return result
// }

func Do(target string, ports []int, workers int) []int {
	p := make(chan int, workers)
	result := make(chan int)
	// var wg sync.WaitGroup
	var openPorts []int

	for i := 0; i < workers; i++ {
		go worker(target, p, result)
	}

	go func() {
		for _, port := range ports {
			p <- port
		}
	}()

	for i := 0; i < len(ports); i++ {
		scanResult := <-result
		if scanResult != -1 {
			openPorts = append(openPorts, scanResult)
		}
	}

	close(p)
	close(result)

	return openPorts
}
