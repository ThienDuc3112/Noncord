import {
  fetchChannelMessages,
  fetchServerById,
  fetchServers,
} from "@/lib/request";
import { GetServerResponse, Message, ServerPreview } from "@/lib/types";
import { useQuery } from "@tanstack/react-query";
import { useAtom, useAtomValue } from "jotai";
import {
  mergeMessages,
  messagesAtom,
  selectedChannelIdAtom,
  selectedServerIdAtom,
} from "./state";
import { useEffect } from "react";

export const useFetchServersQuery = () => {
  return useQuery<ServerPreview[]>({
    queryKey: ["servers"],
    queryFn: fetchServers,
  });
};

export const useFetchServerByIdQuery = () => {
  const selectedServerId = useAtomValue(selectedServerIdAtom);
  return useQuery<GetServerResponse>({
    queryKey: ["server", selectedServerId],
    queryFn: () => fetchServerById(selectedServerId as string),
    staleTime: 15_000,
    enabled: !!selectedServerId,
  });
};

export const useMessages = () => {
  const selectedChannelId = useAtomValue(selectedChannelIdAtom);
  const [messages, setMessages] = useAtom(messagesAtom);

  const data = useQuery<Message[]>({
    queryKey: ["channelMessages", selectedChannelId],
    queryFn: async () =>
      (await fetchChannelMessages(selectedChannelId as string, 100)).result,
    enabled: !!selectedChannelId,
    staleTime: 600_000,
  });

  useEffect(() => {
    if (data.data && selectedChannelId)
      setMessages((prev) => ({
        ...prev,
        [selectedChannelId]: mergeMessages(data.data, prev[selectedChannelId]),
      }));
  }, [data.data, selectedChannelId]);

  return {
    ...data,
    data: selectedChannelId ? messages[selectedChannelId] : undefined,
  };
};
