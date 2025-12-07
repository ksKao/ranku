package anilist

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type graphQLRequest struct {
	Query string `json:"query"`
}

type graphQLResponse struct {
	Data anilistData `json:"data"`
}

type anilistData struct {
	Page page `json:"Page"`
}

type page struct {
	Media []media `json:"media"`
}

type media struct {
	Id         int32     `json:"id"`
	Title      title     `json:"title"`
	Characters character `json:"characters"`
}

type title struct {
	Romaji string `json:"romaji"`
}

type character struct {
	Nodes []characterNode `json:"nodes"`
}

type characterNode struct {
	Id          int32          `json:"id"`
	Image       characterImage `json:"image"`
	Name        characterName  `json:"name"`
	DateOfBirth dateOfBirth    `json:"dateOfBirth"`
	BloodType   *string        `json:"bloodType"`
	Age         *string        `json:"age"` // is string because some values are like "13-14"
	Description *string        `json:"description"`
	Gender      *string        `json:"gender"`
}

type characterImage struct {
	Medium string `json:"medium"`
}

type characterName struct {
	Full string `json:"full"`
}

type dateOfBirth struct {
	Year  *int32 `json:"year"`
	Month *int32 `json:"month"`
	Day   *int32 `json:"day"`
}

func GetAnilistTopAnimeWithCharacters(page int) (graphQLResponse, error) {
	reqBody := graphQLRequest{
		Query: fmt.Sprintf(`
		query Query {
		  Page(page: %d, perPage: 50) {
			media(sort: SCORE_DESC, type: ANIME) {
			  id
			  title {
				romaji
			  }
			  characters {
				nodes {
				  id
				  image {
					medium
				  }
				  name {
					full
				  }
				  dateOfBirth {
					year
					month
					day
				  }
				  bloodType
				  age
				  description
				  gender
				}
			  }
			}
		  }
		}
		`, page),
	}

	jsonBody, err := json.Marshal(reqBody)

	if err != nil {
		return graphQLResponse{}, err
	}

	resp, err := http.Post("https://graphql.anilist.co", "application/json", bytes.NewBuffer([]byte(jsonBody)))

	if err != nil {
		return graphQLResponse{}, nil
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return graphQLResponse{}, fmt.Errorf("anilist API returned non-200 status code (%d)", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return graphQLResponse{}, nil
	}

	var data graphQLResponse
	err = json.Unmarshal(body, &data)

	if err != nil {
		return graphQLResponse{}, nil
	}

	return data, nil
}
