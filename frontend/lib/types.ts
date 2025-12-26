import { z } from "zod";

export const LoginSchema = z.object({
  username: z
    .string({ message: "Username must be a string" })
    .nonempty("Username or email must exist"),
  password: z
    .string({ message: "Password must be a string" })
    .nonempty("Password must exist"),
});

export const TokenSchema = z.object({
  accessToken:
    z
      .string({ message: "server didn't return correct access token type" })
      .jwt?.({ message: "server didn't return jwt token" }) ??
    z.string().nonempty({ message: "server didn't return a access token" }),
  refreshToken: z
    .string({ message: "server didn't return correct refrsh token type" })
    .nonempty({ message: "server didn't return a refresh token" }),
});

export const ServerPreviewSchema = z.object({
  id: z.string(),
  name: z.string(),
  iconUrl: z.string().nullish(),
  bannerUrl: z.string().nullish(),
});

export const ChannelSchema = z.object({
  id: z.string(),
  name: z.string(),
  description: z.string().nullish(),
  serverId: z.string(),
  parentCategory: z.string().nullish(),
  order: z.number().int().optional(),
  createdAt: z.string().optional(),
  updatedAt: z.string().optional(),
});

export const RoleSchema = z.object({
  id: z.string(),
  name: z.string(),
  serverId: z.string(),
  color: z.number().int().optional(),
  allowMention: z.boolean().optional(),
  permissions: z.array(z.string()).optional(),
  priority: z.number().int().optional(),
});

export const GetServersSchema = z.object({
  result: z.array(ServerPreviewSchema),
});

export const MembershipSchema = z.object({
  createdAt: z.string(),
  nickname: z.string(),
  serverId: z.string(),
  userId: z.string(),
  assignedRoles: z.array(z.string().uuid()).nullable(),
});

export const GetServerSchema = z.object({
  id: z.string(),
  name: z.string(),
  bannerUrl: z.string().nullish(),
  iconUrl: z.string().nullish(),
  description: z.string().nullish(),
  announcementChannel: z.string().nullish(),
  defaultRole: z.string().nullish(),
  createdAt: z.string(),
  updatedAt: z.string(),
  channels: z.array(ChannelSchema),
  roles: z.array(RoleSchema),
  selfMembership: MembershipSchema.optional(),
});

export const MessageSchema = z.object({
  id: z.string(),
  author: z.string().nullish(),
  authorType: z.string(),
  channelId: z.string().nullish(),
  groupId: z.string().nullish(),
  message: z.string(),
  createdAt: z.coerce.date(),
  updatedAt: z.coerce.date(),
  avatarUrl: z.string(),
  displayName: z.string(),
});

export const CreateMessageSchema = z.object({
  id: z.string(),
  createdAt: z.string().datetime(),
});

export const GetMessagesSchema = z.object({
  result: z.array(MessageSchema),
  next: z.string().nullish().optional(),
});

export const NewServerSchema = z.object({
  name: z
    .string()
    .min(1, "Server name cannot be empty")
    .max(256, "Server name must be at most 256 characters"),
});

export const NewServerResponseSchema = z.object({
  id: z.string(),
});

export const NewInvitationSchema = z.object({
  bypassApproval: z.boolean().optional(),
  expiresAt: z.string().optional(),
  joinLimit: z.number().int().optional(),
});

export const InvitationSchema = z.object({
  id: z.string(),
  serverId: z.string(),
  bypassApproval: z.boolean(),
  createdAt: z.string(),
  expiresAt: z.string(),
  joinCount: z.number().int(),
  joinLimit: z.number().int(),
});

export const JoinServerResponseSchema = z.object({
  membership: MembershipSchema,
  server: ServerPreviewSchema,
});

export const CreateChannelBodySchema = z.object({
  name: z.string(),
  description: z.string(),
  parentCategory: z.string().nullable().optional(),
  serverId: z.string(),
});

export const CreateChannelResponseSchema = z.object({
  id: z.string(),
  createdAt: z.coerce.date(),
  updatedAt: z.coerce.date(),
  name: z.string(),
  description: z.string(),
  serverId: z.string(),
  order: z.number().gte(0),
  parentCategory: z.string().nullable().optional(),
});

// Types
export type LoginData = z.infer<typeof LoginSchema>;
export type TokenData = z.infer<typeof TokenSchema>;
export type ServerPreview = z.infer<typeof ServerPreviewSchema>;
export type Channel = z.infer<typeof ChannelSchema>;
export type GetServersResponse = z.infer<typeof GetServersSchema>;
export type GetServerResponse = z.infer<typeof GetServerSchema>;
export type Message = z.infer<typeof MessageSchema>;
export type CreateMessageResponse = z.infer<typeof CreateMessageSchema>;
export type GetMessagesResponse = z.infer<typeof GetMessagesSchema>;
export type NewServerResponse = z.infer<typeof NewServerResponseSchema>;
export type NewInvitationData = z.infer<typeof NewInvitationSchema>;
export type Invitation = z.infer<typeof InvitationSchema>;
export type Membership = z.infer<typeof MembershipSchema>;
export type JoinServerResponse = z.infer<typeof JoinServerResponseSchema>;
export type CreateChannelBody = z.infer<typeof CreateChannelBodySchema>;
export type CreateChannelResponse = z.infer<typeof CreateChannelResponseSchema>;

export interface Member {
  id: string;
  name: string;
  status?: "online" | "offline";
  avatarUrl?: string | null;
}

// WS Types
const WS_VERSION = 1;

export const WSEventTypeSchema = z.enum([
  "initialized",
  "auth_failed",
  "incoming_message",
  "message_updated",
  "message_deleted",
]);

export const WSResponseSchema = z.object({
  eventType: WSEventTypeSchema,
  payload: z.any(),
  version: z.number().lte(WS_VERSION, "unsopported version"),
});

export type WSResponse = z.infer<typeof WSResponseSchema>;
