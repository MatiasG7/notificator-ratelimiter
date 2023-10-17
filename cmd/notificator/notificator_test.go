package notificator_test

import (
	"testing"
	"time"

	"github.com/notificator-ratelimiter/cmd/notificator"
	"github.com/stretchr/testify/require"
)

func TestSend(t *testing.T) {
	n := notificator.NewNotificator()

	err := n.Send("project_invitation", "user5", "invitation to project user5")
	require.NoError(t, err)
	err = n.Send("project_invitation", "user5", "invitation to project user5")
	require.NoError(t, err)

	err = n.Send("project_invitation", "user5", "invitation to project user5")
	require.Error(t, err)

	time.Sleep(10 * time.Second)

	err = n.Send("project_invitation", "user5", "invitation to project user5")
	require.NoError(t, err)
}

func TestSend_MultipleUsersAndNotificationTypes(t *testing.T) {
	n := notificator.NewNotificator()
	tt := []struct {
		testName  string
		notifType string
		userID    string
		message   string
		assert    func(err error)
	}{
		{testName: "First status notification user1 - OK",
			notifType: "status",
			userID:    "user1",
			message:   "status to user1",
			assert: func(err error) {
				require.NoError(t, err)
			},
		}, {
			testName:  "Second status notification user1 - OK",
			notifType: "status",
			userID:    "user1",
			message:   "status to user1",
			assert: func(err error) {
				require.NoError(t, err)
			}},
		{testName: "Third status notification user1 - Must fail",
			notifType: "status", userID: "user1",
			message: "status to user1", assert: func(err error) {
				require.Error(t, err)
			},
		}, {
			testName: "First status notification user2 - OK", notifType: "status",
			userID: "user2", message: "status to user2",
			assert: func(err error) {
				require.NoError(t, err)
			}},
		{testName: "First news notification user1 - OK",
			notifType: "news", userID: "user1",
			message: "news to user1", assert: func(err error) {
				require.NoError(t, err)
			},
		}, {
			testName: "Second news notification user1 - Must fail", notifType: "news",
			userID: "user1", message: "news to user1",
			assert: func(err error) {
				require.Error(t, err)
			}},
		{testName: "First news notification user2 - OK",
			notifType: "news", userID: "user2",
			message: "news to user2", assert: func(err error) {
				require.NoError(t, err)
			},
		}, {
			testName: "First marketing notification user1 - OK", notifType: "marketing",
			userID: "user1", message: "marketing to user1",
			assert: func(err error) {
				require.NoError(t, err)
			}},
		{testName: "Second marketing notification user1 - OK",
			notifType: "marketing", userID: "user1",
			message: "marketing to user1", assert: func(err error) {
				require.NoError(t, err)
			},
		}, {
			testName: "Third marketing notification user1 - Must fail", notifType: "marketing",
			userID: "user1", message: "marketing to user1",
			assert: func(err error) {
				require.Error(t, err)
			}},
		{testName: "First marketing notification user2 - OK",
			notifType: "marketing", userID: "user1",
			message: "marketing to user2", assert: func(err error) {
				require.Error(t, err)
			},
		}}
	for _, tc := range tt {
		t.Log(tc.testName)
		err := n.Send(tc.notifType, tc.userID, tc.message)
		tc.assert(err)
	}
}
