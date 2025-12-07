import { atom } from "jotai";

type UserContext = {
  id: string;
};

export const UserAtom = atom<UserContext | null>();
