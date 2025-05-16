package natsclient

import (
	"encoding/json"
	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/repository"
	"github.com/nats-io/nats.go"
	"log"
)

type SubmissionJudgedEvent struct {
	SubmissionID uint   `json:"submission_id"`
	UserID       uint   `json:"user_id"`
	ProblemID    uint   `json:"problem_id"`
	Status       string `json:"status"`
}

type Subscriber struct {
	conn *nats.Conn
}

func NewSubscriber(url string) (*Subscriber, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return &Subscriber{conn: nc}, nil
}

func (s *Subscriber) Subscribe(subject string, handler func(msg []byte)) error {
	_, err := s.conn.Subscribe(subject, func(m *nats.Msg) {
		handler(m.Data)
	})
	return err
}

func (s *Subscriber) HandleSubmissionJudged(repo *repository.ProblemRepo) error {
	return s.Subscribe("submission.judged", func(msg []byte) {
		var event SubmissionJudgedEvent

		// Parse incoming message
		if err := json.Unmarshal(msg, &event); err != nil {
			log.Printf("[submission.judged] JSON unmarshal error: %v", err)
			return
		}

		log.Printf("[submission.judged] Event received: SubmissionID=%d, UserID=%d, ProblemID=%d, Status=%s",
			event.SubmissionID, event.UserID, event.ProblemID, event.Status)

		// Update submission count
		if err := repo.IncrementSubmissionCount(event.ProblemID); err != nil {
			log.Printf("[submission.judged] Failed to increment submission count for ProblemID=%d: %v", event.ProblemID, err)
			return
		}
		log.Printf("[submission.judged] Submission count incremented for ProblemID=%d", event.ProblemID)

		// If accepted, update accepted count
		if event.Status == "Accepted" {
			if err := repo.IncrementAcceptedCount(event.ProblemID); err != nil {
				log.Printf("[submission.judged] Failed to increment accepted count for ProblemID=%d: %v", event.ProblemID, err)
				return
			}
			log.Printf("[submission.judged] Accepted count incremented for ProblemID=%d", event.ProblemID)
		}
	})
}
