"use client";

import { ColumnDef } from "@tanstack/react-table";

import { Avatar, AvatarImage } from "@/components/ui/avatar";

import { formatDate } from "@/lib/utils";

type ExampleTable = {
  userAvatar: string;

  start: Date;

  process: Date | null;

  end: Date | null;
};

export const exampleColumns: ColumnDef<ExampleTable>[] = [
  {
    accessorKey: "userAvatar",
    header: "User",
    cell: ({ row }) => {
      return (
        <Avatar>
          <AvatarImage src={row.getValue("userAvatar")} />
        </Avatar>
      );
    },
  },
  {
    accessorKey: "start",
    header: "Application",
    cell: ({ row }) => {
      const date: Date = row.getValue("start");
      return <span>{formatDate(date)}</span>;
    },
  },
  {
    accessorKey: "process",
    header: "Process",
    cell: ({ row }) => {
      const date: Date = row.getValue("process");
      return <span>{formatDate(date)}</span>;
    },
  },
  {
    accessorKey: "end",
    header: "Finish",
    cell: ({ row }) => {
      const date: Date = row.getValue("end");
      return <span>{formatDate(date)}</span>;
    },
  },
];

export const exampleData: ExampleTable[] = [
  {
    userAvatar: "/avatar-placeholder.jpg",
    start: new Date("2022-12-16"),
    process: new Date("2023-01-05"),
    end: new Date("2023-01-07"),
  },
  {
    userAvatar: "/homePage-me-avatar.jpg",
    start: new Date("2023-01-03"),
    process: null,
    end: null,
  },
  {
    userAvatar: "/avatar-placeholder.jpg",
    start: new Date("2023-01-04"),
    process: new Date("2023-01-31"),
    end: new Date("2023-02-03"),
  },
  {
    userAvatar: "/avatar-placeholder.jpg",
    start: new Date("2023-02-01"),
    process: new Date("2023-03-06"),
    end: null,
  },
];
