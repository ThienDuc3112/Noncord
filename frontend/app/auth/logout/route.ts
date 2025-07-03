import { NextRequest, NextResponse } from "next/server";
import axios from "axios";
import { apiClient } from "@/lib/request";

export async function POST(req: NextRequest) {
  const refreshToken = req.cookies.get("refreshToken");
  if (!refreshToken || refreshToken.value == "")
    return NextResponse.json({ error: "No session found" }, { status: 400 });

  let ret: NextResponse;
  try {
    await apiClient.post<undefined>("/auth/refresh", {
      refreshToken: refreshToken.value,
    });
    ret = new NextResponse(null, { status: 204 });
  } catch (error) {
    if (axios.isAxiosError(error) && error.response) {
      console.error(error.response);
      ret = NextResponse.json(error.response.data, {
        status: error.response.status,
      });
    } else {
      console.error("Auth service error", error);
      ret = NextResponse.json(
        { error: "Internal server error" },
        { status: 500 },
      );
    }
  }

  ret.cookies.set("accessToken", "", {
    path: "/",
    maxAge: 0,
  });
  ret.cookies.set("refreshToken", "", {
    path: "/auth",
    maxAge: 0,
  });

  return ret;
}
