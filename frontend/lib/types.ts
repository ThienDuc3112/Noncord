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
      // assuming you extended zod with .jwt(); if not, you can remove this line
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

export const GetServersSchema = z.object({
  result: z.array(ServerPreviewSchema),
});

export const GetServerSchema = z.object({
  id: z.string(),
  name: z.string(),
  bannerUrl: z.string().nullish(),
  iconUrl: z.string().nullish(),
  description: z.string().nullish(),
  announcementChannel: z.string().nullish(),
  defaultPermission: z.number().int().optional(),
  createdAt: z.string(),
  updatedAt: z.string(),
  channels: z.array(ChannelSchema),
});

export const MessageSchema = z.object({
  id: z.string(),
  author: z.string().nullish(),
  authorType: z.string(),
  channelId: z.string().nullish(),
  groupId: z.string().nullish(),
  message: z.string(),
  createdAt: z.string(),
  updatedAt: z.string(),
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

// Types
export type LoginData = z.infer<typeof LoginSchema>;
export type TokenData = z.infer<typeof TokenSchema>;
export type ServerPreview = z.infer<typeof ServerPreviewSchema>;
export type Channel = z.infer<typeof ChannelSchema>;
export type GetServersResponse = z.infer<typeof GetServersSchema>;
export type GetServerResponse = z.infer<typeof GetServerSchema>;
export type Message = z.infer<typeof MessageSchema>;
export type GetMessagesResponse = z.infer<typeof GetMessagesSchema>;
export type NewServerResponse = z.infer<typeof NewServerResponseSchema>;

export interface Member {
  id: string;
  name: string;
  status?: "online" | "offline";
  avatarUrl?: string | null;
}
