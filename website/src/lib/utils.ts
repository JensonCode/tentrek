import { type ClassValue, clsx } from "clsx";
import { twMerge } from "tailwind-merge";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export const formatDate = (date: Date | null) => {
  if (!date) {
    return "In Progress";
  }

  const result = new Date(date);
  return result.toISOString().split("T")[0];
};
