package types

type Action string

const (
	ActionCreateGroup Action = "create_group"
	ActionJoinGroup   Action = "join_group"
	ActionLiveGroup   Action = "leave_group"

	ActionSendMessage   Action = "send_message"
	ActionEditMessage   Action = "edit_message"
	ActionDeleteMessage Action = "delete_message"

	ActionMessageSeen Action = "message_seen"
)

const (
	ActionMessage Action = "update_message"
)
