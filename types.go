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

type BadRequest struct {
	Status int         `json:"status"`
	Errors interface{} `json:"errors"`
}

type Room struct {
	RoomName    string `json:"roomName,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	//сессионные залы
	SessionHalls []string `json:"sessionHalls,omitempty"`
	Moderators   []User   `json:"moderators,omitempty"`
	//доступно ли приложение для внешних пользователей
	AllowAnonymous                bool      `json:"allowAnonymous,omitempty"`
	AnonymousAccessExpirationDate time.Time `json:"anonymousAccessExpirationDate,omitempty"`
	//политика аудио (none - не задано, muted, disabled)`
	AudioPolicy string `json:"audioPolicy"`
	//политика видео (none - не задано, muted, disabled)
	VideoPolicy string `json:"videoPolicy"`
	//политика показа экрана (none - не задано, muted, disabled)
	ScreenSharePolicy string `json:"screenSharePolicy,omitempty"`
	// информация о не больше чем 3-х уникальных пользователях в комнате
	OnlineUsers []User `json:"onlineUsers,omitempty"`
	// счётчик количества открытых окон\табов (не уникальных пользователей)
	UsersOnline int `json:"usersOnline,omitempty"`
}
