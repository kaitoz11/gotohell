package main

import (
	"fmt"
	"gotohell/scanner"
	"gotohell/utils"
	"strconv"
)

func main() {
	notificationServer := "http://100.77.84.42:3001/"
	target := "firezone.yp2743.me"
	// ports := []int{22, 80, 443, 8080, 3000}

	lo, hi := 1, 1535
	ports := make([]int, hi-lo+1)
	for i := range ports {
		ports[i] = i + lo
	}
	result := scanner.Do(target, ports, 100)
	// result := scanner.ScanWG(target, ports)
	fmt.Println(result)
	ntfy := utils.MakeNotification(notificationServer)

	msg := fmt.Sprintf("Host: %s\nOpen ports: ", target)
	for i, p := range result {
		msg += strconv.Itoa(p)
		if i != len(result)-1 {
			msg += ", "
		}
	}

	ntfy.SetOptions("k4it0z11", "", "Port scanning is Done!", msg, nil, 0)

	fmt.Println(ntfy.SendNotification())
}
