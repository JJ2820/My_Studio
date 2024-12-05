# My_Studio
Platform for boutiques, studios, and gyms which allows business owners to manage their courses, classes, members, memberships etc,.


Using a simple in-memory data structure to maintain the state of classes and bookings.

Implementation Steps:

    Setup the basic structure: Use net/http package to define the API endpoints.
    Define models:
        Class: Represents a class with attributes like name, start_date, end_date, and capacity.
        Booking: Represents a booking with attributes like member_name and date.
    In-memory storage:
        Use slices to store classes and bookings.
    RESTful Endpoints:
        POST /classes: Create a class.
        POST /bookings: Book a class.

        

glofox-api/
├── main.go                # Entry point of the application
├── routes/
│   ├── routes.go          # Defines and initializes all API routes
├── handlers/
│   ├── classes.go         # Handlers for class-related operations
│   ├── bookings.go        # Handlers for booking-related operations
├── models/
│   ├── class.go           # Data model for a class
│   ├── booking.go         # Data model for a booking
├── storage/
│   ├── in_memory.go       # In-memory storage implementation
│   ├── database.go        # (Optional) Persistent database implementation (e.g., SQLite, PostgreSQL)
├── utils/
│   ├── validation.go      # Utility functions like date validation
├── go.mod                 # Module dependencies
└── go.sum                 # Dependency checksum file



