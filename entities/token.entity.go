package entities

import "time"

type TokenObject struct {
	TokenID       string `json:"tokenid,omitempty"`
	Name          string `json:"name,omitempty"`
	Description   string `json:"description,omitempty"`
	UserID        string `json:"userid,omitempty"`
	LastAccess    string `json:"lastaccess,omitempty"`
	Status        int    `json:"status,omitempty"`
	ExpiresAt     string `json:"expires_at,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	CreatorUserID string `json:"creator_userid,omitempty"`
}

type TokenCreate struct {
	Name          string `json:"name,omitempty"`
	Description   string `json:"description,omitempty"`
	UserID        string `json:"userid,omitempty"`
	LastAccess    string `json:"lastaccess,omitempty"`
	Status        int    `json:"status,omitempty"`
	ExpiresAt     string `json:"expires_at,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	CreatorUserID string `json:"creator_userid,omitempty"`
}

type TokenGenerateResponse struct {
	TokenID string `json:"tokenid,omitempty"`
	Token   string `json:"token,omitempty"`
}

type TokenUpdate struct {
	TokenID       string `json:"tokenid,omitempty"`
	Name          string `json:"name,omitempty"`
	Description   string `json:"description,omitempty"`
	UserID        string `json:"userid,omitempty"`
	LastAccess    string `json:"lastaccess,omitempty"`
	Status        int    `json:"status,omitempty"`
	ExpiresAt     string `json:"expires_at,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	CreatorUserID string `json:"creator_userid,omitempty"`
}

type TokenResponse struct {
	TokenIDs []string `json:"tokenids,omitempty"`
}

type TokenGet struct {
	TokenIDs  []string  `json:"tokenids,omitempty"`
	UserIDs   []string  `json:"userids,omitempty"`
	Token     string    `json:"token,omitempty"`
	ValidAt   time.Time `json:"valid_at,omitempty"`
	ExpiredAt time.Time `json:"expired_at,omitempty"`
}
