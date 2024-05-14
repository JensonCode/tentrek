import { emailVerification } from "@/server/data/auth";
import { useMutation } from "@tanstack/react-query";

export const useEmailVerification = () => {
  return useMutation({
    mutationFn: emailVerification,
  });
};
