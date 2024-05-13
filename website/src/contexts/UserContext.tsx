"use client";

import { createContext, useContext } from "react";

export type User = {
  uid: string;
  email: string;
  username: string;
  avatar: string;
};

export const UserContext = createContext<User | null>(null);

function UserProvider({
  user,
  children,
}: {
  user: User | null;
  children: React.ReactNode;
}) {
  return <UserContext.Provider value={user}>{children}</UserContext.Provider>;
}

const useUserInfo = (): User | null => {
  const context = useContext<User | null>(UserContext);

  if (context) return context;
  return null;
};

export { UserProvider, useUserInfo };
