"use client";

import { useState, type ReactNode } from "react";
import { useRouter } from "next/navigation";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";

import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogDescription,
  DialogFooter,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Label } from "@/components/ui/label";

import { NewServerSchema, type ServerPreview } from "@/lib/types";
import { createServer } from "@/lib/request";
import { theme } from "@/lib/theme";

// ---- Schemas & types ----

type NewServerData = z.infer<typeof NewServerSchema>;

const JoinServerSchema = z.object({
  invite: z
    .string()
    .min(1, "Invitation link or ID is required")
    .max(512, "Invitation is too long"),
});

type JoinServerData = z.infer<typeof JoinServerSchema>;

// ---- Utils ----

// Extract UUID-looking string from either raw UUID or full URL
const INVITE_UUID_REGEX =
  /[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}/;

function extractInviteId(input: string): string | null {
  const trimmed = input.trim();
  const match = trimmed.match(INVITE_UUID_REGEX);
  return match ? match[0] : null;
}

interface CreateServerDialogProps {
  onServerCreated: (server: ServerPreview) => void;
  children: ReactNode; // trigger element (the + button)
}

export default function CreateServerDialog({
  onServerCreated,
  children,
}: CreateServerDialogProps) {
  const router = useRouter();
  const [open, setOpen] = useState(false);
  const [mode, setMode] = useState<"create" | "join">("create");

  // Create-server form
  const createForm = useForm<NewServerData>({
    resolver: zodResolver(NewServerSchema),
    defaultValues: { name: "" },
  });

  const {
    register: registerCreate,
    handleSubmit: handleSubmitCreate,
    formState: { errors: createErrors, isSubmitting: isCreating },
    setError: setCreateError,
    reset: resetCreate,
  } = createForm;

  // Join-server form
  const joinForm = useForm<JoinServerData>({
    resolver: zodResolver(JoinServerSchema),
    defaultValues: { invite: "" },
  });

  const {
    register: registerJoin,
    handleSubmit: handleSubmitJoin,
    formState: { errors: joinErrors, isSubmitting: isJoining },
    setError: setJoinError,
    reset: resetJoin,
  } = joinForm;

  // ---- Handlers ----

  const onSubmitCreate = async (data: NewServerData) => {
    try {
      const server = await createServer(data.name);
      onServerCreated(server);
      resetCreate();
      resetJoin();
      setOpen(false);
    } catch (err) {
      console.error("Failed to create server", err);
      setCreateError("root", {
        message: "Failed to create server. Please try again.",
      });
    }
  };

  const onSubmitJoin = async (data: JoinServerData) => {
    const inviteId = extractInviteId(data.invite);

    if (!inviteId) {
      setJoinError("invite", {
        message: "Could not find a valid invitation ID in that input.",
      });
      return;
    }

    // For now, just navigate to the invite page;
    // that route will handle contacting the backend and joining the server.
    router.push(`/invite/${inviteId}`);

    resetJoin();
    resetCreate();
    setOpen(false);
  };

  const createRootError = createForm.formState.errors.root?.message;
  const joinRootError = joinForm.formState.errors.root?.message;

  const closeDialog = () => {
    setOpen(false);
    // Optional: reset on close
    resetCreate();
    resetJoin();
    setMode("create");
  };

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>{children}</DialogTrigger>

      <DialogContent className={theme.classes.card}>
        <DialogHeader>
          <DialogTitle className={theme.colors.text.primary}>
            Create or join a server
          </DialogTitle>
          <DialogDescription className={theme.colors.text.secondary}>
            Create a new Noncord server, or join an existing one with an invite.
          </DialogDescription>
        </DialogHeader>

        {/* Mode toggle */}
        <div className="mb-4 mt-2 flex rounded-md bg-[#24273a] p-1 text-xs font-medium">
          <button
            type="button"
            onClick={() => setMode("create")}
            className={`flex-1 rounded-sm px-2 py-1 transition-colors ${
              mode === "create"
                ? "bg-[#363a4f] text-[#cad3f5]"
                : "text-[#a5adcb] hover:text-[#cad3f5]"
            }`}
          >
            Create
          </button>
          <button
            type="button"
            onClick={() => setMode("join")}
            className={`flex-1 rounded-sm px-2 py-1 transition-colors ${
              mode === "join"
                ? "bg-[#363a4f] text-[#cad3f5]"
                : "text-[#a5adcb] hover:text-[#cad3f5]"
            }`}
          >
            Join
          </button>
        </div>

        {/* Forms */}
        {mode === "create" ? (
          <form
            className="space-y-4"
            onSubmit={handleSubmitCreate(onSubmitCreate)}
            noValidate
          >
            <div className="space-y-1">
              <Label htmlFor="server-name" className={theme.classes.label}>
                Server name
              </Label>
              <Input
                id="server-name"
                type="text"
                placeholder="My awesome server"
                className={
                  createErrors.name
                    ? theme.classes.inputError
                    : theme.classes.input
                }
                {...registerCreate("name")}
              />
              {createErrors.name && (
                <p className={theme.classes.formMessage}>
                  {createErrors.name.message}
                </p>
              )}
            </div>

            {createRootError && (
              <p className={theme.classes.formMessage}>{createRootError}</p>
            )}

            <DialogFooter className="mt-4 flex justify-end gap-2">
              <Button
                type="button"
                variant="ghost"
                onClick={closeDialog}
                className="text-[#a5adcb] hover:text-[#cad3f5]"
              >
                Cancel
              </Button>
              <Button
                type="submit"
                disabled={isCreating}
                className={theme.classes.button.primary}
              >
                {isCreating ? "Creating..." : "Create"}
              </Button>
            </DialogFooter>
          </form>
        ) : (
          <form
            className="space-y-4"
            onSubmit={handleSubmitJoin(onSubmitJoin)}
            noValidate
          >
            <div className="space-y-1">
              <Label htmlFor="invite" className={theme.classes.label}>
                Invitation link or ID
              </Label>
              <Input
                id="invite"
                type="text"
                placeholder="e.g. 123e4567-e89b-12d3-a456-426614174000 or https://noncord.app/invite/123e4567-e89b-12d3-a456-426614174000"
                className={
                  joinErrors.invite
                    ? theme.classes.inputError
                    : theme.classes.input
                }
                {...registerJoin("invite")}
              />
              {joinErrors.invite && (
                <p className={theme.classes.formMessage}>
                  {joinErrors.invite.message}
                </p>
              )}
            </div>

            {joinRootError && (
              <p className={theme.classes.formMessage}>{joinRootError}</p>
            )}

            <DialogFooter className="mt-4 flex justify-end gap-2">
              <Button
                type="button"
                variant="ghost"
                onClick={closeDialog}
                className="text-[#a5adcb] hover:text-[#cad3f5]"
              >
                Cancel
              </Button>
              <Button
                type="submit"
                disabled={isJoining}
                className={theme.classes.button.primary}
              >
                {isJoining ? "Joining..." : "Join"}
              </Button>
            </DialogFooter>
          </form>
        )}
      </DialogContent>
    </Dialog>
  );
}
