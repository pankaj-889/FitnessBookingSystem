package main

import (
	"FitnessClass/model"
	"fmt"
	"time"

	"FitnessClass/booking"
	"FitnessClass/class"
	"FitnessClass/user"
)

func main() {

	bookingService := booking.GetBookingService()

	userService := user.GetUserService()
	user1 := userService.RegisterUser("Alice", "alice@example.com", "password", model.Platinum)
	user2 := userService.RegisterUser("Bob", "bob@example.com", "password", model.Gold)

	classService := class.GetClassService()
	class1 := classService.CreateClass(model.Yoga, 2, time.Now().Add(24*time.Hour))

	err := bookingService.BookClass(user1.ID, class1.ID)
	if err != nil {
		fmt.Println("Error booking class:", err)
	}

	err = bookingService.BookClass(user2.ID, class1.ID)
	if err != nil {
		fmt.Println("Error booking class:", err)
	}

	err = bookingService.BookClass(user2.ID, class1.ID)
	if err != nil {
		fmt.Println("Error booking class:", err)
	}

	err = bookingService.CancelBooking(user1.ID, class1.ID)
	if err != nil {
		fmt.Println("Error canceling booking:", err)
	}

	// Check waitlist promotion
	err = bookingService.BookClass(user2.ID, class1.ID) // Should now book the class
	if err != nil {
		fmt.Println("Error booking class:", err)
	}
}

/*
user
id
name
email
pass
tier
[]booking
BookingLimit



class
classTime
capacity
name
id
[]waitlist one class has multiple user who are waiting
[]booking one class has multiple booking

booking
user
classes

*/
