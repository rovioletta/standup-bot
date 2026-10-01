package teams

import (
	"context"
	"fmt"
	"time"
)

func (s *Service) GetMembersToNotify(ctx context.Context) ([]string, error) {
	notificationHour := uint16(time.Now().Hour())
	members, err := s.queries.GetMembersToNotify(ctx, notificationHour)
	if err != nil {
		return nil, fmt.Errorf("failed to get members from db: %w", err)
	}

	return members, nil
}
