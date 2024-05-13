"use server";

import { ResponseCookie } from "next/dist/compiled/@edge-runtime/cookies";
import { cookies } from "next/headers";

export const setCookies = (
  key: string,
  value: string,
  withExpirySecond?: number,
) => {
  const options: Partial<ResponseCookie> = {
    httpOnly: true,
  };

  if (withExpirySecond) {
    options.maxAge = withExpirySecond;
  }

  cookies().set(key, value, options);
};
