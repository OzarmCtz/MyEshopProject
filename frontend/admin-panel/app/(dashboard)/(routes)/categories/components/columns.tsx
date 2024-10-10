"use client";

import { ColumnDef } from "@tanstack/react-table";
import { CellAction } from "./cell-action";

export type CategoryColumn = {
    id: number;
    name: string;
    description: string;
    pictureUrl: string;
    occurence: number;
    totalItemCount: number;
};

export const columns: ColumnDef<CategoryColumn>[] = [
    {
        accessorKey: "name",
        header: "Name",
    },
    {
        accessorKey: "description",
        header: "Description",
    },
    {
        accessorKey: "pictureUrl",
        header: "Picture Url",
    },
    {
        accessorKey: "occurence",
        header: "Sub Categories Linked",
    },
    {
        accessorKey: "totalItemCount",
        header: "Total Items Linked",
    },
    {
        id: "actions",
        cell: ({ row }) => <CellAction data={row.original} />,
    },
];
