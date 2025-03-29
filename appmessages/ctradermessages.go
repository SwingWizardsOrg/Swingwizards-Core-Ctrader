package appmessages

import (
	"fmt"

	"ctraderapi/messages/github.com/Carlosokumu/messages"

	"google.golang.org/protobuf/proto"

	"github.com/gorilla/websocket"
)

const (
	MESSAGE_TYPE = websocket.BinaryMessage
)

var (
	REQUEST_POSITIONS       = "REQ_POSITIONS"
	REQUEST_SYMBOLS         = "REQ_SYMBOLS"
	REQUEST_SUBSCRIBE_SPOTS = "REQ_SUBSCRIBE_SPOTS"
	REQ_SUBSCRIBE           = "REQ_SUBSCRIBE"
	REQ_SYMBOL_CONVERSION   = "REQ_SYMBOL_CONVERSION"
	REQ_ASSET_LIST          = "REQ_ASSET_LIST"
	REQ_SYMBOL_LIST         = "REQ_SYMBOL_LIST"
)

// SendPositionsRequest requests to get Trader's current open positions and pending orders data.
func SendPositionsRequest(conn *websocket.Conn) error {
	var tt = uint32(messages.ProtoOAPayloadType_PROTO_OA_RECONCILE_REQ)
	id := int64(25675710)

	msgReq := &messages.ProtoOAReconcileReq{
		CtidTraderAccountId: &id,
	}

	msgB, err := proto.Marshal(msgReq)
	if err != nil {
		return err
	}

	msgP := &messages.ProtoMessage{
		PayloadType: &tt,
		Payload:     msgB,
		ClientMsgId: &REQUEST_POSITIONS,
	}

	msgB, err = proto.Marshal(msgP)
	if err != nil {
		return err
	}

	err = conn.WriteMessage(MESSAGE_TYPE, msgB)
	if err != nil {
		return fmt.Errorf("Failed to send positions request: %w", err)
	}
	return nil
}

// SendSymbolRequest  is a request to get a full symbol entity.
func SendSymbolRequest(conn *websocket.Conn) error {
	var payloadtype = uint32(messages.ProtoOAPayloadType_PROTO_OA_SYMBOL_BY_ID_REQ)
	id := int64(25675710)
	ids := []int64{1, 2}

	msgReq := &messages.ProtoOASymbolByIdReq{
		CtidTraderAccountId: &id,
		SymbolId:            ids,
	}

	msgB, err := proto.Marshal(msgReq)
	if err != nil {
		return err
	}

	msgP := &messages.ProtoMessage{
		PayloadType: &payloadtype,
		Payload:     msgB,
		ClientMsgId: &REQUEST_SYMBOLS,
	}

	msgB, err = proto.Marshal(msgP)
	if err != nil {
		return err
	}

	err = conn.WriteMessage(MESSAGE_TYPE, msgB)
	if err != nil {
		return fmt.Errorf("Failed to send symbol request %w:", err)
	}

	return nil
}

/*
SendSubscribeSpotsRequest Request for subscribing on spot events of the specified symbol.
After successful subscription you'll receive technical ProtoOASpotEvent with latest price, after which you'll start receiving updates on prices via consequent ProtoOASpotEvents.
*/
func SendSubscribeSpotsRequest(conn *websocket.Conn) error {
	var payloadtype = uint32(messages.ProtoOAPayloadType_PROTO_OA_SUBSCRIBE_SPOTS_REQ)
	id := int64(25675710)
	ids := []int64{1, 2}

	msgReq := &messages.ProtoOASubscribeSpotsReq{
		CtidTraderAccountId: &id,
		SymbolId:            ids,
	}
	msgB, err := proto.Marshal(msgReq)
	if err != nil {
		return err
	}

	msgP := &messages.ProtoMessage{
		PayloadType: &payloadtype,
		Payload:     msgB,
		ClientMsgId: &REQUEST_SUBSCRIBE_SPOTS,
	}
	msgB, err = proto.Marshal(msgP)
	if err != nil {
		return err
	}
	err = conn.WriteMessage(MESSAGE_TYPE, msgB)
	if err != nil {
		return fmt.Errorf("Failed to subscribe to spots %w", err)
	}
	return nil
}

// SendProtoOAsymbolConversion is a request for getting a conversion chain between two assets that consists of several symbols.
func SendProtoOAsymbolConversion(conn *websocket.Conn) error {
	var payloadtype = uint32(messages.ProtoOAPayloadType_PROTO_OA_SYMBOLS_FOR_CONVERSION_REQ)
	id := int64(25675710)
	firstasset := int64(6)
	lastasset := int64(4)

	msgReq := &messages.ProtoOASymbolsForConversionReq{
		CtidTraderAccountId: &id,
		FirstAssetId:        &firstasset,
		LastAssetId:         &lastasset,
	}
	msgB, err := proto.Marshal(msgReq)
	if err != nil {
		return err
	}
	msgP := &messages.ProtoMessage{
		PayloadType: &payloadtype,
		Payload:     msgB,
		ClientMsgId: &REQ_SYMBOL_CONVERSION,
	}
	msgB, err = proto.Marshal(msgP)
	if err != nil {
		return err
	}
	err = conn.WriteMessage(MESSAGE_TYPE, msgB)
	if err != nil {
		return fmt.Errorf("Failed to send symbol conversation request for %w", err)
	}
	return nil
}

// SendProtoAssetListReq requests for the list of assets available for a trader's account.
func SendProtoAssetListReq(conn *websocket.Conn) error {
	var payloadtype = uint32(messages.ProtoOAPayloadType_PROTO_OA_ASSET_LIST_REQ)
	id := int64(25675710)

	msgReq := &messages.ProtoOAAssetListReq{
		CtidTraderAccountId: &id,
	}

	msgB, err := proto.Marshal(msgReq)
	if err != nil {
		return err
	}

	msgP := &messages.ProtoMessage{
		PayloadType: &payloadtype,
		Payload:     msgB,
		ClientMsgId: &REQ_ASSET_LIST,
	}

	msgB, err = proto.Marshal(msgP)
	if err != nil {
		return err
	}
	err = conn.WriteMessage(MESSAGE_TYPE, msgB)
	if err != nil {
		return fmt.Errorf("Failed to request asset list: %w", err)
	}
	return nil
}

// SendSymbolListRequest request for all the symbols available in a given Trading account
func SendSymbolListRequest(conn *websocket.Conn) error {
	var payloadtype = uint32(messages.ProtoOAPayloadType_PROTO_OA_SYMBOLS_LIST_REQ)
	id := int64(25675710)

	msgReq := &messages.ProtoOASymbolsListReq{
		CtidTraderAccountId: &id,
	}

	msgB, err := proto.Marshal(msgReq)
	if err != nil {
		return err
	}

	msgP := &messages.ProtoMessage{
		PayloadType: &payloadtype,
		Payload:     msgB,
		ClientMsgId: &REQ_SYMBOL_LIST,
	}

	msgB, err = proto.Marshal(msgP)
	if err != nil {
		return err
	}

	err = conn.WriteMessage(MESSAGE_TYPE, msgB)
	if err != nil {
		return fmt.Errorf("Failed to Send symbol list request %w: ", err)
	}
	return nil
}
