package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

)

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "./videos")
	app.Get("/", func(c *fiber.Ctx) error {
		videos, err := os.ReadDir("./videos")
		if err != nil {
			return c.Status(500).SendString("Error reading public directory")
		}
		if len(videos) == 0 {
			return c.Status(404).SendString("No files found in public directory")
		}
		randomIndex := rand.Intn(len(videos))
		selectedFile := videos[randomIndex].Name()
		videoPath := "videos/" + selectedFile
		fmt.Println("Randomly selected videos:", videoPath)
		return c.Render("index", fiber.Map{
			"Title":         "Go Fiber Video Server",
			"Description":   "Serving random videos using Go Fiber",
			"Path":          videoPath,
		})
	})

	
	app.Static("/videos", "./videos", fiber.Static{
		Browse: false, 
		MaxAge: 3600,  // Cache videos for performance
		Download: false, 
	})
	
	log.Fatal(app.Listen(":3000"))
}
