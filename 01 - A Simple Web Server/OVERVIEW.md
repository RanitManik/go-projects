# 01 - A Simple Web Server

This Go program is a basic HTTP web server that serves static files and handles two routes: `/form` (for processing form submissions) and `/hello` (a simple greeting endpoint).

## 🗂️ Overview

This server does three main things:

1. **Serves static files** from a directory (`./static`) on the root URL (`/`).
2. **Handles GET requests** to `/hello`.
3. **Handles POST requests** to `/form`, extracting form values.

## 📦 Key Components Explained

### 1. **`formHandler`**

```go
func formHandler(w http.ResponseWriter, r *http.Request)
```

- **Path Check**: Ensures the request path is `/form`, else responds with a 404 error.

- **Method Check**: Only accepts `POST` requests; other methods get a 404 response.

- **Form Parsing**:

  ```go
  if err := r.ParseForm(); err != nil {
      fmt.Fprintf(w, "ParseForm() err: %v", err)
      return
  }
  ```

  Parses the form values from the request body.

- **Retrieves Form Values**:

  ```go
  name := r.FormValue("name")
  address := r.FormValue("address")
  ```

  Retrieves values for `name` and `address` fields from the form.

- **Responds to Client**: Writes the form values back in the HTTP response.

### 2. **`helloHandler`**

```go
func helloHandler(w http.ResponseWriter, r *http.Request)
```

- **Path Check**: Only handles `/hello`. Anything else results in 404.
- **Method Check**: Only allows `GET` requests.
- **Simple Response**: Responds with `"hello!"`.

### 3. **`main()`**

```go
fileServer := http.FileServer(http.Dir("./static"))
http.Handle("/", fileServer)
```

- Serves static files (like HTML, CSS, JS) from a folder named `static`.
- A request to `/` or `/index.html` would return `static/index.html` if it exists.

```go
http.HandleFunc("/form", formHandler)
http.HandleFunc("/hello", helloHandler)
```

- Registers route handlers.

```go
http.ListenAndServe(":8080", nil)
```

- Starts the server on port **8080**.

## 🌐 Example Use

### 1. **Form Submission Example**

Assuming `static/form.html` contains:

```html
<form action="/form" method="POST">
  Name: <input type="text" name="name" /> Address:
  <input type="text" name="address" />
  <input type="submit" />
</form>
```

When you fill and submit the form:

- The browser sends a POST request to `/form`
- The server parses it and replies with:

```
POST request successful
Name = Alice
Address = Wonderland
```

## ✅ Summary

| Path     | Method | Purpose                  |
| -------- | ------ | ------------------------ |
| `/`      | GET    | Serve static files       |
| `/hello` | GET    | Simple "hello" response  |
| `/form`  | POST   | Handles form submissions |
