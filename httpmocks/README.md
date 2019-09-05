# httpmocks

## HTTP Response Body Mock

Creating a mock io.ReadCloser for a response body:
````go
    s := []string{"1", "2", "3"}
    b, err := json.Marshal(s)
    
    body := httpmocks.NewReadCloserMock(b, nil)
````

### HTTP Response Mock
Creating a successful http.Response with status code 200 with the desired a body:
```go
    s := []string{"1", "2", "3"}
    b, err := json.Marshal(s))
    
    body := httpmocks.NewReadCloserMock(b, nil)
    resp := httpmocks.NewResponseMock(body, 200)
```

Creating an unsuccessful response:
```go
    resp := httpmocks.NewResponseMock(nil, 500)
```