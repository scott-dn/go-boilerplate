package book_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/scott-dn/go-boilerplate/internal/database/entities"
	"github.com/scott-dn/go-boilerplate/internal/pkg/database/query"
	"github.com/scott-dn/go-boilerplate/internal/response"
)

func TestGetAllBookWithInvalidToken(t *testing.T) {
	t.Parallel()

	tokens := []string{"", "Beare", "Bearer", "Bearer 123"}

	for _, token := range tokens {
		req, err := http.NewRequestWithContext(
			context.Background(),
			http.MethodGet,
			"http://localhost:8080/service/api/v1/books",
			nil,
		)
		if err != nil {
			t.Fatalf("Got error: %v", err)
		}
		req.Header.Set("authorization", token)

		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Got error: %v", err)
		}
		defer resp.Body.Close()

		_, err = io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Got error: %v", err)
		}
		if resp.StatusCode != http.StatusUnauthorized {
			t.Fatalf("Got status code: %v But expect: %v", resp.StatusCode, http.StatusUnauthorized)
		}
	}
}

func TestGetAllBookReturnEmptyData(t *testing.T) { //nolint:paralleltest
	if _, err := query.Book.Where(query.Book.ID.IsNotNull()).Delete(); err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		"http://localhost:8080/service/api/v1/books",
		nil,
	)
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}

	//nolint:lll
	req.Header.Set("authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAZXhhbXBsZS5jb20ifQ.3QJBsXLwBNSyLWLDA5nugTzc83x9Ac9zsxKkghKJ__E")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Got status code: %v But expect: %v", resp.StatusCode, http.StatusOK)
	}

	expect := `{"data":[]}`
	actual := strings.Trim(string(respBody), "\n")
	if actual != expect {
		t.Fatalf("Got response: %v\n But expect: %v", actual, expect)
	}
}

func TestGetAllBookWithSomeData(t *testing.T) { //nolint:paralleltest
	if _, err := query.Book.Where(query.Book.ID.IsNotNull()).Delete(); err != nil {
		t.Fatal(err)
	}

	entities := []*entities.Book{
		{Name: "11", Author: "12", Description: "13", CreatedBy: "scott", UpdatedBy: "scott"},
		{Name: "21", Author: "22", Description: "23", CreatedBy: "scott", UpdatedBy: "scott"},
		{Name: "31", Author: "32", Description: "33", CreatedBy: "scott", UpdatedBy: "scott"},
	}

	if err := query.Book.CreateInBatches(entities, 3); err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		"http://localhost:8080/service/api/v1/books",
		nil,
	)
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}

	//nolint:lll
	req.Header.Set("authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAZXhhbXBsZS5jb20ifQ.3QJBsXLwBNSyLWLDA5nugTzc83x9Ac9zsxKkghKJ__E")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Got status code: %v But expect: %v", resp.StatusCode, http.StatusOK)
	}

	var actual response.Books
	json.Unmarshal(respBody, &actual) //nolint:errcheck

	for i, book := range entities {
		if actual.Data[i].Name != book.Name {
			t.Fatalf("Got: %v But expect: %v", actual.Data[i].Name, book.Name)
		}
		if actual.Data[i].Author != book.Author {
			t.Fatalf("Got: %v But expect: %v", actual.Data[i].Author, book.Author)
		}
		if actual.Data[i].Description != book.Description {
			t.Fatalf("Got: %v But expect: %v", actual.Data[i].Description, book.Description)
		}
		if actual.Data[i].Version != book.Version {
			t.Fatalf("Got: %v But expect: %v", actual.Data[i].Version, book.Version)
		}
		if actual.Data[i].CreatedBy != book.CreatedBy {
			t.Fatalf("Got: %v But expect: %v", actual.Data[i].CreatedBy, book.CreatedBy)
		}
		if actual.Data[i].UpdatedBy != book.UpdatedBy {
			t.Fatalf("Got: %v But expect: %v", actual.Data[i].UpdatedBy, book.UpdatedBy)
		}
	}
}
