package saruman

import "github.com/vectorman1/saruman/serve"

func main(){
	err := serve.HandleRequests()

	if err != nil {
		panic(err)
	}
}