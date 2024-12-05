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
