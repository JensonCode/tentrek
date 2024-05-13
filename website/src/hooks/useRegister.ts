import { userRegister } from "@/server/auth";
import { useMutation } from "@tanstack/react-query";

export const useRegister = () => {
  return useMutation({
    mutationFn: userRegister,
  });
};
