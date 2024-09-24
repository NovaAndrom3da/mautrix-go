// GENERATED BY portalinternal_generate.go; DO NOT EDIT

//go:generate go run portalinternal_generate.go
//go:generate goimports -local maunium.net/go/mautrix -w portalinternal.go

package bridgev2

import (
	"context"
	"time"

	"github.com/rs/zerolog"

	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/bridgev2/database"
	"maunium.net/go/mautrix/bridgev2/networkid"
	"maunium.net/go/mautrix/event"
	"maunium.net/go/mautrix/id"
)

type PortalInternals Portal

// Deprecated: portal internals should be used carefully and only when necessary.
func (portal *Portal) Internal() *PortalInternals {
	return (*PortalInternals)(portal)
}

func (portal *PortalInternals) UpdateLogger() {
	(*Portal)(portal).updateLogger()
}

func (portal *PortalInternals) QueueEvent(ctx context.Context, evt portalEvent) {
	(*Portal)(portal).queueEvent(ctx, evt)
}

func (portal *PortalInternals) EventLoop() {
	(*Portal)(portal).eventLoop()
}

func (portal *PortalInternals) HandleSingleEventAsync(idx int, rawEvt any) {
	(*Portal)(portal).handleSingleEventAsync(idx, rawEvt)
}

func (portal *PortalInternals) GetEventCtxWithLog(rawEvt any, idx int) context.Context {
	return (*Portal)(portal).getEventCtxWithLog(rawEvt, idx)
}

func (portal *PortalInternals) HandleSingleEvent(ctx context.Context, rawEvt any, doneCallback func()) {
	(*Portal)(portal).handleSingleEvent(ctx, rawEvt, doneCallback)
}

func (portal *PortalInternals) SendSuccessStatus(ctx context.Context, evt *event.Event, streamOrder int64) {
	(*Portal)(portal).sendSuccessStatus(ctx, evt, streamOrder)
}

func (portal *PortalInternals) SendErrorStatus(ctx context.Context, evt *event.Event, err error) {
	(*Portal)(portal).sendErrorStatus(ctx, evt, err)
}

func (portal *PortalInternals) CheckConfusableName(ctx context.Context, userID id.UserID, name string) bool {
	return (*Portal)(portal).checkConfusableName(ctx, userID, name)
}

func (portal *PortalInternals) HandleMatrixEvent(ctx context.Context, sender *User, evt *event.Event) {
	(*Portal)(portal).handleMatrixEvent(ctx, sender, evt)
}

func (portal *PortalInternals) HandleMatrixReceipts(ctx context.Context, evt *event.Event) {
	(*Portal)(portal).handleMatrixReceipts(ctx, evt)
}

func (portal *PortalInternals) HandleMatrixReadReceipt(ctx context.Context, user *User, eventID id.EventID, receipt event.ReadReceipt) {
	(*Portal)(portal).handleMatrixReadReceipt(ctx, user, eventID, receipt)
}

func (portal *PortalInternals) HandleMatrixTyping(ctx context.Context, evt *event.Event) {
	(*Portal)(portal).handleMatrixTyping(ctx, evt)
}

func (portal *PortalInternals) SendTypings(ctx context.Context, userIDs []id.UserID, typing bool) {
	(*Portal)(portal).sendTypings(ctx, userIDs, typing)
}

func (portal *PortalInternals) PeriodicTypingUpdater() {
	(*Portal)(portal).periodicTypingUpdater()
}

func (portal *PortalInternals) CheckMessageContentCaps(ctx context.Context, caps *NetworkRoomCapabilities, content *event.MessageEventContent, evt *event.Event) bool {
	return (*Portal)(portal).checkMessageContentCaps(ctx, caps, content, evt)
}

func (portal *PortalInternals) HandleMatrixMessage(ctx context.Context, sender *UserLogin, origSender *OrigSender, evt *event.Event) {
	(*Portal)(portal).handleMatrixMessage(ctx, sender, origSender, evt)
}

