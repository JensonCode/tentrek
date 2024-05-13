import { z } from "zod";

export const OTPFormSchema = z.object({
  otp: z
    .string()
    .min(6, "Please enter the 6-digit OTP from your email")
    .max(6, "Please enter the 6-digit OTP from your email"),
});

export type OTPFormData = z.infer<typeof OTPFormSchema>;
