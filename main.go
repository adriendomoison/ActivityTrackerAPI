import (
	"fmt"
	"repository"
	"api"
)

func main() {
	fmt.Println("Hello World!")
	repository.InitDBConnection()
	api.InitApi()
}
