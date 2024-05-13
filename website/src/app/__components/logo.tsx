import { Radio_Canada } from "next/font/google";
import { cn } from "@/lib/utils";

import Image from "next/image";
import Link from "next/link";

const font = Radio_Canada({ subsets: ["latin"] });

export function Logo() {
  return (
    <Link href="/" className={cn(font.className, "flex items-center")}>
      <Image src="/tt-icon.png" height={56} width={56} alt="icon" priority />
      <span className="text-lg font-bold max-md:hidden">Tentrek</span>
    </Link>
  );
}
