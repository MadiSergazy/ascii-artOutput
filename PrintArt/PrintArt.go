package PrintArt

import (
	"fmt"
	"strings"
)

func PrintArt(massivString []string, args string) string {
	text := ""

	args = strings.ReplaceAll(args, "\n", "\\n")

	splitArgs := strings.Split(args, "\\n") // sozder mojno vmesto args

	for _, arg := range splitArgs {
		// fmt.Println(arg)

		for i := 0; i < 8; i++ {

			for index, v := range arg {
				if !(v <= 126) {
					fmt.Println("DONT CORRECT ARGS IN %d", index)
				}

				if v == 10 {
					continue
				}

				text += strings.Split(massivString[v-32], "\n")[i]
			}

			text += "\n"
		}

	}
	return text
}
