import { z } from "zod";

const OTPFormSchema = z.object({
  otp: z
    .string()
    .min(6, "Invalid OTP: Please enter the 6-digit OTP")
    .max(6, "Invalid OTP: Please enter the 6-digit OTP"),
});

const OTPFormFields = [
  {
    name: "otp",
    label: "One Time Passcode",
    type: "otp",
  },
];

export { OTPFormSchema, OTPFormFields };
export type OTPFormData = z.infer<typeof OTPFormSchema>;
