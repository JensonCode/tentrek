import { z } from "zod";

export const registerFormSchema = z
  .object({
    email: z
      .string()
      .min(1, "Email address is requried")
      .max(100, "Email address is over character limit")
      .email("Please enter a valid email"),
    password: z
      .string()
      .min(1, "Password is requried")
      .max(100, "Password is over character limit"),
    confirmPassword: z
      .string()
      .min(1, "Confirm Password is requried")
      .max(100, "Confirm Password is over character limit"),
  })
  .refine((data) => data.password === data.confirmPassword, {
    message: "Passwords do not match",
    path: ["confirmPassword"],
  });

export type RegisterFormData = z.infer<typeof registerFormSchema>;
