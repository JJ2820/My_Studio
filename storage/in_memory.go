package storage

import (
	"glofox-api/models"
	"sync"
)

var (
	classes  []models.Class
	bookings []models.Booking
	mutex    sync.Mutex
)

func AddClass(newClass models.Class) {
	mutex.Lock()
	defer mutex.Unlock()
	classes = append(classes, newClass)
}

func AddBooking(newBooking models.Booking) {
	mutex.Lock()
	defer mutex.Unlock()
	bookings = append(bookings, newBooking)
}
