import { theme } from "@/lib/theme";

export default function DefaultView() {
  return (
    <div className="flex h-full flex-1 flex-col items-center justify-center bg-[#181926]">
      <div className="max-w-md text-center">
        <h1 className="mb-2 text-2xl font-semibold text-[#cad3f5]">
          Welcome to Noncord
        </h1>
        <p className={`${theme.colors.text.secondary} text-sm`}>
          Select a server and a channel on the left sidebar to start chatting.
        </p>
      </div>
    </div>
  );
}
