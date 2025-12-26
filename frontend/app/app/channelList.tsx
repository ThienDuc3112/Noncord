import type { Channel } from "@/lib/types";
import { useFetchServerByIdQuery, useFetchServersQuery } from "./hooks";
import { useMemo } from "react";
import { theme } from "@/lib/theme";
import { useAtom, useAtomValue } from "jotai";
import { selectedChannelIdAtom, selectedServerIdAtom } from "./state";
import { Plus } from "lucide-react";
import CreateChannelDialog from "./createChannelDialog";

const ChannelList = () => {
  const { isLoading: isLoadingServers } = useFetchServersQuery();

  const selectedServerId = useAtomValue(selectedServerIdAtom);

  const { data: serverData, isLoading: isLoadingServer } =
    useFetchServerByIdQuery();

  const currentServer = serverData ?? null;
  const channels: Channel[] = useMemo(
    () => currentServer?.channels ?? [],
    [currentServer?.channels],
  );

  const [selectedChannelId, setSelectedChannelId] = useAtom(
    selectedChannelIdAtom,
  );

  return (
    <div className="flex w-64 flex-col border-r border-[#363a4f] bg-[#1e2030]">
      <div className="flex h-12 items-center border-b border-[#363a4f] px-3 text-sm font-semibold">
        {isLoadingServers || isLoadingServer
          ? "Loading server..."
          : (currentServer?.name ?? "Select a server")}
      </div>

      <div className="flex-1 overflow-y-auto py-2">
        {channels.length === 0 ? (
          <p className={`px-4 text-xs ${theme.colors.text.muted}`}>
            No channels. Create one to get started.
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

      {selectedServerId && (
        <div className="p-1">
          <CreateChannelDialog
            serverId={selectedServerId}
            onChannelCreated={(channel) => setSelectedChannelId(channel.id)}
          >
            <button className={theme.classes.button.ghostCta}>
              <Plus className="h-4 w-4 text-[#c6a0f6]" />
              <span className="truncate">Add channel</span>
            </button>
          </CreateChannelDialog>
        </div>
      )}
    </div>
  );
};

export default ChannelList;
