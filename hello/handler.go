package function

import (
	"fmt"

	"github.com/hypnoglow/somefunc/hello/sub"
)

// Handle a serverless request
func Handle(req []byte) string {
	return fmt.Sprintf("Hello, %s!", sub.Name())
}
