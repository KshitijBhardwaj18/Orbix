import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";
import { AuthRoutes, PublicRoutes } from "./routes";

export function middleware(req: NextRequest) {
  const token = req.cookies.get("authToken")?.value;
  const pathName = req.nextUrl.pathname;

  const isPublicRoute = PublicRoutes.includes(pathName);
  const isAuthRoute = AuthRoutes.includes(pathName);

  // If no token and route is protected → redirect to /signin
  if (!token && !isAuthRoute && !isPublicRoute) {
    const url = req.nextUrl.clone();
    url.pathname = "/signin";

    return NextResponse.redirect(url);
  }

  // If token exists and user visits an unprotected auth page → redirect to home
  if (token && isAuthRoute) {
    const url = req.nextUrl.clone();
    url.pathname = "/home";
    return NextResponse.redirect(url);
  }

  return NextResponse.next();
}

// Optional: exclude Next.js assets and favicon
export const config = {
  matcher: [
    "/((?!_next/static|_next/image|favicon\\.ico|.*\\.(?:png|jpg|jpeg|gif|webp|svg)$).*)",
  ],
};
