package parser

import (
	"Square_Pos/app/parser/dto"
	"strconv"
)

type SquareOrderResponse struct {
	Order struct {
		ID        string `json:"id"`
		State     string `json:"state"`
		CreatedAt string `json:"created_at"`
		TableID   string `josn:"location_id"`
		LineItems []struct {
			Name           string `json:"name"`
			Quantity       string `json:"quantity"`
			BasePriceMoney struct {
				Amount int64 `json:"amount"`
			} `json:"base_price_money"`
			TotalMoney struct {
				Amount int64 `json:"amount"`
			} `json:"total_money"`
			Modifiers []struct {
				BasePriceMoney struct {
					Amount int64 `json:"amount"`
				} `json:"base_prize_money"`
				Name string `json:"name"`
				Quantity string `json:"quantity"`
				TotalMoney struct {
					Amount int64 `json:"amount"`
				} `json:"total_prize_money"`
			} `json:"modifiers"`
		} `json:"line_items"`
		NetAmounts struct {
			DiscountMoney struct {
				Amount int64 `json:"amount"`
			} `json:"discount_money"`
			ServiceChargeMoney struct {
				Amount int64 `json:"amount"`
			} `json:"service_charge_money"`
			TaxMoney struct {
				Amount int64 `json:"amount"`
			} `json:"tax_money"`
			TipMoney struct {
				Amount int64 `json:"amount"`
			} `json:"tip_money"`
			TotalMoney struct {
				Amount int64 `json:"amount"`
			} `json:"total_money"`
		} `json:"net_amounts"`
		NetAmountDueMoney struct {
			Amount int64 `json:"amount"`
		} `json:"net_amount_due_money"`
		TotalTipMoney struct {
			Amount int64 `json:"amount"`
		} `json:"total_tip_money"`
	} `json:"order"`
}

func ParseOrder(squareResp SquareOrderResponse) dto.OrderResponse {
	order := dto.OrderResponse{
		ID: squareResp.Order.ID,
		OpenedAt: squareResp.Order.CreatedAt,
		IsClosed: squareResp.Order.State == "COMPLETED",
		Table: squareResp.Order.TableID,
	}

	for _, item := range squareResp.Order.LineItems {
		quantityInt, _ := strconv.Atoi(item.Quantity)
		var modifiers []dto.Modifier
		for _, modifier := range item.Modifiers{
			modQuantity, _ := strconv.Atoi(modifier.Quantity)
			modifiers = append(modifiers, dto.Modifier{
				Name: modifier.Name,
				UnitPrice: modifier.BasePriceMoney.Amount,
				Quantity: modQuantity,
				Amount: modifier.TotalMoney.Amount,
			})
		}
		order.Items = append(order.Items, dto.Item{
			Name: item.Name,
			Comment: "",
			UnitPrice: item.BasePriceMoney.Amount,
			Quantity: quantityInt,
			Discounts: []dto.DiscountDTO{},
			Modifiers: modifiers,
			Amount: item.TotalMoney.Amount,
		})
	}


	order.Totals = dto.Total{
		Discounts: squareResp.Order.NetAmounts.DiscountMoney.Amount,
		Due:           squareResp.Order.NetAmountDueMoney.Amount,
		Tax:           squareResp.Order.NetAmounts.TaxMoney.Amount,
		ServiceCharge: squareResp.Order.NetAmounts.ServiceChargeMoney.Amount,
		Paid:          0, // Square doesn't provide "paid" directly
		Tips:          squareResp.Order.NetAmounts.TipMoney.Amount,
		Total:         squareResp.Order.NetAmounts.TotalMoney.Amount,
	}
	return order
}