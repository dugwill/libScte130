package scte130

import "testing"

var b = []byte(`<?xml version="1.0" encoding="UTF-8" standalone="yes" ?><adm:PlacementStatusNotification identity="86CF2E98-AEBA-4C3A-A3D4-616CF7D74A79" messageId="46df5072-412f-423c-838d-e87028d182fa" version="1.1" system="QAMGW" xmlns:adm="http://www.scte.org/schemas/130-3/2008a/adm" xmlns:core="http://www.scte.org/schemas/130-2/2008a/core" xmlns:schemaLocation="http://www.scte.org/schemas/130-3/2008a/adm SCTE_130-3_2008a.xsd" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:cmcst="http://www.comcast.com/schemas/NGOD/P1/2015/R1V0"><adm:PlayData identityADS="1ea15585-1075-4833-86c6-9321695b5ce4"><adm:SystemContext><adm:Session>569cb3d1-8016-4043-89be-f7c26a7f2848</adm:Session></adm:SystemContext><adm:Client><adm:TerminalAddress type="private:comcast_qam_gateway">McHenry_ATP-1-0573-a</adm:TerminalAddress></adm:Client><adm:SpotScopedEvents><adm:Spot><core:Content><core:AssetRef assetID="CSEK7055899600110000" providerID="spotlight.tv"/><core:Tracking>b4cc532c-8efe-4181-8c80-be9e8951930f|LINEAR_T6|LM5cJfkXV0ekzzHYoBVzZw==</core:Tracking></core:Content></adm:Spot><adm:Events><adm:PlacementStatusEvent time="2019-04-17T12:21:21.063+00:00" type="startPlacement" messageRef="16830b09-f2fe-4bd1-a96b-c1caab3805a3"><adm:SpotNPT scale="1.0">0.000</adm:SpotNPT></adm:PlacementStatusEvent></adm:Events></adm:SpotScopedEvents></adm:PlayData></adm:PlacementStatusNotification>`)

func TestUnMarshalPSN(t *testing.T) {

	psn := NewPSN()

	err := psn.UnMarshalPSN(b)

	if err != nil {
		t.Error("Unmarshal failled. Returned err:", err)
		t.FailNow()
	}

	if psn.Version != "1.1" {
		t.Errorf("Version incorrect. Want: 1.1 Got: %v", psn.Version)
		//t.FailNow()
	}
}

func TestNumEvents(t *testing.T) {

	psn := NewPSN()

	err := psn.UnMarshalPSN(b)

	if err != nil {
		t.Error("Unmarshal failled. Returned err:", err)
		t.FailNow()
	}

	if psn.NumEvents() != 1 {
		t.Errorf("Num Events Incorrect. Want: 1 Got: %v", psn.NumEvents())
	}

}

func TestGetEvents(t *testing.T) {
	psn := NewPSN()

	err := psn.UnMarshalPSN(b)

	if err != nil {
		t.Error("Unmarshal failled. Returned err:", err)
		t.FailNow()
	}

	if len(psn.GetEvents()) != 1 {
		t.Errorf("Length of Event struct incorrect. Want: 1 Got: %v", psn.NumEvents())
	}
}
