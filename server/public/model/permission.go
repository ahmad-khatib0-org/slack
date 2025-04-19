package model

import (
	"net/http"
)

const (
	PermissionScopeSystem   = "system_scope"
	PermissionScopeTeam     = "team_scope"
	PermissionScopeChannel  = "channel_scope"
	PermissionScopeGroup    = "group_scope"
	PermissionScopePlaybook = "playbook_scope"
	PermissionScopeRun      = "run_scope"
)

type Permission struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Scope       string `json:"scope"`
}

var (
	PermissionInviteUser    *Permission
	PermissionAddUserToTeam *Permission
)

// Deprecated: PermissionUseSlashCommands is not longer used. It's only kept for backwards compatibility.
// See https://mattermost.atlassian.net/browse/MM-52574 for more details.
var (
	PermissionUseSlashCommands                      *Permission
	PermissionManageSlashCommands                   *Permission
	PermissionManageOthersSlashCommands             *Permission
	PermissionCreatePublicChannel                   *Permission
	PermissionCreatePrivateChannel                  *Permission
	PermissionManagePublicChannelMembers            *Permission
	PermissionManagePrivateChannelMembers           *Permission
	PermissionConvertPublicChannelToPrivate         *Permission
	PermissionConvertPrivateChannelToPublic         *Permission
	PermissionAssignSystemAdminRole                 *Permission
	PermissionManageRoles                           *Permission
	PermissionManageTeamRoles                       *Permission
	PermissionManageChannelRoles                    *Permission
	PermissionCreateDirectChannel                   *Permission
	PermissionCreateGroupChannel                    *Permission
	PermissionManagePublicChannelProperties         *Permission
	PermissionManagePrivateChannelProperties        *Permission
	PermissionListPublicTeams                       *Permission
	PermissionJoinPublicTeams                       *Permission
	PermissionListPrivateTeams                      *Permission
	PermissionJoinPrivateTeams                      *Permission
	PermissionListTeamChannels                      *Permission
	PermissionJoinPublicChannels                    *Permission
	PermissionDeletePublicChannel                   *Permission
	PermissionDeletePrivateChannel                  *Permission
	PermissionEditOtherUsers                        *Permission
	PermissionReadChannel                           *Permission
	PermissionReadChannelContent                    *Permission
	PermissionReadPublicChannelGroups               *Permission
	PermissionReadPrivateChannelGroups              *Permission
	PermissionReadPublicChannel                     *Permission
	PermissionAddReaction                           *Permission
	PermissionRemoveReaction                        *Permission
	PermissionRemoveOthersReactions                 *Permission
	PermissionPermanentDeleteUser                   *Permission
	PermissionUploadFile                            *Permission
	PermissionGetPublicLink                         *Permission
	PermissionManageWebhooks                        *Permission
	PermissionManageOthersWebhooks                  *Permission
	PermissionManageIncomingWebhooks                *Permission
	PermissionManageOutgoingWebhooks                *Permission
	PermissionManageOthersIncomingWebhooks          *Permission
	PermissionManageOthersOutgoingWebhooks          *Permission
	PermissionManageOAuth                           *Permission
	PermissionManageSystemWideOAuth                 *Permission
	PermissionManageEmojis                          *Permission
	PermissionManageOthersEmojis                    *Permission
	PermissionCreateEmojis                          *Permission
	PermissionDeleteEmojis                          *Permission
	PermissionDeleteOthersEmojis                    *Permission
	PermissionCreatePost                            *Permission
	PermissionCreatePostPublic                      *Permission
	PermissionCreatePostEphemeral                   *Permission
	PermissionReadDeletedPosts                      *Permission
	PermissionEditPost                              *Permission
	PermissionEditOthersPosts                       *Permission
	PermissionDeletePost                            *Permission
	PermissionDeleteOthersPosts                     *Permission
	PermissionRemoveUserFromTeam                    *Permission
	PermissionCreateTeam                            *Permission
	PermissionManageTeam                            *Permission
	PermissionImportTeam                            *Permission
	PermissionViewTeam                              *Permission
	PermissionListUsersWithoutTeam                  *Permission
	PermissionReadJobs                              *Permission
	PermissionManageJobs                            *Permission
	PermissionCreateUserAccessToken                 *Permission
	PermissionReadUserAccessToken                   *Permission
	PermissionRevokeUserAccessToken                 *Permission
	PermissionCreateBot                             *Permission
	PermissionAssignBot                             *Permission
	PermissionReadBots                              *Permission
	PermissionReadOthersBots                        *Permission
	PermissionManageBots                            *Permission
	PermissionManageOthersBots                      *Permission
	PermissionViewMembers                           *Permission
	PermissionInviteGuest                           *Permission
	PermissionPromoteGuest                          *Permission
	PermissionDemoteToGuest                         *Permission
	PermissionUseChannelMentions                    *Permission
	PermissionUseGroupMentions                      *Permission
	PermissionAddBookmarkPublicChannel              *Permission
	PermissionEditBookmarkPublicChannel             *Permission
	PermissionDeleteBookmarkPublicChannel           *Permission
	PermissionOrderBookmarkPublicChannel            *Permission
	PermissionAddBookmarkPrivateChannel             *Permission
	PermissionEditBookmarkPrivateChannel            *Permission
	PermissionDeleteBookmarkPrivateChannel          *Permission
	PermissionOrderBookmarkPrivateChannel           *Permission
	PermissionReadOtherUsersTeams                   *Permission
	PermissionEditBrand                             *Permission
	PermissionManageSharedChannels                  *Permission
	PermissionManageSecureConnections               *Permission
	PermissionDownloadComplianceExportResult        *Permission
	PermissionCreateDataRetentionJob                *Permission
	PermissionManageDataRetentionJob                *Permission
	PermissionReadDataRetentionJob                  *Permission
	PermissionCreateComplianceExportJob             *Permission
	PermissionManageComplianceExportJob             *Permission
	PermissionReadComplianceExportJob               *Permission
	PermissionReadAudits                            *Permission
	PermissionTestElasticsearch                     *Permission
	PermissionTestSiteURL                           *Permission
	PermissionTestS3                                *Permission
	PermissionReloadConfig                          *Permission
	PermissionInvalidateCaches                      *Permission
	PermissionRecycleDatabaseConnections            *Permission
	PermissionPurgeElasticsearchIndexes             *Permission
	PermissionTestEmail                             *Permission
	PermissionCreateElasticsearchPostIndexingJob    *Permission
	PermissionManageElasticsearchPostIndexingJob    *Permission
	PermissionCreateElasticsearchPostAggregationJob *Permission
	PermissionManageElasticsearchPostAggregationJob *Permission
	PermissionReadElasticsearchPostIndexingJob      *Permission
	PermissionReadElasticsearchPostAggregationJob   *Permission
	PermissionPurgeBleveIndexes                     *Permission
	PermissionCreatePostBleveIndexesJob             *Permission
	PermissionManagePostBleveIndexesJob             *Permission
	PermissionCreateLdapSyncJob                     *Permission
	PermissionManageLdapSyncJob                     *Permission
	PermissionReadLdapSyncJob                       *Permission
	PermissionTestLdap                              *Permission
	PermissionInvalidateEmailInvite                 *Permission
	PermissionGetSamlMetadataFromIdp                *Permission
	PermissionAddSamlPublicCert                     *Permission
	PermissionAddSamlPrivateCert                    *Permission
	PermissionAddSamlIdpCert                        *Permission
	PermissionRemoveSamlPublicCert                  *Permission
	PermissionRemoveSamlPrivateCert                 *Permission
	PermissionRemoveSamlIdpCert                     *Permission
	PermissionGetSamlCertStatus                     *Permission
	PermissionAddLdapPublicCert                     *Permission
	PermissionAddLdapPrivateCert                    *Permission
	PermissionRemoveLdapPublicCert                  *Permission
	PermissionRemoveLdapPrivateCert                 *Permission
	PermissionGetLogs                               *Permission
	PermissionGetAnalytics                          *Permission
	PermissionReadLicenseInformation                *Permission
	PermissionManageLicenseInformation              *Permission
)

var (
	PermissionSysconsoleReadAbout  *Permission
	PermissionSysconsoleWriteAbout *Permission
)

var (
	PermissionSysconsoleReadAboutEditionAndLicense  *Permission
	PermissionSysconsoleWriteAboutEditionAndLicense *Permission
)

var (
	PermissionSysconsoleReadBilling  *Permission
	PermissionSysconsoleWriteBilling *Permission
)

var (
	PermissionSysconsoleReadReporting  *Permission
	PermissionSysconsoleWriteReporting *Permission
)

var (
	PermissionSysconsoleReadReportingSiteStatistics  *Permission
	PermissionSysconsoleWriteReportingSiteStatistics *Permission
)

var (
	PermissionSysconsoleReadReportingTeamStatistics  *Permission
	PermissionSysconsoleWriteReportingTeamStatistics *Permission
)

var (
	PermissionSysconsoleReadReportingServerLogs  *Permission
	PermissionSysconsoleWriteReportingServerLogs *Permission
)

var (
	PermissionSysconsoleReadUserManagementUsers  *Permission
	PermissionSysconsoleWriteUserManagementUsers *Permission
)

var (
	PermissionSysconsoleReadUserManagementGroups  *Permission
	PermissionSysconsoleWriteUserManagementGroups *Permission
)

var (
	PermissionSysconsoleReadUserManagementTeams  *Permission
	PermissionSysconsoleWriteUserManagementTeams *Permission
)

var (
	PermissionSysconsoleReadUserManagementChannels  *Permission
	PermissionSysconsoleWriteUserManagementChannels *Permission
)

var (
	PermissionSysconsoleReadUserManagementPermissions  *Permission
	PermissionSysconsoleWriteUserManagementPermissions *Permission
)

var (
	PermissionSysconsoleReadUserManagementSystemRoles  *Permission
	PermissionSysconsoleWriteUserManagementSystemRoles *Permission
)

// DEPRECATED
var PermissionSysconsoleReadEnvironment *Permission

// DEPRECATED
var PermissionSysconsoleWriteEnvironment *Permission

var (
	PermissionSysconsoleReadEnvironmentWebServer  *Permission
	PermissionSysconsoleWriteEnvironmentWebServer *Permission
)

var (
	PermissionSysconsoleReadEnvironmentDatabase  *Permission
	PermissionSysconsoleWriteEnvironmentDatabase *Permission
)

var (
	PermissionSysconsoleReadEnvironmentElasticsearch  *Permission
	PermissionSysconsoleWriteEnvironmentElasticsearch *Permission
)

