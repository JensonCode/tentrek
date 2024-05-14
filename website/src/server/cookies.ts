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

  if (withExpirySecond !== undefined) {
    options.maxAge = withExpirySecond;
  }

  cookies().set(key, value, options);
};

export const getCookies = (key: string) => {
  return cookies().get(key)?.value;
};
