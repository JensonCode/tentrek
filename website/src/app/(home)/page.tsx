import styles from "@/styles/homePage.module.css";

import { DataTable } from "@/components/data-table";

import { Button } from "@/components/ui/button";

import { Roboto_Slab } from "next/font/google";
import { cn } from "@/lib/utils";

import { exampleColumns, exampleData } from "@/example/example-table";

const font = Roboto_Slab({ subsets: ["latin"] });

export default function HomePage() {
  return (
    <main className="flex min-h-[90dvh] flex-col items-center space-y-10">
      <div
        className={cn(
          font.style,
          "mt-10 flex flex-col space-y-4 text-center sm:space-y-10",
        )}
      >
        <h1 className="text-3xl font-bold lg:text-5xl">Know Where You Are</h1>
        <h2 className="text-xl font-bold lg:text-2xl">
          A Collaborative Platform <br />
          for Progress Tracking
        </h2>
      </div>

      <div className="w-[85%] overflow-x-auto sm:w-[65%]">
        <DataTable
          columns={exampleColumns}
          data={exampleData}
          className={styles.table}
        />
      </div>

      <div className="pb-6 sm:py-6">
        <Button size={"xl"}>Track</Button>
      </div>
    </main>
  );
}
