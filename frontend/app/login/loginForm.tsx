"use client";
import { useState } from "react";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { AlertCircle, Eye, EyeOff } from "lucide-react";
import { theme } from "@/lib/theme";
import z from "zod";
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
import axios, { isAxiosError } from "axios";

const LoginSchema = z.object({
  username: z
    .string({ message: "Username must be a string" })
    .nonempty("Username or email must exist"),
  password: z
    .string({ message: "Password must be a string" })
    .nonempty("Password must exist"),
});

type LoginData = z.infer<typeof LoginSchema>;

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

  const handleSubmit = async (data: LoginData) => {
    try {
      await axios.post("/auth/login", data, { withCredentials: true });
      router.push("/app");
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
                Email or Username
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

        {form.formState.errors.root && (
          <div className="flex items-center space-x-2 p-3 bg-red-500/10 border border-red-500/20 rounded-md">
            <AlertCircle className="w-4 h-4 text-red-400 flex-shrink-0" />
            <p className={theme.classes.formMessage}>
              {form.formState.errors.root.message}
            </p>
          </div>
        )}

        <button
          type="button"
          className={`${theme.colors.interactive.link} text-sm`}
        >
          Forgot your password?
        </button>

        <Button
          type="submit"
          className={`w-full ${theme.classes.button.primary}`}
        >
          Log In
        </Button>

        <p className={`${theme.colors.text.muted} text-sm`}>
          Need an account?{" "}
          <Link href="/register" className={theme.colors.interactive.link}>
            Register
          </Link>
        </p>
      </form>
    </Form>
  );
};

export default LoginForm;
