package observer

import (
	"context"
	"slices"

	goesi "github.com/antihax/goesi/esi"
	"go.uber.org/zap"
)

type router struct {
	in           chan *Killmail
	out          chan *RoutedKillmail
	destinations []Destination
	log          *zap.SugaredLogger
}

func newRouter(in chan *Killmail, out chan *RoutedKillmail, destinations []Destination, log *zap.SugaredLogger) *router {
	return &router{in, out, destinations, log}
}

// run routes incoming killmails to matching destinations until ctx is
// cancelled.
func (r *router) run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case km := <-r.in:
			r.log.Debugw("routing KM", "killmail_id", km.KillmailID)
			routed := r.routeMessage(km)
			select {
			case r.out <- routed:
			case <-ctx.Done():
				return
			}
		}
	}
}

func (r *router) routeMessage(km *Killmail) *RoutedKillmail {
	var matchedRoutes []Route
	for _, dest := range r.destinations {
		if dest.All {
			routes := destinationToRoutes(dest, false)
			matchedRoutes = append(matchedRoutes, routes...)
			continue
		}

		if matchedVictim(dest, km.ESI.Victim) {
			routes := destinationToRoutes(dest, true)
			matchedRoutes = append(matchedRoutes, routes...)
			continue
		}

		if matchedAttackers(dest, km.ESI.Attackers) {
			routes := destinationToRoutes(dest, false)
			matchedRoutes = append(matchedRoutes, routes...)
			continue
		}
	}

	return &RoutedKillmail{Killmail: km, MatchedRoutes: matchedRoutes}
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
	return slices.Contains(ids, id)
}
