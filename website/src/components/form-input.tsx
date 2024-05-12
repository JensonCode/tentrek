import React from "react";

import { Input } from "@/components/ui/input";

import {
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";

import {
  InputOTP,
  InputOTPGroup,
  InputOTPSlot,
} from "@/components/ui/input-otp";

import { Control, FieldValues, Path } from "react-hook-form";

export type Field<FormData> = {
  name: Path<FormData>;
  label: string;
  textType?: "text" | "password" | "email";
  variants?: "text" | "otp";
  placeholder?: string;
  description?: string;
};

type FormInputProps<T extends FieldValues> = {
  control: Control<T>;
} & Field<T>;

export default function FormInput<T extends FieldValues>({
  control,
  name,
  label,
  placeholder,
  description,
  variants,
  textType,
}: FormInputProps<T>) {
  return (
    <FormField
      control={control}
      name={name}
      render={({ field }) => (
        <FormItem>
          <FormLabel>{label}</FormLabel>
          <FormControl>
            {variants === "otp" ? (
              <InputOTP maxLength={6} {...field}>
                <InputOTPGroup>
                  <InputOTPSlot index={0} />
                  <InputOTPSlot index={1} />
                  <InputOTPSlot index={2} />
                  <InputOTPSlot index={3} />
                  <InputOTPSlot index={4} />
                  <InputOTPSlot index={5} />
                </InputOTPGroup>
              </InputOTP>
            ) : (
              <Input placeholder={placeholder} type={textType} {...field} />
            )}
          </FormControl>
          {!!description && <FormDescription>{description}</FormDescription>}
          <FormMessage />
        </FormItem>
      )}
    />
  );
}
