package kafka

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func mainProducer() {
	var err error
	reader := bufio.NewReader(os.Stdin)
	kafka := newKafkaSyncProducer()

	for {
		fmt.Print("->")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		args := strings.Split(text, "-")
		cmd := args[0]

		switch cmd {
		case "create":
			if len(args) == 2 {
				acc
			}
		}
	}
}
