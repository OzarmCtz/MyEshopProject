"use client";

import { ColumnDef } from "@tanstack/react-table";
import { CellAction } from "./cell-action";

export type ItemColumn = {
    id: number;
    title: string;
    description: string;
    price: number;
    quantity: number;
    pictureUrl: string;
    filePath: string;
    disabled: boolean;
    release_date: string;
    sub_category_name: string,
};

export const columns: ColumnDef<ItemColumn>[] = [
    {
        accessorKey: "title",
        header: "Title",
    },
    {
        accessorKey: "description",
        header: "Description",
    },
    {
        accessorKey: "price",
        header: "Price",
        cell: ({ row }) => (
            <>
                {row.original.price} €
            </>
        ),
    },
    {
        accessorKey: "quantity",
        header: "Quantity",
    },
    {
        accessorKey: "pictureUrl",
        header: "Picture Url",
    },
    {
        accessorKey: "filePath",
        header: "File Path",
    },
    {
        accessorKey: "disabled",
        header: "Status",
        cell: ({ row }) => (
            row.original.disabled ? "Disabled" : "Enabled"
        ),
    },
    {
        accessorKey: "release_date",
        header: "Release Date",
    },
    {
        accessorKey: "sub_category_name",
        header: "Lindked Sub Category",
    },
    {
        id: "actions",
        cell: ({ row }) => <CellAction data={row.original} />,
    },
];
