README

Event Management Application

This is a simple Event Management Application developed in Go, using the Gin web framework and MSSQL as the database.

Getting Started

Follow the instructions below to set up the project on your local machine.

Prerequisites Install Go (version 1.16 or later) Install MSSQL Server Clone the repository Clone the repository to your local machine:

Install dependencies: go mod download

Build and run the application Build the application:

go build

API Endpoints

The following API endpoints are available:

Add a user to a specific event: POST /users/:id/events

Create a meeting and its invitations: POST /meetings POST /api/meetings { "event_id": 1, "organizer_id": 1, "datetime": "2023-03-21", "invitee_ids": [1] }

Update an invitation's status: PUT /invitations/:id

Get all invitations for a user: GET /users/:id/invitations Get all meetings for a user: GET /users/:id/meetings Running Tests

To run the tests, execute the following command in the project directory:

go test -v ./routes This will run the test in the routes file.
