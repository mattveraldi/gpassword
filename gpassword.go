/*
Props to theprimegen for his go-htmx tutorial (https://youtu.be/x7v6SNIgJpE?si=E2KNqXkQX9nFL6a0)
*/
package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mattveraldi/gpassword/internal/generator"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

type PasswordInput struct {
	Password string
}

type IndexProps struct {
	Symbols []string
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())  // logs http requests
	e.Use(middleware.Recover()) // recover from panics
	e.Static("/static", "./static")
	e.Renderer = newTemplate()

	// Get index page
	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index.html", IndexProps{
			Symbols: generator.Symbols,
		})
	})

	// Get a password component
	e.GET("/password", func(c echo.Context) error {
		pass, err := generator.GenerateSecurePassword([]string{})
		if err != nil {
			return c.JSON(500, "Server error")
		}
		return c.Render(200, "password-input.html", PasswordInput{
			Password: pass,
		})
	})

	// Get a password component with configuration
	e.POST("/password", func(c echo.Context) error {
		err := c.Request().ParseForm()
		if err != nil {
			return c.JSON(400, "Wrong request parameters")
		}

		formData := c.Request().Form

		symbolsToRemove := []string{}
		for _, symbol := range generator.Symbols {
			if formData.Get(symbol) != "on" {
				symbolsToRemove = append(symbolsToRemove, symbol)
			}
		}

		pass, err := generator.GenerateSecurePassword(symbolsToRemove)
		if err != nil {
			return c.JSON(500, "Server error")
		}
		return c.Render(200, "password-input.html", PasswordInput{
			Password: pass,
		})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
