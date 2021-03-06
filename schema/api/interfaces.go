package api

// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
    "context"
    "generated"
    "github.com/Azure/go-autorest/autorest"
)

        // AttachmentsClientAPI contains the set of methods on the AttachmentsClient type.
        type AttachmentsClientAPI interface {
            GetAttachmentInfoMethod(ctx context.Context, attachmentID string) (result .AttachmentInfoType, err error)
            GetAttachmentMethod(ctx context.Context, attachmentID string, viewID string) (result .ReadCloserType, err error)
        }

        var _ AttachmentsClientAPI = (*.AttachmentsClient)(nil)
        // ConversationsClientAPI contains the set of methods on the ConversationsClient type.
        type ConversationsClientAPI interface {
            CreateConversationMethod(ctx context.Context, parameters .ConversationParametersType) (result .ConversationResourceResponseType, err error)
            DeleteActivityMethod(ctx context.Context, conversationID string, activityID string) (result autorest.Response, err error)
            DeleteConversationMemberMethod(ctx context.Context, conversationID string, memberID string) (result autorest.Response, err error)
            GetActivityMembersMethod(ctx context.Context, conversationID string, activityID string) (result .ListChannelAccountType, err error)
            GetConversationMembersMethod(ctx context.Context, conversationID string) (result .ListChannelAccountType, err error)
            GetConversationPagedMembersMethod(ctx context.Context, conversationID string, pageSize *int32, continuationToken string) (result .PagedMembersResultType, err error)
            GetConversationsMethod(ctx context.Context, continuationToken string) (result .ConversationsResultType, err error)
            ReplyToActivityMethod(ctx context.Context, conversationID string, activityID string, activity .ActivityType) (result .ResourceResponseType, err error)
            SendConversationHistoryMethod(ctx context.Context, conversationID string, history .TranscriptType) (result .ResourceResponseType, err error)
            SendToConversationMethod(ctx context.Context, conversationID string, activity .ActivityType) (result .ResourceResponseType, err error)
            UpdateActivityMethod(ctx context.Context, conversationID string, activityID string, activity .ActivityType) (result .ResourceResponseType, err error)
            UploadAttachmentMethod(ctx context.Context, conversationID string, attachmentUpload .AttachmentDataType) (result .ResourceResponseType, err error)
        }

        var _ ConversationsClientAPI = (*.ConversationsClient)(nil)
