import { z } from "zod";

export const loginFormSchema = z.object({
  email: z
    .string()
    .min(1, "Email address is requried")
    .max(100, "Email address is over character limit")
    .email("Please enter a valid email"),
  password: z
    .string()
    .min(1, "Password is requried")
    .max(100, "Password is over character limit"),
});

export type LoginFormData = z.infer<typeof loginFormSchema>;
