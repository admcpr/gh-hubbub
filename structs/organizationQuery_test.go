package structs

import (
	"fmt"
	"testing"
)

func Test_GetRepository(t *testing.T) {
	t.Run("GetRepository", func(t *testing.T) {
		node := RepositoryQuery{}
		// repository := node.GetRepository()
		// fmt.Println(repository)
		fmt.Print(node)
	})
}
