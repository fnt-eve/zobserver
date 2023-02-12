package observer

import (
	goesi "github.com/antihax/goesi/esi"
	"go.uber.org/zap"
)

type router struct {
	in           chan *ZkilResponse
	out          chan *RoutedZkilResponse
	destinations []Destination
}

func newRouter(in chan *ZkilResponse, out chan *RoutedZkilResponse, destinations []Destination, log *zap.SugaredLogger) *router {
	r := &router{in, out, destinations}
	r.init(log)
	return r
}

func (r *router) init(log *zap.SugaredLogger) {
	go func() {
		for {
			zkr := <-r.in
			log.Debugw("routing KM", "killID", zkr.Package.KillID)
			routed := r.routeMessage(zkr)
			r.out <- routed
		}
	}()
}

func (r *router) routeMessage(zkr *ZkilResponse) *RoutedZkilResponse {
	var matchedRoutes []Route
	for _, dest := range r.destinations {
		if dest.All {
			routes := destinationToRoutes(dest, false)
			matchedRoutes = append(matchedRoutes, routes...)
			continue
		}

		if matchedVictim(dest, zkr.Package.Killmail.Victim) {
			routes := destinationToRoutes(dest, true)
			matchedRoutes = append(matchedRoutes, routes...)
			continue
		}

		if matchedAttackers(dest, zkr.Package.Killmail.Attackers) {
			routes := destinationToRoutes(dest, false)
			matchedRoutes = append(matchedRoutes, routes...)
			continue
		}
	}

	return &RoutedZkilResponse{ZkilResponse: zkr, MatchedRoutes: matchedRoutes}
}

func matchedVictim(d Destination, victim goesi.GetKillmailsKillmailIdKillmailHashVictim) bool {
	return ContainsID(d.CharacterIDs, victim.CharacterId) ||
		ContainsID(d.CorporationIDs, victim.CorporationId) ||
		ContainsID(d.AllianceIDs, victim.AllianceId)
}

func matchedAttacker(d Destination, attacker goesi.GetKillmailsKillmailIdKillmailHashAttacker) bool {
	return ContainsID(d.CharacterIDs, attacker.CharacterId) ||
		ContainsID(d.CorporationIDs, attacker.CorporationId) ||
		ContainsID(d.AllianceIDs, attacker.AllianceId)
}

func matchedAttackers(d Destination, attackers goesi.GetKillmailsKillmailIdKillmailHashAttackerList) bool {
	for _, att := range attackers {
		if matchedAttacker(d, att) {
			return true
		}
	}
	return false
}

func destinationToRoutes(d Destination, isLoss bool) []Route {
	var routes []Route
	for _, dest := range d.DiscordWebhooks {
		routes = append(routes, Route{dest, isLoss})
	}
	return routes
}

func ContainsID(ids []int32, id int32) bool {
	for _, elem := range ids {
		if elem == id {
			return true
		}
	}

	return false
}
