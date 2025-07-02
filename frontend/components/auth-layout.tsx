import type React from "react";
import { theme, backgroundPattern } from "@/lib/theme";
import { Card, CardContent } from "@/components/ui/card";

interface AuthLayoutProps {
  children: React.ReactNode;
}

export function AuthLayout({ children }: AuthLayoutProps) {
  return (
    <div
      className={
        theme.classes.background + " flex items-center justify-center p-4"
      }
    >
      {/* Background Pattern */}
      <div className="absolute inset-0 opacity-10">
        <div
          className="absolute inset-0"
          style={{
            backgroundImage: backgroundPattern,
          }}
        />
      </div>

      <Card className={theme.classes.card + " w-full max-w-md relative z-10"}>
        <CardContent className="p-8">{children}</CardContent>
      </Card>
    </div>
  );
}
