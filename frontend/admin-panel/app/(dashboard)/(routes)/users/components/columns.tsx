"use client";

import { ColumnDef } from "@tanstack/react-table";
import { CellAction } from "./cell-action";

export type UserColumn = {
    id: number;
    email: string;
    uid: string;
    registerDate: string;
    isDisabled: boolean;
    userStatus: string;
};

export const columns: ColumnDef<UserColumn>[] = [
    {
        accessorKey: "email",
        header: "Email",
    },
    {
        accessorKey: "uid",
        header: "Firebase Uid",
    },
    {
        accessorKey: "registerDate",
        header: "Register Date",
    },

    {
        accessorKey: "isDisabled",
        header: "Account Status",
        cell: ({ row }) => (
            row.original.isDisabled ? "Suspended" : "Active"
        ),
    },
    {
        accessorKey: "userStatus",
        header: "Account Level",
        cell: ({ row }) => {
            const status = row.original.userStatus;
            let color = "black";
            let text = "UNKNOWN";

            if (status === 'SUPERADMIN_STATUS') {
                color = "red";
                text = "SUPER ADMIN";
            } else if (status === 'ADMIN_STATUS') {
                color = "green";
                text = "ADMIN";
            } else if (status === 'CLIENT_STATUS') {
                color = "black";
                text = "CLIENT";
            }

            return (
                <span style={{ color }}>
                    {text}
                </span>
            );
        },
    },
    {
        id: "actions",
        cell: ({ row }) => <CellAction data={row.original} />,
    },
];
