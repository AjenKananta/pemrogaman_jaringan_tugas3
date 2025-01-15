package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var db *sql.DB

func main() {
	// Koneksi ke database
	var err error
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/eventdb")
	if err != nil {
		log.Fatal(err)
	}

	// Inisialisasi Fiber
	app := fiber.New()

	// Middleware CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Bisa diatur ke asal spesifik seperti "http://127.0.0.1:5500"
		AllowMethods: "GET,POST,PUT,DELETE",
	}))

	// Routes
	app.Get("/events", getEvents)
	app.Post("/events", createEvent)
	app.Put("/events/:id", updateEvent)
	app.Delete("/events/:id", deleteEvent)

	// Start server
	log.Fatal(app.Listen(":3000"))
}

type Event struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Date     string `json:"date"`
	Time     string `json:"time"`
	Location string `json:"location"`
}

// Get all events
func getEvents(c *fiber.Ctx) error {
	rows, err := db.Query("SELECT id, name, date, time, location FROM events")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()

	events := []Event{}
	for rows.Next() {
		var e Event
		if err := rows.Scan(&e.ID, &e.Name, &e.Date, &e.Time, &e.Location); err != nil {
			return c.Status(500).SendString(err.Error())
		}
		events = append(events, e)
	}

	return c.JSON(events)
}

// Create a new event
func createEvent(c *fiber.Ctx) error {
	var e Event
	if err := c.BodyParser(&e); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	_, err := db.Exec("INSERT INTO events (name, date, time, location) VALUES (?, ?, ?, ?)",
		e.Name, e.Date, e.Time, e.Location)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(201).SendString("Event created")
}

// Update an event
func updateEvent(c *fiber.Ctx) error {
	id := c.Params("id")
	var e Event
	if err := c.BodyParser(&e); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	_, err := db.Exec("UPDATE events SET name = ?, date = ?, time = ?, location = ? WHERE id = ?",
		e.Name, e.Date, e.Time, e.Location, id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendString("Event updated")
}

// Delete an event
func deleteEvent(c *fiber.Ctx) error {
	id := c.Params("id")
	_, err := db.Exec("DELETE FROM events WHERE id = ?", id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendString("Event deleted")
}
