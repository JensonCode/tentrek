import { notFound } from "next/navigation";

import { cookies } from "next/headers";
import Link from "next/link";
import { Button } from "@/components/ui/button";
import Image from "next/image";

import { env } from "@/env";
import LoginForm from "@/forms/login/form";
import RegisterForm from "@/forms/register/form";
import OTPForm from "@/forms/otp/form";

type PageProps = {
  params: {
    service: string;
  };
};

const paths = ["login", "register", "otp"];
const titles = ["Login", "Register", "Email Verification"];

export default function Page({ params }: PageProps) {
  const pathIndex = paths.findIndex((path) => params.service === path);

  if (pathIndex === -1) {
    return notFound();
  }

  if (pathIndex === 2 && !cookies().get("register_id")) {
    return notFound();
  }

  return (
    <div className="mx-auto w-full rounded-md bg-white px-6 py-6 drop-shadow-sm sm:max-w-md sm:px-8 ">
      <h1 className="pb-6 text-center text-2xl font-semibold">
        {titles[pathIndex]}
      </h1>
      {pathIndex === 0 ? (
        <LoginForm />
      ) : pathIndex === 1 ? (
        <RegisterForm />
      ) : (
        <OTPForm />
      )}

      {pathIndex !== 2 && (
        <div className="w-full border-t-2 border-gray-300 pt-4">
          <Link href={`${env.NEXT_PUBLIC_API_BASE_URL}/auth/google`}>
            <Button
              variant={"googleOAuth"}
              className="flex w-full justify-center space-x-4"
              size={"googleOAuth"}
            >
              <Image
                src="/google-oauth-icon.svg"
                alt="Google oauth"
                height={32}
                width={32}
              />
              <span>Login with Google</span>
            </Button>
          </Link>
        </div>
      )}
    </div>
  );
}
