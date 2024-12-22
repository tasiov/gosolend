package obligations

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

func (s *Service) GetObligation(ctx context.Context, wallet string) (*models.Obligation, error) {
	if err := utils.ValidateAddress(wallet); err != nil {
		return nil, err
	}

	// Implement API call
	return nil, common.ErrNotFound
}
