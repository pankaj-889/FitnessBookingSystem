package model

import (
	"sync"
)

type Tier int

const (
	Platinum Tier = iota
	Gold
	Silver
)

type User struct {
	ID           int
	Name         string
	Email        string
	Password     string
	Tier         Tier
	BookingLimit int
	Bookings     []*Booking
	mu           sync.Mutex
}

var idCounter = 1

func NewUser(name, email, password string, tier Tier) *User {
	user := &User{
		ID:       idCounter,
		Name:     name,
		Email:    email,
		Password: password,
		Tier:     tier,
	}
	idCounter++

	switch tier {
	case Platinum:
		user.BookingLimit = 10
	case Gold:
		user.BookingLimit = 5
	case Silver:
		user.BookingLimit = 3
	}

	return user
}

func (u *User) CanBook() bool {
	u.mu.Lock()
	defer u.mu.Unlock()
	return len(u.Bookings) < u.BookingLimit
}

func (u *User) AddBooking(booking *Booking) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.Bookings = append(u.Bookings, booking)
}

func (u *User) RemoveBooking(booking *Booking) {
	u.mu.Lock()
	defer u.mu.Unlock()
	for i, b := range u.Bookings {
		if b == booking {
			u.Bookings = append(u.Bookings[:i], u.Bookings[i+1:]...)
			return
		}
	}
}
