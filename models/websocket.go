package models

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type BaseResponse struct {
	ResponseType ResponseType
	Symbol       string
}

type TickerResponse struct {
	Ticker
	BaseResponse
}

type TradesResponse struct {
	Trades []Trade
	BaseResponse
}

type TradeResponse struct {
	Trade
	BaseResponse
}

type OrderBookResponse struct {
	OrderBook
	BaseResponse
}

type FillResponse struct {
	Fill
	BaseResponse
}

type OrdersResponse struct {
	Order
	BaseResponse
}

type WSRequest struct {
	ChannelType ChannelType `json:"channel"`
	Market      string      `json:"market"`
	Op          Operation   `json:"op"`
}

type WSRequestAuthorize struct {
	Args map[string]interface{} `json:"args"`
	Op   Operation              `json:"op"`
}

type WsResponse struct {
	ChannelType  ChannelType     `json:"channel"`
	Market       string          `json:"market"`
	ResponseType ResponseType    `json:"type"`
	Code         int             `json:"code"`
	Message      string          `json:"msg"`
	Data         json.RawMessage `json:"data"`
}

func (wr *WsResponse) MapToTradesResponse() (*TradesResponse, error) {
	var trades []Trade
	err := json.Unmarshal(wr.Data, &trades)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &TradesResponse{
		Trades: trades,
		BaseResponse: BaseResponse{
			ResponseType: wr.ResponseType,
			Symbol:       wr.Market,
		},
	}, nil
}

func (wr *WsResponse) MapToTickerResponse() (*TickerResponse, error) {
	ticker := Ticker{}
	err := json.Unmarshal(wr.Data, &ticker)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &TickerResponse{
		Ticker: ticker,
		BaseResponse: BaseResponse{
			ResponseType: wr.ResponseType,
			Symbol:       wr.Market,
		},
	}, nil
}

func (wr *WsResponse) MapToOrderBookResponse() (*OrderBookResponse, error) {
	book := OrderBook{}
	err := json.Unmarshal(wr.Data, &book)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &OrderBookResponse{
		OrderBook: book,
		BaseResponse: BaseResponse{
			ResponseType: wr.ResponseType,
			Symbol:       wr.Market,
		},
	}, nil
}

func (wr *WsResponse) MapToFillResponse() (*FillResponse, error) {
	fill := Fill{}
	err := json.Unmarshal(wr.Data, &fill)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &FillResponse{
		Fill: fill,
		BaseResponse: BaseResponse{
			ResponseType: wr.ResponseType,
		},
	}, nil
}

func (wr *WsResponse) MapToOrdersResponse() (*OrdersResponse, error) {
	order := Order{}
	err := json.Unmarshal(wr.Data, &order)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &OrdersResponse{
		Order: order,
		BaseResponse: BaseResponse{
			ResponseType: wr.ResponseType,
		},
	}, nil
}
