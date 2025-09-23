import type { NextRequest } from "next/server";

const HOP_BY_HOP = new Set([
  "connection",
  "keep-alive",
  "proxy-authenticate",
  "proxy-authorization",
  "te",
  "trailers",
  "transfer-encoding",
  "upgrade",
]);

// Allow-list methods you want to forward
const ALLOWED = new Set(["GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"]);

// Which cookie to read
const AUTH_COOKIE = "accessToken"; // change to yours

function stripHopByHop(h: Headers) {
  const out = new Headers(h);
  for (const k of HOP_BY_HOP) out.delete(k);
  return out;
}

async function handle(req: NextRequest, path: string[]) {
  if (!ALLOWED.has(req.method)) {
    return new Response("Method Not Allowed", { status: 405 });
  }

  const base = process.env.BACKEND_URL;
  if (!base) return new Response("Missing UPSTREAM_BASE", { status: 500 });

  // Join the rest of the path and preserve the query string
  const tail = path?.length ? `/api/v1/${path.join("/")}` : "/api/v1";
  const url = new URL(tail, base);
  const qs = new URLSearchParams(req.nextUrl.searchParams);
  qs.delete("path");
  url.search = qs.toString();
  console.log(`url: ${url}`);

  const headers = stripHopByHop(req.headers);
  const token = req.cookies.get(AUTH_COOKIE)?.value;

  if (token) headers.set("authorization", `Bearer ${token}`);

  const hasBody = !["GET", "HEAD"].includes(req.method);
  const body = hasBody ? req.body : undefined;
  const init: RequestInit = {
    method: req.method,
    headers,
    body,
    // Avoid caching at the proxy layer by default
    cache: "no-store",
    redirect: "manual",
    // If your upstream needs credentials for CORS preflight, mirror mode:
    // credentials: "include",
  };

  if (hasBody) {
    init.body = body as any;
    (init as any).duplex = "half";
  }

  // Forward the request
  try {
    const upstreamRes = await fetch(url, init);
    // Mirror upstream response headers/status; strip hop-by-hop on the way back too
    const respHeaders = stripHopByHop(upstreamRes.headers);
    return new Response(upstreamRes.body, {
      status: upstreamRes.status,
      statusText: upstreamRes.statusText,
      headers: respHeaders,
    });
  } catch (error) {
    console.error(error);
    return new Response(null, {
      status: 500,
      statusText: "Internal server error",
    });
  }
}

// Route handlers for all HTTP verbs:
// In Next 14.3+/15, params is a Promise—await it before use.
type Ctx = { params: Promise<{ path?: string[] }> };

export async function GET(req: NextRequest, ctx: Ctx) {
  const { path = [] } = await ctx.params;
  return handle(req, path);
}
export async function POST(req: NextRequest, ctx: Ctx) {
  const { path = [] } = await ctx.params;
  return handle(req, path);
}
export async function PUT(req: NextRequest, ctx: Ctx) {
  const { path = [] } = await ctx.params;
  return handle(req, path);
}
export async function PATCH(req: NextRequest, ctx: Ctx) {
  const { path = [] } = await ctx.params;
  return handle(req, path);
}
export async function DELETE(req: NextRequest, ctx: Ctx) {
  const { path = [] } = await ctx.params;
  return handle(req, path);
}
export async function OPTIONS(req: NextRequest, ctx: Ctx) {
  const { path = [] } = await ctx.params;
  return handle(req, path);
}