var (
	PermissionSysconsoleReadEnvironmentFileStorage  *Permission
	PermissionSysconsoleWriteEnvironmentFileStorage *Permission
)

var (
	PermissionSysconsoleReadEnvironmentImageProxy  *Permission
	PermissionSysconsoleWriteEnvironmentImageProxy *Permission
)

var (
	PermissionSysconsoleReadEnvironmentSMTP  *Permission
	PermissionSysconsoleWriteEnvironmentSMTP *Permission
)

var (
	PermissionSysconsoleReadEnvironmentPushNotificationServer  *Permission
	PermissionSysconsoleWriteEnvironmentPushNotificationServer *Permission
)

var (
	PermissionSysconsoleReadEnvironmentHighAvailability  *Permission
	PermissionSysconsoleWriteEnvironmentHighAvailability *Permission
)

var (
	PermissionSysconsoleReadEnvironmentRateLimiting  *Permission
	PermissionSysconsoleWriteEnvironmentRateLimiting *Permission
)

var (
	PermissionSysconsoleReadEnvironmentLogging  *Permission
	PermissionSysconsoleWriteEnvironmentLogging *Permission
)

var (
	PermissionSysconsoleReadEnvironmentSessionLengths  *Permission
	PermissionSysconsoleWriteEnvironmentSessionLengths *Permission
)

var (
	PermissionSysconsoleReadEnvironmentPerformanceMonitoring  *Permission
	PermissionSysconsoleWriteEnvironmentPerformanceMonitoring *Permission
)

var (
	PermissionSysconsoleReadEnvironmentDeveloper  *Permission
	PermissionSysconsoleWriteEnvironmentDeveloper *Permission
)

var (
	PermissionSysconsoleReadEnvironmentMobileSecurity  *Permission
	PermissionSysconsoleWriteEnvironmentMobileSecurity *Permission
)

var (
	PermissionSysconsoleReadSite  *Permission
	PermissionSysconsoleWriteSite *Permission
)

var (
	PermissionSysconsoleReadSiteCustomization  *Permission
	PermissionSysconsoleWriteSiteCustomization *Permission
)

var (
	PermissionSysconsoleReadSiteLocalization  *Permission
	PermissionSysconsoleWriteSiteLocalization *Permission
)

var (
	PermissionSysconsoleReadSiteUsersAndTeams  *Permission
	PermissionSysconsoleWriteSiteUsersAndTeams *Permission
)

var (
	PermissionSysconsoleReadSiteNotifications  *Permission
	PermissionSysconsoleWriteSiteNotifications *Permission
)

var (
	PermissionSysconsoleReadSiteAnnouncementBanner  *Permission
	PermissionSysconsoleWriteSiteAnnouncementBanner *Permission
)

var (
	PermissionSysconsoleReadSiteEmoji  *Permission
	PermissionSysconsoleWriteSiteEmoji *Permission
)

var (
	PermissionSysconsoleReadSitePosts  *Permission
	PermissionSysconsoleWriteSitePosts *Permission
)

var (
	PermissionSysconsoleReadSiteFileSharingAndDownloads  *Permission
	PermissionSysconsoleWriteSiteFileSharingAndDownloads *Permission
)

var (
	PermissionSysconsoleReadSitePublicLinks  *Permission
	PermissionSysconsoleWriteSitePublicLinks *Permission
)

var (
	PermissionSysconsoleReadSiteNotices  *Permission
	PermissionSysconsoleWriteSiteNotices *Permission
)

var (
	PermissionSysconsoleReadIPFilters  *Permission
	PermissionSysconsoleWriteIPFilters *Permission
)

var (
	PermissionSysconsoleReadAuthentication  *Permission
	PermissionSysconsoleWriteAuthentication *Permission
)

var (
	PermissionSysconsoleReadAuthenticationSignup  *Permission
	PermissionSysconsoleWriteAuthenticationSignup *Permission
)

var (
	PermissionSysconsoleReadAuthenticationEmail  *Permission
	PermissionSysconsoleWriteAuthenticationEmail *Permission
)

var (
	PermissionSysconsoleReadAuthenticationPassword  *Permission
	PermissionSysconsoleWriteAuthenticationPassword *Permission
)

var (
	PermissionSysconsoleReadAuthenticationMfa  *Permission
	PermissionSysconsoleWriteAuthenticationMfa *Permission
)

var (
	PermissionSysconsoleReadAuthenticationLdap  *Permission
	PermissionSysconsoleWriteAuthenticationLdap *Permission
)

var (
	PermissionSysconsoleReadAuthenticationSaml  *Permission
	PermissionSysconsoleWriteAuthenticationSaml *Permission
)

var (
	PermissionSysconsoleReadAuthenticationOpenid  *Permission
	PermissionSysconsoleWriteAuthenticationOpenid *Permission
)

var (
	PermissionSysconsoleReadAuthenticationGuestAccess  *Permission
	PermissionSysconsoleWriteAuthenticationGuestAccess *Permission
)

var (
	PermissionSysconsoleReadPlugins  *Permission
	PermissionSysconsoleWritePlugins *Permission
)

var (
	PermissionSysconsoleReadIntegrations  *Permission
	PermissionSysconsoleWriteIntegrations *Permission
)

var (
	PermissionSysconsoleReadIntegrationsIntegrationManagement  *Permission
	PermissionSysconsoleWriteIntegrationsIntegrationManagement *Permission
)

var (
	PermissionSysconsoleReadIntegrationsBotAccounts  *Permission
	PermissionSysconsoleWriteIntegrationsBotAccounts *Permission
)

var (
	PermissionSysconsoleReadIntegrationsGif  *Permission
	PermissionSysconsoleWriteIntegrationsGif *Permission
)

var (
	PermissionSysconsoleReadIntegrationsCors  *Permission
	PermissionSysconsoleWriteIntegrationsCors *Permission
)

var (
	PermissionSysconsoleReadCompliance  *Permission
	PermissionSysconsoleWriteCompliance *Permission
)

var (
	PermissionSysconsoleReadComplianceDataRetentionPolicy  *Permission
	PermissionSysconsoleWriteComplianceDataRetentionPolicy *Permission
)

var (
	PermissionSysconsoleReadComplianceComplianceExport  *Permission
	PermissionSysconsoleWriteComplianceComplianceExport *Permission
)

var (
	PermissionSysconsoleReadComplianceComplianceMonitoring  *Permission
	PermissionSysconsoleWriteComplianceComplianceMonitoring *Permission
)

var (
	PermissionSysconsoleReadComplianceCustomTermsOfService  *Permission
	PermissionSysconsoleWriteComplianceCustomTermsOfService *Permission
)

var (
	PermissionSysconsoleReadExperimental  *Permission
	PermissionSysconsoleWriteExperimental *Permission
)

var (
	PermissionSysconsoleReadExperimentalFeatures  *Permission
	PermissionSysconsoleWriteExperimentalFeatures *Permission
)

var (
	PermissionSysconsoleReadExperimentalFeatureFlags  *Permission
	PermissionSysconsoleWriteExperimentalFeatureFlags *Permission
)

var (
	PermissionSysconsoleReadExperimentalBleve  *Permission
	PermissionSysconsoleWriteExperimentalBleve *Permission
)

var (
	PermissionPublicPlaybookCreate           *Permission
	PermissionPublicPlaybookManageProperties *Permission
	PermissionPublicPlaybookManageMembers    *Permission
	PermissionPublicPlaybookManageRoles      *Permission
	PermissionPublicPlaybookView             *Permission
	PermissionPublicPlaybookMakePrivate      *Permission
)

var (
	PermissionPrivatePlaybookCreate           *Permission
	PermissionPrivatePlaybookManageProperties *Permission
	PermissionPrivatePlaybookManageMembers    *Permission
	PermissionPrivatePlaybookManageRoles      *Permission
	PermissionPrivatePlaybookView             *Permission
	PermissionPrivatePlaybookMakePublic       *Permission
)

var (
	PermissionRunCreate           *Permission
	PermissionRunManageProperties *Permission
	PermissionRunManageMembers    *Permission
	PermissionRunView             *Permission
)

var (
	PermissionSysconsoleReadProductsBoards  *Permission
	PermissionSysconsoleWriteProductsBoards *Permission
)

// PermissionManageSystem is a general permission that encompasses all system admin functions
// in the future this could be broken up to allow access to some
// admin functions but not others
var PermissionManageSystem *Permission

var (
	PermissionCreateCustomGroup        *Permission
	PermissionManageCustomGroupMembers *Permission
	PermissionEditCustomGroup          *Permission
	PermissionDeleteCustomGroup        *Permission
	PermissionRestoreCustomGroup       *Permission
)

var (
	AllPermissions        []*Permission
	DeprecatedPermissions []*Permission
)

var (
	ChannelModeratedPermissions    []string
	ChannelModeratedPermissionsMap map[string]string
)

var (
	SysconsoleReadPermissions  []*Permission
	SysconsoleWritePermissions []*Permission
)

var (
	PermissionManageOutgoingOAuthConnections *Permission
	ModeratedBookmarkPermissions             []*Permission
)

