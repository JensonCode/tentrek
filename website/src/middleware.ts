import { cookies } from "next/headers";
import { NextRequest, NextResponse } from "next/server";
import { decodeToken, deleteExpired } from "./lib/jwt";

const getAccessToken = () => {
  const accessToken = cookies().get("access_token")?.value;
  const payload = decodeToken(accessToken);
  if (payload) deleteExpired(payload);

  return payload;
};

export default async function middleware(req: NextRequest) {
  const path = req.nextUrl.pathname;
  const response = NextResponse.next();

  const accessToken = getAccessToken();

  if (path.startsWith("/auth")) {
    if (!!accessToken && accessToken.uid) {
      return NextResponse.redirect(
        // new URL(`/dashboard/${accessToken.uid}`, req.nextUrl),
        new URL(`/`, req.nextUrl),
      );
    }
  }

  return response;
}
