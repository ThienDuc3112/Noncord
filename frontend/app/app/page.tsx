// app/app/page.tsx
"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import Cookies from "universal-cookie";
import MainView from "./mainview";
import { theme, backgroundPattern } from "@/lib/theme";

export default function AppPage() {
  const router = useRouter();
  const [checkingAuth, setCheckingAuth] = useState(true);
  const [hasToken, setHasToken] = useState(false);

  useEffect(() => {
    const cookies = new Cookies();
    const accessToken = cookies.get("accessToken");

    if (!accessToken) {
      // No token -> bounce to login
      router.replace("/login");
      setHasToken(false);
      setCheckingAuth(false);
      return;
    }

    setHasToken(true);
    setCheckingAuth(false);
  }, [router]);

  if (checkingAuth) {
    return (
      <div
        className={`flex min-h-screen items-center justify-center ${theme.classes.background} ${theme.colors.text.secondary}`}
        style={{ backgroundImage: backgroundPattern }}
      >
        <span className="text-sm">Loading your workspace...</span>
      </div>
    );
  }

  if (!hasToken) {
    // We already called router.replace("/login"), this is just a visual fallback
    return (
      <div
        className={`flex min-h-screen items-center justify-center ${theme.classes.background} ${theme.colors.text.secondary}`}
        style={{ backgroundImage: backgroundPattern }}
      >
        <span className="text-sm">Redirecting to login...</span>
      </div>
    );
  }

  // Auth ok -> show Discord-style main app
  return <MainView />;
}
