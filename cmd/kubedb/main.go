package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/dlokkers/hello-kube/model"
	pb "github.com/dlokkers/hello-kube/protobuf"
)

type server struct {
	db model.RecipeStore
}

func (s *server) ListRecipes(cxt context.Context, r *pb.Empty) (*pb.RecipeList, error) {
	result := &pb.RecipeList{}
	recipes, err := s.db.GetAllRecipes()

	if err != nil {
		return nil, err
	}

	for _, r := range recipes {
		result.Title = append(result.Title, r)
	}

	return result, nil
}

func (s *server) RetrieveRecipe(cxt context.Context, r *pb.RecipeTitle) (*pb.Recipe, error) {
	recipe, err := s.db.GetRecipe(r.Title)
	if err != nil {
		return nil, err
	}

	result := &pb.Recipe{
		Title:   recipe.Title,
		Process: recipe.Process,
	}

	for _, i := range recipe.Ingredients {
		ingredient := &pb.Ingredient{
			Amount: i.Amount,
			Type:   i.Type,
		}
		result.Ingredients = append(result.Ingredients, ingredient)
	}

	return result, nil
}

func (s *server) AddRecipe(cxt context.Context, r *pb.Recipe) (*pb.Empty, error) {

	add := &model.Recipe{
		Title:   r.Title,
		Process: r.Process,
	}

	for _, i := range r.Ingredients {
		ingredient := model.Ingredient{
			Amount: i.Amount,
			Type:   i.Type,
		}
		add.Ingredients = append(add.Ingredients, ingredient)
	}

	err := s.db.AddRecipe(add.Title, add.Ingredients, add.Process)
	res := &pb.Empty{}

	return res, err
}

func main() {
	// Setup Aerospike cluster
	// TODO: Create these values as configurables
	db, err := model.NewDB("172.28.128.3", 3000)
	if err != nil {
		log.Fatalf("Failed to connect to Aerospike: %v", err)
	}
	defer db.Close()
	server := &server{db}

	// Listen to incoming gRPC calls
	lis, err := net.Listen("tcp", ":5678")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterRecipesServer(s, server)
	reflection.Register(s)

	log.Println("Attempt to listen")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
