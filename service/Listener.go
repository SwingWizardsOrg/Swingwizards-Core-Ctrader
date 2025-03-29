package service

import (
	"fmt"
	"log"
	"math"
	"time"

	"ctraderapi/mappers"
	"ctraderapi/messages/github.com/Carlosokumu/messages"
	"ctraderapi/middlewares"
	"ctraderapi/models"

	"google.golang.org/protobuf/proto"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second
)

func ReadCtraderMessages(conn *websocket.Conn, handler middlewares.Client) {
	defer func() {
		//conn.Close()
	}()

	conn.SetReadDeadline(time.Now().Add(pongWait))
	conn.SetPongHandler(func(string) error { conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		msg := &messages.ProtoMessage{}
		_, readmessage, err := conn.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		unmarsherr := proto.Unmarshal(readmessage, msg)

		if unmarsherr != nil {
			fmt.Println(unmarsherr)
		}

		handler.Hub.Protos <- messages.ProtoMessage{
			PayloadType: msg.PayloadType,
			Payload:     msg.Payload,
			ClientMsgId: msg.ClientMsgId,
		}
	}
}

func CollectAllMessages(h *middlewares.Hub, conn *websocket.Conn, appConn *websocket.Conn) {

	go func() {

		for {

			//Order Matters else channels will block hence no execution
			symbolmodels := <-h.SymbolModelChannel
			swingtrader := <-h.TraderResChannnel
			accountorders := <-h.AccountOrdersChannel

			accountModel := models.AccountModel{}

			accountModel.Symbols = symbolmodels

			Trader := messages.ProtoOATraderRes{}
			err := proto.Unmarshal(swingtrader.Payload, &Trader)
			if err != nil {
				log.Fatal(err)
			}
			accountModel.DepositAsset = mappers.SwingAsset{
				AssetId: *Trader.Trader.DepositAssetId,
			}
			accountModel.SwingTrader = models.Trader{
				Balance:     *Trader.Trader.Balance,
				TraderLogin: *Trader.Trader.TraderLogin,
			}

			var positions []models.MarketOrderModel
			for _, position := range accountorders.Position {
				var positionSymbol models.SymbolModel
				for _, symbol := range accountModel.Symbols {
					if symbol.Id == *position.TradeData.SymbolId {
						positionSymbol = symbol
						break
					}

				}
				mareketOrderModel := models.MarketOrderModel{}
				mareketOrderModel.Ordermodel.Symbol = positionSymbol
				mareketOrderModel.Ordermodel.TradeSide = *position.TradeData.TradeSide
				mareketOrderModel.Ordermodel.Price = *position.Price
				mareketOrderModel.Ordermodel.Volume = *position.TradeData.Volume
				mareketOrderModel.Ordermodel.OpenTime = *position.TradeData.OpenTimestamp
				mareketOrderModel.Commision = *position.Commission
				mareketOrderModel.Swap = *position.Swap
				mareketOrderModel.MoneyDigits = int32(*position.MoneyDigits)
				mareketOrderModel.CommissionMonetary = float64(mareketOrderModel.Commision / int64(math.Pow(10, float64(mareketOrderModel.MoneyDigits))))
				mareketOrderModel.DoubleCommissionMonetary = mareketOrderModel.CommissionMonetary * 2
				power := math.Pow(10, float64(mareketOrderModel.MoneyDigits))
				SwapMoney := float64(*position.Swap) / float64(power)
				mareketOrderModel.SwapMonetary = SwapMoney
				positions = append(positions, mareketOrderModel)
			}
			accountModel.Positions = positions

			h.AccountModelChannel <- accountModel

			for _, symbomodel := range accountModel.Symbols {
				if symbomodel.QuoteAsset.AssetId != accountModel.DepositAsset.AssetId {
					symbomodel.ConversionSymbols = append(symbomodel.ConversionSymbols)
				} else {

				}

			}
		}
	}()

}

func ListenSpots(h *middlewares.Hub, client *middlewares.Client) {
	go func() {

		for {
			select {
			case Event, ok := <-h.SpotEventChannel:

				if !ok {
					// The hub closed the channel.
					client.Appconn.WriteMessage(websocket.CloseMessage, []byte{})
					return
				}
				event := Event
				spotEvent := &messages.ProtoOASpotEvent{}
				unmarsherr := proto.Unmarshal(event.Payload, spotEvent)
				if unmarsherr != nil {
					log.Fatal(unmarsherr)
				}
				client.Appconn.WriteJSON(spotEvent)

			case _, ok := <-h.SubChannel:
				if !ok {
					// The hub closed the channel.
					client.Appconn.WriteMessage(websocket.CloseMessage, []byte{})
					return
				}
				SendSubscribeSpotsRequest(client.Conn)

			}
		}

	}()
}
