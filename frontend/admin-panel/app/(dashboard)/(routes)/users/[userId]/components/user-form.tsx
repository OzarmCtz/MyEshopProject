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
import { User } from "@/app/schema/schema";

const formSchema = z.object({
    email: z.string().email(),
    uid: z.string(),
    registerDate: z.string(),
    disabled: z.string(),
});

type UserFormValues = z.infer<typeof formSchema>;

interface UserFormProps {
    initialData: User | null;
}

export const UserForm: React.FC<UserFormProps> = ({ initialData }) => {
    console.log("initialData");
    console.log(initialData);
    const params = useParams();
    const router = useRouter();

    const [open, setOpen] = useState(false);
    const [loading, setLoading] = useState(false);

    const title = "Update User";
    const description = "Update the user details";
    const toastMessage = "Item updated successfully";
    const action = "Save changes";

    const formMethods = useForm<UserFormValues>({
        resolver: zodResolver(formSchema),
        defaultValues: {
            email: initialData?.u_email,
            uid: initialData?.u_uid,
            disabled: initialData?.u_is_disabled ? "1" : "0",
        }
    });

    const onSubmit = async (data: UserFormValues) => {

        const formattedDataPut = {

        };

        try {
            setLoading(true);
            let response;
            let result;
            response = await fetch(`/api/users?id=${initialData?.u_id}`, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(formattedDataPut)
            });
            result = await response.json();

            if (response.status === 201 || response.status === 200) {
                toast.success(toastMessage);
                router.push('/users');
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
            const response = await fetch(`/api/users?id=${params.itemId}`, {
                method: 'DELETE',
            });

            const result = await response.json();

            if (response.status === 200) {
                toast.success("User deleted successfully");
                router.push('/users');
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
                        name="email"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>User Email</FormLabel>
                                <FormControl>
                                    <Input {...field} disabled={true} />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={formMethods.control}
                        name="uid"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>User Uid</FormLabel>
                                <FormControl>
                                    <Input {...field} disabled={true} />
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
                                        <SelectItem value="0">Active</SelectItem>
                                        <SelectItem value="1">Suspended</SelectItem>
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
