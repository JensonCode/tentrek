"use client";

import React, { useState } from "react";
import { useRouter } from "next/navigation";

import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";

import { Form } from "@/components/ui/form";
import FormInput, { Field } from "@/components/form-input";
import { Button } from "@/components/ui/button";

import { useEmailVerification } from "@/hooks/useEmailVerification";
import { useCounter } from "@/hooks/useCounter";

import { setCookies } from "@/server/cookies";

import { OTPFormData, OTPFormSchema } from "./schema";

const OTPFormFields: Field<OTPFormData>[] = [
  {
    name: "otp",
    label: "OTP",
    variants: "otp",
  },
];

export default function OTPForm() {
  const router = useRouter();

  const [msg, setMsg] = useState<string>("");
  const [enableResend, setEnableResend] = useState(false);

  const form = useForm<OTPFormData>({
    resolver: zodResolver(OTPFormSchema),
  });

  const { mutate, isError } = useEmailVerification();

  const { seconds } = useCounter(100);

  const onSubmit = (formData: OTPFormData) => {
    mutate(formData, {
      onSuccess: (token) => {
        setCookies("access_token", token);
        router.push("/");
      },
      onError: (err) => {
        setMsg(err.message);
        setEnableResend(true);
      },
    });
  };

  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit(onSubmit)}
        className="mt-2 flex flex-col items-center justify-center space-y-8"
        id="otp-form"
      >
        {OTPFormFields.map((field) => (
          <FormInput
            key={`otp-form-${field.name}`}
            control={form.control}
            name={field.name}
            label={field.label}
            variants={field.variants}
            description={`Enter the OTP in ${seconds} seconds`}
          />
        ))}

        {isError && (
          <div className="mt-4 font-semibold text-red-500">
            <span>{msg}</span>
          </div>
        )}

        <div className="flex items-center justify-center space-x-10">
          <Button
            type="submit"
            variant="default"
            className="basis-1/2"
            form="otp-form"
            disabled={seconds === 0}
          >
            Confirm
          </Button>

          {(enableResend || seconds === 0) && (
            <Button
              type="button"
              variant="destructive"
              className="basis-1/2"
              onClick={() => router.back()}
            >
              Resend OTP
            </Button>
          )}
        </div>
      </form>
    </Form>
  );
}
