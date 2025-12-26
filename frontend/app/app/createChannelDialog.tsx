"use client";

import { useState, type ReactNode } from "react";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { useQueryClient } from "@tanstack/react-query";

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

import { createChannel } from "@/lib/request";
import type { Channel } from "@/lib/types";
import { theme } from "@/lib/theme";

// ---- Form schema ----
const CreateChannelFormSchema = z.object({
  name: z.string().min(1, "Channel name is required").max(64, "Too long"),
  description: z.string().max(256, "Too long").optional(),
  parentCategory: z.string().max(128, "Too long").optional(),
});

type CreateChannelFormData = z.infer<typeof CreateChannelFormSchema>;

interface CreateChannelDialogProps {
  serverId: string;
  onChannelCreated?: (channel: Channel) => void;
  children: ReactNode; // trigger element
}

export default function CreateChannelDialog({
  serverId,
  onChannelCreated,
  children,
}: CreateChannelDialogProps) {
  const [open, setOpen] = useState(false);
  const queryClient = useQueryClient();

  const form = useForm<CreateChannelFormData>({
    resolver: zodResolver(CreateChannelFormSchema),
    defaultValues: {
      name: "",
      description: "",
      parentCategory: "",
    },
  });

  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
    setError,
    reset,
  } = form;

  const closeDialog = () => {
    setOpen(false);
    reset();
  };

  const onSubmit = async (data: CreateChannelFormData) => {
    try {
      const name = data.name.trim();
      const description = (data.description ?? "").trim() || undefined;
      const parentCategory = (data.parentCategory ?? "").trim() || undefined;

      const channel = await createChannel(serverId, name, {
        description,
        parentCategory,
      });

      // Refresh the server details so ChannelList gets the new channel
      await queryClient.invalidateQueries({
        queryKey: ["server", serverId],
      });

      onChannelCreated?.(channel);

      reset();
      setOpen(false);
    } catch (err) {
      console.error("Failed to create channel", err);
      setError("root", {
        message: "Failed to create channel. Please try again.",
      });
    }
  };

  const rootError = form.formState.errors.root?.message;

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>{children}</DialogTrigger>

      <DialogContent className={theme.classes.card}>
        <DialogHeader>
          <DialogTitle className={theme.colors.text.primary}>
            Create channel
          </DialogTitle>
          <DialogDescription className={theme.colors.text.secondary}>
            Add a new channel to this server.
          </DialogDescription>
        </DialogHeader>

        <form
          className="space-y-4"
          onSubmit={handleSubmit(onSubmit)}
          noValidate
        >
          <div className="space-y-1">
            <Label htmlFor="channel-name" className={theme.classes.label}>
              Channel name
            </Label>
            <Input
              id="channel-name"
              type="text"
              placeholder="general"
              className={
                errors.name ? theme.classes.inputError : theme.classes.input
              }
              {...register("name")}
            />
            {errors.name && (
              <p className={theme.classes.formMessage}>{errors.name.message}</p>
            )}
          </div>

          <div className="space-y-1">
            <Label htmlFor="channel-desc" className={theme.classes.label}>
              Description (optional)
            </Label>
            <Input
              id="channel-desc"
              type="text"
              placeholder="What is this channel for?"
              className={
                errors.description
                  ? theme.classes.inputError
                  : theme.classes.input
              }
              {...register("description")}
            />
            {errors.description && (
              <p className={theme.classes.formMessage}>
                {errors.description.message}
              </p>
            )}
          </div>

          {
            // <div className="space-y-1">
            //   <Label htmlFor="channel-parent" className={theme.classes.label}>
            //     Parent category (optional)
            //   </Label>
            //   <Input
            //     id="channel-parent"
            //     type="text"
            //     placeholder="category-id (if you use categories)"
            //     className={
            //       errors.parentCategory
            //         ? theme.classes.inputError
            //         : theme.classes.input
            //     }
            //     {...register("parentCategory")}
            //   />
            //   {errors.parentCategory && (
            //     <p className={theme.classes.formMessage}>
            //       {errors.parentCategory.message}
            //     </p>
            //   )}
            // </div>
          }

          {rootError && (
            <p className={theme.classes.formMessage}>{rootError}</p>
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
              disabled={isSubmitting}
              className={theme.classes.button.primary}
            >
              {isSubmitting ? "Creating..." : "Create"}
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
}
