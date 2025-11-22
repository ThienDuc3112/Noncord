"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import Cookies from "universal-cookie";
import { isAxiosError } from "axios";
import { Eye, EyeOff } from "lucide-react";

import { apiClient } from "@/lib/request";
import { LoginSchema, TokenSchema, type LoginData } from "@/lib/types";
import { theme } from "@/lib/theme";

import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Label } from "@/components/ui/label";

const cookies = new Cookies();

const LoginForm = () => {
  const router = useRouter();
  const [showPassword, setShowPassword] = useState(false);

  const form = useForm<LoginData>({
    resolver: zodResolver(LoginSchema),
    defaultValues: {
      username: "",
      password: "",
    },
  });

  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
    setError,
  } = form;

  const onSubmit = async (data: LoginData) => {
    try {
      const res = await apiClient.post("/auth/login", data);
      const body = TokenSchema.safeParse(res.data);

      if (!body.success) {
        return setError("root", {
          message: body.error.errors.map((e) => e.message).join("\n"),
        });
      }

      cookies.set("accessToken", body.data.accessToken, {
        sameSite: true,
        secure: true,
        path: "/",
      });
      cookies.set("refreshToken", body.data.refreshToken, {
        sameSite: true,
        secure: true,
        path: "/",
      });

      router.push("/app");
    } catch (error: unknown) {
      console.error(error);
      if (
        isAxiosError(error) &&
        error.response &&
        error.response.status < 500
      ) {
        const errorMsg: string | undefined = (error.response.data as any)
          ?.error;
        if (errorMsg) {
          return setError("root", { message: errorMsg });
        }
      }
      setError("root", {
        message: "Unknown error occured, please try again some time later",
      });
    }
  };

  const rootError = errors.root?.message;

  return (
    <form className="space-y-5" onSubmit={handleSubmit(onSubmit)} noValidate>
      {/* Username */}
      <div className="space-y-1">
        <Label htmlFor="username" className={theme.classes.label}>
          Username or Email
        </Label>
        <Input
          id="username"
          type="text"
          autoComplete="username"
          className={
            errors.username ? theme.classes.inputError : theme.classes.input
          }
          {...register("username")}
        />
        {errors.username && (
          <p className={theme.classes.formMessage}>{errors.username.message}</p>
        )}
      </div>

      {/* Password */}
      <div className="space-y-1">
        <Label htmlFor="password" className={theme.classes.label}>
          Password
        </Label>
        <div className="relative">
          <Input
            id="password"
            type={showPassword ? "text" : "password"}
            autoComplete="current-password"
            className={
              errors.password
                ? `${theme.classes.inputError} pr-10`
                : `${theme.classes.input} pr-10`
            }
            {...register("password")}
          />
          <button
            type="button"
            onClick={() => setShowPassword((v) => !v)}
            className="absolute inset-y-0 right-0 flex items-center pr-3 text-[#a5adcb]"
          >
            {showPassword ? (
              <EyeOff className="h-4 w-4" />
            ) : (
              <Eye className="h-4 w-4" />
            )}
          </button>
        </div>
        {errors.password && (
          <p className={theme.classes.formMessage}>{errors.password.message}</p>
        )}
      </div>

      {/* Root errors */}
      {rootError && (
        <div className="rounded-sm bg-[#3b1f2b]/80 px-3 py-2 text-sm text-[#ed8796]">
          {rootError}
        </div>
      )}

      <Button
        type="submit"
        disabled={isSubmitting}
        className={`w-full ${theme.classes.button.primary}`}
      >
        {isSubmitting ? "Signing in..." : "Sign in"}
      </Button>
    </form>
  );
};

export default LoginForm;
