package main;
//import "src/test";
import (
    "fmt"
    "net/http"
)

func main() {
    // Define the directory to serve static files from
    publicDir := "../../public"
    // Create a file server handler for the specified directory
    fileServer := http.FileServer(http.Dir(publicDir))
    // Register the file server handler to the root route
    http.Handle("/", fileServer)
    // Start the server on port 8000
    fmt.Println("Server is running on port 8000")
    err := http.ListenAndServe(":8000", nil)
    if err != nil {
        fmt.Println("Error starting server:", err)
    }
}
