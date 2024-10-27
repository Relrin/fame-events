package service

import (
	"context"

	pb "github.com/relrin/fame-events/proto"
)

// ScheduleEvent starts a new round with the given manifest, scenario and a list of teams
func (es *EventService) ScheduleEvent(ctx context.Context, request *pb.ScheduleEventRequest) (*pb.ScheduleEventResponse, error) {
	return &pb.ScheduleEventResponse{}, nil
}
