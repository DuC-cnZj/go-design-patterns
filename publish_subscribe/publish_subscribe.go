package publish_subscribe

import (
	"errors"
	"time"
)

type Message struct {
	Subscription Subscription
}

type Subscription struct {
	ch chan Message

	Inbox chan Message
}

type Topic struct {
	Subscribers    []Session
	MessageHistory []Message
}

type User struct {
	ID   uint64
	Name string
}

type Session struct {
	User      User
	Timestamp time.Time
}

func (s *Subscription) Publish(msg Message) error {
	if _, ok := <-s.ch; !ok {
		return errors.New("Topic has been closed")
	}

	s.ch <- msg

	return nil
}

func (t *Topic) Subscribe(uid uint64) (Subscription, error) {
	// Get session and create one if it's the first
	for _, subscriber := range t.Subscribers {
		if subscriber.User.ID == uid {
			return Subscription{}, nil
		}
	}
	// Add session to the Topic & MessageHistory
	t.Subscribers = append(
		t.Subscribers,
		Session{
			User:      User{ID: uid, Name: string(uid)},
			Timestamp: time.Now(),
		})

	// Create a subscription
	//t.MessageHistory
}

func (t *Topic) Unsubscribe(Subscription) error {
	// Implementation
}

func (t *Topic) Delete() error {
	// Implementation
}
