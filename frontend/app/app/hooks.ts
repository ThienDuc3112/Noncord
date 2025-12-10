import {
  fetchChannelMessages,
  fetchServerById,
  fetchServers,
} from "@/lib/request";
import {
  GetMessagesResponse,
  GetServerResponse,
  ServerPreview,
} from "@/lib/types";
import { useQuery } from "@tanstack/react-query";
import { useAtomValue } from "jotai";
import { selectedChannelIdAtom, selectedServerIdAtom } from "./state";

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

export const useFetchChannelMessagesQuery = () => {
  const selectedChannelId = useAtomValue(selectedChannelIdAtom);
  return useQuery<GetMessagesResponse>({
    queryKey: ["channelMessages", selectedChannelId],
    queryFn: () => fetchChannelMessages(selectedChannelId as string, 100),
    enabled: !!selectedChannelId,
  });
};
