package test

import (
	"context"
	"fmt"
	graphql2 "github.com/machinebox/graphql"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)


func TestResource(t *testing.T){

	req := graphql2.NewRequest(`
		query{
		  todos{
			id
			text
		  }
		}
	`)

	// define a Context for the request
	ctx := context.Background()

	// response structure
	var resp interface{}

	client := graphql2.NewClient("http://localhost:8080/query")

	if err := client.Run(ctx, req, &resp); err != nil {
		t.Error(err)
	}

	fmt.Println(resp)
	log.Print("Log: ", resp)

	require.Equal(t, "1", "1")
}
