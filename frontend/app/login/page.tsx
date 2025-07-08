import { cookies } from "next/headers";
import { redirect } from "next/navigation";
import LoginForm from "./loginForm";

export default async function LoginPage() {
  const cookieStore = await cookies();
  const accessToken = cookieStore.get("accessToken");
  if (accessToken) redirect("/app");
  return <LoginForm />;
}
