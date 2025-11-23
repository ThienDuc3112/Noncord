import type { Channel, Message } from "@/lib/types";
import MessageItem from "./message";
import MessageInput from "./messageInput";
import { theme } from "@/lib/theme";

interface ChatListProps {
  channel: Channel;
  messages: Message[];
  isLoading: boolean;
  onSendMessage: (content: string) => Promise<void> | void;
}

export default function ChatList({
  channel,
  messages,
  isLoading,
  onSendMessage,
}: ChatListProps) {
  return (
    <div className="flex h-full max-h-full min-h-full flex-1 flex-col bg-[#181926]">
      {/* Header */}
      <div className="flex h-12 items-center border-b border-[#363a4f] px-4 text-sm">
        <span className="mr-2 text-[#6e738d]">#</span>
        <span className="font-semibold text-[#cad3f5]">{channel.name}</span>
      </div>

      {/* Messages */}
      <div className="flex-1 overflow-y-scroll px-4 py-3">
        {isLoading && messages.length === 0 ? (
          <p className={`${theme.colors.text.muted} text-xs`}>
            Loading messages...
          </p>
        ) : messages.length === 0 ? (
          <p className="text-xs text-[#8087a2]">
            No messages yet. Be the first to say hi!
          </p>
        ) : (
          <div className="space-y-2 flex gap-2 flex-col-reverse">
            {messages.map((m) => (
              <MessageItem key={m.id} message={m} />
            ))}
          </div>
        )}
      </div>

      {/* Input */}
      <div className="border-t border-[#363a4f] px-4 py-3">
        <MessageInput
          placeholder={`Message #${channel.name}`}
          onSend={onSendMessage}
        />
      </div>
    </div>
  );
}
