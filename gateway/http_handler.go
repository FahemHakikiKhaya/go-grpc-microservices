package main

import (
	"net/http"

	"github.com/fahemhakikikhaya/common"
	menuPb "github.com/fahemhakikikhaya/common/api/menu"
	orderPb "github.com/fahemhakikikhaya/common/api/order"
	"google.golang.org/grpc/status"
)

type handler struct {
	menuService menuPb.MenuServiceClient
	orderService orderPb.OrderServiceClient
}

func NewHandler(orderService orderPb.OrderServiceClient, menuService menuPb.MenuServiceClient) *handler {
	return &handler{orderService: orderService, menuService: menuService}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerID}/orders", h.HandleCreateOrder)
	mux.HandleFunc("POST /api/menus", h.HandleCreateMenu)
}

func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	customerID := r.PathValue(("customerID"))

	var items []*orderPb.OrderItem

	if err := common.ReadJson(r, &items); err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	o, err := h.orderService.CreateOrder(r.Context(), &orderPb.CreateOrderRequest{
		CustomerID: customerID ,
		Items: items,
	})

	rStatus := status.Convert(err)
	
	if rStatus != nil {
		common.WriteError(w, http.StatusBadRequest, rStatus.Message())
		return 
	}

	common.WriteJSON(w, http.StatusOK, o)
}

func (h *handler) HandleCreateMenu(w http.ResponseWriter, r *http.Request) {
	var menu *menuPb.CreateMenuRequest

	if err := common.ReadJson(r, &menu); err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	createdMenu, err := h.menuService.CreateMenu(r.Context(), menu)

	rStatus := status.Convert(err)

	if rStatus != nil {
		common.WriteError(w, http.StatusBadRequest, rStatus.Message())
		return
	}

	common.WriteJSON(w, http.StatusOK, createdMenu)
}