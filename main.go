package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"
)

// create function init
func init() {
	// load variable from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// load template
	engine := html.New("./views", ".html")

	// create go fiber app
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	// seave static assets
	app.Static("/static", "./public")
	// create route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Home",
		})
	})
	// create a post route to accept upload file
	app.Post("/", func(c *fiber.Ctx) error {
		// get file from form
		file, err := c.FormFile("upload")
		if err != nil {
			return err
		}
		// save file to public folder
		// 1st param is file that we get from form
		// 2nd param is path to save file
		err = c.SaveFile(file, "./public/uploads/"+file.Filename)
		if err != nil {
			return err
		}
		// return success message
		return c.Render("index", fiber.Map{
			"Title": "Home",
			"Msg":   "File uploaded successfully!",
			"File":  file,
		})
	})

	// start app with port from .env file
	app.Listen(os.Getenv("HOST") + ":" + os.Getenv("PORT"))
}
