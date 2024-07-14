// Package helloworld provides a set of Cloud Functions samples.
package helloworld

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"cloud.google.com/go/firestore"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/google/uuid"
	"google.golang.org/api/iterator"
	"gopkg.in/alessio/shellescape.v1"
)

func convertContent(content string) string {
	// shell injection対策
	cmd := fmt.Sprintf("echo %s | perl validation_content.pl", shellescape.Quote(content))
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

func generateUUID() string {
	return uuid.New().String()
}

type Todo struct {
	ID      string `json:"id"`
	Done    bool   `json:"done"`
	Content string `json:"content"`
}

type TodoData struct {
	Done    bool   `json:"done"`
	Content string `json:"content"`
}

func init() {
	ctx := context.Background()
	client := createClient(ctx)
	todoapp := createTodoApp(client, ctx)
	functions.HTTP("todoapp", todoapp)
}

func createClient(ctx context.Context) *firestore.Client {
	// Sets your Google Cloud Platform project ID.
	projectID := os.Getenv("PROJECT")

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	// Close client when done with
	// defer client.Close()
	return client
}

func createTodoApp(c *firestore.Client, x context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getTodos(c, x, w, r)
		case "POST":
			postTodos(c, x, w, r)
		case "PUT":
			putTodos(c, x, w, r)
		case "DELETE":
			deleteTodos(c, x, w, r)
		}
	}
}

// DELETE /?id=123
func deleteTodos(c *firestore.Client, x context.Context, w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	_, err := c.Collection("todo").Doc(id).Delete(x)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprint(w, "")
}

func putTodos(c *firestore.Client, x context.Context, w http.ResponseWriter, r *http.Request) {
	var raw Todo
	err := json.NewDecoder(r.Body).Decode(&raw)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = c.Collection("todo").Doc(raw.ID).Update(x, []firestore.Update{{
		Path:  "done",
		Value: raw.Done,
	}, {
		Path:  "content",
		Value: convertContent(raw.Content),
	},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprint(w, "")
}

func postTodos(c *firestore.Client, x context.Context, w http.ResponseWriter, r *http.Request) {
	var raw TodoData
	err := json.NewDecoder(r.Body).Decode(&raw)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := generateUUID()
	_, err = c.Collection("todo").Doc(id).Set(x, map[string]interface{}{
		"id":      id,
		"done":    raw.Done,
		"content": convertContent(raw.Content),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprint(w, "")
}

func getTodos(c *firestore.Client, x context.Context, w http.ResponseWriter, _ *http.Request) {
	iter := c.Collection("todo").Documents(x)
	var ret []map[string]interface{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		ret = append(ret, doc.Data())
	}
	json.NewEncoder(w).Encode(ret)
}
