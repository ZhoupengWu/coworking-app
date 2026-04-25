package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.GET("/credit-cards", GetCreditCards)

	if err := r.Run(":8089"); err != nil {
		panic(err)
	}
}

type creditCard struct {
	Number       string     `json:"number"`
	SecurityCode string     `json:"-"`
	Expiration   *time.Time `json:"expiration,omitempty"`
}

func GetCreditCards(ctx *gin.Context) {
	expirationDate := time.Now().AddDate(1, 0, 0)
	cards := []creditCard{
		{
			Number:       "1234",
			SecurityCode: "999",
			Expiration:   &expirationDate,
		},
		{
			Number:       "1567",
			SecurityCode: "259",
		},
		{
			Number:       "1964",
			SecurityCode: "146",
		},
		{
			Number:       "9610",
			SecurityCode: "670",
		},
	}

	ctx.IndentedJSON(http.StatusOK, cards)
}

/* func main () {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.GET("/safe", SafeHandler)
	r.GET("/unsafe", UnsafeHandler)

	if err := r.Run(":8089"); err != nil {
		panic(err)
	}
}

func SafeHandler (ctx *gin.Context) {
	keyToGet := "foo"
	ctx.Set("foo", "baz")
	val, exist := ctx.Get(keyToGet)

	if !exist {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("The key %s doesn't exist\n", keyToGet))

		return
	}

	ctx.String(http.StatusOK, val.(string))
}

func UnsafeHandler (ctx *gin.Context) {
	keyToGet := "foo2"
	ctx.Set("foo", "baz")
	val := ctx.MustGet(keyToGet)
	ctx.String(http.StatusOK, val.(string))
} */

/* func main () {
	var a1 interface{}
	a1 = dog{}

	d1 := a1.(dog)
	fmt.Printf("d1 makeSound: %s\n", d1.makeSound())

	// It panics because the "ball" is not an "animal"
	// a1 = ball{}
	// b1 := a1.(animal)
	// fmt.Printf("ball makeSound: %s", b1.makeSound())

	a1 = ball{}
	b1, ok := a1.(animal)

	if !ok {
		fmt.Println(`"ball" is not an "animal"`)
	}

	fmt.Printf("b1: %v\n", b1)

	a1 = ball{}
	b2 := a1.(ball)
	fmt.Printf("b2: %s\n", b2.bounce())
}

type animal interface {
	makeSound() string
}

type dog struct {

}

func (d dog) makeSound () string {
	return "Woof!"
}

type ball struct {

}

func (b ball) bounce () string {
	return "Boom!"
} */

/* func main () {
	c1 := circle{
		radius: 5,
	}
	r1 := rect{
		width: 4.5,
		height: 2.5,
	}
	measure(c1)
	measure(r1)
	fmt.Printf("Il perimetro è %f\n", c1.perimetry())
}

func measure (g geometry) {
	fmt.Printf("L'area è %f\n", g.area())
}

type geometry interface {
	area() float64
}

type circle struct {
	radius float64
}

func (c circle) area () float64 {
	return c.radius * c.radius * math.Pi
}

func (c circle) perimetry () float64 {
	return c.radius * 2 * math.Pi
}

type rect struct {
	width, height float64
}

func (r rect) area () float64 {
	return r.width * r.height
} */

/* func main () {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.GET("/example", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, fmt.Sprintln("This is a demo endpoint"))
	})
	r.GET("/health", HandlerHealthCheck)

	if err := r.Run(":8089"); err != nil {
		panic(err)
	}
}

func HandlerHealthCheck (ctx *gin.Context) {
	ctx.String(http.StatusOK, fmt.Sprintln("The system is health"))
} */

/*
func main () {
	router := http.NewServeMux()
	router.HandleFunc("/example", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/example endpoint has been invoked")
		w.Write([]byte(fmt.Sprintf("This is a demo endpoint...")))
	})
	router.HandleFunc("/health", HandlerHealthCheck)

	err := http.ListenAndServe(":8089", router)

	if err != nil {
		panic(err)
	}
}

func HandlerHealthCheck (w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandlerHealthCheck has been invoked")

	if r.Method != http.MethodGet {
		w.Write([]byte(fmt.Sprintf("The request must be a GET")))

		return
	}

	w.Write([]byte(fmt.Sprintf("The system is health")))
} */
