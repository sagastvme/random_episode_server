package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Ruta estática para los videos
	app.Static("/", "./videos")

	app.Get("/", func(c *fiber.Ctx) error {
		// Leer los archivos en el directorio de videos
		videos, err := os.ReadDir("./videos")
		if err != nil {
			return c.Status(500).SendString("Error reading public directory")
		}
		if len(videos) == 0 {
			return c.Status(404).SendString("No files found in public directory")
		}

		// Filtrar solo los archivos .m3u8
		var m3u8Files []string
		for _, file := range videos {
			if strings.HasSuffix(file.Name(), ".m3u8") {
				m3u8Files = append(m3u8Files, file.Name())
			}
		}

		// Si no hay archivos .m3u8, enviar un error
		if len(m3u8Files) == 0 {
			return c.Status(404).SendString("No .m3u8 files found")
		}

		// Seleccionar un archivo .m3u8 aleatorio
		randomIndex := rand.Intn(len(m3u8Files))
		selectedFile := m3u8Files[randomIndex]
		videoPath := "videos/" + selectedFile

		fmt.Println("Randomly selected video:", videoPath)

		// Renderizar la vista con el archivo seleccionado
		return c.Render("index", fiber.Map{
			"Title":       "Go Fiber Video Server",
			"Description": "Serving random videos using Go Fiber",
			"Path":        videoPath,
		})
	})

	// Configuración del servidor
	app.Static("/videos", "./videos", fiber.Static{
		Browse:   false,
		MaxAge:   3600,  // Cache videos for performance
		Download: false,
	})

	// Iniciar el servidor en el puerto 3000
	log.Fatal(app.Listen(":3000"))
}
