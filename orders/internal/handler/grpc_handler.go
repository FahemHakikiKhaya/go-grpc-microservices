package handler

import (
	"context"
	"errors"
	"sync"

	menuPb "github.com/fahemhakikikhaya/common/api/menu"
	orderPb "github.com/fahemhakikikhaya/common/api/order"
	"github.com/fahemhakikikhaya/go-microservices-orders/internal/domain"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	orderPb.UnimplementedOrderServiceServer
	menuServiceClient menuPb.MenuServiceClient
	service domain.OrderService
}

func NewGRPCHandler(grpcServer *grpc.Server, service domain.OrderService, client menuPb.MenuServiceClient) {
	handler := &grpcHandler{service: service, menuServiceClient: client}
	orderPb.RegisterOrderServiceServer(grpcServer, handler)
}

func (h *grpcHandler)  CreateOrder(ctx context.Context, r *orderPb.CreateOrderRequest) (*orderPb.CreateOrderResponse, error) {
	// return nil, fmt.Errorf("some errors !!!")
	items := r.GetItems()

	var wg sync.WaitGroup
	var mu sync.Mutex
	var menus   []*menuPb.Menu
	var orderItems []domain.OrderItem
	var errList []error
	
	for _, item := range items {
		wg.Add(1)

		go func(item *orderPb.OrderItem) {
			defer wg.Done()

			res, err := h.menuServiceClient.GetMenu(ctx, &menuPb.GetMenuRequest{MenuId: item.MenuId})

			if err != nil {
				mu.Lock()
				errList = append(errList, err)
				mu.Unlock()
				return
			}

			mu.Lock()
			menus = append(menus, res.Menu)
			orderItems = append(orderItems, domain.OrderItem{
				MenuID: item.GetMenuId(),
				MenuName: item.GetMenuName(),
				Quantity: item.GetQuantity(),
				UnitPrice: item.GetUnitPrice(),
			})
			mu.Unlock()
		}(item)
	}

	wg.Wait()


	for _, menu := range menus {
		if !menu.Available {
			return nil, errors.New("one or more menu are unavailable")
		}
	}

	order := &domain.Order{
		UserID: r.GetCustomerID(),
		Items: orderItems,
	}

	createdOrder, err := h.service.CreateOrder(ctx, order)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	var pbOrderItems []*orderPb.OrderItem

	for _, orderItem := range createdOrder.Items {
		pbOrderItems = append(pbOrderItems, &orderPb.OrderItem{
			MenuId: orderItem.MenuID,
			MenuName: orderItem.MenuName,
			Quantity: orderItem.Quantity,
			UnitPrice: orderItem.UnitPrice,
		})
	}

	response := &orderPb.CreateOrderResponse{
		Order: &orderPb.Order{
			Id: createdOrder.ID,
			UserId: createdOrder.UserID,
			Status: createdOrder.Status,
			Items: pbOrderItems,
			TotalPrice: createdOrder.TotalPrice,
		},
	}


	return response, nil
}