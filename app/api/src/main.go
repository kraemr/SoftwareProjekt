package main;
import (
    "fmt"
    "net/http"
    "os"
)





func main() {
	argsWithProg := os.Args
	fmt.Println(argsWithProg);

    publicDir := "/opt/app/public"

    fileServer := http.FileServer(http.Dir(publicDir))
    http.Handle("/", fileServer)
    fmt.Println("Server is running on port 8000")
    err := http.ListenAndServe(":8000", nil)
    if err != nil {
        fmt.Println("Error starting File server:", err)
    }



}
