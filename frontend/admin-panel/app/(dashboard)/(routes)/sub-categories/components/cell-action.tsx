"use client";

import { DropdownMenu, DropdownMenuTrigger, DropdownMenuContent, DropdownMenuItem } from "@/components/ui/dropdown-menu";
import { SubCategoryColumn } from "./columns";
import { Button } from "@/components/ui/button";
import { Copy, Edit, MoreHorizontal, Trash } from "lucide-react";
import { DropdownMenuLabel } from "@radix-ui/react-dropdown-menu";
import { toast } from "react-hot-toast";
import { useRouter, useParams } from "next/navigation";
import { useState } from "react";
import { AlertModal } from "@/components/modals/alert-modal";
interface CellActionProps {
    data: SubCategoryColumn;
};


export const CellAction: React.FC<CellActionProps> = ({
    data
}) => {
    const router = useRouter();
    const params = useParams();

    const [loading, setLoading] = useState(false);
    const [open, setOpen] = useState(false);


    const onCopy = (id: string) => {
        navigator.clipboard.writeText(id)
        toast.success('ID copied to clipboard')
    };

    const onDelete = async () => {
   

        try {
            setLoading(true);
            const response = await fetch(`/api/sub-categories?id=${data.id}`, {
                method: 'DELETE',
                
            });

            const result = await response.json();

            if (response.status === 200) {
                toast.success("Sub Categorie deleted successfully");
                router.push('/categories')
                router.refresh();
            } else {
                toast.error(result.error || result.message || JSON.stringify(result));
            }
        } catch (error: any) {
            toast.error(`Error: ${error.message}`);
        } finally {
            setLoading(false);
            setOpen(false)
        }
    }


    return (
        <>
            <AlertModal
                isOpen={open}
                onClose={() => setOpen(false)}
                onConfirm={onDelete}
                loading={loading}
            />
            <DropdownMenu>
                <DropdownMenuTrigger asChild>
                    <Button variant="ghost" className="h-8 w-8 p-0">
                        <span className="sr-only">Open menu</span>
                        <MoreHorizontal className="h-4 w-4" />
                    </Button>
                </DropdownMenuTrigger>
                <DropdownMenuContent align="end">
                    <DropdownMenuLabel>
                        Actions
                    </DropdownMenuLabel>
                    <DropdownMenuItem onClick={() => onCopy(data.id.toString())}>
                        <Copy className="mr-2 h-4 w-4" />
                        Copy ID
                    </DropdownMenuItem>
                    <DropdownMenuItem onClick={() => router.push(`/sub-categories/${data.id}`)}>
                        <Edit className="mr-2 h-4 w-4" />
                        Update
                    </DropdownMenuItem>
                    <DropdownMenuItem onClick={() => setOpen(true)}>
                        <Trash className="mr-2 h-4 w-4" />
                        Delete
                    </DropdownMenuItem>
                </DropdownMenuContent>
            </DropdownMenu>
        </>
    )
}