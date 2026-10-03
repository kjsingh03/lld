package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type Notification interface {
	Send(message string, user *User) (string, error)
}

type User struct {
	Name    string `json:"name"`
	UserId  string `json:"userId"`
	Address string `json:"address,omitempty"`
}

type EmailNotification struct{}
type SmsNotification struct{}
type PushNotification struct{}

func (e *EmailNotification) Send(msg string, user *User) (string, error){
	if rand.Intn(2) == 0 {
		return "", errors.New("failed to send email")
	}

	return fmt.Sprint("Email sent successfully to: ", user.Name,": ", msg), nil
}

func (s *SmsNotification) Send(msg string, user *User) (string, error){
	if rand.Intn(2) == 0 {
		return "", errors.New("failed to send SMS")
	}

	return fmt.Sprint("SMS sent successfully to: ", user.Name,": ", msg), nil
}

func (p *PushNotification) Send(msg string, user *User) (string, error){
	if rand.Intn(2) == 0 {
		return "", errors.New("failed to send Push Notification")
	}

	return fmt.Sprint("Push Notification sent successfully to: ", user.Name,": ", msg), nil
}

var _ Notification = (*EmailNotification)(nil)
var _ Notification = (*SmsNotification)(nil)
var _ Notification = (*PushNotification)(nil)

type NotificationFactory struct {
	creators map[string]func() Notification
}

func NewNotificationFactory() *NotificationFactory {
	return &NotificationFactory{
		creators: make(map[string]func() Notification),
	}
}

func (f *NotificationFactory) Register(notificationType string, creator func() Notification) {
	f.creators[notificationType] = creator
}

func (f *NotificationFactory) Create(notificationType string) Notification {
	creator, ok := f.creators[notificationType]

	if !ok {
		return nil
	}

	return creator()
}

func main() {
	user := &User{
		Name:   "Karan",
		UserId: "u1",
	}

	user2 := &User{
		Name:   "Ramu",
		UserId: "u2",
	}

	factory := NewNotificationFactory()

	factory.Register("email", func() Notification {
		return &EmailNotification{}
	})

	factory.Register("sms", func() Notification {
		return &SmsNotification{}
	})

	factory.Register("push", func() Notification {
		return &PushNotification{}
	})

	email := factory.Create("email")
	sms := factory.Create("sms")
	push := factory.Create("push")

	// Email
	res, err := email.Send("Hello, World!", user)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(res)
	}

	// SMS
	res, err = sms.Send("Hey there", user)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(res)
	}

	// Push
	res, err = push.Send("Kya haal hain....", user2)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(res)
	}
}

// Switch based Factory - O/C principle violated - Not closed for modification :)

// import "fmt"

// type Notification interface {
// 	Send(message string, user *User)
// }

// type EmailNotification struct{}

// func (e *EmailNotification) Send(msg string, user *User) {
// 	fmt.Println("Email notification sent to", user.Name, ": ", msg)
// }

// type SmsNotification struct{}

// func (s *SmsNotification) Send(msg string, user *User) {
// 	fmt.Println("Sms notification sent to", user.Name, ": ", msg)
// }

// type User struct {
// 	Name    string `json:"name"`
// 	UserId  string `json:"userId"`
// 	Address string `json:"address,omitempty"`
// }

// func NewNotification(notificationType string) Notification {
// 	switch notificationType {
// 	case "email":
// 		return &EmailNotification{}
// 	case "sms":
// 		return &SmsNotification{}
// 	default:
// 		return nil
// 	}
// }

// func main() {
// 	user := &User{
// 		Name:   "Karan",
// 		UserId: "u1",
// 	}

// 	user2 := &User{
// 		Name:   "Ramu",
// 		UserId: "u2",
// 	}

// 	email := NewNotification("email")
// 	sms := NewNotification("sms")

// 	email.Send("Hello, World!", user)
// 	sms.Send("Hey there", user)
// 	email.Send("Kya haal hain....", user2)

// 	fmt.Println("Hello, World!")
// }
