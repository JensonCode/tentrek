import axios, { AxiosResponse } from "axios";

import { env } from "@/env";

import { User } from "@/contexts/UserContext";
import { cookies } from "next/headers";
import { decodeToken } from "@/lib/jwt";

export const getUser = async (): Promise<User | null> => {
  try {
    const token = cookies().get("access_token")?.value;
    const payload = decodeToken(token);

    if (!payload) {
      return null;
    }

    const res: AxiosResponse<User> = await axios.get(
      env.NEXT_PUBLIC_API_BASE_URL + "/user/" + payload.uid,
      {
        headers: {
          Authorization: `bearer ${token}`,
        },
      },
    );

    return res.data;
  } catch (error) {
    return null;
  }
};
