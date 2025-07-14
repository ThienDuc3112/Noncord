import { AuthLayout } from "@/components/auth-layout";
import { theme } from "@/lib/theme";
import RegisterForm from "./registerForm";
import { cookies } from "next/headers";
import { redirect } from "next/navigation";

export default async function RegisterPage() {
  const cookieStore = await cookies();
  const accessToken = cookieStore.get("accessToken");
  if (accessToken) redirect("/app");
  return (
    <AuthLayout>
      <div className="text-center mb-6">
        <h1 className={`text-2xl font-bold ${theme.colors.text.primary} mb-2`}>
          Create an account
        </h1>
        <p className={`${theme.colors.text.secondary} text-sm`}>
          Join our amazing community!
        </p>
      </div>

      <RegisterForm />
    </AuthLayout>
  );
}
