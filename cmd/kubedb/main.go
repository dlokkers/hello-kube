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

type server struct{}

func (s *server) ListArticles(cxt context.Context, r *pb.Empty) (*pb.ArticleList, error) {
	result := &pb.ArticleList{}
	articles := model.GetAllArticles()

	for _, v := range articles {
		article := &pb.Article{}
		article.Id = v.ID
		article.Title = v.Title
		article.Content = v.Content

		result.Articles = append(result.Articles, article)
	}

	return result, nil
}

func (s *server) RetrieveArticle(cxt context.Context, r *pb.ArticleID) (*pb.Article, error) {
	result := &pb.Article{}
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
	pb.RegisterArticlesServer(s, &server{})
	reflection.Register(s)
	log.Println("Attempt to listen")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
