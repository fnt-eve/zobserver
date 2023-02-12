package observer

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/antihax/goesi"
	"github.com/antihax/goesi/esi"
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type EntityAffiliation struct {
	CharacterID   int32
	CorporationID int32
	AllianceID    int32
	FactionID     int32
	ShipTypeID    int32
}

type EntityNames struct {
	CharacterName   string
	CorporationName string
	AllianceName    string
	ShipTypeName    string
	FactionName     string
}

type resolvedNames struct {
	systemName string
	victimInfo EntityNames
	fbInfo     *EntityNames
}

type sender struct {
	dg        *discordgo.Session
	APIClient *goesi.APIClient
	in        chan *RoutedZkilResponse
}

func newSender(in chan *RoutedZkilResponse, ESIClient *goesi.APIClient, log *zap.SugaredLogger) *sender {
	dg, _ := discordgo.New("")
	s := &sender{dg, ESIClient, in}
	s.init(log)
	return s
}

func (s *sender) init(log *zap.SugaredLogger) {
	go func() {
		for {
			rm := <-s.in
			l := log.With("killID", rm.ZkilResponse.Package.KillID)
			log.Debugw("sending KM", "killID", rm.ZkilResponse.Package.KillID)
			for _, route := range rm.MatchedRoutes {
				ll := l.With("discordWHID", route.DiscordWebhook.ID)
				embed, err := s.transform(rm.ZkilResponse, route.IsLoss)
				if err != nil {
					ll.Errorw("failed to transform KM", "error", err)
					continue
				}
				msg := discordgo.WebhookParams{Content: "", Username: "Killmail Feed", Embeds: []*discordgo.MessageEmbed{embed}}
				_, err = s.dg.WebhookExecute(route.DiscordWebhook.ID, route.DiscordWebhook.Token, false, &msg)
				if err != nil {
					ll.Errorw("failed to send KM to route", "error", err)
				}
			}
		}
	}()
}

func (s *sender) resolveIds(ids []int32) (map[int32]string, error) {
	// Make sure all ids are unique
	u := make([]int32, 0, len(ids))
	m := make(map[int32]bool)

	for _, val := range ids {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}
	resp, _, err := s.APIClient.ESI.UniverseApi.PostUniverseNames(context.Background(), u, nil)
	if err != nil {
		return nil, err
	}
	ret := make(map[int32]string)
	for _, elem := range resp {
		ret[elem.Id] = elem.Name
	}

	return ret, nil
}

func (s *sender) resolveKMEntities(victimAffiliation, fbAffiliation *EntityAffiliation, systemID int32) (*resolvedNames, error) {
	ids := []int32{
		systemID,
		victimAffiliation.ShipTypeID,
	}

	if victimAffiliation.CharacterID != 0 {
		ids = append(ids, victimAffiliation.CharacterID)
	}

	if victimAffiliation.CorporationID != 0 {
		ids = append(ids, victimAffiliation.CorporationID)
	}

	if victimAffiliation.AllianceID != 0 {
		ids = append(ids, victimAffiliation.AllianceID)
	}

	if victimAffiliation.FactionID != 0 {
		ids = append(ids, victimAffiliation.FactionID)
	}

	if fbAffiliation != nil {
		if fbAffiliation.CharacterID != 0 {
			ids = append(ids, fbAffiliation.CharacterID)
		}
		if fbAffiliation.CorporationID != 0 {
			ids = append(ids, fbAffiliation.CorporationID)
		}
		if fbAffiliation.AllianceID != 0 {
			ids = append(ids, fbAffiliation.AllianceID)
		}
		if fbAffiliation.FactionID != 0 {
			ids = append(ids, fbAffiliation.FactionID)
		}
		ids = append(ids, fbAffiliation.ShipTypeID)
	}

	resolved, err := s.resolveIds(ids)
	if err != nil {
		return nil, err
	}

	rn := resolvedNames{}

	rn.victimInfo = getKMEntityInfo(*victimAffiliation, resolved)
	if fbAffiliation != nil {
		entity := getKMEntityInfo(*fbAffiliation, resolved)
		rn.fbInfo = &entity
	}

	if name, ok := resolved[systemID]; ok {
		rn.systemName = name
	}

	return &rn, nil
}

func getKMEntityInfo(aff EntityAffiliation, resolvedNames map[int32]string) EntityNames {
	names := EntityNames{}

	if name, ok := resolvedNames[aff.CharacterID]; ok {
		names.CharacterName = name
	}

	if name, ok := resolvedNames[aff.CorporationID]; ok {
		names.CorporationName = name
	}

	if name, ok := resolvedNames[aff.AllianceID]; ok {
		names.AllianceName = name
	}

	if name, ok := resolvedNames[aff.ShipTypeID]; ok {
		names.ShipTypeName = name
	}

	if name, ok := resolvedNames[aff.FactionID]; ok {
		names.FactionName = name
	}

	return names
}

