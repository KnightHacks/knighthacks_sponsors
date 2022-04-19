// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type NewSponsor struct {
	Name        string           `json:"name"`
	Tier        SubscriptionTier `json:"tier"`
	Since       *time.Time       `json:"since"`
	Description *string          `json:"description"`
	Website     *string          `json:"website"`
	Logo        *string          `json:"logo"`
}

type Sponsor struct {
	ID          string           `json:"id"`
	Name        string           `json:"name"`
	Tier        SubscriptionTier `json:"tier"`
	Since       time.Time        `json:"since"`
	Description *string          `json:"description"`
	Website     *string          `json:"website"`
	Logo        *string          `json:"logo"`
}

type SponsorFilter struct {
	Tiers []SubscriptionTier `json:"tiers"`
}

type UpdatedSponsor struct {
	Name        *string           `json:"name"`
	Tier        *SubscriptionTier `json:"tier"`
	Since       *time.Time        `json:"since"`
	Description *string           `json:"description"`
	Website     *string           `json:"website"`
	Logo        *string           `json:"logo"`
}

type SubscriptionTier string

const (
	SubscriptionTierBronze   SubscriptionTier = "BRONZE"
	SubscriptionTierSilver   SubscriptionTier = "SILVER"
	SubscriptionTierGold     SubscriptionTier = "GOLD"
	SubscriptionTierPlatinum SubscriptionTier = "PLATINUM"
)

var AllSubscriptionTier = []SubscriptionTier{
	SubscriptionTierBronze,
	SubscriptionTierSilver,
	SubscriptionTierGold,
	SubscriptionTierPlatinum,
}

func (e SubscriptionTier) IsValid() bool {
	switch e {
	case SubscriptionTierBronze, SubscriptionTierSilver, SubscriptionTierGold, SubscriptionTierPlatinum:
		return true
	}
	return false
}

func (e SubscriptionTier) String() string {
	return string(e)
}

func (e *SubscriptionTier) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SubscriptionTier(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SubscriptionTier", str)
	}
	return nil
}

func (e SubscriptionTier) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
