import { cn } from "@/lib/utils";
import ReactQueryProvider from "@/app/__components/ReactQueryProvider";
import "@/styles/globals.css";

import { Inter } from "next/font/google";
import { getUser } from "@/server/data/user";
import { UserProvider } from "@/contexts/UserContext";

const inter = Inter({
  subsets: ["latin"],
  variable: "--font-sans",
});

export const metadata = {
  title: "TenTrek",
  description: "TenTrek",
  icons: [{ rel: "icon", url: "/favicon.ico" }],
};

export default async function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const user = await getUser();

  return (
    <html lang="en">
      <body className={cn(`font-sans ${inter.variable}`, "bg-secondary")}>
        <ReactQueryProvider>
          <UserProvider user={user}>{children}</UserProvider>
        </ReactQueryProvider>
      </body>
    </html>
  );
}
