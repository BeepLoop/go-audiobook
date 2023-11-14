package store

import "github.com/audiobook/types"

func GetStoryData(storyId string) *types.Story {
	var data types.Story

	for _, story := range Stories.Stories {
		if story.Id == storyId {
			data = story
			break
		}
	}

	return &data
}
