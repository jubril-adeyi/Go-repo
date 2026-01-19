package cursing

import "fmt"

func Curse(name string) string {

	message := fmt.Sprintf("fuck you %v.", name)

	return message
}
