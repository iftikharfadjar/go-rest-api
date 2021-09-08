package main
import "flag"
import "fmt"
import "net/http"
import "log"

const version = "1.0.0"

type config struct {
	port int
	env string
}

type AppStatus struct 

func main(){
	var cfg config
	
	flag.IntVar(&cfg.port, "port", 80, "server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "application environment(development|production)")
	flag.Parse()
	
	fmt.Println("running")
	
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request){
		currentStatus
	}) 
	
	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.port), nil)
	if err != nil {
		fmt.Println("Err")
		log.Println(err)
	}
}