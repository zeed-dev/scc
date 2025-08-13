package firebase

import (
	"context"

	"firebase.google.com/go/messaging"
)

// SendNotificationToTopic sends a simple notification to a topic using FCM.
func SendNotificationToTopic(topic, title, body string) error {
	if FirebaseApp == nil {
		return nil
	}

	ctx := context.Background()
	client, err := FirebaseApp.Messaging(ctx)
	if err != nil {
		return err
	}

	msg := &messaging.Message{
		Topic: topic,
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
	}

	_, err = client.Send(ctx, msg)
	return err
}
