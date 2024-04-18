package main

import (
  "fmt"
  "log"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) 
  if r.Method == "GET" 
    w.Header().Set("Content-Type", "text/html")
    fmt.Fprintf(w, "<form method='POST'>"+
      "<label for='name'>What's your name?</label>"+
      "<input type='text' id='name' name='name'>"+
      "<button type='submit'>Submit</button>"+
      "</form>")
  } else if r.Method == "POST" {
    r.ParseForm()
    name := r.Form.Get("name")
    fmt.Fprintf(w, "Hello, %s! Nice to meet you!\n", name)
  }
}

func main() {
  http.HandleFunc("/", handler)
  fmt.Println("Running demo app. Press Ctrl+C to exit...")
  log.Fatal(http.ListenAndServe(":8888", nil))
}

