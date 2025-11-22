"use client";

import { Home, Plus } from "lucide-react";
import type { ServerPreview } from "@/lib/types";
import { theme } from "@/lib/theme";
import CreateServerDialog from "./createServerDialog";

interface SidebarProps {
  servers: ServerPreview[];
  selectedServerId: string | null;
  onSelectServer: (id: string | null) => void;
  onServerCreated: (server: ServerPreview) => void;
}

function getInitials(name: string) {
  return name
    .split(" ")
    .map((part) => part[0])
    .join("")
    .slice(0, 2)
    .toUpperCase();
}

export default function Sidebar({
  servers,
  selectedServerId,
  onSelectServer,
  onServerCreated,
}: SidebarProps) {
  return (
    <aside
      className={`flex h-screen w-16 flex-col items-center gap-3 border-r border-[#363a4f] ${theme.colors.background.card} py-3`}
    >
      {/* Home / default */}
      <button
        onClick={() => onSelectServer(null)}
        className={`flex h-11 w-11 items-center justify-center rounded-3xl ${theme.colors.interactive.primary} transition-all hover:rounded-2xl`}
      >
        <Home className="h-5 w-5" />
      </button>

      <div className="mt-2 h-px w-8 bg-[#363a4f]" />

      {/* Server icons */}
      <div className="flex flex-1 flex-col gap-2 overflow-y-auto">
        {servers.map((server) => {
          const isActive = server.id === selectedServerId;
          return (
            <button
              key={server.id}
              onClick={() => onSelectServer(server.id)}
              className={`relative flex h-11 w-11 items-center justify-center rounded-3xl text-sm font-semibold transition-all ${
                isActive
                  ? "bg-[#c6a0f6] text-[#181926] hover:bg-[#b7bdf8]"
                  : "bg-[#363a4f] text-[#cad3f5] hover:rounded-2xl hover:bg-[#494d64]"
              }`}
            >
              {server.iconUrl ? (
                // eslint-disable-next-line @next/next/no-img-element
                <img
                  src={server.iconUrl}
                  alt={server.name}
                  className="h-full w-full rounded-3xl object-cover"
                />
              ) : (
                <span>{getInitials(server.name)}</span>
              )}
            </button>
          );
        })}
      </div>

      {/* Create server (+) */}
      <CreateServerDialog onServerCreated={onServerCreated}>
        <button
          className="flex h-11 w-11 items-center justify-center rounded-full bg-[#363a4f] text-[#a5adcb] transition-all hover:bg-[#494d64] hover:text-[#cad3f5]"
          aria-label="Create server"
        >
          <Plus className="h-5 w-5" />
        </button>
      </CreateServerDialog>
    </aside>
  );
}
