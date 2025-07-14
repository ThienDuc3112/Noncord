"use client";

import { useState } from "react";
import Link from "next/link";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { AlertCircle, Eye, EyeOff } from "lucide-react";
import { theme } from "@/lib/theme";
import { z } from "zod";
import { useRouter } from "next/navigation";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { apiClient } from "@/lib/request";
import { isAxiosError } from "axios";

const RegisterSchema = z
  .object({
    username: z
      .string({ message: "Username must be a string" })
      .min(3, "Username must be between 3 and 32 characters")
      .max(32, "Username must be between 3 and 32 characters")
      .regex(
        /^[a-zA-Z0-9_-](?:[a-zA-Z0-9_-]{1,30}[a-zA-Z0-9_-])$/,
        "Must contain only alphanumeric and '-' or '_' characters",
      ),
    email: z
      .string({ message: "Email must be a string" })
      .email("Must be a valid email")
      .nonempty("Email cannot be empty"),
    password: z
      .string({ message: "Password must be a string" })
      .min(8, "Password too short, must be between 8 and 72 characters")
      .max(72, "Password too long, must be between 8 and 72 characters"),
    confirmPassword: z.string({ message: "Password must be a string" }),
  })
  .refine((o) => o.password === o.confirmPassword, {
    message: "Confirm password must be the same as password",
    path: ["confirmPassword"],
  });

type RegisterData = z.infer<typeof RegisterSchema>;

const RegisterForm = () => {
  const router = useRouter();
  const form = useForm<RegisterData>({
    resolver: zodResolver(RegisterSchema),
    defaultValues: {
      username: "",
      email: "",
      password: "",
      confirmPassword: "",
    },
  });

  const [showPassword, setShowPassword] = useState(false);

  const handleSubmit = async (data: RegisterData) => {
    try {
      await apiClient.post("/auth/register", data);
      router.push("/login");
    } catch (error) {
      console.error(error);
      if (
        isAxiosError(error) &&
        error.response &&
        error.response.status < 500
      ) {
        const errorMsg: string | undefined = error.response.data?.error;
        if (errorMsg) {
          return form.setError("root", { message: errorMsg });
        }
      }
      form.setError("root", {
        message: "Unknown error occured, please try again some time later",
      });
    }
    // Simulate successful registration
    // navigate.push("/dashboard");
  };

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(handleSubmit)} className="space-y-4">
        <FormField
          control={form.control}
          name="username"
          render={({ field, fieldState }) => (
            <FormItem>
              <FormLabel className={theme.classes.label}>
                Username
                <span className={theme.colors.states.error}>*</span>
              </FormLabel>
              <FormControl>
                <Input
                  {...field}
                  className={
                    fieldState.error
                      ? theme.classes.inputError
                      : theme.classes.input
                  }
                />
              </FormControl>
              <FormMessage className={theme.classes.formMessage} />
            </FormItem>
          )}
        />

        <FormField
          control={form.control}
          name="email"
          render={({ field, fieldState }) => (
            <FormItem>
              <FormLabel className={theme.classes.label}>
                Email<span className={theme.colors.states.error}>*</span>
              </FormLabel>
              <FormControl>
                <Input
                  {...field}
                  type="email"
                  className={
                    fieldState.error
                      ? theme.classes.inputError
                      : theme.classes.input
                  }
                />
              </FormControl>
              <FormMessage className={theme.classes.formMessage} />
            </FormItem>
          )}
        />

        <FormField
          control={form.control}
          name="password"
          render={({ field, fieldState }) => (
            <FormItem>
              <FormLabel className={theme.classes.label}>
                Password<span className={theme.colors.states.error}>*</span>
              </FormLabel>
              <div className="relative">
                <FormControl>
                  <Input
                    {...field}
                    type={showPassword ? "text" : "password"}
                    className={
                      (fieldState.error
                        ? theme.classes.inputError
                        : theme.classes.input) + " pr-10"
                    }
                  />
                </FormControl>
                <button
                  type="button"
                  onClick={() => setShowPassword(!showPassword)}
                  className={`absolute right-3 top-1/2 -translate-y-1/2 ${theme.colors.text.secondary} hover:text-white`}
                >
                  {showPassword ? (
                    <EyeOff className="h-4 w-4" />
                  ) : (
                    <Eye className="h-4 w-4" />
                  )}
                </button>
              </div>
              <FormMessage className={theme.classes.formMessage} />
            </FormItem>
          )}
        />

        <FormField
          control={form.control}
          name="confirmPassword"
          render={({ field, fieldState }) => (
            <FormItem>
              <FormLabel className={theme.classes.label}>
                Confirm password
                <span className={theme.colors.states.error}>*</span>
              </FormLabel>
              <div className="relative">
                <FormControl>
                  <Input
                    {...field}
                    type={showPassword ? "text" : "password"}
                    className={
                      (fieldState.error
                        ? theme.classes.inputError
                        : theme.classes.input) + " pr-10"
                    }
                  />
                </FormControl>
                <button
                  type="button"
                  onClick={() => setShowPassword(!showPassword)}
                  className={`absolute right-3 top-1/2 -translate-y-1/2 ${theme.colors.text.secondary} hover:text-white`}
                >
                  {showPassword ? (
                    <EyeOff className="h-4 w-4" />
                  ) : (
                    <Eye className="h-4 w-4" />
                  )}
                </button>
              </div>
              <FormMessage className={theme.classes.formMessage} />
            </FormItem>
          )}
        />

        {form.formState.errors.root && (
          <div className="flex items-center space-x-2 p-3 bg-red-500/10 border border-red-500/20 rounded-md">
            <AlertCircle className="w-4 h-4 text-red-400 flex-shrink-0" />
            <p className={theme.classes.formMessage}>
              {form.formState.errors.root.message}
            </p>
          </div>
        )}

        <Button
          type="submit"
          className={`w-full ${theme.classes.button.primary}`}
          disabled={form.formState.isSubmitting}
        >
          Continue
        </Button>

        <p className={`${theme.colors.text.muted} text-xs leading-relaxed`}>
          By registering, you agree to Noncord's{" "}
          <button className={theme.colors.interactive.link}>
            Terms of Service
          </button>{" "}
          and{" "}
          <button className={theme.colors.interactive.link}>
            Privacy Policy
          </button>
          .
        </p>

        <Link
          href="/login"
          className={`${theme.colors.interactive.link} text-sm w-full text-center block`}
        >
          Already have an account?
        </Link>
      </form>
    </Form>
  );
};

export default RegisterForm;
