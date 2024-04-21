package school

import (
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/database"
)

func Populate() (err error) {
	var schools = []model.School{
		{Name: "primaire 1"},
		{Name: "primaire 2"},
		{Name: "primaire 3"},
		{Name: "primaire 4"},
		{Name: "primaire 5"},
		{Name: "primaire 6"},
		{Name: "secondaire 1"},
		{Name: "secondaire 2"},
		{Name: "secondaire 3"},
		{Name: "secondaire 4"},
		{Name: "secondaire 5"},
		{Name: "Cégep"},
		{Name: "Université"},
	}

	var primary = []model.SchoolSubject{
		{Name: "Mathématiques"},
		{Name: "Français"},
		{Name: "Anglais"},
		{Name: "Science et technologie"},
		{Name: "Histoire / Géographie"},
		{Name: "Éthique  et culture religieuse"},
		{Name: "Culture et citoyenneté québécoise (CCQ )"},
	}

	var secondaireOne = []model.SchoolSubject{
		{Name: "Français"},
		{Name: "Anglais"},
		{Name: "Histoire / Géographie"},
		{Name: "Sciences et technologies"},
		{Name: "Monde contemporain"},
		{Name: "Education financière"},
		{Name: "Éthique et culture religieuse"},
		{Name: "Mathématiques"},
	}

	var secondaireThree = []model.SchoolSubject{
		{Name: "Sciences ST"},
		{Name: "Sciences ATS"},
	}

	var secondaireFour = []model.SchoolSubject{
		{Name: "Mathématiques CST"},
		{Name: "Mathématiques SN"},
		{Name: "Mathématiques TS"},
		{Name: "Sciences STE"},
		{Name: "Sciences SE"},
	}

	var secondaireFive = []model.SchoolSubject{
		{Name: "Mathématiques CST"},
		{Name: "Mathématiques SN"},
		{Name: "Mathématiques  TS"},
		{Name: "Chimie"},
		{Name: "Physique"},
	}

	var cegep = []model.SchoolSubject{
		{Name: "Mathémathiques"},
		{Name: "Chimie"},
		{Name: "Physique"},
		{Name: "Chimie"},
		{Name: "Biologie"},
	}

	var univ = []model.SchoolSubject{
		{Name: "Chimie"},
		{Name: "Biologie"},
		{Name: "Physique"},
	}

	for j := 0; j < len(schools); j++ {
		id, errF := database.Insert(schools[j])
		if errF != nil {
			return errF
		}

		switch schools[j].Name {
		case "primaire 1", "primaire 2", "primaire 3", "primaire 4", "primaire 5", "primaire 6":
			for i := 0; i < len(primary); i++ {
				primary[i].SchoolNumber = int(id)
				_, err = database.Insert(primary[i])
				if err != nil {
					return err
				}
			}
			break
		case "secondaire 1", "secondaire 2":
			for i := 0; i < len(secondaireOne); i++ {
				secondaireOne[i].SchoolNumber = int(id)
				_, err = database.Insert(secondaireOne[i])
				if err != nil {
					return err
				}
			}
			break
		case "secondaire 3":
			for i := 0; i < len(secondaireThree); i++ {
				secondaireThree[i].SchoolNumber = int(id)
				_, err = database.Insert(secondaireThree[i])
				if err != nil {
					return err
				}
			}
			break
		case "secondaire 4":
			for i := 0; i < len(secondaireFour); i++ {
				secondaireFour[i].SchoolNumber = int(id)
				_, err = database.Insert(secondaireFour[i])
				if err != nil {
					return err
				}
			}
			break
		case "secondaire 5":
			for i := 0; i < len(secondaireFive); i++ {
				secondaireFive[i].SchoolNumber = int(id)
				_, err = database.Insert(secondaireFive[i])
				if err != nil {
					return err
				}
			}
			break
		case "Cégep":
			for i := 0; i < len(cegep); i++ {
				cegep[i].SchoolNumber = int(id)
				_, err = database.Insert(cegep[i])
				if err != nil {
					return err
				}
			}
			break
		case "Université":
			for i := 0; i < len(univ); i++ {
				univ[i].SchoolNumber = int(id)
				_, err = database.Insert(univ[i])
				if err != nil {
					return err
				}
			}
			break
		default:
			break
		}
	}

	return err
}
