package services

import (
	"astrology-services/internal/domain/models"
	"fmt"
	"time"
)

type PlanetaryReport struct {
	name string
	dob  time.Time
	tob  time.Time
	pob  string
}

func (pr *PlanetaryReport) calculatePlanetPosition(planet string) int {
	dayOfYear := pr.dob.YearDay()
	basePosition := (dayOfYear % 12) + 1 // Mock calculation for the position in a sign (1 to 12)
	return basePosition
}

func (pr *PlanetaryReport) getSign(position int) string {
	signs := []string{
		"Aries", "Taurus", "Gemini", "Cancer", "Leo", "Virgo",
		"Libra", "Scorpio", "Sagittarius", "Capricorn", "Aquarius", "Pisces",
	}
	return signs[position-1]
}

func (pr *PlanetaryReport) sunConsideration() string {
	position := pr.calculatePlanetPosition("Sun")
	sign := pr.getSign(position)
	return fmt.Sprintf(`
Sun Consideration
The planet Sun is in the %dth house of your Kundli which is the %s sign. The presence of the Sun in this house could be an ominous omen of the great restlessness and hindrances that you may have to bear over the next few days. You might also face some complaints about a few disorders in the stomach, specifically below the navel. Therefore, you should take care of health and eating habits in the next few days in order to prevent this. Regulating a good workout routine and could also help prevent any sort of physical disorders.

On the upside, you will be successful in getting work done in the house of your foes with the help of diplomatic tactics. These are a result of your wisdom, cleverness and excellent presence of mind. Sometimes, you might be prone to face very severe calamities but you will rely on your strengths to see such times through.
`, position, sign)
}

func (pr *PlanetaryReport) moonConsideration() string {
	position := pr.calculatePlanetPosition("Moon")
	sign := pr.getSign(position)
	return fmt.Sprintf(`
Moon Consideration
The planet Moon is in the %dth house of your Kundli which is the %s sign. Owing to the conjunction of the Moon in this house, you might have to face a bit of weakness in education. You could also feel like you are lacking a bit of intellect in certain situations. You could also incur some loss in your destiny in the coming days. If you have children, this is the time when they need your guidance the most as there is a slight possibility that they need a little bit more attention.

On the other hand, you will toil consistently and your labours will bear fruit as you enjoy financial success over the coming days. You lead a well balanced life and have elegance in your daily routine. This will help you achieve a peaceful and healthy mind. You will also be blessed with domestic bliss as your family progresses in its bonds.
`, position, sign)
}

func (pr *PlanetaryReport) mercuryConsideration() string {
	position := pr.calculatePlanetPosition("Mercury")
	sign := pr.getSign(position)
	return fmt.Sprintf(`
Mercury Consideration
The planet Mercury is in the %dth house of your natal chart with the %s sign. Results may not be favorable for the native with this conjunction. You might feel loss regarding your father, have issues with siblings, and unfriendly terms with your mother. Your family won’t be supportive of you. Thus, conflict of thoughts may arise between them and you. When pursuing a career or business, you might have to bear a lot of trouble and problems to get a raise. Also, to attain a reputable position in society and government, you may go through several hindrances.
Naturewise, you will be steady and energetic, but there are chances that you might be aggressive and rude in speaking terms. With that, you could face delays and problems in owning desired luxuries. Disadvantages and loss may also occur if involved in debts, loans, or court cases.
`, position, sign)
}

func (pr *PlanetaryReport) venusConsideration() string {
	position := pr.calculatePlanetPosition("Venus")
	sign := pr.getSign(position)
	return fmt.Sprintf(`
Venus Consideration
The planet Venus is in the %dth house of your natal chart with the %s sign. The presence of the planet Venus in conjunction with %s suggests that the native has some weakness in the house of destiny. You can be extremely fortunate at times, but not always. Your work in your professional life will bring you some success. This will bring you honour from the government and society. However, you might also experience some weakness in your business and occupation.
In terms of your personal life and family, the conjunction of the planet Venus with %s, in the %dth house, suggests that you might not be very fortunate. You might have a difficult relationship with your father and experience some weakness from his side.
`, position, sign, sign, sign, position)
}

func (pr *PlanetaryReport) marsConsideration() string {
	position := pr.calculatePlanetPosition("Mars")
	sign := pr.getSign(position)
	return fmt.Sprintf(`
Mars Consideration
The planet Mars is in the %dth house of your natal chart with the %s sign. Mars in this house may make the native go through many unfavorable results in life. You might seek hindrances in family life. Bonding with mother won’t be friendly, and thus, misunderstandings between the two might persist.
Settlement in life could come with a gradual delay, and instability regarding other life matters can chase you. With that, obstacles in buying a new home, land, or vehicles might occur as well. Despite having a long life, you will deal with struggles, impatience, and restlessness in it. You might try to look for happiness from baffling means, but hurdles are likely to follow you in the same. Professional stability can become a tough task for you. An imbalance in money management might also be there, and thus, possessing your desired comforts in life won’t exist.
`, position, sign)
}

