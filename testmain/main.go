package main

import (
	"fmt"

	"github.com/jugglechat/im-server/commons/tokens"
)

func main() {
	tokenWrap, err := tokens.ParseTokenString("CgZhcHBrZXkaIDAr072n8uOcw5YBeKCcQ+QCw4m6YWhgt99U787/dEJS")
	fmt.Println(tokenWrap, err)

	token, err := tokens.ParseToken(tokenWrap, []byte("abcdefghijklmnop"))

	fmt.Println(token, err)
}
