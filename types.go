// Контур.Толк API - https://kontur.renote.team/doc/gmVrwgTNW

package ktalk

import "time"

type User struct {
	UserKey   string   `json:"userKey"`
	Login     string   `json:"login,omitempty"`
	HasEmail  bool     `json:"hasEmail,omitempty"`
	Email     string   `json:"email,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Surname   string   `json:"surname,omitempty"`
	Roles     []string `json:"roles,omitempty"`
	RoleInfos []struct {
		Id    string `json:"id,omitempty"`
		Title string `json:"title,omitempty"`
	} `json:"roleInfos,omitempty"`
	Key         string `json:"key"`
	Post        string `json:"post,omitempty"`
	AvatarUrl   string `json:"avatarUrl,omitempty"`
	MobilePhone string `json:"mobilePhone,omitempty"`
	InnerPhone  string `json:"innerPhone,omitempty"`
	Status      string `json:"status,omitempty"`
}

type Users struct {
	Users  []User `json:"users"`
	Offset string `json:"offset"`
}

type BadRequest struct {
	Status int         `json:"status"`
	Errors interface{} `json:"errors"`
}

type Room struct {
	RoomName                      string    `json:"roomName,omitempty"`
	Title                         string    `json:"title,omitempty"`
	Description                   string    `json:"description,omitempty"`
	SessionHalls                  []string  `json:"sessionHalls,omitempty"` //сессионные залы
	Moderators                    []User    `json:"moderators,omitempty"`
	AllowAnonymous                bool      `json:"allowAnonymous,omitempty"` //доступно ли приложение для внешних пользователей
	AnonymousAccessExpirationDate time.Time `json:"anonymousAccessExpirationDate,omitempty"`
	AudioPolicy                   string    `json:"audioPolicy"`                 //политика аудио (none - не задано, muted, disabled)`
	VideoPolicy                   string    `json:"videoPolicy"`                 //политика видео (none - не задано, muted, disabled)
	ScreenSharePolicy             string    `json:"screenSharePolicy,omitempty"` //политика показа экрана (none - не задано, muted, disabled)
	OnlineUsers                   []User    `json:"onlineUsers,omitempty"`       // информация о не больше чем 3-х уникальных пользователях в комнате
	UsersOnline                   int       `json:"usersOnline,omitempty"`       // счётчик количества открытых окон\табов (не уникальных пользователей)
}

type Users2 map[string]map[string]interface{}
