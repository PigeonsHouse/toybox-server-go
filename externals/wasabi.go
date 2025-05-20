package externals

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
)

func Wasabi() {
	op := session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}
	fmt.Println(op)
}
