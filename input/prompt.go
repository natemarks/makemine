package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FromUser() MakeMineInput {
	data := MakeMineInput{}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Full Name (ex. Firstly Lastly):")
	data.Name, _ = reader.ReadString('\n')
	data.Name = strings.TrimSuffix(data.Name, "\n")

	fmt.Println("local computer user account(ex. flastly): ")
	data.LocalUser, _ = reader.ReadString('\n')
	data.LocalUser = strings.TrimSuffix(data.LocalUser, "\n")

	fmt.Println("Email address (ex. flastly@somedomain.com): ")
	data.Email, _ = reader.ReadString('\n')
	data.Email = strings.TrimSuffix(data.Email, "\n")

	return data
}
