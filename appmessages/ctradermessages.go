package appmessages

import (
	"fmt"

	"ctraderapi/messages/github.com/Carlosokumu/messages"

	"google.golang.org/protobuf/proto"

	"github.com/gorilla/websocket"
)

// Will Send an Account Auth Request to Ctrader.
func SendPositionsRequest(conn *websocket.Conn) error {
	var tt = uint32(messages.ProtoOAPayloadType_PROTO_OA_RECONCILE_REQ)
	id := int64(25675710)
	nmess := "443"

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
		ClientMsgId: &nmess,
	}

	msgB, err = proto.Marshal(msgP)
	if err != nil {
		return err
	}

	err = conn.WriteMessage(2, msgB)
	if err != nil {
		return fmt.Errorf("Failed to send positions request %v", err.Error())
	}
	return nil
}

// Request Symbol Information By Id on the Ctrader Platform
func SendSymbolRequest(conn *websocket.Conn) error {
	var payloadtype = uint32(messages.ProtoOAPayloadType_PROTO_OA_SYMBOL_BY_ID_REQ)
	id := int64(25675710)
	ids := []int64{1, 2}
	nmess := "Symbol_request"

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
		ClientMsgId: &nmess,
	}

	msgB, err = proto.Marshal(msgP)
	if err != nil {
		return err
	}

	err = conn.WriteMessage(2, msgB)
	if err != nil {
		return fmt.Errorf("Failed to send symbol request %v", err)
	}
	
	return nil
}

func SendSubscribeSpotsRequest(conn *websocket.Conn) error {
	var payloadtype = uint32(messages.ProtoOAPayloadType_PROTO_OA_SUBSCRIBE_SPOTS_REQ)
	id := int64(25675710)
	ids := []int64{1, 2}
	nmess := "Subscribe_request"

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
		ClientMsgId: &nmess,
	}
	msgB, err = proto.Marshal(msgP)
	if err != nil {
		return err
	}
	err = conn.WriteMessage(2, msgB)
	if err != nil {
		return fmt.Errorf("Failed to subscribe to spots %v", err)
	}
	return nil
}

func SendProtoOAsymbolConversion(conn *websocket.Conn) error {
	var payloadtype = uint32(messages.ProtoOAPayloadType_PROTO_OA_SYMBOLS_FOR_CONVERSION_REQ)
	nmess := "Subscribe_request"
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
		ClientMsgId: &nmess,
	}
	msgB, err = proto.Marshal(msgP)
	if err != nil {
		return err
	}
	err = conn.WriteMessage(2, msgB)
	if err != nil {
		return fmt.Errorf("Failed to send symbol conversation request for %v ")
	}
	return nil
}

func SendProtoAssetListReq(conn *websocket.Conn) error {
	var payloadtype = uint32(messages.ProtoOAPayloadType_PROTO_OA_ASSET_LIST_REQ)
	fmt.Println(payloadtype)
	id := int64(25675710)
	nmess := "asset_req"

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
		ClientMsgId: &nmess,
	}

	msgB, err = proto.Marshal(msgP)
	if err != nil {
		return err
	}
	err = conn.WriteMessage(2, msgB)
	if err != nil {
		return err
	}
	return nil
}

// Request for all the Symbol's available in a given Trading account
func SendSymbolListRequest(conn *websocket.Conn) error {
	var payloadtype = uint32(messages.ProtoOAPayloadType_PROTO_OA_SYMBOLS_LIST_REQ)
	id := int64(25675710)
	nmess := "symbols_req"

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
		ClientMsgId: &nmess,
	}

	msgB, err = proto.Marshal(msgP)
	if err != nil {
		return err
	}

	err = conn.WriteMessage(2, msgB)
	if err != nil {
		return fmt.Errorf("Failed to Send symbol list request %v", err)
	}
	return nil
}
