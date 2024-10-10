"use client";

import { ColumnDef } from "@tanstack/react-table";
import { CellAction } from "./cell-action";

export type SubCategoryColumn = {
    id: number;
    name: string;
    description: string;
    pictureUrl: string;
    categoryName: string;
    item_count: number,
};

export const columns: ColumnDef<SubCategoryColumn>[] = [
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
        accessorKey: "categoryName",
        header: "Linked Category",
    },
    {
        accessorKey: "item_count",
        header: "Number of Items Linked",
    },
    {
        id: "actions",
        cell: ({ row }) => <CellAction data={row.original} />,
    },
];
