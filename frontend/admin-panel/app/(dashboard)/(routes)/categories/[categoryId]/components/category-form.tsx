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
import toast from "react-hot-toast";
import { AlertModal } from "@/components/modals/alert-modal";
import { useRouter, useParams } from "next/navigation";
import { Category } from "@/app/schema/schema";


const formSchema = z.object({
    name: z.string().min(1),
    description: z.string().min(1),
    pictureUrl: z.string().url().optional()
});

type CategoryFormValues = z.infer<typeof formSchema>;

interface CategoryFormProps {
    initialData: Category | null;
}

export const CategoryForm: React.FC<CategoryFormProps> = ({ initialData }) => {
    const params = useParams();
    const router = useRouter();



    const [open, setOpen] = useState(false);
    const [loading, setLoading] = useState(false);

    const title = initialData ? "Edit Category" : "Add a Category";
    const description = initialData ? "Update the category details" : "Add a new category";
    const toastMessage = initialData ? "Category updated successfully" : "Category added successfully";
    const action = initialData ? "Save changes" : "Add Category";

    const formMethods = useForm<CategoryFormValues>({
        resolver: zodResolver(formSchema),
        defaultValues: {
            name: initialData?.ic_name || "",
            description: initialData?.ic_description || "",
            pictureUrl: initialData?.ic_picture_url || ""
        }
    });

    const onSubmit = async (data: CategoryFormValues) => {


        const formattedData = {
            ic_name: data.name,
            ic_description: data.description,
            ic_picture_url: data.pictureUrl
        };

        try {
            setLoading(true);
            let response;
            let result;
            if (initialData) {
                response = await fetch(`/api/categories?id=${initialData?.ic_id}`, {
                    method: 'PUT',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify(formattedData)
                });
            } else {
                response = await fetch('/api/categories', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify(formattedData)
                });
            }
            result = await response.json();

            if (response.status === 201 || response.status === 200) {
                toast.success("Category submitted successfully");
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
            const response = await fetch(`/api/categories?id=${params.categoryId}`, {
                method: 'DELETE',
            });

            const result = await response.json();

            if (response.status === 200) {
                toast.success("Category deleted successfully");
                router.push('/categories')
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
                                <FormLabel>Category Name</FormLabel>
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
                                <FormLabel>Category Description</FormLabel>
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
                    <Button disabled={loading} className="ml-auto" type="submit">
                        {action}
                    </Button>
                </form>
            </FormProvider>
        </>
    );
};