func (portal *PortalInternals) HandleMatrixEdit(ctx context.Context, sender *UserLogin, origSender *OrigSender, evt *event.Event, content *event.MessageEventContent, caps *NetworkRoomCapabilities) {
	(*Portal)(portal).handleMatrixEdit(ctx, sender, origSender, evt, content, caps)
}

func (portal *PortalInternals) HandleMatrixReaction(ctx context.Context, sender *UserLogin, evt *event.Event) {
	(*Portal)(portal).handleMatrixReaction(ctx, sender, evt)
}

func (portal *PortalInternals) GetTargetUser(ctx context.Context, userID id.UserID) (GhostOrUserLogin, error) {
	return (*Portal)(portal).getTargetUser(ctx, userID)
}

func (portal *PortalInternals) HandleMatrixMembership(ctx context.Context, sender *UserLogin, origSender *OrigSender, evt *event.Event) {
	(*Portal)(portal).handleMatrixMembership(ctx, sender, origSender, evt)
}

func (portal *PortalInternals) HandleMatrixPowerLevels(ctx context.Context, sender *UserLogin, origSender *OrigSender, evt *event.Event) {
	(*Portal)(portal).handleMatrixPowerLevels(ctx, sender, origSender, evt)
}

func (portal *PortalInternals) HandleMatrixRedaction(ctx context.Context, sender *UserLogin, origSender *OrigSender, evt *event.Event) {
	(*Portal)(portal).handleMatrixRedaction(ctx, sender, origSender, evt)
}

func (portal *PortalInternals) HandleRemoteEvent(ctx context.Context, source *UserLogin, evtType RemoteEventType, evt RemoteEvent) {
	(*Portal)(portal).handleRemoteEvent(ctx, source, evtType, evt)
}

func (portal *PortalInternals) GetIntentAndUserMXIDFor(ctx context.Context, sender EventSender, source *UserLogin, otherLogins []*UserLogin, evtType RemoteEventType) (intent MatrixAPI, extraUserID id.UserID) {
	return (*Portal)(portal).getIntentAndUserMXIDFor(ctx, sender, source, otherLogins, evtType)
}

func (portal *PortalInternals) GetRelationMeta(ctx context.Context, currentMsg networkid.MessageID, replyToPtr *networkid.MessageOptionalPartID, threadRootPtr *networkid.MessageID, isBatchSend bool) (replyTo, threadRoot, prevThreadEvent *database.Message) {
	return (*Portal)(portal).getRelationMeta(ctx, currentMsg, replyToPtr, threadRootPtr, isBatchSend)
}

func (portal *PortalInternals) ApplyRelationMeta(content *event.MessageEventContent, replyTo, threadRoot, prevThreadEvent *database.Message) {
	(*Portal)(portal).applyRelationMeta(content, replyTo, threadRoot, prevThreadEvent)
}

func (portal *PortalInternals) SendConvertedMessage(ctx context.Context, id networkid.MessageID, intent MatrixAPI, senderID networkid.UserID, converted *ConvertedMessage, ts time.Time, streamOrder int64, logContext func(*zerolog.Event) *zerolog.Event) []*database.Message {
	return (*Portal)(portal).sendConvertedMessage(ctx, id, intent, senderID, converted, ts, streamOrder, logContext)
}

func (portal *PortalInternals) CheckPendingMessage(ctx context.Context, evt RemoteMessage) (bool, *database.Message) {
	return (*Portal)(portal).checkPendingMessage(ctx, evt)
}

func (portal *PortalInternals) HandleRemoteUpsert(ctx context.Context, source *UserLogin, evt RemoteMessageUpsert, existing []*database.Message) bool {
	return (*Portal)(portal).handleRemoteUpsert(ctx, source, evt, existing)
}

func (portal *PortalInternals) HandleRemoteMessage(ctx context.Context, source *UserLogin, evt RemoteMessage) {
	(*Portal)(portal).handleRemoteMessage(ctx, source, evt)
}

