package ktalk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func (c *Client) GetAllUsers(ctx context.Context) (*Users, error) {
	offset := ""
	users := Users{}
	for {
		url := fmt.Sprintf("%s/users/scan?top=1000&offset=%s&includeDisabled=true", c.BaseURL, offset)
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return nil, err
		}

		req = req.WithContext(ctx)
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
		req.Header.Set("Accept", "application/json; charset=utf-8")
		req.Header.Set("X-Auth-Token", c.token)

		res := Users{}
		if err = c.sendRequest(req, &res); err != nil {
			return nil, err
		}
		if len(res.Users) == 0 {
			break
		}

		offset = res.Offset
		users.Users = append(users.Users, res.Users...)
		time.Sleep(100 * time.Millisecond)
	}

	return &users, nil
}

func (u *Users) GetUsersFile(ctx context.Context) (*Users, error) {
	file, _ := os.ReadFile("response.json")

	res := Users{}

	if err := json.Unmarshal(file, &res); err != nil {
		log.Fatal(err)
	}

	return &res, nil
}

func (u *User) Disable(ctx context.Context, c *Client) error {

	url := fmt.Sprintf("%s/users/%s/permissions", c.BaseURL, u.Key)
	jsonBody := []byte(`{"disabled":true}`)
	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(http.MethodPut, url, bodyReader)
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("X-Auth-Token", c.token)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	log.Println("StatusCode:", res.StatusCode, u.Email, "Disabled")
	return nil
}

func (u *User) Enable(ctx context.Context, c *Client) error {

	url := fmt.Sprintf("%s/users/%s/permissions", c.BaseURL, u.Key)
	jsonBody := []byte(`{"disabled":false}`)
	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(http.MethodPut, url, bodyReader)
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("X-Auth-Token", c.token)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	log.Println("StatusCode:", res.StatusCode, u.Email, "Enabled")
	return nil
}
