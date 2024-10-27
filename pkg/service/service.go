package service

import (
	pb "github.com/relrin/fame-events/proto"
)

type EventService struct {
	pb.UnimplementedAppCallbackServer
}

func Init() *EventService {
	return &EventService{}
}
