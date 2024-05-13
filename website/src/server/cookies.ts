"use server";

import { cookies } from "next/headers";

export const setCookies = (key: string, value: string) => {
  cookies().set(key, value);
};

export const deleteCookies = (key: string) => {
  cookies().delete(key);
};
