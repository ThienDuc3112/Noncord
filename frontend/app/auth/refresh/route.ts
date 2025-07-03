import { NextRequest, NextResponse } from "next/server";
import axios from "axios";
import { apiClient } from "@/lib/request";

export async function POST(req: NextRequest) {
  const refreshToken = req.cookies.get("refreshToken");
  if (!refreshToken || refreshToken.value == "")
    return NextResponse.json({ error: "No session found" }, { status: 400 });

  try {
    const tokens = (
      await apiClient.post<{
        accessToken: string;
        refreshToken: string;
      }>("/auth/refresh", { refreshToken: refreshToken.value })
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
