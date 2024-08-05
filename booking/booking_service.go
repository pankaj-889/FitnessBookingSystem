package booking

import (
	"FitnessClass/model"
	"errors"
	"sync"

	"FitnessClass/class"
	"FitnessClass/user"
)

type BookingService struct {
	mu      sync.Mutex
	classes map[int]*model.Class
	users   map[int]*model.User
}

var bookingServiceInstance *BookingService
var bookingServiceOnce sync.Once

func GetBookingService() *BookingService {
	bookingServiceOnce.Do(func() {
		bookingServiceInstance = &BookingService{
			classes: class.GetClassService().Classes,
			users:   user.GetUserService().Users,
		}
	})
	return bookingServiceInstance
}

func (bs *BookingService) BookClass(userID, classID int) error {
	bs.mu.Lock()
	defer bs.mu.Unlock()

	user, ok := bs.users[userID]
	if !ok {
		return errors.New("user not found")
	}

	class, ok := bs.classes[classID]
	if !ok {
		return errors.New("class not found")
	}

	if !user.CanBook() {
		return errors.New("user has reached booking limit")
	}

	booking := &model.Booking{User: user, Class: class}
	if class.AddBooking(booking) {
		user.AddBooking(booking)
		return nil
	} else {
		class.AddToWaitlist(user)
		return errors.New("class is full, added to waitlist")
	}
}

func (bs *BookingService) CancelBooking(userID, classID int) error {
	bs.mu.Lock()
	defer bs.mu.Unlock()

	user, ok := bs.users[userID]
	if !ok {
		return errors.New("user not found")
	}

	class, ok := bs.classes[classID]
	if !ok {
		return errors.New("class not found")
	}

	for _, booking := range user.Bookings {
		if booking.Class.ID == classID {
			class.RemoveBooking(booking)
			user.RemoveBooking(booking)
			return nil
		}
	}

	return errors.New("booking not found")
}
