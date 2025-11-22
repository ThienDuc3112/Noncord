"use client";

import { useCallback, useEffect, useMemo, useState } from "react";
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
  GetServerResponse,
  Member,
  Message,
  ServerPreview,
} from "@/lib/types";

export default function MainView() {
  const [servers, setServers] = useState<ServerPreview[]>([]);
  const [selectedServerId, setSelectedServerId] = useState<string | null>(null);
  const [currentServer, setCurrentServer] = useState<GetServerResponse | null>(
    null,
  );
  const [channels, setChannels] = useState<Channel[]>([]);
  const [selectedChannelId, setSelectedChannelId] = useState<string | null>(
    null,
  );

  const [messages, setMessages] = useState<Message[]>([]);
  const [isLoadingServers, setIsLoadingServers] = useState(true);
  const [isLoadingServer, setIsLoadingServer] = useState(false);
  const [isLoadingMessages, setIsLoadingMessages] = useState(false);

  const [loadError, setLoadError] = useState<string | null>(null);

  // Mock members until you wire a membership endpoint
  const members: Member[] = useMemo(
    () => [
      { id: "you", name: "You", status: "online" },
      { id: "bot", name: "Noncord Bot", status: "online" },
    ],
    [],
  );

  // 1) Fetch list of servers
  useEffect(() => {
    let cancelled = false;

    (async () => {
      try {
        setIsLoadingServers(true);
        const svrs = await fetchServers();
        if (cancelled) return;

        setServers(svrs);
        if (svrs.length > 0) {
          setSelectedServerId(svrs[0].id);
        }
      } catch (err) {
        if (!cancelled) {
          console.error("Failed to fetch servers", err);
          setLoadError(
            "Failed to load servers. Your session may have expired, please try logging in again.",
          );
        }
      } finally {
        if (!cancelled) {
          setIsLoadingServers(false);
        }
      }
    })();

    return () => {
      cancelled = true;
    };
  }, []);

  // 2) When server changes, fetch server details & channels
  useEffect(() => {
    if (!selectedServerId) {
      setCurrentServer(null);
      setChannels([]);
      setSelectedChannelId(null);
      return;
    }

    let cancelled = false;

    (async () => {
      try {
        setIsLoadingServer(true);
        const server = await fetchServerById(selectedServerId);
        if (cancelled) return;

        setCurrentServer(server);
        setChannels(server.channels ?? []);
        setSelectedChannelId(server.channels?.[0]?.id ?? null);
      } catch (err) {
        if (!cancelled) {
          console.error("Failed to fetch server", err);
          setLoadError("Failed to load server details.");
        }
      } finally {
        if (!cancelled) {
          setIsLoadingServer(false);
        }
      }
    })();

    return () => {
      cancelled = true;
    };
  }, [selectedServerId]);

  // 3) When channel changes, fetch messages
  useEffect(() => {
    if (!selectedChannelId) {
      setMessages([]);
      return;
    }

    let cancelled = false;

    (async () => {
      try {
        setIsLoadingMessages(true);
        const data = await fetchChannelMessages(selectedChannelId, 100);
        if (cancelled) return;

        setMessages(data.result ?? []);
      } catch (err) {
        if (!cancelled) {
          console.error("Failed to fetch messages", err);
          setLoadError("Failed to load messages.");
        }
      } finally {
        if (!cancelled) {
          setIsLoadingMessages(false);
        }
      }
    })();

    return () => {
      cancelled = true;
    };
  }, [selectedChannelId]);

  // 4) Send a message
  const handleSendMessage = useCallback(
    async (content: string) => {
      if (!selectedChannelId) return;
      const created = await sendChannelMessage(selectedChannelId, content);
      setMessages((prev) => [...prev, created]);
    },
    [selectedChannelId],
  );

  // 5) When a server is created via the sidebar dialog
  const handleServerCreated = useCallback((server: ServerPreview) => {
    setServers((prev) => {
      const exists = prev.some((s) => s.id === server.id);
      if (exists) return prev;
      return [...prev, server];
    });
    setSelectedServerId(server.id);
  }, []);

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
      <Sidebar
        servers={servers}
        selectedServerId={selectedServerId}
        onSelectServer={setSelectedServerId}
        onServerCreated={handleServerCreated}
      />

      <section className="flex flex-1">
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
