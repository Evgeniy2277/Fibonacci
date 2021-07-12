package web

import (
	"fmt"
	"github.com/Evgeniy2277/fibonacci/cache"
	"github.com/Evgeniy2277/fibonacci/closure"
	"github.com/Evgeniy2277/fibonacci/recursion"
	"math/rand"
	"net/http"
	"strconv"
)

func Serve(fiber Fiber) {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		nSlice, ok := request.URL.Query()["n"] //http://localhost:8080/?n=10&n=20
		FibImpl := []Fiber {recursion.NewRecursion(), cache.NewCache(), closure.NewClosure()}
		var wfiber Fiber
		wfiber = FibImpl[rand.Intn(len(FibImpl))]
		data := fmt.Sprintf("%s\n", wfiber.Name())
		if ok {
			for _, nValue := range nSlice {
				n, err := strconv.Atoi(nValue)
				if err == nil {
					if n < 92 {
						data += fmt.Sprintf("Fib(%d) = %d\n", n, wfiber.Fib(n))
					} else {
						data +=fmt.Sprintf("Число  %d слишком большое ,больше 91\n", n)
					}
				}
			}
		}

		_, err := writer.Write([]byte(data))
		if err != nil {
			return
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}