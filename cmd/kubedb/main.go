package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/dlokkers/hello-kube/model"
	"github.com/dlokkers/hello-kube/protobuf"
)

type server struct{}

func (s *server) Retrieve(cxt context.Context, r *protobuf.GetArticle) (*protobuf.ArticleResponse, error) {
	result := &protobuf.ArticleResponse{}
	article, err := model.GetArticle(r.Id)

	if err != nil {
		return nil, err
	}

	result.Id = r.Id
	result.Title = article.Title
	result.Content = article.Content

	return result, nil
}

func main() {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	protobuf.RegisterRetrieveArticleServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
