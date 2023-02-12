package observer

import (
	goesi "github.com/antihax/goesi/esi"
)

type ZkilResponse struct {
	Package ZkilPackage `json:"package"`
}

type ZkilPackage struct {
	KillID            int64                                       `json:"killID"`
	Killmail          *goesi.GetKillmailsKillmailIdKillmailHashOk `json:"killmail"`
	ZKillmailMetadata *ZKillmailMetadata                          `json:"zkb"`
}

type ZKillmailMetadata struct {
	LocationID   int32   `json:"locationID"`
	Hash         string  `json:"hash"`
	FittedValue  float64 `json:"fittedValue"`
	DroppedValue float64 `json:"droppedValue"`
	TotalValue   float64 `json:"totalValue"`
	Points       int     `json:"points"`
	Npc          bool    `json:"npc"`
	Solo         bool    `json:"solo"`
	Awox         bool    `json:"awox"`
	Href         string  `json:"href"`
}

type Route struct {
	DiscordWebhook DiscordWebhook
	IsLoss         bool
}

type RoutedZkilResponse struct {
	ZkilResponse  *ZkilResponse
	MatchedRoutes []Route
}
