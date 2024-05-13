import "server-only";

import { JwtPayload, jwtDecode } from "jwt-decode";
import { cookies } from "next/headers";

interface Payload extends JwtPayload {
  uid: string;
  exp: number;
  iat: number;
}

export const decodeToken = (token: string | undefined) => {
  if (!token) return;

  try {
    const decoded: Payload = jwtDecode(token);

    return decoded;
  } catch (error) {
    console.log(error);

    return undefined;
  }
};

export const deleteExpired = (payload: Payload) => {
  const now = Math.floor(Date.now() / 1000);
  if (payload.exp < now) {
    cookies().delete("access_token");
  }
};
