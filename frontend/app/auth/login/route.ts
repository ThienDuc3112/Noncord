import { NextRequest, NextResponse } from "next/server";
import { z } from "zod";
import axios from "axios";
import { apiClient } from "@/lib/request";

const loginBody = z.object({
  username: z
    .string({ message: "username must be a string" })
    .nonempty("username is required"),
  password: z
    .string({ message: "password must be a string" })
    .nonempty("password is required"),
});

export async function POST(req: NextRequest) {
  const body = await req.json();
  const parseRes = await loginBody.safeParseAsync(body);

  if (!parseRes.success) {
    return NextResponse.json(
      { error: parseRes.error.flatten().fieldErrors },
      { status: 400 },
    );
  }
  try {
    const tokens = (
      await apiClient.post<{
        accessToken: string;
        refreshToken: string;
      }>("/auth/login", parseRes.data)
    ).data;

    const res = new NextResponse(null, { status: 204 });
    res.cookies.set("accessToken", tokens.accessToken, {
      path: "/",
      maxAge: 30 * 60, // 30 min
      sameSite: "lax",
    });
    res.cookies.set("refreshToken", tokens.refreshToken, {
      httpOnly: true,
      secure: process.env.NODE_ENV === "production",
      sameSite: "lax",
      maxAge: 30 * 24 * 3600,
      path: "/auth",
    });

    return res;
  } catch (error) {
    if (axios.isAxiosError(error) && error.response) {
      console.error(error.response);
      return NextResponse.json(error.response.data, {
        status: error.response.status,
      });
    }

    console.error("Auth service error", error);
    return NextResponse.json(
      { error: "Internal server error" },
      { status: 500 },
    );
  }
}