func (portal *PortalInternals) SendRemoteErrorNotice(ctx context.Context, intent MatrixAPI, err error, ts time.Time, evtTypeName string) {
	(*Portal)(portal).sendRemoteErrorNotice(ctx, intent, err, ts, evtTypeName)
}

func (portal *PortalInternals) HandleRemoteEdit(ctx context.Context, source *UserLogin, evt RemoteEdit) {
	(*Portal)(portal).handleRemoteEdit(ctx, source, evt)
}

func (portal *PortalInternals) SendConvertedEdit(ctx context.Context, targetID networkid.MessageID, senderID networkid.UserID, converted *ConvertedEdit, intent MatrixAPI, ts time.Time, streamOrder int64) {
	(*Portal)(portal).sendConvertedEdit(ctx, targetID, senderID, converted, intent, ts, streamOrder)
}

func (portal *PortalInternals) GetTargetMessagePart(ctx context.Context, evt RemoteEventWithTargetMessage) (*database.Message, error) {
	return (*Portal)(portal).getTargetMessagePart(ctx, evt)
}

func (portal *PortalInternals) GetTargetReaction(ctx context.Context, evt RemoteReactionRemove) (*database.Reaction, error) {
	return (*Portal)(portal).getTargetReaction(ctx, evt)
}

func (portal *PortalInternals) HandleRemoteReactionSync(ctx context.Context, source *UserLogin, evt RemoteReactionSync) {
	(*Portal)(portal).handleRemoteReactionSync(ctx, source, evt)
}

func (portal *PortalInternals) HandleRemoteReaction(ctx context.Context, source *UserLogin, evt RemoteReaction) {
	(*Portal)(portal).handleRemoteReaction(ctx, source, evt)
}

func (portal *PortalInternals) SendConvertedReaction(ctx context.Context, senderID networkid.UserID, intent MatrixAPI, targetMessage *database.Message, emojiID networkid.EmojiID, emoji string, ts time.Time, dbMetadata any, extraContent map[string]any, logContext func(*zerolog.Event) *zerolog.Event) {
	(*Portal)(portal).sendConvertedReaction(ctx, senderID, intent, targetMessage, emojiID, emoji, ts, dbMetadata, extraContent, logContext)
}

func (portal *PortalInternals) GetIntentForMXID(ctx context.Context, userID id.UserID) (MatrixAPI, error) {
	return (*Portal)(portal).getIntentForMXID(ctx, userID)
}

func (portal *PortalInternals) HandleRemoteReactionRemove(ctx context.Context, source *UserLogin, evt RemoteReactionRemove) {
	(*Portal)(portal).handleRemoteReactionRemove(ctx, source, evt)
}

func (portal *PortalInternals) HandleRemoteMessageRemove(ctx context.Context, source *UserLogin, evt RemoteMessageRemove) {
	(*Portal)(portal).handleRemoteMessageRemove(ctx, source, evt)
}

func (portal *PortalInternals) RedactMessageParts(ctx context.Context, parts []*database.Message, intent MatrixAPI, ts time.Time) {
	(*Portal)(portal).redactMessageParts(ctx, parts, intent, ts)
}

func (portal *PortalInternals) HandleRemoteReadReceipt(ctx context.Context, source *UserLogin, evt RemoteReadReceipt) {
	(*Portal)(portal).handleRemoteReadReceipt(ctx, source, evt)
}

func (portal *PortalInternals) HandleRemoteMarkUnread(ctx context.Context, source *UserLogin, evt RemoteMarkUnread) {
	(*Portal)(portal).handleRemoteMarkUnread(ctx, source, evt)
}

func (portal *PortalInternals) HandleRemoteDeliveryReceipt(ctx context.Context, source *UserLogin, evt RemoteDeliveryReceipt) {
	(*Portal)(portal).handleRemoteDeliveryReceipt(ctx, source, evt)
}

