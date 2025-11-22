"use client";
import axios from "axios";
import Cookies from "universal-cookie";
import {
  GetServersSchema,
  GetServerSchema,
  GetMessagesSchema,
  MessageSchema,
  NewServerSchema,
  NewServerResponseSchema,
  type ServerPreview,
  type GetServerResponse,
  type GetMessagesResponse,
  type Message,
} from "./types";

export const apiClient = axios.create({
  baseURL: process.env.API_URL || process.env.NEXT_PUBLIC_BACKEND_URL,
  headers: {
    "Content-Type": "application/json",
  },
});

function getAuthHeaders() {
  if (typeof window === "undefined") return {};

  const cookies = new Cookies();
  const token = cookies.get("accessToken");

  if (!token) return {};
  return {
    Authorization: `Bearer ${token}`,
  };
}

export async function fetchServers(): Promise<ServerPreview[]> {
  const res = await apiClient.get("/server", {
    headers: getAuthHeaders(),
  });

  const parsed = GetServersSchema.safeParse(res.data);
  if (!parsed.success) {
    console.error("Invalid server list response", parsed.error);
    throw new Error("Server returned invalid server list shape");
  }

  return parsed.data.result;
}

export async function fetchServerById(
  serverId: string,
): Promise<GetServerResponse> {
  const res = await apiClient.get(`/server/${serverId}`, {
    headers: getAuthHeaders(),
  });

  const parsed = GetServerSchema.safeParse(res.data);
  if (!parsed.success) {
    console.error("Invalid server response", parsed.error);
    throw new Error("Server returned invalid server shape");
  }

  return parsed.data;
}

export async function fetchChannelMessages(
  channelId: string,
  limit = 100,
): Promise<GetMessagesResponse> {
  const res = await apiClient.get(`/message/channel/${channelId}`, {
    headers: getAuthHeaders(),
    params: { limit },
  });

  const parsed = GetMessagesSchema.safeParse(res.data);
  if (!parsed.success) {
    console.error("Invalid messages response", parsed.error);
    throw new Error("Server returned invalid messages shape");
  }

  return parsed.data;
}

export async function sendChannelMessage(
  channelId: string,
  content: string,
): Promise<Message> {
  const res = await apiClient.post(
    "/message",
    {
      content,
      isTargetChannel: true,
      targetId: channelId,
    },
    {
      headers: getAuthHeaders(),
    },
  );

  const parsed = MessageSchema.safeParse(res.data);
  if (!parsed.success) {
    console.error("Invalid message response", parsed.error);
    throw new Error("Server returned invalid message shape");
  }

  return parsed.data;
}

export async function createServer(name: string): Promise<ServerPreview> {
  // Validate payload on client
  const payload = NewServerSchema.parse({ name });

  const res = await apiClient.post("/server", payload, {
    headers: getAuthHeaders(),
  });

  const parsed = NewServerResponseSchema.safeParse(res.data);
  if (!parsed.success) {
    console.error("Invalid create-server response", parsed.error);
    throw new Error("Server returned invalid create-server shape");
  }

  // Use the returned id to fetch full server details and build a preview
  const server = await fetchServerById(parsed.data.id);

  return {
    id: server.id,
    name: server.name,
    iconUrl: server.iconUrl ?? undefined,
    bannerUrl: server.bannerUrl ?? undefined,
  };
}