func (pr *PlanetaryReport) jupiterConsideration() string {
	position := pr.calculatePlanetPosition("Jupiter")
	sign := pr.getSign(position)
	return fmt.Sprintf(`
Jupiter Consideration
The planet Jupiter is in the %dth house of your Kundli with the %s sign. Owing the conjunction of the Jupiter in this house, there are chances that you may experience some sort of physical weakness in the coming days. This may be due to lack of strength and as you may have been engaged in a meticulous and demanding routine that has been sapping you of your energy. Therefore, it is advisable to take a good amount of rest and delegate some of the work so that you may allow your body to recover completely. In your personal life, you may have to incur some loss through your siblings.

However, as a patronee of the %s sign, you are not someone who sits idle after having suffered losses. You will instead use great effort and draw on your personal resources to recover the lost wealth. Apart from your financial troubles, you may tend to feel worries about your day-to-day activities. In such a case, it would do you good to try and change things up a bit by drawing out a new schedule.
`, position, sign, sign)
}

func (pr *PlanetaryReport) saturnConsideration() string {
	position := pr.calculatePlanetPosition("Saturn")
	sign := pr.getSign(position)
	return fmt.Sprintf(`
Saturn Consideration
The planet Saturn is in the %dth house of your natal chart with the %s sign. The presence of Saturn in the final house suggests that the native might be bad with money. You might develop the habit of making reckless and unnecessary expenditures and might also exhibit weakness in the house of destiny. Therefore it is better to keep grinding and not leave anything up to fate. You will be blessed with the support of labor which will grant you the ability to mold your destiny to better suit your preference.
Speaking of destiny, despite the likelihood of you developing the habit of reckless spending, you may end up with the power to manage expenditures. Moving forward, you might come across various hurdles if you are looking to acquire fame. There might also be some hurdles in your pursuit of happiness, along with some deficiencies in the day-to-day work in your life.
`, position, sign)
}

func (pr *PlanetaryReport) rahuConsideration() string {
	position := pr.calculatePlanetPosition("Rahu")
	sign := pr.getSign(position)
	return fmt.Sprintf(`
Rahu Consideration
The planet Rahu is in the %dth house of your natal chart with the %s sign. The presence of Rahu in conjunction with %s suggests that the native will be extremely wise and respectful. You will be someone who has an inherently influential side in you which will create leadership qualities. Your personality will go as far as to influence your enemies in their lives.
In terms of your family, having Rahu in the %dth house of your natal chart suggests that you will make some progress in rekindling your relationship with your father. If you both had fallen apart due to some reason over the years, this will give you the opportunity to make things right.
`, position, sign, sign, position)
}

func (pr *PlanetaryReport) ketuConsideration() string {
	position := pr.calculatePlanetPosition("Ketu")
	sign := pr.getSign(position)
	return fmt.Sprintf(`
Ketu Consideration
The planet Ketu is in the %dth house of your natal chart with the %s sign. %s is in a neighboring house and thus, will give native all fruitful results. In terms of family, you will attain strength from your siblings and support from your father. You might get involved in a few misunderstandings with your mother, but most things would be fine regarding your family life.
You will be of a hard-working personality, and this would help you in attaining fame in life. You would gain honor from society and government, be successful if involved in business, and diligent in attitude. You would be courageous enough to take a stand regarding things in your life and will focus on work. There are also chances that you involve yourself in religious deeds and benefit from the same.
`, position, sign, sign)
}

func (pr *PlanetaryReport) generateReport() *models.PlanetaryReport {
	return &models.PlanetaryReport{
		SunConsideration:     pr.sunConsideration(),
		MoonConsideration:    pr.moonConsideration(),
		MercuryConsideration: pr.mercuryConsideration(),
		VenusConsideration:   pr.venusConsideration(),
		MarsConsideration:    pr.marsConsideration(),
		JupiterConsideration: pr.jupiterConsideration(),
		SaturnConsideration:  pr.saturnConsideration(),
		RahuConsideration:    pr.rahuConsideration(),
		KetuConsideration:    pr.ketuConsideration(),
	}
}

func (s *AstrologyReportService) GeneratePlanetaryReport(name, dob, tob, placeOfBirth string) (*models.PlanetaryReport, error) {
	dobTime, err := time.Parse("2006-01-02", dob)
	if err != nil {
		return nil, err
	}
	tobTime, err := time.Parse("15:04", tob)
	if err != nil {
		return nil, err
	}

	report := &PlanetaryReport{
		name: name,
		dob:  dobTime,
		tob:  tobTime,
		pob:  placeOfBirth,
	}

	return report.generateReport(), nil
}
