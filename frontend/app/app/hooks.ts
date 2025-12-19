import {
  fetchChannelMessages,
  fetchServerById,
  fetchServers,
} from "@/lib/request";
import { GetServerResponse, Message, ServerPreview } from "@/lib/types";
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
  return useQuery<Message[]>({
    queryKey: ["channelMessages", selectedChannelId],
    queryFn: async () =>
      (await fetchChannelMessages(selectedChannelId as string, 100)).result,
    enabled: !!selectedChannelId,
  });
};