func (s *sender) transform(r *ZkilResponse, isLoss bool) (*discordgo.MessageEmbed, error) {
	victimAff := EntityAffiliation{
		CharacterID:   r.Package.Killmail.Victim.CharacterId,
		CorporationID: r.Package.Killmail.Victim.CorporationId,
		AllianceID:    r.Package.Killmail.Victim.AllianceId,
		ShipTypeID:    r.Package.Killmail.Victim.ShipTypeId,
		FactionID:     r.Package.Killmail.Victim.FactionId,
	}
	fb := findFinalBlow(r.Package.Killmail.Attackers)
	var fbAff *EntityAffiliation
	if fb != nil {
		fbAff = &EntityAffiliation{
			CharacterID:   fb.CharacterId,
			CorporationID: fb.CorporationId,
			AllianceID:    fb.AllianceId,
			ShipTypeID:    fb.ShipTypeId,
			FactionID:     fb.FactionId,
		}
	}

	names, err := s.resolveKMEntities(&victimAff, fbAff, r.Package.Killmail.SolarSystemId)
	if err != nil {
		return nil, err
	}

	var footer *discordgo.MessageEmbedFooter
	if fb != nil {
		footer = genFooter(*names.fbInfo, fb.ShipTypeId, len(r.Package.Killmail.Attackers))
	}

	color, err := getColor(isLoss)
	if err != nil {
		return nil, err
	}

	return &discordgo.MessageEmbed{
		URL:       zKillURL(r.Package.KillID),
		Title:     genTitle(names.victimInfo.CharacterName, names.victimInfo.ShipTypeName, names.systemName),
		Timestamp: r.Package.Killmail.KillmailTime.Format(time.RFC3339),
		Fields:    genFields(names.victimInfo, r.Package.Killmail.Victim.ShipTypeId, names.fbInfo.ShipTypeName, names.systemName, r.Package.ZKillmailMetadata.TotalValue),
		Thumbnail: genThumbnail(r.Package.Killmail.Victim.ShipTypeId),
		Color:     color,
		Footer:    footer,
	}, nil
}

func getColor(isLoss bool) (int, error) {
	// These are hex values, ie 0x...
	red := "990000"
	green := "009900"

	str := green
	if isLoss {
		str = red
	}

	i, err := strconv.ParseInt(str, 16, 0)
	return int(i), err
}

func genTitle(charName, shipTypeName, systemName string) string {
	return fmt.Sprintf("%s | %s | %s", charName, shipTypeName, systemName)
}

func genThumbnail(shipTypeID int32) *discordgo.MessageEmbedThumbnail {
	return &discordgo.MessageEmbedThumbnail{
		URL: itemTypeImageURL(shipTypeID, false, 256),
	}
}

func genFields(victimStrings EntityNames, shipTypeID int32, shipTypeName, systemName string, totalValue float64) []*discordgo.MessageEmbedField {
	fields := []*discordgo.MessageEmbedField{}
	if victimStrings.CharacterName != "" {
		fields = append(fields,
			&discordgo.MessageEmbedField{
				Name:   "Character",
				Value:  victimStrings.CharacterName,
				Inline: true,
			})
	}
	fields = append(fields,
		&discordgo.MessageEmbedField{
			Name:   "Corporation",
			Value:  victimStrings.CorporationName,
			Inline: true,
		})
	if victimStrings.AllianceName != "" {
		fields = append(fields,
			&discordgo.MessageEmbedField{
				Name:   "Alliance",
				Value:  victimStrings.AllianceName,
				Inline: true,
			})
	}

	valuePrinter := message.NewPrinter(language.English)

	fields = append(fields, []*discordgo.MessageEmbedField{
		{
			Name:   "Ship",
			Value:  shipTypeName,
			Inline: true,
		},
		{
			Name:   "Location",
			Value:  systemName,
			Inline: true,
		},
		{
			Name:   "Total Value",
			Value:  valuePrinter.Sprintf("%.2f ISK", totalValue),
			Inline: true,
		},
	}...)

	return fields
}

func genFooter(entityInfo EntityNames, shipTypeID int32, involvedCount int) *discordgo.MessageEmbedFooter {
	entityName := entityInfo.CorporationName
	if entityInfo.AllianceName != "" {
		entityName = entityInfo.AllianceName
	}

	var footText string
	if entityInfo.CharacterName == "" && entityName != "" {
		footText = fmt.Sprintf("Final blow by %s in %s (%v involved)", entityName, entityInfo.ShipTypeName, involvedCount)
	} else if entityInfo.CharacterName == "" && entityName == "" {
		footText = fmt.Sprintf("Final blow by %s in %s (%v involved)", entityInfo.FactionName, entityInfo.ShipTypeName, involvedCount)
	} else {
		footText = fmt.Sprintf("Final blow by %s (%s) in %s (%v involved)", entityInfo.CharacterName, entityName, entityInfo.ShipTypeName, involvedCount)
	}
	return &discordgo.MessageEmbedFooter{
		Text:    footText,
		IconURL: itemTypeImageURL(shipTypeID, true, 32),
	}
}

func zKillURL(killID int64) string {
	return fmt.Sprintf("https://zkillboard.com/kill/%v", killID)
}

func itemTypeImageURL(ID int32, isIcon bool, size int) string {
	imgType := "render"
	if isIcon {
		imgType = "icon"
	}
	URL := fmt.Sprintf("https://images.evetech.net/types/%v/%s", ID, imgType)

	if size != 0 {
		URL += fmt.Sprintf("?size=%v", size)
	}

	return URL
}

func findFinalBlow(attackers esi.GetKillmailsKillmailIdKillmailHashAttackerList) *esi.GetKillmailsKillmailIdKillmailHashAttacker {
	for _, a := range attackers {
		if a.FinalBlow {
			return &a
		}
	}
	return nil
}
