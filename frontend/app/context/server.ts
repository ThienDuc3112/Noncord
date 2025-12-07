import { ServerPreview } from "@/lib/types";
import { atom } from "jotai";

export const ServersAtom = atom<ServerPreview[]>([]);
