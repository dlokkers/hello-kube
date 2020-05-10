package main

import (
	"log"

	pb "github.com/dlokkers/hello-kube/protobuf"
	"github.com/dlokkers/hello-kube/routes"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server
	// TODO: Create these values as configurables
	conn, err := grpc.Dial("localhost:5678", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewRecipesClient(conn)

	// Setup HTTP Server
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	routes.Init(r, client)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
