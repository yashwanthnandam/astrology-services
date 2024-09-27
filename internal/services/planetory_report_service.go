package services

import (
	"astrology-services/internal/domain/models"
	"fmt"
	"time"
)

type PlanetaryReport struct {
	Name string
	Dob  time.Time
	Tob  time.Time
	Pob  string
}

// Struct to represent each planetary consideration
type PlanetConsideration struct {
	PlanetName string
	House      int
	Sign       string
	Message    string
}

// Helper method to calculate a planet's position (this is still a mock)
func (pr *PlanetaryReport) calculatePlanetPosition(planet string) int {
	dayOfYear := pr.Dob.YearDay()
	basePosition := (dayOfYear % 12) + 1 // Mock calculation for the position in a sign (1 to 12)
	return basePosition
}

// Helper method to get the zodiac sign from a position
func (pr *PlanetaryReport) getSign(position int) string {
	signs := []string{
		"Aries", "Taurus", "Gemini", "Cancer", "Leo", "Virgo",
		"Libra", "Scorpio", "Sagittarius", "Capricorn", "Aquarius", "Pisces",
	}
	if position > 0 && position <= len(signs) {
		return signs[position-1]
	}
	return "Unknown"
}

// A general method for generating planet considerations
func (pr *PlanetaryReport) generateConsideration(planet string, template string) string {
	position := pr.calculatePlanetPosition(planet)
	sign := pr.getSign(position)
	return fmt.Sprintf(template, position, sign)
}

// Templates for each planet
var planetTemplates = map[string]string{
	"Sun": `
Sun Consideration
The planet Sun is in the %dth house of your Kundli which is the %s sign...`,
	"Moon": `
Moon Consideration
The planet Moon is in the %dth house of your Kundli which is the %s sign...`,
	"Mercury": `
Mercury Consideration
The planet Mercury is in the %dth house of your natal chart with the %s sign...`,
	"Venus": `
Venus Consideration
The planet Venus is in the %dth house of your natal chart with the %s sign...`,
	"Mars": `
Mars Consideration
The planet Mars is in the %dth house of your natal chart with the %s sign...`,
	"Jupiter": `
Jupiter Consideration
The planet Jupiter is in the %dth house of your Kundli with the %s sign...`,
	"Saturn": `
Saturn Consideration
The planet Saturn is in the %dth house of your natal chart with the %s sign...`,
	"Rahu": `
Rahu Consideration
The planet Rahu is in the %dth house of your natal chart with the %s sign...`,
	"Ketu": `
Ketu Consideration
The planet Ketu is in the %dth house of your natal chart with the %s sign...`,
}

// Method to generate a single planet's consideration dynamically
func (pr *PlanetaryReport) generatePlanetaryConsideration(planet string) string {
	template, exists := planetTemplates[planet]
	if !exists {
		return fmt.Sprintf("No consideration available for %s", planet)
	}
	return pr.generateConsideration(planet, template)
}

// Generates the full planetary report
func (pr *PlanetaryReport) generateReport() *models.PlanetaryReport {
	return &models.PlanetaryReport{
		SunConsideration:     pr.generatePlanetaryConsideration("Sun"),
		MoonConsideration:    pr.generatePlanetaryConsideration("Moon"),
		MercuryConsideration: pr.generatePlanetaryConsideration("Mercury"),
		VenusConsideration:   pr.generatePlanetaryConsideration("Venus"),
		MarsConsideration:    pr.generatePlanetaryConsideration("Mars"),
		JupiterConsideration: pr.generatePlanetaryConsideration("Jupiter"),
		SaturnConsideration:  pr.generatePlanetaryConsideration("Saturn"),
		RahuConsideration:    pr.generatePlanetaryConsideration("Rahu"),
		KetuConsideration:    pr.generatePlanetaryConsideration("Ketu"),
	}
}
