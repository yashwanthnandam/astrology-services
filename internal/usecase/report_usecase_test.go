package usecase_test

import (
	"astrology-services/internal/domain/models"
	"astrology-services/internal/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockReportRepository struct {
	mock.Mock
}

func (m *MockReportRepository) GetGeneralReportByID(id uint) (*models.GeneralReport, error) {
	args := m.Called(id)
	return args.Get(0).(*models.GeneralReport), args.Error(1)
}

func (m *MockReportRepository) GetPlanetaryReportByID(id uint) (*models.PlanetaryReport, error) {
	args := m.Called(id)
	return args.Get(0).(*models.PlanetaryReport), args.Error(1)
}

func (m *MockReportRepository) GetVimshottariReportByID(id uint) (*models.VimshottariReport, error) {
	args := m.Called(id)
	return args.Get(0).(*models.VimshottariReport), args.Error(1)
}

func (m *MockReportRepository) GetYogaReportByID(id uint) (*models.YogaReport, error) {
	args := m.Called(id)
	return args.Get(0).(*models.YogaReport), args.Error(1)
}

type MockReportService struct {
	mock.Mock
}

func (m *MockReportService) GeneratePlanetaryReport(name, dob, tob, placeOfBirth string) (*models.PlanetaryReport, error) {
	args := m.Called(name, dob, tob, placeOfBirth)
	return args.Get(0).(*models.PlanetaryReport), args.Error(1)
}

func (m *MockReportService) GenerateVimshottariReport(name, dob, tob, placeOfBirth string) (*models.VimshottariReport, error) {
	args := m.Called(name, dob, tob, placeOfBirth)
	return args.Get(0).(*models.VimshottariReport), args.Error(1)
}

func (m *MockReportService) GenerateYogaReport(name, dob, tob, placeOfBirth string) (*models.YogaReport, error) {
	args := m.Called(name, dob, tob, placeOfBirth)
	return args.Get(0).(*models.YogaReport), args.Error(1)
}

func TestGetGeneralReport(t *testing.T) {
	mockRepo := new(MockReportRepository)
	mockService := new(MockReportService)

	report := &models.GeneralReport{
		Description:  "This is a description",
		Personality:  "This is a personality",
		Physical:     "This is physical",
		Health:       "This is health",
		Career:       "This is career",
		Relationship: "This is relationship",
	}

	mockRepo.On("GetGeneralReportByID", uint(1)).Return(report, nil)

	reportUsecase := usecase.NewReportUsecase(mockRepo, mockService)
	result, err := reportUsecase.GetGeneralReport(1)

	assert.NoError(t, err)
	assert.Equal(t, report, result)

	mockRepo.AssertExpectations(t)
}

