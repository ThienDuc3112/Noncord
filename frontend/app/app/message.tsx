import type { Message } from "@/lib/types";
import { Avatar, AvatarFallback } from "@/components/ui/avatar";

interface MessageProps {
  message: Message;
}

function formatTime(iso: string) {
  const d = new Date(iso);
  return d.toLocaleTimeString(undefined, {
    hour: "2-digit",
    minute: "2-digit",
  });
}

export default function MessageItem({ message }: MessageProps) {
  const initials = message.author
    ? message.author
        .split(" ")
        .map((p) => p[0])
        .join("")
        .slice(0, 2)
        .toUpperCase()
    : "";

  return (
    <div className="group flex gap-3 text-sm">
      {message.authorType != "system" && (
        <Avatar className="mt-0.5 h-8 w-8">
          <AvatarFallback className="bg-[#363a4f] text-[10px] text-[#cad3f5]">
            {initials || "?"}
          </AvatarFallback>
        </Avatar>
      )}

      <div>
        <div className="flex items-baseline gap-2">
          <span className="font-semibold text-[#cad3f5]">{message.author}</span>
          <span className="text-[11px] text-[#6e738d]">
            {formatTime(message.createdAt)}
          </span>
        </div>
        <div className="text-[#cad3f5]">{message.message}</div>
      </div>
    </div>
  );
}
