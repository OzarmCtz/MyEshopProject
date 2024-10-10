"use client";

import { Heading } from "@/components/ui/heading";
import { Button } from "@/components/ui/button";
import { Trash } from "lucide-react";
import { Separator } from "@/components/ui/separator";
import { z } from "zod";
import { useForm, FormProvider } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useState } from "react";
import { FormControl, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import toast from "react-hot-toast";
import { AlertModal } from "@/components/modals/alert-modal";
import { useRouter, useParams } from "next/navigation";
import { Item, SubCategory } from "@/app/schema/schema";

const formSchema = z.object({
    title: z.string().min(1),
    description: z.string().min(1),
    price: z.preprocess((val) => Number(val), z.number().min(1)),
    quantity: z.preprocess((val) => Number(val), z.number().min(1)),
    pictureUrl: z.string().url().min(1),
    filePath: z.string().url().min(1),
    disabled: z.enum(["0", "1"]).optional(),
    subCategoryId: z.string().min(1)
});

type ItemFormValues = z.infer<typeof formSchema>;

interface ItemFormProps {
    initialData: Item | null;
    subCategory: SubCategory[];
    initialSubCategory: SubCategory | null;
}

export const ItemForm: React.FC<ItemFormProps> = ({ initialData, subCategory, initialSubCategory }) => {
    const params = useParams();
    const router = useRouter();

    const [open, setOpen] = useState(false);
    const [loading, setLoading] = useState(false);

    const title = initialData ? "Edit Item" : "Add an Item";
    const description = initialData ? "Update the item details" : "Add a new item";
    const toastMessage = initialData ? "Item updated successfully" : "Item added successfully";
    const action = initialData ? "Save changes" : "Add Item";

    const formMethods = useForm<ItemFormValues>({
        resolver: zodResolver(formSchema),
        defaultValues: {
            title: initialData?.i_title || "",
            description: initialData?.i_description || "",
            price: initialData?.i_price || 0,
            quantity: initialData?.i_quantity || 0,
            pictureUrl: initialData?.i_picture_url || "",
            filePath: initialData?.i_file_path || "",
            disabled: initialData?.i_is_disabled ? "1" : "0",
            subCategoryId: initialSubCategory?.isc_id.toString() || ""
        }
    });

    const onSubmit = async (data: ItemFormValues) => {

        const formattedData = {
            i_title: data.title,
            i_description: data.description,
            i_price: data.price.toString(),
            i_quantity: data.quantity,
            i_picture_url: data.pictureUrl,
            i_file_path: data.filePath,
            i_is_disabled: data.disabled === "1",
            iscl_sub_category_id: Number(data.subCategoryId)

        };


        const formattedDataPut = {
            i_title: data.title,
            i_description: data.description,
            i_price: data.price.toString(),
            i_quantity: data.quantity,
            i_picture_url: data.pictureUrl,
            i_file_path: data.filePath,
            i_is_disabled: data.disabled === "1",
            iscl_sub_category_id: Number(data.subCategoryId),
        };

        try {
            setLoading(true);
            let response;
            let result;
            if (initialData) {
                response = await fetch(`/api/items?id=${initialData?.i_id}`, {
                    method: 'PUT',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify(formattedDataPut)
                });
            } else {
                response = await fetch('/api/items', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify(formattedData)
                });
            }
            result = await response.json();

            if (response.status === 201 || response.status === 200) {
                toast.success("Item submitted successfully");
                router.push('/items');
                router.refresh();
            } else {
                toast.error(result.error || result.message || JSON.stringify(result));
            }
        } catch (error: any) {
            toast.error(`Error: ${error.message}`);
        } finally {
            setLoading(false);
        }
    };

    const onDelete = async () => {
        try {
            setLoading(true);
            const response = await fetch(`/api/items?id=${params.itemId}`, {
                method: 'DELETE',
            });

            const result = await response.json();

            if (response.status === 200) {
                toast.success("Item deleted successfully");
                router.push('/items');
                router.refresh();
            } else {
                toast.error(result.error || result.message || JSON.stringify(result));
            }
        } catch (error: any) {
            toast.error(`Error: ${error.message}`);
        } finally {
            setLoading(false);
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
            <div className="flex items-center justify-between">
                <Heading
                    title={title}
                    description={description}
                />
                {initialData && (
                    <Button
                        disabled={loading}
                        variant="destructive"
                        size="icon"
                        onClick={() => setOpen(true)}
                    >
                        <Trash className="h-4 w-4" />
                    </Button>
                )}
            </div>
            <Separator />
            <FormProvider {...formMethods}>
                <form onSubmit={formMethods.handleSubmit(onSubmit)} className="space-y-8 w-full">
                    <FormField
                        control={formMethods.control}
                        name="title"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Item Title</FormLabel>
                                <FormControl>
                                    <Input {...field} disabled={loading} />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={formMethods.control}
                        name="description"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Item Description</FormLabel>
                                <FormControl>
                                    <Input {...field} disabled={loading} />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={formMethods.control}
                        name="price"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Price</FormLabel>
                                <FormControl>
                                    <Input type="number" {...field} disabled={loading} />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={formMethods.control}
                        name="quantity"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Quantity</FormLabel>
                                <FormControl>
                                    <Input type="number" {...field} disabled={loading} />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={formMethods.control}
                        name="pictureUrl"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Picture URL</FormLabel>
                                <FormControl>
                                    <Input {...field} disabled={loading} />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={formMethods.control}
                        name="filePath"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>File Path</FormLabel>
                                <FormControl>
                                    <Input {...field} disabled={loading} />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={formMethods.control}
                        name="disabled"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Status</FormLabel>
                                <Select
                                    disabled={loading}
                                    onValueChange={field.onChange}
                                    value={field.value}
                                >
                                    <FormControl>
                                        <SelectTrigger>
                                            <SelectValue placeholder="Select status" />
                                        </SelectTrigger>
                                    </FormControl>
                                    <SelectContent>
                                        <SelectItem value="0">Enabled</SelectItem>
                                        <SelectItem value="1">Disabled</SelectItem>
                                    </SelectContent>
                                </Select>
                                <FormMessage />
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={formMethods.control}
                        name="subCategoryId"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Sub Category Link</FormLabel>
                                <Select
                                    disabled={loading}
                                    onValueChange={field.onChange}
                                    value={field.value}
                                    defaultValue={field.value}
                                >
                                    <FormControl>
                                        <SelectTrigger>
                                            <SelectValue placeholder="Select a category" />
                                        </SelectTrigger>
                                    </FormControl>
                                    <SelectContent>
                                        {subCategory.map((item) => (
                                            <SelectItem key={item.isc_id} value={item.isc_id.toString()}>
                                                {item.isc_name}
                                            </SelectItem>
                                        ))}
                                    </SelectContent>
                                </Select>
                                <FormMessage />
                            </FormItem>
                        )}
                    />
                    <Button disabled={loading} className="ml-auto" type="submit">
                        {action}
                    </Button>
                </form>
            </FormProvider>
        </>
    );
};
