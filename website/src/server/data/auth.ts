import axios, { AxiosResponse } from "axios";

import { env } from "@/env";

import { LoginFormData } from "@/forms/login/schema";
import { RegisterFormData } from "@/forms/register/schema";
import { OTPFormData } from "@/forms/otp/schema";

import { getServerError } from "../errors";
import { getCookies } from "../cookies";

const Routes = {
  login: "/auth/user/login",
  register: "/auth/user/register",
  emailVerification: "/auth/user/otp",
};

type AuthRouteReponse<T> = AxiosResponse<T | { error: string }>;

export const userLogin = async (formData: LoginFormData): Promise<string> => {
  type LoginResponse = AuthRouteReponse<{
    access_token: string;
  }>;

  try {
    const res: LoginResponse = await axios.post(
      env.NEXT_PUBLIC_API_BASE_URL + Routes.login,
      formData,
    );

    if ("error" in res.data) {
      throw res.data;
    }

    return res.data.access_token;
  } catch (err) {
    throw getServerError(err);
  }
};

export const userRegister = async (
  formData: RegisterFormData,
): Promise<string> => {
  type RegisterResponse = AuthRouteReponse<{
    register_id: string;
  }>;

  try {
    const res: RegisterResponse = await axios.post(
      env.NEXT_PUBLIC_API_BASE_URL + Routes.register,
      {
        email: formData.email,
        password: formData.password,
        provider: "app",
      },
    );

    if ("error" in res.data) {
      throw res.data;
    }

    return res.data.register_id;
  } catch (err) {
    throw getServerError(err);
  }
};

export const emailVerification = async (
  formData: OTPFormData,
): Promise<string> => {
  type verificationResponse = AuthRouteReponse<{
    access_token: string;
  }>;

  try {
    const registerID = getCookies("register_id");

    if (!registerID) throw Error("Register failed. Please retry.");

    const res: verificationResponse = await axios.post(
      env.NEXT_PUBLIC_API_BASE_URL + Routes.emailVerification,
      {
        otp: formData.otp,
        register_id: registerID,
      },
    );

    if ("error" in res.data) {
      throw res.data;
    }

    return res.data.access_token;
  } catch (err) {
    throw getServerError(err);
  }
};
