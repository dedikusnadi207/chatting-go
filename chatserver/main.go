package chatserver

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

type messageUnit struct {
	ClientName        string
	ClientUniqueCode  int
	MessageBody       string
	MessageUniqueCode int
}

type messageHandle struct {
	MQue []messageUnit
	mu   sync.Mutex
}

var messageHandleObject = messageHandle{}

type ChatServer struct {
}

func (cs *ChatServer) ChatService(csi Services_ChatServiceServer) error {
	clientUniqueCode := rand.Intn(1e6)
	errch := make(chan error)

	// receive message - init go routine
	go receiveFromStream(csi, clientUniqueCode, errch)

	// send message - init go routine
	go sendToStream(csi, clientUniqueCode, errch)

	return <-errch
}

// Receive Messages
func receiveFromStream(csi Services_ChatServiceServer, clientUniqueCode int, errch chan error) {
	for {
		msg, err := csi.Recv()
		if err != nil {
			log.Printf("Error receiveFromStream :: %v", err)
			errch <- err
		} else {
			messageHandleObject.mu.Lock()

			messageHandleObject.MQue = append(messageHandleObject.MQue, messageUnit{
				ClientName:        msg.Name,
				MessageBody:       msg.Body,
				MessageUniqueCode: rand.Intn(1e8),
				ClientUniqueCode:  clientUniqueCode,
			})
			messageHandleObject.mu.Unlock()
			log.Print(messageHandleObject.MQue[len(messageHandleObject.MQue)-1])
		}
	}
}

// Send Messages
func sendToStream(csi Services_ChatServiceServer, clientUniqueCode int, errch chan error) {
	for {
		for {
			time.Sleep(500 * time.Millisecond)
			messageHandleObject.mu.Lock()
			if len(messageHandleObject.MQue) == 0 {
				messageHandleObject.mu.Unlock()
				break
			}

			msg := messageHandleObject.MQue[0]
			senderUniqueCode := msg.ClientUniqueCode
			senderName4Client := msg.ClientName
			message4Client := msg.MessageBody
			messageHandleObject.mu.Unlock()

			if senderUniqueCode != clientUniqueCode {
				err := csi.Send(&FromServer{Name: senderName4Client, Body: message4Client})
				if err != nil {
					errch <- err
				}
				messageHandleObject.mu.Lock()
				if len(messageHandleObject.MQue) > 1 {
					messageHandleObject.MQue = messageHandleObject.MQue[1:]
				} else {
					messageHandleObject.MQue = []messageUnit{}
				}
				messageHandleObject.mu.Unlock()
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
}
