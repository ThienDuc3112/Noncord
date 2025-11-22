"use client";

import { useState } from "react";
import { Send } from "lucide-react";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { theme } from "@/lib/theme";

interface MessageInputProps {
  placeholder?: string;
  onSend: (content: string) => Promise<void> | void;
}

export default function MessageInput({
  placeholder,
  onSend,
}: MessageInputProps) {
  const [value, setValue] = useState("");
  const [sending, setSending] = useState(false);

  const handleSubmit = async () => {
    const trimmed = value.trim();
    if (!trimmed) return;
    setSending(true);
    try {
      await onSend(trimmed);
      setValue("");
    } finally {
      setSending(false);
    }
  };

  const handleKeyDown: React.KeyboardEventHandler<HTMLInputElement> = (e) => {
    if (e.key === "Enter" && !e.shiftKey) {
      e.preventDefault();
      void handleSubmit();
    }
  };

  return (
    <div className="flex items-center gap-2">
      <Input
        value={value}
        onChange={(e) => setValue(e.target.value)}
        onKeyDown={handleKeyDown}
        placeholder={placeholder ?? "Message"}
        className={`flex-1 text-sm ${theme.classes.input}`}
      />
      <Button
        size="icon"
        variant="ghost"
        disabled={sending || !value.trim()}
        onClick={handleSubmit}
        className="text-[#a5adcb] hover:text-[#cad3f5]"
      >
        <Send className="h-4 w-4" />
      </Button>
    </div>
  );
}
