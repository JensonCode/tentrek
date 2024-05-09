import { Radio_Canada } from "next/font/google";
import { cn } from "@/lib/utils";

import Image from "next/image";
import Link from "next/link";

const font = Radio_Canada({ subsets: ["latin"] });

export default function Logo() {
  return (
    <Link href="/" className={cn(font.className, "flex items-center")}>
      <Image src="/tt-icon.png" height={56} width={56} alt="icon" />
      <span className="text-lg font-bold max-md:hidden">Track Together</span>
    </Link>
  );
}