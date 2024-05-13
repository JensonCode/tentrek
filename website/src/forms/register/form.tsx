"use client";

import React, { useState } from "react";
import { useRouter } from "next/navigation";
import Link from "next/link";

import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";

import { Form } from "@/components/ui/form";
import FormInput, { Field } from "@/components/form-input";
import { Button } from "@/components/ui/button";

import { useRegister } from "@/hooks/useRegister";

import { setCookies } from "@/server/cookies";

import { RegisterFormData, registerFormSchema } from "./schema";

const registerFormFields: Field<RegisterFormData>[] = [
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
  {
    name: "confirmPassword",
    label: "Confirm Password",
    textType: "password",
  },
];

export default function RegisterForm() {
  const router = useRouter();

  const [msg, setMsg] = useState<string>("");

  const form = useForm<RegisterFormData>({
    resolver: zodResolver(registerFormSchema),
  });

  const { mutate, isError } = useRegister();

  const onSubmit = (formData: RegisterFormData) => {
    mutate(formData, {
      onSuccess: (registerID) => {
        setCookies("register_id", registerID, 100);
        router.push("/auth/otp");
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
        id="register-form"
      >
        {registerFormFields.map((field) => (
          <FormInput
            key={`register-form-${field.name}`}
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
          form="register-form"
          className="basis-1/2"
        >
          Sign Up
        </Button>

        <Link href={"/auth/login"}>
          <Button
            type="button"
            variant="secondary"
            size="lg"
            className="w-full"
          >
            Already have account
          </Button>
        </Link>
      </div>
    </Form>
  );
}
