package hackernews

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

const baseUrl = "https://hacker-news.firebaseio.com/v0/"
const suffix = ".json"

func GetItem(id int) (*Item, error) {
	data, err := query(fmt.Sprintf("item/%d", id))
	if err != nil {
		return nil, err
	}

	var item *Item
	err = json.Unmarshal(data, &item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func GetUser(id string) (*User, error) {
	data, err := query(fmt.Sprintf("user/%s", id))
	if err != nil {
		return nil, err
	}

	var user *User
	err = json.Unmarshal(data, &user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetTopStories(n int) ([]*Item, error) {
	return queryAll("topstories", n)
}

func GetNewStories(n int) ([]*Item, error) {
	return queryAll("newstories", n)
}

func GetBestStories(n int) ([]*Item, error) {
	return queryAll("beststories", n)
}

func queryAll(path string, n int) ([]*Item, error) {
	data, err := query(path)
	if err != nil {
		return nil, err
	}

	var topIDs []int
	err = json.Unmarshal(data, &topIDs)
	if err != nil {
		return nil, err
	}

	if n > 0 {
		topIDs = topIDs[:n]
	}

	items := make([]*Item, len(topIDs))

	var wg sync.WaitGroup
	wg.Add(len(items))

	for i := range items {
		go func(idx int) {
			item, err := GetItem(topIDs[idx])
			if err != nil {
				items[idx] = &Item{ID: topIDs[idx], Title: "Error: " + err.Error()}
			} else {
				items[idx] = item
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
	return items, nil
}

func query(path string) ([]byte, error) {
	res, err := http.Get(baseUrl + path + suffix)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf(http.StatusText(http.StatusNotFound))
	}

	return body, nil
}
