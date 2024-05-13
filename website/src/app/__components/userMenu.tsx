"use client";

import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";

import { Avatar, AvatarImage } from "@/components/ui/avatar";

import { useUserInfo } from "@/contexts/UserContext";

import { setCookies } from "@/server/cookies";
import { useRouter } from "next/navigation";

export default function UserMenu() {
  const user = useUserInfo();
  const router = useRouter();

  const avatarImageURL =
    user && !!user.avatar ? user.avatar : "/avatar-placeholder.jpg";

  const handleLogout = async (e: React.MouseEvent<HTMLDivElement>) => {
    e.preventDefault();
    setCookies("access_token", "", 0);
    router.refresh();
  };

  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild className="border-red-500 hover:border">
        <Avatar>
          <AvatarImage src={avatarImageURL} />
        </Avatar>
      </DropdownMenuTrigger>

      <DropdownMenuContent className="mr-2">
        <DropdownMenuLabel>My Account</DropdownMenuLabel>
        <DropdownMenuSeparator />
        <DropdownMenuItem onClick={handleLogout}>Logout</DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
}