func TestGeneratePlanetaryReport(t *testing.T) {
	mockRepo := new(MockReportRepository)
	mockService := new(MockReportService)

	report := &models.PlanetaryReport{
		SunConsideration:     "Sun Consideration The planet Sun is in the 16th house of your Kundli which is the Aries sign. The presence of the Sun in this house could be an ominous omen of the great restlessness and hindrances that you may have to bear over the next few days. You might also face some complaints about a few disorders in the stomach, specifically below the navel. Therefore, you should take care of health and eating habits in the next few days in order to prevent this. Regulating a good workout routine and could also help prevent any sort of physical disorders. On the upside, you will be successful in getting work done in the house of your foes with the help of diplomatic tactics. These are a result of your wisdom, cleverness and excellent presence of mind. Sometimes, you might be prone to face very severe calamities but you will rely on your strengths to see such times through.",
		MoonConsideration:    "Moon Consideration The planet Moon is in the 16th house of your Kundli which is the Aries sign. Owing to the conjunction of the Moon in this house, you might have to face a bit of weakness in education. You could also feel like you are lacking a bit of intellect in certain situations. You could also incur some loss in your destiny in the coming days. If you have children, this is the time when they need your guidance the most as there is a slight possibility that they need a little bit more attention. On the other hand, you will toil consistently and your labours will bear fruit as you enjoy financial success over the coming days. You lead a well balanced life and have elegance in your daily routine. This will help you achieve a peaceful and healthy mind. You will also be blessed with domestic bliss as your family progresses in its bonds.",
		MercuryConsideration: "Mercury Consideration The planet Mercury is in the 16th house of your natal chart with the Aries sign. Results may not be favorable for the native with this conjunction. You might feel loss regarding your father, have issues with siblings, and unfriendly terms with your mother. Your family won’t be supportive of you. Thus, conflict of thoughts may arise between them and you. When pursuing a career or business, you might have to bear a lot of trouble and problems to get a raise. Also, to attain a reputable position in society and government, you may go through several hindrances. Naturewise, you will be steady and energetic, but there are chances that you might be aggressive and rude in speaking terms. With that, you could face delays and problems in owning desired luxuries. Disadvantages and loss may also occur if involved in debts, loans, or court cases.",
		VenusConsideration:   "Venus Consideration The planet Venus is in the 16th house of your natal chart with the Aries sign. The presence of the planet Venus in conjunction with Aries suggests that the native has some weakness in the house of destiny. You can be extremely fortunate at times, but not always. Your work in your professional life will bring you some success. This will bring you honour from the government and society. However, you might also experience some weakness in your business and occupation. In terms of your personal life and family, the conjunction of the planet Venus with Aries, in the 16th house, suggests that you might not be very fortunate. You might have a difficult relationship with your father and experience some weakness from his side.",
		MarsConsideration:    "Mars Consideration The planet Mars is in the 16th house of your natal chart with the Aries sign. Mars in this house may make the native go through many unfavorable results in life. You might seek hindrances in family life. Bonding with mother won’t be friendly, and thus, misunderstandings between the two might persist. Settlement in life could come with a gradual delay, and instability regarding other life matters can chase you. With that, obstacles in buying a new home, land, or vehicles might occur as well. Despite having a long life, you will deal with struggles, impatience, and restlessness in it. You might try to look for happiness from baffling means, but hurdles are likely to follow you in the same. Professional stability can become a tough task for you. An imbalance in money management might also be there, and thus, possessing your desired comforts in life won’t exist.",
		JupiterConsideration: "Jupiter Consideration The planet Jupiter is in the 16th house of your Kundli with the Aries sign. Owing the conjunction of the Jupiter in this house, there are chances that you may experience some sort of physical weakness in the coming days. This may be due to lack of strength and as you may have been engaged in a meticulous and demanding routine that has been sapping you of your energy. Therefore, it is advisable to take a good amount of rest and delegate some of the work so that you may allow your body to recover completely. In your personal life, you may have to incur some loss through your siblings. However, as a patronee of the Aries sign, you are not someone who sits idle after having suffered losses. You will instead use great effort and draw on your personal resources to recover the lost wealth. Apart from your financial troubles, you may tend to feel worries about your day-to-day activities. In such a case, it would do you good to try and change things up a bit by drawing out a new schedule.",
		SaturnConsideration:  "Saturn Consideration The planet Saturn is in the 16th house of your natal chart with the Aries sign. The presence of Saturn in the final house suggests that the native might be bad with money. You might develop the habit of making reckless and unnecessary expenditures and might also exhibit weakness in the house of destiny. Therefore it is better to keep grinding and not leave anything up to fate. You will be blessed with the support of labor which will grant you the ability to mold your destiny to better suit your preference. Speaking of destiny, despite the likelihood of you developing the habit of reckless spending, you may end up with the power to manage expenditures. Moving forward, you might come across various hurdles if you are looking to acquire fame. There might also be some hurdles in your pursuit of happiness, along with some deficiencies in the day-to-day work in your life.",
		RahuConsideration:    "Rahu Consideration The planet Rahu is in the 16th house of your natal chart with the Aries sign. The presence of Rahu in conjunction with Aries suggests that the native will be extremely wise and respectful. You will be someone who has an inherently influential side in you which will create leadership qualities. Your personality will go as far as to influence your enemies in their lives. In terms of your family, having Rahu in the 16th house of your natal chart suggests that you will make some progress in rekindling your relationship with your father. If you both had fallen apart due to some reason over the years, this will give you the opportunity to make things right.",
		KetuConsideration:    "Ketu Consideration The planet Ketu is in the 16th house of your natal chart with the Aries sign. Aries is in a neighboring house and thus, will give native all fruitful results. In terms of family, you will attain strength from your siblings and support from your father. You might get involved in a few misunderstandings with your mother, but most things would be fine regarding your family life. You will be of a hard-working personality, and this would help you in attaining fame in life. You would gain honor from society and government, be successful if involved in business, and diligent in attitude. You would be courageous enough to take a stand regarding things in your life and will focus on work. There are also chances that you involve yourself in religious deeds and benefit from the same.",
	}

	mockService.On("GeneratePlanetaryReport", "John Doe", "1990-01-01", "12:00", "New York").Return(report, nil)

	reportUsecase := usecase.NewReportUsecase(mockRepo, mockService)
	result, err := reportUsecase.GeneratePlanetaryReport("John Doe", "1990-01-01", "12:00", "New York")

	assert.NoError(t, err)
	assert.Equal(t, report, result)

	mockService.AssertExpectations(t)
}

func TestGenerateVimshottariReport(t *testing.T) {
	mockRepo := new(MockReportRepository)
	mockService := new(MockReportService)

	report := &models.VimshottariReport{
		MercuryMahadasha: "Mercury Mahadasha is good",
		MoonMahadasha:    "Moon Mahadasha is good",
		VenusMahadasha:   "Venus Mahadasha is good",
		MarsMahadasha:    "Mars Mahadasha is good",
		JupiterMahadasha: "Jupiter Mahadasha is good",
		SaturnMahadasha:  "Saturn Mahadasha is good",
		RahuMahadasha:    "Rahu Mahadasha is good",
		KetuMahadasha:    "Ketu Mahadasha is good",
	}

	mockService.On("GenerateVimshottariReport", "John Doe", "1990-01-01", "12:00", "New York").Return(report, nil)

	reportUsecase := usecase.NewReportUsecase(mockRepo, mockService)
	result, err := reportUsecase.GenerateVimshottariReport("John Doe", "1990-01-01", "12:00", "New York")

	assert.NoError(t, err)
	assert.Equal(t, report, result)

	mockService.AssertExpectations(t)
}

func TestGenerateYogaReport(t *testing.T) {
	mockRepo := new(MockReportRepository)
	mockService := new(MockReportService)

	report := &models.YogaReport{
		GajakesariYoga:  "Gajakesari Yoga is present",
		SunaphaYoga:     "Sunapha Yoga is present",
		AdhiYoga:        "Adhi Yoga is present",
		BudhaAdityaYoga: "Budha-Aditya Yoga is present",
	}

	mockService.On("GenerateYogaReport", "John Doe", "1990-01-01", "12:00", "New York").Return(report, nil)

	reportUsecase := usecase.NewReportUsecase(mockRepo, mockService)
	result, err := reportUsecase.GenerateYogaReport("John Doe", "1990-01-01", "12:00", "New York")

	assert.NoError(t, err)
	assert.Equal(t, report, result)

	mockService.AssertExpectations(t)
}
