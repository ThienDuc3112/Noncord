import { Message } from "@/lib/types";
import { atom } from "jotai";

export const selectedServerIdAtom = atom<string | null>(null);
export const selectedChannelIdAtom = atom<string | null>(null);
export const tempWsMessages = atom<Record<string, Message[]>>({});

export const messagesAtom = atom<Record<string, Message[]>>({});

export const mergeMessages = (
  a: Message[] | null | undefined,
  b: Message[] | null | undefined,
): Message[] => {
  if (!a && !b) return [];
  if (!a) return b!;
  if (!b) return a;
  const map = new Map<string, Message>();
  a.forEach((m) => map.set(m.id, m));
  b.forEach((m) => map.set(m.id, m));

  return map
    .values()
    .toArray()
    .toSorted((a, b) => b.createdAt.getTime() - a.createdAt.getTime());
};
