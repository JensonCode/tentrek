"use client";

import React, { useState } from "react";
import { useRouter } from "next/navigation";
import Link from "next/link";

import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";

import { Form } from "@/components/ui/form";
import FormInput, { Field } from "@/components/form-input";
import { Button } from "@/components/ui/button";

import { useLogin } from "@/hooks/useLogin";

import { setCookies } from "@/server/cookies";

import { LoginFormData, loginFormSchema } from "./schema";

const loginFormFields: Field<LoginFormData>[] = [
  {
    name: "email",
    label: "Email",
    textType: "email",
  },
  {
    name: "password",
    label: "Password",
    textType: "password",
  },
];

export default function LoginForm() {
  const router = useRouter();

  const [msg, setMsg] = useState<string>("");

  const form = useForm<LoginFormData>({
    resolver: zodResolver(loginFormSchema),
  });

  const { mutate, isError } = useLogin();

  const onSubmit = (formData: LoginFormData) => {
    mutate(formData, {
      onSuccess: (token) => {
        setCookies("access_token", token, 24 * 60 * 60);
        router.push("/");
      },
      onError: (err) => {
        setMsg(err.message);
      },
    });
  };

  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit(onSubmit)}
        className="space-y-4"
        id="login-form"
      >
        {loginFormFields.map((field) => (
          <FormInput
            key={`login-form-${field.name}`}
            control={form.control}
            name={field.name}
            label={field.label}
            textType={field.textType}
          />
        ))}

        {isError && (
          <div className="mt-4 font-semibold text-red-500">
            <span>{msg}</span>
          </div>
        )}
      </form>

      <div className="mb-4 mt-8 flex items-center justify-center space-x-10">
        <Button
          type="submit"
          variant="default"
          size="lg"
          className="basis-1/2"
          form="login-form"
        >
          Login
        </Button>

        <Link href={"/auth/register"} className="w-full basis-1/2">
          <Button
            type="button"
            variant="secondary"
            size="lg"
            className="w-full"
          >
            Sign Up
          </Button>
        </Link>
      </div>
    </Form>
  );
}
