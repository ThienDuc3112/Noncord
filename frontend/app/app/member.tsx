import type { Member } from "@/lib/types";
import { Avatar, AvatarFallback } from "@/components/ui/avatar";

interface MemberListProps {
  members: Member[];
}

export default function MemberList({ members }: MemberListProps) {
  const online = members.filter((m) => m.status !== "offline");
  const offline = members.filter((m) => m.status === "offline");

  const renderMember = (m: Member) => {
    const initials = m.name
      .split(" ")
      .map((p) => p[0])
      .join("")
      .slice(0, 2)
      .toUpperCase();

    return (
      <div
        key={m.id}
        className="flex items-center gap-2 rounded-md px-2 py-1.5 text-sm hover:bg-[#24273a]"
      >
        <Avatar className="h-7 w-7">
          <AvatarFallback className="bg-[#363a4f] text-[10px] text-[#cad3f5]">
            {initials || "?"}
          </AvatarFallback>
        </Avatar>
        <span className="truncate text-[#cad3f5]">{m.name}</span>
      </div>
    );
  };

  return (
    <div className="flex h-full flex-1 flex-col bg-[#1e2030]/90">
      <div className="flex h-12 items-center border-b border-[#363a4f] px-3 text-xs font-semibold uppercase tracking-wide text-[#8087a2]">
        Members
      </div>

      <div className="flex-1 overflow-y-auto px-2 py-2 text-xs">
        {online.length > 0 && (
          <div className="mb-3">
            <div className="mb-1 px-1 text-[10px] font-semibold uppercase tracking-wide text-[#6e738d]">
              Online — {online.length}
            </div>
            <div className="space-y-0.5">{online.map(renderMember)}</div>
          </div>
        )}

        {offline.length > 0 && (
          <div>
            <div className="mb-1 px-1 text-[10px] font-semibold uppercase tracking-wide text-[#6e738d]">
              Offline — {offline.length}
            </div>
            <div className="space-y-0.5">{offline.map(renderMember)}</div>
          </div>
        )}

        {members.length === 0 && (
          <p className="px-1 text-[11px] text-[#8087a2]">
            No members loaded yet. Wire up your membership API here.
          </p>
        )}
      </div>
    </div>
  );
}
