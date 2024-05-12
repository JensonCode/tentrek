"use server";

import axios, { AxiosResponse } from "axios";

import { env } from "@/env";

import { LoginFormData } from "@/schema/loginForm";
import { RegisterFormData } from "@/schema/registerForm";
import { OTPFormData } from "@/schema/otpForm";
import { cookies } from "next/headers";
import { getServerError } from "./errors";

const Routes = {
  login: "/auth/user/login",
  otp: "/auth/user/otp",
  register: "/auth/user/register",
};

export const userLogin = async (formData: LoginFormData): Promise<string> => {
  type LoginResponse = AxiosResponse<{
    access_token: string;
  }>;

  try {
    const res: LoginResponse = await axios.post(
      env.NEXT_PUBLIC_API_BASE_URL + Routes.login,
      formData,
    );

    if (!res.data.access_token) {
      throw res.data;
    }

    return res.data.access_token;
  } catch (err) {
    throw getServerError(err);
  }
};

export const sendOTPEmail = async (
  formData: RegisterFormData,
): Promise<string> => {
  type RegisterResponse = AxiosResponse<{
    register_id: string;
  }>;

  try {
    const res: RegisterResponse = await axios.post(
      env.NEXT_PUBLIC_API_BASE_URL + Routes.otp,
      {
        email: formData.email,
        password: formData.password,
        provider: "app",
      },
    );

    return res.data.register_id;
  } catch (err) {
    throw getServerError(err);
  }
};

export const userRegister = async (formData: OTPFormData): Promise<string> => {
  type RegisterResponse = AxiosResponse<{
    access_token: string;
  }>;

  try {
    const registerID = cookies().get("register_id")?.value;

    const requestBody = {
      otp: formData.otp,
      register_id: registerID,
    };

    const res: RegisterResponse = await axios.post(
      env.NEXT_PUBLIC_API_BASE_URL + Routes.register,
      requestBody,
    );

    return res.data.access_token;
  } catch (err) {
    throw getServerError(err);
  }
};
