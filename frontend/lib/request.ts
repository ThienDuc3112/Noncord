"use client";
import axios, { AxiosError, AxiosRequestConfig } from "axios";
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
  CreateMessageSchema,
  CreateMessageResponse,
  TokenData,
  NewInvitationSchema,
  InvitationSchema,
  JoinServerResponseSchema,
  type NewInvitationData,
  type Invitation,
  type JoinServerResponse,
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

apiClient.interceptors.request.use(
  (config) => {
    config.headers.Authorization = getAuthHeaders().Authorization;
    return config;
  },
  (err) => Promise.reject(err),
);

let refreshPromise: Promise<TokenData> | null = null;

const refreshToken = async (originalRequestConfig: AxiosRequestConfig) => {
  if (typeof window === "undefined")
    return Promise.reject(new Error("Cannot refresh token on server"));
  const cookies = new Cookies();
  const refreshToken = cookies.get("refreshToken");

  if (!refreshToken) {
    (window as any).location = "/login";
    return Promise.reject(new Error("No refresh token"));
  }

  if (!refreshPromise) {
    refreshPromise = axios
      .post<TokenData>(process.env.NEXT_PUBLIC_BACKEND_URL + "/auth/refresh", {
        refreshToken: refreshToken,
      })
      .then((refreshRes) => {
        cookies.set("accessToken", refreshRes.data.accessToken, {
          path: "/",
          secure: true,
          sameSite: true,
        });
        cookies.set("refreshToken", refreshRes.data.refreshToken, {
          path: "/",
          secure: true,
          sameSite: true,
        });
        return refreshRes.data;
      })
      .catch((refreshError) => {
        console.error(refreshError);
        cookies.remove("refreshToken");
        cookies.remove("accessToken");
        window.alert("Token Expired. Please login");
        (window as any).location = "/login";
        throw refreshError;
      })
      .finally(() => {
        refreshPromise = null;
      });
  }

  const newData = await refreshPromise;

  if (!originalRequestConfig.headers) {
    originalRequestConfig.headers = {};
  }

  originalRequestConfig.headers.Authorization = `Bearer ${newData.accessToken}`;

  return axios(originalRequestConfig);
};

// Add a response interceptor
apiClient.interceptors.response.use(
  (response) => response,
  async (error: AxiosError) => {
    const status = error.response?.status;
    const data = error.response?.data as { error?: string } | undefined;

    if (status === 401 && data?.error === "token expired" && error.config) {
      return refreshToken(error.config!);
    }

    return Promise.reject(error);
  },
);

export async function fetchServers(): Promise<ServerPreview[]> {
  const res = await apiClient.get("/server");

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
  const res = await apiClient.get(`/server/${serverId}`);

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
): Promise<CreateMessageResponse> {
  const res = await apiClient.post("/message", {
    content,
    isTargetChannel: true,
    targetId: channelId,
  });

  const parsed = CreateMessageSchema.safeParse(res.data);
  if (!parsed.success) {
    console.error("Invalid message response", parsed.error);
    throw new Error("Server returned invalid message shape");
  }

  return parsed.data;
}

export async function createServer(name: string): Promise<ServerPreview> {
  // Validate payload on client
  const payload = NewServerSchema.parse({ name });

  const res = await apiClient.post("/server", payload);

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

export async function createInvitation(
  serverId: string,
  data: NewInvitationData,
): Promise<Invitation> {
  // Client-side validation of payload
  const payload = NewInvitationSchema.parse(data);

  const res = await apiClient.post(`/server/${serverId}/invitations`, payload);

  const parsed = InvitationSchema.safeParse(res.data);
  if (!parsed.success) {
    console.error("Invalid create-invitation response", parsed.error);
    throw new Error("Server returned invalid invitation shape");
  }

  return parsed.data;
}

export async function joinServer(
  invitationId: string,
): Promise<JoinServerResponse> {
  const res = await apiClient.post(`/invitations/${invitationId}/join`);

  const parsed = JoinServerResponseSchema.safeParse(res.data);
  if (!parsed.success) {
    console.error("Invalid join-server response", parsed.error);
    throw new Error("Server returned invalid join-server shape");
  }

  return parsed.data;
}