func (portal *PortalInternals) HandleRemoteTyping(ctx context.Context, source *UserLogin, evt RemoteTyping) {
	(*Portal)(portal).handleRemoteTyping(ctx, source, evt)
}

func (portal *PortalInternals) HandleRemoteChatInfoChange(ctx context.Context, source *UserLogin, evt RemoteChatInfoChange) {
	(*Portal)(portal).handleRemoteChatInfoChange(ctx, source, evt)
}

func (portal *PortalInternals) HandleRemoteChatResync(ctx context.Context, source *UserLogin, evt RemoteChatResync) {
	(*Portal)(portal).handleRemoteChatResync(ctx, source, evt)
}

func (portal *PortalInternals) HandleRemoteChatDelete(ctx context.Context, source *UserLogin, evt RemoteChatDelete) {
	(*Portal)(portal).handleRemoteChatDelete(ctx, source, evt)
}

func (portal *PortalInternals) HandleRemoteBackfill(ctx context.Context, source *UserLogin, backfill RemoteBackfill) {
	(*Portal)(portal).handleRemoteBackfill(ctx, source, backfill)
}

func (portal *PortalInternals) UpdateName(ctx context.Context, name string, sender MatrixAPI, ts time.Time) bool {
	return (*Portal)(portal).updateName(ctx, name, sender, ts)
}

func (portal *PortalInternals) UpdateTopic(ctx context.Context, topic string, sender MatrixAPI, ts time.Time) bool {
	return (*Portal)(portal).updateTopic(ctx, topic, sender, ts)
}

func (portal *PortalInternals) UpdateAvatar(ctx context.Context, avatar *Avatar, sender MatrixAPI, ts time.Time) bool {
	return (*Portal)(portal).updateAvatar(ctx, avatar, sender, ts)
}

func (portal *PortalInternals) GetBridgeInfo() (string, event.BridgeEventContent) {
	return (*Portal)(portal).getBridgeInfo()
}

func (portal *PortalInternals) SendStateWithIntentOrBot(ctx context.Context, sender MatrixAPI, eventType event.Type, stateKey string, content *event.Content, ts time.Time) (resp *mautrix.RespSendEvent, err error) {
	return (*Portal)(portal).sendStateWithIntentOrBot(ctx, sender, eventType, stateKey, content, ts)
}

func (portal *PortalInternals) SendRoomMeta(ctx context.Context, sender MatrixAPI, ts time.Time, eventType event.Type, stateKey string, content any) bool {
	return (*Portal)(portal).sendRoomMeta(ctx, sender, ts, eventType, stateKey, content)
}

func (portal *PortalInternals) GetInitialMemberList(ctx context.Context, members *ChatMemberList, source *UserLogin, pl *event.PowerLevelsEventContent) (invite, functional []id.UserID, err error) {
	return (*Portal)(portal).getInitialMemberList(ctx, members, source, pl)
}

func (portal *PortalInternals) UpdateOtherUser(ctx context.Context, members *ChatMemberList) (changed bool) {
	return (*Portal)(portal).updateOtherUser(ctx, members)
}

func (portal *PortalInternals) SyncParticipants(ctx context.Context, members *ChatMemberList, source *UserLogin, sender MatrixAPI, ts time.Time) error {
	return (*Portal)(portal).syncParticipants(ctx, members, source, sender, ts)
}

func (portal *PortalInternals) UpdateUserLocalInfo(ctx context.Context, info *UserLocalPortalInfo, source *UserLogin, didJustCreate bool) {
	(*Portal)(portal).updateUserLocalInfo(ctx, info, source, didJustCreate)
}

func (portal *PortalInternals) UpdateParent(ctx context.Context, newParentID networkid.PortalID, source *UserLogin) bool {
	return (*Portal)(portal).updateParent(ctx, newParentID, source)
}

func (portal *PortalInternals) LockedUpdateInfoFromGhost(ctx context.Context, ghost *Ghost) {
	(*Portal)(portal).lockedUpdateInfoFromGhost(ctx, ghost)
}

