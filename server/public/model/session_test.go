package model

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSessionIsValid(t *testing.T) {
	tcs := []struct {
		name          string
		input         Session
		expectedError string
	}{
		{
			"Invalid Id",
			Session{},
			"model.session.is_valid.id.app_error",
		},
		{
			"Invalid UserId",
			Session{
				Id: NewId(),
			},
			"model.session.is_valid.user_id.app_error",
		},
		{
			"Invalid CreateAt",
			Session{
				Id:     NewId(),
				UserId: NewId(),
			},
			"model.session.is_valid.create_at.app_error",
		},
		{
			"Invalid Roles",
			Session{
				Id:       NewId(),
				UserId:   NewId(),
				CreateAt: 1000,
				Roles:    strings.Repeat("a", UserRolesMaxLength+1),
			},
			"model.session.is_valid.roles_limit.app_error",
		},
		{
			"Valid",
			Session{
				Id:       NewId(),
				UserId:   NewId(),
				CreateAt: 1000,
				Roles:    strings.Repeat("a", UserRolesMaxLength),
			},
			"",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			appErr := tc.input.IsValid()
			if tc.expectedError != "" {
				require.NotNil(t, appErr)
				require.Equal(t, tc.expectedError, appErr.Id)
			} else {
				require.Nil(t, appErr)
			}
		})
	}
}

func TestSessionDeepCopy(t *testing.T) {
	sessionId := NewId()
	userId := NewId()
	mapKey := "key"
	mapValue := "val"

	session := &Session{Id: sessionId, Props: map[string]string{mapKey: mapValue}, TeamMembers: []*TeamMember{{UserId: userId, TeamId: "someteamId"}}}

	copySession := session.DeepCopy()
	copySession.Id = "changed"
	copySession.Props[mapKey] = "changed"
	copySession.TeamMembers[0].UserId = "changed"

	assert.Equal(t, sessionId, session.Id)
	assert.Equal(t, mapValue, session.Props[mapKey])
	assert.Equal(t, userId, session.TeamMembers[0].UserId)

	session = &Session{Id: sessionId}
	copySession = session.DeepCopy()

	assert.Equal(t, sessionId, copySession.Id)

	session = &Session{TeamMembers: []*TeamMember{}}
	copySession = session.DeepCopy()

	assert.Empty(t, copySession.TeamMembers)
}

func TestSessionCSRF(t *testing.T) {
	s := Session{}
	token := s.GetCSRF()
	assert.Empty(t, token)

	token = s.GenerateCSRF()
	assert.NotEmpty(t, token)

	token2 := s.GetCSRF()
	assert.NotEmpty(t, token2)
	assert.Equal(t, token, token2)
}

func TestSessionIsOAuthUser(t *testing.T) {
	testCases := []struct {
		Description string
		Session     Session
		isOAuthUser bool
	}{
		{"False on empty props", Session{}, false},
		{"True when key is set to true", Session{Props: StringMap{UserAuthServiceIsOAuth: strconv.FormatBool(true)}}, true},
		{"False when key is set to false", Session{Props: StringMap{UserAuthServiceIsOAuth: strconv.FormatBool(false)}}, false},
		{"Not affected by Session.IsOAuth being true", Session{IsOAuth: true}, false},
		{"Not affected by Session.IsOAuth being false", Session{IsOAuth: false, Props: StringMap{UserAuthServiceIsOAuth: strconv.FormatBool(true)}}, true},
	}

	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {
			require.Equal(t, tc.isOAuthUser, tc.Session.IsOAuthUser())
		})
	}
}

func TestIsIntegration(t *testing.T) {
	testCases := []struct {
		Description   string
		Session       Session
		IsIntegration bool
	}{
		{"False on empty props", Session{}, false},
		{"True when is OAuth App", Session{IsOAuth: true}, true},
		{"True when session is bot", Session{Props: StringMap{SessionPropIsBot: SessionPropIsBotValue}}, true},
		{"True when session is user access token", Session{Props: StringMap{SessionPropType: SessionTypeUserAccessToken}}, true},
		{"Not affected by Props[UserAuthServiceIsOAuth]", Session{Props: StringMap{UserAuthServiceIsOAuth: strconv.FormatBool(true)}}, false},
	}

	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {
			require.Equal(t, tc.IsIntegration, tc.Session.IsIntegration())
		})
	}
}
