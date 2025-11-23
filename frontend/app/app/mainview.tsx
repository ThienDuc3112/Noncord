"use client";

import { useCallback, useEffect, useMemo } from "react";
import Sidebar from "./sidebar";
import DefaultView from "./defaultView";
import ChatList from "./chatList";
import MemberList from "./member";

import {
  fetchServers,
  fetchServerById,
  fetchChannelMessages,
  sendChannelMessage,
} from "@/lib/request";
import { theme, backgroundPattern } from "@/lib/theme";

import type {
  Channel,
  GetMessagesResponse,
  GetServerResponse,
  Member,
  Message,
  ServerPreview,
} from "@/lib/types";
import { useQuery, useQueryClient } from "@tanstack/react-query";
import { useAtom } from "jotai";
import { selectedChannelIdAtom, selectedServerIdAtom } from "./state";

export default function MainView() {
  const queryClient = useQueryClient();

  const [selectedServerId, setSelectedServerId] = useAtom(selectedServerIdAtom);
  const [selectedChannelId, setSelectedChannelId] = useAtom(
    selectedChannelIdAtom,
  );

  // 1) Fetch list of servers
  const {
    data: serversData,
    isLoading: isLoadingServers,
    error: serversError,
  } = useQuery<ServerPreview[]>({
    queryKey: ["servers"],
    queryFn: fetchServers,
  });

  const servers = serversData ?? [];

  // 2) When server changes, fetch server details & channels
  const {
    data: serverData,
    isLoading: isLoadingServer,
    error: serverError,
  } = useQuery<GetServerResponse>({
    queryKey: ["server", selectedServerId],
    queryFn: () => fetchServerById(selectedServerId as string),
    enabled: !!selectedServerId,
  });

  const currentServer = serverData ?? null;
  const channels: Channel[] = currentServer?.channels ?? [];

  // Ensure a channel is selected when channels change
  useEffect(() => {
    if (!channels.length) {
      setSelectedChannelId(null);
      return;
    }

    if (
      !selectedChannelId ||
      !channels.some((c) => c.id === selectedChannelId)
    ) {
      setSelectedChannelId(channels[0].id);
    }
  }, [channels, selectedChannelId]);

  // 3) When channel changes, fetch messages
  const {
    data: messagesData,
    isLoading: isLoadingMessages,
    error: messagesError,
  } = useQuery<GetMessagesResponse>({
    queryKey: ["channelMessages", selectedChannelId],
    queryFn: () => fetchChannelMessages(selectedChannelId as string, 100),
    enabled: !!selectedChannelId,
  });

  const messages: Message[] = messagesData?.result ?? [];

  // Unified load error message similar to your previous setLoadError usage
  const loadError = useMemo(() => {
    if (serversError) {
      return "Failed to load servers. Your session may have expired, please try logging in again.";
    }
    if (serverError) {
      return "Failed to load server details.";
    }
    if (messagesError) {
      return "Failed to load messages.";
    }
    return null;
  }, [serversError, serverError, messagesError]);

  // Mock members until membership endpoint is wired
  const members: Member[] = useMemo(
    () => [
      { id: "you", name: "You", status: "online" },
      { id: "bot", name: "Noncord Bot", status: "online" },
    ],
    [],
  );

  // 4) Send a message (update React Query cache instead of local state)
  const handleSendMessage = useCallback(
    async (content: string) => {
      if (!selectedChannelId) return;

      const created = await sendChannelMessage(selectedChannelId, content);

      queryClient.setQueryData<GetMessagesResponse | undefined>(
        ["channelMessages", selectedChannelId],
        (old) => {
          const prev = old?.result ?? [];
          return {
            ...(old ?? { result: [] as Message[] }),
            result: [
              {
                id: created.id,
                authorType: "user",
                author: "",
                displayName: "",
                avatarUrl: "",
                createdAt: created.createdAt,
                updatedAt: created.createdAt,
                message: content,
                channelId: selectedChannelId,
                groupId: null,
              },
              ...prev,
            ],
          };
        },
      );
    },
    [selectedChannelId, queryClient],
  );

  // 5) When a server is created via the sidebar dialog
  const handleServerCreated = useCallback(
    (server: ServerPreview) => {
      queryClient.setQueryData<ServerPreview[] | undefined>(
        ["servers"],
        (old) => {
          const list = old ?? [];
          const exists = list.some((s) => s.id === server.id);
          if (exists) return list;
          return [...list, server];
        },
      );
      setSelectedServerId(server.id);
    },
    [queryClient],
  );

  const currentChannel = useMemo(
    () => channels.find((c) => c.id === selectedChannelId) ?? null,
    [channels, selectedChannelId],
  );

  // Top-level layout
  return (
    <div
      className={`flex min-h-screen ${theme.classes.background} ${theme.colors.text.primary}`}
      style={{ backgroundImage: backgroundPattern }}
    >
      <Sidebar servers={servers} onServerCreated={handleServerCreated} />

      <section className="flex flex-1 max-h-screen">
        {/* Left: channels */}
        <div className="flex w-64 flex-col border-r border-[#363a4f] bg-[#1e2030]">
          <div className="flex h-12 items-center border-b border-[#363a4f] px-3 text-sm font-semibold">
            {isLoadingServers || isLoadingServer
              ? "Loading server..."
              : (currentServer?.name ?? "Select a server")}
          </div>

          <div className="flex-1 overflow-y-auto py-2">
            {channels.length === 0 ? (
              <p className={`px-4 text-xs ${theme.colors.text.muted}`}>
                No channels. Create one in your backend / admin UI.
              </p>
            ) : (
              <ul className="space-y-0.5 px-2">
                {channels
                  .slice()
                  .sort(
                    (a, b) =>
                      (a.order ?? Number.MAX_SAFE_INTEGER) -
                      (b.order ?? Number.MAX_SAFE_INTEGER),
                  )
                  .map((channel) => {
                    const isActive = channel.id === selectedChannelId;
                    return (
                      <li key={channel.id}>
                        <button
                          onClick={() => setSelectedChannelId(channel.id)}
                          className={`flex w-full items-center gap-2 rounded-md px-2 py-1.5 text-left text-sm transition-colors ${
                            isActive
                              ? "bg-[#363a4f] text-[#cad3f5]"
                              : "text-[#a5adcb] hover:bg-[#24273a] hover:text-[#cad3f5]"
                          }`}
                        >
                          <span className="text-[#6e738d]">#</span>
                          <span className="truncate">{channel.name}</span>
                        </button>
                      </li>
                    );
                  })}
              </ul>
            )}
          </div>
        </div>

        {/* Middle: chat */}
        <div className="flex flex-1 flex-col">
          {loadError && (
            <div className="border-b border-[#363a4f] bg-[#3b1f2b]/80 px-4 py-2 text-xs text-[#ed8796]">
              {loadError}
            </div>
          )}

          {selectedChannelId && currentChannel ? (
            <ChatList
              channel={currentChannel}
              messages={messages}
              isLoading={isLoadingMessages}
              onSendMessage={handleSendMessage}
            />
          ) : (
            <DefaultView />
          )}
        </div>

        {/* Right: members */}
        <div className="hidden w-60 border-l border-[#363a4f] bg-[#1e2030]/90 md:flex">
          <MemberList members={members} />
        </div>
      </section>
    </div>
  );
}