func (portal *PortalInternals) CreateMatrixRoomInLoop(ctx context.Context, source *UserLogin, info *ChatInfo, backfillBundle any) error {
	return (*Portal)(portal).createMatrixRoomInLoop(ctx, source, info, backfillBundle)
}

func (portal *PortalInternals) RemoveInPortalCache(ctx context.Context) {
	(*Portal)(portal).removeInPortalCache(ctx)
}

func (portal *PortalInternals) UnlockedDelete(ctx context.Context) error {
	return (*Portal)(portal).unlockedDelete(ctx)
}

func (portal *PortalInternals) UnlockedDeleteCache() {
	(*Portal)(portal).unlockedDeleteCache()
}

func (portal *PortalInternals) DoForwardBackfill(ctx context.Context, source *UserLogin, lastMessage *database.Message, bundledData any) {
	(*Portal)(portal).doForwardBackfill(ctx, source, lastMessage, bundledData)
}

func (portal *PortalInternals) FetchThreadBackfill(ctx context.Context, source *UserLogin, anchor *database.Message) *FetchMessagesResponse {
	return (*Portal)(portal).fetchThreadBackfill(ctx, source, anchor)
}

func (portal *PortalInternals) DoThreadBackfill(ctx context.Context, source *UserLogin, threadID networkid.MessageID) {
	(*Portal)(portal).doThreadBackfill(ctx, source, threadID)
}

func (portal *PortalInternals) CutoffMessages(ctx context.Context, messages []*BackfillMessage, aggressiveDedup, forward bool, lastMessage *database.Message) []*BackfillMessage {
	return (*Portal)(portal).cutoffMessages(ctx, messages, aggressiveDedup, forward, lastMessage)
}

func (portal *PortalInternals) SendBackfill(ctx context.Context, source *UserLogin, messages []*BackfillMessage, forceForward, markRead, inThread bool, done func()) {
	(*Portal)(portal).sendBackfill(ctx, source, messages, forceForward, markRead, inThread, done)
}

func (portal *PortalInternals) CompileBatchMessage(ctx context.Context, source *UserLogin, msg *BackfillMessage, out *compileBatchOutput, inThread bool) {
	(*Portal)(portal).compileBatchMessage(ctx, source, msg, out, inThread)
}

func (portal *PortalInternals) FetchThreadInsideBatch(ctx context.Context, source *UserLogin, dbMsg *database.Message, out *compileBatchOutput) {
	(*Portal)(portal).fetchThreadInsideBatch(ctx, source, dbMsg, out)
}

func (portal *PortalInternals) SendBatch(ctx context.Context, source *UserLogin, messages []*BackfillMessage, forceForward, markRead, inThread bool) {
	(*Portal)(portal).sendBatch(ctx, source, messages, forceForward, markRead, inThread)
}

func (portal *PortalInternals) SendLegacyBackfill(ctx context.Context, source *UserLogin, messages []*BackfillMessage, markRead bool) {
	(*Portal)(portal).sendLegacyBackfill(ctx, source, messages, markRead)
}

func (portal *PortalInternals) UnlockedReID(ctx context.Context, target networkid.PortalKey) error {
	return (*Portal)(portal).unlockedReID(ctx, target)
}

func (portal *PortalInternals) CreateParentAndAddToSpace(ctx context.Context, source *UserLogin) {
	(*Portal)(portal).createParentAndAddToSpace(ctx, source)
}

func (portal *PortalInternals) AddToParentSpaceAndSave(ctx context.Context, save bool) {
	(*Portal)(portal).addToParentSpaceAndSave(ctx, save)
}

func (portal *PortalInternals) ToggleSpace(ctx context.Context, spaceID id.RoomID, canonical, remove bool) error {
	return (*Portal)(portal).toggleSpace(ctx, spaceID, canonical, remove)
}

func (portal *PortalInternals) SetMXIDToExistingRoom(roomID id.RoomID) bool {
	return (*Portal)(portal).setMXIDToExistingRoom(roomID)
}
