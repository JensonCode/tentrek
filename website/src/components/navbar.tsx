import { twMerge } from "tailwind-merge";

import { Button } from "@/components/ui/button";
import Logo from "./logo";
import { Link } from "lucide-react";

export default function Navbar() {
  return (
    <header
      className={twMerge(
        "flex items-center justify-between px-6",
        "w-full rounded-md drop-shadow-sm",
        "bg-white text-black",
      )}
    >
      <Logo />

      <div className="flex justify-end space-x-4">
        <Link href="/auth/login">
          <Button variant="secondary">Login</Button>
        </Link>
        <Link href="/auth/register">
          <Button>Register</Button>
        </Link>
      </div>
    </header>
  );
}
