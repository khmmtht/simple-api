package main

import (
    "bytes"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "io"
    "net/http"
)

type ForwardRequest struct {
    Method  string            `json:"method"`
    Url     string            `json:"url"`
    Body    string            `json:"body"`
    Headers map[string]string `json:"headers"`
}

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
    e := echo.New()
    e.Use(middleware.Logger())
    e.GET("/hello", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
    e.POST("/send_request", func(c echo.Context) error {
        body := new(ForwardRequest)
        if err := c.Bind(body); err != nil {
            return err
        }

        var b []byte
        if body.Body != "" {
            b = []byte(body.Body)
        }
        req, err := http.NewRequest(body.Method, body.Url, bytes.NewBuffer(b))
        if err != nil {
            return err
        }

        for k, v := range body.Headers {
            req.Header.Set("Content-Type", "application/json")
            req.Header.Set(k, v)
        }

        client := &http.Client{}
        resp, err := client.Do(req)
        if err != nil {
            panic(err)
        }
        defer func(Body io.ReadCloser) {
            err := Body.Close()
            if err != nil {
                panic(err)
            }
        }(resp.Body)

        respBody, _ := io.ReadAll(resp.Body)

        return c.String(resp.StatusCode, string(respBody))
    })

    e.Logger.Fatal(e.Start(":8080"))
}
