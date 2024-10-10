"use client";

import { Heading } from "@/components/ui/heading";
import { Button } from "@/components/ui/button";
import { Trash } from "lucide-react";
import { Separator } from "@/components/ui/separator";
import { z } from "zod";
import { useForm, FormProvider, Form } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useState } from "react";
import { FormControl, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import toast from "react-hot-toast";
import { AlertModal } from "@/components/modals/alert-modal";
import { useRouter, useParams } from "next/navigation";
import { SubCategory, Category } from "@/app/schema/schema";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";

const formSchema = z.object({
    name: z.string().min(1),
    description: z.string().min(1),
    pictureUrl: z.string().url().optional(),
    categoryId: z.string().min(1)
});

type SubCategoryFormValues = z.infer<typeof formSchema>;

interface SubCategoryFormProps {
    initialData: SubCategory | null;
    category: Category[];
    initialCategory: Category | null;
}

export const SubCategoryForm: React.FC<SubCategoryFormProps> = ({ initialData, category, initialCategory }) => {

    const params = useParams();
    const router = useRouter();
    const [open, setOpen] = useState(false);
    const [loading, setLoading] = useState(false);

    const title = initialData ? "Edit Sub Category" : "Add a Sub Category";
    const description = initialData ? "Update the Sub category details" : "Add a new Sub Category";
    const toastMessage = initialData ? "Sub Category updated successfully" : "Sub Category added successfully";
    const action = initialData ? "Save changes" : "Add Sub Category";

    const formMethods = useForm<SubCategoryFormValues>({
        resolver: zodResolver(formSchema),
        defaultValues: {
            name: initialData?.isc_name || "",
            description: initialData?.isc_description || "",
            pictureUrl: initialData?.isc_picture_url || "",
            categoryId: initialCategory?.ic_id.toString() || ""
        }
    });

    const onSubmit = async (data: SubCategoryFormValues) => {

        const formattedData = {
            isc_name: data.name,
            isc_description: data.description,
            isc_picture_url: data.pictureUrl,
            icl_items_category_id: Number(data.categoryId)
        };

        const putformattedData = {
            isc_name: data.name,
            isc_description: data.description,
            isc_picture_url: data.pictureUrl,
            icl_items_category_id: Number(data.categoryId),
        };


        try {
            setLoading(true);
            let response;
            let result;
            if (initialData) {
                response = await fetch(`/api/sub-categories?id=${initialData?.isc_id}`, {
                    method: 'PUT',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify(putformattedData)
                });
            } else {
                response = await fetch('/api/sub-categories', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify(formattedData)
                });
            }
            result = await response.json();

            if (response.status === 201 || response.status === 200) {
                toast.success(toastMessage);
                router.push('/categories');
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
            const response = await fetch(`/api/sub-categories?id=${initialData?.isc_id}`, {
                method: 'DELETE',
            });

            const result = await response.json();

            if (response.status === 200) {
                toast.success("Sub Category deleted successfully");
                router.push('/categories');
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
                        name="name"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Sub Category Name</FormLabel>
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
                                <FormLabel>Sub Category Description</FormLabel>
                                <FormControl>
                                    <Input {...field} disabled={loading} />
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
                        name="categoryId"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Category Link</FormLabel>
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
                                        {category.map((item) => (
                                            <SelectItem key={item.ic_id} value={item.ic_id.toString()}>
                                                {item.ic_name}
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
