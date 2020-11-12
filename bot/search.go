package bot

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/davidleitw/leetcodeBOT/model"
)

func Search(command []string) (*discordgo.MessageEmbed, error) {
	if len(command) == 1 {
		return nil, FORMAT_ERROR
	}

	if command[1] != "-d" && len(command) == 2 {
		if command[1] == "c8763" {
			return nil, KIRITO_ERROR
		}

		problemID, err := strconv.Atoi(command[1])
		if err != nil {
			return nil, SEARCH_NUMBER_ERROR
		}

		if problemID == 48763 {
			return nil, KIRITO_ERROR
		}

		problem, err := model.SearchWithProblemID(problemID)
		if err != nil {
			return nil, SEARCH_NOT_FOUND_ERROR
		}

		msgs := SearchProblemsMsg(problem)
		return msgs, nil

	} else if (command[1] == "-d" || command[1] == "--detail") && len(command) == 3 {
		problemID, err := strconv.Atoi(command[2])
		if err != nil {
			return nil, SEARCH_NUMBER_ERROR
		}

		problem, err := model.SearchWithProblemID(problemID)
		if err != nil {
			return nil, SEARCH_NOT_FOUND_ERROR
		}
		msgs := SearchProblemsDetailMsg(problem)
		return msgs, nil
	}

	return nil, FORMAT_ERROR
}
