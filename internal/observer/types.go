package observer

import (
	goesi "github.com/antihax/goesi/esi"
)

// Killmail is one R2Z2 killmail file. It is also the wire-decode target for
// GET /ephemeral/{sequence_id}.json. The full ESI killmail body is inline
// under ESI; Zkb carries zKillboard metadata.
type Killmail struct {
	KillmailID int64                                       `json:"killmail_id"`
	ESI        *goesi.GetKillmailsKillmailIdKillmailHashOk `json:"esi"`
	Zkb        *ZKillmailMetadata                          `json:"zkb"`
}

// ZKillmailMetadata is zKillboard's zkb metadata block. Only the total ISK
// value of the kill is consumed.
type ZKillmailMetadata struct {
	TotalValue float64 `json:"totalValue"`
}

type Route struct {
	DiscordWebhook DiscordWebhook
	IsLoss         bool
}

// RoutedKillmail pairs a killmail with the destination routes it matched.
type RoutedKillmail struct {
	Killmail      *Killmail
	MatchedRoutes []Route
}