func initializePermissions() {
	PermissionInviteUser = &Permission{
		"invite_user",
		"authentication.permissions.team_invite_user.name",
		"authentication.permissions.team_invite_user.description",
		PermissionScopeTeam,
	}
	PermissionAddUserToTeam = &Permission{
		"add_user_to_team",
		"authentication.permissions.add_user_to_team.name",
		"authentication.permissions.add_user_to_team.description",
		PermissionScopeTeam,
	}
	PermissionUseSlashCommands = &Permission{
		"use_slash_commands",
		"authentication.permissions.team_use_slash_commands.name",
		"authentication.permissions.team_use_slash_commands.description",
		PermissionScopeChannel,
	}
	PermissionManageSlashCommands = &Permission{
		"manage_slash_commands",
		"authentication.permissions.manage_slash_commands.name",
		"authentication.permissions.manage_slash_commands.description",
		PermissionScopeTeam,
	}
	PermissionManageOthersSlashCommands = &Permission{
		"manage_others_slash_commands",
		"authentication.permissions.manage_others_slash_commands.name",
		"authentication.permissions.manage_others_slash_commands.description",
		PermissionScopeTeam,
	}
	PermissionCreatePublicChannel = &Permission{
		"create_public_channel",
		"authentication.permissions.create_public_channel.name",
		"authentication.permissions.create_public_channel.description",
		PermissionScopeTeam,
	}
	PermissionCreatePrivateChannel = &Permission{
		"create_private_channel",
		"authentication.permissions.create_private_channel.name",
		"authentication.permissions.create_private_channel.description",
		PermissionScopeTeam,
	}
	PermissionManagePublicChannelMembers = &Permission{
		"manage_public_channel_members",
		"authentication.permissions.manage_public_channel_members.name",
		"authentication.permissions.manage_public_channel_members.description",
		PermissionScopeChannel,
	}
	PermissionManagePrivateChannelMembers = &Permission{
		"manage_private_channel_members",
		"authentication.permissions.manage_private_channel_members.name",
		"authentication.permissions.manage_private_channel_members.description",
		PermissionScopeChannel,
	}
	PermissionConvertPublicChannelToPrivate = &Permission{
		"convert_public_channel_to_private",
		"authentication.permissions.convert_public_channel_to_private.name",
		"authentication.permissions.convert_public_channel_to_private.description",
		PermissionScopeChannel,
	}
	PermissionConvertPrivateChannelToPublic = &Permission{
		"convert_private_channel_to_public",
		"authentication.permissions.convert_private_channel_to_public.name",
		"authentication.permissions.convert_private_channel_to_public.description",
		PermissionScopeChannel,
	}
	PermissionAssignSystemAdminRole = &Permission{
		"assign_system_admin_role",
		"authentication.permissions.assign_system_admin_role.name",
		"authentication.permissions.assign_system_admin_role.description",
		PermissionScopeSystem,
	}
	PermissionManageRoles = &Permission{
		"manage_roles",
		"authentication.permissions.manage_roles.name",
		"authentication.permissions.manage_roles.description",
		PermissionScopeSystem,
	}
	PermissionManageTeamRoles = &Permission{
		"manage_team_roles",
		"authentication.permissions.manage_team_roles.name",
		"authentication.permissions.manage_team_roles.description",
		PermissionScopeTeam,
	}
	PermissionManageChannelRoles = &Permission{
		"manage_channel_roles",
		"authentication.permissions.manage_channel_roles.name",
		"authentication.permissions.manage_channel_roles.description",
		PermissionScopeChannel,
	}
	PermissionManageSystem = &Permission{
		"manage_system",
		"authentication.permissions.manage_system.name",
		"authentication.permissions.manage_system.description",
		PermissionScopeSystem,
	}
	PermissionCreateDirectChannel = &Permission{
		"create_direct_channel",
		"authentication.permissions.create_direct_channel.name",
		"authentication.permissions.create_direct_channel.description",
		PermissionScopeSystem,
	}
	PermissionCreateGroupChannel = &Permission{
		"create_group_channel",
		"authentication.permissions.create_group_channel.name",
		"authentication.permissions.create_group_channel.description",
		PermissionScopeSystem,
	}
	PermissionManagePublicChannelProperties = &Permission{
		"manage_public_channel_properties",
		"authentication.permissions.manage_public_channel_properties.name",
		"authentication.permissions.manage_public_channel_properties.description",
		PermissionScopeChannel,
	}
	PermissionManagePrivateChannelProperties = &Permission{
		"manage_private_channel_properties",
		"authentication.permissions.manage_private_channel_properties.name",
		"authentication.permissions.manage_private_channel_properties.description",
		PermissionScopeChannel,
	}
	PermissionListPublicTeams = &Permission{
		"list_public_teams",
		"authentication.permissions.list_public_teams.name",
		"authentication.permissions.list_public_teams.description",
		PermissionScopeSystem,
	}
	PermissionJoinPublicTeams = &Permission{
		"join_public_teams",
		"authentication.permissions.join_public_teams.name",
		"authentication.permissions.join_public_teams.description",
		PermissionScopeSystem,
	}
	PermissionListPrivateTeams = &Permission{
		"list_private_teams",
		"authentication.permissions.list_private_teams.name",
		"authentication.permissions.list_private_teams.description",
		PermissionScopeSystem,
	}
	PermissionJoinPrivateTeams = &Permission{
		"join_private_teams",
		"authentication.permissions.join_private_teams.name",
		"authentication.permissions.join_private_teams.description",
		PermissionScopeSystem,
	}
	PermissionListTeamChannels = &Permission{
		"list_team_channels",
		"authentication.permissions.list_team_channels.name",
		"authentication.permissions.list_team_channels.description",
		PermissionScopeTeam,
	}
	PermissionJoinPublicChannels = &Permission{
		"join_public_channels",
		"authentication.permissions.join_public_channels.name",
		"authentication.permissions.join_public_channels.description",
		PermissionScopeTeam,
	}
	PermissionDeletePublicChannel = &Permission{
		"delete_public_channel",
		"authentication.permissions.delete_public_channel.name",
		"authentication.permissions.delete_public_channel.description",
		PermissionScopeChannel,
	}
	PermissionDeletePrivateChannel = &Permission{
		"delete_private_channel",
		"authentication.permissions.delete_private_channel.name",
		"authentication.permissions.delete_private_channel.description",
		PermissionScopeChannel,
	}
	PermissionEditOtherUsers = &Permission{
		"edit_other_users",
		"authentication.permissions.edit_other_users.name",
		"authentication.permissions.edit_other_users.description",
		PermissionScopeSystem,
	}
	PermissionReadChannel = &Permission{
		"read_channel",
		"authentication.permissions.read_channel.name",
		"authentication.permissions.read_channel.description",
		PermissionScopeChannel,
	}
	PermissionReadChannelContent = &Permission{
		"read_channel_content",
		"authentication.permissions.read_channel_content.name",
		"authentication.permissions.read_channel_content.description",
		PermissionScopeChannel,
	}
	PermissionReadPublicChannelGroups = &Permission{
		"read_public_channel_groups",
		"authentication.permissions.read_public_channel_groups.name",
		"authentication.permissions.read_public_channel_groups.description",
		PermissionScopeChannel,
	}
	PermissionReadPrivateChannelGroups = &Permission{
		"read_private_channel_groups",
		"authentication.permissions.read_private_channel_groups.name",
		"authentication.permissions.read_private_channel_groups.description",
		PermissionScopeChannel,
	}
	PermissionReadPublicChannel = &Permission{
		"read_public_channel",
		"authentication.permissions.read_public_channel.name",
		"authentication.permissions.read_public_channel.description",
		PermissionScopeTeam,
	}
	PermissionAddReaction = &Permission{
		"add_reaction",
		"authentication.permissions.add_reaction.name",
		"authentication.permissions.add_reaction.description",
		PermissionScopeChannel,
	}
	PermissionRemoveReaction = &Permission{
		"remove_reaction",
		"authentication.permissions.remove_reaction.name",
		"authentication.permissions.remove_reaction.description",
		PermissionScopeChannel,
	}
	PermissionRemoveOthersReactions = &Permission{
		"remove_others_reactions",
		"authentication.permissions.remove_others_reactions.name",
		"authentication.permissions.remove_others_reactions.description",
		PermissionScopeChannel,
	}
	// DEPRECATED
	PermissionPermanentDeleteUser = &Permission{
		"permanent_delete_user",
		"authentication.permissions.permanent_delete_user.name",
		"authentication.permissions.permanent_delete_user.description",
		PermissionScopeSystem,
	}
	PermissionUploadFile = &Permission{
		"upload_file",
		"authentication.permissions.upload_file.name",
		"authentication.permissions.upload_file.description",
		PermissionScopeChannel,
	}
	PermissionGetPublicLink = &Permission{
		"get_public_link",
		"authentication.permissions.get_public_link.name",
		"authentication.permissions.get_public_link.description",
		PermissionScopeSystem,
	}
	// DEPRECATED
	PermissionManageWebhooks = &Permission{
		"manage_webhooks",
		"authentication.permissions.manage_webhooks.name",
		"authentication.permissions.manage_webhooks.description",
		PermissionScopeTeam,
	}
	// DEPRECATED
	PermissionManageOthersWebhooks = &Permission{
		"manage_others_webhooks",
		"authentication.permissions.manage_others_webhooks.name",
		"authentication.permissions.manage_others_webhooks.description",
		PermissionScopeTeam,
	}
	PermissionManageIncomingWebhooks = &Permission{
		"manage_incoming_webhooks",
		"authentication.permissions.manage_incoming_webhooks.name",
		"authentication.permissions.manage_incoming_webhooks.description",
		PermissionScopeTeam,
	}
	PermissionManageOutgoingWebhooks = &Permission{
		"manage_outgoing_webhooks",
		"authentication.permissions.manage_outgoing_webhooks.name",
		"authentication.permissions.manage_outgoing_webhooks.description",
		PermissionScopeTeam,
	}
	PermissionManageOthersIncomingWebhooks = &Permission{
		"manage_others_incoming_webhooks",
		"authentication.permissions.manage_others_incoming_webhooks.name",
		"authentication.permissions.manage_others_incoming_webhooks.description",
		PermissionScopeTeam,
	}
	PermissionManageOthersOutgoingWebhooks = &Permission{
		"manage_others_outgoing_webhooks",
		"authentication.permissions.manage_others_outgoing_webhooks.name",
		"authentication.permissions.manage_others_outgoing_webhooks.description",
		PermissionScopeTeam,
	}
	PermissionManageOAuth = &Permission{
		"manage_oauth",
		"authentication.permissions.manage_oauth.name",
		"authentication.permissions.manage_oauth.description",
		PermissionScopeSystem,
	}
	PermissionManageSystemWideOAuth = &Permission{
		"manage_system_wide_oauth",
		"authentication.permissions.manage_system_wide_oauth.name",
		"authentication.permissions.manage_system_wide_oauth.description",
		PermissionScopeSystem,
	}
	// DEPRECATED
	PermissionManageEmojis = &Permission{
		"manage_emojis",
		"authentication.permissions.manage_emojis.name",
		"authentication.permissions.manage_emojis.description",
		PermissionScopeTeam,
	}
	// DEPRECATED
	PermissionManageOthersEmojis = &Permission{
		"manage_others_emojis",
		"authentication.permissions.manage_others_emojis.name",
		"authentication.permissions.manage_others_emojis.description",
		PermissionScopeTeam,
	}
	PermissionCreateEmojis = &Permission{
		"create_emojis",
		"authentication.permissions.create_emojis.name",
		"authentication.permissions.create_emojis.description",
		PermissionScopeTeam,
	}
	PermissionDeleteEmojis = &Permission{
		"delete_emojis",
		"authentication.permissions.delete_emojis.name",
		"authentication.permissions.delete_emojis.description",
		PermissionScopeTeam,
	}
	PermissionDeleteOthersEmojis = &Permission{
		"delete_others_emojis",
		"authentication.permissions.delete_others_emojis.name",
		"authentication.permissions.delete_others_emojis.description",
		PermissionScopeTeam,
	}
	PermissionCreatePost = &Permission{
		"create_post",
		"authentication.permissions.create_post.name",
		"authentication.permissions.create_post.description",
		PermissionScopeChannel,
	}
	PermissionCreatePostPublic = &Permission{
		"create_post_public",
		"authentication.permissions.create_post_public.name",
		"authentication.permissions.create_post_public.description",
		PermissionScopeChannel,
	}
	PermissionCreatePostEphemeral = &Permission{
		"create_post_ephemeral",
		"authentication.permissions.create_post_ephemeral.name",
		"authentication.permissions.create_post_ephemeral.description",
		PermissionScopeChannel,
	}
	PermissionReadDeletedPosts = &Permission{
		"read_deleted_posts",
		"authentication.permissions.read_deleted_posts.name",
		"authentication.permissions.read_deleted_posts.description",
		PermissionScopeChannel,
	}
	PermissionEditPost = &Permission{
		"edit_post",
		"authentication.permissions.edit_post.name",
		"authentication.permissions.edit_post.description",
		PermissionScopeChannel,
	}
	PermissionEditOthersPosts = &Permission{
		"edit_others_posts",
		"authentication.permissions.edit_others_posts.name",
		"authentication.permissions.edit_others_posts.description",
		PermissionScopeChannel,
	}
	PermissionDeletePost = &Permission{
		"delete_post",
		"authentication.permissions.delete_post.name",
		"authentication.permissions.delete_post.description",
		PermissionScopeChannel,
	}
	PermissionDeleteOthersPosts = &Permission{
		"delete_others_posts",
		"authentication.permissions.delete_others_posts.name",
		"authentication.permissions.delete_others_posts.description",
		PermissionScopeChannel,
	}
	PermissionManageSharedChannels = &Permission{
		"manage_shared_channels",
		"authentication.permissions.manage_shared_channels.name",
		"authentication.permissions.manage_shared_channels.description",
		PermissionScopeSystem,
	}
	PermissionManageSecureConnections = &Permission{
		"manage_secure_connections",
		"authentication.permissions.manage_secure_connections.name",
		"authentication.permissions.manage_secure_connections.description",
		PermissionScopeSystem,
	}

	PermissionCreateDataRetentionJob = &Permission{
		"create_data_retention_job",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionManageDataRetentionJob = &Permission{
		"manage_data_retention_job",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionReadDataRetentionJob = &Permission{
		"read_data_retention_job",
		"",
		"",
		PermissionScopeSystem,
	}

	PermissionCreateComplianceExportJob = &Permission{
		"create_compliance_export_job",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionManageComplianceExportJob = &Permission{
		"manage_compliance_export_job",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionReadComplianceExportJob = &Permission{
		"read_compliance_export_job",
		"",
		"",
		PermissionScopeSystem,
	}

	PermissionReadAudits = &Permission{
		"read_audits",
		"",
		"",
		PermissionScopeSystem,
	}

	PermissionPurgeBleveIndexes = &Permission{
		"purge_bleve_indexes",
		"",
		"",
		PermissionScopeSystem,
	}

	PermissionCreatePostBleveIndexesJob = &Permission{
		"create_post_bleve_indexes_job",
		"",
		"",
		PermissionScopeSystem,
	}

	PermissionManagePostBleveIndexesJob = &Permission{
		"manage_post_bleve_indexes_job",
		"",
		"",
		PermissionScopeSystem,
	}

	PermissionCreateLdapSyncJob = &Permission{
		"create_ldap_sync_job",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionManageLdapSyncJob = &Permission{
		"manage_ldap_sync_job",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionReadLdapSyncJob = &Permission{
		"read_ldap_sync_job",
		"",
		"",
		PermissionScopeSystem,
	}

	PermissionTestLdap = &Permission{
		"test_ldap",
		"",
		"",
		PermissionScopeSystem,
	}

	PermissionInvalidateEmailInvite = &Permission{
		"invalidate_email_invite",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionGetSamlMetadataFromIdp = &Permission{
		"get_saml_metadata_from_idp",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionAddSamlPublicCert = &Permission{
		"add_saml_public_cert",
		"",
		"",
		PermissionScopeSystem,
	}

	PermissionAddSamlPrivateCert = &Permission{
		"add_saml_private_cert",
		"",
		"",
		PermissionScopeSystem,
	}

	PermissionAddSamlIdpCert = &Permission{
		"add_saml_idp_cert",
		"",
		"",
		PermissionScopeSystem,
	}

	PermissionRemoveSamlPublicCert = &Permission{
		"remove_saml_public_cert",
		"",
		"",
		PermissionScopeSystem,
	}

	PermissionRemoveSamlPrivateCert = &Permission{
		"remove_saml_private_cert",
		"",
		"",
		PermissionScopeSystem,
	}

	PermissionRemoveSamlIdpCert = &Permission{
		"remove_saml_idp_cert",
		"",
		"",
		PermissionScopeSystem,
	}

	PermissionGetSamlCertStatus = &Permission{
		"get_saml_cert_status",
		"",
		"",
		PermissionScopeSystem,
	}

	PermissionAddLdapPublicCert = &Permission{
		"add_ldap_public_cert",
		"",
		"",
		PermissionScopeSystem,
	}

	PermissionAddLdapPrivateCert = &Permission{
		"add_ldap_private_cert",
		"",
		"",
		PermissionScopeSystem,
	}

	PermissionRemoveLdapPublicCert = &Permission{
		"remove_ldap_public_cert",
		"",
		"",
		PermissionScopeSystem,
	}

	PermissionRemoveLdapPrivateCert = &Permission{
		"remove_ldap_private_cert",
		"",
		"",
		PermissionScopeSystem,
	}

	PermissionGetLogs = &Permission{
		"get_logs",
		"",
		"",
		PermissionScopeSystem,
	}

	PermissionReadLicenseInformation = &Permission{
		"read_license_information",
		"",
		"",
		PermissionScopeSystem,
	}

	PermissionGetAnalytics = &Permission{
		"get_analytics",
		"",
		"",
		PermissionScopeSystem,
	}

	PermissionManageLicenseInformation = &Permission{
		"manage_license_information",
		"",
		"",
		PermissionScopeSystem,
	}

	PermissionDownloadComplianceExportResult = &Permission{
		"download_compliance_export_result",
		"authentication.permissions.download_compliance_export_result.name",
		"authentication.permissions.download_compliance_export_result.description",
		PermissionScopeSystem,
	}

	PermissionTestSiteURL = &Permission{
		"test_site_url",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionTestElasticsearch = &Permission{
		"test_elasticsearch",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionTestS3 = &Permission{
		"test_s3",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionReloadConfig = &Permission{
		"reload_config",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionInvalidateCaches = &Permission{
		"invalidate_caches",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionRecycleDatabaseConnections = &Permission{
		"recycle_database_connections",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionPurgeElasticsearchIndexes = &Permission{
		"purge_elasticsearch_indexes",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionTestEmail = &Permission{
		"test_email",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionCreateElasticsearchPostIndexingJob = &Permission{
		"create_elasticsearch_post_indexing_job",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionManageElasticsearchPostIndexingJob = &Permission{
		"manage_elasticsearch_post_indexing_job",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionCreateElasticsearchPostAggregationJob = &Permission{
		"create_elasticsearch_post_aggregation_job",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionManageElasticsearchPostAggregationJob = &Permission{
		"manage_elasticsearch_post_aggregation_job",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionReadElasticsearchPostIndexingJob = &Permission{
		"read_elasticsearch_post_indexing_job",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionReadElasticsearchPostAggregationJob = &Permission{
		"read_elasticsearch_post_aggregation_job",
		"",
		"",
		PermissionScopeSystem,
	}

	PermissionRemoveUserFromTeam = &Permission{
		"remove_user_from_team",
		"authentication.permissions.remove_user_from_team.name",
		"authentication.permissions.remove_user_from_team.description",
		PermissionScopeTeam,
	}
	PermissionCreateTeam = &Permission{
		"create_team",
		"authentication.permissions.create_team.name",
		"authentication.permissions.create_team.description",
		PermissionScopeSystem,
	}
	PermissionManageTeam = &Permission{
		"manage_team",
		"authentication.permissions.manage_team.name",
		"authentication.permissions.manage_team.description",
		PermissionScopeTeam,
	}
	PermissionImportTeam = &Permission{
		"import_team",
		"authentication.permissions.import_team.name",
		"authentication.permissions.import_team.description",
		PermissionScopeTeam,
	}
	PermissionViewTeam = &Permission{
		"view_team",
		"authentication.permissions.view_team.name",
		"authentication.permissions.view_team.description",
		PermissionScopeTeam,
	}
	PermissionListUsersWithoutTeam = &Permission{
		"list_users_without_team",
		"authentication.permissions.list_users_without_team.name",
		"authentication.permissions.list_users_without_team.description",
		PermissionScopeSystem,
	}
	PermissionCreateUserAccessToken = &Permission{
		"create_user_access_token",
		"authentication.permissions.create_user_access_token.name",
		"authentication.permissions.create_user_access_token.description",
		PermissionScopeSystem,
	}
	PermissionReadUserAccessToken = &Permission{
		"read_user_access_token",
		"authentication.permissions.read_user_access_token.name",
		"authentication.permissions.read_user_access_token.description",
		PermissionScopeSystem,
	}
	PermissionRevokeUserAccessToken = &Permission{
		"revoke_user_access_token",
		"authentication.permissions.revoke_user_access_token.name",
		"authentication.permissions.revoke_user_access_token.description",
		PermissionScopeSystem,
	}
	PermissionCreateBot = &Permission{
		"create_bot",
		"authentication.permissions.create_bot.name",
		"authentication.permissions.create_bot.description",
		PermissionScopeSystem,
	}
	PermissionAssignBot = &Permission{
		"assign_bot",
		"authentication.permissions.assign_bot.name",
		"authentication.permissions.assign_bot.description",
		PermissionScopeSystem,
	}
	PermissionReadBots = &Permission{
		"read_bots",
		"authentication.permissions.read_bots.name",
		"authentication.permissions.read_bots.description",
		PermissionScopeSystem,
	}
	PermissionReadOthersBots = &Permission{
		"read_others_bots",
		"authentication.permissions.read_others_bots.name",
		"authentication.permissions.read_others_bots.description",
		PermissionScopeSystem,
	}
	PermissionManageBots = &Permission{
		"manage_bots",
		"authentication.permissions.manage_bots.name",
		"authentication.permissions.manage_bots.description",
		PermissionScopeSystem,
	}
	PermissionManageOthersBots = &Permission{
		"manage_others_bots",
		"authentication.permissions.manage_others_bots.name",
		"authentication.permissions.manage_others_bots.description",
		PermissionScopeSystem,
	}
	PermissionReadJobs = &Permission{
		"read_jobs",
		"authentication.permisssions.read_jobs.name",
		"authentication.permisssions.read_jobs.description",
		PermissionScopeSystem,
	}
	PermissionManageJobs = &Permission{
		"manage_jobs",
		"authentication.permisssions.manage_jobs.name",
		"authentication.permisssions.manage_jobs.description",
		PermissionScopeSystem,
	}
	PermissionViewMembers = &Permission{
		"view_members",
		"authentication.permisssions.view_members.name",
		"authentication.permisssions.view_members.description",
		PermissionScopeTeam,
	}
	PermissionInviteGuest = &Permission{
		"invite_guest",
		"authentication.permissions.invite_guest.name",
		"authentication.permissions.invite_guest.description",
		PermissionScopeTeam,
	}
	PermissionPromoteGuest = &Permission{
		"promote_guest",
		"authentication.permissions.promote_guest.name",
		"authentication.permissions.promote_guest.description",
		PermissionScopeSystem,
	}
	PermissionDemoteToGuest = &Permission{
		"demote_to_guest",
		"authentication.permissions.demote_to_guest.name",
		"authentication.permissions.demote_to_guest.description",
		PermissionScopeSystem,
	}
	PermissionUseChannelMentions = &Permission{
		"use_channel_mentions",
		"authentication.permissions.use_channel_mentions.name",
		"authentication.permissions.use_channel_mentions.description",
		PermissionScopeChannel,
	}
	PermissionUseGroupMentions = &Permission{
		"use_group_mentions",
		"authentication.permissions.use_group_mentions.name",
		"authentication.permissions.use_group_mentions.description",
		PermissionScopeChannel,
	}

	// Channel bookmarks
	PermissionAddBookmarkPublicChannel = &Permission{
		"add_bookmark_public_channel",
		"",
		"",
		PermissionScopeChannel,
	}
	PermissionEditBookmarkPublicChannel = &Permission{
		"edit_bookmark_public_channel",
		"",
		"",
		PermissionScopeChannel,
	}
	PermissionDeleteBookmarkPublicChannel = &Permission{
		"delete_bookmark_public_channel",
		"",
		"",
		PermissionScopeChannel,
	}
	PermissionOrderBookmarkPublicChannel = &Permission{
		"order_bookmark_public_channel",
		"",
		"",
		PermissionScopeChannel,
	}
	PermissionAddBookmarkPrivateChannel = &Permission{
		"add_bookmark_private_channel",
		"",
		"",
		PermissionScopeChannel,
	}
	PermissionEditBookmarkPrivateChannel = &Permission{
		"edit_bookmark_private_channel",
		"",
		"",
		PermissionScopeChannel,
	}
	PermissionDeleteBookmarkPrivateChannel = &Permission{
		"delete_bookmark_private_channel",
		"",
		"",
		PermissionScopeChannel,
	}
	PermissionOrderBookmarkPrivateChannel = &Permission{
		"order_bookmark_private_channel",
		"",
		"",
		PermissionScopeChannel,
	}

	PermissionReadOtherUsersTeams = &Permission{
		"read_other_users_teams",
		"authentication.permissions.read_other_users_teams.name",
		"authentication.permissions.read_other_users_teams.description",
		PermissionScopeSystem,
	}
	PermissionEditBrand = &Permission{
		"edit_brand",
		"authentication.permissions.edit_brand.name",
		"authentication.permissions.edit_brand.description",
		PermissionScopeSystem,
	}
	// DEPRECATED
	PermissionSysconsoleReadAbout = &Permission{
		"sysconsole_read_about",
		"authentication.permissions.use_group_mentions.name",
		"authentication.permissions.use_group_mentions.description",
		PermissionScopeSystem,
	}
	// DEPRECATED
	PermissionSysconsoleWriteAbout = &Permission{
		"sysconsole_write_about",
		"authentication.permissions.use_group_mentions.name",
		"authentication.permissions.use_group_mentions.description",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadAboutEditionAndLicense = &Permission{
		"sysconsole_read_about_edition_and_license",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteAboutEditionAndLicense = &Permission{
		"sysconsole_write_about_edition_and_license",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadBilling = &Permission{
		"sysconsole_read_billing",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteBilling = &Permission{
		"sysconsole_write_billing",
		"",
		"",
		PermissionScopeSystem,
	}
	// DEPRECATED
	PermissionSysconsoleReadReporting = &Permission{
		"sysconsole_read_reporting",
		"authentication.permissions.use_group_mentions.name",
		"authentication.permissions.use_group_mentions.description",
		PermissionScopeSystem,
	}
	// DEPRECATED
	PermissionSysconsoleWriteReporting = &Permission{
		"sysconsole_write_reporting",
		"authentication.permissions.use_group_mentions.name",
		"authentication.permissions.use_group_mentions.description",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadReportingSiteStatistics = &Permission{
		"sysconsole_read_reporting_site_statistics",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteReportingSiteStatistics = &Permission{
		"sysconsole_write_reporting_site_statistics",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadReportingTeamStatistics = &Permission{
		"sysconsole_read_reporting_team_statistics",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteReportingTeamStatistics = &Permission{
		"sysconsole_write_reporting_team_statistics",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadReportingServerLogs = &Permission{
		"sysconsole_read_reporting_server_logs",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteReportingServerLogs = &Permission{
		"sysconsole_write_reporting_server_logs",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadUserManagementUsers = &Permission{
		"sysconsole_read_user_management_users",
		"authentication.permissions.use_group_mentions.name",
		"authentication.permissions.use_group_mentions.description",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteUserManagementUsers = &Permission{
		"sysconsole_write_user_management_users",
		"authentication.permissions.use_group_mentions.name",
		"authentication.permissions.use_group_mentions.description",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadUserManagementGroups = &Permission{
		"sysconsole_read_user_management_groups",
		"authentication.permissions.use_group_mentions.name",
		"authentication.permissions.use_group_mentions.description",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteUserManagementGroups = &Permission{
		"sysconsole_write_user_management_groups",
		"authentication.permissions.use_group_mentions.name",
		"authentication.permissions.use_group_mentions.description",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadUserManagementTeams = &Permission{
		"sysconsole_read_user_management_teams",
		"authentication.permissions.use_group_mentions.name",
		"authentication.permissions.use_group_mentions.description",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteUserManagementTeams = &Permission{
		"sysconsole_write_user_management_teams",
		"authentication.permissions.use_group_mentions.name",
		"authentication.permissions.use_group_mentions.description",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadUserManagementChannels = &Permission{
		"sysconsole_read_user_management_channels",
		"authentication.permissions.use_group_mentions.name",
		"authentication.permissions.use_group_mentions.description",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteUserManagementChannels = &Permission{
		"sysconsole_write_user_management_channels",
		"authentication.permissions.use_group_mentions.name",
		"authentication.permissions.use_group_mentions.description",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadUserManagementPermissions = &Permission{
		"sysconsole_read_user_management_permissions",
		"authentication.permissions.use_group_mentions.name",
		"authentication.permissions.use_group_mentions.description",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteUserManagementPermissions = &Permission{
		"sysconsole_write_user_management_permissions",
		"authentication.permissions.use_group_mentions.name",
		"authentication.permissions.use_group_mentions.description",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadUserManagementSystemRoles = &Permission{
		"sysconsole_read_user_management_system_roles",
		"authentication.permissions.use_group_mentions.name",
		"authentication.permissions.use_group_mentions.description",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteUserManagementSystemRoles = &Permission{
		"sysconsole_write_user_management_system_roles",
		"authentication.permissions.use_group_mentions.name",
		"authentication.permissions.use_group_mentions.description",
		PermissionScopeSystem,
	}
	// DEPRECATED
	PermissionSysconsoleReadEnvironment = &Permission{
		"sysconsole_read_environment",
		"authentication.permissions.use_group_mentions.name",
		"authentication.permissions.use_group_mentions.description",
		PermissionScopeSystem,
	}
	// DEPRECATED
	PermissionSysconsoleWriteEnvironment = &Permission{
		"sysconsole_write_environment",
		"authentication.permissions.use_group_mentions.name",
		"authentication.permissions.use_group_mentions.description",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadEnvironmentWebServer = &Permission{
		"sysconsole_read_environment_web_server",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteEnvironmentWebServer = &Permission{
		"sysconsole_write_environment_web_server",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadEnvironmentDatabase = &Permission{
		"sysconsole_read_environment_database",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteEnvironmentDatabase = &Permission{
		"sysconsole_write_environment_database",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadEnvironmentElasticsearch = &Permission{
		"sysconsole_read_environment_elasticsearch",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteEnvironmentElasticsearch = &Permission{
		"sysconsole_write_environment_elasticsearch",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadEnvironmentFileStorage = &Permission{
		"sysconsole_read_environment_file_storage",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteEnvironmentFileStorage = &Permission{
		"sysconsole_write_environment_file_storage",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadEnvironmentImageProxy = &Permission{
		"sysconsole_read_environment_image_proxy",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteEnvironmentImageProxy = &Permission{
		"sysconsole_write_environment_image_proxy",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadEnvironmentSMTP = &Permission{
		"sysconsole_read_environment_smtp",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteEnvironmentSMTP = &Permission{
		"sysconsole_write_environment_smtp",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadEnvironmentPushNotificationServer = &Permission{
		"sysconsole_read_environment_push_notification_server",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteEnvironmentPushNotificationServer = &Permission{
		"sysconsole_write_environment_push_notification_server",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadEnvironmentHighAvailability = &Permission{
		"sysconsole_read_environment_high_availability",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteEnvironmentHighAvailability = &Permission{
		"sysconsole_write_environment_high_availability",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadEnvironmentRateLimiting = &Permission{
		"sysconsole_read_environment_rate_limiting",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteEnvironmentRateLimiting = &Permission{
		"sysconsole_write_environment_rate_limiting",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadEnvironmentLogging = &Permission{
		"sysconsole_read_environment_logging",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteEnvironmentLogging = &Permission{
		"sysconsole_write_environment_logging",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadEnvironmentSessionLengths = &Permission{
		"sysconsole_read_environment_session_lengths",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteEnvironmentSessionLengths = &Permission{
		"sysconsole_write_environment_session_lengths",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadEnvironmentPerformanceMonitoring = &Permission{
		"sysconsole_read_environment_performance_monitoring",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteEnvironmentPerformanceMonitoring = &Permission{
		"sysconsole_write_environment_performance_monitoring",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadEnvironmentDeveloper = &Permission{
		"sysconsole_read_environment_developer",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteEnvironmentDeveloper = &Permission{
		"sysconsole_write_environment_developer",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadEnvironmentMobileSecurity = &Permission{
		"sysconsole_read_environment_mobile_security",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteEnvironmentMobileSecurity = &Permission{
		"sysconsole_write_environment_mobile_security",
		"",
		"",
		PermissionScopeSystem,
	}
	// DEPRECATED
	PermissionSysconsoleReadSite = &Permission{
		"sysconsole_read_site",
		"authentication.permissions.use_group_mentions.name",
		"authentication.permissions.use_group_mentions.description",
		PermissionScopeSystem,
	}
	// DEPRECATED
	PermissionSysconsoleWriteSite = &Permission{
		"sysconsole_write_site",
		"authentication.permissions.use_group_mentions.name",
		"authentication.permissions.use_group_mentions.description",
		PermissionScopeSystem,
	}

	PermissionSysconsoleReadSiteCustomization = &Permission{
		"sysconsole_read_site_customization",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteSiteCustomization = &Permission{
		"sysconsole_write_site_customization",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadSiteLocalization = &Permission{
		"sysconsole_read_site_localization",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteSiteLocalization = &Permission{
		"sysconsole_write_site_localization",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadSiteUsersAndTeams = &Permission{
		"sysconsole_read_site_users_and_teams",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteSiteUsersAndTeams = &Permission{
		"sysconsole_write_site_users_and_teams",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadSiteNotifications = &Permission{
		"sysconsole_read_site_notifications",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteSiteNotifications = &Permission{
		"sysconsole_write_site_notifications",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadSiteAnnouncementBanner = &Permission{
		"sysconsole_read_site_announcement_banner",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteSiteAnnouncementBanner = &Permission{
		"sysconsole_write_site_announcement_banner",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadSiteEmoji = &Permission{
		"sysconsole_read_site_emoji",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteSiteEmoji = &Permission{
		"sysconsole_write_site_emoji",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadSitePosts = &Permission{
		"sysconsole_read_site_posts",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteSitePosts = &Permission{
		"sysconsole_write_site_posts",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadSiteFileSharingAndDownloads = &Permission{
		"sysconsole_read_site_file_sharing_and_downloads",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteSiteFileSharingAndDownloads = &Permission{
		"sysconsole_write_site_file_sharing_and_downloads",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadSitePublicLinks = &Permission{
		"sysconsole_read_site_public_links",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteSitePublicLinks = &Permission{
		"sysconsole_write_site_public_links",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadSiteNotices = &Permission{
		"sysconsole_read_site_notices",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteSiteNotices = &Permission{
		"sysconsole_write_site_notices",
		"",
		"",
		PermissionScopeSystem,
	}

	PermissionSysconsoleReadIPFilters = &Permission{
		"sysconsole_read_site_ip_filters",
		"",
		"",
		PermissionScopeSystem,
	}

	PermissionSysconsoleWriteIPFilters = &Permission{
		"sysconsole_write_site_ip_filters",
		"",
		"",
		PermissionScopeSystem,
	}

	// Deprecated
	PermissionSysconsoleReadAuthentication = &Permission{
		"sysconsole_read_authentication",
		"authentication.permissions.use_group_mentions.name",
		"authentication.permissions.use_group_mentions.description",
		PermissionScopeSystem,
	}
	// Deprecated
	PermissionSysconsoleWriteAuthentication = &Permission{
		"sysconsole_write_authentication",
		"authentication.permissions.use_group_mentions.name",
		"authentication.permissions.use_group_mentions.description",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadAuthenticationSignup = &Permission{
		"sysconsole_read_authentication_signup",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteAuthenticationSignup = &Permission{
		"sysconsole_write_authentication_signup",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadAuthenticationEmail = &Permission{
		"sysconsole_read_authentication_email",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteAuthenticationEmail = &Permission{
		"sysconsole_write_authentication_email",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadAuthenticationPassword = &Permission{
		"sysconsole_read_authentication_password",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteAuthenticationPassword = &Permission{
		"sysconsole_write_authentication_password",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadAuthenticationMfa = &Permission{
		"sysconsole_read_authentication_mfa",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteAuthenticationMfa = &Permission{
		"sysconsole_write_authentication_mfa",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadAuthenticationLdap = &Permission{
		"sysconsole_read_authentication_ldap",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteAuthenticationLdap = &Permission{
		"sysconsole_write_authentication_ldap",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadAuthenticationSaml = &Permission{
		"sysconsole_read_authentication_saml",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteAuthenticationSaml = &Permission{
		"sysconsole_write_authentication_saml",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadAuthenticationOpenid = &Permission{
		"sysconsole_read_authentication_openid",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteAuthenticationOpenid = &Permission{
		"sysconsole_write_authentication_openid",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadAuthenticationGuestAccess = &Permission{
		"sysconsole_read_authentication_guest_access",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteAuthenticationGuestAccess = &Permission{
		"sysconsole_write_authentication_guest_access",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadPlugins = &Permission{
		"sysconsole_read_plugins",
		"authentication.permissions.use_group_mentions.name",
		"authentication.permissions.use_group_mentions.description",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWritePlugins = &Permission{
		"sysconsole_write_plugins",
		"authentication.permissions.use_group_mentions.name",
		"authentication.permissions.use_group_mentions.description",
		PermissionScopeSystem,
	}
	// DEPRECATED
	PermissionSysconsoleReadIntegrations = &Permission{
		"sysconsole_read_integrations",
		"authentication.permissions.use_group_mentions.name",
		"authentication.permissions.use_group_mentions.description",
		PermissionScopeSystem,
	}
	// DEPRECATED
	PermissionSysconsoleWriteIntegrations = &Permission{
		"sysconsole_write_integrations",
		"authentication.permissions.use_group_mentions.name",
		"authentication.permissions.use_group_mentions.description",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadIntegrationsIntegrationManagement = &Permission{
		"sysconsole_read_integrations_integration_management",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteIntegrationsIntegrationManagement = &Permission{
		"sysconsole_write_integrations_integration_management",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadIntegrationsBotAccounts = &Permission{
		"sysconsole_read_integrations_bot_accounts",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteIntegrationsBotAccounts = &Permission{
		"sysconsole_write_integrations_bot_accounts",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadIntegrationsGif = &Permission{
		"sysconsole_read_integrations_gif",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteIntegrationsGif = &Permission{
		"sysconsole_write_integrations_gif",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadIntegrationsCors = &Permission{
		"sysconsole_read_integrations_cors",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteIntegrationsCors = &Permission{
		"sysconsole_write_integrations_cors",
		"",
		"",
		PermissionScopeSystem,
	}
	// DEPRECATED
	PermissionSysconsoleReadCompliance = &Permission{
		"sysconsole_read_compliance",
		"authentication.permissions.use_group_mentions.name",
		"authentication.permissions.use_group_mentions.description",
		PermissionScopeSystem,
	}
	// DEPRECATED
	PermissionSysconsoleWriteCompliance = &Permission{
		"sysconsole_write_compliance",
		"authentication.permissions.use_group_mentions.name",
		"authentication.permissions.use_group_mentions.description",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadComplianceDataRetentionPolicy = &Permission{
		"sysconsole_read_compliance_data_retention_policy",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteComplianceDataRetentionPolicy = &Permission{
		"sysconsole_write_compliance_data_retention_policy",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadComplianceComplianceExport = &Permission{
		"sysconsole_read_compliance_compliance_export",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteComplianceComplianceExport = &Permission{
		"sysconsole_write_compliance_compliance_export",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadComplianceComplianceMonitoring = &Permission{
		"sysconsole_read_compliance_compliance_monitoring",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteComplianceComplianceMonitoring = &Permission{
		"sysconsole_write_compliance_compliance_monitoring",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadComplianceCustomTermsOfService = &Permission{
		"sysconsole_read_compliance_custom_terms_of_service",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteComplianceCustomTermsOfService = &Permission{
		"sysconsole_write_compliance_custom_terms_of_service",
		"",
		"",
		PermissionScopeSystem,
	}
	// DEPRECATED
	PermissionSysconsoleReadExperimental = &Permission{
		"sysconsole_read_experimental",
		"authentication.permissions.use_group_mentions.name",
		"authentication.permissions.use_group_mentions.description",
		PermissionScopeSystem,
	}
	// DEPRECATED
	PermissionSysconsoleWriteExperimental = &Permission{
		"sysconsole_write_experimental",
		"authentication.permissions.use_group_mentions.name",
		"authentication.permissions.use_group_mentions.description",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadExperimentalFeatures = &Permission{
		"sysconsole_read_experimental_features",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteExperimentalFeatures = &Permission{
		"sysconsole_write_experimental_features",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadExperimentalFeatureFlags = &Permission{
		"sysconsole_read_experimental_feature_flags",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteExperimentalFeatureFlags = &Permission{
		"sysconsole_write_experimental_feature_flags",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleReadExperimentalBleve = &Permission{
		"sysconsole_read_experimental_bleve",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteExperimentalBleve = &Permission{
		"sysconsole_write_experimental_bleve",
		"",
		"",
		PermissionScopeSystem,
	}

	PermissionCreateCustomGroup = &Permission{
		"create_custom_group",
		"authentication.permissions.create_custom_group.name",
		"authentication.permissions.create_custom_group.description",
		PermissionScopeSystem,
	}

	PermissionManageCustomGroupMembers = &Permission{
		"manage_custom_group_members",
		"authentication.permissions.manage_custom_group_members.name",
		"authentication.permissions.manage_custom_group_members.description",
		PermissionScopeGroup,
	}

	PermissionEditCustomGroup = &Permission{
		"edit_custom_group",
		"authentication.permissions.edit_custom_group.name",
		"authentication.permissions.edit_custom_group.description",
		PermissionScopeGroup,
	}

	PermissionDeleteCustomGroup = &Permission{
		"delete_custom_group",
		"authentication.permissions.delete_custom_group.name",
		"authentication.permissions.delete_custom_group.description",
		PermissionScopeGroup,
	}

	PermissionRestoreCustomGroup = &Permission{
		"restore_custom_group",
		"authentication.permissions.restore_custom_group.name",
		"authentication.permissions.restore_custom_group.description",
		PermissionScopeGroup,
	}

	// Playbooks
	PermissionPublicPlaybookCreate = &Permission{
		"playbook_public_create",
		"",
		"",
		PermissionScopeTeam,
	}

	PermissionPublicPlaybookManageProperties = &Permission{
		"playbook_public_manage_properties",
		"",
		"",
		PermissionScopePlaybook,
	}

	PermissionPublicPlaybookManageMembers = &Permission{
		"playbook_public_manage_members",
		"",
		"",
		PermissionScopePlaybook,
	}

	PermissionPublicPlaybookManageRoles = &Permission{
		"playbook_public_manage_roles",
		"",
		"",
		PermissionScopePlaybook,
	}

	PermissionPublicPlaybookView = &Permission{
		"playbook_public_view",
		"",
		"",
		PermissionScopePlaybook,
	}

	PermissionPublicPlaybookMakePrivate = &Permission{
		"playbook_public_make_private",
		"",
		"",
		PermissionScopePlaybook,
	}

	PermissionPrivatePlaybookCreate = &Permission{
		"playbook_private_create",
		"",
		"",
		PermissionScopeTeam,
	}

	PermissionPrivatePlaybookManageProperties = &Permission{
		"playbook_private_manage_properties",
		"",
		"",
		PermissionScopePlaybook,
	}

	PermissionPrivatePlaybookManageMembers = &Permission{
		"playbook_private_manage_members",
		"",
		"",
		PermissionScopePlaybook,
	}

	PermissionPrivatePlaybookManageRoles = &Permission{
		"playbook_private_manage_roles",
		"",
		"",
		PermissionScopePlaybook,
	}

	PermissionPrivatePlaybookView = &Permission{
		"playbook_private_view",
		"",
		"",
		PermissionScopePlaybook,
	}

	PermissionPrivatePlaybookMakePublic = &Permission{
		"playbook_private_make_public",
		"",
		"",
		PermissionScopePlaybook,
	}

	PermissionRunCreate = &Permission{
		"run_create",
		"",
		"",
		PermissionScopePlaybook,
	}

	PermissionRunManageProperties = &Permission{
		"run_manage_properties",
		"",
		"",
		PermissionScopeRun,
	}

	PermissionRunManageMembers = &Permission{
		"run_manage_members",
		"",
		"",
		PermissionScopeRun,
	}

	PermissionRunView = &Permission{
		"run_view",
		"",
		"",
		PermissionScopeRun,
	}

	PermissionSysconsoleReadProductsBoards = &Permission{
		"sysconsole_read_products_boards",
		"",
		"",
		PermissionScopeSystem,
	}
	PermissionSysconsoleWriteProductsBoards = &Permission{
		"sysconsole_write_products_boards",
		"",
		"",
		PermissionScopeSystem,
	}

	PermissionManageOutgoingOAuthConnections = &Permission{
		"manage_outgoing_oauth_connections",
		"authentication.permissions.manage_outgoing_oauth_connections.name",
		"authentication.permissions.manage_outgoing_oauth_connections.description",
		PermissionScopeSystem,
	}

	SysconsoleReadPermissions = []*Permission{
		PermissionSysconsoleReadAboutEditionAndLicense,
		PermissionSysconsoleReadBilling,
		PermissionSysconsoleReadReportingSiteStatistics,
		PermissionSysconsoleReadReportingTeamStatistics,
		PermissionSysconsoleReadReportingServerLogs,
		PermissionSysconsoleReadUserManagementUsers,
		PermissionSysconsoleReadUserManagementGroups,
		PermissionSysconsoleReadUserManagementTeams,
		PermissionSysconsoleReadUserManagementChannels,
		PermissionSysconsoleReadUserManagementPermissions,
		PermissionSysconsoleReadUserManagementSystemRoles,
		PermissionSysconsoleReadEnvironmentWebServer,
		PermissionSysconsoleReadEnvironmentDatabase,
		PermissionSysconsoleReadEnvironmentElasticsearch,
		PermissionSysconsoleReadEnvironmentFileStorage,
		PermissionSysconsoleReadEnvironmentImageProxy,
		PermissionSysconsoleReadEnvironmentSMTP,
		PermissionSysconsoleReadEnvironmentPushNotificationServer,
		PermissionSysconsoleReadEnvironmentHighAvailability,
		PermissionSysconsoleReadEnvironmentRateLimiting,
		PermissionSysconsoleReadEnvironmentLogging,
		PermissionSysconsoleReadEnvironmentSessionLengths,
		PermissionSysconsoleReadEnvironmentPerformanceMonitoring,
		PermissionSysconsoleReadEnvironmentDeveloper,
		PermissionSysconsoleReadEnvironmentMobileSecurity,
		PermissionSysconsoleReadSiteCustomization,
		PermissionSysconsoleReadSiteLocalization,
		PermissionSysconsoleReadSiteUsersAndTeams,
		PermissionSysconsoleReadSiteNotifications,
		PermissionSysconsoleReadSiteAnnouncementBanner,
		PermissionSysconsoleReadSiteEmoji,
		PermissionSysconsoleReadSitePosts,
		PermissionSysconsoleReadSiteFileSharingAndDownloads,
		PermissionSysconsoleReadSitePublicLinks,
		PermissionSysconsoleReadSiteNotices,
		PermissionSysconsoleReadAuthenticationSignup,
		PermissionSysconsoleReadAuthenticationEmail,
		PermissionSysconsoleReadAuthenticationPassword,
		PermissionSysconsoleReadAuthenticationMfa,
		PermissionSysconsoleReadAuthenticationLdap,
		PermissionSysconsoleReadAuthenticationSaml,
		PermissionSysconsoleReadAuthenticationOpenid,
		PermissionSysconsoleReadAuthenticationGuestAccess,
		PermissionSysconsoleReadPlugins,
		PermissionSysconsoleReadIntegrationsIntegrationManagement,
		PermissionSysconsoleReadIntegrationsBotAccounts,
		PermissionSysconsoleReadIntegrationsGif,
		PermissionSysconsoleReadIntegrationsCors,
		PermissionSysconsoleReadComplianceDataRetentionPolicy,
		PermissionSysconsoleReadComplianceComplianceExport,
		PermissionSysconsoleReadComplianceComplianceMonitoring,
		PermissionSysconsoleReadComplianceCustomTermsOfService,
		PermissionSysconsoleReadExperimentalFeatures,
		PermissionSysconsoleReadExperimentalFeatureFlags,
		PermissionSysconsoleReadExperimentalBleve,
		PermissionSysconsoleReadProductsBoards,
		PermissionSysconsoleReadIPFilters,
	}

	SysconsoleWritePermissions = []*Permission{
		PermissionSysconsoleWriteAboutEditionAndLicense,
		PermissionSysconsoleWriteBilling,
		PermissionSysconsoleWriteReportingSiteStatistics,
		PermissionSysconsoleWriteReportingTeamStatistics,
		PermissionSysconsoleWriteReportingServerLogs,
		PermissionSysconsoleWriteUserManagementUsers,
		PermissionSysconsoleWriteUserManagementGroups,
		PermissionSysconsoleWriteUserManagementTeams,
		PermissionSysconsoleWriteUserManagementChannels,
		PermissionSysconsoleWriteUserManagementPermissions,
		PermissionSysconsoleWriteUserManagementSystemRoles,
		PermissionSysconsoleWriteEnvironmentWebServer,
		PermissionSysconsoleWriteEnvironmentDatabase,
		PermissionSysconsoleWriteEnvironmentElasticsearch,
		PermissionSysconsoleWriteEnvironmentFileStorage,
		PermissionSysconsoleWriteEnvironmentImageProxy,
		PermissionSysconsoleWriteEnvironmentSMTP,
		PermissionSysconsoleWriteEnvironmentPushNotificationServer,
		PermissionSysconsoleWriteEnvironmentHighAvailability,
		PermissionSysconsoleWriteEnvironmentRateLimiting,
		PermissionSysconsoleWriteEnvironmentLogging,
		PermissionSysconsoleWriteEnvironmentSessionLengths,
		PermissionSysconsoleWriteEnvironmentPerformanceMonitoring,
		PermissionSysconsoleWriteEnvironmentDeveloper,
		PermissionSysconsoleWriteEnvironmentMobileSecurity,
		PermissionSysconsoleWriteSiteCustomization,
		PermissionSysconsoleWriteSiteLocalization,
		PermissionSysconsoleWriteSiteUsersAndTeams,
		PermissionSysconsoleWriteSiteNotifications,
		PermissionSysconsoleWriteSiteAnnouncementBanner,
		PermissionSysconsoleWriteSiteEmoji,
		PermissionSysconsoleWriteSitePosts,
		PermissionSysconsoleWriteSiteFileSharingAndDownloads,
		PermissionSysconsoleWriteSitePublicLinks,
		PermissionSysconsoleWriteSiteNotices,
		PermissionSysconsoleWriteAuthenticationSignup,
		PermissionSysconsoleWriteAuthenticationEmail,
		PermissionSysconsoleWriteAuthenticationPassword,
		PermissionSysconsoleWriteAuthenticationMfa,
		PermissionSysconsoleWriteAuthenticationLdap,
		PermissionSysconsoleWriteAuthenticationSaml,
		PermissionSysconsoleWriteAuthenticationOpenid,
		PermissionSysconsoleWriteAuthenticationGuestAccess,
		PermissionSysconsoleWritePlugins,
		PermissionSysconsoleWriteIntegrationsIntegrationManagement,
		PermissionSysconsoleWriteIntegrationsBotAccounts,
		PermissionSysconsoleWriteIntegrationsGif,
		PermissionSysconsoleWriteIntegrationsCors,
		PermissionSysconsoleWriteComplianceDataRetentionPolicy,
		PermissionSysconsoleWriteComplianceComplianceExport,
		PermissionSysconsoleWriteComplianceComplianceMonitoring,
		PermissionSysconsoleWriteComplianceCustomTermsOfService,
		PermissionSysconsoleWriteExperimentalFeatures,
		PermissionSysconsoleWriteExperimentalFeatureFlags,
		PermissionSysconsoleWriteExperimentalBleve,
		PermissionSysconsoleWriteProductsBoards,
		PermissionSysconsoleWriteIPFilters,
	}

	SystemScopedPermissionsMinusSysconsole := []*Permission{
		PermissionAssignSystemAdminRole,
		PermissionManageRoles,
		PermissionManageSystem,
		PermissionCreateDirectChannel,
		PermissionCreateGroupChannel,
		PermissionListPublicTeams,
		PermissionJoinPublicTeams,
		PermissionListPrivateTeams,
		PermissionJoinPrivateTeams,
		PermissionEditOtherUsers,
		PermissionReadOtherUsersTeams,
		PermissionGetPublicLink,
		PermissionManageOAuth,
		PermissionManageSystemWideOAuth,
		PermissionCreateTeam,
		PermissionListUsersWithoutTeam,
		PermissionCreateUserAccessToken,
		PermissionReadUserAccessToken,
		PermissionRevokeUserAccessToken,
		PermissionCreateBot,
		PermissionAssignBot,
		PermissionReadBots,
		PermissionReadOthersBots,
		PermissionManageBots,
		PermissionManageOthersBots,
		PermissionReadJobs,
		PermissionManageJobs,
		PermissionPromoteGuest,
		PermissionDemoteToGuest,
		PermissionEditBrand,
		PermissionManageSharedChannels,
		PermissionManageSecureConnections,
		PermissionDownloadComplianceExportResult,
		PermissionCreateDataRetentionJob,
		PermissionManageDataRetentionJob,
		PermissionReadDataRetentionJob,
		PermissionCreateComplianceExportJob,
		PermissionManageComplianceExportJob,
		PermissionReadComplianceExportJob,
		PermissionReadAudits,
		PermissionTestSiteURL,
		PermissionTestElasticsearch,
		PermissionTestS3,
		PermissionReloadConfig,
		PermissionInvalidateCaches,
		PermissionRecycleDatabaseConnections,
		PermissionPurgeElasticsearchIndexes,
		PermissionTestEmail,
		PermissionCreateElasticsearchPostIndexingJob,
		PermissionManageElasticsearchPostIndexingJob,
		PermissionCreateElasticsearchPostAggregationJob,
		PermissionManageElasticsearchPostAggregationJob,
		PermissionReadElasticsearchPostIndexingJob,
		PermissionReadElasticsearchPostAggregationJob,
		PermissionPurgeBleveIndexes,
		PermissionCreatePostBleveIndexesJob,
		PermissionManagePostBleveIndexesJob,
		PermissionCreateLdapSyncJob,
		PermissionManageLdapSyncJob,
		PermissionReadLdapSyncJob,
		PermissionTestLdap,
		PermissionInvalidateEmailInvite,
		PermissionGetSamlMetadataFromIdp,
		PermissionAddSamlPublicCert,
		PermissionAddSamlPrivateCert,
		PermissionAddSamlIdpCert,
		PermissionRemoveSamlPublicCert,
		PermissionRemoveSamlPrivateCert,
		PermissionRemoveSamlIdpCert,
		PermissionGetSamlCertStatus,
		PermissionAddLdapPublicCert,
		PermissionAddLdapPrivateCert,
		PermissionRemoveLdapPublicCert,
		PermissionRemoveLdapPrivateCert,
		PermissionGetAnalytics,
		PermissionGetLogs,
		PermissionReadLicenseInformation,
		PermissionManageLicenseInformation,
		PermissionCreateCustomGroup,
		PermissionManageOutgoingOAuthConnections,
	}

	TeamScopedPermissions := []*Permission{
		PermissionInviteUser,
		PermissionAddUserToTeam,
		PermissionManageSlashCommands,
		PermissionManageOthersSlashCommands,
		PermissionCreatePublicChannel,
		PermissionCreatePrivateChannel,
		PermissionManageTeamRoles,
		PermissionListTeamChannels,
		PermissionJoinPublicChannels,
		PermissionReadPublicChannel,
		PermissionManageIncomingWebhooks,
		PermissionManageOutgoingWebhooks,
		PermissionManageOthersIncomingWebhooks,
		PermissionManageOthersOutgoingWebhooks,
		PermissionCreateEmojis,
		PermissionDeleteEmojis,
		PermissionDeleteOthersEmojis,
		PermissionRemoveUserFromTeam,
		PermissionManageTeam,
		PermissionImportTeam,
		PermissionViewTeam,
		PermissionViewMembers,
		PermissionInviteGuest,
		PermissionPublicPlaybookCreate,
		PermissionPrivatePlaybookCreate,
	}

	ChannelScopedPermissions := []*Permission{
		PermissionUseSlashCommands,
		PermissionManagePublicChannelMembers,
		PermissionManagePrivateChannelMembers,
		PermissionManageChannelRoles,
		PermissionManagePublicChannelProperties,
		PermissionManagePrivateChannelProperties,
		PermissionConvertPublicChannelToPrivate,
		PermissionConvertPrivateChannelToPublic,
		PermissionDeletePublicChannel,
		PermissionDeletePrivateChannel,
		PermissionReadChannel,
		PermissionReadChannelContent,
		PermissionReadPublicChannelGroups,
		PermissionReadPrivateChannelGroups,
		PermissionAddReaction,
		PermissionRemoveReaction,
		PermissionRemoveOthersReactions,
		PermissionUploadFile,
		PermissionCreatePost,
		PermissionCreatePostPublic,
		PermissionCreatePostEphemeral,
		PermissionReadDeletedPosts,
		PermissionEditPost,
		PermissionEditOthersPosts,
		PermissionDeletePost,
		PermissionDeleteOthersPosts,
		PermissionUseChannelMentions,
		PermissionUseGroupMentions,
		PermissionAddBookmarkPublicChannel,
		PermissionEditBookmarkPublicChannel,
		PermissionDeleteBookmarkPublicChannel,
		PermissionOrderBookmarkPublicChannel,
		PermissionAddBookmarkPrivateChannel,
		PermissionEditBookmarkPrivateChannel,
		PermissionDeleteBookmarkPrivateChannel,
		PermissionOrderBookmarkPrivateChannel,
	}

	GroupScopedPermissions := []*Permission{
		PermissionManageCustomGroupMembers,
		PermissionEditCustomGroup,
		PermissionDeleteCustomGroup,
		PermissionRestoreCustomGroup,
	}

	DeprecatedPermissions = []*Permission{
		PermissionPermanentDeleteUser,
		PermissionManageWebhooks,
		PermissionManageOthersWebhooks,
		PermissionManageEmojis,
		PermissionManageOthersEmojis,
		PermissionSysconsoleReadAuthentication,
		PermissionSysconsoleWriteAuthentication,
		PermissionSysconsoleReadSite,
		PermissionSysconsoleWriteSite,
		PermissionSysconsoleReadEnvironment,
		PermissionSysconsoleWriteEnvironment,
		PermissionSysconsoleReadReporting,
		PermissionSysconsoleWriteReporting,
		PermissionSysconsoleReadAbout,
		PermissionSysconsoleWriteAbout,
		PermissionSysconsoleReadExperimental,
		PermissionSysconsoleWriteExperimental,
		PermissionSysconsoleReadIntegrations,
		PermissionSysconsoleWriteIntegrations,
		PermissionSysconsoleReadCompliance,
		PermissionSysconsoleWriteCompliance,
	}

	PlaybookScopedPermissions := []*Permission{
		PermissionPublicPlaybookManageProperties,
		PermissionPublicPlaybookManageMembers,
		PermissionPublicPlaybookManageRoles,
		PermissionPublicPlaybookView,
		PermissionPublicPlaybookMakePrivate,
		PermissionPrivatePlaybookManageProperties,
		PermissionPrivatePlaybookManageMembers,
		PermissionPrivatePlaybookManageRoles,
		PermissionPrivatePlaybookView,
		PermissionPrivatePlaybookMakePublic,
		PermissionRunCreate,
	}

	RunScopedPermissions := []*Permission{
		PermissionRunManageProperties,
		PermissionRunManageMembers,
		PermissionRunView,
	}

	AllPermissions = []*Permission{}
	AllPermissions = append(AllPermissions, SystemScopedPermissionsMinusSysconsole...)
	AllPermissions = append(AllPermissions, TeamScopedPermissions...)
	AllPermissions = append(AllPermissions, ChannelScopedPermissions...)
	AllPermissions = append(AllPermissions, SysconsoleReadPermissions...)
	AllPermissions = append(AllPermissions, SysconsoleWritePermissions...)
	AllPermissions = append(AllPermissions, GroupScopedPermissions...)
	AllPermissions = append(AllPermissions, PlaybookScopedPermissions...)
	AllPermissions = append(AllPermissions, RunScopedPermissions...)

	ChannelModeratedPermissions = []string{
		PermissionCreatePost.Id,
		"create_reactions",
		"manage_members",
		PermissionUseChannelMentions.Id,
		"manage_bookmarks",
	}

	ChannelModeratedPermissionsMap = map[string]string{
		PermissionCreatePost.Id:                  ChannelModeratedPermissions[0],
		PermissionAddReaction.Id:                 ChannelModeratedPermissions[1],
		PermissionRemoveReaction.Id:              ChannelModeratedPermissions[1],
		PermissionManagePublicChannelMembers.Id:  ChannelModeratedPermissions[2],
		PermissionManagePrivateChannelMembers.Id: ChannelModeratedPermissions[2],
		PermissionUseChannelMentions.Id:          ChannelModeratedPermissions[3],
	}

	ModeratedBookmarkPermissions = []*Permission{
		PermissionAddBookmarkPublicChannel,
		PermissionEditBookmarkPublicChannel,
		PermissionDeleteBookmarkPublicChannel,
		PermissionOrderBookmarkPublicChannel,
		PermissionAddBookmarkPrivateChannel,
		PermissionEditBookmarkPrivateChannel,
		PermissionDeleteBookmarkPrivateChannel,
		PermissionOrderBookmarkPrivateChannel,
	}

	for _, mbp := range ModeratedBookmarkPermissions {
		ChannelModeratedPermissionsMap[mbp.Id] = ChannelModeratedPermissions[4]
	}
}

func init() {
	initializePermissions()
}

func MakePermissionError(s *Session, permissions []*Permission) *AppError {
	return MakePermissionErrorForUser(s.UserId, permissions)
}

func MakePermissionErrorForUser(userId string, permissions []*Permission) *AppError {
	permissionsStr := "permission="
	for i, permission := range permissions {
		permissionsStr += permission.Id
		if i != len(permissions)-1 {
			permissionsStr += ","
		}
	}
	return NewAppError("Permissions", "api.context.permissions.app_error", nil, "userId="+userId+", "+permissionsStr, http.StatusForbidden)
}
