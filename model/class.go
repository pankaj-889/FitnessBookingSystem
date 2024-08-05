package model

import (
	"sync"
	"time"
)

type ClassType int

const (
	Yoga ClassType = iota
	Gym
	Dance
)

type Class struct {
	ID       int
	Type     ClassType
	Capacity int
	Schedule time.Time
	Bookings []*Booking
	Waitlist []*User
	mu       sync.Mutex
}

var classIDCounter = 1

func NewClass(classType ClassType, capacity int, schedule time.Time) *Class {
	class := &Class{
		ID:       classIDCounter,
		Type:     classType,
		Capacity: capacity,
		Schedule: schedule,
	}
	classIDCounter++
	return class
}

func (c *Class) AddBooking(booking *Booking) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if len(c.Bookings) < c.Capacity {
		c.Bookings = append(c.Bookings, booking)
		return true
	}

	return false
}

func (c *Class) AddToWaitlist(user *User) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Waitlist = append(c.Waitlist, user)
}

func (c *Class) RemoveBooking(booking *Booking) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for i, b := range c.Bookings {
		if b == booking {
			c.Bookings = append(c.Bookings[:i], c.Bookings[i+1:]...)
			if len(c.Waitlist) > 0 {
				nextUser := c.Waitlist[0]
				c.Waitlist = c.Waitlist[1:]
				newBooking := &Booking{User: nextUser, Class: c}
				c.Bookings = append(c.Bookings, newBooking)
				nextUser.AddBooking(newBooking)
			}
			return
		}
	}
}
