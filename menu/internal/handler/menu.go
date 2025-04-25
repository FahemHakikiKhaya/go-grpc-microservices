package handler

import (
	"context"

	menuPb "github.com/fahemhakikikhaya/common/api/menu"
	"github.com/fahemhakikikhaya/go-microservices-menu/internal/domain"
	"google.golang.org/grpc"
)

type menuHandler struct {
	service domain.MenuService
	menuPb.UnimplementedMenuServiceServer
}

func NewMenuHandler(grpcServer *grpc.Server, service domain.MenuService) {
	handler := &menuHandler{service: service}
	menuPb.RegisterMenuServiceServer(grpcServer, handler)
}

func (h *menuHandler) GetMenu(ctx context.Context, r *menuPb.GetMenuRequest) (*menuPb.GetMenuResponse, error) {
	menuId := r.GetMenuId()
	menu, err := h.service.GetMenuByID(ctx, menuId)
	

	if err != nil {
		return nil, err
	}

	pbMenu := &menuPb.Menu{
		Id: menu.ID,
		Name: menu.Name,
		Description: menu.Description,
		Price: menu.Price,
		Available: menu.Available,
	}

	return &menuPb.GetMenuResponse{
		Menu: pbMenu,
	}, nil
}

func (h *menuHandler) CreateMenu(ctx context.Context, r *menuPb.CreateMenuRequest) (*menuPb.CreateMenuResponse, error) {
	menu := &domain.Menu{
		Name: r.GetName(),
		Description: r.GetDescription(),
		Price: r.GetPrice(),
		Available: r.GetAvailable(),
	}

	createdMenu, err := h.service.CreateMenu(ctx, menu)

	if err != nil {
		return nil, err
	}


	pbMenu := &menuPb.Menu{
		Id: createdMenu.ID,
		Name: createdMenu.Name,
		Description: createdMenu.Description,
		Available: createdMenu.Available,
		Price: createdMenu.Price,
	}

	return &menuPb.CreateMenuResponse{
		Menu: pbMenu,
	}, nil
}