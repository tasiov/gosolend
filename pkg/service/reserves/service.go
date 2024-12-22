package reserves

import (
	"context"
	"time"

	"github.com/tasiov/gosolend/internal/common"
	"github.com/tasiov/gosolend/internal/utils"
	"github.com/tasiov/gosolend/pkg/models"
)

type Service struct {
	endpoint string
	timeout  time.Duration
}

func NewService(endpoint string, timeout time.Duration) *Service {
	return &Service{
		endpoint: endpoint,
		timeout:  utils.ValidateTimeout(timeout),
	}
}

func (s *Service) GetReserve(ctx context.Context, address string) (*models.Reserve, error) {
	if err := utils.ValidateAddress(address); err != nil {
		return nil, err
	}

	// Implement API call
	return nil, common.ErrNotFound
}
